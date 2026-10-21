package benchmark

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateMarkdownReport formats the benchmark summary into a structured markdown document
func GenerateMarkdownReport(s *BenchmarkSummary) string {
	var b strings.Builder

	b.WriteString("# 🧠 LLM Context Vault: Benchmark Report\n\n")
	b.WriteString(fmt.Sprintf("**Date:** %s  \n", s.Timestamp))
	b.WriteString(fmt.Sprintf("**Runner Backend:** `%s`  \n", s.RunnerName))
	if s.Model != "" {
		b.WriteString(fmt.Sprintf("**Model:** `%s`  \n", s.Model))
	}
	b.WriteString(fmt.Sprintf("**Tasks Evaluated:** %d (Successful: %d)  \n\n", s.TotalTasks, s.SuccessfulTasks))
	b.WriteString(fmt.Sprintf("**Repeats:** %d (each task; odd runs baseline first, even runs vault first)  \n\n", s.Repeats))
	b.WriteString("**Scope:** Blind question versus question with retrieved vault context. No repository exploration or tool calls are measured. Keyword coverage is a proxy, not verified factual accuracy. Token counts may be estimated by the backend. Offline mode uses fixed synthetic answers and is only a pipeline smoke test.\n\n")

	b.WriteString("---\n\n")
	b.WriteString("## 📊 Executive Scorecard\n\n")

	b.WriteString("| Metric | Baseline (Arm A) | Vault-Assisted (Arm B) | Improvement / Impact |\n")
	b.WriteString("| :--- | :--- | :--- | :--- |\n")
	if s.SuccessfulTasks == 0 {
		b.WriteString("| No successful paired runs | — | — | — |\n")
	} else {
		b.WriteString(fmt.Sprintf("| **Average Tokens / Task** | `%.0f` tokens | `%.0f` tokens | **%+.1f%%** tokens |\n",
			s.AvgBaselineTokens, s.AvgVaultTokens, s.TokenChangePct))
		b.WriteString(fmt.Sprintf("| **Average Generation Time (retrieval excluded)** | `%.0f ms` | `%.0f ms` | **%.2fx ratio** |\n",
			s.AvgBaselineDurationMs, s.AvgVaultDurationMs, s.GenerationTimeRatio))
		b.WriteString(fmt.Sprintf("| **Keyword Coverage (proxy)** | `%.1f%%` | `%.1f%%` | **%+.1f%% relative change** |\n",
			s.AvgBaselineKeywordScore*100.0, s.AvgVaultKeywordScore*100.0, s.KeywordCoverageChangePct))
	}

	if s.AvgVaultJudgeScore > 0 {
		judgeDiff := s.AvgVaultJudgeScore - s.AvgBaselineJudgeScore
		b.WriteString(fmt.Sprintf("| **LLM Judge Score (1–5)** | `%.2f / 5.0` | `%.2f / 5.0` | **%+.2f points** |\n",
			s.AvgBaselineJudgeScore, s.AvgVaultJudgeScore, judgeDiff))
	}

	b.WriteString("\n> [!NOTE]\n")
	b.WriteString("> A failed task is excluded from averages. Retrieval latency is shown separately and is not included in the generation-time ratio. For end-to-end latency, add retrieval to the vault arm. No tool-use or exploration savings are inferred.\n\n")

	b.WriteString("---\n\n")
	b.WriteString("## 📋 Per-Task Comparison Table\n\n")

	b.WriteString("| Task ID | Run | Project | FTS Latency | Baseline Tokens | Vault Tokens | Baseline Keywords | Vault Keywords | Δ Score |\n")
	b.WriteString("| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n")

	for _, r := range s.Results {
		status := "✅"
		if r.Error != "" {
			status = "⚠️ Error"
		}
		proj := r.Task.Project
		if proj == "" {
			proj = "general"
		}

		baseScore, vaultScore, diff := fmt.Sprintf("%.0f%%", r.Baseline.KeywordScore*100), fmt.Sprintf("%.0f%%", r.VaultAssisted.KeywordScore*100), fmt.Sprintf("%+.0f%%", r.ScoreDiff*100)
		if r.Error != "" {
			baseScore, vaultScore, diff = "—", "—", "—"
		}
		b.WriteString(fmt.Sprintf("| `%s` %s | `%d` | `%s` | `%d ms` | `%d` | `%d` | `%s` | `%s` | `%s` |\n",
			r.Task.ID, status, r.Repeat, proj, r.ContextRetrievalMs,
			r.Baseline.TotalTokens, r.VaultAssisted.TotalTokens,
			baseScore, vaultScore, diff))
	}

	b.WriteString("\n---\n\n")
	b.WriteString("## 🔍 Detailed Task Breakdown\n\n")

	for i, r := range s.Results {
		b.WriteString(fmt.Sprintf("### Task %d: `%s`\n\n", i+1, r.Task.ID))
		if r.Error != "" {
			b.WriteString(fmt.Sprintf("**Failed:** %s  \n", r.Error))
		}
		b.WriteString(fmt.Sprintf("**Query:** *\"%s\"*  \n", r.Task.Query))
		b.WriteString(fmt.Sprintf("**Ground Truth:** `%s`  \n\n", r.Task.GroundTruth))

		b.WriteString(fmt.Sprintf("- **Vault SQLite FTS Retrieval:** `%d ms` (%d chars snippet injected)\n",
			r.ContextRetrievalMs, r.ContextChars))
		b.WriteString(fmt.Sprintf("- **Tokens:** Baseline = `%d` (duration `%d ms`), Vault-Assisted = `%d` (duration `%d ms`)\n",
			r.Baseline.TotalTokens, r.Baseline.DurationMs,
			r.VaultAssisted.TotalTokens, r.VaultAssisted.DurationMs))
		b.WriteString(fmt.Sprintf("- **Matched Keywords:** Baseline = `%s` | Vault = `%s`\n",
			strings.Join(r.Baseline.KeywordsMatched, ", "),
			strings.Join(r.VaultAssisted.KeywordsMatched, ", ")))

		if r.Baseline.JudgeScore > 0 || r.VaultAssisted.JudgeScore > 0 {
			b.WriteString(fmt.Sprintf("- **LLM Judge Score:** Baseline = `%.1f / 5` | Vault = `%.1f / 5`\n",
				r.Baseline.JudgeScore, r.VaultAssisted.JudgeScore))
			if r.VaultAssisted.JudgeFeedback != "" {
				b.WriteString(fmt.Sprintf("- **Judge Feedback:** *%s*\n", r.VaultAssisted.JudgeFeedback))
			}
		}

		b.WriteString("\n<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>\n\n")
		b.WriteString("#### Baseline Response (Arm A):\n```\n")
		b.WriteString(truncateText(r.Baseline.Response, 1500))
		b.WriteString("\n```\n\n")
		b.WriteString("#### Vault-Assisted Response (Arm B):\n```\n")
		b.WriteString(truncateText(r.VaultAssisted.Response, 1500))
		b.WriteString("\n```\n</details>\n\n")
	}

	return b.String()
}

// SaveReport writes the markdown report and optional JSON summary to disk
func SaveReport(s *BenchmarkSummary, mdPath, jsonPath string) error {
	if mdPath != "" {
		if err := os.MkdirAll(filepath.Dir(mdPath), 0755); err != nil {
			return err
		}
		mdContent := GenerateMarkdownReport(s)
		if err := os.WriteFile(mdPath, []byte(mdContent), 0644); err != nil {
			return err
		}
	}

	if jsonPath != "" {
		if err := os.MkdirAll(filepath.Dir(jsonPath), 0755); err != nil {
			return err
		}
		jsonData, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
			return err
		}
	}

	return nil
}

func truncateText(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes]) + "\n[...truncated]"
}
