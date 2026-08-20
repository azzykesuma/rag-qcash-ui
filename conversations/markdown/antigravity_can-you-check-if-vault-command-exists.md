# can you check if vault command exists?

- **ID**: `8b68858a-0b4e-4999-b133-e51e32e70fdf`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-20 08:05:00`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

can you check if vault command exists?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **run_command**: 
- **view_file**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 6: Assistant

Yes, the **`vault`** command exists on your system:

### 1. PowerShell Command / Alias
* **Type:** PowerShell Alias
* **Target Binary:** `D:\code\llm-context-vault\vault.exe`
* **Application:** `llm-context-vault` (tool to manage, scan, search, and sanitize local LLM conversation context from Antigravity, OpenCode, Codex, and Aider)

### 2. Available Commands
```text
USAGE:
  vault <command> [options]

COMMANDS:
  scan                   Auto-detects & extracts all local LLMs (AGY, OpenCode, Codex, Aider)
  scan-agy [dir]         Scan only Antigravity sessions
  scan-codex [dir]       Scan only OpenAI Codex sessions
  scan-opencode [db]     Scan only OpenCode SQLite database
  import <path>          Extract, sanitize, and add a single file/folder
  search <query>         Search indexed conversations for solutions or code
  context <query>        Generate context snippet to inject into local LLM prompt
  audit                  Audit all stored conversations for secrets / leaked paths
  stats                  Display statistics about stored conversations
```

### 3. Integrated Skill
The [`vault`](file:///C:~/.gemini/config/skills/vault/SKILL.md) skill is also active in Antigravity to retrieve past debugging sessions, architecture decisions, and code solutions directly from `D:\code\llm-context-vault`.

---

