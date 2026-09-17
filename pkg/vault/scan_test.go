package vault

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"llm-context-vault/pkg/extractor"
	"llm-context-vault/pkg/models"
	"llm-context-vault/pkg/sanitizer"
)

type changingExtractor struct{}

func (changingExtractor) Name() string          { return "aider" }
func (changingExtractor) CanHandle(string) bool { return true }
func (changingExtractor) Extract(path string) (*models.Conversation, error) {
	if err := os.WriteFile(path, []byte("changed while reading"), 0600); err != nil {
		return nil, err
	}
	return &models.Conversation{SourceTool: "aider", Title: "Session expiration", Messages: []models.Message{
		{Role: "user", Content: "Explain session expiration handling."},
		{Role: "assistant", Content: "Refresh the token before session expiration."},
	}}, nil
}

func writeAGY(t testing.TB, root, id, answer string) string {
	t.Helper()
	path := filepath.Join(root, id, ".system_generated", "logs", "transcript.jsonl")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	var lines []string
	for _, step := range []map[string]string{{"type": "USER_INPUT", "content": "How do session expiration and refresh tokens interact?"}, {"type": "PLANNER_RESPONSE", "content": answer}} {
		data, _ := json.Marshal(step)
		lines = append(lines, string(data))
	}
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestIncrementalScanRecoveryAndInvalidation(t *testing.T) {
	base, source := t.TempDir(), t.TempDir()
	writeAGY(t, source, "one", "Session expiration uses a refresh token and an explicit timeout. SecretClient")
	sources := []ScanSource{{Tool: "antigravity", Path: source}}
	newVault := func() *Vault { return New(base, sanitizer.New(sanitizer.DefaultConfig())) }
	scan := func(v *Vault) ToolScanReport {
		t.Helper()
		reports, err := v.ScanSources(sources)
		if err != nil || len(reports) != 1 || reports[0].FailedCount != 0 {
			t.Fatalf("scan failed: %+v %v", reports, err)
		}
		return reports[0]
	}
	if r := scan(newVault()); r.ImportedCount != 1 {
		t.Fatalf("cold scan: %+v", r)
	}
	dataset := filepath.Join(base, "conversations", "dataset.jsonl")
	before, _ := os.ReadFile(dataset)
	beforeStamp, _ := stamp(dataset)
	if r := scan(newVault()); r.UnchangedCount != 1 || r.ImportedCount != 0 {
		t.Fatalf("repeat scan: %+v", r)
	}
	afterStamp, _ := stamp(dataset)
	if beforeStamp != afterStamp {
		t.Fatal("unchanged scan rewrote dataset")
	}
	v := newVault()
	v.ScanOptions.Full = true
	if r := scan(v); r.ImportedCount != 1 || r.UnchangedCount != 0 {
		t.Fatalf("full scan: %+v", r)
	}
	after, _ := os.ReadFile(dataset)
	if string(before) != string(after) {
		t.Fatal("full scan changed unchanged content")
	}
	files, _ := os.ReadDir(filepath.Join(base, "conversations", "markdown"))
	if err := os.Remove(filepath.Join(base, "conversations", "markdown", files[0].Name())); err != nil {
		t.Fatal(err)
	}
	if r := scan(newVault()); r.ImportedCount != 1 {
		t.Fatalf("missing export not restored: %+v", r)
	}
	writeAGY(t, source, "two", "Refresh tokens renew the session before expiration; handle revocation explicitly.")
	if r := scan(newVault()); r.NewCount != 1 || r.UnchangedCount != 1 {
		t.Fatalf("new session: %+v", r)
	}
	writeAGY(t, source, "one", "Updated session expiration explanation: refresh tokens must rotate and renew. SecretClient")
	if r := scan(newVault()); r.ChangedCount != 1 || r.UnchangedCount != 1 {
		t.Fatalf("changed session: %+v", r)
	}
	assertExportCount(t, base, 2)
	cfg := sanitizer.DefaultConfig()
	cfg.CustomKeywords = []string{"SecretClient"}
	if err := os.Remove(filepath.Join(base, ".vault", "scan-state.json")); err != nil {
		t.Fatal(err)
	}
	if r := scan(New(base, sanitizer.New(cfg))); r.ImportedCount != 2 {
		t.Fatalf("redaction settings did not invalidate cache: %+v", r)
	}
	assertExportCount(t, base, 2)
	findings, err := New(base, sanitizer.New(cfg)).AuditAll()
	if err != nil || len(findings) != 0 {
		t.Fatalf("redaction left stale export: %v %v", findings, err)
	}
	if err := os.Remove(dataset); err != nil {
		t.Fatal(err)
	}
	if r := scan(newVault()); r.ImportedCount != 2 {
		t.Fatalf("missing dataset not rebuilt: %+v", r)
	}
}

func assertExportCount(t *testing.T, base string, want int) {
	t.Helper()
	for _, dir := range []string{"markdown", "sharegpt"} {
		entries, err := os.ReadDir(filepath.Join(base, "conversations", dir))
		if err != nil || len(entries) != want {
			t.Fatalf("%s exports=%d want=%d err=%v", dir, len(entries), want, err)
		}
	}
	if got := countJSONL(t, filepath.Join(base, "conversations", "dataset.jsonl")); got != want {
		t.Fatalf("dataset records=%d want=%d", got, want)
	}
}

func TestScanFailedFlushDoesNotCheckpoint(t *testing.T) {
	base, source := t.TempDir(), t.TempDir()
	writeAGY(t, source, "one", "The session expiration policy needs a refresh timeout.")
	if err := os.MkdirAll(filepath.Join(base, "conversations"), 0755); err != nil {
		t.Fatal(err)
	}
	dataset := filepath.Join(base, "conversations", "dataset.jsonl")
	if err := os.WriteFile(dataset, []byte("invalid\n"), 0600); err != nil {
		t.Fatal(err)
	}
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	_, err := v.ScanSources([]ScanSource{{Tool: "antigravity", Path: source}})
	if err == nil {
		t.Fatal("expected merge failure")
	}
	if _, err := os.Stat(filepath.Join(base, ".vault", "scan-state.json")); !os.IsNotExist(err) {
		t.Fatal("failed merge wrote checkpoint")
	}
	data, _ := os.ReadFile(dataset)
	if string(data) != "invalid\n" {
		t.Fatal("failed scan damaged dataset")
	}
}

func TestScanWriterLock(t *testing.T) {
	v := New(t.TempDir(), sanitizer.New(sanitizer.DefaultConfig()))
	release, err := v.lockWriter()
	if err != nil {
		t.Fatal(err)
	}
	defer release()
	if _, err := v.ScanSources(nil); err == nil {
		t.Fatal("second writer acquired the same vault")
	}
}

func TestChangedSourceIsNotPersisted(t *testing.T) {
	base := t.TempDir()
	source := filepath.Join(t.TempDir(), "history.md")
	if err := os.WriteFile(source, []byte("initial"), 0600); err != nil {
		t.Fatal(err)
	}
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	v.extractors = []extractor.Extractor{changingExtractor{}}
	reports, err := v.ScanSources([]ScanSource{{Tool: "aider", Path: source}})
	if err != nil {
		t.Fatal(err)
	}
	if len(reports) != 1 || reports[0].FailedCount != 1 {
		t.Fatalf("concurrent modification was not reported: %+v", reports)
	}
	for _, path := range []string{
		filepath.Join(base, "conversations", "dataset.jsonl"),
		filepath.Join(base, "conversations", "markdown"),
		filepath.Join(base, "conversations", "sharegpt"),
	} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("changed source was persisted at %s", path)
		}
	}
}

func BenchmarkIncrementalScan(b *testing.B) {
	base, source := b.TempDir(), b.TempDir()
	for i := 0; i < 30; i++ {
		writeAGY(b, source, fmt.Sprint(i), strings.Repeat("Refresh tokens control session expiration. ", 50)+fmt.Sprint(i))
	}
	sources := []ScanSource{{Tool: "antigravity", Path: source}}
	v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	if _, err := v.ScanSources(sources); err != nil {
		b.Fatal(err)
	}
	for _, full := range []bool{true, false} {
		name := "Unchanged"
		if full {
			name = "Full"
		}
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				v := New(base, sanitizer.New(sanitizer.DefaultConfig()))
				v.ScanOptions.Full = full
				r, err := v.ScanSources(sources)
				if err != nil || r[0].FailedCount != 0 {
					b.Fatalf("%v %+v", err, r)
				}
			}
		})
	}
}

func TestCacheCorruptionAndDatasetChanges(t *testing.T) {
	base, source := t.TempDir(), t.TempDir()
	writeAGY(t, source, "one", "Session expiration requires a refresh token timeout.")
	sources := []ScanSource{{Tool: "antigravity", Path: source}}
	scan := func() ToolScanReport {
		t.Helper()
		r, err := New(base, sanitizer.New(sanitizer.DefaultConfig())).ScanSources(sources)
		if err != nil {
			t.Fatal(err)
		}
		return r[0]
	}
	scan()
	cache := filepath.Join(base, ".vault", "scan-state.json")
	if err := os.WriteFile(cache, []byte(`{"version":123}`), 0600); err != nil {
		t.Fatal(err)
	}
	if r := scan(); r.ImportedCount != 1 {
		t.Fatalf("bad cache wasn't rebuilt: %+v", r)
	}
	dataset := filepath.Join(base, "conversations", "dataset.jsonl")
	if err := os.Chtimes(dataset, time.Unix(123456789, 0), time.Unix(123456789, 0)); err != nil {
		t.Fatal(err)
	}
	if r := scan(); r.ImportedCount != 1 {
		t.Fatalf("external dataset modification didn't invalidate cache: %+v", r)
	}
}

func TestToolSpecificInvalidationDoesNotValidateUntouchedSources(t *testing.T) {
	base, first, second := t.TempDir(), t.TempDir(), t.TempDir()
	writeAGY(t, first, "one", "Session expiration uses refresh token rotation. SecretClient")
	writeAGY(t, second, "two", "Cache expiration uses an explicit retention timeout.")
	sources := []ScanSource{{Tool: "antigravity", Path: first}, {Tool: "antigravity", Path: second}}
	if _, err := New(base, sanitizer.New(sanitizer.DefaultConfig())).ScanSources(sources); err != nil {
		t.Fatal(err)
	}
	cfg := sanitizer.DefaultConfig()
	cfg.CustomKeywords = []string{"SecretClient"}
	if reports, err := New(base, sanitizer.New(cfg)).ScanSources(sources[:1]); err != nil || reports[0].ImportedCount != 1 {
		t.Fatalf("tool-specific invalidation failed: %+v %v", reports, err)
	}
	reports, err := New(base, sanitizer.New(cfg)).ScanSources(sources)
	if err != nil || reports[0].UnchangedCount != 1 || reports[1].ImportedCount != 1 {
		t.Fatalf("untouched source was incorrectly validated: %+v %v", reports, err)
	}
}

func TestFailedRetryPreservesOwnershipForRetirement(t *testing.T) {
	base := t.TempDir()
	source := filepath.Join(t.TempDir(), ".aider.chat.history.md")
	writeAider := func(answer string) {
		t.Helper()
		content := "#### USER\nExplain session expiration behavior.\n#### ASSISTANT\n" + answer + "\n"
		if err := os.WriteFile(source, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	inputs := []ScanSource{{Tool: "aider", Path: source}}
	writeAider("Refresh tokens renew the original session.")
	if _, err := New(base, sanitizer.New(sanitizer.DefaultConfig())).ScanSources(inputs); err != nil {
		t.Fatal(err)
	}
	writeAider("Refresh tokens must rotate before session expiration.")
	failing := New(base, sanitizer.New(sanitizer.DefaultConfig()))
	failing.extractors = []extractor.Extractor{changingExtractor{}}
	if reports, err := failing.ScanSources(inputs); err != nil || reports[0].FailedCount != 1 {
		t.Fatalf("expected failed retry: %+v %v", reports, err)
	}
	writeAider("Refresh tokens must rotate before session expiration.")
	if reports, err := New(base, sanitizer.New(sanitizer.DefaultConfig())).ScanSources(inputs); err != nil || reports[0].ChangedCount != 1 {
		t.Fatalf("successful retry failed: %+v %v", reports, err)
	}
	assertExportCount(t, base, 1)
}
