package extractor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/models"
)

// AntigravityLogStep represents a step in AGY's transcript.jsonl
type AntigravityLogStep struct {
	StepIndex int               `json:"step_index"`
	Source    string            `json:"source"` // USER_EXPLICIT, MODEL, SYSTEM
	Type      string            `json:"type"`   // USER_INPUT, PLANNER_RESPONSE, etc.
	Status    string            `json:"status"`
	Content   string            `json:"content"`
	ToolCalls []AGYToolCallItem `json:"tool_calls,omitempty"`
}

type AGYToolCallItem struct {
	Name        string `json:"name"`
	ToolAction  string `json:"toolAction,omitempty"`
	ToolSummary string `json:"toolSummary,omitempty"`
	Arguments   any    `json:"arguments,omitempty"`
}

// AGYExtractor extracts conversations from Google Antigravity (AGY) transcript logs
type AGYExtractor struct{}

func NewAGYExtractor() *AGYExtractor {
	return &AGYExtractor{}
}

func (e *AGYExtractor) Name() string {
	return "antigravity"
}

func (e *AGYExtractor) CanHandle(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if info.IsDir() {
		// Check if it's a brain conversation dir
		transcriptPath := filepath.Join(path, ".system_generated", "logs", "transcript.jsonl")
		if _, err := os.Stat(transcriptPath); err == nil {
			return true
		}
		transcriptFullPath := filepath.Join(path, ".system_generated", "logs", "transcript_full.jsonl")
		if _, err := os.Stat(transcriptFullPath); err == nil {
			return true
		}
	}

	base := strings.ToLower(filepath.Base(path))
	return strings.Contains(base, "transcript") && strings.HasSuffix(base, ".jsonl")
}

func (e *AGYExtractor) Extract(targetPath string) (*models.Conversation, error) {
	resolvedFile := targetPath

	info, err := os.Stat(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect path %s: %w", targetPath, err)
	}

	convID := filepath.Base(targetPath)

	if info.IsDir() {
		transcriptFull := filepath.Join(targetPath, ".system_generated", "logs", "transcript_full.jsonl")
		transcript := filepath.Join(targetPath, ".system_generated", "logs", "transcript.jsonl")

		if _, err := os.Stat(transcriptFull); err == nil {
			resolvedFile = transcriptFull
		} else if _, err := os.Stat(transcript); err == nil {
			resolvedFile = transcript
		} else {
			return nil, fmt.Errorf("no transcript.jsonl found in directory %s", targetPath)
		}
	} else {
		// Infer ID from parent directory structure
		parts := strings.Split(filepath.ToSlash(targetPath), "/")
		for i, part := range parts {
			if part == "brain" && i+1 < len(parts) {
				convID = parts[i+1]
				break
			}
		}
	}

	file, err := os.Open(resolvedFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open transcript file %s: %w", resolvedFile, err)
	}
	defer file.Close()

	var messages []models.Message
	var firstUserInput string
	var languages []string
	langMap := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	// Allow large tokens for lengthy assistant responses / tool outputs (up to 16MB per line)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 16*1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var step AntigravityLogStep
		if err := json.Unmarshal([]byte(line), &step); err != nil {
			// Skip unparseable lines gracefully
			continue
		}

		switch step.Type {
		case "USER_INPUT":
			cleanContent := strings.TrimSpace(step.Content)
			if cleanContent == "" {
				continue
			}
			// Strip XML wrappers if present
			cleanContent = stripXMLTag(cleanContent, "USER_REQUEST")

			if firstUserInput == "" {
				firstUserInput = cleanContent
			}

			messages = append(messages, models.Message{
				Role:    "user",
				Content: cleanContent,
			})

		case "PLANNER_RESPONSE":
			cleanContent := strings.TrimSpace(step.Content)
			var toolCalls []models.ToolCall

			for _, tc := range step.ToolCalls {
				argsJSON, _ := json.Marshal(tc.Arguments)
				toolCalls = append(toolCalls, models.ToolCall{
					Name:      tc.Name,
					Summary:   tc.ToolSummary,
					Arguments: string(argsJSON),
				})

				// Infer languages from tool calls / extensions
				detectLanguages(string(argsJSON), langMap)
			}

			if cleanContent != "" || len(toolCalls) > 0 {
				messages = append(messages, models.Message{
					Role:      "assistant",
					Content:   cleanContent,
					ToolCalls: toolCalls,
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading transcript %s: %w", resolvedFile, err)
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("no user/assistant messages extracted from %s", resolvedFile)
	}

	for l := range langMap {
		languages = append(languages, l)
	}

	// Generate title from first user prompt
	title := generateTitle(firstUserInput)

	return &models.Conversation{
		ID:         convID,
		SourceTool: "antigravity",
		Title:      title,
		CreatedAt:  time.Now(),
		Languages:  languages,
		Tags:       []string{"coding", "assistant", "antigravity"},
		Messages:   messages,
	}, nil
}

func stripXMLTag(content, tag string) string {
	openTag := "<" + tag + ">"
	closeTag := "</" + tag + ">"
	start := strings.Index(content, openTag)
	if start != -1 {
		end := strings.Index(content, closeTag)
		if end != -1 && end > start {
			return strings.TrimSpace(content[start+len(openTag) : end])
		}
	}
	return content
}

func generateTitle(prompt string) string {
	if prompt == "" {
		return "Untitled Conversation"
	}
	lines := strings.Split(prompt, "\n")
	firstLine := strings.TrimSpace(lines[0])
	if len(firstLine) > 70 {
		return firstLine[:67] + "..."
	}
	return firstLine
}

func detectLanguages(content string, langMap map[string]bool) {
	lower := strings.ToLower(content)
	if strings.Contains(lower, ".go") || strings.Contains(lower, "go.mod") {
		langMap["go"] = true
	}
	if strings.Contains(lower, ".ts") || strings.Contains(lower, ".tsx") {
		langMap["typescript"] = true
	}
	if strings.Contains(lower, ".js") || strings.Contains(lower, "package.json") {
		langMap["javascript"] = true
	}
	if strings.Contains(lower, ".py") || strings.Contains(lower, "requirements.txt") {
		langMap["python"] = true
	}
	if strings.Contains(lower, ".rs") || strings.Contains(lower, "cargo.toml") {
		langMap["rust"] = true
	}
	if strings.Contains(lower, ".java") || strings.Contains(lower, "pom.xml") {
		langMap["java"] = true
	}
}
