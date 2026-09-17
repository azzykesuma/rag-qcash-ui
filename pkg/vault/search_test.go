package vault

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"llm-context-vault/pkg/sanitizer"
)

func writeSearchDoc(t *testing.T, base, name, title, project, body string) string {
	t.Helper()
	dir := filepath.Join(base, "conversations", "markdown")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, name+".md")
	text := fmt.Sprintf("# %s\n\n- **Source Tool**: `opencode`\n- **Date**: `2026-09-14 12:00:00`\n- **Project**: `%s`\n\n---\n\n%s", title, project, body)
	if err := os.WriteFile(path, []byte(text), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRankedSearchContextAndFilters(t *testing.T) {
	base := t.TempDir()
	writeSearchDoc(t, base, "aaa-unrelated", "General expiration discussion", "other", "## Turn 1: User\n\nExplain expiration.\n\n## Turn 2: Assistant\n\nCache expiration is configured by TTL.")
	body := "## Turn 1: User\n\nWhy does the session expire early?\n\n## Turn 2: Assistant\n\nSession expiration occurs when refresh token rotation fails. Renew the token before expiry.\n\n```go\nrenewToken()\n```"
	path := writeSearchDoc(t, base, "zzz-relevant", "Session expiration troubleshooting", "example-project", body)
	writeSearchDoc(t, base, "duplicate", "Session expiration troubleshooting", "example-project", body)
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	results, err := v.SearchWithOptions("why session expiration", SearchOptions{})
	if err != nil || len(results) != 2 {
		t.Fatalf("results=%+v err=%v", results, err)
	}
	if !strings.Contains(results[0].Snippet, "renewToken") {
		t.Fatalf("solution did not rank first: %+v", results)
	}
	results, err = v.SearchWithOptions("expiration", SearchOptions{Project: "example-project", Tool: "opencode", After: "2026-09-01", Before: "2026-09-30"})
	if err != nil || len(results) != 1 {
		t.Fatalf("filters or deduplication failed: %+v %v", results, err)
	}
	content, _ := os.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	r := results[0]
	if r.FirstTurn != 1 || r.LastTurn != 2 || !strings.Contains(strings.Join(lines[r.StartLine-1:r.EndLine], "\n"), "renewToken") {
		t.Fatalf("bad source references: %+v", r)
	}
	context, err := v.ContextWithOptions("session expiration", SearchOptions{Project: "example-project"}, 900)
	if err != nil || utf8.RuneCountInString(context) > 900 || !strings.Contains(context, "Historical reference") || !strings.Contains(context, "turns 1–2") || !strings.Contains(context, "renewToken") {
		t.Fatalf("bad context: %s %v", context, err)
	}
	results, err = v.SearchWithOptions("session", SearchOptions{Before: "2025-01-01"})
	if err != nil || len(results) != 0 {
		t.Fatalf("date filter failed: %v %v", results, err)
	}
	if _, err = v.Search(`session OR " : * NOT`); err != nil {
		t.Fatalf("query syntax escaped incorrectly: %v", err)
	}
}

func TestSearchIndexReconcilesChangesDeletesAndAliases(t *testing.T) {
	base := t.TempDir()
	path := writeSearchDoc(t, base, "original", "Refresh policy", "", "## Turn 1: User\n\nExplain refresh policy.\n\n## Turn 2: Assistant\n\nRotate the refresh token.")
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	if _, err := v.Search("refresh"); err != nil {
		t.Fatal(err)
	}
	aliasPath := filepath.Join(base, ".vault", "projects.json")
	if err := os.WriteFile(aliasPath, []byte(`{"original.md":"legacy-project"}`), 0600); err != nil {
		t.Fatal(err)
	}
	r, err := v.SearchWithOptions("refresh", SearchOptions{Project: "legacy-project"})
	if err != nil || len(r) != 1 {
		t.Fatalf("alias failed: %v %v", r, err)
	}
	writeSearchDoc(t, base, "original", "Refresh policy", "", "## Turn 1: User\n\nExplain revocation.\n\n## Turn 2: Assistant\n\nRevocation invalidates credentials immediately.")
	r, err = v.Search("revocation")
	if err != nil || len(r) != 1 {
		t.Fatalf("index didn't update: %v %v", r, err)
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	writeSearchDoc(t, base, "pulled-from-git", "Revocation policy", "", "Imported revocation explanation from another machine.")
	r, err = v.Search("revocation")
	if err != nil || len(r) != 1 || r[0].ConversationID != "pulled-from-git" {
		t.Fatalf("index didn't reconcile: %v %v", r, err)
	}
	if err := os.Remove(filepath.Join(base, ".vault", "search.db")); err != nil {
		t.Fatal(err)
	}
	r, err = v.Search("revocation")
	if err != nil || len(r) != 1 {
		t.Fatalf("index didn't rebuild: %v %v", r, err)
	}
}

func TestContextBudgetAndValidation(t *testing.T) {
	base := t.TempDir()
	writeSearchDoc(t, base, "large", "Session expiration", "", "## Turn 1: User\n\nExplain session expiration.\n\n## Turn 2: Assistant\n\n"+strings.Repeat("Session expiration 認証 requires renewal. ", 200))
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	for _, budget := range []int{512, 900, 12000} {
		text, err := v.ContextWithOptions("session expiration", SearchOptions{}, budget)
		if err != nil || utf8.RuneCountInString(text) > budget || !utf8.ValidString(text) || !strings.Contains(text, "Source:") {
			t.Fatalf("budget %d failed: %d %v", budget, utf8.RuneCountInString(text), err)
		}
	}
	for _, opts := range []SearchOptions{{After: "yesterday"}, {Before: "2026-13-10"}, {After: "2026-09-01", Before: "2026-08-01"}, {Limit: 101}} {
		if _, err := v.SearchWithOptions("session", opts); err == nil {
			t.Fatalf("accepted invalid options: %+v", opts)
		}
	}
	if _, err := v.ContextWithOptions("session", SearchOptions{}, 10); err == nil {
		t.Fatal("accepted invalid budget")
	}
}
