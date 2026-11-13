# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-22b9b7e44ff81c16bd8faee2d3890cabeb6a856fac191d3535e09a2353da6433`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 09:30:53`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

What is the vault sanitization policy for Windows and Unix machine file paths before publishing?

---

## Turn 2: Assistant

Based on the vault sanitization policy for publishing conversational transcripts and artifacts:

1. **User Directory Redaction**:
   - **Windows:** Absolute paths matching user home directories (e.g., `C:\Users\<username>\...`) are stripped or replaced with standard placeholders such as `~` or `<HOME>`.
   - **Unix / macOS:** Paths referencing `/home/<username>/` or `/Users/<username>/` are normalized to `~/` or `<HOME>/`.

2. **Workspace & Repository Root Abstraction**:
   - Machine-specific workspace checkouts (e.g., `C:\code\...` or `/workspace/...`) are converted to relative repository paths or replaced with repository root tokens (e.g., `<repo-root>/...` or `<WORKSPACE>/`).

3. **Temporary Directories & Environment Paths**:
   - Machine-specific scratch paths such as `AppData\Local\Temp`, `/tmp`, and `/var/tmp` (or `/var/folders/...`) are masked or excised.

4. **Separator Normalization**:
   - Windows backslashes (`\`) in sanitized file paths are normalized to POSIX forward slashes (`/`) to maintain cross-platform uniformity and avoid escaping anomalies across markdown consumers.

---

