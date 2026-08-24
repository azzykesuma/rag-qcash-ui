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

// GenericExtractor handles standardized JSON and ShareGPT format files
type GenericExtractor struct{}

func NewGenericExtractor() *GenericExtractor {
	return &GenericExtractor{}
}

func (e *GenericExtractor) Name() string {
	return "generic"
}

func (e *GenericExtractor) CanHandle(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json" || ext == ".jsonl"
}

func (e *GenericExtractor) Extract(targetPath string) (*models.Conversation, error) {
	conversations, err := e.ExtractAll(targetPath)
	if err != nil {
		return nil, err
	}
	if len(conversations) != 1 {
		return nil, fmt.Errorf("expected one conversation in %s, found %d; use import to process all records", targetPath, len(conversations))
	}
	return conversations[0], nil
}

// ExtractAll parses JSON documents and JSONL files containing one conversation per line.
func (e *GenericExtractor) ExtractAll(targetPath string) ([]*models.Conversation, error) {
	if strings.EqualFold(filepath.Ext(targetPath), ".jsonl") {
		return e.extractJSONL(targetPath)
	}

	data, err := os.ReadFile(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read generic json %s: %w", targetPath, err)
	}
	if conversations, err := parseConversationList(data, targetPath); err == nil {
		return conversations, nil
	}
	conversation, err := parseConversation(data, targetPath)
	if err != nil {
		return nil, err
	}
	return []*models.Conversation{conversation}, nil
}

func (e *GenericExtractor) extractJSONL(targetPath string) ([]*models.Conversation, error) {
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read generic JSONL %s: %w", targetPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 16*1024*1024)
	var conversations []*models.Conversation
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		conversation, err := parseConversation([]byte(scanner.Text()), fmt.Sprintf("%s:%d", targetPath, lineNumber))
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read generic JSONL %s: %w", targetPath, err)
	}
	if len(conversations) == 0 {
		return nil, fmt.Errorf("no conversation records found in %s", targetPath)
	}
	return conversations, nil
}

func parseConversationList(data []byte, sourcePath string) ([]*models.Conversation, error) {
	var records []json.RawMessage
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no conversation records found in %s", sourcePath)
	}
	conversations := make([]*models.Conversation, 0, len(records))
	for i, record := range records {
		conversation, err := parseConversation(record, fmt.Sprintf("%s[%d]", sourcePath, i))
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	return conversations, nil
}

func parseConversation(data []byte, sourcePath string) (*models.Conversation, error) {

	// 1. Try parsing directly as models.Conversation
	var conv models.Conversation
	if err := json.Unmarshal(data, &conv); err == nil && len(conv.Messages) > 0 {
		if conv.ID == "" {
			conv.ID = filepath.Base(sourcePath)
		}
		if conv.CreatedAt.IsZero() {
			conv.CreatedAt = time.Now()
		}
		return &conv, nil
	}

	// 2. Try parsing as ShareGPT format
	var shareGPT models.ShareGPTFormat
	if err := json.Unmarshal(data, &shareGPT); err == nil && len(shareGPT.Conversations) > 0 {
		var messages []models.Message
		for _, msg := range shareGPT.Conversations {
			role := "user"
			if msg.From == "gpt" || msg.From == "assistant" {
				role = "assistant"
			} else if msg.From == "system" {
				role = "system"
			}
			messages = append(messages, models.Message{
				Role:    role,
				Content: msg.Value,
			})
		}

		title := "Shared Conversation"
		if len(messages) > 0 {
			title = generateTitle(messages[0].Content)
		}

		return &models.Conversation{
			ID:         shareGPT.ID,
			SourceTool: "sharegpt",
			Title:      title,
			CreatedAt:  time.Now(),
			Tags:       shareGPT.Tags,
			Messages:   messages,
		}, nil
	}

	return nil, fmt.Errorf("unrecognized conversation JSON structure in %s", sourcePath)
}
