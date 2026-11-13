# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-fa3864800847ffad5a0aaa4eb192a8012ce872cc68ff6640ecac602359f7e219`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 09:27:25`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

What is the vault sanitization policy for Windows and Unix machine file paths before publishing?

---

## Turn 2: Assistant

Based on the vault sanitization and privacy standards, the sanitization policy for machine file paths before publishing enforces:

1. **User Home & Username Masking**:
   - **Windows:** Absolute paths matching user directories (e.g., `C:\Users\<username>\...`) are scrubbed or replaced with a generic placeholder (such as `<USER_HOME>`, `C:\Users\<USER>\`, or project-relative paths).
   - **Unix / macOS:** Paths referencing user home directories (e.g., `/home/<username>/...` or `/Users/<username>/...`) are masked and replaced with `~` or `<USER_HOME>/`.

2. **Repository & Workspace Normalization**:
   - Host-specific absolute prefixes (drive letters, custom workspace paths, and temp directories like `C:\Temp\` or `/tmp/`) are stripped and normalized to project-relative paths (e.g., `./` or `<WORKSPACE>/`).
   - Path separators are typically normalized to consistent forward slashes (`/`).

3. **PII & Environment Redaction**:
   - Machine names, internal host/domain identifiers, internal corporate IDs, and any embedded credentials or authentication tokens residing in path strings are identified and redacted prior to vault publication.

---

