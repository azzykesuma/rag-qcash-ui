# Contributing to LLM Context Vault 🤝

Thank you for your interest in contributing to **LLM Context Vault**! 

The mission of this repository is to build an open, crowdsourced knowledge base of real-world problem-solving, debugging, and architecture discussions with local AI coding assistants (Antigravity, OpenCode, Codex, Aider, etc.) — **with 100% of sensitive information, tokens, and machine paths strictly redacted**.

---

## 🛠️ Supported AI Assistants

The built-in `vault` CLI automatically discovers and parses sessions from:
- **Antigravity (AGY)** (`~/.gemini/antigravity-cli/brain/`)
- **OpenCode** (`~/.local/share/opencode/opencode.db` or `~/.opencode/`)
- **OpenAI Codex** (`~/.codex/sessions/YYYY/MM/DD/`)
- **Aider** (`.aider.chat.history.md`)
- **ShareGPT / Generic JSON** (`vault import <path>`)

---

## 🚀 Quick Contribution Guide (5 Steps)

### Step 1: Fork & Clone
Fork this repository on GitHub, then clone your fork locally:
```bash
git clone https://github.com/<your-username>/rag-qcash-ui.git
cd rag-qcash-ui
```

---

### Step 2: Download the `vault` CLI
Download the standalone executable for your operating system from our [Latest Releases](https://github.com/azzykesuma/rag-qcash-ui/releases/latest):

- **Windows**: [`vault-windows-amd64.exe`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-windows-amd64.exe) (rename to `vault.exe`)
- **Linux**: [`vault-linux-amd64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-linux-amd64) (`chmod +x vault-linux-amd64 && mv vault-linux-amd64 vault`)
- **macOS (Apple Silicon)**: [`vault-darwin-arm64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-darwin-arm64)
- **macOS (Intel)**: [`vault-darwin-amd64`](https://github.com/azzykesuma/rag-qcash-ui/releases/latest/download/vault-darwin-amd64)

*(Alternatively, build directly from source if you have Go installed: `go build -o vault ./cmd/vault`)*

---

### Step 3: Run the Auto-Scanner & Sanitizer
Run the unified scan command inside the repository root:

```bash
# Standard scan across all local assistants
./vault scan

# Scan with custom proprietary keywords or project names scrubbed:
./vault scan --redact-words "MyCompany,ClientAlpha,InternalService"
```

#### What the scanner handles automatically:
1. **Low-Value Filtering**: Skips empty turns, aborted runs, and greeting-only chats (`"hi"`, `"test"`, `"halo"`).
2. **Multi-Layer Secret Redaction**: Replaces OpenAI keys (`sk-...`), Anthropic keys, GitHub PATs, Bitbucket tokens, SonarQube tokens, JWTs, AWS credentials, and database passwords with safe placeholders like `[REDACTED_SECRET]`.
3. **Machine Path Anonymization**: Converts `C:\Users\username\...` and `/home/username/...` to relative `~/` paths.
4. **Readable Slugs**: Generates human-readable Markdown files in `conversations/markdown/` based on your initial prompt.

---

### Step 4: Run the Local Privacy Audit
Before committing, verify that zero secrets or unauthorized paths exist:

```bash
./vault audit
```

**Expected Output:**
```text
✅ Safe to publish! No secrets or unauthorized paths detected.
```

---

### Step 5: Commit & Open a Pull Request

Create a new branch, stage your new conversation files, and push:

```bash
# 1. Create a feature branch
git checkout -b add-sessions-web-frameworks

# 2. Stage new conversation files
git add conversations/

# 3. Commit your changes
git commit -m "feat(dataset): add sanitized debugging sessions for Next.js and Go"

# 4. Push to your fork
git push origin add-sessions-web-frameworks
```

Finally, open a **Pull Request (PR)** on GitHub targeting the `main` branch.

---

## 🛡️ Automated CI Security Gate

Every Pull Request is automatically audited by our GitHub Actions workflow ([`audit.yml`](.github/workflows/audit.yml)):
- Runs `./vault audit` across all added/modified files.
- Blocks merging if any unredacted token, private key, or machine path is detected.

Thank you for helping build a better, shared context library for local LLMs!
