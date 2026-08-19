# Inspect issuance routing (@explore subagent)

- **ID**: `ses_0261875d0ffegFFIInx7td3mS4`
- **Source Tool**: `opencode`
- **Date**: `2026-08-07 08:47:37`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Analyze D:\work\qcash-ui-issuance-lc for routing/navigation behavior that could cause this exact issue when federated into a Next.js Pages Router host: entering Issuance LC works; clicking host header links changes browser URL but page remains Issuance LC. Search for browser history monkeypatches, beforeunload/popstate listeners, router events, navigation blockers, React Router/Next router use, global event listeners, unsaved-change guards, layout portals, and cleanup bugs. Also inspect Module Federation config/shared dependencies and exposed modules. Compare branch/recent code if useful via git log but do not edit. Return ranked root cause(s), exact file/line references, and a minimal fix recommendation. Thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

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

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 28: Assistant

## Conclusion

The exact symptom is most consistent with navigation being performed by the federated header’s `next/router` instance rather than the host Pages Router instance:

1. `router.push()` updates `window.history`, so the URL changes.
2. The host Pages Router does not receive/complete the route transition.
3. Its current page component—Issuance LC—therefore remains mounted.

There is no active navigation blocker in the Issuance LC page itself. A previously committed header/footer fix explicitly recognizes this exact Issuance-LC-only problem and forces a document navigation when leaving it.

## Ranked root causes

### 1. Known header navigation/context defect; its hard-navigation fix is currently disabled

Confidence: very high.

The header/footer repository contains an Issuance-specific hard-navigation list:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:37-39`

```ts
const REGISTERED_HARD_NAV_PAGES = [
  "/v2/import-lc/issuance-lc",
];
```

The corresponding workaround is currently commented out in the working tree:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:93-101`

Consequently, navigation falls through to:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:108-117`

```ts
router.push(`/${_href}`);
```

The header gets that router from its own federated bundle:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:2`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:43`

This is precisely capable of producing “URL changes, old page remains” when its `next/router` resolves to a different router/module context than the host’s.

This was already fixed by:

- Commit `8161f7cc`: `fix: document navigation when leaving issuance LC`
- Commit `8a867362`: `fix[NQLA-60919] - adjust navigation issue next route missing context`

The committed fix was:

```ts
if (globalThis.window.location.pathname.startsWith("/v2/import-lc/issuance-lc")) {
  globalThis.window.location.href = href;
  return;
}
```

The current header/footer working tree has locally commented that fix out while adding navigation diagnostics. If that state reflects the failing deployment, this is the immediate cause. Otherwise, verify that the deployed/cached header `remoteEntry.js` actually contains `8161f7cc` or later.

### 2. Multiple Next/Module Federation runtime contexts make the header defect possible

Confidence: high as the underlying structural cause.

The Issuance remote exposes raw Pages Router page modules:

- `D:\work\qcash-ui-issuance-lc\next.config.js:27-33`

It deliberately does not expose complete Next pages:

- `D:\work\qcash-ui-issuance-lc\next.config.js:34-36`

Those exposed modules still import Next routing primitives:

- `D:\work\qcash-ui-issuance-lc\containers\new-skin.tsx:8`
- `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:5,32`
- `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:9,28`
- Numerous internal `next/link` and `router.push()` calls under `components` and `containers`.

There is significant runtime/version skew:

| Application | Next | nextjs-mf | MF runtime |
|---|---:|---:|---:|
| Issuance remote | 13.5.6 | 8.2.5 | `^0.1.1` |
| Pages host | 13.5.11 | 8.2.6 | 0.18.1 |
| Header/footer | 13.5.6 | 8.2.5 | `^0.2.3` |

References:

- `D:\work\qcash-ui-issuance-lc\package.json:19-20,34,38-40`
- `D:\work\qcash-ui\package.json:24-26,40-43`
- `D:\work\qcash-ui-header-footer\package.json:22-23,39-45`

The applications also have different base paths:

- Issuance: `D:\work\qcash-ui-issuance-lc\next.config.js:69-70` — `/issuance-lc-v2`
- Host: `D:\work\qcash-ui\next.config.js:66-70` — `/v2`
- Header/footer: `D:\work\qcash-ui-header-footer\next.config.mjs:10-14` — `/header-footer`

The remote also calls `init()` under its own application name from several places:

- Standalone `_app`: `D:\work\qcash-ui-issuance-lc\pages\_app.tsx:26-45`
- Create page: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\create\index.tsx:107-171`
- Edit page: `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:87-119`

Repeated federation initialization and mixed runtime versions increase the chance that a federated `useRouter()` does not use the host’s router singleton.

### 3. Auth-provider singleton/version skew adds a route-event side effect

Confidence: medium.

The remote uses an older auth provider:

- `D:\work\qcash-ui-issuance-lc\package.json:22` — `0.9.119-mfe`

It is shared with `requiredVersion: false`:

- `D:\work\qcash-ui-issuance-lc\next.config.js:50-53`

The installed remote version itself depends on Next 12:

- `D:\work\qcash-ui-issuance-lc\node_modules\@ordentco\addons-auth-provider\package.json:30`

It imports and uses `next/router`:

- `D:\work\qcash-ui-issuance-lc\node_modules\@ordentco\addons-auth-provider\dist\src\auth.js:74,175`

It also subscribes to every route start and clears authorization readiness:

- `D:\work\qcash-ui-issuance-lc\node_modules\@ordentco\addons-auth-provider\dist\src\auth.js:849-868`

The list page responds to that state by replacing its contents with a skeleton:

- `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\index.tsx:55-57`

This is not sufficient by itself to explain a fully unchanged page, but it introduces another Next router instance and route-transition side effect. Because the package is a singleton with no version constraint, runtime selection order matters.

### 4. Current `_app` history monkeypatch is risky, but not the original federated cause

Confidence: low for this incident.

The dirty Issuance working tree currently wraps global history:

- `D:\work\qcash-ui-issuance-lc\pages\_app.tsx:97-109`

It restores captured methods during cleanup:

- `D:\work\qcash-ui-issuance-lc\pages\_app.tsx:124-125`

The wrappers delegate correctly, so they do not directly block navigation. However, cleanup can overwrite a wrapper installed later by the host or another MFE.

More importantly:

- This code is not in tag `1.0.260`; it is an uncommitted diagnostic change.
- `_app.tsx` is not one of the exposed modules in `next.config.js:27-33`.

Therefore it should not run merely because the host imports `qcash-ui-issuance-lc/issuance-lc`. Remove the monkeypatch after diagnosis, but it is not the primary cause.

### 5. Cleanup leaks exist, but not on the list page involved in this reproduction

Confidence: low.

Three page-3 implementations add a scroll listener during render and never remove it:

- `D:\work\qcash-ui-issuance-lc\components\create\page-3.tsx:441-443`
- `D:\work\qcash-ui-issuance-lc\components\create-new-old\page-3.tsx:1003-1005`
- `D:\work\qcash-ui-issuance-lc\components\create-new-skin\page-3.tsx:1011-1013`

This creates one new listener per render and leaves stale state setters after unmount. It should be moved into a `useEffect` with cleanup, but it is not active on the Issuance list page.

Other listeners inspected clean up correctly:

- `D:\work\qcash-ui-issuance-lc\components\create-new-skin\page-4.tsx:222-230`
- `D:\work\qcash-ui-issuance-lc\components\ui\QuiCalenderPicker.tsx:59-67`
- `D:\work\qcash-ui-issuance-lc\components\ui\qui-currency-amount.tsx:123-136`

## Negative findings

I found no Issuance implementation of:

- `beforeunload` or `onbeforeunload`
- `beforePopState`
- React Router, `BrowserRouter`, `MemoryRouter`, or navigation blockers
- Unsaved-change prompts or route-change exceptions
- `popstate` handlers that prevent or reverse navigation
- Global click handlers intercepting header links
- A permanent full-screen portal on the Issuance list page

Portals exist for selects/modals, for example:

- `D:\work\qcash-ui-issuance-lc\components\ui\table\table-footer.tsx:75-87`

But no list-page portal code explains the URL changing while the host route remains unchanged.

## Minimal fix

Restore the already committed header/footer hard-navigation branch:

```ts
if (window.location.pathname.startsWith("/v2/import-lc/issuance-lc")) {
  window.location.assign(href);
  return;
}
```

Location:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:93-101`

This is the smallest, already-tested fix and guarantees that the Issuance remote unmounts.

## Recommended durable fix

1. Make the host the sole owner of cross-feature navigation.
   - Pass a host `navigate()` callback into the header remote.
   - Do not let the header resolve and call its own federated `next/router`.
2. Align Next, `@module-federation/nextjs-mf`, and `@module-federation/runtime` versions across host, header, and Issuance.
3. Remove repeated page-level `init()` calls; initialize federation once in the host.
4. Align the auth-provider version and remove its dependency on `next/router`.
5. Make exposed Issuance components router-agnostic where possible; use host callbacks for navigation outside Issuance.
6. Fix the three uncleaned scroll listeners.

## Quick confirmation

Using the current diagnostics:

- If `history.pushState` logs but the **host** emits no `routeChangeStart`, the header is using a foreign/direct history path.
- If `routeChangeStart` appears but neither `routeChangeComplete` nor Issuance’s unmount log appears, inspect the router context or route-load error.
- The relevant unmount diagnostic is:
  - `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\index.tsx:45-52`

No files were edited.

---

