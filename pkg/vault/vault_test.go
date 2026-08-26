package vault

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"llm-context-vault/pkg/models"
	"llm-context-vault/pkg/sanitizer"
)

func TestStoreConversationIsIdempotentAndAvoidsSlugCollisions(t *testing.T) {
	baseDir := t.TempDir()
	v := New(baseDir, sanitizer.New(sanitizer.DefaultConfig()))
	first := conversation("first", "How do I configure a resilient service?")
	second := conversation("second", "How do I configure a resilient service with retries?")

	storedFirst, _, err := v.StoreConversation(first)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := v.StoreConversation(first); err != nil {
		t.Fatal(err)
	}
	storedSecond, _, err := v.StoreConversation(second)
	if err != nil {
		t.Fatal(err)
	}
	if storedFirst.ID == storedSecond.ID {
		t.Fatal("different conversation content received the same stable ID")
	}

	entries, err := os.ReadDir(filepath.Join(baseDir, "conversations", "markdown"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected two unique markdown exports, got %d", len(entries))
	}
	if countJSONL(t, filepath.Join(baseDir, "conversations", "dataset.jsonl")) != 2 {
		t.Fatal("dataset was not deduplicated by stable ID")
	}
}

func TestStoreConversationAuditsAllSanitizedFields(t *testing.T) {
	v := New(t.TempDir(), sanitizer.New(sanitizer.DefaultConfig()))
	stored, warnings, err := v.StoreConversation(&models.Conversation{
		SourceTool: "test",
		Title:      "metadata audit",
		Messages: []models.Message{
			{Role: "user", Content: "Please review this implementation thoroughly."},
			{Role: "assistant", Content: "The implementation is ready for review."},
		},
		Metadata: map[string]string{"owner": "dev@example.com"},
	})
	if err != nil || stored == nil {
		t.Fatalf("failed to store sanitized conversation: %v", err)
	}
	if len(warnings) != 0 {
		t.Fatalf("sanitized conversation still has audit warnings: %v", warnings)
	}
}

func TestAuditAllIgnoresNonPublishableFiles(t *testing.T) {
	baseDir := t.TempDir()
	conversationDir := filepath.Join(baseDir, "conversations")
	if err := os.MkdirAll(conversationDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(conversationDir, "index.db"), []byte("ghp_abcdefghijklmnopqrstuvwxyz0123456789AB"), 0644); err != nil {
		t.Fatal(err)
	}

	findings, err := New(baseDir, sanitizer.New(sanitizer.DefaultConfig())).AuditAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 0 {
		t.Fatalf("unexpected audit findings for ignored binary artifact: %v", findings)
	}
}

func conversation(id, prompt string) *models.Conversation {
	return &models.Conversation{
		ID:         id,
		SourceTool: "test",
		Title:      prompt,
		Messages: []models.Message{
			{Role: "user", Content: prompt},
			{Role: "assistant", Content: "Use a retry policy and an explicit timeout."},
		},
	}
}

func countJSONL(t *testing.T, path string) int {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var conversation models.Conversation
		if err := json.Unmarshal(scanner.Bytes(), &conversation); err != nil {
			t.Fatal(err)
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return count
}
