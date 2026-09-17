package exporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"llm-context-vault/pkg/models"
)

func TestBatchMergeAndUnchangedWrites(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.jsonl")
	e := NewJSONLExporter()
	first := &models.Conversation{ID: "first", Title: "original"}
	if err := e.UpsertJSONL(first, path); err != nil {
		t.Fatal(err)
	}
	e.BeginBatch()
	if err := e.UpsertJSONL(&models.Conversation{ID: "first", Title: "updated"}, path); err != nil {
		t.Fatal(err)
	}
	if err := e.UpsertJSONL(&models.Conversation{ID: "second"}, path); err != nil {
		t.Fatal(err)
	}
	before, _ := readJSONL(path)
	if len(before) != 1 || before[0].Title != "original" {
		t.Fatal("batch wrote before flush")
	}
	if err := e.Flush(path); err != nil {
		t.Fatal(err)
	}
	after, err := readJSONL(path)
	if err != nil || len(after) != 2 || after[0].Title != "updated" {
		t.Fatalf("bad merge: %+v %v", after, err)
	}
	old := time.Unix(1234567890, 0)
	if err := os.Chtimes(path, old, old); err != nil {
		t.Fatal(err)
	}
	if err := e.UpsertJSONL(&after[0], path); err != nil {
		t.Fatal(err)
	}
	info, _ := os.Stat(path)
	if !info.ModTime().Equal(old) {
		t.Fatal("identical dataset was rewritten")
	}
}

func TestFailedMergePreservesDataset(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.jsonl")
	original := []byte("{\"id\":\"valid\"}\ninvalid json\n")
	if err := os.WriteFile(path, original, 0600); err != nil {
		t.Fatal(err)
	}
	e := NewJSONLExporter()
	e.BeginBatch()
	if err := e.UpsertJSONL(&models.Conversation{ID: "new"}, path); err != nil {
		t.Fatal(err)
	}
	if err := e.Flush(path); err == nil {
		t.Fatal("expected invalid JSONL error")
	}
	actual, _ := os.ReadFile(path)
	if string(actual) != string(original) {
		t.Fatal("failed merge modified original dataset")
	}
	if err := os.WriteFile(path, []byte(""), 0600); err != nil {
		t.Fatal(err)
	}
	if err := e.Flush(path); err != nil {
		t.Fatalf("batch was not retryable: %v", err)
	}
	conversations, err := readJSONL(path)
	if err != nil || len(conversations) != 1 || conversations[0].ID != "new" {
		t.Fatalf("retry lost pending update: %+v %v", conversations, err)
	}
}

func TestBatchDeletesSupersededConversation(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.jsonl")
	e := NewJSONLExporter()
	if err := e.UpsertJSONL(&models.Conversation{ID: "old"}, path); err != nil {
		t.Fatal(err)
	}
	e.BeginBatch()
	e.DeleteJSONL("old")
	if err := e.UpsertJSONL(&models.Conversation{ID: "new"}, path); err != nil {
		t.Fatal(err)
	}
	if err := e.Flush(path); err != nil {
		t.Fatal(err)
	}
	conversations, err := readJSONL(path)
	if err != nil || len(conversations) != 1 || conversations[0].ID != "new" {
		t.Fatalf("superseded record retained: %+v %v", conversations, err)
	}
}

func TestDatasetRecoversWindowsStyleBackupBeforeMerge(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.jsonl")
	e := NewJSONLExporter()
	if err := e.UpsertJSONL(&models.Conversation{ID: "old"}, path); err != nil {
		t.Fatal(err)
	}
	backup := backupPath(path)
	if err := os.MkdirAll(filepath.Dir(backup), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(path, backup); err != nil {
		t.Fatal(err)
	}
	if err := e.UpsertJSONL(&models.Conversation{ID: "new"}, path); err != nil {
		t.Fatal(err)
	}
	conversations, err := readJSONL(path)
	if err != nil || len(conversations) != 2 {
		t.Fatalf("backup record was lost: %+v %v", conversations, err)
	}
}

func BenchmarkDatasetImport(b *testing.B) {
	for _, batch := range []bool{false, true} {
		name := "PerConversation"
		if batch {
			name = "Batched"
		}
		b.Run(name, func(b *testing.B) {
			dir := b.TempDir()
			convs := make([]*models.Conversation, 100)
			for i := range convs {
				convs[i] = &models.Conversation{ID: fmt.Sprint(i), Messages: []models.Message{{Role: "assistant", Content: strings.Repeat("A useful explanation with code. ", 500)}}}
			}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				path := filepath.Join(dir, fmt.Sprintf("dataset-%d.jsonl", n))
				e := NewJSONLExporter()
				if batch {
					e.BeginBatch()
				}
				for _, conv := range convs {
					if err := e.UpsertJSONL(conv, path); err != nil {
						b.Fatal(err)
					}
				}
				if batch {
					if err := e.Flush(path); err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}
