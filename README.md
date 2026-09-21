# LLM Context Vault 🧠

An open-source repository and CLI tool to extract, sanitize, and share problem-solving coding conversations across local LLM assistants (Antigravity/AGY, Aider, OpenCode, Codex, Cursor, etc.).

Anyone who clones this repo can search past real-world developer sessions, inject relevant debugging context directly into their local LLM prompts, or fine-tune local models.

---

## 🚀 Why Go for this Project?

| Criteria | **Go (Chosen)** | Node.js |
| :--- | :--- | :--- |
| **Distribution** | **Zero dependencies**. Produces a standalone binary (`vault` / `vault.exe`) that users run immediately without installing runtimes. | Requires Node.js runtime, `npm install`, and massive `node_modules`. |
| **Speed & I/O** | Extremely fast streaming line-by-line JSONL parsing and multi-pass regex redactions across gigabytes of chat logs. | Slower streaming, high memory overhead on large file sets. |
| **Concurrency** | Goroutines scan multi-agent sessions concurrently with zero overhead. | Event-loop based, requires worker threads for heavy regex processing. |

---

## 📂 Which Files Get Stored in Git?

Never commit raw database files or unredacted system logs. The vault normalizes conversations into two formats:

1. **`conversations/markdown/*.md`**: Human-readable GitHub markdown with code blocks, timestamps, and tool summaries.
2. **`conversations/sharegpt/*.json` & `conversations/dataset.jsonl`**: Standardized dataset format compatible with HuggingFace, OpenAI fine-tuning, and Ollama RAG pipelines.

```text
llm-context-vault/
├── cmd/vault/main.go               # Standalone CLI entrypoint
├── pkg/
│   ├── models/types.go             # Unified conversation schema
│   ├── sanitizer/sanitizer.go      # Secret, PII & path redaction engine
│   ├── extractor/                  # Tool-specific extractors (AGY, Aider, JSON)
│   ├── exporter/                   # Markdown & ShareGPT JSONL exporters
│   └── vault/vault.go              # Storage, search, and context generator
├── conversations/                  # Sanitized datasets ready for Git
│   ├── markdown/
│   ├── sharegpt/
│   └── dataset.jsonl
├── go.mod
└── README.md
```

---

## 🛡️ Sanitization & Redaction Engine

Before any conversation is added to the repository, the engine executes a multi-layer scrubbing pass:

1. **API Keys & Tokens**:
   - OpenAI (`sk-...`, `sk-proj-...`)
   - Anthropic (`sk-ant-...`)
   - Google AI Studio (`AIza...`)
   - GitHub PATs (`ghp_...`, `github_pat_...`)
   - AWS Keys (`AKIA...` and secret keys)
   - JWT / Bearer tokens & RSA/EC private key blocks
   - Database connection strings (`postgres://user:pass@host` -> `postgres://user:[PASSWORD_REDACTED]@host`)
2. **Paths & Hostnames**:
   - Strips Windows paths (`C:\Users\username\...` -> `~/...`)
   - Strips Unix paths (`/home/username/...` -> `~/...`)
3. **PII**:
   - Emails -> `[REDACTED_EMAIL]`
   - Non-loopback IP addresses -> `[REDACTED_IP]`
4. **Custom Keywords**:
   - Support `--redact-words "CompanyInc,InternalSecretProject"`

---

## 📥 Installation & Global Setup

### 1. Download Pre-Compiled Binary
Grab the executable for your OS from [GitHub Releases](https://github.com/azzykesuma/llm-context-vault/releases/latest):
- **Windows**: [`vault-windows-amd64.exe`](https://github.com/azzykesuma/llm-context-vault/releases/latest/download/vault-windows-amd64.exe) (rename to `vault.exe`)
- **Linux**: [`vault-linux-amd64`](https://github.com/azzykesuma/llm-context-vault/releases/latest/download/vault-linux-amd64)
- **macOS (Apple Silicon)**: [`vault-darwin-arm64`](https://github.com/azzykesuma/llm-context-vault/releases/latest/download/vault-darwin-arm64)
- **macOS (Intel)**: [`vault-darwin-amd64`](https://github.com/azzykesuma/llm-context-vault/releases/latest/download/vault-darwin-amd64)

### 2. Make `vault` Available Globally (Terminal & VS Code)

#### 🪟 Windows (PowerShell & VS Code)
To run `vault` from any terminal or inside VS Code:

```powershell
# 1. Add vault's directory to your User PATH permanently
[System.Environment]::SetEnvironmentVariable('Path', [System.Environment]::GetEnvironmentVariable('Path', 'User') + ';D:\code\llm-context-vault', 'User')

# 2. Add alias to PowerShell Profile (ensures instant availability in VS Code)
if (!(Test-Path $PROFILE)) { New-Item -ItemType File -Path $PROFILE -Force }
Add-Content -Path $PROFILE -Value 'Set-Alias -Name vault -Value "D:\code\llm-context-vault\vault.exe"'
```

#### 🐧 Linux / 🍎 macOS
```bash
# Move to system bin
chmod +x vault-linux-amd64
sudo mv vault-linux-amd64 /usr/local/bin/vault
```

### 3. (Optional) Set Central Vault Directory Environment Variable
When running `vault scan` or `vault search` from outside the repository (e.g. inside another workspace like `my-web-app/`), `vault` automatically discovers your central repository. You can also explicitly define it:
```powershell
# Windows
[System.Environment]::SetEnvironmentVariable('LLM_VAULT_DIR', 'D:\code\llm-context-vault', 'User')

# Linux / macOS (in ~/.bashrc or ~/.zshrc)
export LLM_VAULT_DIR="$HOME/code/llm-context-vault"
```

---

## 🛠️ Usage Guide

### 1. 🌟 Single Unified Scan (All Local Assistants)
Auto-detects and harvests sessions from **Antigravity (AGY)**, **OpenCode**, **OpenAI Codex**, and **Aider** all in one command:
```bash
vault scan
```

Scans are incremental and bounded by default. Unchanged sessions are skipped and at most 10 new or changed sessions are processed across all assistants per invocation. Deferred sessions remain pending in the scan state, so run the command again to process the next batch. Changed sessions are processed concurrently, and `dataset.jsonl` is merged once at the end instead of being rewritten for every conversation.

Choose a different batching policy when needed:

```bash
vault scan --limit 10             # 10 changed sessions total (the default)
vault scan --limit-per-tool 10    # Up to 10 for each assistant
vault scan --daily                # One completed 10-session batch per local day
vault scan --limit 0              # Process the entire pending backlog
vault scan --full                 # Reprocess every source without a batch limit
vault scan --workers 8
```

`--limit` and `--limit-per-tool` cannot be combined. A bounded scan audits every conversation before storing it but defers the repository-wide privacy audit so it does not re-read the entire vault after each small batch. Run `vault audit` before committing or publishing; `vault publish` also performs the full audit.

The local scan state and search index live under `.vault/` and are excluded from Git. They can be deleted safely; the next scan or search rebuilds them from source sessions and sanitized exports. A per-vault writer lock prevents concurrent scans from corrupting output.

Output:
```text
Assistant          New Changed  Unchanged Deferred Trivial Failed       Time
antigravity          2       1        115        7       0      0      1.2s
opencode             1       0        121        0       0      0     540ms
codex                0       0         14        0       0      0      12ms
Import phase: 1.8s
Search index: 24ms

Full privacy audit deferred for this bounded scan; run `vault audit` before publishing.
```

### 2. Custom Redaction Keywords
Redact specific client names, internal projects, or company code names:
```bash
vault scan --redact-words "SecretClient,ProjectTitan,InternalService"
```

### Vault Location And Repeatable Imports
Use `--vault-dir` from any working directory to select the repository that receives exports:
```bash
vault scan --vault-dir /path/to/llm-context-vault
```

Exports use a stable ID derived from sanitized content. Re-importing unchanged content updates the same Markdown, ShareGPT, and JSONL records instead of creating duplicates. New exports receive an offline, conversation-aware filename based on recurring topics across the session, rather than the opening words alone. Existing filenames are preserved. Generic JSON and JSONL imports support one or many normalized or ShareGPT conversations:
```bash
vault import conversations.jsonl --tool generic
```

### 3. Tool-Specific Scans
```bash
vault scan-agy       # Antigravity sessions only
vault scan-opencode  # OpenCode SQLite DB only
vault scan-codex     # OpenAI Codex sessions only
```

### 4. Search and Generate Context for Local LLM
```bash
# Search across all past sessions (AGY, OpenCode, Codex)
vault search "JWT validation"

# Output context snippet to pipe into an LLM prompt
vault context "how to fix splash screen crash"

# Rank passages within one project and date range
vault search "session expiration" --project qcash-ui --after 2026-01-01 --limit 10

# Generate bounded, source-cited context for an assistant
vault context "why does the session expire early?" --project qcash-ui --max-chars 12000
```

Search uses a local SQLite FTS index and ranks conversation passages instead of requiring an exact full-query substring. Results are deduplicated by session and content. Context output includes the export path, source line range, and turn range, and is explicitly marked as historical reference rather than current instructions.

Project aliases are captured from source metadata when available. For older exports, create a local `.vault/projects.json` mapping from the Markdown filename to your preferred project alias:

```json
{
  "opencode_session-expiration-a1b2c3d4e5f6.md": "qcash-ui"
}
```

This mapping remains local and is not committed.

### 5. Local Web Dashboard
Launch the embedded dashboard to browse conversations, filter by assistant/project/date, run ranked search, generate bounded prompt context, and run a privacy audit:
```bash
vault ui
vault ui --port 3000 --no-browser
vault ui --vault-dir /path/to/llm-context-vault
```
By default, the dashboard opens your browser at `http://127.0.0.1:8080`. It runs locally with no external web assets or runtime dependencies. Conversations are displayed as exported Markdown text. Press `Ctrl+C` in the terminal to stop it.

### 6. Audit Before Git Push
Verify that 0 secrets or machine paths exist in your repository:
```bash
vault audit
```

### 7. Pull Latest Changes
Fetch and fast-forward the vault repository to the latest changes from GitHub:
```bash
vault pull
```
Extra arguments are forwarded to `git pull` (e.g. `vault pull --rebase`).

---

## 🤝 Contributing
Contributions are welcome! Please check out our [Contributing Guidelines](CONTRIBUTING.md) to see how you can safely scan and submit your own local assistant problem-solving sessions via Pull Request.

Dataset publication, retention, removal, schema, and attribution requirements are documented in [DATA_GOVERNANCE.md](DATA_GOVERNANCE.md).

---

## 📄 License
MIT License. Free for open-source research, developer tooling, and dataset enrichment.

