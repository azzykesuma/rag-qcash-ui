# Analyze codebase (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_fe18de0c4ffe4Z2im8TTmgKeH0`
- **Source Tool**: `opencode`
- **Date**: `2026-08-20 16:13:10`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository at D:\code\llm-context-vault with medium thoroughness. Do not modify files. Identify: project purpose, languages/frameworks/dependencies, top-level structure, core execution/data flows, tests/tooling, documentation, and notable risks or gaps. Inspect relevant manifests/readmes/source and return a concise but specific architecture summary with exact file paths. Also report git worktree state but do not alter it.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

## Architecture summary

**Purpose.** A Go CLI and Git-backed dataset for harvesting local coding-assistant conversations, sanitizing them, and exporting searchable Markdown plus ShareGPT/JSONL records for reuse in prompts or model/RAG training. Primary docs: `D:\code\llm-context-vault\README.md`, `D:\code\llm-context-vault\CONTRIBUTING.md`.

**Technology.**
- Go 1.25 module: `D:\code\llm-context-vault\go.mod`
- Only direct dependency is pure-Go SQLite driver `modernc.org/sqlite`, used to read OpenCode databases: `D:\code\llm-context-vault\pkg\extractor\opencode.go`
- GitHub Actions builds Windows/Linux/macOS static binaries and publishes tag releases: `D:\code\llm-context-vault\.github\workflows\release.yml`

**Top-level structure.**
- CLI: `D:\code\llm-context-vault\cmd\vault\main.go`
- Domain models: `D:\code\llm-context-vault\pkg\models\types.go`
- Extractors: `D:\code\llm-context-vault\pkg\extractor\`
- Sanitization: `D:\code\llm-context-vault\pkg\sanitizer\sanitizer.go`
- Export/storage/search orchestration: `D:\code\llm-context-vault\pkg\vault\vault.go`, `D:\code\llm-context-vault\pkg\exporter\`
- Published conversation corpus: `D:\code\llm-context-vault\conversations\markdown\`, `D:\code\llm-context-vault\conversations\sharegpt\`, and `D:\code\llm-context-vault\conversations\dataset.jsonl` (736 tracked conversation artifacts).
- Admin script: `D:\code\llm-context-vault\scripts\set-github-token.ps1`

## Core flow

1. `vault scan` discovers AGY, OpenCode, Codex, and Aider locations; individual scan/import commands are dispatched in `D:\code\llm-context-vault\cmd\vault\main.go`.
2. Tool-specific extractors normalize source data into `models.Conversation`:
   - Antigravity JSONL: `D:\code\llm-context-vault\pkg\extractor\agy.go`
   - Codex rollout JSONL: `D:\code\llm-context-vault\pkg\extractor\codex.go`
   - OpenCode SQLite `session`/`message`/`part` tables: `D:\code\llm-context-vault\pkg\extractor\opencode.go`
   - Aider history Markdown: `D:\code\llm-context-vault\pkg\extractor\aider.go`
   - Generic normalized/ShareGPT JSON: `D:\code\llm-context-vault\pkg\extractor\generic.go`
3. `Vault.StoreConversation` filters trivial chats, sanitizes message and tool-call text, performs a post-sanitize audit, then writes Markdown and ShareGPT files and appends one normalized record to JSONL: `D:\code\llm-context-vault\pkg\vault\vault.go`.
4. Search is case-insensitive literal scanning of stored Markdown; `context` formats up to three matching snippets rather than using embeddings or the SQLite index: `D:\code\llm-context-vault\pkg\vault\vault.go`.
5. Sanitization uses regex rules for common credentials, paths, a specific internal domain, emails, IPs, and optional custom terms: `D:\code\llm-context-vault\pkg\sanitizer\sanitizer.go`.

## Tests, tooling, and documentation

- No `*_test.go` files or test framework/configuration found.
- CI only builds and runs `vault audit` on pushes/PRs to `main`: `D:\code\llm-context-vault\.github\workflows\audit.yml`.
- No Makefile, Docker configuration, lint configuration, or dependency-update tooling found.
- README and contribution guide are substantial, though release/clone URLs still use the legacy repository name `rag-qcash-ui`.

## Notable risks/gaps

- **No automated tests**, including for extractors, redaction patterns, malformed input, and export idempotency.
- **Potential duplicate/data-loss behavior:** Markdown/ShareGPT filenames derive from truncated prompt slugs and can overwrite each other; JSONL always appends, so repeat scans can duplicate records. See `D:\code\llm-context-vault\pkg\vault\vault.go` and `D:\code\llm-context-vault\pkg\exporter\jsonl.go`.
- **Privacy coverage is incomplete:** sanitizer does not sanitize `Conversation.Metadata` / `Message.Metadata`, and audit does not check IP patterns despite redacting them during sanitization. `D:\code\llm-context-vault\pkg\models\types.go`, `D:\code\llm-context-vault\pkg\sanitizer\sanitizer.go`.
- Per-session extraction/storage errors are generally swallowed during scans, yielding incomplete imports without detailed reporting: `D:\code\llm-context-vault\pkg\vault\vault.go`.
- `--vault-dir` is advertised but not parsed/implemented; directory resolution includes a hard-coded Windows path, reducing portability: `D:\code\llm-context-vault\cmd\vault\main.go`.
- Generic extractor claims JSONL support but reads and unmarshals the whole file as one JSON document, so multi-record JSONL is not supported: `D:\code\llm-context-vault\pkg\extractor\generic.go`.
- `conversations\index.db` is untracked/ignored and no code references it; current search does not use it.
- The token setup script persists PATs in user environment variables and writes plaintext config values; it is operationally sensitive: `D:\code\llm-context-vault\scripts\set-github-token.ps1`.

## Git worktree state

- Worktree: `D:\code\llm-context-vault`, branch `main` tracking `origin/main`, HEAD `b21c50a`.
- Modified tracked file: `D:\code\llm-context-vault\.gitignore` (adds database ignore rules).
- Ignored untracked local artifacts: `D:\code\llm-context-vault\vault.exe`, `D:\code\llm-context-vault\vault.exe~`, `D:\code\llm-context-vault\conversations\index.db`.
- No files were modified during this review.

---

