package extractor

import (
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
	data, err := os.ReadFile(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read generic json %s: %w", targetPath, err)
	}

	// 1. Try parsing directly as models.Conversation
	var conv models.Conversation
	if err := json.Unmarshal(data, &conv); err == nil && len(conv.Messages) > 0 {
		if conv.ID == "" {
			conv.ID = filepath.Base(targetPath)
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

	return nil, fmt.Errorf("unrecognized conversation json structure in %s", targetPath)
}
