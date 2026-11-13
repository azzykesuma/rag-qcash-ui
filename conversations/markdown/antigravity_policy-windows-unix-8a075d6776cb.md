# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-e45f903718a52beb51e457ae87522dd5f36b9b802a96718bf0768a075d6776cb`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:29:07`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

What is the vault sanitization policy for Windows and Unix machine file paths before publishing?

---

## Turn 2: Assistant

In the vault sanitization and publishing pipeline, the path sanitization policy ensures no machine-specific identities, user credentials, or local environment details are leaked into shared or published context:

1. **User Home Directory Scrubbing**:
   - **Windows:** Absolute paths matching user profiles (e.g., `C:\Users\<username>\...` or `C:\Users\<employee_id>\...`) are scrubbed and replaced with generic placeholders (e.g., `~` or `<USER_HOME>`).
   - **Unix/macOS:** Home directories (e.g., `/home/<username>/...` or `/Users/<username>/...`) are similarly normalized to `~` or `<USER_HOME>`.

2. **Absolute Path & Drive Letter Normalization**:
   - Local drive prefixes (e.g., `C:\`, `D:\`) and machine-specific root mounts are stripped.
   - Code and project paths are converted to repository-relative paths (e.g., relative to the repository root) or prefixed with generic root markers (e.g., `<WORKSPACE>/...` or `<REPO_ROOT>/...`).

3. **System & Temporary Path Scrubbing**:
   - Paths referencing local temp directories (e.g., `AppData\Local\Temp`, `/tmp`, `/var/tmp`) are stripped or generalized to `<TEMP>/...`.
   - Local tool configuration and secrets directories (e.g., `.gemini`, `.ssh`, `.aws`, `.env`) are redacted or excluded from indexed and published context.

---

