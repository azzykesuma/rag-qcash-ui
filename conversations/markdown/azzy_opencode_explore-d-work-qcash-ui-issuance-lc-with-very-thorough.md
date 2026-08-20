# Inspect issuance routing (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_049ba44d1ffed1TVzl7eiD5MCd`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 10:44:11`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore D:\work\qcash-ui-issuance-lc with very thorough focus on its Next/router setup, auth bridge integration, module federation mount/export, route synchronization, rewrites/basePath, effects/listeners, and any guards that could block navigation away while URL changes. Research only, no edits. Explain whether the app handles host navigation away from /v2/import-lc/issuance-lc and identify root cause with exact paths/line refs plus minimal fix recommendation. Also inspect recent git status/log if useful via available tools.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 17: Assistant

## Verdict

The issuance-LC remote does **not independently synchronize itself to host URL changes**. It exports React page components and relies on the host’s Next router to replace/unmount them.

For normal host navigation using `router.push`, the checked integration should navigate away correctly:

- Host pages mount the remote through standard Next page components:
  - `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:15-18`
  - Child routes use separate wrappers at the equivalent `create`, `detail`, and `edit` paths.
- The host remote wrapper aborts loading when unmounted:
  - `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:68-76`
- Nothing in issuance-LC calls `preventDefault`, `beforePopState`, emits `routeChangeError`, or returns `false` from a routing listener.

Therefore, **issuance-LC does not contain a guard capable of changing the browser URL while deliberately canceling navigation**. If the entire old page remains after a completed host URL change, the cause is in the host navigation/mount lifecycle or an unresolved Next transition—not an unsaved-form guard in this repository.

## Most likely issuance-side cause of the observed “URL changed but page looks stuck”

The installed auth-provider version has a problematic global route listener:

- `D:\work\qcash-ui-issuance-lc\node_modules\@ordentco\addons-auth-provider\src\auth.tsx:769-791`
- Specifically:
  - It subscribes to `routeChangeStart` at lines `786-790`.
  - Every pathname change runs `setIsAuthoritiesReady(false)` at lines `780-783`.

Every exposed issuance page immediately replaces its content with a skeleton when that flag becomes false:

- List: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\index.tsx:45-51`
- Create: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\create\index.tsx:56-62`
- Detail: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:155-160`
- Edit: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:73-75,125-127`

Their `guard()` effects run only on initial mount (`[]`). Consequently, if a host transition is delayed, fails, or is canceled after `routeChangeStart`, the existing remote can remain mounted in a permanent skeleton state because it does not re-run `guard()`.

This listener does **not literally cancel navigation**, but it can create exactly the visual impression that the old module is stuck after the URL begins changing.

### Version mismatch

Issuance-LC pins the older provider:

- `D:\work\qcash-ui-issuance-lc\package.json:19` — `0.9.119-mfe`

The checked host pins:

- `D:\work\qcash-ui\package.json:27` — `0.9.138-mfe`

The newer provider no longer resets readiness in `routeChangeStart`:

- `D:\work\qcash-ui\node_modules\@ordentco\addons-auth-provider\src\auth.tsx:1064-1077`

Because both applications declare this dependency as a Module Federation singleton with `requiredVersion: false`:

- Remote: `D:\work\qcash-ui-issuance-lc\next.config.js:26-30`
- Host: `D:\work\qcash-ui\next.config.js:15-22`

the runtime version depends on share-scope/loading behavior. The mismatch makes routing/auth behavior deployment-dependent.

## Module Federation and routing contract

The remote exposes page modules, not a lifecycle API:

- `D:\work\qcash-ui-issuance-lc\next.config.js:18-24`

There is no `mount()`, `unmount()`, navigation callback, pathname prop, or host-route listener. This is valid for the checked host because the host owns actual Next pages and swaps the exposed React component.

`next/router` is automatically shared as a singleton by `nextjs-mf`:

- `D:\work\qcash-ui-issuance-lc\node_modules\@module-federation\nextjs-mf\README.md:61-72`

Thus internal unprefixed routes such as:

- `D:\work\qcash-ui-issuance-lc\containers\index.tsx:538`
- `D:\work\qcash-ui-issuance-lc\containers\index.tsx:988`
- `D:\work\qcash-ui-issuance-lc\components\create\form-footer.tsx:49`

use the host router and receive the host’s `/v2` base path when federated.

## Base paths and rewrites

Remote standalone base path:

- `D:\work\qcash-ui-issuance-lc\next.config.js:46-50` — `/issuance-lc-v2`

Host base path:

- `D:\work\qcash-ui\next.config.js:65-70` — `/v2`

There are no rewrites or redirects in issuance-LC. The host’s only rewrite is a development design-system CSS proxy:

- `D:\work\qcash-ui\next.config.js:120-144`

That is unrelated to issuance navigation.

The host federation registry correctly associates the remote with all issuance subpaths by prefix:

- `D:\work\qcash-ui\constants\features\registry.ts:516-520`

## Auth bridge integration

The remote does not consume `__QCASH_AUTH_BRIDGE__` or `qc-bridge-sync`; it directly calls `useAuth()`.

Its standalone `_app` wraps pages in an `AuthProvider`:

- `D:\work\qcash-ui-issuance-lc\pages\_app.tsx:48-50`

But `_app` is not exposed, so under federation the remote relies on the host’s shared Ordent provider.

In the checked host, issuance-LC is not a BRI-provider migration route:

- `D:\work\qcash-ui\components\providers\auth-provider-migration-routes.ts:1-14`
- Provider selection: `D:\work\qcash-ui\pages\_app.tsx:170-183`

Therefore `AuthBridgeSync` is not the normal auth path for `/v2/import-lc/issuance-lc`; the shared Ordent context is. The bridge is not the navigation blocker.

## Navigation guards and listeners found

No genuine navigation blocker was found.

- No `beforeunload`
- No `popstate`
- No `beforePopState`
- No `routeChangeError`
- No router-event exception/cancellation
- No use of `setOnLeaveAction` in issuance-LC

The create form does not register an unsaved-form leave guard. Its Cancel button merely opens its own confirmation modal and then navigates:

- `D:\work\qcash-ui-issuance-lc\components\create\form-footer.tsx:34-51`

One unrelated listener leak exists:

- `D:\work\qcash-ui-issuance-lc\components\create\page-3.tsx:450-452`

It adds a scroll listener during render and never removes it, but it cannot block routing.

## Separate broken “Back” destination

The list page’s Back link targets `/import-lc`:

- `D:\work\qcash-ui-issuance-lc\containers\index.tsx:958-977`

The checked host has no `pages/import-lc/index.tsx`; it only has issuance, incoming, and amendment child routes. Thus that Back link points at a nonexistent host route. It may result in a 404 or failed-looking transition, although it does not preserve the remote intentionally.

## Minimal fix recommendation

1. **Primary fix:** align issuance-LC with the host auth-provider version, preferably `0.9.138-mfe` or newer, and update its lockfile. This removes the `setIsAuthoritiesReady(false)` route-start behavior.
2. Keep the provider a singleton, but avoid `requiredVersion: false` across materially different versions; pin the same version in host and remote.
3. Do not add a remote `popstate` or `history` synchronization workaround. With standard host Next pages, the host should own unmounting.
4. Fix the Back link at `containers\index.tsx:958` to a real host route such as `/homepage`, or provide an explicit host callback for opening the Import LC selector.

## Git state

Issuance-LC local checkout:

- `master` is **behind `origin/master` by 762 commits**.
- Modified files:
  - `D:\work\qcash-ui-issuance-lc\containers\form.tsx`
  - `D:\work\qcash-ui-issuance-lc\containers\index.tsx`
  - `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\create\index.tsx`
  - `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\index.tsx`
  - `D:\work\qcash-ui-issuance-lc\styles\globals.css`

The relevant source changes are only temporary `console.log`/whitespace changes; none adds a navigation guard. I also inspected the current `origin/master` versions: they still pin `@ordentco/addons-auth-provider` `0.9.119-mfe`, so the auth-version issue remains in the latest fetched remote ref.

---

