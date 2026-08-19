# Inspect deployed shell route (@general subagent)

- **ID**: `ses_049a3da01ffeRygvkv1P0U15fv`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 11:08:40`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Analyze the saved deployed shell page chunk ~\.local\share\opencode\tool-output\tool_fb6522ea0001o3rwPsGi7ASQHL and app chunk ~\.local\share\opencode\tool-output\tool_fb6522da1001Ubd0J7sLXQXcrx. Research only. Determine how /import-lc/issuance-lc mounts/unmounts the remote, provider selection in deployed dev, actual auth provider and Next versions, and any code/effect that could block host route component replacement after URL push. Search for router events, dynamic loading lifecycle, error boundaries, stale component caching. Return root-cause likelihoods based only on deployed evidence.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **read**: Tool: read

</details>

I’ll inspect both saved chunks, map their module boundaries and routing/runtime references, then correlate only evidence present in the deployed code.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The chunks are minified into oversized lines, so ordinary indexed reads cannot expose the relevant sections. I’m switching to bounded context extraction around deployed symbols and literals, without modifying either artifact.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The deployed page is a Pages Router entry with two independent client-only remotes: the issuance remote and its header/layout remote. The auth tree mounts both BRI and Ordentco providers; route-specific authority is chosen later by the bridge rather than by conditionally mounting one provider in `_app`.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

**Conclusion**
The deployed host contains no mechanism that should prevent a genuinely different Next.js page component from replacing `/import-lc/issuance-lc`. The strongest stale-page mechanisms are a URL update that does not constitute a different Next route, or same-page/shallow navigation: the remote wrapper has no route key and its loading effect runs only once.

**Mount Lifecycle**
Evidence: page chunk modules `80101`, `34275`, `81367`.

- `/import-lc/issuance-lc` is registered through `window.__NEXT_P`, confirming the Pages Router.
- Module `34275` creates two client-only remotes:
  - `qcash-ui-issuance-lc/issuance-lc`
  - `qcash-ui-header-footer/default`
- `getLayout` wraps the issuance page in the header remote with `isNewUI: true`.
- `toggleIssuanceLc` is `true` in every deployed environment, including dev.
- On first React mount, `createRemoteDynamic` runs a `useEffect(..., [])`.
- It first loads `qcash-ui-issuance-lc/globals`.
- The globals load has a 15-second default timeout, waits five seconds after failure, and retries once.
- It then constructs a client-only `next/dynamic` loader for `qcash-ui-issuance-lc/issuance-lc`.
- On React unmount, the wrapper aborts an `AbortController`, but the signal is only checked before setting state. It is not passed to `loadRemote`, the timeout, or the retry.
- Therefore, route departure unmounts the React remote UI but does not cancel its network/runtime load and does not unload the federation container or cached JavaScript.
- A later mount can reuse Module Federation and Next dynamic caches.

The delayed load cannot resurrect the old component after unmount because the aborted signal suppresses its state update.

**Provider Selection**
Evidence: app chunk modules `58013`, `64520`, `27084`.

The persistent `_app` tree always mounts both custom auth systems:

```text
ModuleFederationProvider
  BriProviders
    qui/providers
      AuthBridgeSync
      OnboardingTourProvider
      page layout/page
```

There is no NextAuth/Auth.js usage. Neither `next-auth` nor `@auth/*` occurs in either artifact.

For dev, local, and `gym-master` hostnames:

- `toggleAddonsAuthProviderMigration` evaluates to `false`.
- Despite that toggle, `_app` still mounts both provider wrappers. The toggle is not used for provider selection in the supplied `_app` or bridge code.
- `AuthBridgeSync` performs the effective route/session selection.

For `/import-lc/issuance-lc`, which is not one of the BRI-special paths:

- Ordentco wins whenever it has a non-guest session.
- BRI is selected only when Ordentco is not non-guest and BRI is non-guest.
- If both are guests, the exposed guest bridge defaults to Ordentco.
- The BRI-special paths are `/helpdesk`, `/plafon-pegadaian`, and `/systems`.
- The same `access-token` is synchronized into both providers.
- Ordentco `guard()` is invoked after route changes when its token exists and authority state needs hydration.

Embedded provider package versions are:

- `@bri/addons-auth-provider`: `0.10.6-mfe`
- `@ordentco/addons-auth-provider`: `0.9.142-mfe`

The bridge separately inspects `globalThis.__FEDERATION__` to report versions actually consumed by each remote. The saved artifact contains that inspector but not its runtime output, so it cannot prove which shared provider version the issuance remote ultimately consumed.

**Next Version**
The evidence proves use of:

- Next.js Pages Router
- `window.__NEXT_P`
- `/_app`
- `useRouter`
- `next/dynamic` with `ssr: false`

No Next.js version string, build metadata, or federation share declaration for `next` exists in either saved chunk. An exact Next version cannot be established from these artifacts alone.

**Routing Effects**
Only one router event subscription exists:

```text
router.events.on("routeChangeComplete", checkTokenAndModal)
```

It controls session-expiration modal state. There are no deployed references to:

- `routeChangeStart`
- `routeChangeError`
- `beforeHistoryChange`
- `beforePopState`
- `pushState`
- `replaceState`
- `popstate`

Navigation-capable effects found in `_app` are:

- A stored `auth_redirect_path` can call `router.replace(storedPath)` after an access token appears.
- Pending or draft onboarding status can call `router.replace("/onboarding-tour...")`.
- Completed onboarding can replace an onboarding route with `/`.
- `?logout=true` performs a full `window.location.replace`.
- Certain API authentication failures perform full redirects to `/landing-page`.
- Ordentco `guard()` runs on `asPath` changes; its implementation is in an external provider chunk and cannot be evaluated here.

None of these effects explicitly cancels a route transition. Some can supersede a completed push with a second navigation.

**Replacement And Caching**
Important deployed characteristics:

- The issuance remote string is fixed at module evaluation.
- Its loading effect has an empty dependency array.
- Neither `_app` nor `getLayout` keys the page or remote by `pathname` or `asPath`.
- A query-only, shallow, rewritten, or otherwise same-page navigation can preserve the existing wrapper and remote instance.
- A raw History API update originating inside the unavailable issuance remote would change the URL without informing the Next router. The supplied host chunks contain no listener that would compensate for that.
- Navigating to a genuinely different Next page component should change the rendered component type and unmount issuance normally.
- There is no custom host page cache, component map, `memo` wrapper, or stale-element store in these artifacts.
- Module Federation and `next/dynamic` can cache loaded code, but no deployed code uses that cache to choose which host route component remains rendered.

**Errors And Overlays**
No custom React error boundary appears in either artifact:

- No `ErrorBoundary`
- No `componentDidCatch`
- No `getDerivedStateFromError`

`createRemoteDynamic` catches remote-loading failures and renders an `ErrorAlert`, but it does not catch errors thrown while the loaded remote renders or runs effects. Such errors fall through to Next/React runtime handling outside the supplied chunks.

Persistent `_app` modals can visually cover or hide a destination:

- MFA state applies `display: none` to the page container.
- Session-expiration and MFA portals remain mounted across routes.
- Loading overlays are fixed and high-z-index, but the issuance overlay is removed when its page wrapper unmounts.
- The header may legitimately persist when destination layouts reuse the same header remote.

**Root-Cause Likelihoods**
1. **High, conditional: URL changed without a different Next page transition.** Raw `history.pushState`, shallow navigation, query-only navigation, or a rewrite resolving to the same page would retain the issuance wrapper because there is no route key and its effect is mount-only. The issuance remote code itself is not present, so raw History API use cannot be confirmed.

2. **Medium: a competing auth/onboarding redirect supersedes the push.** `auth_redirect_path`, onboarding replacement, and external Ordentco `guard()` all react to route/session changes. This should normally produce another URL transition rather than silently retain the issuance URL.

3. **Medium-low: destination load or render failure.** There is no local render error boundary, and only loader promises are caught. A destination error could prevent a clean replacement, but no actual destination error is present in these artifacts.

4. **Medium-low as a visual explanation: persistent modal or header UI.** Persistent `_app` portals or a shared header can make navigation appear stale even when the page changed.

5. **Low: stale remote/module cache blocking host replacement.** Federation code remains cached and in-flight loads are not cancelled, but the aborted wrapper cannot update after unmount, and no cache is wired to host route selection.

6. **Very low: host router-event interception.** The only router listener observes `routeChangeComplete`; no start/error/history hook capable of blocking replacement exists in the supplied deployed code.

---

