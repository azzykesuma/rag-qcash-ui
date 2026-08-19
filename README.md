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
1. Run `vault scan-agy` or `vault import <path>`.
2. Run `vault audit` to ensure all secrets are scrubbed.
3. Commit the updated `conversations/` directory and open a Pull Request!
