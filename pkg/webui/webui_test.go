package webui

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDashboardRoutes(t *testing.T) {
	dir := t.TempDir()
	mdDir := filepath.Join(dir, "conversations", "markdown")
	if err := os.MkdirAll(mdDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := "# Example session\n\n- **Source Tool**: `opencode`\n- **Date**: `2026-09-14`\n\n---\n\nUnique search phrase\n<script>alert('untrusted')</script>"
	if err := os.WriteFile(filepath.Join(mdDir, "example.md"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	handler, err := NewHandler(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		path     string
		status   int
		contains string
	}{
		{"/", 200, "LLM Context Vault"},
		{"/api/sessions", 200, `"tool":"opencode"`},
		{"/api/conversation?file=example.md", 200, content},
		{"/api/search?q=Unique", 200, "example"},
		{"/api/context?q=Unique", 200, "Reference Session"},
		{"/api/search?q=does-not-exist", 200, "null"},
		{"/api/audit", 200, "{}"},
		{"/api/conversation?file=../outside.md", 400, "Invalid"},
		{"/api/conversation?file=..%5Coutside.md", 400, "Invalid"},
		{"/api/conversation?file=missing.md", 404, "not found"},
		{"/unknown", 404, "not found"},
	} {
		t.Run(tc.path, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, httptest.NewRequest("GET", tc.path, nil))
			if rr.Code != tc.status || !strings.Contains(rr.Body.String(), tc.contains) {
				t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
			}
			if strings.HasPrefix(tc.path, "/api/conversation?file=example") && rr.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
				t.Fatal("conversation must be served as plain text")
			}
		})
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, httptest.NewRequest("POST", "/api/audit", nil))
	if rr.Code != 405 {
		t.Fatalf("unexpected POST status: %d", rr.Code)
	}
}

func TestMissingVault(t *testing.T) {
	if _, err := NewHandler(t.TempDir()); err == nil {
		t.Fatal("expected missing conversations error")
	}
}

func TestInvalidPort(t *testing.T) {
	for _, port := range []int{-1, 65536} {
		if err := Run(t.TempDir(), port, false); err == nil {
			t.Fatalf("expected invalid port error for %d", port)
		}
	}
}

func TestSessionsUseLocalProjectAliases(t *testing.T) {
	dir := t.TempDir()
	mdDir := filepath.Join(dir, "conversations", "markdown")
	if err := os.MkdirAll(mdDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(mdDir, "example.md"), []byte("# Example\n\n- **Project**: `old-name`\n\n---\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, ".vault"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".vault", "projects.json"), []byte(`{"example.md":"preferred-name"}`), 0600); err != nil {
		t.Fatal(err)
	}
	handler, err := NewHandler(dir)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, httptest.NewRequest("GET", "/api/sessions", nil))
	if rr.Code != 200 || !strings.Contains(rr.Body.String(), `"project":"preferred-name"`) {
		t.Fatalf("alias not applied: status=%d body=%s", rr.Code, rr.Body.String())
	}
}
