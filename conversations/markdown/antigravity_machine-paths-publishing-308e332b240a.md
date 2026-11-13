# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-537fb9fd7151b8fc277123cfb8a000cba624d76779bccc2857fd308e332b240a`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:32:52`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

What is the vault sanitization policy for Windows and Unix machine file paths before publishing?

---

## Turn 2: Assistant

The vault sanitization policy for machine file paths before publishing enforces strict redaction and normalization to prevent information disclosure (such as local usernames, machine hostnames, and proprietary local directory structures):

1. **User Directory Redaction (PII Protection)**:
   - **Windows**: Paths matching `C:\Users\<username>\...` (or other drive letters) have the personal user account identifier redacted or replaced with a generic placeholder (e.g., `<USER_HOME>`, `~`, or `[REDACTED]`).
   - **Unix/macOS**: Paths matching `/home/<username>/...` and `/Users/<username>/...` are scrubbed to strip specific machine account names.

2. **Workspace Root & Drive Letter Normalization**:
   - Drive letters (e.g., `C:`, `D:`) and absolute machine-specific workspace prefixes are normalized to generic root placeholders (e.g., `<WORKSPACE>/` or relative project paths) so conversations remain portable and agnostic of individual machine setups.

3. **Path Separator Uniformity**:
   - Windows backslashes (`\`) are normalized to POSIX-compliant forward slashes (`/`) across all indexed code snippets, logs, and stack traces to maintain cross-platform search consistency.

4. **Temporary and System Cache Scrubbing**:
   - Machine-specific scratch/temp directories (e.g., `AppData\Local\Temp`, `/tmp`, `/var/tmp`) and process IDs/session-bound hashes in path strings are sanitized or stripped.

---

