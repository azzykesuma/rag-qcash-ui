# Trace host routing (@explore subagent)

- **ID**: `ses_02620099bffeZjI7DtkQ7yCCCx`
- **Source Tool**: `opencode`
- **Date**: `2026-08-07 08:39:21`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Analyze the qcash-ui repository routing and module federation integration, specifically Issuance LC. Find all relevant route declarations, lazy/remote loading wrappers, navigation handling, and lifecycle code. Compare Issuance LC integration with at least one working page or remote. Do not edit files. Return likely causes for URL changing without rendered page changing, with exact file paths and line references. Thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

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
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

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
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
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

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
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

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 16: Assistant

# Repository analysis: Issuance LC routing and Module Federation

No files were edited.

## 1. Issuance LC route declarations

This is a Next.js Pages Router application. The configured base path is `/v2`:

- `D:\work\qcash-ui\next.config.js:66-70`

Therefore, the actual browser routes are:

| Browser URL | Next page | Federated expose |
|---|---|---|
| `/v2/import-lc/issuance-lc` | `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:15-18` | `qcash-ui-issuance-lc/issuance-lc` |
| `/v2/import-lc/issuance-lc/create` | `D:\work\qcash-ui\pages\import-lc\issuance-lc\create\index.tsx:15-18` | `qcash-ui-issuance-lc/issuance-lc/create` |
| `/v2/import-lc/issuance-lc/edit/:transactionEventId` | `D:\work\qcash-ui\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:15-18` | `qcash-ui-issuance-lc/issuance-lc/edit` |
| `/v2/import-lc/issuance-lc/detail/:transactionEventId` | `D:\work\qcash-ui\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:15-18` | `qcash-ui-issuance-lc/issuance-lc/detail` |

Important route-contract details:

- Create has **no ID segment**. `/create/:id` is not declared.
- Edit and detail require an ID path segment.
- There is no Issuance LC cancellation route.
- Query-only URLs such as `/issuance-lc?id=123`, `/issuance-lc?mode=detail`, or `/detail?id=123` do not select the dynamic detail page.
- `router.pathname` for a detail page is the template `/import-lc/issuance-lc/detail/[transactionEventId]`, not the concrete URL.

The Module Federation feature registry declares one prefix route:

- `D:\work\qcash-ui\constants\features\registry.ts:516-520`

```ts
"qcash-ui-issuance-lc": {
  toggleKey: "toggleIssuanceLc",
  routes: ["/import-lc/issuance-lc"],
}
```

Prefix matching is the default and correctly includes all declared children:

- Prefix implementation: `D:\work\qcash-ui\services\federation\utils\registry.ts:102-120`
- Feature filtering/loading: `D:\work\qcash-ui\services\federation\utils\registry.ts:160-196`

The toggle is enabled in every configured environment:

- `D:\work\qcash-ui\config.ts:507-513`

The pages independently gate rendering with the same toggle:

- Index: `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:21-27`
- Create: `D:\work\qcash-ui\pages\import-lc\issuance-lc\create\index.tsx:21-27`
- Edit: `D:\work\qcash-ui\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:21-27`
- Detail: `D:\work\qcash-ui\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:21-27`

There are no local Issuance LC `router.push`, `router.replace`, `<Link>`, or click handlers. Internal list/create/edit/detail navigation belongs to the remote `qcash-ui-issuance-lc`, whose source is not present in this repository.

---

## 2. Remote entry resolution

Issuance LC has no `customEntry`:

- `D:\work\qcash-ui\constants\features\registry.ts:517-520`

Its URL is synthesized from the remote name:

- `D:\work\qcash-ui\services\federation\utils\registry.ts:81-84`
- `D:\work\qcash-ui\services\federation\utils\registry.ts:90-100`
- Name stripping: `D:\work\qcash-ui\services\federation\utils\name.ts:1-4`

That produces:

```text
{baseFederation}/issuance-lc-v2/_next/static/chunks/remoteEntry.js
```

`baseFederation` is the selected API origin with `/api` removed:

- `D:\work\qcash-ui\config.ts:16-26`
- Environment selection: `D:\work\qcash-ui\utils\index.ts:23-58`
- URL selection: `D:\work\qcash-ui\utils\index.ts:65-109`

All remotes are registered manually at runtime:

- `D:\work\qcash-ui\services\federation\init.ts:32-46`
- Runtime `init`: `D:\work\qcash-ui\services\federation\init.ts:48-66`
- Federation-ready flag: `D:\work\qcash-ui\services\federation\init.ts:68-69`

The shell also installs `NextFederationPlugin`, but its webpack `remotes` are empty:

- `D:\work\qcash-ui\next.config.js:8-13`
- Plugin installation: `D:\work\qcash-ui\next.config.js:181`

The runtime registry is therefore essential.

---

## 3. Lazy-loading and lifecycle

All four Issuance pages use the custom `createRemoteDynamic` wrapper.

### Loading sequence

For a remote such as `qcash-ui-issuance-lc/issuance-lc/detail`:

1. Extract module name `qcash-ui-issuance-lc`.
2. Load `qcash-ui-issuance-lc/globals`.
3. Create a client-only `next/dynamic` component.
4. Load the actual expose.
5. Render it with route/page props.

Relevant code:

- Module extraction and globals load:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:9-12`
- Actual `loadRemote`:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:13-25`
- State initialization:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:31-39`
- Load and state transition:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:41-65`
- Mount-only effect and abort cleanup:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:68-76`
- Loading/error/component rendering:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:78-102`

The key lifecycle property is:

```ts
useEffect(() => {
  ...
}, []);
```

The wrapper loads its expose only once per wrapper mount.

The outer loading component remains until `/globals` completes. After that, `next/dynamic` has a second loading phase while the actual expose loads.

### Globals loading

`dynamicLoadStyles` actually invokes:

```ts
loadRemote(`${moduleName}/globals`)
```

- `D:\work\qcash-ui\services\federation\utils\loader.ts:19-37`

It has:

- Default 15-second timeout:  
  `D:\work\qcash-ui\services\federation\types.ts:134-149`
- One retry after a fixed five-second delay:  
  `D:\work\qcash-ui\services\federation\utils\loader.ts:48-75`

A globals failure can therefore make a page appear stuck for roughly 35 seconds before final failure.

### Page and layout lifecycle

Every Issuance page also loads the header/footer as a remote:

- `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:10-13`
- Equivalent declarations exist at lines `10-13` in create, edit, and detail.

Each passes `isNewUI={true}`:

- Index: `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:21-27`
- Create: `D:\work\qcash-ui\pages\import-lc\issuance-lc\create\index.tsx:21-27`
- Edit: `D:\work\qcash-ui\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:21-27`
- Detail: `D:\work\qcash-ui\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:21-27`

`_app` selects the current page's `getLayout` and renders the page without a route key:

- Layout selection: `D:\work\qcash-ui\pages\_app.tsx:151`
- Render: `D:\work\qcash-ui\pages\_app.tsx:187-191`

```tsx
<Component {...pageProps} />
```

There is no `key={router.asPath}` or equivalent to force a remount.

---

## 4. Route-aware federation lifecycle

`ModuleFederationProvider` observes `router.pathname`, not the concrete `router.asPath`:

- Router acquisition: `D:\work\qcash-ui\components\federation\provider.tsx:68-75`
- Active feature update: `D:\work\qcash-ui\components\federation\provider.tsx:114-125`
- Route module load: `D:\work\qcash-ui\components\federation\provider.tsx:143-158`

It suppresses duplicate work with:

```ts
if (modulesLoadedForPath === router.pathname) return;
```

Consequences:

- Query-only changes do not reload the route module.
- Changing only `transactionEventId` retains the same `router.pathname`.
- `/detail/1` to `/detail/2` does not rerun the provider's loading lifecycle.
- The same remote React component instance may be retained unless the remote reacts correctly to `router.query`/`router.asPath`.

The page wrapper independently loads the expose, so provider preloading does not directly gate child rendering. The provider renders `children` regardless of module status:

- `D:\work\qcash-ui\components\federation\provider.tsx:171-178`

---

## 5. Navigation and surrounding route handlers

### Header/menu navigation

The menu itself is inside `qcash-ui-header-footer/default`, not this repository. The shell only mounts it around the page.

Menu and authority data are synchronized for remotes:

- Menu data extraction: `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:224-241`
- Route-dependent auth hydration: `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:256-300`
- Bridge refresh on `router.asPath`: `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:383-452`

The Next router is passed into both auth providers:

- `D:\work\qcash-ui\components\providers\BRIProviders.tsx:43-50`

No host-side handler translates raw `history.pushState` calls into Next navigation.

### Other route effects

These can affect what appears after a route change:

- Onboarding redirect using `router.replace`:  
  `D:\work\qcash-ui\components\providers\OnboardingTourProvider.tsx:11-43`
- Session-expiration route listener:  
  `D:\work\qcash-ui\hooks\use-session-expiration.ts:78-106`
- Saved auth redirect:  
  `D:\work\qcash-ui\pages\_app.tsx:116-149`
- Component hidden while MFA modal is open:  
  `D:\work\qcash-ui\pages\_app.tsx:174-190`

---

## 6. Comparison with working integrations

### Amend Cancel LC: closest structural comparator

Amend Cancel LC uses the same wrapper, same header/footer layout, separate host pages, and matching exposes:

- Index: `D:\work\qcash-ui\pages\import-lc\amend-cancel-lc\index.tsx:10-26`
- Create: `D:\work\qcash-ui\pages\import-lc\amend-cancel-lc\create\[id].tsx:10-26`
- Edit: `D:\work\qcash-ui\pages\import-lc\amend-cancel-lc\edit\[id].tsx:10-26`
- Detail: `D:\work\qcash-ui\pages\import-lc\amend-cancel-lc\detail\[id].tsx:9-25`
- Cancellation: `D:\work\qcash-ui\pages\import-lc\amend-cancel-lc\cancellation\[id].tsx:10-26`

Differences from Issuance LC:

1. **Explicit remote entry**

   Amend:

   - `D:\work\qcash-ui\constants\features\registry.ts:615-620`

   ```ts
   customEntry: remoteEntry("amend-cancel-lc-v2")
   ```

   Issuance relies on inferred naming.

2. **Create route contract**

   - Amend: `/create/:id`
   - Issuance: `/create`

3. **Layout mode**

   - Amend legacy pages use `isNewUI={false}`.
   - Issuance uses `isNewUI={true}`.

4. **Cancellation**

   - Amend has a host route and expose.
   - Issuance does not.

Otherwise, the patterns are nearly identical. If Amend works but Issuance does not, the common shell wrapper is less likely to be the sole cause. Entry URL, expose names, remote routing, and `isNewUI` handling become stronger suspects.

### Incoming Document: simpler loader comparator

Incoming uses direct `next/dynamic(() => loadRemote(...))` rather than `createRemoteDynamic`:

- `D:\work\qcash-ui\pages\import-lc\incoming\index.tsx:9-19`
- `D:\work\qcash-ui\pages\import-lc\incoming\detail\[docDetailEventId].tsx:9-19`
- `D:\work\qcash-ui\pages\import-lc\incoming\create\[docEventId].tsx:9-19`

It does not block component creation on a `/globals` load and has only one dynamic-loading layer. If Incoming works while Issuance remains loading, `createRemoteDynamic` or Issuance's `/globals` expose is a primary differentiator.

---

# Likely causes of “URL changes, rendered page does not”

## 1. Remote navigation changes browser history without using the shell's Next router

**Likelihood: high; strongest fit for the exact symptom.**

There is no local Issuance navigation implementation and no host bridge for raw `history.pushState`. If the remote or header uses:

```ts
window.history.pushState(...)
```

the address bar changes, but:

- Next does not switch `Component`.
- `router.pathname` does not update.
- Federation route effects do not run.
- The currently mounted page remains visible.

Evidence that host lifecycle depends on Next router state:

- `D:\work\qcash-ui\pages\_app.tsx:79-95`
- `D:\work\qcash-ui\components\federation\provider.tsx:114-125`
- `D:\work\qcash-ui\components\federation\provider.tsx:143-158`

The source of `qcash-ui-issuance-lc` and `qcash-ui-header-footer` should be checked for `pushState`, raw anchors, custom history, or a separate router instance.

## 2. Navigation URL does not match the host route contract

**Likelihood: high.**

Examples:

- Navigating to `/issuance-lc/create/:id` although only `/create` exists.
- Navigating to `/detail?id=123` instead of `/detail/123`.
- Navigating by query parameters while staying on the index route.
- Missing or incorrectly duplicating `/v2`.
- Using a route name not declared by the host.

The exact accepted routes are listed in section 1. Query-only changes leave `router.pathname` unchanged, and the provider suppresses reloading based on pathname:

- `D:\work\qcash-ui\components\federation\provider.tsx:143-158`

## 3. Same dynamic page is reused when only the transaction ID changes

**Likelihood: high for `/detail/1` → `/detail/2` or `/edit/1` → `/edit/2`.**

Next preserves page component state for navigation between instances of the same dynamic route. The shell does not provide a route key:

- `D:\work\qcash-ui\pages\_app.tsx:187-190`

The remote wrapper loads once:

- `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:68-76`

The provider also sees the same pathname template:

- `D:\work\qcash-ui\components\federation\provider.tsx:145-158`

If the remote reads `transactionEventId` only on mount, or uses an effect without the ID in its dependencies, the URL updates while old transaction content remains.

## 4. Issuance remote is using a different Next/router or React singleton

**Likelihood: medium-high.**

The remote must observe the shell's Next router context. The host shared configuration does not explicitly list `react`, `react-dom`, `next`, or `next/router`:

- `D:\work\qcash-ui\next.config.js:15-63`

`nextjs-mf` normally injects Next defaults, but a differently built Issuance remote can still bundle its own router or incompatible shared versions. That would explain why host navigation or `useRouter()` updates work in another remote but not Issuance.

Host versions are:

- Next `13.5.11`
- React/React DOM `18.3.1`
- Federation runtime `0.18.1`

See `D:\work\qcash-ui\package.json:19-50`.

## 5. Remote entry or expose contract mismatch

**Likelihood: medium.**

Issuance assumes:

```text
/issuance-lc-v2/_next/static/chunks/remoteEntry.js
```

and requires these exact exposes:

```text
./globals
./issuance-lc
./issuance-lc/create
./issuance-lc/edit
./issuance-lc/detail
```

Host references:

- `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:15-18`
- `D:\work\qcash-ui\pages\import-lc\issuance-lc\create\index.tsx:15-18`
- `D:\work\qcash-ui\pages\import-lc\issuance-lc\edit\[transactionEventId].tsx:15-18`
- `D:\work\qcash-ui\pages\import-lc\issuance-lc\detail\[transactionEventId].tsx:15-18`
- `/globals`: `D:\work\qcash-ui\services\federation\utils\loader.ts:31`

Unlike Amend, Issuance has no explicit `customEntry`. A deployment-folder discrepancy would affect only this remote.

## 6. Stale `remoteEntry.js` or cached remote implementation

**Likelihood: medium.**

Remote entry URLs are static and have no cache-busting query:

- `D:\work\qcash-ui\constants\features\registry.ts:30-32`
- `D:\work\qcash-ui\services\federation\utils\registry.ts:81-84`

A timestamp-stripping plugin exists but is disabled:

- `D:\work\qcash-ui\services\federation\init.ts:17-30`
- `D:\work\qcash-ui\services\federation\init.ts:63`

A stale host/remote combination could retain old expose mappings or old navigation behavior. Repository history also shows that edit/detail expose paths were previously swapped; current source is correct, but an older deployed host chunk could still contain that mapping.

## 7. `/globals` loading prevents the actual page from loading

**Likelihood: medium if a skeleton is visible; lower if the old page remains completely unchanged.**

The wrapper will not even construct the page dynamic component until `/globals` succeeds:

- `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:9-12`
- `D:\work\qcash-ui\components\federation\__tests__\remote-dynamic.test.tsx:543-562`

With timeout plus retry, the delay can be substantial. Incoming does not have this dependency.

## 8. `isNewUI={true}` behavior inside header/footer

**Likelihood: medium if the problem began after the new-layout migration.**

All Issuance pages now pass `isNewUI={true}`. Legacy Amend pages pass false. Because the actual header/footer implementation is remote, this repository cannot verify whether its new-UI branch:

- renders updated `children`,
- memoizes stale children,
- owns a nested router,
- or uses different navigation handling.

Relevant Issuance lines are `21-27` in each page file.

## 9. Global error boundary retains an earlier error across route changes

**Likelihood: lower, but concrete lifecycle defect.**

`RemoteErrorBoundary` is above the page tree:

- `D:\work\qcash-ui\components\federation\provider.tsx:171-178`

Once it enters `hasError`, it does not automatically reset on route changes:

- Error state: `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:22-29`
- Error render: `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:64-103`
- Reset only happens through retry: `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:47-60`

Thus, URL changes can occur while the same global error UI remains.

---

## Bottom line

The host has valid, distinct Issuance LC pages and correct current expose mappings. A correct Next navigation from index to create/detail/edit should replace or at least show a loading state; it should not leave the old page untouched.

The highest-probability explanations are:

1. The Issuance remote updates browser history directly or uses a separate router instance.
2. The target URL does not match the host's exact file-system route contract.
3. The transition changes only query/ID while the same dynamic page and remote component remain mounted.
4. Issuance's remote entry/exposes or shared router dependencies differ from working remotes.
5. The `isNewUI` header/footer branch retains or fails to update its children.

---

