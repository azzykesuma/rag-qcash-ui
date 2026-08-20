# Issuance LC navigation issue

- **Author**: `azzy`
- **ID**: `ses_049c2d650ffev6DVonzwHa9G1G`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 10:34:49`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

okay there is some issues i am having
1. when i navigate to http://internal-service.example.com/v2/import-lc/issuance-lc, then i tried to navigate away by clicking other menu, the url changes  but the page didn't navigate away
2. this only happen to that particular url, any other url navigation works just fine

you can check
1. qcash ui for the auth bridge sync
2. check issuance lc repo

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll trace the auth-bridge route synchronization in this UI, identify the issuance-LC integration point, then inspect the issuance-LC repository if it is available alongside this workspace.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **task**: Tool: task
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **task**: Tool: task
- **task**: Tool: task
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>

The header/footer uses `next/router.push`, which can update history without remounting an independently routed microfrontend. I’m now checking the main `qcash-ui` bridge synchronization and the issuance-LC app’s route listeners to determine which side is failing to react, rather than applying a blanket hard reload prematurely.

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
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Two concrete inconsistencies are now confirmed: the issuance remote pins `@ordentco/addons-auth-provider@0.9.119-mfe`, whose global `routeChangeStart` handler marks authorities unready on every navigation, while the shell/header use `0.9.138-mfe`, where that behavior was removed. I’m checking the deployed bundles and route transition behavior before deciding whether the fix belongs only in issuance-LC or also needs a defensive header fallback.

---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **task**: Tool: task
- **task**: Tool: task
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **task**: Tool: task

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The deployed shell has no route blocker, and the auth bridge only synchronizes auth state, not navigation. The issuance remote is the outlier: it publishes an older auth singleton with route-start side effects; I’ll make header navigation resilient when leaving that route by forcing a shell-owned document navigation, while preserving soft routing everywhere else, and add a regression test.

---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The regression test passes with the new hard-navigation branch. I’m running the targeted lint/test once more and checking the final worktree to ensure the pre-existing generated CSS change remains untouched.

---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 20: Assistant

Implemented a targeted fix in `MenuModalProvider.tsx`:

- Leaving `/v2/import-lc/issuance-lc` now uses full document navigation.
- Other routes retain existing soft navigation through `router.push`.
- Added a regression test covering this behavior.

Root cause findings:

- `AuthBridgeSync` synchronizes authentication state only; it does not synchronize navigation.
- Issuance-LC deploys `@ordentco/addons-auth-provider@0.9.119-mfe`, while the deployed shell uses `0.9.142-mfe`.
- The older version resets authority readiness during `routeChangeStart`, making issuance-LC uniquely vulnerable to incomplete soft transitions.

Verification:

- Targeted tests: 4 passed.
- `git diff --check`: passed.
- Targeted lint is blocked by pre-existing CRLF and formatting violations in these files.
- Existing `styles/globals.css` changes were untouched.

Long term, issuance-LC should align its auth-provider version with the shell.

---

