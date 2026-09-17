package vault

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
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
	Project        string
	Date           string
	StartLine      int
	EndLine        int
	FirstTurn      int
	LastTurn       int
	Score          float64
}

// ToolScanReport represents the result of scanning a specific assistant
type ToolScanReport struct {
	ToolName        string
	Location        string
	DiscoveredCount int
	ImportedCount   int
	SkippedCount    int
	Warnings        []string
	UnchangedCount  int
	NewCount        int
	ChangedCount    int
	FailedCount     int
	Duration        time.Duration
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
	storeMu           sync.Mutex
	indexMu           sync.Mutex
	exportNames       map[string]existingExport
	identityAliases   map[string]string
	retireEntries     []scanEntry
	ScanOptions       ScanOptions
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
	rawConv, err := v.ExtractConversation(sourcePath, explicitTool)
	if err != nil {
		return nil, nil, err
	}
	return v.StoreConversation(rawConv)
}

// ExtractConversation reads and normalizes a source without persisting it.
func (v *Vault) ExtractConversation(sourcePath string, explicitTool string) (*models.Conversation, error) {
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
		return nil, fmt.Errorf("no suitable extractor found for path: %s", sourcePath)
	}

	rawConv, err := matched.Extract(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("extraction error: %w", err)
	}
	return rawConv, nil
}

// ProcessAndStoreAll imports every record from generic JSON/JSONL input.
func (v *Vault) ProcessAndStoreAll(sourcePath string, explicitTool string) ([]*models.Conversation, []string, int, error) {
	if explicitTool != "" && !strings.EqualFold(explicitTool, "generic") {
		release, err := v.lockWriter()
		if err != nil {
			return nil, nil, 0, err
		}
		defer release()
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
	release, err := v.lockWriter()
	if err != nil {
		return nil, nil, 0, err
	}
	defer release()
	v.jsonExport.BeginBatch()
	var stored []*models.Conversation
	var warnings []string
	skipped := 0
	for _, conversation := range conversations {
		cleanConversation, auditWarnings, err := v.StoreConversation(conversation)
		if err != nil {
			flushErr := v.jsonExport.Flush(filepath.Join(v.BaseDir, "conversations", "dataset.jsonl"))
			return stored, warnings, skipped, fmt.Errorf("failed to store %s: %w (dataset flush: %v)", conversation.ID, err, flushErr)
		}
		if cleanConversation == nil {
			skipped++
			continue
		}
		stored = append(stored, cleanConversation)
		warnings = append(warnings, auditWarnings...)
	}
	err = v.jsonExport.Flush(filepath.Join(v.BaseDir, "conversations", "dataset.jsonl"))
	return stored, warnings, skipped, err
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

	// Extraction and sanitization can run concurrently; output naming and writes
	// have one owner so workers cannot race on the dataset or filename registry.
	v.storeMu.Lock()
	defer v.storeMu.Unlock()
	if err := v.loadExportNames(); err != nil {
		return nil, nil, err
	}
	if existingID := v.identityAliases[cleanConv.ID]; existingID != "" {
		cleanConv.ID = existingID
	}

	fileSlug := fmt.Sprintf("%s_%s-%s", toKebabCase(cleanConv.SourceTool, 20), conversationTopic(cleanConv), cleanConv.ID[len(cleanConv.ID)-12:])
	cleanConv.CreatedAt = cleanConv.CreatedAt.Truncate(time.Second)
	if existing, ok := v.exportNames[cleanConv.ID]; ok {
		fileSlug = existing.Slug
		if !existing.CreatedAt.IsZero() {
			cleanConv.CreatedAt = existing.CreatedAt
		}
	}

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
	v.exportNames[cleanConv.ID] = existingExport{Slug: fileSlug, CreatedAt: cleanConv.CreatedAt}

	return cleanConv, auditWarnings, nil
}

func stableConversationID(conv *models.Conversation) string {
	languages := append([]string(nil), conv.Languages...)
	sort.Strings(languages)
	// Exclude volatile source IDs and timestamps so re-importing unchanged content is idempotent.
	payload := struct {
		SourceTool  string            `json:"source_tool"`
		Title       string            `json:"title"`
		Project     string            `json:"project"`
		Description string            `json:"description"`
		Languages   []string          `json:"languages"`
		Tags        []string          `json:"tags"`
		Messages    []models.Message  `json:"messages"`
		Metadata    map[string]string `json:"metadata"`
	}{
		SourceTool:  conv.SourceTool,
		Title:       conv.Title,
		Project:     conv.Project,
		Description: conv.Description,
		Languages:   languages,
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
	if runes := []rune(result); len(runes) > maxLength {
		result = string(runes[:maxLength])
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
	var sources []ScanSource
	agyBrain := filepath.Join(homeDir, ".gemini", "antigravity-cli", "brain")
	if _, err := os.Stat(agyBrain); err == nil {
		sources = append(sources, ScanSource{Tool: "antigravity", Path: agyBrain})
	}
	opencodePaths := []string{
		filepath.Join(homeDir, ".local", "share", "opencode", "opencode.db"),
		filepath.Join(homeDir, ".opencode", "opencode.db"),
		filepath.Join(homeDir, ".config", "opencode", "opencode.db"),
	}
	for _, p := range opencodePaths {
		if _, err := os.Stat(p); err == nil {
			sources = append(sources, ScanSource{Tool: "opencode", Path: p})
			break
		}
	}

	codexSessionsDir := filepath.Join(homeDir, ".codex", "sessions")
	if _, err := os.Stat(codexSessionsDir); err == nil {
		sources = append(sources, ScanSource{Tool: "codex", Path: codexSessionsDir})
	}

	aiderPaths := []string{
		filepath.Join(homeDir, ".aider.chat.history.md"),
		filepath.Join(".", ".aider.chat.history.md"),
	}
	for _, ap := range aiderPaths {
		if _, err := os.Stat(ap); err == nil {
			sources = append(sources, ScanSource{Tool: "aider", Path: ap})
		}
	}

	return v.ScanSources(sources)
}

// ScanAGYBrain scans the user's Antigravity brain folder
func (v *Vault) ScanAGYBrain(brainDir string) (int, int, []string, error) {
	return v.scanOne("antigravity", brainDir)
}

// ScanOpenCodeDB scans and extracts all conversations from OpenCode SQLite DB
func (v *Vault) ScanOpenCodeDB(dbPath string) (int, int, []string, error) {
	return v.scanOne("opencode", dbPath)
}

// ScanCodexSessions recursively scans and extracts all Codex rollout-*.jsonl session files
func (v *Vault) ScanCodexSessions(sessionsDir string) (int, int, []string, error) {
	return v.scanOne("codex", sessionsDir)
}

// Search uses the local ranked passage index with default options.
func (v *Vault) Search(query string) ([]SearchResult, error) {
	return v.SearchWithOptions(query, SearchOptions{})
}

// GenerateContextSnippet formats top matching conversation turns to feed into a local LLM prompt
func (v *Vault) GenerateContextSnippet(query string, maxEntries int) (string, error) {
	return v.ContextWithOptions(query, SearchOptions{Limit: maxEntries}, 12000)
}

// AuditAll audits all stored markdown and json files to verify zero secret leaks
func (v *Vault) AuditAll() (map[string][]string, error) {
	violations := make(map[string][]string)
	convDir := filepath.Join(v.BaseDir, "conversations")

	err := filepath.Walk(convDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		extension := strings.ToLower(filepath.Ext(path))
		if extension != ".md" && extension != ".json" && extension != ".jsonl" {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		warnings := v.Sanitizer.AuditText(string(contentBytes))
		if len(warnings) > 0 {
			violations[path] = warnings
		}
		return nil
	})

	return violations, err
}
