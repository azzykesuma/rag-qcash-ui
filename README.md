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
Grab the executable for your OS from [GitHub Releases](https://github.com/azzykesuma/rag-qcash-ui/releases/latest):
- **Windows**: [`vault-windows-amd64.exe`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-windows-amd64.exe) (rename to `vault.exe`)
- **Linux**: [`vault-linux-amd64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-linux-amd64)
- **macOS (Apple Silicon)**: [`vault-darwin-arm64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-darwin-arm64)
- **macOS (Intel)**: [`vault-darwin-amd64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-darwin-amd64)

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

Output:
```text
🔍 Unified Scanner: Auto-detecting and harvesting local AI assistant sessions...

-----------------------------------------------------------------------------------------
Assistant            Discovered Location                              Status     Imported
-----------------------------------------------------------------------------------------
Antigravity (AGY)    ...sers\dev\.gemini\antigravity-cli\brain        SUCCESS    18 session(s)
OpenCode             ...\dev\.local\share\opencode\opencode.db        SUCCESS    121 session(s)
Codex                C:\Users\dev\.codex\sessions                     SUCCESS    4 session(s)
-----------------------------------------------------------------------------------------
🎉 Total conversations imported & sanitized: 143

🛡️ Running instant security & privacy audit...
✅ Privacy Audit PASSED: 0 secrets, 0 private keys, 0 user paths detected.
```

### 2. Custom Redaction Keywords
Redact specific client names, internal projects, or company code names:
```bash
vault scan --redact-words "SecretClient,ProjectTitan,InternalService"
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
```

### 5. Audit Before Git Push
Verify that 0 secrets or machine paths exist in your repository:
```bash
vault audit
```

---

## 🤝 Contributing
Contributions are welcome! Please check out our [Contributing Guidelines](CONTRIBUTING.md) to see how you can safely scan and submit your own local assistant problem-solving sessions via Pull Request.

---

## 📄 License
MIT License. Free for open-source research, developer tooling, and dataset enrichment.

