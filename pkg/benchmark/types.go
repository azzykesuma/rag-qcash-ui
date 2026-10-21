package benchmark

import "time"

// Task represents a single benchmark query case with ground truth criteria
type Task struct {
	ID               string   `json:"id"`
	Project          string   `json:"project,omitempty"`
	Category         string   `json:"category,omitempty"` // e.g. "architecture", "bugfix", "pipeline", "negative_control"
	Query            string   `json:"query"`
	GroundTruth      string   `json:"ground_truth"`
	ExpectedKeywords []string `json:"expected_keywords,omitempty"`
	NegativeKeywords []string `json:"negative_keywords,omitempty"`
}

// ArmExecution holds the run metrics for one condition (Arm A or Arm B)
type ArmExecution struct {
	Prompt           string        `json:"prompt"`
	Response         string        `json:"response"`
	PromptTokens     int           `json:"prompt_tokens"`
	CompletionTokens int           `json:"completion_tokens"`
	TotalTokens      int           `json:"total_tokens"`
	Duration         time.Duration `json:"duration"`
	DurationMs       int64         `json:"duration_ms"`
	KeywordsMatched  []string      `json:"keywords_matched,omitempty"`
	KeywordScore     float64       `json:"keyword_score"` // 0.0 to 1.0
	JudgeScore       float64       `json:"judge_score"`   // 1.0 to 5.0 (if LLM judge enabled)
	JudgeFeedback    string        `json:"judge_feedback,omitempty"`
}

// TaskResult holds the comparative benchmark results for one task
type TaskResult struct {
	Task               Task         `json:"task"`
	Repeat             int          `json:"repeat"`
	ContextSnippet     string       `json:"context_snippet,omitempty"`
	ContextChars       int          `json:"context_chars"`
	ContextTokens      int          `json:"context_tokens"`
	ContextRetrievalMs int64        `json:"context_retrieval_ms"`
	Baseline           ArmExecution `json:"baseline"`       // Arm A: without vault context
	VaultAssisted      ArmExecution `json:"vault_assisted"` // Arm B: with vault context
	TokenDiff          int          `json:"token_diff"`     // VaultAssisted - Baseline
	DurationDiffMs     int64        `json:"duration_diff_ms"`
	ScoreDiff          float64      `json:"score_diff"`
	Error              string       `json:"error,omitempty"`
}

// BenchmarkSummary provides aggregate metrics across all evaluated tasks
type BenchmarkSummary struct {
	TotalTasks               int          `json:"total_tasks"`
	SuccessfulTasks          int          `json:"successful_tasks"`
	Repeats                  int          `json:"repeats"`
	Model                    string       `json:"model"`
	RunnerName               string       `json:"runner_name"`
	Timestamp                string       `json:"timestamp"`
	AvgBaselineTokens        float64      `json:"avg_baseline_tokens"`
	AvgVaultTokens           float64      `json:"avg_vault_tokens"`
	AvgBaselineDurationMs    float64      `json:"avg_baseline_duration_ms"`
	AvgVaultDurationMs       float64      `json:"avg_vault_duration_ms"`
	AvgBaselineKeywordScore  float64      `json:"avg_baseline_keyword_score"`
	AvgVaultKeywordScore     float64      `json:"avg_vault_keyword_score"`
	AvgBaselineJudgeScore    float64      `json:"avg_baseline_judge_score,omitempty"`
	AvgVaultJudgeScore       float64      `json:"avg_vault_judge_score,omitempty"`
	TokenChangePct           float64      `json:"token_change_pct"`
	GenerationTimeRatio      float64      `json:"generation_time_ratio"`
	KeywordCoverageChangePct float64      `json:"keyword_coverage_change_pct"`
	Results                  []TaskResult `json:"results"`
}

// TokenUsage tracks token consumption from LLM backends
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
