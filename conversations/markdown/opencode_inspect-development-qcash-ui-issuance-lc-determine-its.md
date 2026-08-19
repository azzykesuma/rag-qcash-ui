# Audit issuance routing (@explore subagent)

- **ID**: `ses_ffefda48dffegLSc12PwBCZYNE`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 23:02:06`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Inspect ~/Development/qcash-ui-issuance-lc. Determine its current routing/navigation architecture, especially use of Next router versus TanStack Router and QUI design-system navigation components. Trace remote exports/custom entry and identify concrete navigation coupling bugs or assumptions when mounted in qcash-ui. Do not edit. Return findings with exact file/line references and explain whether host auth-provider lazy/static loading could plausibly fix them. Desired thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
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

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

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

## Turn 8: Assistant

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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 15: Assistant

# Executive summary

`qcash-ui-issuance-lc` is a **Next.js 13 Pages Router micro-frontend**. It does not use TanStack Router. Its current production paths use:

- `next/router` for imperative navigation and dynamic route parameters.
- `next/link` for declarative navigation.
- QUI primarily for visual controls such as `QuiButton`, `DataTable`, and dropdowns—not for routing.
- Module Federation page exports consumed by matching Next Pages Router proxy pages in `qcash-ui`.

When mounted in `qcash-ui`, the remote assumes that the host supplies a compatible Next router context and compatible shared Next internals. That assumption is especially important because:

- The host has `basePath: "/v2"`.
- The remote standalone app has `basePath: "/issuance-lc-v2"`.
- All business navigation paths are hard-coded without either prefix.

There is no Next-versus-TanStack Router bridge. A TanStack Router host would not be compatible without an explicit navigation adapter and explicit route-param props.

Changing host auth-provider loading from lazy to static could plausibly fix an **auth-context/share-scope race or duplicate auth-provider instance**, but it cannot fix the navigation coupling described below. The current host already statically imports and mounts its auth providers before rendering the remote, so further auth-provider loading changes are not a credible fix for current link, router, base-path, or QUI interaction problems.

---

# 1. Routing architecture in the remote

## 1.1 It is a Next Pages Router application

Evidence:

- Next 13.5.6 dependency:
  - `~/Development/qcash-ui-issuance-lc/package.json:34`
- Pages directory and route files:
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx`
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx`
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx`
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx`
- `_app.tsx` uses `AppProps` and `NextPage` from the Pages Router APIs:
  - `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:3-4`
  - `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:18-24`

The active dynamic pages read route state through `next/router`:

- Detail:
  - import: `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:5`
  - `useRouter()`: `:32`
  - `router.query.transactionEventId`: `:45-50`
- Edit:
  - import: `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:7`
  - `useRouter()`: `:26`
  - `router.query.transactionEventId`: `:27`

The active list, form, and detail UI also navigate directly through Next:

- List links:
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:8`
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:65-69`
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:75-79`
- Form navigation:
  - `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:2`
  - `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:85`
  - `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:1242-1245`
  - `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:1393-1399`
- Detail navigation:
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:13`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:83`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:267-269`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:320-322`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:491-494`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:650-652`
- Table links:
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:9`
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:51-60`
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:196-203`

## 1.2 There is no TanStack Router

An exhaustive source search found no:

- `@tanstack/react-router`
- `@tanstack/router`
- `RouterProvider` from TanStack
- `useNavigate`
- TanStack route-tree creation

This is true in both the remote and current host.

The only TanStack dependency in the remote is React Query:

- `~/Development/qcash-ui-issuance-lc/package.json:23`

Examples of its use:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:6-7`
- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:38-40`
- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:7-9`

Therefore, any diagnosis that attributes current routing to TanStack Router is incorrect. If another branch or future host mounts this under TanStack Router, the current remote will not automatically participate in that router.

---

# 2. QUI design-system navigation usage

## 2.1 QUI is not the routing layer

The remote dynamically loads QUI components:

- QUI remote registration:
  - `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:29-38`
- QUI global/provider loading:
  - `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:40`
  - `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:43-71`
- Component loader:
  - `~/Development/qcash-ui-issuance-lc/components/global.tsx:568-590`

But it does not use QUI navigation components such as:

- `Breadcrumbs`
- `Sidebar`
- `HeaderItem`
- `MenuItem` as a route abstraction

Type declarations for those components exist under `@mf-types`, but there are no corresponding production imports/usages in this feature.

The QUI provider itself has no navigation callback or router prop. Its available props are query-client and transaction-authorization configuration:

- `~/Development/qcash-ui-issuance-lc/@mf-types/qui/node_modules/@qui/components/.rslib/declarations/esm/providers/provider.d.ts:5-10`

Thus QUI cannot translate the remote’s route operations into host navigation.

## 2.2 QUI buttons are placed inside Next links

The remote commonly composes navigation as:

```tsx
<Link href="...">
  <QuiButton>...</QuiButton>
</Link>
```

Examples:

- Back to homepage:
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:65-69`
- Create transaction:
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:75-79`
- Detail table link:
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:51-60`
- Edit action:
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:196-203`
- Dropdown detail action:
  - `~/Development/qcash-ui-issuance-lc/components/dropdown.tsx:129-136`

This means Next owns navigation while QUI owns the inner interaction control.

---

# 3. Remote exports and host entry resolution

## 3.1 Remote exports

The remote federation configuration exposes four route components:

- `~/Development/qcash-ui-issuance-lc/next.config.js:19-36`

Specifically:

- `./issuance-lc` → list page, line 31
- `./issuance-lc/detail` → dynamic detail page, line 32
- `./issuance-lc/create` → create page, line 33
- `./issuance-lc/edit` → dynamic edit page, line 34
- `./globals` → global CSS, line 35

The remote entry filename is:

- `static/chunks/remoteEntry.js`
- `~/Development/qcash-ui-issuance-lc/next.config.js:20-22`

The standalone deployment base path is:

- `/issuance-lc-v2`
- `~/Development/qcash-ui-issuance-lc/next.config.js:61-65`

Therefore the standalone remote entry is expected at:

```text
/issuance-lc-v2/_next/static/chunks/remoteEntry.js
```

## 3.2 Host registration does not use a custom entry

The host registry entry is:

- `~/Development/qcash-ui/constants/features/registry.ts:519-523`

It contains:

- name: `qcash-ui-issuance-lc`
- route prefix: `/import-lc/issuance-lc`
- no `customEntry`

The host resolves a remote without `customEntry` using:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:78-100`

The name is stripped from `qcash-ui-issuance-lc` to `issuance-lc` here:

- `~/Development/qcash-ui/services/federation/utils/name.ts:1-4`

Then the standard entry builder appends `-v2`:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:81-84`

Result:

```text
${baseFederation}/issuance-lc-v2/_next/static/chunks/remoteEntry.js
```

That currently matches the remote’s base path. It is nevertheless a naming convention coupling: renaming the service folder or deploying its entry elsewhere requires adding `customEntry`.

The host initializes every registry remote before rendering:

- collects all features: `~/Development/qcash-ui/services/federation/init.ts:37-45`
- initializes the runtime as `qcash-ui`: `:48-66`
- marks federation loaded: `:68`

## 3.3 Host proxy pages mirror the remote routes

Host wrappers:

- List:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:15-18`
- Create:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/create/index.tsx:15-18`
- Detail:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:15-18`
- Edit:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:15-18`

Each page is loaded client-side through `createRemoteDynamic`:

- remote CSS is loaded first:
  - `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:9-16`
- SSR is disabled:
  - `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:21-24`
- the remote component is rendered with host page props:
  - `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:97-100`

All wrappers also place the page under the remote header/footer layout:

- e.g. `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:20-27`

The host uses prefix route matching, so one registry route covers create/detail/edit subpaths:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:104-120`
- `~/Development/qcash-ui/services/federation/utils/registry.ts:169-187`

---

# 4. Concrete navigation bugs and coupling assumptions

## 4.1 Critical assumption: shared Next router and conflicting base paths

The host base path is:

- `/v2`
- `~/Development/qcash-ui/next.config.js:66-70`

The remote standalone base path is:

- `/issuance-lc-v2`
- `~/Development/qcash-ui-issuance-lc/next.config.js:61-65`

The remote hard-codes unprefixed application routes, for example:

- `/homepage`
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:65`
- `/import-lc/issuance-lc/create`
  - `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:76`
- detail:
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:53`
- edit:
  - `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:199`

This works in the host only if the remote’s `next/link` and `next/router` resolve to compatible host-side Next modules/router context so `/v2` is applied. If remote-local Next internals apply the remote’s build-time base path, links can point to:

```text
/issuance-lc-v2/import-lc/issuance-lc/...
```

instead of the host route:

```text
/v2/import-lc/issuance-lc/...
```

Conversely, standalone behavior relies on the remote’s Next runtime adding `/issuance-lc-v2`.

The host federation config does not explicitly list Next itself in `shared`; it relies on `@module-federation/nextjs-mf` handling Next internals:

- `~/Development/qcash-ui/next.config.js:8-63`
- `~/Development/qcash-ui-issuance-lc/next.config.js:19-58`

This is an architectural coupling even if it currently works.

## 4.2 Dynamic parameters are obtained from host Next router context, not props

The host wrappers do not pass `transactionEventId` explicitly. The remote reads it from `router.query`:

- detail:
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:32-50`
- edit:
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:26-27`

This works because the host happens to use identically named dynamic Pages Router files:

- `~/Development/qcash-ui/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx`

Concrete assumption: the host route must continue to use the exact parameter name `transactionEventId` and expose a Next router context. A TanStack route such as `$transactionEventId` would not populate `next/router.query`.

## 4.3 There is no injectable navigation API

None of the exposed page components accepts navigation functions or route parameters. They call `useRouter()` internally.

Examples:

- form: `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:85`
- detail: `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:83`
- page-level detail query: `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:32`

Consequences:

- The host cannot redirect those operations into TanStack Router.
- The host cannot centrally add or normalize route prefixes.
- The host cannot distinguish standalone versus embedded navigation except indirectly through whichever Next router wins the shared runtime.

## 4.4 Invalid/fragile nested interactive elements

Several links wrap QUI buttons. If `QuiButton` renders a `<button>`, the resulting `<a><button /></a>` is invalid nested interactive HTML.

References:

- `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:65-69`
- `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:75-79`
- `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:51-60`
- `~/Development/qcash-ui-issuance-lc/components/table/ongoing-columns.tsx:196-203`
- `~/Development/qcash-ui-issuance-lc/components/dropdown.tsx:129-136`

Potential effects include inconsistent keyboard activation, duplicate event paths, accessibility problems, and behavior differences when QUI or React Aria suppresses click defaults.

## 4.5 Dropdown detail navigation relies on a nested link rather than the menu action

The dropdown action handler recognizes:

- delete
- reject
- rework
- workflow
- download
- report

but not `detail`:

- `~/Development/qcash-ui-issuance-lc/components/dropdown.tsx:39-61`

The detail menu item has `id="detail"` but navigation is delegated to a nested Next link/button:

- `~/Development/qcash-ui-issuance-lc/components/dropdown.tsx:129-136`

Thus selecting the QUI menu item invokes `handleAction("detail")`, which only closes the menu and performs no route operation. Navigation succeeds only if the nested link’s browser click is processed before closing/unmounting. This is a concrete coupling bug between QUI menu event semantics and Next Link DOM behavior.

## 4.6 Direct browser redirects bypass the host router and host base path

The Axios interceptor performs hard browser navigation:

- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:71-85`
- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:124-128`

Specifically:

```ts
window.location.replace("/main-page");
window.location.href = "/main-page?logout=true";
```

These bypass Next’s base-path handling. Under the current host, canonical host routes are under `/v2`, so these target root-level legacy routes rather than `/v2/main-page` unless the reverse proxy separately aliases them.

This logic also compares `window.location.pathname` against unprefixed paths:

- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:71-75`

Under the host, `pathname` can be `/v2/homepage` or `/v2/main-page`, so `isDashboard` and `isMain` will be false.

There is also a definite query-string bug:

- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:73`

`window.location.pathname` never includes `?logout=true`, so:

```ts
pathname === "/main-page?logout=true"
```

cannot be true.

The same defects remain in the duplicate original interceptor:

- `~/Development/qcash-ui-issuance-lc/services/custom-axios-original.ts:65-79`
- `~/Development/qcash-ui-issuance-lc/services/custom-axios-original.ts:116-120`

## 4.7 Hard-coded product-shell destinations

The feature assumes the containing application has these exact routes:

- `/homepage`
- `/import-lc/issuance-lc`
- `/import-lc/issuance-lc/create`
- `/import-lc/issuance-lc/detail/:id`
- `/import-lc/issuance-lc/edit/:id`
- `/main-page`

There is no route configuration object or host-provided route contract. Examples include:

- `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:65-76`
- `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:1244`
- `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:1396`
- `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:491`
- `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:650`
- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:85`
- `~/Development/qcash-ui-issuance-lc/services/custom-axios.ts:127`

The current host happens to mirror the issuance paths, but this remains shell-specific coupling.

## 4.8 Mixed QUI event conventions increase interaction fragility

The same `QuiButton` abstraction is used with both `onPress` and `onClick`:

- `onPress`:
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:491`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:650`
- `onClick`:
  - `~/Development/qcash-ui-issuance-lc/containers/form-new-skin.tsx:1395`
  - `~/Development/qcash-ui-issuance-lc/containers/detail/view-new-skin.tsx:580`

This is not necessarily broken if QUI supports both, but it makes navigation behavior dependent on the exact design-system Button implementation/version loaded by the host.

---

# 5. Standalone provider architecture versus embedded architecture

## 5.1 Remote standalone mode

The remote’s `_app.tsx` statically imports its own auth provider:

- `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:7`

It mounts:

- `AuthProvider`
- a dynamically loaded QUI provider
- `GlobalStoreProvider`

at:

- `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:99-105`

The installed remote auth-provider version is `0.9.119-mfe`:

- `~/Development/qcash-ui-issuance-lc/package.json:22`
- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/package.json:1-5`

That old installed version directly imports and calls `next/router` internally:

- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/dist/src/auth.js:73-75`
- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/dist/src/auth.js:173-176`

## 5.2 Embedded mode in qcash-ui

The exposed page modules do not include remote `_app.tsx`. Therefore, when embedded:

- the remote’s `AuthProvider` is not mounted;
- the remote’s `StandaloneQuiProvider` is not mounted;
- the host providers wrap the remote instead.

The host statically imports `BriProviders`:

- `~/Development/qcash-ui/pages/_app.tsx:19`

It mounts it before QUI and before the remote page:

- `~/Development/qcash-ui/pages/_app.tsx:160-193`

`BriProviders` statically imports both auth packages and supplies the host Next router:

- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:1-12`
- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:18-33`

The host also statically shares the Ordent auth provider as a singleton:

- `~/Development/qcash-ui/next.config.js:15-22`

The remote declares the same singleton:

- `~/Development/qcash-ui-issuance-lc/next.config.js:41-45`

The host package version and remote package version differ:

- remote installed: `0.9.119-mfe`
  - `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/package.json:3`
- host installed: `0.9.138-mfe`
  - `~/Development/qcash-ui/node_modules/@ordentco/addons-auth-provider/package.json:3`

Both federation configs use `requiredVersion: false`, so whichever singleton is established is accepted despite the version mismatch.

---

# 6. Could auth-provider lazy/static loading fix the navigation failures?

## Short answer

**Not the navigation failures themselves.**

## What static auth loading could plausibly fix

Static host loading can help if the symptom is one of these:

- `useAuth()` receives a default/guest context because the remote resolved a different package instance.
- The remote and host created separate auth contexts.
- The older remote auth package won the federation singleton race.
- An auth provider was mounted after the remote consumer.
- Auth `guard()` or authority hydration runs against the wrong package version.

That is plausible because:

- host and remote carry different auth-provider versions;
- both declare the package a singleton with version checks disabled;
- the remote’s old implementation internally couples to `next/router`;
- the remote pages immediately call `useAuth()` and `guard()`:
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:15-28`
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx:25-27`
  - `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx:68-71`

## Why it is unlikely to be the current fix

The current host already:

1. statically imports both auth packages through `BriProviders`;
2. mounts those providers above the remote;
3. supplies the host Next router to both providers;
4. loads the issuance component client-side afterward.

References:

- static provider import:
  - `~/Development/qcash-ui/pages/_app.tsx:19`
- provider mounting:
  - `~/Development/qcash-ui/pages/_app.tsx:166-193`
- static package imports and injected router:
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:1-12`
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:18-33`
- remote is dynamically loaded later:
  - `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:41-76`

Changing the host providers to lazy loading would generally make provider/share-scope ordering less deterministic unless the entire application waits for them before rendering children.

## What auth loading cannot fix

Auth loading cannot repair:

- remote `next/link` using the wrong build/base path;
- `router.query` depending on a host Next dynamic route;
- absence of a TanStack Router adapter;
- hard-coded `/homepage`, `/main-page`, and issuance paths;
- root-relative `window.location` redirects;
- `pathname` comparisons that omit `/v2`;
- the impossible `pathname === "/main-page?logout=true"` test;
- nested QUI button/Next link markup;
- the QUI dropdown’s missing `"detail"` action.

Those operations do not flow through the auth provider’s router prop. They call Next Router or browser location APIs directly.

---

# Final assessment

The current mount succeeds because `qcash-ui` and the issuance remote are both Pages Router applications and the host deliberately mirrors the remote route structure. It is not router-neutral.

The most important risks are:

1. **Conflicting host and remote base paths coupled through shared Next internals.**
2. **Dynamic route parameters read from host `next/router.query` rather than passed as an MFE contract.**
3. **Hard browser redirects that ignore `/v2`, including a definite `pathname`/query bug.**
4. **QUI interactive components nested inside Next links and a dropdown detail action that has no action-handler navigation.**
5. **No TanStack Router or host navigation adapter at all.**

Static auth-provider loading can stabilize auth singleton/context behavior, but the host already does that. It is not a plausible general fix for these navigation issues.

---

