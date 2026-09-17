package vault

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"llm-context-vault/pkg/models"
	"llm-context-vault/pkg/sanitizer"
)

func TestConversationTopicUsesWholeConversation(t *testing.T) {
	conv := &models.Conversation{Title: "hello", Messages: []models.Message{
		{Role: "user", Content: "hello"},
		{Role: "assistant", Content: "Hello! How can I help?"},
		{Role: "user", Content: "Investigate session expiration after refreshing tokens."},
		{Role: "assistant", Content: "Session expiration occurs because the refresh token is not rotated."},
		{Role: "user", Content: "Explain the session expiration timeout."},
	}}
	name := conversationTopic(conv)
	if !strings.Contains(name, "session-expiration") || strings.Contains(name, "hello") {
		t.Fatalf("unhelpful topic: %s", name)
	}
	conv.Title = "Session expiration troubleshooting"
	if name := conversationTopic(conv); !strings.Contains(name, "session-expiration") {
		t.Fatal(name)
	}
	conv.Title = "https://internal-service.example.com/123456789"
	if name := conversationTopic(conv); strings.Contains(name, "https") || strings.Contains(name, "123456789") {
		t.Fatal(name)
	}
}

func TestExistingFilenamesSurviveCacheLoss(t *testing.T) {
	base := t.TempDir()
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	conv := conversation("one", "Hello, could you investigate the session expiration policy?")
	stored, _, err := v.StoreConversation(conv)
	if err != nil {
		t.Fatal(err)
	}
	files, _ := os.ReadDir(filepath.Join(base, "conversations", "markdown"))
	oldSlug := strings.TrimSuffix(files[0].Name(), ".md")
	legacySlug := "test_hello-gibberish-existing"
	for dir, ext := range map[string]string{"markdown": ".md", "sharegpt": ".json"} {
		if err := os.Rename(filepath.Join(base, "conversations", dir, oldSlug+ext), filepath.Join(base, "conversations", dir, legacySlug+ext)); err != nil {
			t.Fatal(err)
		}
	}
	v = New(base, sanitizer.New(sanitizer.DefaultConfig()))
	if _, _, err = v.StoreConversation(conv); err != nil {
		t.Fatal(err)
	}
	files, _ = os.ReadDir(filepath.Join(base, "conversations", "markdown"))
	if len(files) != 1 || files[0].Name() != legacySlug+".md" {
		t.Fatalf("renamed existing export: %+v", files)
	}
	if v.exportNames[stored.ID].Slug != legacySlug {
		t.Fatal("lost existing ID association")
	}
}

func TestTopicIsWindowsSafeAndStable(t *testing.T) {
	conv := conversation("one", strings.Repeat("認証セッションの有効期限 ", 20))
	conv.Title = "https://example.com/a?token=abc"
	for i := 0; i < 10; i++ {
		name := conversationTopic(conv)
		if name != conversationTopic(conv) || !utf8.ValidString(name) || strings.ContainsAny(name, `<>:"/\|?*`) || len([]rune(name)) > 60 {
			t.Fatalf("unsafe filename %q", name)
		}
	}
}

func TestLanguageOrderingDoesNotChangeIdentity(t *testing.T) {
	conv := conversation("one", "Investigate authentication middleware retries")
	conv.Languages = []string{"go", "typescript"}
	a := stableConversationID(conv)
	conv.Languages = []string{"typescript", "go"}
	b := stableConversationID(conv)
	if a != b {
		t.Fatal("language ordering changed content identity")
	}
}

func TestProjectSeparatesOtherwiseIdenticalConversations(t *testing.T) {
	conv := conversation("one", "Investigate authentication middleware retries")
	conv.Project = "payments"
	first := stableConversationID(conv)
	conv.Project = "onboarding"
	if second := stableConversationID(conv); first == second {
		t.Fatal("different projects collapsed into one conversation identity")
	}
}
