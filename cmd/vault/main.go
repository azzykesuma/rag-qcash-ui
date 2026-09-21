package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/sanitizer"
	"llm-context-vault/pkg/vault"
	"llm-context-vault/pkg/webui"
)

func printUsage() {
	fmt.Print(`
llm-context-vault: Open-source repository for sanitized local LLM conversations

USAGE:
  vault <command> [options]

COMMANDS:
  scan                   🌟 Unified Scan: Auto-detects & extracts all local LLMs (AGY, OpenCode, Codex, Aider)
  scan-agy [dir]         Scan only Antigravity sessions
  scan-codex [dir]       Scan only OpenAI Codex sessions
  scan-opencode [db]     Scan only OpenCode SQLite database
  import <path>          Extract, sanitize, and add a single file/folder
  search <query>         Search indexed conversations for solutions or code
  context <query>        Generate context snippet to inject into local LLM prompt
  audit                  Audit all stored conversations for secrets / leaked paths
  stats                  Display statistics about stored conversations
  ui                     Launch the local web dashboard (--port 8080, --no-browser)
  pull [git args...]     Pull the latest changes from GitHub into the vault repository
  publish [git args...]  Audit, commit, and push sanitized conversations to GitHub

OPTIONS:
  --vault-dir <path>     Explicit path to the vault repository (defaults to current dir or LLM_VAULT_DIR)
  --redact-words <w1,w2> Comma-separated list of custom words/company names to redact
  --tool <name>          Explicit tool name for 'import' (agy, codex, opencode, aider)

SCAN OPTIONS:
  --full                 Reprocess every source instead of using the incremental cache
  --workers <n>          Parallel session workers (0 = automatic, up to 4)
  --limit <n>            Changed sessions per scan (default 10, 0 = unlimited)
  --limit-per-tool <n>   Changed sessions per assistant instead of one global limit
  --daily                Run at most one completed scan per local calendar day

SEARCH / CONTEXT OPTIONS:
  --project <alias>      Filter by exact project alias
  --tool <name>          Filter by source assistant
  --after/--before <day> Inclusive date filters in YYYY-MM-DD format
  --limit <n>            Maximum distinct conversations (search: 10, context: 3)
  --max-chars <n>        Context output budget in Unicode characters (context only)
`)
}

func resolveVaultDir(explicitDir string) string {
	if explicitDir != "" {
		return explicitDir
	}
	// 1. Environment variable
	if envDir := os.Getenv("LLM_VAULT_DIR"); envDir != "" {
		if _, err := os.Stat(envDir); err == nil {
			return envDir
		}
	}

	// 2. Current working directory if it contains conversations/
	if _, err := os.Stat("conversations"); err == nil {
		cwd, _ := os.Getwd()
		return cwd
	}

	// 3. Executable's own directory
	if exePath, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exePath)
		if _, err := os.Stat(filepath.Join(exeDir, "conversations")); err == nil {
			return exeDir
		}
	}

	// 4. Fallback to CWD
	cwd, _ := os.Getwd()
	return cwd
}

func main() {
	args, vaultDir := extractGlobalOptions(os.Args[1:])
	if len(args) == 0 {
		printUsage()
		return
	}

	workDir := resolveVaultDir(vaultDir)

	command := strings.ToLower(args[0])

	switch command {
	case "ui":
		uiCmd := flag.NewFlagSet("ui", flag.ExitOnError)
		port := uiCmd.Int("port", 8080, "Local dashboard port (0 selects an available port)")
		noBrowser := uiCmd.Bool("no-browser", false, "Do not automatically open the browser")
		_ = uiCmd.Parse(args[1:])
		if uiCmd.NArg() != 0 {
			fmt.Println("Error: ui accepts only --port, --no-browser, and --vault-dir options.")
			os.Exit(1)
		}
		if err := webui.Run(workDir, *port, !*noBrowser); err != nil {
			fmt.Printf("❌ UI failed: %v\n", err)
			os.Exit(1)
		}

	case "scan", "scan-all", "scan-agy", "scan-codex", "scan-opencode":
		if err := runScan(workDir, command, args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

	case "import":
		importCmd := flag.NewFlagSet("import", flag.ExitOnError)
		toolFlag := importCmd.String("tool", "", "Explicit tool name: agy, codex, opencode, aider, generic")
		customWordsFlag := importCmd.String("redact-words", "", "Comma-separated words to redact")

		if len(args) < 2 {
			fmt.Println("Error: Target path required for import.")
			fmt.Println("Example: vault import C:/path/to/transcript.jsonl --tool agy")
			os.Exit(1)
		}

		targetPath := args[1]
		_ = importCmd.Parse(args[2:])

		var customKeywords []string
		if *customWordsFlag != "" {
			for _, w := range strings.Split(*customWordsFlag, ",") {
				if trimmed := strings.TrimSpace(w); trimmed != "" {
					customKeywords = append(customKeywords, trimmed)
				}
			}
		}

		cfg := sanitizer.DefaultConfig()
		cfg.CustomKeywords = customKeywords
		s := sanitizer.New(cfg)
		v := vault.New(workDir, s)

		fmt.Printf("📦 Importing and sanitizing from %s...\n", targetPath)
		conversations, warnings, skipped, err := v.ProcessAndStoreAll(targetPath, *toolFlag)
		if err != nil {
			fmt.Printf("❌ Failed to process conversation: %v\n", err)
			os.Exit(1)
		}
		if len(conversations) == 0 {
			fmt.Println("ℹ️ Skipped: Conversation was classified as a trivial greeting or empty session.")
			return
		}

		fmt.Printf("✅ Successfully imported %d conversation(s) (%d skipped).\n", len(conversations), skipped)
		for _, conv := range conversations {
			fmt.Printf("   - %s (%s), %d messages from %s\n", conv.Title, conv.ID, len(conv.Messages), conv.SourceTool)
		}

		if len(warnings) > 0 {
			fmt.Printf("⚠️ Warning: Detected sensitive patterns during audit:\n")
			for _, w := range warnings {
				fmt.Printf("   - %s\n", w)
			}
		}

	case "search", "context":
		if err := runQuery(workDir, command, args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "%s failed: %v\n", command, err)
			os.Exit(1)
		}

	case "audit":
		fmt.Println("🛡️ Auditing stored conversations for potential secret leaks...")
		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))

		violations, err := v.AuditAll()
		if err != nil {
			fmt.Printf("❌ Error running audit: %v\n", err)
			os.Exit(1)
		}

		if len(violations) == 0 {
			fmt.Println("✅ Safe to publish! No secrets or unauthorized paths detected.")
		} else {
			fmt.Printf("⚠️ Found %d file(s) with potential sensitive content:\n", len(violations))
			for file, warnings := range violations {
				fmt.Printf("  - %s:\n", file)
				for _, w := range warnings {
					fmt.Printf("      * %s\n", w)
				}
			}
		}

	case "stats":
		showStats(workDir)

	case "pull":
		if err := pullRepo(workDir, args[1:]); err != nil {
			fmt.Printf("❌ Pull failed: %v\n", err)
			os.Exit(1)
		}

	case "publish":
		if err := publishRepo(workDir, args[1:]); err != nil {
			fmt.Printf("❌ Publish failed: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func extractGlobalOptions(args []string) ([]string, string) {
	result := make([]string, 0, len(args))
	var vaultDir string
	for i := 0; i < len(args); i++ {
		if args[i] == "--vault-dir" && i+1 < len(args) {
			vaultDir = args[i+1]
			i++
			continue
		}
		if strings.HasPrefix(args[i], "--vault-dir=") {
			vaultDir = strings.TrimPrefix(args[i], "--vault-dir=")
			continue
		}
		result = append(result, args[i])
	}
	return result, vaultDir
}

func publishRepo(baseDir string, extraArgs []string) error {
	if _, err := os.Stat(filepath.Join(baseDir, ".git")); err != nil {
		return fmt.Errorf("%s is not a git repository (no .git directory)", baseDir)
	}
	v := vault.New(baseDir, sanitizer.New(sanitizer.DefaultConfig()))
	release, err := v.AcquireWriterLock()
	if err != nil {
		return err
	}
	defer release()

	// 1. Run Pre-Publish Security Audit
	fmt.Println("🛡️ Pre-publish Security & Privacy Audit...")
	violations, err := v.AuditAll()
	if err != nil {
		return fmt.Errorf("audit failed: %w", err)
	}
	strict := false
	for _, arg := range extraArgs {
		if arg == "--strict" {
			strict = true
			break
		}
	}
	if len(violations) > 0 {
		if strict {
			fmt.Printf("❌ Publish ABORTED (--strict enabled)! Found %d file(s) with potential secrets/paths:\n", len(violations))
			for file, warnings := range violations {
				fmt.Printf("  - %s:\n", file)
				for _, w := range warnings {
					fmt.Printf("      * %s\n", w)
				}
			}
			return fmt.Errorf("cannot publish while security warnings exist")
		}
		fmt.Printf("⚠️ Warning: Found %d file(s) with potential sensitive content during audit:\n", len(violations))
		for file, warnings := range violations {
			fmt.Printf("  - %s:\n", file)
			for _, w := range warnings {
				fmt.Printf("      * %s\n", w)
			}
		}
	} else {
		fmt.Println("✅ Privacy Audit PASSED: 0 secrets or sensitive paths detected.")
	}
	fmt.Println()

	// 2. Stage conversations directory
	fmt.Println("📦 Staging sanitized conversations...")
	for _, addArgs := range [][]string{
		{"add", "-u", "--", "conversations/"},
		{"add", "--", "conversations/dataset.jsonl", ":(glob)conversations/markdown/*.md", ":(glob)conversations/sharegpt/*.json"},
	} {
		addCmd := exec.Command("git", addArgs...)
		addCmd.Dir = baseDir
		addCmd.Stdout = os.Stdout
		addCmd.Stderr = os.Stderr
		if err := addCmd.Run(); err != nil {
			return fmt.Errorf("git add failed: %w", err)
		}
	}

	// 3. Check status to see if anything needs committing
	statusCmd := exec.Command("git", "status", "--porcelain", "conversations/")
	statusCmd.Dir = baseDir
	statusOutput, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("git status check failed: %w", err)
	}

	customMsg := ""
	for i := 0; i < len(extraArgs); i++ {
		if (extraArgs[i] == "-m" || extraArgs[i] == "--message") && i+1 < len(extraArgs) {
			customMsg = extraArgs[i+1]
			break
		}
	}

	if len(strings.TrimSpace(string(statusOutput))) > 0 {
		commitMsg := customMsg
		if commitMsg == "" {
			mdDir := filepath.Join(baseDir, "conversations", "markdown")
			mdFiles, _ := os.ReadDir(mdDir)
			timestamp := time.Now().Format("2006-01-02 15:04")
			commitMsg = fmt.Sprintf("chore(vault): sync %d sanitized conversation session(s) [%s]", len(mdFiles), timestamp)
		}

		fmt.Printf("💾 Committing changes: %s\n", commitMsg)
		commitCmd := exec.Command("git", "commit", "-m", commitMsg)
		commitCmd.Dir = baseDir
		commitCmd.Stdout = os.Stdout
		commitCmd.Stderr = os.Stderr
		if err := commitCmd.Run(); err != nil {
			return fmt.Errorf("git commit failed: %w", err)
		}
	} else {
		fmt.Println("ℹ️ No new changes to commit in conversations/ directory.")
	}

	// 4. Push to remote
	fmt.Println()
	fmt.Printf("🚀 Pushing to remote repository...\n")
	pushArgs := []string{"push"}
	if len(extraArgs) > 0 && customMsg == "" {
		pushArgs = append(pushArgs, extraArgs...)
	}
	pushCmd := exec.Command("git", pushArgs...)
	pushCmd.Dir = baseDir
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("git push failed: %w", err)
	}

	fmt.Println()
	fmt.Println("🎉 Successfully published sanitized conversations to remote!")
	return nil
}

func pullRepo(baseDir string, extraArgs []string) error {
	if _, err := os.Stat(filepath.Join(baseDir, ".git")); err != nil {
		return fmt.Errorf("%s is not a git repository (no .git directory)", baseDir)
	}
	v := vault.New(baseDir, sanitizer.New(sanitizer.DefaultConfig()))
	release, err := v.AcquireWriterLock()
	if err != nil {
		return err
	}
	defer release()

	fmt.Printf("🔁 Pulling latest changes from GitHub into %s...\n", baseDir)
	fmt.Println()

	pullArgs := append([]string{"pull", "--ff-only"}, extraArgs...)
	cmd := exec.Command("git", pullArgs...)
	cmd.Dir = baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull failed: %w", err)
	}

	fmt.Println()
	fmt.Println("✅ Vault is up to date.")
	return nil
}

func showStats(baseDir string) {
	mdDir := filepath.Join(baseDir, "conversations", "markdown")
	shareDir := filepath.Join(baseDir, "conversations", "sharegpt")

	mdFiles, _ := os.ReadDir(mdDir)
	shareFiles, _ := os.ReadDir(shareDir)

	fmt.Println("📊 Repository Statistics:")
	fmt.Printf("   - Markdown Sessions: %d\n", len(mdFiles))
	fmt.Printf("   - ShareGPT Sessions: %d\n", len(shareFiles))
	datasetPath := filepath.Join(baseDir, "conversations", "dataset.jsonl")
	if info, err := os.Stat(datasetPath); err == nil {
		fmt.Printf("   - Master Dataset (dataset.jsonl): %.2f KB\n", float64(info.Size())/1024.0)
	} else {
		fmt.Println("   - Master Dataset: None yet")
	}
}

func indent(s, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = prefix + l
	}
	return strings.Join(lines, "\n")
}
