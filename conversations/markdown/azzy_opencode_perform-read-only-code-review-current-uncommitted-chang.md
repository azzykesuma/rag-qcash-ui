# Review auth gate changes (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_fc3fb6155ffeUst5JdsZk7OlRU`
- **Source Tool**: `opencode`
- **Date**: `2026-08-26 10:02:10`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a read-only code review of the current uncommitted changes in D:\work\qcash-ui related to HostAuthGate centralization. Inspect git diff and surrounding code. Focus on correctness bugs, infinite effects/retries, token races, public/protected routing, BRI/Ordent behavior, event handling, and whether AuthBridgeSync truly remains untouched. Do not edit. Return findings ordered by severity with exact file:line references; say explicitly if no blocking issues.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>

I’ll inspect the worktree diff first, then trace the changed auth/routing code and verify whether `AuthBridgeSync` has any uncommitted modifications. No files will be changed.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **bash**: Tool: bash

</details>

The central gate now controls mounting of everything beneath it, so I’m tracing provider nesting and route ownership next. This is important because a gate cannot hydrate a provider that is mounted inside the children it is withholding.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The provider implementations reveal that `guard()` can resolve after internally clearing or refreshing the browser token, rather than consistently rejecting. I’m validating the resulting gate behavior against the targeted tests and checking whether current tests cover that real provider contract.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

**Findings**
1. **High, blocking: token replacement does not cancel the previous provider-validation sequence.** `components/providers/HostAuthGate.tsx:193-206` blocks new validation whenever any token is in flight, while `components/providers/HostAuthGate.tsx:208-210` checks token freshness only after both guards finish. On BRI routes, an Ordent guard for token A can finish after token B arrives and still invoke the BRI guard using the captured token-A context (`components/providers/HostAuthGate.tsx:74-81`). Because these provider guards can refresh or clear tokens internally, the stale sequence can rotate or invalidate the new session before token B is validated.

2. **Medium: a failed redirect leaves protected routes on an unrecoverable loading screen.** `redirectingRef` is set before navigation but never cleared when `router.replace()` rejects at `components/providers/HostAuthGate.tsx:163-167`. The unauthorized path has the same problem at `components/providers/HostAuthGate.tsx:223-225`. Subsequent effects see `redirectingRef.current === true`, suppress further attempts, and no error UI is shown.

3. **Medium: session re-login can silently bypass gate invalidation when `setToken` throws.** The `token-change` event is inside the same `try` block after `setToken` at `hooks/use-modal-session-expired.tsx:109-114`; the legacy branch similarly cannot dispatch if the setter throws at `hooks/use-modal-session-expired.tsx:152-153`. Local storage has already received the new token, but same-tab storage writes emit no `storage` event. The gate can therefore remain validated against the old session while requests use the new token.

4. **Low: BRI route ownership uses an unbounded prefix match.** The gate applies BRI validation at `components/providers/HostAuthGate.tsx:109`, while `components/providers/auth-provider-migration-routes.ts:13-14` treats paths such as `/helpdesk-old` or `/systems-admin` as BRI-owned. This should use the same exact-or-child boundary logic used for public routes unless every such prefix is intentional.

**Verification**
- Blocking issues found: **1**.
- `AuthBridgeSync` truly remains untouched: neither staged nor unstaged diffs include `components/providers/AuthBridgeSync.ts`.
- Targeted tests passed: 4 files, 27 tests.
- ESLint reported no errors; warnings were line-ending/Prettier warnings.

---

## Turn 14: User

Re-review the updated uncommitted diff after the reported fixes. Confirm whether the stale-token-to-BRI blocking race and redirect/token-notification findings are resolved. Look for any new blocking correctness issue. Read-only, no edits. Return concise findings with file:line references.

---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’ll compare the updated diff against the previously reported paths, then inspect the surrounding gate/provider behavior and rerun the targeted tests without editing.

---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>

The stale Ordent-to-BRI handoff is now guarded correctly, but the redirect failure handling introduces a retry loop because the error state itself retriggers the no-token effect. I’m checking exact lines and test coverage for that path before finalizing.

---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

**Findings**
1. **High, blocking: failed unauthenticated redirects now retry indefinitely.** The rejection handler resets `redirectingRef` and sets `authError` at `components/providers/HostAuthGate.tsx:173-176`. Because `authError` is an effect dependency (`components/providers/HostAuthGate.tsx:262`), the effect reruns, enters the no-token branch before checking `authError` (`components/providers/HostAuthGate.tsx:166-179`), and calls `router.replace()` again. Unauthorized redirect failures feed into the same loop via `components/providers/HostAuthGate.tsx:223-225`.

**Confirmed Resolved**
- Stale Ordent validation no longer proceeds into BRI after token replacement: token is rechecked between guards at `components/providers/HostAuthGate.tsx:77-85`.
- Redirect failures now produce an error state, but the retry loop above prevents this fix from being correct.
- Token notification now occurs even if `setToken` throws at `hooks/use-modal-session-expired.tsx:109-114` and `hooks/use-modal-session-expired.tsx:152-157`.
- `AuthBridgeSync` remains untouched.
- Targeted suite passed: 28 tests across 4 files.

---

## Turn 19: User

Final read-only review after moving the authError check before the no-token redirect and adding the regression test. Confirm whether any blocking correctness findings remain in the current diff. Be concise and include exact references for any remaining issue.

---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>

I’ll verify the reordered effect and regression test in the current diff, then run the targeted auth tests and check for any remaining blocking path.

---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 22: Assistant

No blocking correctness findings remain.

Confirmed:
- Redirect retry loop fixed by checking `authError` before no-token redirect: `components/providers/HostAuthGate.tsx:166-171`.
- Stale token cannot proceed from Ordent into BRI: `components/providers/HostAuthGate.tsx:77-85`.
- Token-change notification remains outside setter error handling: `hooks/use-modal-session-expired.tsx:109-114`, `hooks/use-modal-session-expired.tsx:152-157`.
- Regression test covers failed redirect without automatic retry: `components/providers/__tests__/HostAuthGate.test.tsx:122-129`.
- Targeted suite passed: 29 tests across 4 files.

---

