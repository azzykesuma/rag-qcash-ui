package benchmark

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"llm-context-vault/pkg/sanitizer"
	"llm-context-vault/pkg/vault"
)

func TestEvaluateKeywords(t *testing.T) {
	text := "This implementation uses HostAuthGate and checks isAuthoritiesReady before loading /auth/me."
	expected := []string{"HostAuthGate", "isAuthoritiesReady", "/auth/me", "/menu/me"}
	negative := []string{"NextAuth", "Django"}

	matched, score := evaluateKeywords(text, expected, negative)
	if len(matched) != 3 {
		t.Fatalf("expected 3 matched keywords, got %d: %v", len(matched), matched)
	}
	expectedScore := 3.0 / 4.0
	if score != expectedScore {
		t.Fatalf("expected score %f, got %f", expectedScore, score)
	}

	// Test negative penalty
	textWithNeg := text + " We also integrated NextAuth."
	_, scoreWithNeg := evaluateKeywords(textWithNeg, expected, negative)
	if scoreWithNeg >= score {
		t.Fatalf("negative keyword should reduce score, got %f vs %f", scoreWithNeg, score)
	}
}

func TestEstimateTokens(t *testing.T) {
	sample := "Hello world, this is a test of token estimation."
	tokens := estimateTokens(sample)
	if tokens < 8 || tokens > 20 {
		t.Fatalf("unexpected token count %d for sample %q", tokens, sample)
	}
}

func TestLoadTasks(t *testing.T) {
	tempDir := t.TempDir()
	taskFile := filepath.Join(tempDir, "tasks.json")
	content := `[
		{
			"id": "t1",
			"query": "How to optimize?",
			"ground_truth": "Use incremental cache.",
			"expected_keywords": ["incremental", "cache"]
		}
	]`
	if err := os.WriteFile(taskFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write tasks file: %v", err)
	}

	tasks, err := LoadTasks(taskFile)
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}
	if len(tasks) != 1 || tasks[0].ID != "t1" {
		t.Fatalf("unexpected tasks parsed: %+v", tasks)
	}
}

func TestOfflineBenchmarkExecution(t *testing.T) {
	tempDir := t.TempDir()
	convDir := filepath.Join(tempDir, "conversations", "markdown")
	if err := os.MkdirAll(convDir, 0755); err != nil {
		t.Fatalf("failed to create conv dir: %v", err)
	}

	// Seed a test markdown session
	mdSession := `# Test Session
## Turn 1: User
How is authentication handled in standalone mode?
## Turn 2: Assistant
HostAuthGate owns /auth/me and page readiness uses isAuthoritiesReady.
`
	if err := os.WriteFile(filepath.Join(convDir, "auth-test.md"), []byte(mdSession), 0644); err != nil {
		t.Fatalf("failed to write test session: %v", err)
	}

	v := vault.New(tempDir, sanitizer.New(sanitizer.DefaultConfig()))
	if err := v.RefreshSearchIndex(); err != nil {
		t.Fatalf("failed to refresh index: %v", err)
	}

	taskFile := filepath.Join(tempDir, "tasks.json")
	taskJSON := `[
		{
			"id": "auth-gate",
			"query": "authentication standalone",
			"ground_truth": "HostAuthGate owns /auth/me and page readiness uses isAuthoritiesReady.",
			"expected_keywords": ["HostAuthGate", "isAuthoritiesReady"]
		}
	]`
	if err := os.WriteFile(taskFile, []byte(taskJSON), 0644); err != nil {
		t.Fatalf("failed to write task file: %v", err)
	}

	reportFile := filepath.Join(tempDir, "report.md")
	jsonFile := filepath.Join(tempDir, "results.json")

	cfg := RunnerConfig{
		Vault:      v,
		TasksFile:  taskFile,
		ReportFile: reportFile,
		JSONFile:   jsonFile,
		RunnerType: "offline",
		Timeout:    5 * time.Second,
	}

	runner, err := NewRunner(cfg)
	if err != nil {
		t.Fatalf("NewRunner failed: %v", err)
	}

	summary, err := runner.Run(context.Background())
	if err != nil {
		t.Fatalf("Runner.Run failed: %v", err)
	}

	if summary.TotalTasks != 1 || summary.SuccessfulTasks != 1 {
		t.Fatalf("expected 1 successful task, got %d/%d", summary.SuccessfulTasks, summary.TotalTasks)
	}

	if err := SaveReport(summary, reportFile, jsonFile); err != nil {
		t.Fatalf("SaveReport failed: %v", err)
	}

	repBytes, err := os.ReadFile(reportFile)
	if err != nil {
		t.Fatalf("failed to read report file: %v", err)
	}
	reportContent := string(repBytes)
	if !strings.Contains(reportContent, "LLM Context Vault: Benchmark Report") {
		t.Fatalf("report missing title: %s", reportContent)
	}
	if !strings.Contains(reportContent, "auth-gate") {
		t.Fatalf("report missing task ID: %s", reportContent)
	}
}
