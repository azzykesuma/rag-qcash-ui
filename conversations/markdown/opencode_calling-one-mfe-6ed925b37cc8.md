# Trace host me lifecycle (@explore subagent)

- **ID**: `opencode-8836d43ede06e871fbd94d40480b23492084e169036815f8e4966ed925b37cc8`
- **Source Tool**: `opencode`
- **Date**: `2026-09-01 10:46:32`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Very thorough read-only investigation in D:\work\qcash-ui. Trace every path that calls the `me` API (especially host-network-event initiators), auth singleton/provider initialization, route/redirection behavior, and MFE mount/unmount/render handling. Inspect recent commits c94e5db5, ab8ae0f0, 0e146981, c0f353e5 and relevant diffs/history. Determine concrete mechanisms capable of repeatedly calling /me when one MFE rerenders or emits events. Return file:line references, commit evidence, and likely causal chain. Do not edit files.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 18: Assistant

## Conclusion

The shell has one intended `/me` owner, `HostAuthGate`, but it does not enforce exclusivity over federated code. The most credible causes of repeated `/auth/me` are:

1. **A remote still calls `guard()` during mount/render/effects.**
2. **Commit `ab8ae0f0` removed the only request-level `/auth/me` and `/menu/me` coalescer.**
3. **That same commit made `HostAuthGate` stop rendering the entire MFE subtree whenever `briAuth.isAuthoritiesReady` becomes false**, allowing guard-driven unmount/remount cycles.
4. **The session-expired flow can replay the same failed `/auth/me` twice and can concurrently trigger host revalidation.**
5. **A remote calling `checkToken()` can recursively call `/auth/me?accessToken=...` indefinitely on HTTP 504.**
6. **Separate/nested auth-provider instances are still possible even though the package is configured as a Module Federation singleton.** Singleton package code does not guarantee one React provider instance.
7. In development, `reactStrictMode: true` can double-run mount effects, exposing any non-idempotent remote `guard()` effect.

No literal `host-network-event`, `network-event`, or corresponding initiator implementation exists in the current source, `.next` output, relevant dependencies searched, or the four commits. The host event buses involved are `qc-bridge-sync` and `showDialog`.

---

# 1. Every host-visible `/me` path

## 1.1 Normal protected-route bootstrap

Provider hierarchy:

- `D:\work\qcash-ui\pages\_app.tsx:171-207`
- `D:\work\qcash-ui\components\providers\BRIProviders.tsx:14-22`

`BRIProviders` mounts exactly one shell-owned BRI provider:

```text
BRIAuthProvider
  BRIGlobalStoreProvider
    HostAuthGate
      QuiProviders
        page/remotes
```

`HostAuthGate`:

- Reads `access-token`:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:35`
- Synchronizes it into provider state:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:37-44`
- Calls `auth.guard()`:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:89-94`
- Starts validation:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:223-248`

Inside `@bri/addons-auth-provider`:

- `guard()` calls `validateToken()`:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:332-339`
- `validateToken()` performs `GET /auth/me`:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\services\auth.js:79-87`
- The guard then always proceeds to `validateMenu()`:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:356-359`
- `validateMenu()` performs `POST /menu/me`:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\services\auth.js:115-127`

Therefore a normal uncached host guard produces this pair:

```text
GET  /auth/me
POST /menu/me
```

The shell’s window-level state prevents additional **host-owned** guards for the same token:

- State key/type:  
  `D:\work\qcash-ui\components\providers\host-auth-events.ts:1-2`  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:22-28`
- Validated-token check:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:200-216`
- In-flight promise installation:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:223-245`

This does **not** intercept a remote calling `guard()` directly.

## 1.2 Auth monitor manual health check

The federation auth monitor has a second explicit caller:

- `D:\work\qcash-ui\components\federation\monitor\auth\index.tsx:110-134`

It runs:

```ts
customAxios.get("/auth/me")
```

This is user-triggered by the monitor UI, not automatic on `qc-bridge-sync`. The event listener only refreshes bridge/version state:

- `D:\work\qcash-ui\components\federation\monitor\auth\index.tsx:78-108`

## 1.3 Session-expired retry path

Failed requests dispatch `showDialog`:

- Shell axios:  
  `D:\work\qcash-ui\services\custom-axios.ts:95-123`
- Provider’s own axios:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\services\custom-axios.js:88-149`

The shell listens and opens the relogin UI:

- `D:\work\qcash-ui\hooks\use-session-expiration.ts:46-76`
- `D:\work\qcash-ui\pages\_app.tsx:188-195`

The relogin hook specially queues `auth/me`:

- `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:202-216`

After successful login it retries every queued request:

- `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:218-253`

It then enters a second fallback retry block:

- `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:255-267`

That block excludes only `menu/me`, not `auth/me`:

```ts
if (!originalRequest?.url?.endsWith("menu/me") && resolve && reject) {
  await retryExceptionV2(...)
}
```

Consequently, when the failed original request is `/auth/me`:

1. It is pushed into `requestQueue`.
2. It is retried by `requestQueue.current.map(...)`.
3. The queue is cleared.
4. The second block retries the same original `/auth/me` again.

The Promise may already be settled after the first resolution, but the second `customAxios(originalRequest)` is still constructed and can still send another HTTP request.

At the same time, successful login:

- Stores the new token:  
  `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:97-102`
- Calls provider `setToken`:  
  `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:108-113`
- Emits `qc-bridge-sync`:  
  `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:114`

The token-state change can make `HostAuthGate` validate the new token as well. With no request coalescer, the observable result can be:

```text
retry original /auth/me       # queue map
retry original /auth/me again # fallback block
host guard /auth/me           # token-specific validation
```

Whether the third call reaches the network depends on the provider’s in-memory freshness state.

There is no test covering an `auth/me` original request. The existing test only covers `menu/me` and explicitly expects no retry:

- `D:\work\qcash-ui\hooks\__tests__\use-modal-session-expired.test.tsx:269-321`

## 1.4 Public `checkToken()` API exposed to remotes

The provider exposes `checkToken` through context:

- Declaration:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.d.ts:55`
- Context value:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:1030-1123`

`checkToken()` performs:

```text
GET /auth/me?accessToken=<token>
```

References:

- Call:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:953-968`
- Endpoint implementation:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\services\auth.js:88-96`

On 504 it recursively invokes itself with no delay or limit:

- `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:969-978`

Thus any remote invoking `checkToken()` can generate an unbounded `/auth/me?accessToken=...` sequence during persistent 504 responses. No shell source currently calls it, so this path can only be activated by provider consumers/remotes.

## 1.5 Federated remote calls

Remote source code is not contained in this checkout. The host has 409 page files using `loadRemote`, so the opaque remote surface is large.

The repository’s migration documentation explicitly records that remotes historically called `guard()`:

- `D:\work\qcash-ui\docs\centralized-auth-guard-migration.md:7-13`
- `D:\work\qcash-ui\docs\centralized-auth-guard-migration.md:19-25`
- `D:\work\qcash-ui\docs\ordent-to-bri-provider-migration.md:19-28`

The required contract is clear: remotes must consume auth state and must not call `guard()`. The shell cannot statically verify that deployed remotes comply.

---

# 2. Auth singleton/provider initialization

## Shell provider

- Package/version:  
  `D:\work\qcash-ui\package.json:19-21` — `0.10.6-mfe`
- Host Federation singleton declaration:  
  `D:\work\qcash-ui\next.config.js:15-17`
- Provider mount:  
  `D:\work\qcash-ui\components\providers\BRIProviders.tsx:17-21`

Inside the package:

- Token initializes from local storage:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:235-240`
- API service is recreated for `apiUrl` or token changes:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:241`
- `apiUrl` is published on provider mount/change:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:269-275`
- Context provider is rendered at:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:1124-1128`

Important distinction:

> Module Federation `singleton: true` means one selected module implementation in a share scope. It does not prevent a remote from rendering another `<AuthProvider>` instance.

A nested remote provider has separate token/readiness/freshness state and can independently call `/auth/me`.

Also, every remote must configure the package as singleton. The host setting alone cannot force a remote using another share scope/version to consume the host copy. Runtime detection exists here:

- `D:\work\qcash-ui\components\federation\monitor\auth\provider-versions.ts:7-49`

## Provider cache behavior

The provider has two freshness mechanisms:

- Session storage cache is used only when `guard(true)` is called:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:283-318`
- In-provider five-minute timestamp:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:319-326`

The host calls `guard()` without `true`, so it skips the session-storage restore path. A repeated guard in the same mounted provider may still be stopped by the in-memory timestamp. Concurrent guards before the first one sets that timestamp are not coalesced.

There is no guard-level in-flight Promise inside the provider.

---

# 3. Event initiators

## `qc-bridge-sync`

All production dispatchers found:

1. Host hydration:
   - `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:46-84`
2. Session relogin:
   - `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:114`
   - `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:157`
3. Onboarding refresh-token flow:
   - `D:\work\qcash-ui\components\onboarding-tour\notif-onboarding-tour.tsx:73-86`

Host listener:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:131-150`

It only resets validation when:

```ts
detail.source === "logout"
```

Therefore:

- `source: "host-auth-gate"` does not invoke guard.
- `source: "token-change"` does not directly invoke guard.
- The auth monitor rerenders, but does not automatically perform its health check.
- A relogin causes host validation through `setToken`/context changes, not through the event itself.

The onboarding flow only changes local storage and emits the event; it does not call provider `setToken`:

- `D:\work\qcash-ui\components\onboarding-tour\notif-onboarding-tour.tsx:76-86`

So that event alone does not cause `/me` in current shell code.

## `showDialog`

This is the meaningful network-related event:

1. A 401/499/408 dispatches `showDialog`.
2. `HostAuthGate` sets `sessionDialogPending`:  
   `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:152-161`
3. `useSessionExpiration` opens the modal:  
   `D:\work\qcash-ui\hooks\use-session-expiration.ts:46-76`
4. Successful submit logs in, changes token, and replays requests.

A notable mount hazard exists at:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:272-274`

```tsx
if (sessionDialogPending) {
  return <>{children}</>;
}
```

This renders children unconditionally, even if protected children had never successfully mounted. During an initial expired-token guard, `showDialog` can therefore expose/mount remotes while the host guard is still suspended. Any remote mount-level `guard()` can then execute concurrently with the host’s pending validation.

---

# 4. MFE render, mount and unmount behavior

## Host rendering

The host does not explicitly call React `createRoot`, `render`, or `unmount` for remotes. It uses `next/dynamic` and Module Federation components.

Examples:

- Dashboard and header/footer:  
  `D:\work\qcash-ui\pages\homepage\index.tsx:9-27`
- Helpdesk and alternate header/footer:  
  `D:\work\qcash-ui\pages\helpdesk\portal\index.tsx:10-29`
- Account management:  
  `D:\work\qcash-ui\pages\account-management\index.tsx:10-27`

These dynamic component definitions are module-scoped and normally stable across ordinary parent rerenders. Route changes replace the page component and unmount the old route subtree.

## Host gate can unmount all remotes

Current render condition:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:276-288`

```tsx
if (!token || !isHostValidated || !briAuth.isAuthoritiesReady || !hasMountedRef.current) {
  return null;
}
```

Once any shared-context guard sets `isAuthoritiesReady` false, the host returns `null` and unmounts:

- `QuiProviders`
- header/footer remote
- current page remote
- onboarding/promo children
- all page-level effects

When readiness returns true, those components mount again.

That is capable of converting an MFE’s “guard once on mount” code into repeated guard execution.

The test suite checks initial readiness and explicit gate remounts but does not test this critical transition after children have already mounted:

- `D:\work\qcash-ui\components\providers\__tests__\HostAuthGate.test.tsx:101-151`

## Error retries force remounts

`RemoteErrorBoundary` increments a keyed wrapper on retry:

- `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:47-60`
- `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:105-106`

That deliberately performs a full subtree remount. After three failures it performs a page reload:

- `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:53-56`

Federation fallback paths also reload the entire page:

- `D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:70-75`
- `D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:97-102`

A full reload destroys `window.__QCASH_HOST_AUTH_GUARD__`, so the next protected boot calls `/auth/me` again.

## `createRemoteDynamic`

- Loads styles and creates a `next/dynamic` component:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:9-25`
- Loads once per wrapper mount:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:68-76`
- Retry creates/replaces the dynamic component:  
  `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:82-99`

Thus parent unmounts or error retries can rerun remote mount effects.

## React Strict Mode

- `D:\work\qcash-ui\next.config.js:67`

With `reactStrictMode: true`, development mount effects can execute twice. The host’s window-level `validating` state protects the host caller, but an MFE’s own direct `guard()` effect has no shell-level protection.

## Hiding the page for MFA does not unmount it

The response/MFA modal merely applies `display: none`:

- `D:\work\qcash-ui\pages\_app.tsx:196-200`

That path does not itself rerun mount effects. The dangerous unmounts are the `HostAuthGate` null return, route changes, keyed error retries, and full reloads.

---

# 5. Route and redirection behavior

## Protected routes

- Public route list:  
  `D:\work\qcash-ui\constants\routes.ts:1-10`
- Route normalization:  
  `D:\work\qcash-ui\components\providers\auth-provider-migration-routes.ts:1-9`
- No token on protected route → `router.replace("/landing-page")`:  
  `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:185-195`

## Root route

SSR redirects based on `loggedIn` cookie:

- `D:\work\qcash-ui\pages\index.tsx:23-28`

The client gate instead uses local storage. A stale cookie or missing local token can therefore yield SSR access followed by a client-side redirect.

## Saved redirect path

- Save protected path when cookie exists but token does not:  
  `D:\work\qcash-ui\pages\_app.tsx:79-92`
- Poll every 500 ms on landing page and redirect once token appears:  
  `D:\work\qcash-ui\pages\_app.tsx:113-146`

That navigation causes the host’s protected-route guard once for the active token.

## Logout reload

- `D:\work\qcash-ui\pages\_app.tsx:94-99`

`?logout=true` triggers `window.location.replace(window.location.pathname)` to flush stale singleton/provider state. That full reload also clears the window guard and guarantees a future protected-route bootstrap will call `/me`.

## Landing/main-page redirect cycle possibility

Landing page, when disabled:

- `D:\work\qcash-ui\pages\landing-page\index.tsx:16-20`
- Redirects via full `location.href` to `/main-page`.

Main page, when disabled:

- `D:\work\qcash-ui\pages\main-page\index.tsx:27-35`
- Redirects back to `/landing-page`.

Configuration has both disabled in several environments:

- `D:\work\qcash-ui\config.ts:49-55`
- `D:\work\qcash-ui\config.ts:578-584`

This can form a navigation cycle. Both routes are public, so the shell does not directly guard them, but the main-page layout is still assigned at:

- `D:\work\qcash-ui\pages\main-page\index.tsx:38-40`

A public remote/layout that owns auth logic could repeatedly mount during the cycle.

---

# 6. Commit evidence

## `c0f353e5` — centralized bridge/gate

Commit: `c0f353e5268b2f5d75bb5be7571439736b51d867`  
Date: 2026-08-31 08:57 +0700

This commit:

- Removed `AuthBridgeSync`.
- Added host BRI/Ordent gate selection and bridge publication.
- Added persistent `hasMountedRef`.
- Removed readiness from the final persistent render condition once children had mounted.

Most importantly, its migration documentation explicitly recorded:

- Concurrent host/remote guards caused `/auth/me` and `/menu/me` storms.
- A remote guard set readiness false.
- Host stopped rendering children.
- The remote unmounted and remounted.
- Its mount effect called guard again.

Historical references at that commit:

- `c0f353e5:docs/centralized-auth-guard-migration.md:12`
- `c0f353e5:docs/centralized-auth-guard-migration.md:146-165`

This is direct repository evidence for the exact class of failure under investigation.

## `0e146981` — BRI-only provider

Commit: `0e14698114477bfd6c7b992efe49e2c6c0ed28a6`  
Date: 2026-08-31 10:41 +0700

It:

- Removed the Ordent provider/package.
- Made BRI the guard owner for all protected routes.
- Retained the `/auth/me` and `/menu/me` coalescer.
- Reduced final rendering to `!token || !hasMountedRef.current`, thereby keeping already-mounted children alive during transient readiness changes.
- Replaced the long migration document with a strict “remotes never call guard” contract.

## `ab8ae0f0` — deduper removal and readiness regression

Commit: `ab8ae0f046b4163e48d1edfa9addda1bbdc0cfdd`  
Date: 2026-08-31 15:13 +0700

It deleted:

- `D:\work\qcash-ui\components\providers\host-auth-guard-interceptor.ts`
- Initialization from `BRIProviders`.
- Initialization from `_app`.

The deleted interceptor:

- Coalesced in-flight `/auth/me` and `/menu/me`.
- Cached successful responses for five seconds.
- Was explicitly intended to suppress remote rerender/guard storms.

Historical lines:

- `ab8ae0f0^:components/providers/host-auth-guard-interceptor.ts:12-17`
- `ab8ae0f0^:components/providers/host-auth-guard-interceptor.ts:27-83`
- `ab8ae0f0^:components/providers/host-auth-guard-interceptor.ts:85-103`

The same commit changed the final gate from:

```ts
if (!token || !hasMountedRef.current) return null;
```

to the current volatile condition:

```ts
if (!token || !isHostValidated || !briAuth.isAuthoritiesReady || !hasMountedRef.current) return null;
```

Current location:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:276-285`

This is the strongest regression relevant to MFE-driven unmount/remount behavior.

## `c94e5db5` — redirect fixes, not a `/me` dedupe fix

Commit: `c94e5db5c76f85071787ae4ebd44ecb1090350a8`  
Date: 2026-08-31 19:58 +0700  
Message: `fix : fixing redirection and double hit me`

Actual changes were only:

1. `/landing-page/?logout=true` → `/landing-page?logout=true`
2. Landing component selection no longer depended on `globalThis.window !== undefined`.

It did not change:

- `HostAuthGate`
- Provider guard behavior
- `/auth/me` calls
- Request coalescing
- Session-expired retry logic

The landing selection change can prevent the landing remote from being selected during server/client disagreement when the feature is disabled, so it may remove one indirect mount path. It does not solve concurrent host/remote guard calls.

---

# 7. Most likely causal chains

## Chain A: continuous calls correlated with an MFE rendering

Most likely when the remote has an executable guard effect:

```text
Host validates
  -> GET /auth/me
  -> POST /menu/me
  -> host marks validated/readiness true
  -> protected subtree mounts
  -> remote mount/rerender calls guard()
  -> provider sets isAuthoritiesReady false
  -> current HostAuthGate returns null
  -> entire MFE subtree unmounts
  -> guard completes, readiness becomes true
  -> subtree mounts again
  -> remote mount effect calls guard again
```

Amplifiers:

- Deduper removed by `ab8ae0f0`.
- Remote uses a separate/nested provider, so freshness state is recreated.
- Remote package is not actually sharing the host package.
- Remote effect depends on readiness or unstable auth values.
- React Strict Mode double-runs mount effects.
- Error-boundary retry or route navigation remounts the remote.

## Chain B: exactly two or three `/auth/me` calls after relogin

```text
Original /auth/me gets 401
  -> provider axios emits showDialog
  -> relogin succeeds
  -> hook setToken(newToken)
  -> queued original /auth/me is retried once
  -> fallback block retries original /auth/me again
  -> HostAuthGate sees new provider token and may run guard
```

This chain is fully present in current code and no longer protected by the deleted coalescer.

## Chain C: calls begin during session dialog

```text
Initial host guard gets auth failure
  -> showDialog event
  -> HostAuthGate sets sessionDialogPending
  -> HostAuthGate renders children unconditionally
  -> previously unmounted remotes mount during pending host guard
  -> remote mount guard runs concurrently
```

Reference:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:152-161`
- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:272-274`

## Chain D: repeated `/auth/me?accessToken` during backend timeout

```text
Remote invokes checkToken()
  -> /auth/me?accessToken
  -> response/error status 504
  -> checkToken() recursively invokes itself
  -> no limit or delay
```

Reference:

- `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:953-978`

---

# Assessment

The leading structural cause is **not the `qc-bridge-sync` event itself**. It is a remote or second provider invoking auth lifecycle functions, combined with:

- removal of request coalescing in `ab8ae0f0`,
- readiness-dependent subtree unmounting introduced by that commit,
- and retry/remount paths that execute remote mount effects again.

For a network trace labeled as a host/event initiator:

- If it occurs after `showDialog`/relogin, the duplicate retry path is the strongest exact explanation.
- If it tracks MFE remounts or readiness flicker, an unmigrated remote `guard()` plus the current render gate is the strongest explanation.
- If the URL includes `?accessToken=`, investigate remote `checkToken()` immediately.
- `qc-bridge-sync` alone does not call `/me` in this shell; a remote listener not present in this checkout would be required for a direct event-to-network chain.

The working tree remained unmodified and clean.

---

