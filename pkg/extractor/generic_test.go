package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenericExtractorExtractAllJSONL(t *testing.T) {
	path := filepath.Join(t.TempDir(), "conversations.jsonl")
	data := "{\"id\":\"one\",\"source_tool\":\"test\",\"messages\":[{\"role\":\"user\",\"content\":\"first request\"}]}\n" +
		"{\"id\":\"two\",\"source_tool\":\"test\",\"messages\":[{\"role\":\"user\",\"content\":\"second request\"}]}\n"
	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	conversations, err := NewGenericExtractor().ExtractAll(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(conversations) != 2 || conversations[0].ID != "one" || conversations[1].ID != "two" {
		t.Fatalf("unexpected conversations: %#v", conversations)
	}
}

func TestGenericExtractorReportsJSONLLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "invalid.jsonl")
	if err := os.WriteFile(path, []byte("{bad}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := NewGenericExtractor().ExtractAll(path); err == nil {
		t.Fatal("expected malformed JSONL error")
	}
}
