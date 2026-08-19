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

// CodexLogItem represents a line in Codex's rollout-*.jsonl
type CodexLogItem struct {
	Timestamp string          `json:"timestamp"`
	Type      string          `json:"type"` // session_meta, event_msg, response_item, turn_context
	Payload   json.RawMessage `json:"payload"`
}

type CodexSessionMetaPayload struct {
	SessionID string `json:"session_id"`
	ID        string `json:"id"`
	Cwd       string `json:"cwd"`
	Timestamp string `json:"timestamp"`
}

type CodexEventMsgPayload struct {
	Type    string `json:"type"`    // user_message, agent_message, task_started, task_complete
	Message string `json:"message"` // Text for user_message / agent_message
	Phase   string `json:"phase"`
}

type CodexResponseItemPayload struct {
	Type      string                 `json:"type"` // message, reasoning, custom_tool_call
	Role      string                 `json:"role"` // developer, user, assistant
	Content   []CodexContentFragment `json:"content"`
	Name      string                 `json:"name"` // For custom_tool_call
	Input     string                 `json:"input"`
	CallID    string                 `json:"call_id"`
	Status    string                 `json:"status"`
}

type CodexContentFragment struct {
	Type string `json:"type"` // input_text, output_text
	Text string `json:"text"`
}

// CodexExtractor extracts conversations from OpenAI Codex CLI session logs
type CodexExtractor struct{}

func NewCodexExtractor() *CodexExtractor {
	return &CodexExtractor{}
}

func (e *CodexExtractor) Name() string {
	return "codex"
}

func (e *CodexExtractor) CanHandle(path string) bool {
	base := strings.ToLower(filepath.Base(path))
	return (strings.HasPrefix(base, "rollout-") && strings.HasSuffix(base, ".jsonl")) || strings.Contains(strings.ToLower(path), ".codex")
}

func (e *CodexExtractor) Extract(targetPath string) (*models.Conversation, error) {
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open codex session file %s: %w", targetPath, err)
	}
	defer file.Close()

	var messages []models.Message
	var firstUserInput string
	var sessionID string
	var sessionTime time.Time

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 20*1024*1024) // 20MB buffer for large logs

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var item CodexLogItem
		if err := json.Unmarshal([]byte(line), &item); err != nil {
			continue
		}

		switch item.Type {
		case "session_meta":
			var meta CodexSessionMetaPayload
			if err := json.Unmarshal(item.Payload, &meta); err == nil {
				sessionID = meta.SessionID
				if sessionID == "" {
					sessionID = meta.ID
				}
				if meta.Timestamp != "" {
					if t, err := time.Parse(time.RFC3339, meta.Timestamp); err == nil {
						sessionTime = t
					}
				}
			}

		case "event_msg":
			var evt CodexEventMsgPayload
			if err := json.Unmarshal(item.Payload, &evt); err != nil {
				continue
			}

			if evt.Type == "user_message" && strings.TrimSpace(evt.Message) != "" {
				text := strings.TrimSpace(evt.Message)
				if !isInternalCodexInstruction(text) {
					if firstUserInput == "" {
						firstUserInput = text
					}
					messages = append(messages, models.Message{
						Role:    "user",
						Content: text,
					})
				}
			} else if evt.Type == "agent_message" && strings.TrimSpace(evt.Message) != "" {
				text := strings.TrimSpace(evt.Message)
				messages = append(messages, models.Message{
					Role:    "assistant",
					Content: text,
				})
			}

		case "response_item":
			var resp CodexResponseItemPayload
			if err := json.Unmarshal(item.Payload, &resp); err != nil {
				continue
			}

			if resp.Type == "message" {
				var sb strings.Builder
				for _, frag := range resp.Content {
					if frag.Text != "" {
						sb.WriteString(frag.Text)
					}
				}
				text := strings.TrimSpace(sb.String())
				if text != "" && !isInternalCodexInstruction(text) {
					role := strings.ToLower(resp.Role)
					if role == "developer" || role == "system" {
						// Skip system prompts if desired, or keep as system
						continue
					}
					if role == "user" {
						if firstUserInput == "" {
							firstUserInput = text
						}
					}
					messages = append(messages, models.Message{
						Role:    role,
						Content: text,
					})
				}
			} else if resp.Type == "custom_tool_call" && resp.Name != "" {
				messages = append(messages, models.Message{
					Role: "assistant",
					ToolCalls: []models.ToolCall{
						{
							Name:      resp.Name,
							Summary:   fmt.Sprintf("Executed tool %s", resp.Name),
							Arguments: resp.Input,
						},
					},
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning codex file %s: %w", targetPath, err)
	}

	// Deduplicate consecutive messages with identical content
	deduped := deduplicateMessages(messages)
	if len(deduped) == 0 {
		return nil, fmt.Errorf("no user/assistant conversation turns found in %s", targetPath)
	}

	if sessionID == "" {
		sessionID = strings.TrimSuffix(filepath.Base(targetPath), ".jsonl")
	}
	if sessionTime.IsZero() {
		sessionTime = time.Now()
	}

	title := generateTitle(firstUserInput)
	if title == "Untitled Conversation" {
		title = fmt.Sprintf("Codex Session %s", sessionID[:min(8, len(sessionID))])
	}

	return &models.Conversation{
		ID:         sessionID,
		SourceTool: "codex",
		Title:      title,
		CreatedAt:  sessionTime,
		Tags:       []string{"coding", "assistant", "codex", "gpt-5"},
		Messages:   deduped,
	}, nil
}

func isInternalCodexInstruction(s string) bool {
	// Skip environment context injections, skills instructions, and permissions
	trimmed := strings.TrimSpace(s)
	return strings.HasPrefix(trimmed, "<permissions instructions>") ||
		strings.HasPrefix(trimmed, "<collaboration_mode>") ||
		strings.HasPrefix(trimmed, "<skills_instructions>") ||
		strings.HasPrefix(trimmed, "<apps_instructions>") ||
		strings.HasPrefix(trimmed, "<plugins_instructions>") ||
		strings.HasPrefix(trimmed, "<environment_context>") ||
		strings.HasPrefix(trimmed, "# AGENTS.md instructions")
}

func deduplicateMessages(msgs []models.Message) []models.Message {
	if len(msgs) == 0 {
		return msgs
	}
	var out []models.Message
	for _, m := range msgs {
		if len(out) > 0 && out[len(out)-1].Role == m.Role && out[len(out)-1].Content == m.Content && len(m.ToolCalls) == 0 {
			continue
		}
		out = append(out, m)
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
