# please do research at when access-token gone ?

- **ID**: `codex-59506eeb67b03cef82a056dc9596674e473f6f61e8341def365d49b1ef5a5d7c`
- **Source Tool**: `codex`
- **Date**: `2026-09-08 09:55:26`
- **Tags**: `coding`, `assistant`, `codex`, `gpt-5`

---

## Turn 1: User

please do research at when access-token gone ?

---

## Turn 2: Assistant

Which system or app do you mean? Are you asking what happens when an OAuth access token **expires, is revoked, or is deleted**—and how the app should recover?

---

## Turn 3: User

is expired makes it gone ?

---

## Turn 4: Assistant

No. An expired access token usually still exists in the app’s storage, but the server no longer accepts it.

Typical flow:

1. The access token expires.
2. An API request returns `401 Unauthorized`.
3. The app uses a refresh token to obtain a new access token.
4. If refreshing fails, the app deletes the old credentials and asks the user to sign in again.

So **expired ≠ deleted**. Expiration makes the token unusable; your application should remove or replace it.

---

## Turn 5: User

can access-token got deleted somehow by user ?

---

## Turn 6: Assistant

Yes. A user can indirectly or directly remove an access token in several ways:

- Logging out—the app typically deletes its locally stored token.
- Clearing browser cookies, site data, or app storage.
- Uninstalling the app or clearing its data.
- Revoking the app’s authorization through the account/security settings.
- Changing a password, if the provider invalidates existing sessions.
- Removing a connected device or session.

Important distinction: deleting a token from the device only removes that local copy. It may remain valid elsewhere until it expires or is revoked. For security-sensitive logout, delete the local token **and** call the provider’s token-revocation endpoint when available.

---

## Turn 7: User

please do research can it makes no undefined blank like that

---

## Turn 8: Assistant

I’ll verify the distinction between a deleted/missing token, an expired token, and why code may display `undefined` or a blank value.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 10: Assistant

Yes—if the user logs out, clears browser/app data, or your code removes the token, the application can end up showing a blank, `null`, or `undefined` value.

The exact result depends on how the token is read:

```js
localStorage.removeItem("access_token");

localStorage.getItem("access_token"); // null
localStorage.access_token;            // usually undefined
```

`removeItem()` itself also returns `undefined`; it does not return the removed token. This is standard Web Storage behavior. [MDN documentation](https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem)

Prevent malformed requests such as `Authorization: Bearer undefined`:

```js
const token = localStorage.getItem("access_token");

if (!token) {
  // Treat as logged out
  redirectToLogin();
  return;
}

fetch("/api/data", {
  headers: {
    Authorization: `Bearer ${token}`,
  },
});
```

A missing local token means the application cannot authenticate that request. A revoked or expired token might still exist as text locally, but the server should reject it—normally with `401 Unauthorized`. [OAuth Bearer Token standard](https://www.rfc-editor.org/info/rfc6750) Token revocation is separately defined by [RFC 7009](https://www.rfc-editor.org/info/rfc7009/).

Therefore, a suddenly blank token is possible, but it usually indicates local deletion, cleared storage, a different browser/origin, or an application state problem—not normal token expiration by itself.

---

## Turn 11: User

no i mean while the access-token got deleted

---

## Turn 12: Assistant

Yes. **While the access token is being deleted**, another part of the app may read it immediately afterward and receive no value:

```js
localStorage.removeItem("access_token");
const token = localStorage.getItem("access_token"); // null
```

It should not become a blank string or `undefined` automatically:

- `getItem()` returns `null` when the token is missing.
- `localStorage.access_token` may return `undefined`.
- `removeItem()` returns `undefined`.
- A blank string (`""`) appears only if code explicitly stores an empty value, such as `setItem("access_token", "")`.

If the UI shows `undefined` or blank during deletion, it is probably rendering the missing value:

```jsx
<span>{token}</span>             // null/undefined appears blank
<input value={token ?? ""} />    // deliberately becomes blank
```

This is expected during logout. Delete the token, clear authentication state, and redirect or render the signed-out screen without making another authenticated request.

---

## Turn 13: User

please check the current code of qcash-ui. cant it handled by the new changes ?

---

## Turn 14: Assistant

I’ll inspect `qcash-ui`’s current auth/token handling and the latest changes, then trace what happens if the access-token disappears while the app is running.

---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 17: Assistant

I found the relevant recent commit, “fixing blank screen on token deleted,” plus an uncommitted auth-guard interceptor. I’m comparing both with the active request flow so I can tell whether the new change fully covers deletion or only one code path.

---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 21: Assistant

The new commit does address the original blank-screen sequence: a `showDialog` event now remains actionable even after the token has already been removed, and the gate keeps the previously mounted page alive for that relogin modal. I’m running the focused tests and checking edge cases where deletion occurs without that event.

---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 24: Assistant

The new changes handle the specific **401 → token deleted → `showDialog` emitted** flow, preventing the blank screen:

- The modal now opens even when `access-token` is already missing: [use-session-expiration.ts](~/Development/qcash-ui/hooks/use-session-expiration.ts:27)
- `HostAuthGate` keeps the previously mounted protected content alive while the session dialog is pending: [HostAuthGate.tsx](~/Development/qcash-ui/components/providers/HostAuthGate.tsx:265)
- After the page has successfully mounted once, a missing token no longer makes the gate return `null`: [HostAuthGate.tsx](~/Development/qcash-ui/components/providers/HostAuthGate.tsx:277)

However, it does **not fully handle every token-deletion scenario**:

- Manual deletion from DevTools does not emit `showDialog`.
- Same-tab `localStorage.removeItem()` does not trigger a `storage` event.
- The previous storage listener was removed.
- After direct deletion, protected content can remain visible until an API request triggers the session-expiration or redirect flow.
- On a fresh page load without a token, it correctly redirects to `/landing-page`.
- The uncommitted [host-auth-guard-interceptor.ts](~/Development/qcash-ui/components/providers/host-auth-guard-interceptor.ts:83) is not imported or installed anywhere, so it currently has no effect.

The focused test suite passes: **50/50 tests**. But there is no integrated test covering “successfully mounted page → token removed directly → expected recovery.”

So the answer is: **yes, the new change fixes the reported blank-screen path, but no, it does not provide immediate detection for arbitrary token deletion.**

---

## Turn 25: User

then should add and event listener ?

---

## Turn 26: Assistant

Yes—but a `storage` listener alone is insufficient.

It only fires when another tab/window changes `localStorage`; it does **not** fire in the same tab that deletes the token. The robust approach is:

- Listen for `storage` to catch deletion from another tab.
- Dispatch a custom auth event whenever application code removes the token in the current tab.
- Have `HostAuthGate` handle both events through one function.
- Redirect to the landing page for explicit logout.
- Preserve the mounted page when deletion belongs to the session-expired relogin flow, avoiding the original blank screen.

The existing `HOST_AUTH_EVENT` is suitable for same-tab notification. For example:

```ts
localStorage.removeItem("access-token");

window.dispatchEvent(
  new CustomEvent(HOST_AUTH_EVENT, {
    detail: { source: "token-removed", reason: "session-expired" },
  })
);
```

Then `HostAuthGate` can listen for:

```ts
window.addEventListener(HOST_AUTH_EVENT, handleAuthEvent);
window.addEventListener("storage", handleStorage);
```

The important part is distinguishing `logout` from `session-expired`. Immediately resetting the gate for every deletion could reintroduce the blank screen while the relogin modal is opening.

---

## Turn 27: User

please audit the robust implementation

---

## Turn 28: Assistant

I’ll audit the complete token-removal lifecycle: every write/removal site, same-tab and cross-tab signaling, `HostAuthGate` state transitions, session-expired modal behavior, and existing test coverage. I’ll report concrete failure cases and recommended changes without modifying the code.

---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 32: Assistant

A key issue is emerging: `qc-bridge-sync` currently mixes three meanings—auth data refreshed, token changed, and logout—while deletion paths do not consistently emit it. I’m checking the installed auth provider because its `logout()` and 401 interceptor may remove the token outside this repository’s visible helpers.

---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 35: Assistant

The current implementation fixes the reported modal-related blank screen, but it is not yet a robust general token-lifecycle implementation.

## Audit findings

### High: arbitrary token deletion is not observed

`HostAuthGate` no longer listens to the browser `storage` event. It only reacts to `qc-bridge-sync` when `detail.source === "logout"`: [HostAuthGate.tsx](~/Development/qcash-ui/components/providers/HostAuthGate.tsx:131).

Consequences:

- Deletion in another tab is not handled.
- Manual DevTools deletion is not handled immediately.
- Same-tab code that removes the token without updating `briAuth.token` may leave the application authenticated in memory.
- The protected page can remain visible until some later render or API request.

Severity: high for lifecycle consistency, although API authorization still depends on the server.

### High: token removal is not centralized

There are several independent removal paths:

- [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:83)
- [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:125)
- The installed auth provider’s `logout()`
- The installed auth provider’s refresh-token failure
- Its Axios interceptors

These paths do not consistently emit a same-tab event. Some redirect immediately, some update React auth state, and some only remove storage.

A listener cannot compensate for this because `storage` does not fire in the tab that performed `removeItem()`.

### Medium: `qc-bridge-sync` has too many responsibilities

`HOST_AUTH_EVENT` is defined as `qc-bridge-sync`: [host-auth-events.ts](~/Development/qcash-ui/components/providers/host-auth-events.ts:1).

That event currently means several different things:

- Bridge snapshot updated
- Access token refreshed
- Onboarding refresh completed
- Potential logout
- Cache invalidation

This makes it difficult to determine whether the application should:

- Revalidate the new token
- Open the relogin modal
- Sign out and redirect
- Merely refresh bridge consumers

A dedicated event such as `qcash:auth-token-change` would be safer.

### Medium: mounted protected content intentionally remains visible

After successful authentication, `hasMountedRef.current` stays `true`. Therefore a missing token does not unmount the protected application: [HostAuthGate.tsx](~/Development/qcash-ui/components/providers/HostAuthGate.tsx:273).

This prevents the original blank screen while relogin is active, but it also means direct token deletion can leave stale protected information displayed indefinitely if no modal or redirect follows.

The “keep mounted” behavior should only apply when the state is explicitly `session-expired/relogin-pending`, rather than whenever the page mounted previously.

### Medium: session-dialog state never explicitly returns to false

`sessionDialogPending` becomes `true` on `showDialog`: [HostAuthGate.tsx](~/Development/qcash-ui/components/providers/HostAuthGate.tsx:145).

It is cleared during successful or failed guard validation and logout reset, but not directly when the dialog closes or relogin completes. Successful token replacement usually triggers other state changes, so it may work incidentally, but the lifecycle is implicit.

A `relogin-complete` or `dialog-closed` event should clear it deterministically.

### Medium: behavior varies by feature flags

In [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:110), generic `401/499/408` errors dispatch `showDialog` only when `toggleUpdateLoginService` is false.

That flag is true in development/prestage but false in later environments. Therefore the local behavior being tested is not necessarily the production behavior.

### High if installed: the uncommitted interceptor is unsafe

The untracked [host-auth-guard-interceptor.ts](~/Development/qcash-ui/components/providers/host-auth-guard-interceptor.ts:21) is currently unused, so it has no runtime effect.

It should not be installed as-is:

- It monkey-patches the Axios `request` method.
- Its cache key reads `config.headers.Authorization` before Axios request interceptors necessarily inject the token.
- Requests from different tokens could consequently share the same cache key.
- Successful `/auth/me` or `/menu/me` responses could be reused for another session for up to five seconds.
- Its anonymous event listeners cannot be removed during teardown.

That creates a potential cross-session authorization/menu data leak.

## Recommended robust design

Use an explicit auth lifecycle with a shared helper:

```ts
type AuthTokenChangeDetail =
  | { action: "removed"; reason: "logout" | "expired" | "revoked" }
  | { action: "replaced"; token: string };

const AUTH_TOKEN_CHANGE_EVENT = "qcash:auth-token-change";
```

All same-tab writes should go through functions such as:

```ts
export function removeAccessToken(reason: "logout" | "expired" | "revoked") {
  localStorage.removeItem("access-token");

  window.dispatchEvent(
    new CustomEvent(AUTH_TOKEN_CHANGE_EVENT, {
      detail: { action: "removed", reason },
    })
  );
}
```

`HostAuthGate` should then:

- Listen to `AUTH_TOKEN_CHANGE_EVENT` for same-tab changes.
- Listen to `storage` for other-tab changes.
- On `logout` or `revoked`: clear auth context, reset guard state, and redirect.
- On `expired`: preserve children only while the relogin modal is active.
- On token replacement: update provider state, reset the old guard result, and validate the new token.
- On relogin cancellation: redirect instead of leaving stale content visible.
- Remove every listener during cleanup.

## Test coverage needed

The present focused suite passes—50 tests—but these scenarios are missing:

1. Successfully mounted page → same-tab token deletion → logout redirect.
2. Successfully mounted page → cross-tab storage deletion → logout redirect.
3. Expired token removed → `showDialog` → protected content and modal remain mounted.
4. Relogin succeeds → new token is validated.
5. Relogin is cancelled → protected content is removed and user is redirected.
6. Old validation resolves after deletion → result is ignored.
7. Token A replaced by token B → no cached identity/menu response from token A.
8. Event listeners are removed on unmount.

Verdict: **add both listeners, but only together with centralized token mutations and explicit reasons.** Adding a raw `storage` listener that resets the gate on every change could reintroduce the blank screen the latest commit fixed.

---

## Turn 36: User

why the before changes just work fine on local run ?

---

## Turn 37: Assistant

The old behavior can work locally because the local execution path is different from the affected deployed environment.

Most likely reasons:

- Feature flags differ by environment. In [config.ts](~/Development/qcash-ui/config.ts:606):

  - `toggleRedirectLandingPage`: local/dev `true`, production `false`
  - `toggleNewModalRelog`: local/dev `true`, production `false`
  - `toggleUpdateLoginService`: local/dev `true`, production `false`

  These flags change whether a `401` redirects, opens `showDialog`, or enters the legacy relogin flow.

- Local tokens are normally replaced, not removed. During development you likely log in normally, so the auth provider calls `setToken()` and React rerenders. The defect appears when an interceptor removes `access-token` before the dialog state is established.

- The old `storage` listener does not detect deletion in the same tab. It only reacts when another browser tab modifies storage. Therefore testing token deletion from another tab can behave differently from deletion by the active tab’s Axios interceptor.

- Different Axios implementations are involved. The shell has [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:68), while the installed BRI provider also has its own interceptors. Local pages may use one client, while a deployed remote MFE can use another.

- The URL matters. On `/homepage` or `/`, a `401` immediately deletes tokens and redirects: [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:83). On other protected routes, the legacy configuration dispatches `showDialog`: [custom-axios.ts](~/Development/qcash-ui/services/custom-axios.ts:110).

The original race was approximately:

```text
API returns 401
→ interceptor deletes access-token
→ React rerenders HostAuthGate
→ HostAuthGate sees no token and returns null
→ session modal is inside HostAuthGate, so it cannot render
→ blank screen
```

Locally, an immediate redirect, different flag branch, token still being present, or the modal event winning the race would hide the defect.

So “works locally” does not establish that the old implementation was safe; it means the exact environment, route, Axios client, and event ordering needed to trigger the race were probably absent. The best reproduction is to run with the affected environment’s three feature-flag values and simulate the actual MFE’s `401` response—not merely remove the token manually.

---

## Turn 38: User

try to implement it

---

## Turn 39: Assistant

I’ll implement a dedicated token-change event, wire same-tab writes/removals through it, restore cross-tab detection in `HostAuthGate`, and preserve the session-expired modal path so the blank-screen race does not return. I’ll also add focused tests for replacement, expiration, and cross-tab deletion.

---

## Turn 40: User

please find the logout function on qcash product

---

## Turn 41: Assistant

I’ll trace every logout implementation and call site in `qcash-ui`, including the installed BRI auth provider, and identify which one the product UI actually uses.

---

## Turn 42: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 43: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 44: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 45: Assistant

The main shell only owns session-expired and MFA sign-out. The ordinary product-header logout comes from the `qcash-ui-header-footer` remote, and both its source repositories are present locally, so I’m tracing that implementation next.

---

## Turn 46: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 47: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 48: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **exec**: Executed tool exec

</details>



---

## Turn 49: Assistant

The normal QCash product logout button is implemented in the `qcash-ui-header-footer` remote.

Flow:

1. The header logout button calls `_logout(destination)` in [ListItem.tsx](~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:160).
2. `_logout()`:

   - Disables the button.
   - Calls `logoutLog()` without waiting for it.
   - Waits 500 ms.
   - Clears menu-related cache.
   - Calls `logout(href)` from `useAuth()`.

   See [ListItem.tsx](~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:60).

3. `logoutLog()` only records the logout through `GET auth/logoutlog`: [auth.ts](~/Development/qcash-ui-header-footer/services/auth.ts:193).

4. The actual logout is the BRI provider’s `useAuth().logout()`. It:

   - Removes `access-token`
   - Removes `refresh-token`
   - Removes login/menu storage
   - Clears authentication cookies
   - Clears all `sessionStorage`
   - Navigates to the supplied path
   - Sets the provider token to `null`
   - Clears menus and authorities

   See the installed provider implementation: [auth.js](~/Development/qcash-ui/node_modules/@bri/addons-auth-provider/dist/src/auth.js:650).

Important findings:

- The active header-footer `dev` branch does not emit a logout event.
- It relies on the provider’s React state update and navigation after deleting the token.
- The destination is selected using `sessionStorage["new-qlola-flag"]`:

  ```ts
  const destination =
    newQlola === "true" ? "/landing-page" : "/main-page";
  ```

- It does not add `?logout=true`. Therefore the singleton-reset logic in `qcash-ui/pages/_app.tsx` is not triggered for this header logout.
- A separate commit named `fix: auth bridge reset while logout` exists on branch `rxn-bridge-fix`, but it has not been merged into the current `dev` branch. That version adds `resetAuthBridge()` immediately before `logout(href)`.

There are two other sign-out paths inside the main shell:

- Session-expired modal: [use-modal-session-expired.tsx](~/Development/qcash-ui/hooks/use-modal-session-expired.tsx:181)
- MFA error modal: [MFAErrorModal.tsx](~/Development/qcash-ui/components/ui/MFAErrorModal.tsx:69)

Both call `logoutLog()` and then the same provider `logout()` function. The ordinary product logout, however, comes from `qcash-ui-header-footer`.

---

## Turn 50: User

please can you edit the logout instead of using router.push it should use router.replace throught the dist node_moduels/@bri

---

## Turn 51: Assistant

I’ll update the installed `@bri/addons-auth-provider` distribution so logout navigation uses `router.replace()` instead of `router.push()`, limiting the change to logout/forced-logout paths and then verifying every affected occurrence. Note that this is a local `node_modules` patch and can be overwritten by the next install unless it is persisted upstream or with a patch mechanism.

---

