package extractor

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/models"
	_ "modernc.org/sqlite"
)

type OpenCodeMessageData struct {
	Role string `json:"role"`
}
type OpenCodePartData struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	Tool   string `json:"tool,omitempty"`
	Input  any    `json:"input,omitempty"`
	Output any    `json:"output,omitempty"`
}

type OpenCodeExtractor struct{}

func NewOpenCodeExtractor() *OpenCodeExtractor { return &OpenCodeExtractor{} }
func (e *OpenCodeExtractor) Name() string      { return "opencode" }
func (e *OpenCodeExtractor) CanHandle(path string) bool {
	return strings.EqualFold(filepath.Base(path), "opencode.db") || strings.Contains(strings.ToLower(path), "opencode")
}
func (e *OpenCodeExtractor) Extract(path string) (*models.Conversation, error) {
	convs, err := e.ExtractAll(path)
	if err != nil {
		return nil, err
	}
	if len(convs) == 0 {
		return nil, fmt.Errorf("no conversations found in opencode database: %s", path)
	}
	return convs[len(convs)-1], nil
}
func (e *OpenCodeExtractor) ExtractAll(path string) ([]*models.Conversation, error) {
	return e.ExtractIncremental(path, func(_, _ string) bool { return true })
}

// ExtractIncremental checks session AND child revisions within a read snapshot.
// Schemas without update timestamps safely fall back to extraction on every scan.
// Selected sessions use one joined query each instead of one query per message.
func (e *OpenCodeExtractor) ExtractIncremental(path string, selectSession func(id, revision string) bool) ([]*models.Conversation, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	u := url.URL{Scheme: "file", Path: "/" + strings.TrimPrefix(filepath.ToSlash(abs), "/"), RawQuery: "mode=ro"}
	db, err := sql.Open("sqlite", u.String())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	cacheable := true
	hasDirectory := false
	for _, table := range []string{"session", "message", "part"} {
		rows, err := tx.Query("PRAGMA table_info(" + table + ")")
		if err != nil {
			return nil, err
		}
		hasUpdated := false
		for rows.Next() {
			var cid, notNull, pk int
			var name, kind string
			var def any
			if err := rows.Scan(&cid, &name, &kind, &notNull, &def, &pk); err != nil {
				rows.Close()
				return nil, err
			}
			if name == "time_updated" {
				hasUpdated = true
			}
			if table == "session" && name == "directory" {
				hasDirectory = true
			}
		}
		err = rows.Err()
		rows.Close()
		if err != nil {
			return nil, err
		}
		cacheable = cacheable && hasUpdated
	}
	updated := "0"
	if cacheable {
		updated = "COALESCE(time_updated, 0)"
	}
	directory := "''"
	if hasDirectory {
		directory = "COALESCE(directory, '')"
	}
	rows, err := tx.Query("SELECT id, COALESCE(title, ''), time_created, " + updated + ", " + directory + " FROM session ORDER BY time_created, id")
	if err != nil {
		return nil, fmt.Errorf("query sessions: %w", err)
	}
	type session struct {
		id, title, directory string
		created, updated     int64
	}
	var sessions []session
	for rows.Next() {
		var s session
		if err := rows.Scan(&s.id, &s.title, &s.created, &s.updated, &s.directory); err != nil {
			rows.Close()
			return nil, err
		}
		sessions = append(sessions, s)
	}
	err = rows.Err()
	rows.Close()
	if err != nil {
		return nil, err
	}
	revisions := make(map[string]string)
	if cacheable {
		// Counts detect deletion; sums detect edits even when the maximum timestamp
		// belongs to another row. IDs detect replacement with preserved timestamps.
		for _, query := range []string{
			"SELECT session_id, COUNT(*), COALESCE(SUM(time_updated),0), COALESCE(MAX(time_created),0), COALESCE(MAX(id),'') FROM message GROUP BY session_id",
			"SELECT m.session_id, COUNT(*), COALESCE(SUM(p.time_updated),0), COALESCE(MAX(p.time_created),0), COALESCE(MAX(p.id),'') FROM part p JOIN message m ON m.id=p.message_id GROUP BY m.session_id",
		} {
			rows, err := tx.Query(query)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var id, maxID string
				var count, sum, created int64
				if err := rows.Scan(&id, &count, &sum, &created, &maxID); err != nil {
					rows.Close()
					return nil, err
				}
				revisions[id] += fmt.Sprintf("|%d:%d:%d:%s", count, sum, created, maxID)
			}
			err = rows.Err()
			rows.Close()
			if err != nil {
				return nil, err
			}
		}
	}
	var results []*models.Conversation
	for _, s := range sessions {
		revision := ""
		if cacheable {
			revision = fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s:%d:%d:%s:%s", s.title, s.created, s.updated, s.directory, revisions[s.id]))))
		}
		if !selectSession(s.id, revision) {
			continue
		}
		rows, err := tx.Query(`SELECT m.id, m.time_created, COALESCE(m.data, '{}'), COALESCE(p.data, '{}')
			FROM message m LEFT JOIN part p ON p.message_id=m.id
			WHERE m.session_id=? ORDER BY m.time_created, m.id, p.time_created, p.id`, s.id)
		if err != nil {
			return results, err
		}
		conv := &models.Conversation{ID: s.id, SourceTool: "opencode", Title: s.title, CreatedAt: time.UnixMilli(s.created), Tags: []string{"coding", "assistant", "opencode"}}
		conv.Project = projectName(s.directory)
		var currentID, firstUser string
		var message models.Message
		var textParts []string
		flush := func() {
			message.Content = strings.TrimSpace(strings.Join(textParts, "\n\n"))
			if message.Role == "user" && firstUser == "" {
				firstUser = message.Content
			}
			if message.Content != "" || len(message.ToolCalls) > 0 {
				conv.Messages = append(conv.Messages, message)
			}
		}
		for rows.Next() {
			var id, messageData, partData string
			var created int64
			if err := rows.Scan(&id, &created, &messageData, &partData); err != nil {
				rows.Close()
				return results, err
			}
			if id != currentID {
				if currentID != "" {
					flush()
				}
				var data OpenCodeMessageData
				if err := json.Unmarshal([]byte(messageData), &data); err != nil {
					rows.Close()
					return results, fmt.Errorf("invalid message %s: %w", id, err)
				}
				t := time.UnixMilli(created)
				role := strings.ToLower(data.Role)
				if role == "" {
					role = "assistant"
				}
				message = models.Message{Role: role, Timestamp: &t}
				textParts = nil
				currentID = id
			}
			var part OpenCodePartData
			if err := json.Unmarshal([]byte(partData), &part); err != nil {
				rows.Close()
				return results, fmt.Errorf("invalid part for message %s: %w", id, err)
			}
			if part.Type == "text" && strings.TrimSpace(part.Text) != "" {
				textParts = append(textParts, strings.TrimSpace(part.Text))
			} else if part.Type == "tool" || part.Tool != "" {
				input, _ := json.Marshal(part.Input)
				output, _ := json.Marshal(part.Output)
				message.ToolCalls = append(message.ToolCalls, models.ToolCall{Name: part.Tool, Summary: "Tool: " + part.Tool, Arguments: string(input), Output: string(output)})
			}
		}
		err = rows.Err()
		rows.Close()
		if err != nil {
			return results, err
		}
		if currentID != "" {
			flush()
		}
		if conv.Title == "" || strings.HasPrefix(conv.Title, "New session -") {
			conv.Title = generateTitle(firstUser)
		}
		results = append(results, conv)
	}
	return results, tx.Commit()
}
