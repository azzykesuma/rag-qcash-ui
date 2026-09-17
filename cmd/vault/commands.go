package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/sanitizer"
	"llm-context-vault/pkg/vault"
)

// parseInterspersed accepts both `query --limit 5` and `--limit 5 query`.
func parseInterspersed(fs *flag.FlagSet, args []string) error {
	var options, positional []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") || arg == "-" {
			positional = append(positional, arg)
			continue
		}
		name, _, hasValue := strings.Cut(strings.TrimLeft(arg, "-"), "=")
		f := fs.Lookup(name)
		if f == nil {
			return fmt.Errorf("unknown option: %s", arg)
		}
		options = append(options, arg)
		boolean, ok := f.Value.(interface{ IsBoolFlag() bool })
		if !hasValue && !(ok && boolean.IsBoolFlag()) {
			if i+1 >= len(args) {
				return fmt.Errorf("missing value for %s", arg)
			}
			i++
			options = append(options, args[i])
		}
	}
	return fs.Parse(append(append(options, "--"), positional...))
}

func runScan(baseDir, command string, args []string) error {
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	words := flags.String("redact-words", "", "Comma-separated custom redaction words")
	full := flags.Bool("full", false, "Reprocess every source, bypassing scan cache")
	workers := flags.Int("workers", 0, "Parallel session workers (0 = automatic, up to 4)")
	if err := parseInterspersed(flags, args); err != nil {
		return err
	}
	if *workers < 0 || *workers > 64 {
		return fmt.Errorf("workers must be between 0 and 64")
	}
	unified := command == "scan" || command == "scan-all"
	if flags.NArg() > 1 || unified && flags.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments")
	}
	cfg := sanitizer.DefaultConfig()
	for _, word := range strings.Split(*words, ",") {
		if word = strings.TrimSpace(word); word != "" {
			cfg.CustomKeywords = append(cfg.CustomKeywords, word)
		}
	}
	v := vault.New(baseDir, sanitizer.New(cfg))
	v.ScanOptions = vault.ScanOptions{Full: *full, Workers: *workers, Progress: func(message string) { fmt.Fprintln(os.Stderr, message) }}
	started := time.Now()
	var reports []vault.ToolScanReport
	var err error
	if unified {
		reports, err = v.ScanAll()
	} else {
		home, homeErr := os.UserHomeDir()
		if homeErr != nil {
			return homeErr
		}
		source := vault.ScanSource{}
		switch command {
		case "scan-agy":
			source = vault.ScanSource{Tool: "antigravity", Path: filepath.Join(home, ".gemini", "antigravity-cli", "brain")}
		case "scan-codex":
			source = vault.ScanSource{Tool: "codex", Path: filepath.Join(home, ".codex", "sessions")}
		case "scan-opencode":
			source = vault.ScanSource{Tool: "opencode", Path: filepath.Join(home, ".local", "share", "opencode", "opencode.db")}
		}
		if flags.NArg() == 1 {
			source.Path = flags.Arg(0)
		}
		reports, err = v.ScanSources([]vault.ScanSource{source})
	}
	if err != nil {
		return err
	}
	fmt.Printf("\n%-14s %7s %7s %10s %7s %7s %10s\n", "Assistant", "New", "Changed", "Unchanged", "Trivial", "Failed", "Time")
	failures := 0
	for _, r := range reports {
		fmt.Printf("%-14s %7d %7d %10d %7d %7d %10s\n", r.ToolName, r.NewCount, r.ChangedCount, r.UnchangedCount, r.SkippedCount, r.FailedCount, r.Duration.Round(time.Millisecond))
		failures += r.FailedCount
		for _, warning := range r.Warnings {
			fmt.Fprintf(os.Stderr, "[%s] %s\n", r.ToolName, warning)
		}
	}
	fmt.Printf("Import phase: %s\n", time.Since(started).Round(time.Millisecond))
	if len(reports) == 0 {
		fmt.Println("No local assistant sources discovered.")
	}
	if _, err := os.Stat(filepath.Join(baseDir, "conversations", "markdown")); err == nil {
		indexStarted := time.Now()
		fmt.Fprintln(os.Stderr, "Updating local search index...")
		if err := v.RefreshSearchIndex(); err != nil {
			return err
		}
		fmt.Printf("Search index: %s\n", time.Since(indexStarted).Round(time.Millisecond))
	}
	if unified {
		if _, err := os.Stat(filepath.Join(baseDir, "conversations")); err == nil {
			auditStarted := time.Now()
			fmt.Fprintln(os.Stderr, "Running full privacy audit...")
			violations, err := v.AuditAll()
			if err != nil {
				return err
			}
			if len(violations) == 0 {
				fmt.Println("Privacy audit passed.")
			} else {
				for path, warnings := range violations {
					fmt.Printf("%s: %s\n", path, strings.Join(warnings, ", "))
				}
				return fmt.Errorf("privacy audit found potential sensitive content in %d file(s)", len(violations))
			}
			fmt.Printf("Audit phase: %s\n", time.Since(auditStarted).Round(time.Millisecond))
		}
	}
	fmt.Printf("Total: %s\n", time.Since(started).Round(time.Millisecond))
	if failures > 0 {
		return fmt.Errorf("%d source/session failures; unsuccessful imports will be retried", failures)
	}
	return nil
}

func runQuery(baseDir, command string, args []string) error {
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	var opts vault.SearchOptions
	flags.StringVar(&opts.Project, "project", "", "Exact project alias")
	flags.StringVar(&opts.Tool, "tool", "", "Assistant name")
	flags.StringVar(&opts.After, "after", "", "Inclusive earliest date, YYYY-MM-DD")
	flags.StringVar(&opts.Before, "before", "", "Inclusive latest date, YYYY-MM-DD")
	limit := 10
	if command == "context" {
		limit = 3
	}
	flags.IntVar(&opts.Limit, "limit", limit, "Maximum number of distinct conversations")
	maxChars := 12000
	if command == "context" {
		flags.IntVar(&maxChars, "max-chars", 12000, "Maximum context size in Unicode characters")
	}
	if err := parseInterspersed(flags, args); err != nil {
		return err
	}
	query := strings.TrimSpace(strings.Join(flags.Args(), " "))
	if query == "" {
		return fmt.Errorf("query required")
	}
	if opts.Limit < 1 {
		return fmt.Errorf("limit must be between 1 and 100")
	}
	v := vault.New(baseDir, sanitizer.New(sanitizer.DefaultConfig()))
	if command == "context" {
		text, err := v.ContextWithOptions(query, opts, maxChars)
		if err != nil {
			return err
		}
		fmt.Print(text)
		return nil
	}
	results, err := v.SearchWithOptions(query, opts)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		fmt.Println("No matching conversations found.")
		return nil
	}
	for i, r := range results {
		preview := r.Snippet
		if runes := []rune(preview); len(runes) > 1200 {
			preview = string(runes[:1200]) + "\n[…truncated; use `vault context` for a larger excerpt]"
		}
		fmt.Printf("[%d] %s | %s | project: %s\n    %s:%d-%d\n%s\n\n", i+1, r.Title, r.SourceTool, r.Project, r.Path, r.StartLine, r.EndLine, indent(preview, "    "))
	}
	return nil
}
