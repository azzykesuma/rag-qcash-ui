package exporter

import (
	"bufio"
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

// UpsertJSONL replaces an existing conversation with the same stable ID or appends it.
func (e *JSONLExporter) UpsertJSONL(conv *models.Conversation, jsonlPath string) error {
	if err := os.MkdirAll(filepath.Dir(jsonlPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	conversations, err := readJSONL(jsonlPath)
	if err != nil {
		return err
	}

	found := false
	for i, existing := range conversations {
		if existing.ID == conv.ID {
			conversations[i] = *conv
			found = true
			break
		}
	}
	if !found {
		conversations = append(conversations, *conv)
	}

	file, err := os.Create(jsonlPath)
	if err != nil {
		return fmt.Errorf("failed to rewrite dataset file: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, conversation := range conversations {
		if err := encoder.Encode(conversation); err != nil {
			return fmt.Errorf("failed to encode dataset conversation: %w", err)
		}
	}
	return nil
}

func readJSONL(path string) ([]models.Conversation, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open dataset file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 16*1024*1024)
	var conversations []models.Conversation
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if len(scanner.Bytes()) == 0 {
			continue
		}
		var conversation models.Conversation
		if err := json.Unmarshal(scanner.Bytes(), &conversation); err != nil {
			return nil, fmt.Errorf("invalid dataset JSONL at line %d: %w", lineNumber, err)
		}
		conversations = append(conversations, conversation)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read dataset file: %w", err)
	}
	return conversations, nil
}
