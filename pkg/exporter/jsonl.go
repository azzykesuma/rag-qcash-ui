package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"llm-context-vault/pkg/models"
)

// JSONLExporter exports conversations to ShareGPT and standard JSONL formats
type JSONLExporter struct{}

func NewJSONLExporter() *JSONLExporter {
	return &JSONLExporter{}
}

// ExportShareGPT writes a single conversation in ShareGPT format
func (e *JSONLExporter) ExportShareGPT(conv *models.Conversation, outPath string) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	shareGPT := models.ShareGPTFormat{
		ID:     conv.ID,
		Source: conv.SourceTool,
		Tags:   conv.Tags,
	}

	for _, msg := range conv.Messages {
		from := "human"
		if msg.Role == "assistant" {
			from = "gpt"
		} else if msg.Role == "system" {
			from = "system"
		}

		shareGPT.Conversations = append(shareGPT.Conversations, models.ShareGPTMessage{
			From:  from,
			Value: msg.Content,
		})
	}

	data, err := json.MarshalIndent(shareGPT, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode sharegpt json: %w", err)
	}

	return os.WriteFile(outPath, data, 0644)
}

// AppendToJSONL appends normalized conversation as a single line in a dataset JSONL file
func (e *JSONLExporter) AppendToJSONL(conv *models.Conversation, jsonlPath string) error {
	if err := os.MkdirAll(filepath.Dir(jsonlPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.OpenFile(jsonlPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open dataset file: %w", err)
	}
	defer file.Close()

	line, err := json.Marshal(conv)
	if err != nil {
		return fmt.Errorf("failed to marshal conversation: %w", err)
	}

	_, err = file.Write(append(line, '\n'))
	return err
}
