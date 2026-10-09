# Role name replaceAll undefined crash

- **ID**: `opencode-b90db32777417ef58c922f17f815cf418d8a27a324828aaedf5ad1e19fa94885`
- **Source Tool**: `opencode`
- **Date**: `2026-10-02 09:37:28`
- **Project**: `qcash-ui-role-management`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

[Image 1] check this component. there is a slight chance i got error replaceAll is undefined crash when i tried to type the role name. can you investigate ?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>

I’ll trace the role-name input and every `replaceAll` call involved, then reproduce the failing value path and patch it with focused coverage.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The crash path is confirmed in `TemplateSelectionSection`: `systemConfig` can be absent while auth configuration is still loading, leaving `allowedCharacterRoleName` undefined, and both typing handlers immediately call `.replaceAll()` on it. The same code also depends on `String.prototype.replaceAll`, so I’ll remove both failure modes and add a regression test for typing before config arrives.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The regression test now passes and verifies that typing during the configuration-loading window no longer crashes; unsupported characters are still removed and repeated spaces collapse correctly. `git diff --check` only reports pre-existing trailing whitespace in `styles/globals.css`, which I did not modify. I’m running a focused lint check on the two changed files next.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 12: Assistant

Fixed the role-name crash.

**Root cause**
`allowed_character_role_name` can be temporarily unavailable while `systemConfig` loads. Typing then called `.replaceAll()` on `undefined`. Older browsers may also lack `replaceAll`.

**Changes**
- Default missing or invalid config values to `""`.
- Replaced `replaceAll()` with broadly supported regex-based `replace()`.
- Added regression coverage for typing before configuration loads.

**Verification**
- Focused tests: **5 passed**
- Lint only reported pre-existing issues.
- Unrelated existing workspace changes were left untouched.

---

## Turn 13: User

what happen when the allowed character is empty string?

---

## Turn 14: Assistant

When `allowedCharacterRoleName` is `""`:

- Allowed: `A-Z`, `a-z`, `0-9`, and a single space.
- Removed/blocked: all symbols such as `-`, `_`, `.`, `@`.
- Leading spaces are removed.
- Multiple spaces are collapsed into one.

The generated filter becomes:

```ts
/[^A-Za-z0-9 ]/g
```

Once `systemConfig` loads, configured special characters become allowed on the next render. However, symbols typed before loading will be blocked. If preferred, we can disable the field until config loads instead of using this fallback.

---

