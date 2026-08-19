package extractor

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/models"
	_ "modernc.org/sqlite"
)

type OpenCodeMessageData struct {
	Role string `json:"role"` // "user", "assistant"
}

type OpenCodePartData struct {
	Type   string `json:"type"` // "text", "tool", "reasoning", "patch", "file", etc.
	Text   string `json:"text,omitempty"`
	Tool   string `json:"tool,omitempty"`
	Input  any    `json:"input,omitempty"`
	Output any    `json:"output,omitempty"`
}

// OpenCodeExtractor extracts conversations from OpenCode SQLite database
type OpenCodeExtractor struct{}

func NewOpenCodeExtractor() *OpenCodeExtractor {
	return &OpenCodeExtractor{}
}

func (e *OpenCodeExtractor) Name() string {
	return "opencode"
}

func (e *OpenCodeExtractor) CanHandle(path string) bool {
	base := strings.ToLower(filepath.Base(path))
	return base == "opencode.db" || strings.Contains(strings.ToLower(path), "opencode")
}

func (e *OpenCodeExtractor) Extract(targetPath string) (*models.Conversation, error) {
	convs, err := e.ExtractAll(targetPath)
	if err != nil {
		return nil, err
	}
	if len(convs) == 0 {
		return nil, fmt.Errorf("no conversations found in opencode database: %s", targetPath)
	}
	return convs[len(convs)-1], nil
}

// ExtractAll extracts all conversations stored in the OpenCode SQLite database
func (e *OpenCodeExtractor) ExtractAll(dbPath string) ([]*models.Conversation, error) {
	if _, err := os.Stat(dbPath); err != nil {
		return nil, fmt.Errorf("opencode db not found at %s: %w", dbPath, err)
	}

	dsn := fmt.Sprintf("file:%s?mode=ro", filepath.ToSlash(dbPath))
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database %s: %w", dbPath, err)
	}
	defer db.Close()

	// 1. Fetch sessions
	sessionRows, err := db.Query("SELECT id, COALESCE(title, ''), time_created FROM session ORDER BY time_created ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to query sessions: %w", err)
	}
	defer sessionRows.Close()

	type SessionRecord struct {
		ID          string
		Title       string
		TimeCreated int64
	}

	var sessions []SessionRecord
	for sessionRows.Next() {
		var s SessionRecord
		if err := sessionRows.Scan(&s.ID, &s.Title, &s.TimeCreated); err == nil {
			sessions = append(sessions, s)
		}
	}

	var results []*models.Conversation

	// 2. For each session, fetch messages and parts
	for _, sess := range sessions {
		msgRows, err := db.Query("SELECT id, time_created, COALESCE(data, '{}') FROM message WHERE session_id = ? ORDER BY time_created ASC", sess.ID)
		if err != nil {
			continue
		}

		type MsgRecord struct {
			ID          string
			TimeCreated int64
			Data        string
		}
		var msgList []MsgRecord
		for msgRows.Next() {
			var m MsgRecord
			if err := msgRows.Scan(&m.ID, &m.TimeCreated, &m.Data); err == nil {
				msgList = append(msgList, m)
			}
		}
		msgRows.Close()

		var conversationMessages []models.Message
		var firstUserInput string

		for _, m := range msgList {
			var msgData OpenCodeMessageData
			_ = json.Unmarshal([]byte(m.Data), &msgData)

			role := strings.ToLower(msgData.Role)
			if role == "" {
				role = "assistant"
			}

			partRows, err := db.Query("SELECT COALESCE(data, '{}') FROM part WHERE message_id = ? ORDER BY time_created ASC", m.ID)
			if err != nil {
				continue
			}

			var textParts []string
			var toolCalls []models.ToolCall

			for partRows.Next() {
				var dataStr string
				if err := partRows.Scan(&dataStr); err == nil {
					var part OpenCodePartData
					if err := json.Unmarshal([]byte(dataStr), &part); err == nil {
						if part.Type == "text" && strings.TrimSpace(part.Text) != "" {
							textParts = append(textParts, strings.TrimSpace(part.Text))
						} else if part.Type == "tool" || part.Tool != "" {
							inputJSON, _ := json.Marshal(part.Input)
							outputJSON, _ := json.Marshal(part.Output)
							toolCalls = append(toolCalls, models.ToolCall{
								Name:      part.Tool,
								Summary:   fmt.Sprintf("Tool: %s", part.Tool),
								Arguments: string(inputJSON),
								Output:    string(outputJSON),
							})
						}
					}
				}
			}
			partRows.Close()

			content := strings.TrimSpace(strings.Join(textParts, "\n\n"))
			if role == "user" && firstUserInput == "" && content != "" {
				firstUserInput = content
			}

			if content != "" || len(toolCalls) > 0 {
				t := time.UnixMilli(m.TimeCreated)
				conversationMessages = append(conversationMessages, models.Message{
					Role:      role,
					Content:   content,
					Timestamp: &t,
					ToolCalls: toolCalls,
				})
			}
		}

		if len(conversationMessages) == 0 {
			continue
		}

		title := sess.Title
		if title == "" || strings.HasPrefix(title, "New session -") {
			if firstUserInput != "" {
				title = generateTitle(firstUserInput)
			}
		}
		if title == "" {
			title = fmt.Sprintf("OpenCode Session %s", sess.ID[:min(8, len(sess.ID))])
		}

		results = append(results, &models.Conversation{
			ID:         sess.ID,
			SourceTool: "opencode",
			Title:      title,
			CreatedAt:  time.UnixMilli(sess.TimeCreated),
			Tags:       []string{"coding", "assistant", "opencode"},
			Messages:   conversationMessages,
		})
	}

	return results, nil
}
