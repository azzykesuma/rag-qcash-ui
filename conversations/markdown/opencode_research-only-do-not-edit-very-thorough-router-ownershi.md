# Audit router provider coupling (@explore subagent)

- **ID**: `ses_016acd7d8ffeb110gAFD5tsD7l`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 08:39:28`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Very thorough router ownership/coupling audit focused on qcash-ui, qcash-ui-header-footer, qcash-ui-issuance-lc, addons-auth-provider, and relevant Module Federation configs. Determine whether multiple Next router module instances/providers exist, whether next/router is shared or bundled, whether remote _app/provider runs hosted, whether route changes are globally patched/intercepted, and whether issuance-LC uniquely adds providers/listeners/history handling that can stale subsequent header navigation. Compare a similar working modal-selected route if identifiable. Return exact file:line evidence, root cause ranking, and concrete browser verification probes.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

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

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

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

## Turn 11: Assistant

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

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 22: Assistant

# Router ownership/coupling audit

## Executive conclusion

1. **In the normal Webpack Module Federation path, there should be one effective hosted `next/router` singleton: the `qcash-ui` host router.** `nextjs-mf` automatically shares `next/router` as a singleton even though the application configs do not list it explicitly.
2. **The issuance-LC remote `_app` does not execute when its exposed pages are hosted by `qcash-ui`.** Therefore its standalone `AuthProvider`, `GlobalStoreProvider`, `I18nextProvider`, and Next router provider are not mounted in the hosted tree.
3. **There are multiple Next installations on disk**, most importantly issuance-LC’s Next 13.5.6 plus a nested Next 12.3.4 pulled by its old `@ordentco/addons-auth-provider@0.9.119-mfe`. Under the normal `NextFederationPlugin` path, `next/router` should nevertheless resolve through the singleton share scope.
4. **Issuance-LC is explicitly singled out for hard navigation when leaving it.** The header-footer code says a document navigation is required to guarantee that the federated issuance/QUI page unmounts. This is the strongest repository evidence for the reported “subsequent header navigation is stale” behavior.
5. **Issuance-LC does not patch browser history or install popstate/beforePopState handlers.** It does have several leaked `scroll` listeners added during render, plus additional QUI and QueryClient providers, but those do not directly own routing.
6. **The current addons-auth-provider repository contains a global router/location interception facility**, but it is activated only by `createQcashAxios()`/`attachAuthInterceptor()`. None of the three focused applications directly calls that API. Another previously loaded remote can activate it globally, however, and the patch then persists for the lifetime of the document.
7. **The issuance Rspack branch is unsafe:** it uses a raw federation plugin with no automatic React/Next sharing. It appears dormant in this checkout because `next-rspack` is neither installed nor locked, but if enabled in another build environment it can produce separate React/router module instances.

---

# 1. Router ownership and module-instance findings

## Hosted router owner

The host is a Pages Router Next application:

- `~/Development/qcash-ui/pages/_app.tsx:23-24` imports `useRouter` from `next/router`.
- `~/Development/qcash-ui/pages/_app.tsx:79-80` obtains the host router.
- `~/Development/qcash-ui/package.json:40-43` uses Next `13.5.11`, React/React DOM `18.3.1`.
- `~/Development/qcash-ui/next.config.js:69` gives the host a `/v2` base path.

There is no custom `RouterContext.Provider` or second router provider in `qcash-ui`. Next’s own `_app` wrapper supplies the router context.

## Remote pages are components, not complete Next applications

The host loads the issuance page modules directly:

- `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:15-18`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/create/index.tsx:15-18`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:15-18`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:15-18`

The generic loader calls `loadRemote(remote)` and renders the returned component:

- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:13-24`
- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:97-100`

Issuance exposes those page files, not `_app`:

- `~/Development/qcash-ui-issuance-lc/next.config.js:30-35`

Therefore, hosted rendering does **not** run:

- `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:63-67`, which would otherwise add `I18nextProvider`, `AuthProvider`, and `GlobalStoreProvider`.

The same is true for the hosted header-footer remote: the host loads `qcash-ui-header-footer/default`, exposed from the layout file:

- `~/Development/qcash-ui-header-footer/module-federation.config.mjs:9-13`
- `~/Development/qcash-ui-header-footer/layouts/default/index.tsx:4-8`

Its standalone `_app` does not wrap the hosted layout. That `_app` currently has no active auth provider anyway:

- `~/Development/qcash-ui-header-footer/pages/_app.tsx:62-66`

## Hosted auth providers

Although there is only one hosted Next router provider, the host deliberately runs **two auth context providers**:

- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:7-18` defines the BRI provider wrapper.
- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:20-37` defines the Ordent provider wrapper.
- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:43-51` nests both and passes the same host `useRouter()` result to each.
- `~/Development/qcash-ui/pages/_app.tsx:169-171` mounts those providers around all hosted remotes.

The source comments explicitly describe two concurrent auth contexts:

- `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:6-20`
- `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:22-35`

This is provider duplication, but **not router-provider duplication**. Both current host auth packages receive the host router by prop.

## Issuance’s old auth package does own a router when standalone

Issuance-LC uses old Ordent auth `0.9.119-mfe`:

- `~/Development/qcash-ui-issuance-lc/package.json:22`
- `~/Development/qcash-ui-issuance-lc/yarn.lock:1354-1369`

That package directly imports and calls its own `next/router`:

- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/src/auth.tsx:174-181`

It also subscribes to route changes:

- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/src/auth.tsx:769-791`

But this provider executes only in issuance’s standalone `_app`. When hosted, the exposed page calls `useAuth()` and consumes the host’s shared Ordent context instead:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:1-17`
- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:1-33`

That context crossing depends on the auth package being a Module Federation singleton.

---

# 2. Is `next/router` shared or bundled?

## Normal Webpack path: automatically shared singleton

None of the application `shared` blocks explicitly names `next/router`:

- Host: `~/Development/qcash-ui/next.config.js:15-63`
- Header-footer: `~/Development/qcash-ui-header-footer/module-federation.config.mjs:17-37`
- Issuance-LC: `~/Development/qcash-ui-issuance-lc/next.config.js:41-58`

Nevertheless, `@module-federation/nextjs-mf` supplies a default share:

- `~/Development/qcash-ui/node_modules/@module-federation/nextjs-mf/dist/src/internal.js:30-34` declares `next/router` with `singleton: true`, `requiredVersion: false`.
- `~/Development/qcash-ui/node_modules/@module-federation/nextjs-mf/dist/src/plugins/NextFederationPlugin/index.js:106-135` merges default shares before application shares.
- `~/Development/qcash-ui/node_modules/@module-federation/nextjs-mf/dist/src/plugins/NextFederationPlugin/index.js:119-124` also automatically exposes `./next/router`.

The same default code exists in the installed issuance and header-footer `nextjs-mf` packages.

Thus, for builds taking `new NextFederationPlugin(...)`, hosted header and issuance router imports should consume one singleton—normally the host-provided Next 13 router module.

## Multiple physical Next installations exist

Issuance itself uses Next 13.5.6:

- `~/Development/qcash-ui-issuance-lc/package.json:34`

Its old auth package declares Next 12 as a regular dependency:

- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/package.json:19-35`
- `~/Development/qcash-ui-issuance-lc/yarn.lock:1365-1374`

That resolves to nested Next 12.3.4:

- `~/Development/qcash-ui-issuance-lc/node_modules/@ordentco/addons-auth-provider/node_modules/next/package.json:1-4`
- `~/Development/qcash-ui-issuance-lc/yarn.lock:5634-5644`

So there are definitely multiple module copies available to the bundler. The singleton federation default is what prevents them from becoming independently active hosted router modules.

## Rspack branch can defeat this protection

Issuance conditionally chooses a raw plugin:

- `~/Development/qcash-ui-issuance-lc/next.config.js:78-88`

The `ModuleFederationPluginV1` branch receives only `federationConfig`, whose shared list excludes:

- `react`
- `react-dom`
- `next/router`
- Next internals

Therefore, if `hasRspackMF` becomes true, the automatic `NextFederationPlugin` defaults shown above are bypassed. Separate React and router module instances are then plausible.

In this checkout, the branch appears inactive:

- `~/Development/qcash-ui-issuance-lc/next.config.js:5-10` falls back when `next-rspack` cannot be required.
- `next-rspack` is absent from `~/Development/qcash-ui-issuance-lc/package.json` and `yarn.lock`.

This should still be verified against the actual deployed build image.

---

# 3. Issuance-specific lifecycle/providers/listeners

## Explicit hard-navigation treatment

This is the strongest finding.

Header-footer singles out issuance as the only registered hard-navigation page:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:37-39`

When the current path starts with issuance, subsequent header navigation uses `window.location.href` instead of `router.push`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:90-100`

The comment states why:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:96-97` says document navigation guarantees that the federated issuance page is unmounted when using new QUI components.

That is direct evidence that soft-navigation out of issuance has already been recognized as unsafe or stale.

## Additional hosted providers

Issuance list and detail add a TanStack Query provider:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:38-50`
- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:179-198`

The query client is module-global and is also published to `window`:

- `~/Development/qcash-ui-issuance-lc/services/query-client.ts:3-11`
- `~/Development/qcash-ui-issuance-lc/services/query-client.ts:19-22`

Create dynamically loads and mounts the remote QUI providers:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx:112-162`
- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx:193-208`

These providers do not directly manipulate routing, but the hard-navigation comment specifically associates issuance’s new QUI usage with the need to guarantee unmounting.

## Listener leaks

Several issuance form implementations add `scroll` listeners during render without cleanup:

- `~/Development/qcash-ui-issuance-lc/components/create/page-3.tsx:441-443`
- `~/Development/qcash-ui-issuance-lc/components/create-new-skin/page-3.tsx:1011-1013`
- `~/Development/qcash-ui-issuance-lc/components/create-new-old/page-3.tsx:1003-1005`

Because these statements are in render scope, rerenders can install the same or newly closed-over callbacks repeatedly. They can survive if the remote is not properly unmounted and are concrete evidence of lifecycle residue. They do not intercept router calls, but they reinforce the hard-unmount rationale.

Other issuance listeners inspected are cleaned up correctly, for example:

- `~/Development/qcash-ui-issuance-lc/components/ui/qui-currency-amount.tsx:123-136`

## No issuance history ownership

The focused production issuance code has no:

- `window.history.pushState`
- `window.history.replaceState`
- `beforePopState`
- `popstate`
- direct `Router.events` patch
- global click interception

The only route behavior found is ordinary `useRouter` calls and page-local `push`, `replace`, `reload`, or `back`.

---

# 4. Global route interception audit

## Current addons-auth-provider has a process-wide patch, conditionally activated

`attachAuthInterceptor()` applies browser-global patches once:

- `~/Development/addons-auth-provider/src/global-custom-axios/interceptors/auth-interceptors.ts:302-317`

It patches:

- `Location.prototype.replace`: lines `319-334`
- `Location.prototype.assign`: lines `336-349`
- `Location.prototype.href`: lines `351-365`
- capturing `showDialog` listener: lines `371-408`
- `next/router` singleton `Router.push`: lines `410-430`
- `window.next.router.push`: lines `432-456`

A blocked router call returns `Promise.resolve(true)` without navigating:

- `~/Development/addons-auth-provider/src/global-custom-axios/interceptors/auth-interceptors.ts:415-425`
- `~/Development/addons-auth-provider/src/global-custom-axios/interceptors/auth-interceptors.ts:437-448`

This can exactly look like a successful but stale header click.

### Scope of blocking

It only suppresses destinations containing:

- `landing-page`
- `dashboard`
- `/`
- empty URL

and only while:

- `window.__reloginActive` is true, or
- `Date.now() < window.__blockRedirectUntil`

See lines `416-420` and `438-442`.

Therefore it cannot explain every arbitrary stale header route. It is highly relevant if the failing header target is homepage/root/dashboard, especially around session-expired re-login.

## Activation is not automatic from `AuthProvider`

The global patch is installed by `createQcashAxios()`:

- `~/Development/addons-auth-provider/src/global-custom-axios/create-axios.ts:87-109`

The provider’s legacy internal Axios service is separate:

- `~/Development/addons-auth-provider/services/custom-axios.ts:7-13`
- `~/Development/addons-auth-provider/services/custom-axios.ts:35-56`

No direct `createQcashAxios()` or `attachAuthInterceptor()` call was found in:

- `qcash-ui`
- `qcash-ui-header-footer`
- `qcash-ui-issuance-lc`

However, any other remote loaded earlier can apply the patch to the same window. Examples elsewhere in the workspace use this API, so runtime verification is necessary.

The patch does **not** override `history.pushState` or `history.replaceState`.

---

# 5. Similar working modal-selected route

The closest identifiable comparison is the LC Export modal.

## Working soft-navigation path

LC Export:

- Selects a modal method.
- Closes/reset the modal.
- Awaits `router.push()`.

Evidence:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcExportModal.tsx:37-54`

It is not registered for hard navigation.

## Issuance modal path

LC Import constructs the issuance URL and normally calls `router.push`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:97-113`

But once on issuance, generic header navigation is diverted to a full document navigation by `REGISTERED_HARD_NAV_PAGES`.

This comparison strongly suggests the modal itself is not the primary problem. The differentiator is the lifecycle of the selected issuance remote after it has mounted—particularly the QUI provider/component stack and/or federation-instance divergence—not modal state management.

---

# Root-cause ranking

## 1. High confidence: issuance soft-unmount/lifecycle residue

**Evidence:** issuance is the only hard-navigation page, and the code explicitly says a document navigation is needed to guarantee unmounting.

Likely manifestations:

- stale federated component/provider remains mounted,
- leaked form listeners remain,
- subsequent header navigation uses stale closures/context,
- soft route transition starts but old remote UI or state persists.

This is the repository’s own current diagnosis/workaround.

## 2. Medium-high risk: issuance federation build can produce duplicate React/router modules

Normal Webpack builds are protected by `nextjs-mf` singleton defaults. The raw Rspack branch is not.

Risk factors:

- raw `ModuleFederationPluginV1`,
- no `react`, `react-dom`, or `next/router` shares in issuance config,
- issuance Next 13 plus nested Next 12 from its old auth package,
- auth singleton versions are relaxed with `requiredVersion: false`.

If deployment takes this branch, multiple module instances and broken context/router identity become a leading root cause.

## 3. Medium: globally patched router suppresses successful-looking navigation

If another remote has called `createQcashAxios()`, root/dashboard header clicks can be swallowed while re-login state is active. The patch returns a successful promise, making the UI appear stale without an error.

This is global but **not issuance-specific**. Issuance could merely be the route where session refresh/re-login occurs.

## 4. Medium-low: issuance scroll-listener leaks amplify stale-page behavior

The render-time listeners are defects and can accumulate. They do not own navigation, so they are more likely an amplifier or proof of incomplete cleanup than the sole cause.

## 5. Low: multiple hosted Next router providers

No evidence of multiple hosted Next `RouterContext.Provider`s was found. Remote `_app` wrappers do not run when modules are hosted. This becomes plausible only if the deployed federation build bundles an independent React/router pair, particularly through the Rspack branch.

---

# Concrete browser verification probes

Run these before entering issuance, while on issuance, and after attempting header navigation.

## A. Detect the global auth patch

```js
({
  applied: window.__qcashAuthPatchesApplied,
  reloginActive: window.__reloginActive,
  blockRedirectUntil: window.__blockRedirectUntil,
  remainingMs: (window.__blockRedirectUntil || 0) - Date.now(),
  routerPatched: window.next?.router?.__isPatched,
  pushSource: String(window.next?.router?.push),
  replaceSource: String(Location.prototype.replace),
  assignSource: String(Location.prototype.assign),
})
```

Look for `isBlockedUrl`, `__reloginActive`, or `__blockRedirectUntil` in function source.

As a temporary diagnostic only:

```js
window.__reloginActive = false;
window.__blockRedirectUntil = 0;
```

Retry the same header destination. If root/dashboard starts working, the global patch is implicated.

## B. Trace host router events and actual history writes

```js
const r = window.next?.router;

for (const event of [
  "routeChangeStart",
  "beforeHistoryChange",
  "routeChangeComplete",
  "routeChangeError",
]) {
  r?.events?.on(event, (...args) => console.log("[router]", event, ...args));
}

const originalPushState = history.pushState;
const originalReplaceState = history.replaceState;

history.pushState = function (...args) {
  console.trace("[history.pushState]", ...args);
  return originalPushState.apply(this, args);
};

history.replaceState = function (...args) {
  console.trace("[history.replaceState]", ...args);
  return originalReplaceState.apply(this, args);
};
```

Interpretation:

- No `routeChangeStart`: click handler or patched `push` suppressed it.
- `routeChangeStart`, no history write/complete: router-instance or loading failure.
- `routeChangeComplete` with old UI still visible: stale remote/unmount problem.
- History changes but header’s `useRouter()` remains old: multiple router module/context identity is likely.

## C. Confirm that leaving issuance is a hard document navigation

Before clicking:

```js
sessionStorage.setItem("__nav_probe__", crypto.randomUUID());
performance.getEntriesByType("navigation").map(x => ({
  type: x.type,
  startTime: x.startTime,
  name: x.name
}));
```

Enable DevTools “Preserve log,” then click another header menu.

Expected current behavior from the code:

- a new document request,
- navigation timing type `navigate`,
- all prior JS globals/listeners reset.

If only a client-side route transition occurs, the deployed header-footer does not contain or is not executing the hard-navigation guard.

## D. Inspect federation share scope

```js
const scope = globalThis.__webpack_share_scopes__?.default;
Object.fromEntries(
  Object.entries(scope || {})
    .filter(([k]) =>
      /^(next\/router|react|react-dom|@ordentco\/addons-auth-provider|@bri\/addons-auth-provider)$/.test(k)
    )
    .map(([k, versions]) => [k, Object.keys(versions)])
);
```

Expected healthy result:

- one chosen `next/router` singleton provider,
- host React selected,
- one effective Ordent package version.

Multiple available versions are not automatically a failure; inspect which entry has `loaded: 1` and `from`.

## E. Compare automatically exposed router modules

If the runtime is available in the console/debug harness:

```js
const hostRouterModule =
  await __FEDERATION__.__INSTANCES__[0].loadRemote("qcash-ui/next/router");

const issuanceRouterModule =
  await __FEDERATION__.__INSTANCES__[0].loadRemote(
    "qcash-ui-issuance-lc/next/router"
  );

console.log({
  sameModule: hostRouterModule === issuanceRouterModule,
  sameRouterExport: hostRouterModule.Router === issuanceRouterModule.Router,
  hostSingleton: hostRouterModule.Router?.router,
  issuanceSingleton: issuanceRouterModule.Router?.router,
});
```

Exact runtime API may vary, but the key identity check is whether both remotes expose the same `Router`/singleton object.

## F. Detect listener accumulation

Chrome DevTools-only:

```js
({
  scroll: getEventListeners(window).scroll?.length || 0,
  popstate: getEventListeners(window).popstate?.length || 0,
  storage: getEventListeners(window).storage?.length || 0,
  showDialog: getEventListeners(document).showDialog?.length || 0,
})
```

Record counts:

1. Before issuance.
2. After opening the create wizard.
3. After changing wizard steps or causing rerenders.
4. After soft navigation, if hard navigation is temporarily bypassed.

An increasing `scroll` count confirms the issuance render-time listener leak.

## G. Verify provider ownership in React DevTools

Inspect the hosted tree under `MyApp`:

Expected:

- host BRI `AuthProvider`
- host Ordent `AuthProvider`
- host global stores
- one Next router context
- issuance `QueryClientProvider`
- issuance/QUI provider on relevant pages

Unexpected:

- issuance `_app`’s additional AuthProvider,
- a second independent Next router context,
- duplicate React roots around issuance.

## H. Check deployed build mode

Search loaded resources:

```js
performance.getEntriesByType("resource")
  .map(x => x.name)
  .filter(x => /issuance-lc|remoteEntry|rspack|webpack/i.test(x));
```

Also inspect the issuance `remoteEntry.js` for:

- `next/router` as a shared consume,
- `react`/`react-dom` as shared consumes,
- Rspack runtime markers.

If issuance embeds its own `next/dist/client/router` or React implementation rather than consuming the share scope, the unsafe Rspack/raw-federation path is active.

---

## Bottom line

The codebase does not support a conclusion that two Next router providers normally run in the hosted tree. The expected architecture is one host router with automatically shared `next/router`. The strongest confirmed issue is instead **issuance-specific federated lifecycle residue**, already mitigated by forcing a hard navigation when leaving `/v2/import-lc/issuance-lc`. The highest-value production checks are: confirm that this hard-navigation branch is deployed and reached, inspect the federation share scope for a single loaded router/React, and determine whether the global auth router patch is active when a header click is swallowed.

---

