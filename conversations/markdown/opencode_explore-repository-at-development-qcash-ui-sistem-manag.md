# Trace auth bridge lifecycle (@explore subagent)

- **ID**: `ses_0062ac9cfffe8F9QBxjjpRsoM0`
- **Source Tool**: `opencode`
- **Date**: `2026-08-13 13:35:27`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore the repository at ~/Development/qcash-ui-sistem-management with medium thoroughness. Trace all definitions, writes, reads, and effects involving window.__QCASH_AUTH_BRIDGE__, especially header logout, landing-page login, plafon-pegadaian, and helpdesk portal pages. Identify why the bridge value does not refresh and recommend the smallest correct code change. Do not edit files. Return exact file paths, relevant line ranges, and any existing tests/build commands.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
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
- **grep**: Tool: grep
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read

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
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
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

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

## Finding

`window.__QCASH_AUTH_BRIDGE__` is not defined, written, or read anywhere in the requested repository:

`~/Development/qcash-ui-sistem-management`

That repository only consumes `@bri/addons-auth-provider`. The bridge is owned by the surrounding shell repository, while the header, landing page, Plafon Pegadaian, and Helpdesk Portal are separate federated repositories.

## Bridge lifecycle

### Definition and writes

**Shell writer**

`~/Development/qcash-ui/components/providers/AuthBridgeSync.ts`

- Lines 50–52: bridge event, property name, and internal symbol.
- Lines 77–94: serializes authorities and writes the snapshot with `Object.defineProperty`.
  - The property is `configurable: true`, so `writable: false` does not prevent replacement.
- Lines 101–112: `clearBridgeData()` publishes a guest snapshot, dispatches `qc-bridge-sync`, and removes cached authorization data.
- Lines 143–159: guest snapshot definition.
- Lines 166–193: token synchronization and route/source selection.
- Lines 196–212: creates the normal bridge snapshot.
- Lines 256–300: synchronizes tokens and invokes the Ordent guard.
- Lines 302–381: copies Ordent state into the BRI provider where appropriate.
- Lines 383–452: publishes the selected provider snapshot and dispatches `qc-bridge-sync`.
- Lines 454–462: clears on unmount, but only when no token exists.

The component is mounted globally at:

`~/Development/qcash-ui/pages/_app.tsx:163-189`

Specifically line 171.

**Route selection**

`~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1-15`

- `/helpdesk`
- `/plafon-pegadaian`
- `/systems`

These routes prefer the BRI provider.

### Reads and reactions

**Header bridge adapter**

`~/Development/qcash-ui-header-footer/layouts/default/AuthBridge.tsx:3-97`

- Defines the expected bridge shape and fallback.
- Lines 81–97 read the current window property.
- Lines 86–93 rehydrate serialized authority entries into a `Map`.

`~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:4-31`

- Initializes state from the current bridge.
- Lines 19–27 listen for `qc-bridge-sync`, reread the property, and rerender consumers.
- Lines 8–17 copy a bridged token to local storage, although the shell’s reduced bridge snapshot currently does not include `token`.

**Header layout and menus**

`~/Development/qcash-ui-header-footer/layouts/default/container/index.tsx:30-46`

- Reads bridge state through `useAuthBridge`.
- Lines 34–36 select the bridge whenever the window property is truthy.
- A guest bridge remains truthy, intentionally preventing fallback to a stale local provider.

`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:40-96`

- Uses bridged authorities and readiness.
- Recomputes menu entries when the bridge hook receives a sync event.

**Diagnostics**

`~/Development/qcash-ui/components/federation/monitor/auth/index.tsx:61-113`

- Lines 83–101 read the property.
- Lines 108–113 subscribe to `qc-bridge-sync`.

**Shared authority hook**

`~/Development/qcash-global-component/hooks/use-authorities-ready.ts:4-45`

- Lines 9–12 read the bridge.
- Lines 20–28 retain it in React state and subscribe to `qc-bridge-sync`.
- Lines 30–42 prefer ready bridge authorities over local context.

## Relevant user flow

### Header logout

`~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:43-78`

- Lines 60–74 clear authorization-related local-storage entries and call the Ordent provider’s `logout()`.
- Lines 76–77 choose `/landing-page` or `/main-page`.
- Lines 160–166 invoke logout from the menu.

The installed Ordent provider used by the shell behaves as follows:

`~/Development/qcash-ui/node_modules/@ordentco/addons-auth-provider/dist/src/auth.js:637-671`

It removes tokens and resets token/menu/authority state, but it does **not** reset:

- `username` to `"Guest"`
- `isAuthoritiesReady` to `false`
- company/user identity fields

This is significant because the bridge clear condition currently requires both providers to report guest:

`~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:394-404`

```ts
if (briIsGuest && ordentIsGuest && !hasToken) {
  clearBridgeData();
  ...
}
```

After header logout, `!hasToken` is true, but the Ordent context may still retain the old username. Consequently, this branch can be skipped and the old bridge snapshot remains exposed.

### Landing-page login

`~/Development/qcash-ui-landing-page/pages/landing-page/index.tsx:26-58,107-121`

- Handles SSO login.
- Handles `?logout=true`.
- Calls the Ordent provider logout at lines 111–116.

Normal login writes tokens at:

`~/Development/qcash-ui-landing-page/zustand/useLogin.tsx:572-624`

Provider state is updated at:

- `~/Development/qcash-ui-landing-page/components/newskin/hero-section/hero-login.tsx:99-122,200-224`
- `~/Development/qcash-ui-landing-page/components/newskin/hero-section/index.tsx:235-249`
- `~/Development/qcash-ui-landing-page/components/hero-section/login/login-section.tsx:288-327`

The shell sees the new token and eventually guards/hydrates the providers. However, until hydration completes, consumers can continue seeing the old bridge if logout failed to publish the guest snapshot.

### Plafon Pegadaian

`~/Development/qcash-ui-plafon-pegadaian/pages/plafon-pegadaian/index.tsx:25-81`

- Reads BRI context, not the window bridge directly.
- Lines 29–31 compare the BRI username with `localStorage["login"]`.
- Lines 38–53 remove `session-user-data` and hard reload on mismatch.
- Lines 55–57 redundantly invoke `guard()` a second time on mount.
- Lines 65–80 gate rendering by menu and `PLAFON_PEGADAIAN` authority.
- The reload flag is named `helpdesk_reload_busted`, despite being Plafon code.

Because `/plafon-pegadaian` is a BRI-preferred shell route, stale bridge/provider state is especially visible here as loading or unauthorized rendering.

### Helpdesk Portal

`~/Development/qcash-ui-helpdesk-portal/pages/helpdesk/portal/index.tsx:6-52`

`~/Development/qcash-ui-helpdesk-portal/pages/helpdesk/portal/tickets/index.tsx:6-57`

- Consume BRI context directly.
- Compare username against `localStorage["login"]`.
- Remove `session-user-data` and reload once on mismatch.
- Call `guard()` once through a ref.
- Remain loading while identity is guest/stale/not ready.

Ticket detail differs:

`~/Development/qcash-ui-helpdesk-portal/pages/helpdesk/portal/tickets/view/[userID]/[id].tsx:5-16`

It calls `guard()` but has no readiness or stale-user protection.

## Root cause

The bridge can be replaced correctly; `writable: false` is not the problem because the property is configurable.

The primary failure is the logout invalidation condition in `AuthBridgeSync`:

1. Header logout removes the access token.
2. Ordent’s logout implementation leaves the old username/readiness in context.
3. `AuthBridgeSync` only clears when **both** providers are guest **and** no token exists.
4. Therefore the old window snapshot may survive logout.
5. Landing login changes the token, but consumers can continue reading the stale snapshot until a complete guard/hydration cycle publishes another snapshot.

There is also a secondary refresh gap: the publishing effect at `AuthBridgeSync.ts:432-452` does not explicitly depend on either provider’s `productAuthorities`. A product-authority-only update may not republish unless another listed dependency changes.

## Smallest correct change

In:

`~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:394-404`

clear the bridge whenever the access token is absent:

```ts
if (!hasToken) {
  clearBridgeData();
  lastValidatedTokenRef.current = {};
  return;
}
```

This is preferable to waiting for both provider usernames to become guest because token removal is the authoritative logout signal already used throughout the application.

As a defensive follow-up, add these dependencies to the publishing effect at lines 432–452:

```ts
briAuth?.productAuthorities,
ordentAuth?.productAuthorities,
```

That follow-up is useful for authority-only updates, but the no-token clear is the smallest fix for the observed logout/login stale-session sequence.

## Existing tests

Bridge tests:

- `~/Development/qcash-ui/components/providers/__tests__/AuthBridgeSync.test.tsx:86-245`
  - Publishes BRI/Ordent snapshots.
  - Tests guest clearing at lines 178–204.
  - Tests BRI route selection at lines 206–227.
  - Tests unmount clearing at lines 229–244.
  - Missing regression: token absent while one provider still retains an old non-guest username.

Header tests:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/__tests__/LcImportModal.test.tsx:33-189`

Shared hook tests:

- `~/Development/qcash-global-component/hooks/__tests__/use-authorities-ready.test.ts:12-96`
  - Lines 63–84 verify refresh after `qc-bridge-sync`.

Plafon tests:

- `~/Development/qcash-ui-plafon-pegadaian/test/pages/plafon-pegadaian/index.test.tsx`
  - Includes stale-cache/session removal coverage around lines 142–184.

System-management tests:

- `~/Development/qcash-ui-sistem-management/__tests__/pages/systems/index.test.tsx:118-169`
- Equivalent edit/view tests:
  - `~/Development/qcash-ui-sistem-management/__tests__/pages/systems/edit/[id].test.tsx`
  - `~/Development/qcash-ui-sistem-management/__tests__/pages/systems/view/[id].test.tsx`

## Build/test commands

From each corresponding repository:

- Shell: `yarn test`, `yarn test:coverage`, `yarn build`
- Header/footer: `yarn test`, `yarn test:coverage`, `yarn build`
- Landing page: `yarn test`, `yarn test:coverage`, `yarn build`
- Plafon Pegadaian: `yarn test`, `yarn test:coverage`, `yarn build`
- Helpdesk Portal: `yarn test`, `yarn test:coverage`, `yarn build`
- System management: `yarn test`, `yarn test:coverage`, `yarn build`

No files were modified.

---

