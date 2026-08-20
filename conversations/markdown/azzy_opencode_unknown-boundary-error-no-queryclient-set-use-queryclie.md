# QueryClientProvider missing error investigation

- **Author**: `azzy`
- **ID**: `ses_02a844d58ffeyiPAq7sAy4ATqv`
- **Source Tool**: `opencode`
- **Date**: `2026-08-06 12:11:21`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

unknown-boundary Error: No QueryClient set, use QueryClientProvider to set one
    at useQueryClient (6612.76886c3b.js:1:25051)
    at useBaseQuery (1593.7551be96.js:1:15159)
    at useQuery (1593.7551be96.js:1:17072)
    at useDataTableQuery (__federation_expose_components__data_table.bd94d1c9.js:1:53044)
    at DataTableRemoteContent (__federation_expose_components__data_table.bd94d1c9.js:1:78580)
    at Nh (__federation_expose_react_dom.9dbcb3173fd3634c.js:9:58490)
    at x (__federation_expose_react_dom.9dbcb3173fd3634c.js:9:119387)
    at Uk (__federation_expose_react_dom.9dbcb3173fd3634c.js:9:99077)
    at __federation_expose_react_dom.9dbcb3173fd3634c.js:9:98944
    at Ik (__federation_expose_react_dom.9dbcb3173fd3634c.js:9:98951) Object  check why i got this error

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll trace the application providers, the federated data-table integration, and TanStack Query dependency sharing to identify where the query context is being lost.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The failing hook is inside the federated `qui` DataTable, not this app’s own query hook. The page does wrap the table, but `@tanstack/react-query` is absent from the federation `shared` configuration, so the provider and the remote hook can resolve to different React Query module instances and therefore different contexts. I’m checking the remote provider API and every affected route before applying the fix.

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

There are two concrete provider defects:

1. The list page removed `QuiProviders` in commit `962a9ad`, while `qui/components/data-table` now executes React Query from the design-system remote. The local `QueryClientProvider` cannot satisfy that remote context when React Query is duplicated.
2. The detail page calls `useIssuanceDetailQuery` before its provider is rendered, so that route also depends accidentally on an outer host provider.

I’ll restore the remote provider around the list using the same client, and move the detail provider above its query hook.

---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The provider hierarchy is now explicit for both query implementations, and the detail hook is guaranteed to execute under its provider. I’m running the focused page tests and TypeScript validation to catch provider prop or lifecycle regressions.

---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id
- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **context7_query-docs**: Tool: context7_query-docs
- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 18: Assistant

**Root Cause**
`DataTable` is loaded from the federated `qui` design system and uses its own React Query context. The page only provided the application’s `QueryClientProvider`, so the remote `useDataTableQuery` could not find a client.

The host currently resolves React Query `5.100.11`, while `qui` declares `5.96.2`, increasing the likelihood of separate contexts.

**Fixes**
- Wrapped the list with `QuiProviders` and passed the existing `queryClient` in `pages/import-lc/issuance-lc/index.tsx`.
- Moved the detail page provider above `useIssuanceDetailQuery`, which was previously called outside its provider.
- Added test coverage for the `QuiProviders` hierarchy.

**Verification**
- List page tests: 4 passed.
- Detail page test: 1 passed.
- `git diff --check`: passed.
- Full TypeScript validation remains blocked by unrelated existing test typing errors; none reference the changed files.

---

## Turn 19: User

can you switch to master and check if this issue happen as well?

---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The current branch has the three uncommitted provider fixes from the previous step. I’ll preserve them, inspect `master` in an isolated worktree, and verify the provider structure there without contaminating either branch.

---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The workspace is currently detached at `c77dc44`, not on a named branch, and `styles/globals.css` also changed independently of my provider edits. I’ll leave this worktree untouched and check out `master` into the approved temporary workspace.

---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

`master` has the same vulnerable structure: the list uses only the local `QueryClientProvider`, `@tanstack/react-query` is not shared by federation, and the detail hook runs before its provider. This means the same DataTable error can occur on `master`; I’m confirming there is no branch-specific difference in those files.

---

## Turn 25: Assistant

Yes, `master` at `893b9cf` has the same issue:

- List page lacks `QuiProviders` around the federated DataTable.
- React Query is not configured as a shared federation singleton.
- Detail page calls `useIssuanceDetailQuery` before rendering `QueryClientProvider`.
- Relevant files are identical between `c77dc44` and `master`.

I checked `master` in an isolated worktree because the current workspace has uncommitted changes, including `styles/globals.css`. Those changes remain untouched.

---

