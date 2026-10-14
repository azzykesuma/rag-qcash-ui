package webui

import (
	"encoding/json"
	"net/http"
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
		{"/api/sessions", 200, `"topic":"General development"`},
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

func TestSessionTopicUsesSpecificGroupsBeforeBugFallback(t *testing.T) {
	tests := []struct {
		title string
		file  string
		want  string
	}{
		{"Fix this Bamboo build error", "build.md", "Bamboo & CI fixes"},
		{"Fix failing unit tests", "tests.md", "Testing & coverage"},
		{"Resolve Sonar issue", "quality.md", "Sonar & code quality"},
		{"Fix the broken submit handler", "bug.md", "Bug fixing"},
		{"Discuss implementation approach", "general.md", "General development"},
		{"Build failed", "antigravity_bamboo-run.md", "Bamboo & CI fixes"},
	}
	for _, tc := range tests {
		t.Run(tc.want+tc.title, func(t *testing.T) {
			got := sessionTopic(session{Title: tc.title, File: tc.file})
			if got != tc.want {
				t.Fatalf("sessionTopic(%q, %q) = %q, want %q", tc.title, tc.file, got, tc.want)
			}
		})
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

func TestTopicOverridesAndInsights(t *testing.T) {
	dir := t.TempDir()
	mdDir := filepath.Join(dir, "conversations", "markdown")
	if err := os.MkdirAll(mdDir, 0755); err != nil {
		t.Fatal(err)
	}
	for name, body := range map[string]string{
		"first.md":  "# Fix bamboo build\n\n- **Source Tool**: `opencode`\n\n---\n\n## Turn 1: User\n\nWhy is my bamboo pipeline failing with docker?\n\n---\n\n## Turn 2: Assistant\n\nThe Docker stage may require a newer image.\n",
		"second.md": "# Bamboo docker pipeline\n\n---\n\n## Turn 1: User\n\nDocker pipeline failures in bamboo.\n",
	} {
		if err := os.WriteFile(filepath.Join(mdDir, name), []byte(body), 0644); err != nil {
			t.Fatal(err)
		}
	}
	handler, err := NewHandler(dir)
	if err != nil {
		t.Fatal(err)
	}
	call := func(method, path, body string) *httptest.ResponseRecorder {
		t.Helper()
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, httptest.NewRequest(method, path, strings.NewReader(body)))
		return rr
	}
	if rr := call("PUT", "/api/topic", `{"file":"first.md","topic":"Authentication"}`); rr.Code != 200 {
		t.Fatalf("topic update: %d %s", rr.Code, rr.Body.String())
	}
	if rr := call("GET", "/api/sessions", ""); rr.Code != 200 || !strings.Contains(rr.Body.String(), `"topic":"Authentication"`) {
		t.Fatalf("topic not applied: %d %s", rr.Code, rr.Body.String())
	}
	if data, err := os.ReadFile(filepath.Join(dir, ".vault", "topics.json")); err != nil || !strings.Contains(string(data), `"first.md":"Authentication"`) {
		t.Fatalf("topic not saved: %v %s", err, data)
	}
	for _, body := range []string{`{"file":"../first.md","topic":"Bug fixing"}`, `{"file":"first.md","topic":"<script>"}`} {
		if rr := call("PUT", "/api/topic", body); rr.Code != http.StatusBadRequest {
			t.Fatalf("accepted invalid topic: %s", rr.Body.String())
		}
	}
	if rr := call("PUT", "/api/topic", `{"file":"missing.md","topic":"Bug fixing"}`); rr.Code != 404 {
		t.Fatalf("expected missing file: %d", rr.Code)
	}
	rr := call("GET", "/api/insights?file=first.md", "")
	if rr.Code != 200 {
		t.Fatalf("insights: %d %s", rr.Code, rr.Body.String())
	}
	var info insights
	if err := json.Unmarshal(rr.Body.Bytes(), &info); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(info.Problem, "pipeline failing") || !strings.Contains(info.Finding, "Docker stage") || !strings.Contains(info.Outcome, "Unknown") || len(info.Related) == 0 || info.Related[0].File != "second.md" {
		t.Fatalf("unexpected insights: %+v", info)
	}
	if rr := call("GET", "/api/insights?file=..%5Cfirst.md", ""); rr.Code != 400 {
		t.Fatalf("expected invalid file: %d", rr.Code)
	}
	confirmed := "# Fixed a login issue\n\n---\n\n## Turn 1: User\n\nPlease fix login.\n\n## Turn 2: Assistant\n\nTry clearing the cookie.\n\n## Turn 3: User\n\nIt works now, thank you.\n"
	if err := os.WriteFile(filepath.Join(mdDir, "confirmed.md"), []byte(confirmed), 0644); err != nil {
		t.Fatal(err)
	}
	rr = call("GET", "/api/insights?file=confirmed.md", "")
	if rr.Code != 200 || !strings.Contains(rr.Body.String(), "Confirmed by user") {
		t.Fatalf("confirmed outcome: %d %s", rr.Code, rr.Body.String())
	}
	if rr := call("PUT", "/api/topic", `{"file":"first.md","topic":""}`); rr.Code != 200 {
		t.Fatalf("reset topic: %d %s", rr.Code, rr.Body.String())
	}
	if rr := call("GET", "/api/sessions", ""); !strings.Contains(rr.Body.String(), `"topic":"Bamboo \u0026 CI fixes"`) {
		t.Fatalf("topic reset failed: %s", rr.Body.String())
	}
}
