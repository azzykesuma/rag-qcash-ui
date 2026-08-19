# Trace QUI TanStack integration (@explore subagent)

- **ID**: `ses_ffefda47dffe6sX07vXcr7430Q`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 23:02:06`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Inspect ~/Development/qui and ~/Development/qcash-ui for TanStack Router integration exposed by the QUI design system. Trace providers, navigation hooks/components, router adapters, and qcash-ui host mounting. Focus on whether a remote using new QUI can navigate correctly under Next.js and what context/provider is required. Do not edit. Return concrete findings with exact file/line references and compatibility risks. Desired thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **glob**: Tool: glob
- **read**: Tool: read

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 16: Assistant

# Verdict

The checked-out `~/Development/qui` does **not** contain the QUI Module Federation remote, a TanStack Router integration, navigation hooks/components, router adapters, or a `providers` exposure. It appears to be a newer React 19 component-library/docs rewrite, while `qcash-ui` consumes a separately deployed and opaque remote named `qui`.

Consequently:

- `qcash-ui` does mount `qui/providers` around every client-rendered page.
- It passes **no Next.js router, TanStack router, navigation adapter, or configuration props** into that provider.
- A remote can navigate correctly only if the deployed `qui/providers` implementation internally adapts QUI navigation to the shared Next.js Pages Router, or if the remote uses `next/router` directly.
- A remote that calls TanStack Router hooks such as `useNavigate()` will **not** receive a TanStack Router context from any source visible in these checkouts.
- A remote importing a new QUI navigation context as an npm/workspace package is at significant risk of receiving a different context instance from the remotely loaded `qui/providers`.

## 1. What exists in the new `/qui` checkout

### No TanStack Router dependency or adapter

The UI package has only styles, `clsx`, React Aria, and Tailwind utilities as dependencies:

- `~/Development/qui/packages/ui/package.json:16-21`

Its peer contract is React 19 or newer:

- `~/Development/qui/packages/ui/package.json:30-32`

The docs app has TanStack Form, Pacer, and Virtual—not TanStack Router:

- `~/Development/qui/apps/docs/package.json:17-19`

The only `@tanstack/react-router` reference in the repository is an optional peer declaration of the third-party `fumadocs-core` package in `bun.lock`; there is no direct dependency or source import.

An exhaustive source search found no:

- `@tanstack/react-router`
- `RouterProvider`
- `useNavigate`
- `useRouter`
- navigation context
- Next router adapter
- TanStack router adapter
- history adapter

### No navigation component is exported

The entire public UI surface is currently:

- Button
- Badge
- Tabs

References:

- `~/Development/qui/packages/ui/src/index.ts:1-2`
- `~/Development/qui/packages/ui/src/components/index.ts:1-3`

The Button is a React Aria button rather than a router-aware link:

- `~/Development/qui/packages/ui/src/components/button/button.tsx:3-10`
- `~/Development/qui/packages/ui/src/components/button/button.tsx:23-26`

The sole library-owned context is private Tabs styling/orientation state:

- `~/Development/qui/packages/ui/src/components/tabs/tabs.tsx:28-33`
- `~/Development/qui/packages/ui/src/components/tabs/tabs.tsx:43-52`

It is unrelated to application routing.

### No Module Federation remote or `providers` exposure

The new QUI docs app has a conventional Next configuration containing only MDX, compression, React Compiler, and strict mode:

- `~/Development/qui/apps/docs/next.config.mjs:1-16`

There is no Module Federation plugin or exposure for:

- `./providers`
- `./globals`
- `remoteEntry.js`

The UI package exports source TypeScript directly:

- `~/Development/qui/packages/ui/package.json:5-12`

Therefore this checkout cannot, as currently configured, produce the remote consumed by `qcash-ui` at `qui/providers`.

## 2. How `qcash-ui` registers and loads the deployed QUI remote

QUI is registered as a global remote:

- `~/Development/qcash-ui/constants/features/registry.ts:37-53`

Its entry is expected at:

```text
${config.baseFederation}/design-system/static/remoteEntry.js
```

Specifically:

- `~/Development/qcash-ui/constants/features/registry.ts:49-52`

`baseFederation` is derived by removing `/api` from the selected environment API URL:

- `~/Development/qcash-ui/config.ts:16-26`

Federation initialization:

1. Enumerates global and route-specific features.
2. Resolves each remote entry.
3. Registers them through `@module-federation/runtime`.
4. Loads global remote styles.

References:

- `~/Development/qcash-ui/services/federation/init.ts:32-46`
- `~/Development/qcash-ui/services/federation/init.ts:48-66`
- `~/Development/qcash-ui/services/federation/init.ts:72-84`

QUI is specially listed as an ETag-cached remote, although the timestamp-stripping plugin is currently commented out:

- `~/Development/qcash-ui/services/federation/init.ts:10`
- `~/Development/qcash-ui/services/federation/init.ts:17-29`
- `~/Development/qcash-ui/services/federation/init.ts:63`

In local development, root-relative `/design-system/...` assets are proxied to the design-system host:

- `~/Development/qcash-ui/next.config.js:120-143`

That rewrite explicitly uses `basePath: false`:

- `~/Development/qcash-ui/next.config.js:137-142`

This confirms that the deployed QUI artifact is a distinct design-system service, not the source currently present under `~/Development/qui`.

## 3. Exact host mounting and provider chain

Federation is initialized only in the browser and is awaited before normal module evaluation continues:

- `~/Development/qcash-ui/pages/_app.tsx:40-51`

The host then dynamically loads:

```tsx
loadRemote("qui/providers")
```

with SSR disabled:

- `~/Development/qcash-ui/pages/_app.tsx:60-68`

The root provider order is:

```text
ModuleFederationProvider
  BriProviders
    QuiProviders
      AuthBridgeSync
      OnboardingTourProvider
        PromoCarouselProvider
          page/layout/remotes
```

References:

- `~/Development/qcash-ui/pages/_app.tsx:160-170`
- `~/Development/qcash-ui/pages/_app.tsx:184-194`

Thus all normal page components and feature remotes are descendants of `QuiProviders`.

### Important: the provider receives no adapter

The declaration is typed as `any`:

- `~/Development/qcash-ui/pages/_app.tsx:66-68`

The mount supplies only children:

- `~/Development/qcash-ui/pages/_app.tsx:167`
- `~/Development/qcash-ui/pages/_app.tsx:192`

There is no:

```tsx
<QuiProviders router={router}>
<QuiProviders adapter={...}>
<QuiProviders navigate={...}>
<QuiProviders basePath="/v2">
```

Therefore any Next.js integration must be implemented internally by the deployed provider. The `any` erases the provider contract, making missing or newly required props invisible to TypeScript.

### Other router providers do not supply QUI navigation context

`BriProviders` obtains the Next Pages Router and passes it into the BRI and Ordent **authentication** providers:

- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:18-29`

This router is passed through auth-provider props, not through a general navigation context. It does not create a TanStack Router provider.

## 4. The host router environment

`qcash-ui` is a Next.js 13.5 Pages Router application:

- `~/Development/qcash-ui/package.json:40`
- Pages Router usage: `~/Development/qcash-ui/pages/_app.tsx:22-24`

It uses a `/v2` base path:

- `~/Development/qcash-ui/next.config.js:65-70`

It runs React 18.3.1:

- `~/Development/qcash-ui/package.json:42-43`

The host’s route-sensitive behavior depends on the Next router:

- Federation module selection uses `router.pathname`:  
  `~/Development/qcash-ui/components/federation/provider.tsx:114-125`
- Route changes trigger loading based on `router.pathname`:  
  `~/Development/qcash-ui/components/federation/provider.tsx:143-158`
- Session handling listens for Next `routeChangeComplete`:  
  `~/Development/qcash-ui/hooks/use-session-expiration.ts:97-106`
- The MFE development monitor also listens for `routeChangeComplete`:  
  `~/Development/qcash-ui/components/federation/monitor/sections/DevModeIndicator.tsx:50-63`

Feature pages generally mount remote components client-side and place them inside a remote header/footer layout. Examples:

- `~/Development/qcash-ui/pages/index.tsx:11-20`
- `~/Development/qcash-ui/pages/index.tsx:31-36`
- `~/Development/qcash-ui/pages/user-basic/index.tsx:10-18`
- `~/Development/qcash-ui/pages/user-basic/index.tsx:20-23`

## 5. What context is actually required

### For the currently deployed QUI contract

The only concrete provider requirement visible from the host is:

```tsx
<QuiProviders>{all hosted content}</QuiProviders>
```

loaded from `qui/providers`.

`qcash-ui` already meets that structural requirement at:

- `~/Development/qcash-ui/pages/_app.tsx:167-192`

For navigation to work, the deployed `QuiProviders` must do one of the following internally:

1. Consume the host’s `next/router` context and expose a QUI navigation abstraction backed by `router.push`, `router.replace`, and `router.back`; or
2. Supply a documented adapter/context that all consuming remotes obtain from the exact same federated module instance.

There is no host code explicitly satisfying either contract beyond mounting the opaque provider.

### For a remote using TanStack Router hooks directly

A component calling TanStack Router’s `useNavigate`, `useRouter`, `Link`, or route hooks requires a matching TanStack:

```tsx
<RouterProvider router={router} />
```

from the same `@tanstack/react-router` module/context instance.

No such provider exists in either inspected source tree. `qcash-ui` does not even declare `@tanstack/react-router` as a dependency:

- `~/Development/qcash-ui/package.json:19-50`

Therefore direct TanStack hooks in a hosted remote are not supported by the visible host setup.

Mounting a standalone TanStack `RouterProvider` inside an individual remote would satisfy the hook technically, but it would establish a second router competing with Next.js and would not automatically drive the host’s Pages Router lifecycle.

## 6. Compatibility risks

### Critical: source/deployment mismatch

The checked-out `/qui` cannot explain the deployed `qui/providers` behavior:

- No federation configuration.
- No `providers` module.
- No global-style exposure.
- No router code.
- No TanStack Router dependency.

The navigation contract of the actually deployed design-system remote cannot be audited from this checkout.

### Critical: React 19 QUI versus React 18 host

New QUI requires React 19:

- `~/Development/qui/packages/ui/package.json:30-32`

Its docs application also uses React 19 and Next 16:

- `~/Development/qui/apps/docs/package.json:26-29`

`qcash-ui` uses React 18 and Next 13:

- `~/Development/qcash-ui/package.json:40-43`

The new Tabs implementation uses React 19’s context-provider shorthand:

```tsx
<TabsContext value={...}>
```

- `~/Development/qui/packages/ui/src/components/tabs/tabs.tsx:43-52`

That is not a safe React 18 API contract. A remote built against this new QUI may fail under the host’s shared React 18 runtime. Bundling React 19 privately instead risks duplicate React instances, invalid hooks, and broken context propagation.

### High: context identity across federation boundaries

`QuiProviders` comes from the `qui` remote, but a feature remote using a newly packaged QUI hook could bundle its own copy of the context module.

The host’s explicit federation share list does not include:

- `@qui/ui`
- a QUI navigation package
- `@tanstack/react-router`

See the complete explicit shared configuration beginning at:

- `~/Development/qcash-ui/next.config.js:15-63`

If provider and consumer resolve different copies, the consumer will not see the provider value even though it is visually nested under `<QuiProviders>`.

The safest model is for provider and navigation hooks to be loaded from the same `qui` federated container, or to share the navigation/context package explicitly as a strict singleton.

### High: direct history/TanStack navigation can desynchronize Next.js

If QUI navigation calls `history.pushState`, manipulates `window.location`, or lets an independently mounted TanStack router own history:

- Next’s `router.pathname` may remain stale.
- `routeChangeComplete` may not fire.
- Route-specific remotes may not be loaded.
- Session-expiration behavior may not run.
- Header/footer and page layouts may remain mounted for the previous route.

Relevant dependencies on Next route state:

- `~/Development/qcash-ui/components/federation/provider.tsx:114-125`
- `~/Development/qcash-ui/components/federation/provider.tsx:143-158`
- `~/Development/qcash-ui/hooks/use-session-expiration.ts:97-106`

Navigation under this host should ultimately go through the Pages Router’s `push`, `replace`, or `back`.

### High: `/v2` base-path handling

Next router navigation with application-relative paths can apply the configured base path. Raw links or history writes such as:

```ts
window.location.href = "/feature"
history.pushState({}, "", "/feature")
```

can escape `/v2` or produce a URL the host does not recognize.

Base path:

- `~/Development/qcash-ui/next.config.js:69`

QUI’s own assets already require special root-path treatment:

- `~/Development/qcash-ui/next.config.js:120-143`

A navigation adapter must distinguish logical Next routes such as `/user-basic` from browser URLs such as `/v2/user-basic`.

### High: erased provider contract

The host declares the remote provider as `any`:

- `~/Development/qcash-ui/pages/_app.tsx:66`

A change in new QUI requiring an adapter, router, base path, or context configuration would compile without warning and fail only at runtime.

### Medium: client-only provider and navigation

`QuiProviders` is loaded with `ssr: false`:

- `~/Development/qcash-ui/pages/_app.tsx:66-68`

Any QUI navigation hook/context is consequently unavailable during SSR. This currently matches the predominantly client-only feature-remotes, but a future server-rendered remote or component cannot rely on that provider.

### Medium: provider failure replaces the entire application subtree

Because `QuiProviders` wraps nearly all page content, failure to load `qui/providers` affects the entire application. The generic federation fallback renders an error component rather than preserving navigation context:

- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:41-47`
- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:89-120`

## Bottom line

A remote using the currently checked-out “new QUI” cannot assume TanStack Router navigation will work in `qcash-ui`. The visible new QUI has no router integration at all, while the host supplies only the opaque deployed `<QuiProviders>` wrapper and no adapter props.

For reliable Next.js hosting, the navigation implementation must:

1. Live under `qui/providers`.
2. Adapt to `next/router`, not own browser history independently.
3. Respect the `/v2` base path.
4. Ensure provider and hook/context modules are federation singletons.
5. Support the host’s React 18 runtime—or require the host to upgrade.
6. Avoid direct TanStack hooks unless a matching TanStack `RouterProvider` is deliberately mounted and synchronized with Next.js.

No files were edited.

---

