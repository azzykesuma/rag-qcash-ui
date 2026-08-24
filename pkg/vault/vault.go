package vault

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"llm-context-vault/pkg/exporter"
	"llm-context-vault/pkg/extractor"
	"llm-context-vault/pkg/models"
	"llm-context-vault/pkg/sanitizer"
)

// SearchResult contains matching conversation turns
type SearchResult struct {
	ConversationID string
	Title          string
	SourceTool     string
	Role           string
	Snippet        string
	Path           string
}

// ToolScanReport represents the result of scanning a specific assistant
type ToolScanReport struct {
	ToolName        string
	Location        string
	DiscoveredCount int
	ImportedCount   int
	SkippedCount    int
	Warnings        []string
}

// Vault manages conversation storage, sanitization, indexing, and export
type Vault struct {
	BaseDir           string
	Sanitizer         *sanitizer.Sanitizer
	extractors        []extractor.Extractor
	agyExtractor      *extractor.AGYExtractor
	codexExtractor    *extractor.CodexExtractor
	opencodeExtractor *extractor.OpenCodeExtractor
	aiderExtractor    *extractor.AiderExtractor
	genericExtractor  *extractor.GenericExtractor
	mdExport          *exporter.MarkdownExporter
	jsonExport        *exporter.JSONLExporter
}

func New(baseDir string, s *sanitizer.Sanitizer) *Vault {
	agy := extractor.NewAGYExtractor()
	codex := extractor.NewCodexExtractor()
	opencode := extractor.NewOpenCodeExtractor()
	aider := extractor.NewAiderExtractor()
	generic := extractor.NewGenericExtractor()

	return &Vault{
		BaseDir:           baseDir,
		Sanitizer:         s,
		agyExtractor:      agy,
		codexExtractor:    codex,
		opencodeExtractor: opencode,
		aiderExtractor:    aider,
		genericExtractor:  generic,
		extractors: []extractor.Extractor{
			agy,
			codex,
			opencode,
			aider,
			generic,
		},
		mdExport:   exporter.NewMarkdownExporter(),
		jsonExport: exporter.NewJSONLExporter(),
	}
}

// IsTrivialConversation checks if a conversation is just a greeting, ping, or lacks substantive interaction
func IsTrivialConversation(conv *models.Conversation) bool {
	if conv == nil || len(conv.Messages) == 0 {
		return true
	}

	var firstUser string
	for _, m := range conv.Messages {
		if strings.ToLower(m.Role) == "user" {
			firstUser = strings.TrimSpace(m.Content)
			break
		}
	}

	if firstUser == "" {
		return true
	}

	// Regex to match greetings, pings, single-word tests in English and Indonesian
	greetingPattern := regexp.MustCompile(`(?i)^(hi|hello|hey|halo|hai|test|tes|testing|ping|p|yo|sup|morning|good morning|selamat pagi|selamat siang|selamat sore|selamat malam|apa kabar|howdy|start|menu|log)[!.,?\s]*$`)

	hasToolCalls := false
	totalLength := 0
	hasAssistantResponse := false

	for _, m := range conv.Messages {
		if len(m.ToolCalls) > 0 {
			hasToolCalls = true
		}
		if strings.ToLower(m.Role) == "assistant" && strings.TrimSpace(m.Content) != "" {
			hasAssistantResponse = true
		}
		totalLength += len(m.Content)
	}

	// 1. If no assistant response and no tool calls, it's an aborted/empty session
	if !hasAssistantResponse && !hasToolCalls {
		return true
	}

	// 2. If first message is a greeting and there are no tool calls with minimal exchange (< 300 chars)
	if greetingPattern.MatchString(firstUser) && !hasToolCalls && totalLength < 300 {
		return true
	}

	// 3. If single turn with ultra short prompt (< 10 chars) and no code/tools
	if len(conv.Messages) <= 2 && len(firstUser) < 10 && !hasToolCalls && totalLength < 200 {
		return true
	}

	return false
}

// ProcessAndStore extracts, sanitizes, audits, and persists a conversation
func (v *Vault) ProcessAndStore(sourcePath string, explicitTool string) (*models.Conversation, []string, error) {
	var matched extractor.Extractor

	for _, ext := range v.extractors {
		if explicitTool != "" && strings.EqualFold(ext.Name(), explicitTool) {
			matched = ext
			break
		} else if explicitTool == "" && ext.CanHandle(sourcePath) {
			matched = ext
			break
		}
	}

	if matched == nil {
		return nil, nil, fmt.Errorf("no suitable extractor found for path: %s", sourcePath)
	}

	rawConv, err := matched.Extract(sourcePath)
	if err != nil {
		return nil, nil, fmt.Errorf("extraction error: %w", err)
	}

	return v.StoreConversation(rawConv)
}

// ProcessAndStoreAll imports every record from generic JSON/JSONL input.
func (v *Vault) ProcessAndStoreAll(sourcePath string, explicitTool string) ([]*models.Conversation, []string, int, error) {
	if explicitTool != "" && !strings.EqualFold(explicitTool, "generic") {
		conversation, warnings, err := v.ProcessAndStore(sourcePath, explicitTool)
		if conversation == nil {
			return nil, warnings, 1, err
		}
		return []*models.Conversation{conversation}, warnings, 0, err
	}

	conversations, err := v.genericExtractor.ExtractAll(sourcePath)
	if err != nil {
		return nil, nil, 0, err
	}
	var stored []*models.Conversation
	var warnings []string
	skipped := 0
	for _, conversation := range conversations {
		cleanConversation, auditWarnings, err := v.StoreConversation(conversation)
		if err != nil {
			return stored, warnings, skipped, fmt.Errorf("failed to store %s: %w", conversation.ID, err)
		}
		if cleanConversation == nil {
			skipped++
			continue
		}
		stored = append(stored, cleanConversation)
		warnings = append(warnings, auditWarnings...)
	}
	return stored, warnings, skipped, nil
}

// StoreConversation filters trivial sessions, sanitizes, audits, and persists a stable export.
func (v *Vault) StoreConversation(rawConv *models.Conversation) (*models.Conversation, []string, error) {
	if rawConv == nil {
		return nil, nil, nil
	}

	// Filter out greeting-only or trivial sessions
	if IsTrivialConversation(rawConv) {
		return nil, nil, nil // Silently skip trivial sessions
	}

	cleanConv := v.Sanitizer.SanitizeConversation(rawConv)
	cleanConv.SchemaVersion = 1
	cleanConv.ID = stableConversationID(cleanConv)

	var auditWarnings []string
	if serialized, err := json.Marshal(cleanConv); err == nil {
		auditWarnings = v.Sanitizer.AuditText(string(serialized))
	}

	// Generate human-readable filename from first chat/query
	fileSlug := fmt.Sprintf("%s-%s", generateConversationSlug(cleanConv), cleanConv.ID[len(cleanConv.ID)-12:])

	mdDir := filepath.Join(v.BaseDir, "conversations", "markdown")
	shareDir := filepath.Join(v.BaseDir, "conversations", "sharegpt")
	datasetPath := filepath.Join(v.BaseDir, "conversations", "dataset.jsonl")

	// 1. Export Markdown with human-readable name
	mdPath := filepath.Join(mdDir, fmt.Sprintf("%s.md", fileSlug))
	if err := v.mdExport.Export(cleanConv, mdPath); err != nil {
		return nil, nil, fmt.Errorf("failed to write markdown: %w", err)
	}

	// 2. Export ShareGPT
	sharePath := filepath.Join(shareDir, fmt.Sprintf("%s.json", fileSlug))
	if err := v.jsonExport.ExportShareGPT(cleanConv, sharePath); err != nil {
		return nil, nil, fmt.Errorf("failed to write sharegpt format: %w", err)
	}

	// 3. Upsert master dataset JSONL by stable content ID.
	if err := v.jsonExport.UpsertJSONL(cleanConv, datasetPath); err != nil {
		return nil, nil, fmt.Errorf("failed to update dataset.jsonl: %w", err)
	}

	return cleanConv, auditWarnings, nil
}

func stableConversationID(conv *models.Conversation) string {
	// Exclude volatile source IDs and timestamps so re-importing unchanged content is idempotent.
	payload := struct {
		SourceTool  string            `json:"source_tool"`
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Languages   []string          `json:"languages"`
		Tags        []string          `json:"tags"`
		Messages    []models.Message  `json:"messages"`
		Metadata    map[string]string `json:"metadata"`
	}{
		SourceTool:  conv.SourceTool,
		Title:       conv.Title,
		Description: conv.Description,
		Languages:   conv.Languages,
		Tags:        conv.Tags,
		Messages:    conv.Messages,
		Metadata:    conv.Metadata,
	}
	data, _ := json.Marshal(payload)
	hash := sha256.Sum256(data)
	prefix := toKebabCase(conv.SourceTool, 20)
	if prefix == "" {
		prefix = "llm"
	}
	return fmt.Sprintf("%s-%s", prefix, hex.EncodeToString(hash[:]))
}

// generateConversationSlug builds a clean, readable filename slug from the first user message
func generateConversationSlug(conv *models.Conversation) string {
	var queryText string

	for _, m := range conv.Messages {
		if strings.ToLower(m.Role) == "user" && strings.TrimSpace(m.Content) != "" {
			queryText = m.Content
			break
		}
	}

	if queryText == "" {
		queryText = conv.Title
	}

	// Clean XML tags, markdown markers, code fences
	queryText = cleanTextForSlug(queryText)

	slug := toKebabCase(queryText, 55)
	if slug == "" {
		slug = "session"
	}

	// Format: <tool>_<slug>
	toolPrefix := strings.ToLower(conv.SourceTool)
	if toolPrefix == "" {
		toolPrefix = "llm"
	}

	return fmt.Sprintf("%s_%s", toolPrefix, slug)
}

func cleanTextForSlug(s string) string {
	// Remove XML tags like <USER_REQUEST>
	xmlTagRegex := regexp.MustCompile(`<[^>]+>`)
	s = xmlTagRegex.ReplaceAllString(s, " ")

	// Remove markdown headers and quotes
	s = strings.ReplaceAll(s, "#", " ")
	s = strings.ReplaceAll(s, ">", " ")
	s = strings.ReplaceAll(s, "`", " ")
	s = strings.ReplaceAll(s, "*", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")

	return strings.TrimSpace(s)
}

func toKebabCase(s string, maxLength int) string {
	var words []string
	var current strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			current.WriteRune(unicode.ToLower(r))
		} else if current.Len() > 0 {
			w := current.String()
			// Filter out stop words for cleaner filenames
			if !isStopWord(w) {
				words = append(words, w)
			}
			current.Reset()
		}
	}
	if current.Len() > 0 {
		w := current.String()
		if !isStopWord(w) {
			words = append(words, w)
		}
	}

	result := strings.Join(words, "-")
	if len(result) > maxLength {
		result = result[:maxLength]
		// Avoid trailing hyphen
		result = strings.TrimRight(result, "-")
	}

	return result
}

func isStopWord(w string) bool {
	switch w {
	case "a", "an", "the", "and", "or", "is", "are", "i", "my", "to", "in", "of", "it":
		return true
	default:
		return false
	}
}

// ScanAll automatically discovers and scans all installed local AI coding tools
func (v *Vault) ScanAll() ([]ToolScanReport, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to determine user home directory: %w", err)
	}

	var reports []ToolScanReport

	// 1. Scan Antigravity (AGY)
	agyBrain := filepath.Join(homeDir, ".gemini", "antigravity-cli", "brain")
	if _, err := os.Stat(agyBrain); err == nil {
		imported, skipped, warnings, err := v.ScanAGYBrain(agyBrain)
		rep := ToolScanReport{
			ToolName:        "Antigravity (AGY)",
			Location:        agyBrain,
			ImportedCount:   imported,
			SkippedCount:    skipped,
			DiscoveredCount: imported + skipped,
			Warnings:        warnings,
		}
		if err != nil {
			rep.Warnings = append(rep.Warnings, err.Error())
		}
		reports = append(reports, rep)
	}

	// 2. Scan OpenCode
	opencodePaths := []string{
		filepath.Join(homeDir, ".local", "share", "opencode", "opencode.db"),
		filepath.Join(homeDir, ".opencode", "opencode.db"),
		filepath.Join(homeDir, ".config", "opencode", "opencode.db"),
	}
	for _, p := range opencodePaths {
		if _, err := os.Stat(p); err == nil {
			imported, skipped, warnings, err := v.ScanOpenCodeDB(p)
			rep := ToolScanReport{
				ToolName:        "OpenCode",
				Location:        p,
				ImportedCount:   imported,
				SkippedCount:    skipped,
				DiscoveredCount: imported + skipped,
				Warnings:        warnings,
			}
			if err != nil {
				rep.Warnings = append(rep.Warnings, err.Error())
			}
			reports = append(reports, rep)
			break
		}
	}

	// 3. Scan OpenAI Codex
	codexSessionsDir := filepath.Join(homeDir, ".codex", "sessions")
	if _, err := os.Stat(codexSessionsDir); err == nil {
		imported, skipped, warnings, err := v.ScanCodexSessions(codexSessionsDir)
		rep := ToolScanReport{
			ToolName:        "Codex",
			Location:        codexSessionsDir,
			ImportedCount:   imported,
			SkippedCount:    skipped,
			DiscoveredCount: imported + skipped,
			Warnings:        warnings,
		}
		if err != nil {
			rep.Warnings = append(rep.Warnings, err.Error())
		}
		reports = append(reports, rep)
	}

	// 4. Scan Aider chat history in home or project
	aiderPaths := []string{
		filepath.Join(homeDir, ".aider.chat.history.md"),
		filepath.Join(".", ".aider.chat.history.md"),
	}
	for _, ap := range aiderPaths {
		if _, err := os.Stat(ap); err == nil {
			conv, warnings, err := v.ProcessAndStore(ap, "aider")
			rep := ToolScanReport{
				ToolName: "Aider",
				Location: ap,
				Warnings: warnings,
			}
			if err == nil && conv != nil {
				rep.DiscoveredCount = 1
				rep.ImportedCount = 1
			} else if conv == nil {
				rep.DiscoveredCount = 1
				rep.SkippedCount = 1
			} else if err != nil {
				rep.Warnings = append(rep.Warnings, err.Error())
			}
			reports = append(reports, rep)
		}
	}

	return reports, nil
}

// ScanAGYBrain scans the user's Antigravity brain folder
func (v *Vault) ScanAGYBrain(brainDir string) (int, int, []string, error) {
	entries, err := os.ReadDir(brainDir)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("failed to read brain dir %s: %w", brainDir, err)
	}

	importedCount := 0
	skippedCount := 0
	var allWarnings []string

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		convDir := filepath.Join(brainDir, entry.Name())
		transcript := filepath.Join(convDir, ".system_generated", "logs", "transcript.jsonl")
		transcriptFull := filepath.Join(convDir, ".system_generated", "logs", "transcript_full.jsonl")

		var targetFile string
		if _, err := os.Stat(transcriptFull); err == nil {
			targetFile = transcriptFull
		} else if _, err := os.Stat(transcript); err == nil {
			targetFile = transcript
		}

		if targetFile != "" {
			conv, warnings, err := v.ProcessAndStore(targetFile, "antigravity")
			if err == nil {
				if conv != nil {
					importedCount++
					if len(warnings) > 0 {
						allWarnings = append(allWarnings, fmt.Sprintf("[%s] %s", conv.ID, strings.Join(warnings, ", ")))
					}
				} else {
					skippedCount++
				}
			} else {
				allWarnings = append(allWarnings, fmt.Sprintf("[%s] %v", targetFile, err))
			}
		}
	}

	return importedCount, skippedCount, allWarnings, nil
}

// ScanOpenCodeDB scans and extracts all conversations from OpenCode SQLite DB
func (v *Vault) ScanOpenCodeDB(dbPath string) (int, int, []string, error) {
	convs, err := v.opencodeExtractor.ExtractAll(dbPath)
	if err != nil {
		return 0, 0, nil, err
	}

	importedCount := 0
	skippedCount := 0
	var allWarnings []string

	for _, conv := range convs {
		cleanConv, warnings, err := v.StoreConversation(conv)
		if err == nil {
			if cleanConv != nil {
				importedCount++
				if len(warnings) > 0 {
					allWarnings = append(allWarnings, fmt.Sprintf("[%s] %s", cleanConv.ID, strings.Join(warnings, ", ")))
				}
			} else {
				skippedCount++
			}
		} else {
			allWarnings = append(allWarnings, fmt.Sprintf("[%s] %v", conv.ID, err))
		}
	}

	return importedCount, skippedCount, allWarnings, nil
}

// ScanCodexSessions recursively scans and extracts all Codex rollout-*.jsonl session files
func (v *Vault) ScanCodexSessions(sessionsDir string) (int, int, []string, error) {
	var sessionFiles []string

	err := filepath.Walk(sessionsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		base := strings.ToLower(info.Name())
		if strings.HasPrefix(base, "rollout-") && strings.HasSuffix(base, ".jsonl") {
			sessionFiles = append(sessionFiles, path)
		}
		return nil
	})

	if err != nil {
		return 0, 0, nil, fmt.Errorf("error walking codex sessions: %w", err)
	}

	importedCount := 0
	skippedCount := 0
	var allWarnings []string

	for _, f := range sessionFiles {
		conv, warnings, err := v.ProcessAndStore(f, "codex")
		if err == nil {
			if conv != nil {
				importedCount++
				if len(warnings) > 0 {
					allWarnings = append(allWarnings, fmt.Sprintf("[%s] %s", conv.ID, strings.Join(warnings, ", ")))
				}
			} else {
				skippedCount++
			}
		} else {
			allWarnings = append(allWarnings, fmt.Sprintf("[%s] %v", f, err))
		}
	}

	return importedCount, skippedCount, allWarnings, nil
}

// Search looks for keywords across stored markdown conversations
func (v *Vault) Search(query string) ([]SearchResult, error) {
	mdDir := filepath.Join(v.BaseDir, "conversations", "markdown")
	queryLower := strings.ToLower(query)
	var results []SearchResult

	err := filepath.Walk(mdDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".md") {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		content := string(contentBytes)
		if strings.Contains(strings.ToLower(content), queryLower) {
			baseName := strings.TrimSuffix(filepath.Base(path), ".md")
			toolName := "unknown"
			if idx := strings.Index(baseName, "_"); idx != -1 {
				toolName = baseName[:idx]
			}

			lines := strings.Split(content, "\n")
			for i, line := range lines {
				if strings.Contains(strings.ToLower(line), queryLower) {
					start := i - 2
					if start < 0 {
						start = 0
					}
					end := i + 3
					if end > len(lines) {
						end = len(lines)
					}
					snippet := strings.Join(lines[start:end], "\n")

					results = append(results, SearchResult{
						ConversationID: baseName,
						SourceTool:     toolName,
						Snippet:        snippet,
						Path:           path,
					})
					break
				}
			}
		}
		return nil
	})

	return results, err
}

// GenerateContextSnippet formats top matching conversation turns to feed into a local LLM prompt
func (v *Vault) GenerateContextSnippet(query string, maxEntries int) (string, error) {
	results, err := v.Search(query)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "No relevant previous conversations found in the vault.", nil
	}

	var b strings.Builder
	b.WriteString("# Context from Local Knowledge Vault\n\n")
	b.WriteString("The following are relevant past problem-solving sessions from developers:\n\n")

	count := 0
	for _, res := range results {
		if count >= maxEntries {
			break
		}
		b.WriteString(fmt.Sprintf("### Reference Session: `%s`\n", res.ConversationID))
		b.WriteString("```markdown\n")
		b.WriteString(res.Snippet)
		b.WriteString("\n```\n\n")
		count++
	}

	return b.String(), nil
}

// AuditAll audits all stored markdown and json files to verify zero secret leaks
func (v *Vault) AuditAll() (map[string][]string, error) {
	violations := make(map[string][]string)
	convDir := filepath.Join(v.BaseDir, "conversations")

	err := filepath.Walk(convDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		warnings := v.Sanitizer.AuditText(string(contentBytes))
		if len(warnings) > 0 {
			violations[path] = warnings
		}
		return nil
	})

	return violations, err
}
