# Inspect shell release (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_00205cb91ffeg2PSEVMDJKTRYs`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 08:54:20`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only production investigation. In D:\work\qcash-ui inspect exactly git ref refs/remotes/origin/v1.630.3-release-com (commit b6ce2a6), NOT the checked-out branch and do not edit/checkout. Use git show/git grep as needed. Trace: package versions, next/module federation config including shared entries, provider tree in pages/_app and provider components, dynamic federation initialization order, all guard calls, qcash-global/header/footer/mass-transfer remote loading, and error boundaries. This production branch is Ordent-only; ignore current-branch BRI migration code. Determine how a descendant can receive the auth package's default throwing guard, and assess the user's 'remote arrives before auth code' theory. Return exact ref:path:line references and a concrete timeline.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’ll anchor every read to `refs/remotes/origin/v1.630.3-release-com`, verify it resolves to `b6ce2a6`, then reconstruct dependencies, provider composition, federation startup, remote imports, guards, and boundary behavior without touching the worktree.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The ref is verified at `b6ce2a6bf3cf6564b0aa3f56dd396e4fe84b5a1b`. The initial search shows both compile-time federation configuration and a separate runtime federation layer, so I’m separating those paths before judging initialization order.

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Production selects `OrdentProviders` because the auth migration flag is `false` in prod, and both `AuthProvider` and `GlobalStoreProvider` come from dynamically imported Ordent package code. The key risk now is module identity: the auth package is shared with `requiredVersion: false` but not declared `singleton`, so I’m tracing whether host consumers and each remote can bind to different context instances rather than assuming a network race.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The timing split is now clear: global remotes’ `/globals` exposes are awaited before `_app` can render, while the actual header/footer and mass-transfer components remain descendants and cannot render until the dynamic Ordent provider resolves. I’m verifying the exact Ordent artifact behavior and the boundary that converts a context mismatch into the production symptom, while keeping the checked-out branch out of the evidence.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 41: Assistant

Investigated only `refs/remotes/origin/v1.630.3-release-com`, verified at `b6ce2a6bf3cf6564b0aa3f56dd396e4fe84b5a1b`. No checkout or worktree edits were made.

**Verdict**
A descendant can receive `"Function not implemented."` only when its `useAuth()` resolves a different `AuthContext` object from the one provided by the shell:

```tsx
<AuthContext_host.Provider>
  <RemoteUsingAuthContext_remote />
</AuthContext_host.Provider>
```

Physical React ancestry is insufficient when `AuthContext_host !== AuthContext_remote`.

The strongest supported cause is therefore a duplicated Ordent auth-package instance or separate share scope, not a child rendering before the provider. The host explicitly shares Ordent with `requiredVersion: false` and without `singleton: true`, allowing multiple versions/instances.

The user's timing observation is partly correct:

- `qcash-global-component/globals` and `qcash-ui-header-footer/globals` are loaded before the auth provider mounts.
- The actual `qcash-global-component/modal`, header/footer `default`, and mass-transfer page components cannot render before the auth provider.
- Early remote arrival can influence shared-version registration or selection, but it cannot itself make `useContext` return the default. A context-identity split is still required.

**Versions And Config**
- Shell version: `1.0.562`. `refs/remotes/origin/v1.630.3-release-com:package.json:2-3`
- Active production auth dependency: `@ordentco/addons-auth-provider@0.9.136-mfe`, pinned in both manifest and lock. `refs/remotes/origin/v1.630.3-release-com:package.json:27`, `refs/remotes/origin/v1.630.3-release-com:yarn.lock:696-704`
- React/React DOM: `18.3.1`; Next: `13.5.11`; Webpack: `5.94.0`. `refs/remotes/origin/v1.630.3-release-com:package.json:40-43`, `refs/remotes/origin/v1.630.3-release-com:package.json:50`
- Federation packages: `nextjs-mf@8.2.6`, runtime/retry plugin `0.18.1`. `refs/remotes/origin/v1.630.3-release-com:package.json:24-26`
- The lock also contains `@module-federation/runtime@0.0.17` through `nextjs-mf`, alongside direct runtime `0.18.1`. `refs/remotes/origin/v1.630.3-release-com:yarn.lock:433-442`, `refs/remotes/origin/v1.630.3-release-com:yarn.lock:483-497`
- Compile-time federation has empty `remotes` and `exposes`; all remotes are registered dynamically. `refs/remotes/origin/v1.630.3-release-com:next.config.js:8-13`
- `NextFederationPlugin` is installed in every webpack build. `refs/remotes/origin/v1.630.3-release-com:next.config.js:154`

All 47 explicit shared entries use `requiredVersion: false`. `refs/remotes/origin/v1.630.3-release-com:next.config.js:15-63`

Singleton entries are: `@hookform/error-message`, `classnames`, `cross-env`, `html2canvas`, `ni18n`, `react-circular-progressbar`, `react-currency-flags`, `react-infinite-scroll-component`, `react-laag`, and `sass`.

Non-singleton entries include both auth packages, plus `@headlessui/react`, `@heroicons/react`, `@hookform/resolvers`, `ahooks`, `axios`, `chart.js`, `collect.js`, `country-flag-icons`, `crypto-js`, `dayjs`, `downloadjs`, `formik`, `framer-motion`, `i18next`, `jspdf`, `jspdf-autotable`, `lodash`, `moment`, `ramda`, `react-chartjs-2`, `react-currency-input-field`, `react-datepicker`, `react-dayjs`, `react-dropzone`, `react-hook-form`, `react-i18next`, `react-moment`, `react-number-format`, `react-select`, `react-table`, `styled-components`, `use-debounce`, `uuid`, `xlsx`, and `yup`.

Production is Ordent-only:

- Auth migration flag is `false` for prod. `refs/remotes/origin/v1.630.3-release-com:config.ts:472-477`
- Header/footer BRI flag is `false` for prod. `refs/remotes/origin/v1.630.3-release-com:config.ts:479-484`
- Production environment selection returns the final `prod` argument. `refs/remotes/origin/v1.630.3-release-com:utils/index.ts:139-163`
- Mass transfer is enabled in prod. `refs/remotes/origin/v1.630.3-release-com:config.ts:189-194`
- MFE devtools are disabled in prod. `refs/remotes/origin/v1.630.3-release-com:config.ts:521-527`

**Provider Tree**
```text
ModuleFederationProvider
  RemoteErrorBoundary
    Head
    ExternalScripts
    OrdentProviders
      Ordent AuthProvider
        Ordent GlobalStoreProvider
          qui/providers
            OnboardingTourProvider
              PromoCarouselProvider
                ErrorModal
                conditional SessionExpiredNewQlola
                getLayout(page)
                  qcash-ui-header-footer/default
                    page remote
                NotifOnboardingTour
```

The shell composition is at `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:99-136`.

`OrdentProviders` dynamically resolves `AuthProvider` and `GlobalStoreProvider` from the same package, then nests global store inside auth. `refs/remotes/origin/v1.630.3-release-com:components/providers/OrdentProviders.tsx:8-22`, `refs/remotes/origin/v1.630.3-release-com:components/providers/OrdentProviders.tsx:28-38`

Several shell descendants statically import Ordent `useAuth`:

- `refs/remotes/origin/v1.630.3-release-com:components/onboarding-tour/notif-onboarding-tour.tsx:4`
- `refs/remotes/origin/v1.630.3-release-com:components/providers/OnboardingTourProvider.tsx:3`
- `refs/remotes/origin/v1.630.3-release-com:components/providers/PromoCarouselProvider.tsx:3`
- `refs/remotes/origin/v1.630.3-release-com:components/ui/MFAErrorModal.tsx:3`
- `refs/remotes/origin/v1.630.3-release-com:hooks/use-modal-session-expired.tsx:2`

Because static dependencies must resolve before `_app` executes its top-level federation initialization, the shell's auth context module is present before global remote loading begins. The provider mount is deferred, but not all auth code.

**Exact Default Guard**
The auth source is not vendored in Git. I inspected the exact `0.9.136-mfe` package artifact whose SHA and integrity match `yarn.lock`.

- Package root exports both `AuthProvider` and `useAuth` from the same `./auth` module: `@ordentco/addons-auth-provider@0.9.136-mfe:src/index.ts:1`
- Default context guard throws `"Function not implemented."`: `@ordentco/addons-auth-provider@0.9.136-mfe:src/auth.tsx:150-152`
- Context is created with that default object: `@ordentco/addons-auth-provider@0.9.136-mfe:src/auth.tsx:221`
- Real provider guard is created at `src/auth.tsx:295`
- Real guard is placed in the provider value at `src/auth.tsx:1029`
- Provider wraps children at `src/auth.tsx:1106-1113`
- `useAuth` directly calls `useContext(AuthContext)` without validating the provider: `src/auth.tsx:1117`

Therefore, a mismatched package instance silently returns default fields until `guard()` is called.

**Concrete Timeline**
1. `_app` static dependencies load. Some already statically consume Ordent `useAuth`, so the shell auth module/context is resolved before `_app` body execution. `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:21-26`
2. On the browser, `_app` performs a top-level `await initFederation()`. `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:38-49`
3. `federationInit()` collects every global and route feature, computes entries, and registers them with runtime `init`. It does not pass an explicit runtime `shared` map. `refs/remotes/origin/v1.630.3-release-com:services/federation/init.ts:32-68`, `refs/remotes/origin/v1.630.3-release-com:services/federation/utils/registry.ts:11-12`
4. It sets `__FEDERATION_LOADED = true`, then concurrently awaits `/globals` from `qcash-global-component`, `qcash-ui-header-footer`, and `qui`. `refs/remotes/origin/v1.630.3-release-com:services/federation/init.ts:70-88`
5. Those global remotes are defined at `refs/remotes/origin/v1.630.3-release-com:constants/features/registry.ts:37-53`. Their remote entries are `/global-component/_next/static/chunks/remoteEntry.js` and `/header-footer/_next/static/chunks/remoteEntry.js`. `refs/remotes/origin/v1.630.3-release-com:constants/features/registry.ts:30-46`
6. Only after all global `/globals` loads or fallbacks complete does `_app` render. `ModuleFederationProvider` loads first, then starts the dynamic Ordent provider subtree. `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:66-79`, `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:103-109`
7. On a mass-transfer route, `ModuleFederationProvider`'s post-commit effect triggers route preloading. `refs/remotes/origin/v1.630.3-release-com:components/federation/provider.tsx:143-158`
8. Route preloading starts `qcash-ui-mass-transfer/globals`; it can overlap the dynamic provider request and may finish first. `refs/remotes/origin/v1.630.3-release-com:components/federation/hooks/useModuleLoading.ts:81-95`, `refs/remotes/origin/v1.630.3-release-com:components/federation/hooks/useModuleLoading.ts:121-129`
9. The mass-transfer remote URL follows the standard `/mass-transfer-v2/_next/static/chunks/remoteEntry.js` rule. `refs/remotes/origin/v1.630.3-release-com:constants/features/registry.ts:249-253`, `refs/remotes/origin/v1.630.3-release-com:services/federation/utils/registry.ts:81-83`, `refs/remotes/origin/v1.630.3-release-com:services/federation/utils/name.ts:1-3`
10. Once Ordent AuthProvider, GlobalStoreProvider, and `qui/providers` have resolved, the page layout renders `qcash-ui-header-footer/default`. The mass-transfer page is its child and renders afterward. Representative: `refs/remotes/origin/v1.630.3-release-com:pages/mass-transfer/maker-confirmation/index.tsx:13-28`
11. Mass-transfer exposes cover maker confirmation, transaction action, transfer inquiry, their view/detail variants, and upload transaction. The corresponding loaders are at `pages/mass-transfer/**:10` or `:13`; every page loads header/footer at `:16` or `:20`.
12. `NotifOnboardingTour` then loads `qcash-global-component/modal`; its welcome component is nested beneath it. `refs/remotes/origin/v1.630.3-release-com:components/onboarding-tour/notif-onboarding-tour.tsx:28-34`, `refs/remotes/origin/v1.630.3-release-com:pages/_app.tsx:130-131`

**All Shell Guard Calls**
There are exactly three call expressions in this ref:

- Onboarding notification close: `await guard()`. It is enclosed by the handler's catch. `refs/remotes/origin/v1.630.3-release-com:components/onboarding-tour/notif-onboarding-tour.tsx:73-95`
- Session-expiration production path: `await guard()`, locally caught and logged. `refs/remotes/origin/v1.630.3-release-com:hooks/use-modal-session-expired.tsx:239-246`
- Session-expiration alternate path: unawaited `guard()`. `refs/remotes/origin/v1.630.3-release-com:hooks/use-modal-session-expired.tsx:247-249`

The alternate path is inactive in production because `toggleUpdateLoginService` is `true`. `refs/remotes/origin/v1.630.3-release-com:config.ts:676-681`

This inventory cannot include guard calls compiled inside separately deployed remote repositories.

**Error Handling**
- The root auth/page subtree is wrapped by `RemoteErrorBoundary`. `refs/remotes/origin/v1.630.3-release-com:components/federation/provider.tsx:171-176`
- In production it renders the generic full-page error because MFE devtools are off. `refs/remotes/origin/v1.630.3-release-com:components/ui/RemoteErrorBoundary.tsx:75-113`
- Federation load failures are handled separately by `fallbackPlugin`; header/footer gets a special fallback. `refs/remotes/origin/v1.630.3-release-com:services/federation/fallbackPlugin.tsx:19-47`, `refs/remotes/origin/v1.630.3-release-com:services/federation/fallbackPlugin.tsx:123-127`
- That plugin cannot handle a guard thrown after a remote loaded successfully. A synchronous render/commit error reaches `RemoteErrorBoundary`; async event-handler errors do not.
- Mass-transfer pages have no local `RemoteErrorBoundary`; they use raw `next/dynamic(loadRemote)` wrappers.
- `pages/_error.tsx` is a Next error page, not a React error boundary. `refs/remotes/origin/v1.630.3-release-com:pages/_error.tsx:7-16`

**Theory Assessment**
“Remote arrives before auth code” is not the direct cause:

- Shell auth context code is statically consumed before `initFederation()` can run.
- Dynamic auth wrappers do not render their children while loading, so there is no transient render outside the provider.
- Context propagation does not depend on fetch completion order when provider and consumer share the same context object.

Early remote arrival is still a plausible trigger for the actual problem:

- Global remote containers initialize before the provider mounts.
- Mass-transfer `/globals` can overlap provider initialization.
- The Ordent package is non-singleton with unrestricted version acceptance.
- A remote can therefore consume or bundle a different Ordent instance, or operate in another runtime/share scope. Its `useAuth()` then sees its own default context and returns the throwing guard.

This ref proves that configuration permits the failure, but it cannot identify which deployed remote/version caused it because remote build configs and production share-scope contents are external. The built-in diagnostic already reads all runtime share maps and groups auth versions by remote at `refs/remotes/origin/v1.630.3-release-com:components/federation/monitor/auth/provider-versions.ts:10-56`; a production share-scope snapshot is the missing confirmation.

---

