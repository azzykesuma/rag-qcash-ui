# Audit staged provider change (@explore subagent)

- **ID**: `ses_0009aad9affeNijzEZU54FUtGd`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 15:30:58`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Review ONLY the staged diff in ~/Development/qcash-ui related to removing lazy loading from auth providers. Analyze behavioral/runtime regressions, SSR/client initialization, provider ordering, hooks/context assumptions, and test adequacy. Do not edit files. Return findings ordered by severity with exact file and line references in the current worktree, plus concise reasoning. Desired thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
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
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

## Findings

### High — Existing provider test suite now fails because its dynamic-import mock is obsolete
- `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:2`
- `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:52-55`
- `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:76-79`

`OrdentProviders` now imports the real package directly, but the existing tests still mock `next/dynamic` and expect that mock to supply provider components. Consequently, all three existing `OrdentProviders` tests instantiate the real auth provider rather than their spies. The assertions are no longer meaningful and currently fail before reaching them (`localStorage.getItem` throws in this test environment). Even where `localStorage` is available, the expected test IDs/spies would not come from `dynamicImportMock`.

Confirmed with:

```text
vitest run components/providers/__tests__/OrdentProviders.test.tsx
3 tests failed
```

The new staged test passes, but does not repair the existing failing suite.

### Medium — Tests do not exercise the SSR/client-boundary behavior that this change specifically alters
- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:26-38`
- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:50-96`
- `~/Development/qcash-ui/pages/_app.tsx:19`
- `~/Development/qcash-ui/pages/_app.tsx:161-166`

The new tests replace both packages with inert DOM wrappers and render only under jsdom. They therefore cannot detect:

- browser-global access during package evaluation or provider initialization,
- server-render failures or hydration differences,
- initialization relative to the remaining client-only `ModuleFederationProvider`,
- context identity problems involving actual package exports.

This is especially relevant because removing `{ ssr: false }` is the core change. In the current app, the providers still sit beneath the client-only `ModuleFederationProvider` at line 161, so they are not actually server-rendered through `_app`; however, the standalone providers are now SSR-capable components and no test verifies that contract. A server-render/import test using the real packages, plus an `_app` boundary test, would cover the changed behavior rather than only prop forwarding and nesting.

## Other reviewed areas

No staged regression found in provider ordering or context assumptions: the effective nesting remains BRI auth → BRI global store → Ordent auth → Ordent global store, and `AuthBridgeSync` remains beneath all four contexts. The installed packages also import and server-render successfully in a direct smoke check. TypeScript completed successfully with `tsc --noEmit`.

---

