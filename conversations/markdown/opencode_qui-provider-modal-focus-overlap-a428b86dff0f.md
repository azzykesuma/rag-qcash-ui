# Fix QUi Provider modal focus and overlap

- **ID**: `opencode-a94f651def602d00e40cfc3c4439233fa30d34776ada721d4d8aa428b86dff0f`
- **Source Tool**: `opencode`
- **Date**: `2026-09-17 16:46:11`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

there is an issue with the modal placement and the overlap with modal dialog in qui provider
1. when the session expired modal appear on top of the qui provider, the session expired is not interactive, meaning that, even if the session expired appear in front of the modal dialog, the interaction belong to the modal dialog behind it

i assume that this is the aria focus trap, can we fix it

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

I’ll trace both modal implementations and the provider’s focus/portal behavior, then reproduce the stacking and interaction path before making the smallest fix. I’ll also check the project’s prior context for any modal or QUi decisions.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The session modal already has a dedicated top-level portal and removes `inert`, but that only fixes pointer blocking. The remaining symptom matches a document-level focus scope from the QUi dialog redirecting focus back to the underlying dialog; I also found uncommitted edits in this same file that comment out event-capture guards, so I’ll preserve that work while validating which event boundary actually controls the trap.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The root cause is confirmed: the session UI is only styled as a `<dialog>`; it never calls `showModal()`, so it is not in the browser’s modal top layer and does not make the QUi dialog inert. Removing `aria-hidden` and raising `z-index` changes appearance but cannot defeat QUi’s active focus scope; I’ll activate the native modal lifecycle instead and remove the capture handlers that would prevent child controls from receiving events.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

The targeted test run exposed only a JSDOM limitation: it does not implement `HTMLDialogElement.showModal()`. I’m adding a test-only dialog polyfill and assertions that verify the native modal opens/closes and that child focus/pointer handlers remain reachable, which directly guards against the original regression and the previously commented capture approach.

---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The regression coverage now passes: the session dialog is opened through the browser modal top layer, closes cleanly with state changes, remains non-dismissible by Escape, and its input controls receive pointer/focus events. I’m doing the final worktree and whitespace review now; no provider-wide changes were needed.

---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 20: Assistant

Fixed the modal interaction conflict.

- `SessionExpiredNewLayout` now calls native `dialog.showModal()`, placing it in the browser top layer and making the underlying QUi dialog inert.
- Removed capture handlers that blocked events from reaching session controls.
- Prevented Escape from dismissing the session modal.
- Added regression tests for modal lifecycle, pointer/focus interaction, and Escape handling.

Verification:
- 74 targeted tests passed.
- ESLint passed.
- `git diff --check` passed.

---

