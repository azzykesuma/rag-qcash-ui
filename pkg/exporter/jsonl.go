package exporter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"llm-context-vault/pkg/models"
)

// JSONLExporter exports conversations to ShareGPT and standard JSONL formats
type JSONLExporter struct {
	pending        map[string]*models.Conversation
	pendingDeletes map[string]bool
	rebuild        bool
}

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

	return WriteIfChanged(outPath, data)
}

// UpsertJSONL replaces an existing conversation with the same stable ID or appends it.
func (e *JSONLExporter) UpsertJSONL(conv *models.Conversation, jsonlPath string) error {
	if e.pending != nil {
		e.pending[conv.ID] = conv
		delete(e.pendingDeletes, conv.ID)
		return nil
	}
	return mergeJSONL(map[string]*models.Conversation{conv.ID: conv}, nil, jsonlPath, false)
}

// BeginBatch defers dataset updates until Flush. Calls must be serialized by the owner.
func (e *JSONLExporter) BeginBatch() {
	e.pending = make(map[string]*models.Conversation)
	e.pendingDeletes = make(map[string]bool)
	e.rebuild = false
}

// BeginRebuild replaces the dataset from authoritative source records on Flush.
func (e *JSONLExporter) BeginRebuild() {
	e.BeginBatch()
	e.rebuild = true
}

// DeleteJSONL retires a superseded content ID in the current batch.
func (e *JSONLExporter) DeleteJSONL(id string) {
	if e.pendingDeletes == nil {
		e.pendingDeletes = make(map[string]bool)
	}
	delete(e.pending, id)
	e.pendingDeletes[id] = true
}

func (e *JSONLExporter) Flush(path string) error {
	if !e.rebuild && len(e.pending) == 0 && len(e.pendingDeletes) == 0 {
		return nil
	}
	updates := make(map[string]*models.Conversation, len(e.pending))
	for id, conversation := range e.pending {
		updates[id] = conversation
	}
	deletes := make(map[string]bool, len(e.pendingDeletes))
	for id := range e.pendingDeletes {
		deletes[id] = true
	}
	if err := mergeJSONL(updates, deletes, path, e.rebuild); err != nil {
		return err
	}
	e.pending = nil
	e.pendingDeletes = nil
	e.rebuild = false
	return nil
}

// mergeJSONL streams the existing dataset once, retaining untouched records byte
// for byte. A failed read/write leaves the previous dataset intact.
func mergeJSONL(updates map[string]*models.Conversation, deletes map[string]bool, jsonlPath string, rebuild bool) error {
	if err := os.MkdirAll(filepath.Dir(jsonlPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	if err := recoverFile(jsonlPath); err != nil {
		return fmt.Errorf("recover previous dataset: %w", err)
	}
	tempDir := filepath.Join(filepath.Dir(jsonlPath), ".vault-tmp")
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return err
	}
	file, err := os.CreateTemp(tempDir, "dataset-*")
	if err != nil {
		return err
	}
	defer os.Remove(tempDir)
	defer file.Close()
	defer os.Remove(file.Name())
	w := bufio.NewWriter(file)
	changed := rebuild
	var input *os.File
	if !rebuild {
		input, err = os.Open(jsonlPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	if input != nil {
		scanner := bufio.NewScanner(input)
		scanner.Buffer(make([]byte, 64*1024), 64*1024*1024)
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(bytes.TrimSpace(line)) == 0 {
				continue
			}
			var header struct {
				ID string `json:"id"`
			}
			if err := json.Unmarshal(line, &header); err != nil {
				input.Close()
				return fmt.Errorf("invalid dataset JSONL: %w", err)
			}
			if conv, ok := updates[header.ID]; ok {
				data, err := json.Marshal(conv)
				if err != nil {
					input.Close()
					return err
				}
				changed = changed || !bytes.Equal(data, line)
				line = data
				delete(updates, header.ID)
			}
			if deletes[header.ID] {
				changed = true
				continue
			}
			if _, err := w.Write(line); err != nil {
				input.Close()
				return err
			}
			if err := w.WriteByte('\n'); err != nil {
				input.Close()
				return err
			}
		}
		err = scanner.Err()
		input.Close()
		if err != nil {
			return err
		}
	}
	ids := make([]string, 0, len(updates))
	for id := range updates {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	encoder := json.NewEncoder(w)
	for _, id := range ids {
		changed = true
		if err := encoder.Encode(updates[id]); err != nil {
			return err
		}
	}
	if !changed {
		return nil
	}
	if err := w.Flush(); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return replaceFile(file.Name(), jsonlPath)
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
