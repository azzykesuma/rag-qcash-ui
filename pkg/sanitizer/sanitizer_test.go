package sanitizer

import (
	"strings"
	"testing"

	"llm-context-vault/pkg/models"
)

func TestSanitizeConversationSanitizesMetadata(t *testing.T) {
	s := New(DefaultConfig())
	conversation := &models.Conversation{
		Title:    "test",
		Metadata: map[string]string{"owner": "dev@example.com", "path": `C:\Users\alice\repo`},
		Messages: []models.Message{{
			Role: "user",
			Metadata: map[string]any{
				"contact": "dev@example.com",
				"nested":  map[string]any{"ip": "10.20.30.40"},
			},
		}},
	}

	clean := s.SanitizeConversation(conversation)
	if strings.Contains(clean.Metadata["owner"], "example.com") || strings.Contains(clean.Metadata["path"], "alice") {
		t.Fatalf("conversation metadata was not sanitized: %#v", clean.Metadata)
	}
	nested := clean.Messages[0].Metadata["nested"].(map[string]any)
	if clean.Messages[0].Metadata["contact"] != "[REDACTED_EMAIL]" || nested["ip"] != "[REDACTED_IP]" {
		t.Fatalf("message metadata was not sanitized: %#v", clean.Messages[0].Metadata)
	}
}

func TestAuditTextDetectsPublicIPButNotLoopback(t *testing.T) {
	s := New(DefaultConfig())
	if len(s.AuditText("service at 10.20.30.40")) == 0 {
		t.Fatal("expected public IP audit finding")
	}
	if warnings := s.AuditText("http://127.0.0.1:8080"); len(warnings) != 0 {
		t.Fatalf("unexpected loopback finding: %v", warnings)
	}
}
