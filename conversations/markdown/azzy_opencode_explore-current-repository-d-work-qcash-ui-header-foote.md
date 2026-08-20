# Trace host route sync (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_049c214abffeOAcKy69pCZDs0e`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 10:35:39`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore the current repository D:\work\qcash-ui-header-footer with medium thoroughness. Find all auth bridge, URL/path synchronization, history/popstate, microfrontend navigation, menu click, and issuance-lc/import-lc integration code. Do research only; do not edit. Return likely root causes with exact file paths and line references, especially mechanisms that can change the browser URL without rerendering/navigating the child page.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

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

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
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

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 18: Assistant

## Summary

The strongest root cause is split routing ownership in a module-federated application. The header/footer remote performs navigation through its own `next/router`, but exposes no navigation callback or URL synchronization contract to the host/child microfrontend. A `router.push()` can therefore update browser history and the address bar while the child application remains on its previous rendered page.

This is particularly likely for Import LC issuance and incoming navigation, which use soft Next.js routing, while amendment and same-page cases use hard browser navigation and consequently reload correctly.

No files were edited. The repository already had an unrelated modification to `D:\work\qcash-ui-header-footer\styles\globals.css`; it was untouched.

## Likely root causes

### 1. No host/child navigation bridge

The remote exposes a layout that accepts only `children` and `isNewUI`; there is no `navigate`, route state, or host router callback:

- `D:\work\qcash-ui-header-footer\module-federation.config.mjs:9-13`
- `D:\work\qcash-ui-header-footer\layouts\default\index.tsx:4-5`

The project explicitly says it is intended to run inside a host:

- `D:\work\qcash-ui-header-footer\pages\index.tsx:5-11`

Nevertheless, menu navigation directly uses the remote component’s `next/router`:

- Router acquired: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:39`
- Soft navigation: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:92-98`
- Import LC soft navigation: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:97-113`

There is no production `popstate`, `routeChangeComplete`, or custom navigation listener in this repository. If the child application owns a separate router or microfrontend mount state, nothing tells it to switch pages after the header router changes history.

The federation configuration also does not explicitly share or inject `next/router`; its explicit shared modules are auth providers and i18n libraries:

- `D:\work\qcash-ui-header-footer\module-federation.config.mjs:17-37`

Plugin defaults may share some Next internals, but this repository provides no explicit routing guarantee.

### 2. Issuance LC uses the URL-only-prone branch

Import LC options are defined here:

- Issuance: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:61-67`
- Amendment: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:68-73`
- Incoming: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:74-79`

For a different issuance page, navigation is:

```ts
router.push(redirectUrl);
```

at:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:111-113`

`redirectUrl` is an absolute URL:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:98-100`

A same-origin absolute URL passed to the relevant Next router can be processed as internal navigation, writing browser history without forcing the independently mounted child application to remount.

By contrast:

- Clicking the already-current Import LC item sets `location.href`, forcing a reload: `LcImportModal.tsx:102-105`
- Legacy amendment sets `location.href`: `LcImportModal.tsx:108-110`

This can explain a pattern where amendment works but issuance changes only the address bar.

The production configuration makes the difference even more pronounced: amendment is disabled for preploy, piloting, and production, causing its value to become `"amendment"` and take the hard-navigation branch:

- `D:\work\qcash-ui-header-footer\config.ts:76-81`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:69-72`

Issuance always takes `router.push()` when it is not considered the current page.

### 3. Three incompatible path namespaces are mixed

The code mixes:

1. Menu trigger path: `/v2/import-lc`
2. Child/canonical route: `/import-lc/...`
3. Header-footer remote base path: `/header-footer`

References:

- Import LC modal opens only for exact `/v2/import-lc`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:68-70`
- Canonical Import LC routes omit `/v2`:
  `D:\work\qcash-ui-header-footer\features\other-modules\import-lc.route.ts:4-19`
- Remote Next.js base path:
  `D:\work\qcash-ui-header-footer\next.config.mjs:10-14`
- LC target explicitly removes `/v2`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:98-99`

The general menu handler also rewrites `/v2` paths based on substring checks:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:88-98`

Problems include:

- `href.includes("v2")` and `fullUrl.includes("v2")` are not path-segment checks.
- If already under a URL containing `v2`, `/v2/import-lc` is changed to `/import-lc`.
- Exact modal matching fails if backend `rootPath` or `BaseHreff` supplies `/import-lc` instead of `/v2/import-lc`.
- `config.baseFederation?.replace("/v2", "")` performs an unstructured string replacement.
- `baseFederation` is itself derived by removing `/api`:  
  `D:\work\qcash-ui-header-footer\config.ts:16-26`

### 4. Path-derived UI is not synchronized with raw History API changes

Different components read route state from different sources:

- Layout reads `globalThis.location.pathname` during render:
  `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:43-50`
- Menu selection reads `router.asPath`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:28,113-133`
- Breadcrumbs read `router.pathname`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\breadcrumb\utils\useRouterPathV2.tsx:30-56`
- LC current-page detection reads `router.asPath` and manually adds `/v2`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:100-102`

There is no common route state and no listener that rerenders these components after an external `history.pushState`. Since `pushState` does not emit `popstate`, even a conventional popstate-only child listener would not observe a Next router’s push.

The only direct `history.pushState` calls in the repository are test setup:

- `D:\work\qcash-ui-header-footer\layouts\default\container\__tests__\index.onboarding-tour.test.tsx:39-40`
- `D:\work\qcash-ui-header-footer\layouts\default\container\__tests__\index.loan-application.test.tsx:45-46`

There is no production `popstate` implementation.

### 5. Import LC “routes” only drive breadcrumbs

The Import LC route module does not mount or navigate the LC child. It supplies `handles`, `subPaths`, and breadcrumb definitions:

- `D:\work\qcash-ui-header-footer\features\other-modules\import-lc.route.ts:4-20`
- Registered under other breadcrumb modules:
  `D:\work\qcash-ui-header-footer\features\other-modules\index.ts:12,45`
- Consumed to calculate breadcrumbs:
  `D:\work\qcash-ui-header-footer\layouts\default\container\breadcrumb\utils\useRouterPathV2.tsx:30-40`

Consequently, adding `/import-lc/issuance-lc` here does not ensure that the host switches the rendered child microfrontend.

### 6. Tests verify router calls, not child navigation

The LC tests mock `next/router`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\LcImportModal.test.tsx:4-6`

They only verify that issuance calls `mockPush`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\LcImportModal.test.tsx:103-116`

The same-page and amendment hard-navigation tests merely assert that `router.push` was not called:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\LcImportModal.test.tsx:118-166`

The general menu test similarly mocks the router and verifies only the argument:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\MenuModalProvider.test.tsx:4-5,43-53`

There is no host-plus-child integration test confirming that the child content changes.

## Auth bridge findings

The bridge is a window-global auth snapshot, not a navigation bridge:

- Data contract: `D:\work\qcash-ui-header-footer\layouts\default\AuthBridge.tsx:3-50`
- Fallback values: `D:\work\qcash-ui-header-footer\layouts\default\AuthBridge.tsx:52-82`
- Reads `window.__QCASH_AUTH_BRIDGE__`: `D:\work\qcash-ui-header-footer\layouts\default\AuthBridge.tsx:84-99`
- React hook listens only for `qc-bridge-sync`:  
  `D:\work\qcash-ui-header-footer\hooks\useAuthBridge.ts:5-28`

Notable risks:

- The global object is non-reactive; the host must dispatch `qc-bridge-sync` after every update.
- The exported `AUTH_BRIDGE_EVENT` constant is unused by the hook, which hardcodes the string.
- No writer or dispatcher for the auth bridge exists in this repository.
- The bridge contains no route or navigation field.
- `isBridgeActive` is read directly from the global during render and cannot itself trigger rerender:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:37-39`
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:41-48`

Auth sources are also inconsistent:

- Layout and LC modal fall back safely to Ordent auth.
- `MenuModalProvider` uses only the bridge for leave-page state:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:37-39,115-139`
- `useMenu` mixes bridge identity/token with BRI/Ordent global stores and localStorage:
  `D:\work\qcash-ui-header-footer\hooks\useMenu.tsx:183-203`
- Header dropdown still directly uses Ordent auth/global store:
  `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\ListItem.tsx:44-50`

These inconsistencies can produce stale menu/authority state, although they are less directly responsible for the URL-only symptom.

## Navigation mechanisms inventory

### Soft navigation — capable of URL change without child remount

- General menu: `MenuModalProvider.tsx:92-98`
- Import LC issuance/incoming: `LcImportModal.tsx:111-113`
- LC Export: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcExportModal.tsx:51-54`
- Local Taxes: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LocalTaxesModal.tsx:34-43`
- Account Receivable: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\AccountReceivableUploadMandateModal.tsx:41-46`
- Payroll: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\PayrollModal.tsx:28-31`
- Profile: `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\ListItem.tsx:82-85`
- Breadcrumb `router.push`/`Link`: `D:\work\qcash-ui-header-footer\layouts\default\container\breadcrumb\components\Breadcrumb.tsx:43-51,73-80`
- Complaint `Link`: `D:\work\qcash-ui-header-footer\components\contact-support\index.tsx:80-95`

### Hard navigation — expected to reload/remount

- General non-v2 menu: `MenuModalProvider.tsx:100-105`
- LC same-page and amendment: `LcImportModal.tsx:102-110`
- Global external transfer: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\GlobalExternalTransferModal.tsx:53-64`
- Header anchor links: `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\NavLink.tsx:52-55`
- Auth failure redirects: `D:\work\qcash-ui-header-footer\services\custom-axios.ts:41-55,85-89`

## Most probable failure sequence

1. Backend menu data resolves Import LC to `/v2/import-lc`.
2. `MenuModalProvider` opens `LcImportModal`.
3. User selects Issuance LC.
4. The modal constructs an absolute `/import-lc/issuance-lc` URL after removing `/v2`.
5. It invokes the header remote’s `router.push`.
6. Next/router updates browser history/address bar.
7. The child LC microfrontend has no shared navigation callback, route event, synthetic popstate event, or reactive URL observer.
8. The old child page stays rendered despite the new browser URL.

---

