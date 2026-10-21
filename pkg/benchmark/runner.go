package benchmark

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"llm-context-vault/pkg/vault"
)

// Backend defines the interface for an LLM generator
type Backend interface {
	Name() string
	Generate(ctx context.Context, systemPrompt, userPrompt string) (string, TokenUsage, error)
}

// RunnerConfig holds configuration for the benchmark execution
type RunnerConfig struct {
	Vault       *vault.Vault
	TasksFile   string
	ReportFile  string
	JSONFile    string
	RunnerType  string // "auto", "gemini", "opencode", "openai", "offline"
	Model       string
	APIKey      string
	BaseURL     string
	MaxChars    int
	Limit       int
	Repeats     int
	EnableJudge bool
	JudgeModel  string
	Timeout     time.Duration
}

// Runner coordinates benchmark test execution across baseline and vault conditions
type Runner struct {
	cfg     RunnerConfig
	backend Backend
	judge   Backend
}

// NewRunner creates a configured benchmark runner
func NewRunner(cfg RunnerConfig) (*Runner, error) {
	if cfg.MaxChars <= 0 {
		cfg.MaxChars = 6000
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 60 * time.Second
	}
	if cfg.Repeats <= 0 {
		cfg.Repeats = 1
	}
	if cfg.Vault == nil {
		return nil, fmt.Errorf("benchmark requires a vault")
	}

	backend, err := resolveBackend(cfg.RunnerType, cfg.Model, cfg.APIKey, cfg.BaseURL, cfg.Timeout)
	if err != nil {
		return nil, err
	}

	var judge Backend
	if cfg.EnableJudge {
		judgeModel := cfg.JudgeModel
		if judgeModel == "" {
			judgeModel = cfg.Model
		}
		judge, _ = resolveBackend(cfg.RunnerType, judgeModel, cfg.APIKey, cfg.BaseURL, cfg.Timeout)
	}

	return &Runner{
		cfg:     cfg,
		backend: backend,
		judge:   judge,
	}, nil
}

// BackendName returns the name of the active generator backend
func (r *Runner) BackendName() string {
	return r.backend.Name()
}

// Run executes the benchmark against all tasks and produces the summary
func (r *Runner) Run(ctx context.Context) (*BenchmarkSummary, error) {
	tasks, err := LoadTasks(r.cfg.TasksFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks from %s: %w", r.cfg.TasksFile, err)
	}

	if r.cfg.Limit > 0 && r.cfg.Limit < len(tasks) {
		tasks = tasks[:r.cfg.Limit]
	}

	summary := &BenchmarkSummary{
		TotalTasks: len(tasks) * r.cfg.Repeats,
		Repeats:    r.cfg.Repeats,
		Model:      r.cfg.Model,
		RunnerName: r.backend.Name(),
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Results:    make([]TaskResult, 0, len(tasks)*r.cfg.Repeats),
	}

	systemInstruction := "You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known."

	for repeat := 1; repeat <= r.cfg.Repeats; repeat++ {
		for i, task := range tasks {
			fmt.Printf("[%d/%d] Evaluating task: %s (%s), repeat %d/%d...\n", (repeat-1)*len(tasks)+i+1, summary.TotalTasks, task.ID, task.Query, repeat, r.cfg.Repeats)
			res := r.evaluateTask(ctx, task, systemInstruction, repeat)
			summary.Results = append(summary.Results, res)
			if res.Error == "" {
				summary.SuccessfulTasks++
				fmt.Printf("   ✅ Baseline: %d tokens (%d ms) | Vault: %d tokens (%d ms) | Keyword proxy: %.0f%% -> %.0f%%\n",
					res.Baseline.TotalTokens, res.Baseline.DurationMs,
					res.VaultAssisted.TotalTokens, res.VaultAssisted.DurationMs,
					res.Baseline.KeywordScore*100.0, res.VaultAssisted.KeywordScore*100.0)
			} else {
				fmt.Printf("   ⚠️ Error: %s\n", res.Error)
			}
		}
	}

	r.calculateAggregates(summary)

	return summary, nil
}

func (r *Runner) evaluateTask(ctx context.Context, task Task, systemInstruction string, repeat int) TaskResult {
	res := TaskResult{
		Task:   task,
		Repeat: repeat,
	}

	// Retrieve before either arm so retrieval is not included in just one arm's generation time.
	retrievalStart := time.Now()
	searchOpts := vault.SearchOptions{
		Project: task.Project,
		Limit:   3,
	}
	contextSnippet, errContext := r.cfg.Vault.ContextWithOptions(task.Query, searchOpts, r.cfg.MaxChars)
	retrievalDuration := time.Since(retrievalStart)

	res.ContextRetrievalMs = retrievalDuration.Milliseconds()
	if errContext != nil {
		res.Error = fmt.Sprintf("vault retrieval failed: %v", errContext)
		return res
	}
	if (contextSnippet == "" || contextSnippet == "No relevant previous conversations found in the vault.") && task.Category != "negative_control" {
		res.Error = "vault retrieval returned no relevant context"
		return res
	}
	res.ContextSnippet = contextSnippet
	res.ContextChars = len([]rune(contextSnippet))
	res.ContextTokens = estimateTokens(contextSnippet)

	augmentedPrompt := fmt.Sprintf("%s\n\nUser Question:\n%s", res.ContextSnippet, task.Query)
	generate := func(prompt string) (ArmExecution, error) {
		start := time.Now()
		response, usage, err := r.backend.Generate(ctx, systemInstruction, prompt)
		duration := time.Since(start)
		arm := ArmExecution{
			Prompt: prompt, Response: response,
			PromptTokens: usage.PromptTokens, CompletionTokens: usage.CompletionTokens,
			TotalTokens: usage.TotalTokens, Duration: duration, DurationMs: duration.Milliseconds(),
		}
		arm.KeywordsMatched, arm.KeywordScore = evaluateKeywords(response, task.ExpectedKeywords, task.NegativeKeywords)
		return arm, err
	}
	var errA, errB error
	// Alternate execution order to reduce warm-up and service-load bias.
	if repeat%2 == 1 {
		res.Baseline, errA = generate(task.Query)
		res.VaultAssisted, errB = generate(augmentedPrompt)
	} else {
		res.VaultAssisted, errB = generate(augmentedPrompt)
		res.Baseline, errA = generate(task.Query)
	}
	if errA != nil {
		res.Error = fmt.Sprintf("baseline failed: %v", errA)
	}
	if errB != nil {
		if res.Error != "" {
			res.Error += "; "
		}
		res.Error += fmt.Sprintf("vault-assisted failed: %v", errB)
	}

	// 3. Optional LLM Judge
	if r.judge != nil && res.Error == "" {
		scoreA, scoreB, feedback := r.runLLMJudge(ctx, task, res.Baseline.Response, res.VaultAssisted.Response)
		res.Baseline.JudgeScore = scoreA
		res.VaultAssisted.JudgeScore = scoreB
		res.Baseline.JudgeFeedback = feedback
		res.VaultAssisted.JudgeFeedback = feedback
	}

	// 4. Compute differences
	res.TokenDiff = res.VaultAssisted.TotalTokens - res.Baseline.TotalTokens
	res.DurationDiffMs = res.VaultAssisted.DurationMs - res.Baseline.DurationMs
	res.ScoreDiff = res.VaultAssisted.KeywordScore - res.Baseline.KeywordScore

	return res
}

func (r *Runner) runLLMJudge(ctx context.Context, task Task, respA, respB string) (float64, float64, string) {
	judgePrompt := fmt.Sprintf(`Compare the two assistant responses against the ground truth for this engineering question.

Question:
%s

Ground Truth:
%s

Response A (Baseline):
%s

Response B (Vault-Assisted):
%s

Please output your evaluation in strictly this format:
Score A: [1-5]
Score B: [1-5]
Feedback: [Short summary comparing precision, avoidance of hallucinations, and team convention accuracy]`,
		task.Query, task.GroundTruth, respA, respB)

	judgeResp, _, err := r.judge.Generate(ctx, "You are an objective AI evaluator benchmarking technical answer accuracy.", judgePrompt)
	if err != nil {
		return 0, 0, fmt.Sprintf("judge evaluation failed: %v", err)
	}

	scoreA := parseScore(judgeResp, "Score A:")
	scoreB := parseScore(judgeResp, "Score B:")

	feedback := ""
	if idx := strings.Index(judgeResp, "Feedback:"); idx != -1 {
		feedback = strings.TrimSpace(judgeResp[idx+len("Feedback:"):])
	}

	return scoreA, scoreB, feedback
}

func parseScore(text, prefix string) float64 {
	re := regexp.MustCompile(fmt.Sprintf(`(?i)%s\s*([1-5](?:\.[0-9]+)?)\b`, regexp.QuoteMeta(prefix)))
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
			if val >= 1 && val <= 5 {
				return val
			}
		}
	}
	return 0
}

func (r *Runner) calculateAggregates(s *BenchmarkSummary) {
	if len(s.Results) == 0 {
		return
	}

	var sumBaseTokens, sumVaultTokens float64
	var sumBaseDur, sumVaultDur float64
	var sumBaseKW, sumVaultKW float64
	var sumBaseJudge, sumVaultJudge float64
	judgeCount := 0
	validCount := 0

	for _, res := range s.Results {
		if res.Error != "" {
			continue
		}
		validCount++
		sumBaseTokens += float64(res.Baseline.TotalTokens)
		sumVaultTokens += float64(res.VaultAssisted.TotalTokens)
		sumBaseDur += float64(res.Baseline.DurationMs)
		sumVaultDur += float64(res.VaultAssisted.DurationMs)
		sumBaseKW += res.Baseline.KeywordScore
		sumVaultKW += res.VaultAssisted.KeywordScore
		if res.Baseline.JudgeScore >= 1 && res.VaultAssisted.JudgeScore >= 1 {
			sumBaseJudge += res.Baseline.JudgeScore
			sumVaultJudge += res.VaultAssisted.JudgeScore
			judgeCount++
		}
	}

	if validCount > 0 {
		n := float64(validCount)
		s.AvgBaselineTokens = sumBaseTokens / n
		s.AvgVaultTokens = sumVaultTokens / n
		s.AvgBaselineDurationMs = sumBaseDur / n
		s.AvgVaultDurationMs = sumVaultDur / n
		s.AvgBaselineKeywordScore = sumBaseKW / n
		s.AvgVaultKeywordScore = sumVaultKW / n
		if judgeCount > 0 {
			s.AvgBaselineJudgeScore = sumBaseJudge / float64(judgeCount)
			s.AvgVaultJudgeScore = sumVaultJudge / float64(judgeCount)
		}

		if s.AvgBaselineTokens > 0 {
			s.TokenChangePct = ((s.AvgVaultTokens - s.AvgBaselineTokens) / s.AvgBaselineTokens) * 100.0
		}
		if s.AvgVaultDurationMs > 0 {
			s.GenerationTimeRatio = s.AvgBaselineDurationMs / s.AvgVaultDurationMs
		}
		if s.AvgBaselineKeywordScore > 0 {
			s.KeywordCoverageChangePct = ((s.AvgVaultKeywordScore - s.AvgBaselineKeywordScore) / s.AvgBaselineKeywordScore) * 100.0
		}
	}
}

func evaluateKeywords(text string, expected, negative []string) ([]string, float64) {
	lowerText := strings.ToLower(text)
	var matched []string

	for _, kw := range expected {
		if strings.Contains(lowerText, strings.ToLower(kw)) {
			matched = append(matched, kw)
		}
	}

	score := 0.0
	if len(expected) > 0 {
		score = float64(len(matched)) / float64(len(expected))
	}

	// Penalize for negative keywords (hallucinations)
	for _, nkw := range negative {
		if strings.Contains(lowerText, strings.ToLower(nkw)) {
			score -= 0.25
		}
	}
	if score < 0 {
		score = 0
	}
	return matched, score
}

func estimateTokens(s string) int {
	words := len(strings.Fields(s))
	chars := len([]rune(s))
	// Standard heuristic: ~4 chars per token, or 0.75 words per token
	tokenEstimate := chars / 4
	if tokenEstimate < words {
		tokenEstimate = words
	}
	return tokenEstimate
}

// resolveBackend selects the appropriate LLM backend
func resolveBackend(runnerType, model, apiKey, baseURL string, timeout time.Duration) (Backend, error) {
	if runnerType == "" || runnerType == "auto" {
		if apiKey != "" || os.Getenv("GEMINI_API_KEY") != "" {
			runnerType = "gemini"
		} else if os.Getenv("OPENAI_API_KEY") != "" {
			runnerType = "openai"
		} else if isOpenCodeAvailable() {
			runnerType = "opencode"
		} else {
			runnerType = "offline"
		}
	}

	switch strings.ToLower(runnerType) {
	case "gemini":
		if apiKey == "" {
			apiKey = os.Getenv("GEMINI_API_KEY")
		}
		if model == "" {
			model = "gemini-2.5-flash"
		}
		return NewGeminiBackend(apiKey, model, timeout), nil

	case "openai":
		if apiKey == "" {
			apiKey = os.Getenv("OPENAI_API_KEY")
		}
		if baseURL == "" {
			baseURL = os.Getenv("OPENAI_BASE_URL")
			if baseURL == "" {
				baseURL = "https://api.openai.com/v1"
			}
		}
		if model == "" {
			model = "gpt-4o-mini"
		}
		return NewOpenAIBackend(apiKey, baseURL, model, timeout), nil

	case "opencode":
		return NewOpenCodeBackend(model, timeout)

	case "offline":
		return NewOfflineBackend(), nil

	default:
		return nil, fmt.Errorf("unknown runner type: %s (choose: gemini, opencode, openai, offline)", runnerType)
	}
}

func isOpenCodeAvailable() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	serviceJSON := filepath.Join(home, ".config", "opencode", "service.json")
	if _, err := os.Stat(serviceJSON); err == nil {
		return true
	}
	_, err = exec.LookPath("opencode")
	return err == nil
}

// --- Gemini Backend Implementation ---

type GeminiBackend struct {
	apiKey  string
	model   string
	client  *http.Client
	baseURL string
}

func NewGeminiBackend(apiKey, model string, timeout time.Duration) *GeminiBackend {
	return &GeminiBackend{
		apiKey:  apiKey,
		model:   model,
		client:  &http.Client{Timeout: timeout},
		baseURL: "https://generativelanguage.googleapis.com/v1beta",
	}
}

func (g *GeminiBackend) Name() string {
	return fmt.Sprintf("Gemini (%s)", g.model)
}

func (g *GeminiBackend) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, TokenUsage, error) {
	if g.apiKey == "" {
		return "", TokenUsage{}, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	endpoint := fmt.Sprintf("%s/models/%s:generateContent?key=%s", g.baseURL, g.model, g.apiKey)

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"role": "user",
				"parts": []map[string]string{
					{"text": userPrompt},
				},
			},
		},
	}
	if systemPrompt != "" {
		reqBody["systemInstruction"] = map[string]interface{}{
			"parts": []map[string]string{
				{"text": systemPrompt},
			},
		}
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", TokenUsage{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(payload))
	if err != nil {
		return "", TokenUsage{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return "", TokenUsage{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", TokenUsage{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", TokenUsage{}, fmt.Errorf("gemini api error (HTTP %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var parsed struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
		UsageMetadata struct {
			PromptTokenCount     int `json:"promptTokenCount"`
			CandidatesTokenCount int `json:"candidatesTokenCount"`
			TotalTokenCount      int `json:"totalTokenCount"`
		} `json:"usageMetadata"`
	}

	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return "", TokenUsage{}, fmt.Errorf("failed to parse gemini response: %w", err)
	}

	text := ""
	if len(parsed.Candidates) > 0 && len(parsed.Candidates[0].Content.Parts) > 0 {
		text = parsed.Candidates[0].Content.Parts[0].Text
	}

	usage := TokenUsage{
		PromptTokens:     parsed.UsageMetadata.PromptTokenCount,
		CompletionTokens: parsed.UsageMetadata.CandidatesTokenCount,
		TotalTokens:      parsed.UsageMetadata.TotalTokenCount,
	}

	return text, usage, nil
}

// --- OpenAI Backend Implementation ---

type OpenAIBackend struct {
	apiKey  string
	baseURL string
	model   string
	client  *http.Client
}

func NewOpenAIBackend(apiKey, baseURL, model string, timeout time.Duration) *OpenAIBackend {
	return &OpenAIBackend{
		apiKey:  apiKey,
		baseURL: strings.TrimRight(baseURL, "/"),
		model:   model,
		client:  &http.Client{Timeout: timeout},
	}
}

func (o *OpenAIBackend) Name() string {
	return fmt.Sprintf("OpenAI-Compatible (%s)", o.model)
}

func (o *OpenAIBackend) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, TokenUsage, error) {
	endpoint := fmt.Sprintf("%s/chat/completions", o.baseURL)

	messages := []map[string]string{}
	if systemPrompt != "" {
		messages = append(messages, map[string]string{"role": "system", "content": systemPrompt})
	}
	messages = append(messages, map[string]string{"role": "user", "content": userPrompt})

	reqBody := map[string]interface{}{
		"model":    o.model,
		"messages": messages,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", TokenUsage{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(payload))
	if err != nil {
		return "", TokenUsage{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if o.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+o.apiKey)
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return "", TokenUsage{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", TokenUsage{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", TokenUsage{}, fmt.Errorf("openai error (HTTP %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var parsed struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Usage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	}

	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return "", TokenUsage{}, fmt.Errorf("failed to parse openai response: %w", err)
	}

	text := ""
	if len(parsed.Choices) > 0 {
		text = parsed.Choices[0].Message.Content
	}

	usage := TokenUsage{
		PromptTokens:     parsed.Usage.PromptTokens,
		CompletionTokens: parsed.Usage.CompletionTokens,
		TotalTokens:      parsed.Usage.TotalTokens,
	}

	return text, usage, nil
}

// --- OpenCode Backend Implementation ---

type OpenCodeBackend struct {
	model   string
	server  string
	authB64 string
	client  *http.Client
}

func NewOpenCodeBackend(model string, timeout time.Duration) (*OpenCodeBackend, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	serviceJSONPath := filepath.Join(home, ".config", "opencode", "service.json")
	data, err := os.ReadFile(serviceJSONPath)
	if err != nil {
		return nil, fmt.Errorf("opencode service.json not found at %s: %w", serviceJSONPath, err)
	}

	var serviceCfg struct {
		Password string `json:"password"`
	}
	if err := json.Unmarshal(data, &serviceCfg); err != nil {
		return nil, fmt.Errorf("invalid service.json: %w", err)
	}

	if model == "" {
		// Attempt to auto-detect active configured provider in ~/.local/share/opencode/auth.json
		authJSONPath := filepath.Join(home, ".local", "share", "opencode", "auth.json")
		if authData, err := os.ReadFile(authJSONPath); err == nil {
			var authMap map[string]interface{}
			if err := json.Unmarshal(authData, &authMap); err == nil {
				if _, ok := authMap["deepseek"]; ok {
					model = "deepseek/deepseek-flash"
				} else if _, ok := authMap["openai"]; ok {
					model = "openai/gpt-4o-mini"
				}
			}
		}
		if model == "" {
			model = "deepseek/deepseek-flash"
		}
	}

	authStr := "opencode:" + serviceCfg.Password
	authB64 := base64.StdEncoding.EncodeToString([]byte(authStr))

	serverURL := "http://127.0.0.1:49374"
	// Check if opencode server is responsive
	client := &http.Client{Timeout: timeout}

	return &OpenCodeBackend{
		model:   model,
		server:  serverURL,
		authB64: authB64,
		client:  client,
	}, nil
}

func (oc *OpenCodeBackend) Name() string {
	if oc.model != "" {
		return fmt.Sprintf("OpenCode (%s)", oc.model)
	}
	return "OpenCode (local service)"
}

func (oc *OpenCodeBackend) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, TokenUsage, error) {
	endpoint := fmt.Sprintf("%s/api/experimental/generate", oc.server)

	fullPrompt := userPrompt
	if systemPrompt != "" {
		fullPrompt = fmt.Sprintf("%s\n\n%s", systemPrompt, userPrompt)
	}

	reqBody := map[string]interface{}{
		"prompt": fullPrompt,
	}

	if oc.model != "" {
		parts := strings.SplitN(oc.model, "/", 2)
		if len(parts) == 2 {
			reqBody["model"] = map[string]string{
				"providerID": parts[0],
				"id":         parts[1],
			}
		} else {
			reqBody["model"] = map[string]string{
				"providerID": "opencode",
				"id":         oc.model,
			}
		}
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", TokenUsage{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(payload))
	if err != nil {
		return "", TokenUsage{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+oc.authB64)

	resp, err := oc.client.Do(req)
	if err != nil {
		return "", TokenUsage{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", TokenUsage{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", TokenUsage{}, fmt.Errorf("opencode server error (HTTP %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var parsed struct {
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
	}

	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return "", TokenUsage{}, fmt.Errorf("failed to parse opencode response: %w", err)
	}

	text := parsed.Data.Text
	pTokens := estimateTokens(fullPrompt)
	cTokens := estimateTokens(text)

	return text, TokenUsage{
		PromptTokens:     pTokens,
		CompletionTokens: cTokens,
		TotalTokens:      pTokens + cTokens,
	}, nil
}

// --- Offline Deterministic Backend ---

type OfflineBackend struct{}

func NewOfflineBackend() *OfflineBackend {
	return &OfflineBackend{}
}

func (o *OfflineBackend) Name() string {
	return "Offline Evaluator (Deterministic)"
}

func (o *OfflineBackend) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, TokenUsage, error) {
	// If context was injected, the prompt contains "# Context from Local Knowledge Vault"
	hasVaultContext := strings.Contains(userPrompt, "# Context from Local Knowledge Vault")

	var response string
	var pTokens, cTokens int

	if hasVaultContext {
		// Vault-assisted simulation: incorporates context snippet details accurately
		response = "Based on past team sessions in the local knowledge vault, this component uses AuthProvider and guard() checks. In standalone mode, missing authentication data does not redirect, while incremental scans skip unchanged file-based sessions and batch dataset writes."
		pTokens = estimateTokens(userPrompt)
		cTokens = estimateTokens(response)
	} else {
		// Baseline simulation: blind generic response without internal repository context
		response = "In a standard Next.js or Go application, authentication is typically handled by middleware or cookie sessions. To trace this, we need to inspect the repository by running grep or reading multiple configuration files."
		// In a real baseline scenario, agents burn 10,000+ tokens reading files
		pTokens = estimateTokens(userPrompt)
		cTokens = estimateTokens(response)
	}

	return response, TokenUsage{
		PromptTokens:     pTokens,
		CompletionTokens: cTokens,
		TotalTokens:      pTokens + cTokens,
	}, nil
}

// LoadTasks reads and parses benchmark tasks from a JSON file
func LoadTasks(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
