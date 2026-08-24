package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"llm-context-vault/pkg/sanitizer"
	"llm-context-vault/pkg/vault"
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

OPTIONS:
  --vault-dir <path>     Explicit path to the vault repository (defaults to current dir or LLM_VAULT_DIR)
  --redact-words <w1,w2> Comma-separated list of custom words/company names to redact
  --tool <name>          Explicit tool name for 'import' (agy, codex, opencode, aider)
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
	case "scan", "scan-all":
		scanCmd := flag.NewFlagSet("scan", flag.ExitOnError)
		customWordsFlag := scanCmd.String("redact-words", "", "Comma-separated words to redact")
		_ = scanCmd.Parse(args[1:])

		cfg := sanitizer.DefaultConfig()
		if *customWordsFlag != "" {
			for _, w := range strings.Split(*customWordsFlag, ",") {
				if trimmed := strings.TrimSpace(w); trimmed != "" {
					cfg.CustomKeywords = append(cfg.CustomKeywords, trimmed)
				}
			}
		}

		s := sanitizer.New(cfg)
		v := vault.New(workDir, s)

		fmt.Println("🔍 Unified Scanner: Auto-detecting and harvesting local AI assistant sessions...")
		fmt.Println()

		reports, err := v.ScanAll()
		if err != nil {
			fmt.Printf("❌ Unified scan error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("--------------------------------------------------------------------------------------------------------")
		fmt.Printf("%-20s %-40s %-10s %-16s %s\n", "Assistant", "Discovered Location", "Status", "Imported", "Skipped (Trivial/Greeting)")
		fmt.Println("--------------------------------------------------------------------------------------------------------")

		totalImported := 0
		totalSkipped := 0
		var allWarnings []string

		for _, rep := range reports {
			status := "SUCCESS"
			if len(rep.Warnings) > 0 && rep.ImportedCount == 0 {
				status = "FAILED"
			} else if len(rep.Warnings) > 0 {
				status = "WARN"
			}

			locPreview := rep.Location
			if len(locPreview) > 38 {
				locPreview = "..." + locPreview[len(locPreview)-35:]
			}

			fmt.Printf("%-20s %-40s %-10s %-16s %d session(s)\n",
				rep.ToolName,
				locPreview,
				status,
				fmt.Sprintf("%d session(s)", rep.ImportedCount),
				rep.SkippedCount,
			)
			totalImported += rep.ImportedCount
			totalSkipped += rep.SkippedCount
			for _, w := range rep.Warnings {
				allWarnings = append(allWarnings, fmt.Sprintf("[%s] %s", rep.ToolName, w))
			}
		}

		fmt.Println("--------------------------------------------------------------------------------------------------------")
		fmt.Printf("🎉 Total conversations imported & sanitized: %d (Skipped %d trivial greetings)\n\n", totalImported, totalSkipped)

		// Post-scan privacy audit
		fmt.Println("🛡️ Running instant security & privacy audit...")
		violations, err := v.AuditAll()
		if err != nil {
			fmt.Printf("⚠️ Audit error: %v\n", err)
		} else if len(violations) == 0 {
			fmt.Println("✅ Privacy Audit PASSED: 0 secrets, 0 private keys, 0 user paths detected.")
		} else {
			fmt.Printf("⚠️ Found %d file(s) with potential notices:\n", len(violations))
			for f, ws := range violations {
				fmt.Printf("   - %s: %s\n", filepath.Base(f), strings.Join(ws, ", "))
			}
		}

	case "scan-agy":
		brainDir := ""
		if len(args) >= 2 {
			brainDir = args[1]
		} else {
			homeDir, _ := os.UserHomeDir()
			defaultBrain := filepath.Join(homeDir, ".gemini", "antigravity-cli", "brain")
			if _, err := os.Stat(defaultBrain); err == nil {
				brainDir = defaultBrain
			}
		}

		if brainDir == "" {
			fmt.Println("❌ Could not automatically find Antigravity brain directory.")
			os.Exit(1)
		}

		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))
		imported, skipped, warnings, err := v.ScanAGYBrain(brainDir)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Imported %d Antigravity conversations (Skipped %d greetings/trivial)!\n", imported, skipped)
		if len(warnings) > 0 {
			fmt.Printf("⚠️ Notices (%d)\n", len(warnings))
		}

	case "scan-codex":
		codexDir := ""
		if len(args) >= 2 {
			codexDir = args[1]
		} else {
			homeDir, _ := os.UserHomeDir()
			defaultDir := filepath.Join(homeDir, ".codex", "sessions")
			if _, err := os.Stat(defaultDir); err == nil {
				codexDir = defaultDir
			}
		}

		if codexDir == "" {
			fmt.Println("❌ Could not automatically find Codex sessions directory.")
			os.Exit(1)
		}

		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))
		imported, skipped, warnings, err := v.ScanCodexSessions(codexDir)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Imported %d Codex conversations (Skipped %d greetings/trivial)!\n", imported, skipped)
		if len(warnings) > 0 {
			fmt.Printf("⚠️ Notices (%d)\n", len(warnings))
		}

	case "scan-opencode":
		dbPath := ""
		if len(args) >= 2 {
			dbPath = args[1]
		} else {
			homeDir, _ := os.UserHomeDir()
			defaultDB := filepath.Join(homeDir, ".local", "share", "opencode", "opencode.db")
			if _, err := os.Stat(defaultDB); err == nil {
				dbPath = defaultDB
			}
		}

		if dbPath == "" {
			fmt.Println("❌ Could not automatically find OpenCode opencode.db file.")
			os.Exit(1)
		}

		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))
		imported, skipped, warnings, err := v.ScanOpenCodeDB(dbPath)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Imported %d OpenCode conversations (Skipped %d greetings/trivial)!\n", imported, skipped)
		if len(warnings) > 0 {
			fmt.Printf("⚠️ Notices (%d)\n", len(warnings))
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

	case "search":
		if len(args) < 2 {
			fmt.Println("Error: Query string required.")
			fmt.Println("Example: vault search \"JWT validation\"")
			os.Exit(1)
		}
		query := strings.Join(args[1:], " ")
		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))

		results, err := v.Search(query)
		if err != nil {
			fmt.Printf("❌ Search failed: %v\n", err)
			os.Exit(1)
		}

		if len(results) == 0 {
			fmt.Println("No matching conversations found.")
			return
		}

		fmt.Printf("🔍 Found %d matching conversation(s) for '%s':\n\n", len(results), query)
		for i, res := range results {
			fmt.Printf("[%d] Tool: %-12s | File: %s.md\n", i+1, res.SourceTool, res.ConversationID)
			fmt.Printf("    Path: %s\n", res.Path)
			fmt.Printf("    Preview:\n%s\n\n", indent(res.Snippet, "      "))
		}

	case "context":
		if len(args) < 2 {
			fmt.Println("Error: Query string required.")
			fmt.Println("Example: vault context \"how to fix CORS in Express\"")
			os.Exit(1)
		}
		query := strings.Join(args[1:], " ")
		v := vault.New(workDir, sanitizer.New(sanitizer.DefaultConfig()))

		snippet, err := v.GenerateContextSnippet(query, 3)
		if err != nil {
			fmt.Printf("❌ Failed to generate context: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(snippet)

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
