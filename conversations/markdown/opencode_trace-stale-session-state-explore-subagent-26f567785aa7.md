# Trace stale session state (@explore subagent)

- **ID**: `opencode-74d2f47b774703be0384ea9b4adf06605c3ee035a932e325e72026f567785aa7`
- **Source Tool**: `opencode`
- **Date**: `2026-09-28 09:28:44`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only investigation in D:\work\qcash-ui. Thoroughness: very thorough. Trace authentication/session lifecycle for this bug: login as BA, logout, login as CU, but header/footer and menu still show BA. Find login/logout handlers, both auth providers/stores, persistence (cookies/localStorage/sessionStorage), window bridges, federation remotes, layout/header/footer/menu derivation, effects and cache/reset behavior. Identify likely root cause(s), exact files/functions/lines, and propose smallest robust fix plus tests. Do not edit anything. Return evidence-backed findings and note uncertainties.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

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

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
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

# Authentication/session lifecycle investigation

No files were edited. The Git working tree was clean. I did not run tests because this was a read-only investigation.

## Executive conclusion

The strongest root cause is a **split authentication context across the Module Federation boundary**, compounded by incomplete cache/state invalidation:

1. The host shares `@bri/addons-auth-provider` **without `singleton: true`**.
2. The header/footer is a federated remote and consumes auth/global-store state.
3. Therefore, the host and header/footer may resolve different auth-provider factories and different React contexts.
4. Logout can reset the header/footer provider while the host provider remains hydrated as BA.
5. When CU logs in, the host provider’s token changes, but its token-agnostic five-minute in-memory freshness check can return without loading CU.
6. `HostAuthGate` then records the CU token as “validated” while still publishing BA identity/menu state.

The stale `session-user-data`, menu local-storage values, bridge snapshot, and global store are additional propagation paths.

---

# Lifecycle trace

## 1. Provider ownership

The current host mounts only the BRI provider:

- `D:\work\qcash-ui\pages\_app.tsx:177-202`
- `D:\work\qcash-ui\components\providers\BRIProviders.tsx:14-22`

Provider hierarchy:

```text
ModuleFederationProvider
  BriProviders
    BRI AuthProvider
      BRI GlobalStoreProvider
        HostAuthGate
          Header/footer remote and page remotes
```

The legacy Ordent host provider was removed in commit `6970751d`. Current source contains no active `@ordentco/addons-auth-provider` import.

However, this does not prove that loaded remotes use the same BRI context. A remote can bring:

- another BRI provider factory/version, or
- a legacy Ordent provider.

Historical repository audits found that header/footer had both packages during migration. That evidence is historical, not a verification of the currently deployed remote.

## 2. Federation loading

Header/footer is registered as a global remote:

- `D:\work\qcash-ui\constants\features\registry.ts:37-52`
- entry URL at `D:\work\qcash-ui\constants\features\registry.ts:44-47`

Federation initializes all remotes concurrently:

- `D:\work\qcash-ui\services\federation\init.ts:32-68`
- global remote styles are preloaded at `D:\work\qcash-ui\services\federation\init.ts:72-84`

Almost every protected page then loads:

```ts
loadRemote("qcash-ui-header-footer/default")
```

Example:

- `D:\work\qcash-ui\pages\account-summary\index.tsx:14-25`
- `D:\work\qcash-ui\pages\index.tsx:11-15,32-36`
- shared wrapper: `D:\work\qcash-ui\components\ui\SecureFeatureWrapper.tsx:7-10,48-59`

## 3. Critical federation configuration defect

The host currently declares:

- `D:\work\qcash-ui\next.config.js:15-17`

```js
"@bri/addons-auth-provider": { requiredVersion: false },
```

It does **not** declare `singleton: true`.

That contradicts the repository’s own architecture contract:

- `D:\work\qcash-ui\docs\centralized-auth-guard-migration.md:19-24`
- `D:\work\qcash-ui\docs\ordent-to-bri-provider-migration.md:19-28,32-38`

Both documents require BRI auth to be a Module Federation singleton.

Git history strengthens this finding:

- Commit `4ea23672` explicitly changed BRI and Ordent to singletons.
- The current line is blamed to the older non-singleton configuration.
- The present branch has therefore lost the host-side singleton setting while the migration documentation still depends on it.

Without a singleton, the header/footer’s `useAuth()` and `useGlobalStore()` are not guaranteed to refer to the contexts mounted by `BriProviders`.

## 4. Login handling

The primary login UI is remote, so its implementation is not in this repository:

- landing login remote: `D:\work\qcash-ui\pages\landing-page\index.tsx:9-13`
- legacy main-page remote: `D:\work\qcash-ui\pages\main-page\index.tsx:15-25`

The host’s local login service is:

- `D:\work\qcash-ui\services\auth.ts:14-40`

It is used by the session-expired relogin flow:

- `D:\work\qcash-ui\pages\_app.tsx:188-194`
- `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:81-176`

That flow:

1. Stores new access/refresh tokens:
   - `use-modal-session-expired.tsx:97-105`
   - alternate flow: `144-150`
2. Removes `session-user-data`.
3. Calls BRI `setToken`.
4. Dispatches `qc-bridge-sync` with `source: "token-change"`:
   - `113`
   - `156`
5. Reloads only for selected conditions:
   - `118-128`
   - `160-169`

The installed provider’s own login methods only set new tokens and navigate. They do not reset prior identity, authority, menu, or freshness state:

- password login:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:577-603`
- SSO login:  
  `...\auth.js:620-639`
- login:  
  `...\auth.js:912-940`

## 5. Host validation

`HostAuthGate` reads the access token from local storage:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:35`

It first synchronizes it to the host BRI context:

- `HostAuthGate.tsx:37-44`
- invocation: `180-187`

It then invokes `briAuth.guard()`:

- `HostAuthGate.tsx:237-262`

After guard completion it:

1. Publishes the auth bridge.
2. Stores `{ token, status: "validated" }`.
3. Allows protected children to render.

Relevant lines:

- `HostAuthGate.tsx:239-246`
- final gate: `290-302`

## 6. Provider cache behavior

The installed dependency is:

- `D:\work\qcash-ui\package.json:20`
- `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\package.json:1-8`

### Session-storage cache

`guard(true)` reads `sessionStorage["session-user-data"]` and restores the whole identity if it is less than five minutes old:

- `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:283-318`

The cache includes:

- username
- user type
- company
- role IDs
- menus/menu data
- authorities
- product authorities

It does **not** include a token identity:

- cache creation: `...\auth.js:458-482`

Thus it cannot verify that cached BA data belongs to the current token.

The host currently invokes `guard()` without `true`, so this session-storage fast path is not normally used by `HostAuthGate`. It remains dangerous if a remote calls `guard(true)`.

### In-memory freshness cache

More importantly, the provider always checks `sessionLastValidatedAt`, even when `useCache` is false:

- `...\auth.js:319-326`

```js
if (
  sessionLastValidatedAt &&
  Date.now() - sessionLastValidatedAt < SESSION_VALIDITY_MS
) {
  if (!isAuthoritiesReady) setIsAuthoritiesReady(true);
  return;
}
```

This cache is also not tied to the token or user.

Consequently:

1. Host provider validates BA.
2. A separate header/footer provider performs logout.
3. Host provider’s `sessionLastValidatedAt`, username and authorities remain BA.
4. CU token is stored.
5. `HostAuthGate` copies CU token into the host provider.
6. Host `guard()` returns at the freshness check.
7. `HostAuthGate` marks the CU token validated while host state is still BA.

This precisely matches “token/login changed, but header/footer and menu still show BA.”

## 7. Logout handling

Two local logout callers were found:

- session-expired sign-out:  
  `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx:178-190`
- MFA error sign-out:  
  `D:\work\qcash-ui\components\ui\MFAErrorModal.tsx:67-83`

Both call provider `logout()`.

The installed provider logout:

- `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\auth.js:650-685`

It clears:

- access token
- refresh token
- login
- `productMenu`
- `validateMenu`
- cookies
- all session storage
- token
- validation timestamp
- menus/menu data
- authorities

But it does **not** reset:

- `username`
- `userType`
- company fields
- role fields
- `productAuthorities`
- `isAuthoritiesReady`
- the separate BRI global store
- `localStorage["productRoles"]`

Therefore, even a correctly shared provider retains part of BA state until a successful new guard overwrites it.

If logout is called on an isolated remote provider, even the fields it does reset are reset only in that provider instance.

## 8. Missing logout event

`HostAuthGate` listens for:

```ts
qc-bridge-sync, detail.source === "logout"
```

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:123-143`

But no current application code dispatches that event with `source: "logout"`.

Found producers are only:

- `"host-auth-gate"`: `HostAuthGate.tsx:83`
- `"token-change"`:
  - `use-modal-session-expired.tsx:113,156`
  - `components/onboarding-tour/notif-onboarding-tour.tsx:82-86`

The installed provider logout also does not dispatch it.

As a result, the intended reset of:

- `window.__QCASH_HOST_AUTH_GUARD__`
- `isValidated`
- `hasMountedRef`
- auth error/session-dialog state

is normally not triggered by logout.

## 9. Stale bridge behavior

The bridge contains identity, menus and authorities:

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:46-84`

No logout path clears it.

There is also a stale-snapshot race at:

- `HostAuthGate.tsx:239-245`

`currentBriAuth` was captured before `await guard()`. Immediately after guard completes, the code publishes that captured object. Normally a subsequent context rerender corrects it. But when guard exits through the in-memory freshness check, there is no identity state change and the BA snapshot can remain authoritative.

The host’s auth monitor reads this bridge:

- `D:\work\qcash-ui\components\federation\monitor\auth\index.tsx:78-108`

Any passive remote bridge consumer can therefore continue receiving BA.

## 10. Global-store lifecycle

The BRI global store is a separate React context:

- provider mount: `D:\work\qcash-ui\components\providers\BRIProviders.tsx:19-21`
- implementation:  
  `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\dist\src\global-store\index.js:61-66`

It loads products/system configuration whenever token truthiness changes:

- `...\global-store\index.js:183-194`

But when the token becomes false, it does not dispatch `CLEAN_DATA`. The reducer supports that action:

- `...\global-store\reducer.js:17-18,54-55`

No caller in this repository uses it.

Thus BA `products`, `bricamsUser`, `systemConfig`, errors, and dynamic state can survive logout until CU requests replace them. If a CU request fails, stale BA store values can remain visible.

Historical header/footer migration evidence says its menu/layout consumes `useGlobalStore`, so this store persistence is relevant, although the current deployed remote source was not available inside the requested repository scope.

## 11. Browser persistence

### Local storage

Authentication-related keys found:

- `access-token`
- `refresh-token`
- `login`
- `productMenu`
- `productRoles`
- `validateMenu`
- `agent`

The host landing-page cleanup removes:

- `productMenu`
- `productRoles`
- `validateMenu`
- `session-user-data`

at:

- `D:\work\qcash-ui\pages\_app.tsx:101-111`

But this applies only to `/landing-page` variants, not `/main-page`.

Local logout wrappers omit `productRoles`:

- `use-modal-session-expired.tsx:186`
- `MFAErrorModal.tsx:79-81`

### Session storage

The important key is `session-user-data`. It is:

- read by provider guard: `auth.js:283`
- written by guard: `auth.js:482`
- cleared by provider logout through `sessionStorage.clear()`: `auth.js:674`

### Cookies

The provider sets `loggedIn=true` at login:

- `auth.js:924-930`

and tries to expire it on logout:

- `auth.js:672-673`

The dashboard SSR check trusts the cookie:

- `D:\work\qcash-ui\pages\index.tsx:23-28`

Cookie writes do not specify `Path`, `SameSite`, or `Secure`. A same-name cookie on another path could survive. This can cause routing/session confusion, but it is less likely to be the direct cause of BA identity appearing in the header.

## 12. Existing hard-reload workaround

The application tries to destroy stale singleton state after logout:

- `D:\work\qcash-ui\pages\_app.tsx:94-99`

```ts
if (router.query["logout"] === "true") {
  window.location.replace(window.location.pathname);
}
```

It also clears selected caches on landing:

- `_app.tsx:101-111`

This explains why some environments may not reproduce the problem: a hard reload destroys all provider instances, refs and in-memory request caches.

It is not a robust lifecycle contract because:

- it depends on every logout reaching a route with `?logout=true`;
- cache cleanup excludes `/main-page`;
- it does not address token changes/relogin without that redirect;
- it masks rather than fixes split contexts and token-agnostic caching.

---

# Ranked root causes

## 1. Host auth provider is not a federation singleton — highest confidence

**Evidence:**

- `D:\work\qcash-ui\next.config.js:16`
- federated header/footer loads throughout the application.
- repository documentation explicitly requires singleton sharing.
- Git history shows singleton was deliberately added previously.

This permits separate auth and global-store contexts, making it possible for the header/footer to remain BA while the host token becomes CU.

## 2. Provider freshness cache is not bound to the token — high confidence

**Evidence:**

- `...\node_modules\@bri\addons-auth-provider\dist\src\auth.js:319-326`
- cached session object also lacks token identity at `460-482`.

If logout operates on a different provider instance, the host provider returns early with BA state for up to five minutes.

## 3. Logout does not fully reset state — high confidence

**Evidence:**

- provider logout only resets a subset at `auth.js:667-681`;
- no reset for identity, product authorities, readiness, or global store;
- no `source: "logout"` event exists in current application code.

## 4. Header/footer menu caches can independently retain BA — medium confidence

Historical source analysis found header/footer had:

- one-time menu-fetch refs;
- static request cache keys;
- local-storage fallbacks using `productMenu`/`validateMenu`;
- identity derived directly from its auth context.

This strongly fits the symptom, but the current header/footer remote source/deployment was outside this repository and was not revalidated here.

## 5. Bridge publication can preserve a pre-guard snapshot — medium confidence

- `D:\work\qcash-ui\components\providers\HostAuthGate.tsx:239-245`

This becomes significant when guard returns without producing context updates.

---

# Smallest robust fix

## Required first fix

Restore strict singleton sharing in the host:

```js
"@bri/addons-auth-provider": {
  singleton: true,
  requiredVersion: false,
},
```

at:

- `D:\work\qcash-ui\next.config.js:16`

The same exact share contract must be present in `qcash-ui-header-footer` and every auth-consuming remote. A host-only change is necessary but not sufficient if a remote bundles or initializes a private provider.

## Provider-level correctness fix

In `@bri/addons-auth-provider`:

1. Associate freshness with the exact active token/session:
   - keep `lastValidatedToken`;
   - only use `sessionLastValidatedAt` when `lastValidatedToken === currentToken`.
2. Include a non-secret token/session fingerprint in `session-user-data`, and reject the cache when it does not match.
3. Before every successful login token assignment:
   - remove `session-user-data`;
   - reset validation timestamp/token;
   - clear menu storage;
   - reset identity, menus, authorities, product authorities and readiness.
4. On logout, reset every auth field and dispatch `CLEAN_DATA` to the global store.

This removes the five-minute cross-user window even if logout is imperfect.

## Host safety belt

Create one host session-reset path and use it from all host logout flows:

- clear `session-user-data`;
- clear `productMenu`, `productRoles`, and `validateMenu`;
- replace `window.__QCASH_AUTH_BRIDGE__` with an explicit guest/not-ready snapshot;
- clear `window.__QCASH_HOST_AUTH_GUARD__`;
- dispatch `qc-bridge-sync` with `{ source: "logout" }`;
- then invoke provider logout.

The normal header/footer logout must call the same contract or the provider itself must emit this event.

Do not rely solely on the current hard reload.

---

# Test recommendations

## HostAuthGate tests

Extend:

- `D:\work\qcash-ui\components\providers\__tests__\HostAuthGate.test.tsx`

Add:

1. **BA → logout → CU in one mounted tree**
   - initial BA identity/menus/bridge;
   - dispatch logout;
   - replace storage token with CU;
   - assert children remain gated until CU identity is ready;
   - assert bridge never reports CU token validation with BA identity.

2. **Token change with stale readiness**
   - begin with `isAuthoritiesReady=true` and BA fields;
   - change token;
   - guard initially returns without mutating identity;
   - verify the gate must not mark the new token validated.

3. **Logout event**
   - assert `__QCASH_HOST_AUTH_GUARD__` becomes idle;
   - assert bridge is cleared;
   - assert a subsequent token invokes guard again.

4. **Stale async guard**
   - existing test checks only the guard token;
   - additionally assert stale guard A cannot overwrite bridge/identity after token B.

Current gap: `HostAuthGate.test.tsx:120-137` checks only that guard is called twice; it does not verify which user is rendered or bridged.

## Logout tests

Extend:

- `D:\work\qcash-ui\hooks\__tests__\use-modal-session-expired.test.tsx:172-196`
- `D:\work\qcash-ui\components\ui\__tests__\MFAErrorModal.test.tsx:92-114`

Assert:

- `productRoles` is cleared;
- `session-user-data` is cleared;
- logout event source is `"logout"`;
- host guard and bridge are reset even if `logoutLog()` fails.

## Provider tests

In the auth-provider repository:

1. Validate BA with token A.
2. Set token B within five minutes.
3. Call `guard()` and `guard(true)`.
4. Assert `/auth/me` and `/menu/me` run for B.
5. Assert no BA field is restored.
6. Assert logout returns the complete context to initial state.
7. Assert global store receives `CLEAN_DATA`.

## Federation contract test

Add a configuration test or CI script asserting that the host and auth-consuming remotes declare:

```js
"@bri/addons-auth-provider": {
  singleton: true,
  requiredVersion: false
}
```

An end-to-end test should use the actual header/footer remote and verify:

```text
login BA → capture header/menu
logout without closing tab
login CU within five minutes
header, footer, profile, menu, authorities and bridge all show CU
```

Also run the scenario with hard reload suppression so lifecycle correctness is tested rather than the workaround.

---

# Uncertainties

1. The actual normal login and header/footer logout handlers are inside federated remotes, not this repository. Their currently deployed source and version could not be proven from `D:\work\qcash-ui`.
2. Historical vault evidence shows header/footer migration, dual-provider configurations, and menu caches, but that is historical context rather than proof of the currently deployed build.
3. Runtime share resolution depends on loaded remote manifests and order. The existing auth monitor can inspect `window.__FEDERATION__`, but no live browser runtime was attached during this investigation.
4. The package identifies itself as `0.10.7-mfe`, while its code writes `qcash_auth_provider_version = "0.10.6-mfe-patched-local"` at `auth.js:270-275`. This suggests the artifact contains patched older code and should be verified with the provider maintainers.

---

