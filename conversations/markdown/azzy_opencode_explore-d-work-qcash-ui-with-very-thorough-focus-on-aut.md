# Inspect qcash auth sync (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_049ba457affe1VEaCmhCi4Cqqq`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 10:44:10`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore D:\work\qcash-ui with very thorough focus on the auth bridge sync to module federated remotes, routing/navigation synchronization, routeChange events, popstate/history interception, and loading/mounting qcash-ui-issuance-lc. Research only, no edits. Explain the exact navigation lifecycle when leaving /v2/import-lc/issuance-lc via header menu. Identify bugs that could let URL change but retain issuance content. Return exact file paths/line refs and minimal fix recommendation.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

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
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
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

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

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

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

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

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 19: Assistant

## Conclusion

A normal header click that reaches the host’s shared `next/router` should unmount Issuance LC and render the destination. The strongest explanation for “URL changes but Issuance LC remains” is that navigation bypassed or failed after updating the browser history while the host router/component state stayed on the issuance page.

There is no shell-level synchronization for native `history.pushState`, `replaceState`, or non-Next history entries.

## Exact lifecycle leaving `/v2/import-lc/issuance-lc`

### 1. Existing Issuance LC mount

1. The shell route is backed by:
   - `D:\work\qcash-ui\pages\import-lc\issuance-lc\index.tsx:10-18`
   - Header/layout remote: `qcash-ui-header-footer/default`
   - Content remote: `qcash-ui-issuance-lc/issuance-lc`

2. `createRemoteDynamic` first loads the remote’s `globals`, then `loadRemote(remote)`:
   - `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:9-24`
   - Outer loading/state lifecycle: `:34-76`
   - Remote render: `:97-100`

3. The host registry maps the whole issuance subtree to that remote:
   - `D:\work\qcash-ui\constants\features\registry.ts:516-520`
   - Prefix route matching: `D:\work\qcash-ui\services\federation\utils\registry.ts:104-120`

4. The exposed issuance component consumes the shared Ordent auth singleton and invokes `guard()` on mount:
   - `D:\work\qcash-ui-issuance-lc\pages\import-lc\issuance-lc\index.tsx:34-47`
   - Authorization/loading gates: `:49-55`
   - Actual issuance content: `:57-78`

5. `next/router` is a Module Federation singleton by default, even though it is not explicitly listed in each project:
   - `D:\work\qcash-ui\node_modules\@module-federation\nextjs-mf\dist\src\internal.js:14-34`
   - Plugin merges that default share scope: `D:\work\qcash-ui\node_modules\@module-federation\nextjs-mf\dist\src\plugins\NextFederationPlugin\index.js:106-135`

### 2. Header menu click

For an ordinary header item:

1. `MenuMapper` invokes the supplied item callback:
   - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:45-59`
   - Old-skin equivalent: `:76-95`

2. `Menu` wires that callback to `onMenuClick`:
   - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\Menu.tsx:8-18`

3. `MenuModalProvider.onMenuClick` checks the bridged leave action, optional special modals, then navigates:
   - Leave confirmation: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:115-139`
   - Modal interception: `:54-84`
   - Navigation: `:86-113`

4. From the current `/v2/...` URL, a target such as `/v2/homepage` is stripped to `/homepage`, then sent to `router.push`:
   - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:88-99`
   - Dashboard supplies `/v2/homepage`: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\utils\defaultHome.ts:3-9`
   - The host router restores its configured `/v2` base path:
     `D:\work\qcash-ui\next.config.js:65-70`

For the **Import LC modal**, the route is handled differently:

- Parent `/v2/import-lc` opens `LcImportModal`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:68-70`
- A different LC subtype calls `router.push(redirectUrl)`:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:97-113`

### 3. Next.js route transition

For a successful shared-router navigation:

1. `router.push` enters `change("pushState", ...)`:
   - `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:431-452`

2. `routeChangeStart` fires before destination resolution:
   - `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:771-789`

3. Both active auth providers listen to this event, but their handlers only update a private previous-path ref; they do not drive rendering or intercept navigation:
   - Ordent: `D:\work\qcash-ui\node_modules\@ordentco\addons-auth-provider\src\auth.tsx:1064-1077`
   - BRI: `D:\work\qcash-ui\node_modules\@bri\addons-auth-provider\src\auth.tsx:1069-1082`

4. After destination information loads, Next emits `beforeHistoryChange`, calls browser `pushState`, and then updates the React/router state:
   - `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:955-967`
   - History state includes `__N: true`: `:996-1020`

5. The host receives a new `Component`; `_app` evaluates the destination’s `getLayout` and renders it:
   - `D:\work\qcash-ui\pages\_app.tsx:158-172`
   - Page/layout render: `:174-208`

6. The old issuance wrappers unmount. Their abort cleanup prevents late remote loading from restoring them:
   - `D:\work\qcash-ui\components\federation\remote-dynamic.tsx:68-76`

7. `routeChangeComplete` fires:
   - `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:979-981`
   - Session-expiration logic reacts here:
     `D:\work\qcash-ui\hooks\use-session-expiration.ts:78-106`

8. The federation provider reacts to the new `router.pathname` and loads destination globals:
   - `D:\work\qcash-ui\components\federation\provider.tsx:143-158`

9. If mounted, `AuthBridgeSync` republishes its snapshot because `router.asPath` is a dependency:
   - `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:382-441`

## Popstate/history behavior

The shell has no custom production `popstate`, `pushState`, or `replaceState` interception. Searches found those only in tests.

Next itself installs the single production `popstate` listener:

- `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:1550-1625`
- Listener registration: `:1693-1716`

Critically:

- A history entry with `state.__N` is routed through Next.
- A raw/non-Next entry without `__N` is ignored:
  `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:1571-1578`
- `pushState` and `replaceState` do not emit `popstate`.

Therefore, a remote doing raw `history.pushState({}, "", target)` will change the address bar while leaving the host’s `router.pathname`, `_app` `Component`, and issuance content untouched.

## Bugs and risk points

### 1. No native-history synchronization — primary stale-content mechanism

There is no host adapter that turns remote-native history changes into `router.push`/`router.replace`. The federation provider and page selection observe only Next router state:

- `D:\work\qcash-ui\components\federation\provider.tsx:69-75`
- `D:\work\qcash-ui\components\federation\provider.tsx:114-125`
- `D:\work\qcash-ui\components\federation\provider.tsx:143-158`

Consequences:

- Raw `pushState` can produce the exact symptom immediately.
- Raw history entries are ignored by Next during back/forward.
- Route-change listeners do not repair the mismatch.

The current checked-in header and issuance sources do not directly call raw history APIs, so this points to another deployed remote/version, an isolated router instance, or external navigation code if the symptom occurs on a regular menu click.

### 2. Header ignores navigation promise/failure

Header calls are not awaited or caught:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:92-98`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:108-113`

Next changes browser history before its React state update:

- `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:955-967`

A rejection/cancellation during the subsequent update can therefore leave a changed URL without a successfully replaced page. The header has no recovery or hard-navigation fallback.

### 3. `LcImportModal` builds a malformed host route

It removes `/v2` from `baseFederation`, constructs an absolute URL without the shell base path, then passes it to `router.push`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:97-113`

For a host configured at `/v2`, the desired argument should be `/import-lc/incoming` or `/import-lc/amend-cancel-lc`; Next will add `/v2`. The current absolute URL lacks `/v2` and is classified as non-local because local absolute URLs must contain the base path:

- `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\utils\is-local-url.js:13-21`

In production this falls into hard navigation:

- `D:\work\qcash-ui\node_modules\next\dist\shared\lib\router\router.js:732-740`

This may hit a legacy route, gateway fallback, or 404 rather than the intended host page.

### 4. Auth bridge contract does not match the header contract

The host intentionally publishes a reduced snapshot:

- Host type: `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:54-73`
- Snapshot construction: `:195-211`
- Test explicitly requires no `products`, `dispatch`, or `onLeaveAction`:
  `D:\work\qcash-ui\components\providers\__tests__\AuthBridgeSync.test.tsx:101-133`

The header expects a much broader contract, including `token`, `userID`, `products`, `onLeaveAction`, and `setOnLeaveAction`:

- `D:\work\qcash-ui-header-footer\layouts\default\AuthBridge.tsx:6-50`
- Header navigation reads leave actions:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:37-39,115-129`

Thus, when the bridge is active, unsaved-change navigation interception cannot work through the advertised header contract. Issuance currently does not set `onLeaveAction`, so this is not the direct current-page retention cause, but it is a real cross-MFE synchronization defect.

### 5. Stale bridge remains when leaving a BRI route while logged in

The cleanup comment says it resets on leaving a BRI route, but cleanup only runs when local storage has no token:

- `D:\work\qcash-ui\components\providers\AuthBridgeSync.ts:443-451`

With an active session, the old window snapshot remains. On a subsequent non-BRI route where `AuthBridgeSync` is not mounted, the header detects that stale object and prefers it:

- `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:33-39`
- Same logic in LC modal:
  `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:40-48`

This can leave stale menu/auth behavior across provider-route transitions.

## Minimal fix recommendation

1. **Fix header navigation first:** always pass host-relative routes to the shared router and await the result:
   ```ts
   await router.push(`/import-lc/${menuValue}`);
   ```
   Remove the `baseFederation`/absolute-URL construction in `LcImportModal`.

2. **Handle failed navigation:** catch rejection or a `false` result and use a correctly base-pathed hard-navigation fallback, rather than silently leaving a URL/content mismatch.

3. **Do not let remotes mutate history directly.** Expose a small shell navigation API/custom event that calls the host `router.push`/`replace`. If legacy remotes cannot be updated immediately, add a narrowly scoped host history bridge that detects raw `pushState`/`replaceState`, while avoiding recursion.

4. **Align the auth bridge contract:** either restore the leave-action capability required by header-footer or remove those fields from the header bridge contract and obtain them from the shared auth context. Clear the bridge unconditionally when its owner unmounts/leaves the bridged provider scope.

No files were edited.

---

