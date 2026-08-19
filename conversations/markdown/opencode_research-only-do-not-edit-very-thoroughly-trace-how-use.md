# Trace issuance navigation (@explore subagent)

- **ID**: `ses_016acd7eeffezqs6GjGHR54cb0`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 08:39:28`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Very thoroughly trace how a user reaches issuance-LC from qcash-ui-header-footer, especially any modal/select interaction used before accessing the page, through qcash-ui shell route wrappers/feature registry, into qcash-ui-issuance-lc. Then trace what happens on the NEXT header navigation after issuance-LC has mounted. Search for router.push/replace, href assignment, route events, modal state, menu click handlers, local/session storage, window events, and stale memo/effect state. Identify likely route-specific causes with exact file:line references and distinguish proven defects from hypotheses. Inspect current worktrees and note branch/dirty state without modifying anything.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
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
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

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

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
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
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

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

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 17: Assistant

## Executive finding

The issuance-LC route has a documented, route-specific history of breaking the **next client-side header-menu navigation** because the federated issuance page uses QUI context/components. The current `qcash-ui-header-footer` branch mitigates that by forcing a full document navigation when leaving `/v2/import-lc/issuance-lc...`.

However, that mitigation is incomplete:

1. Normal horizontal-menu destinations go through the hard-navigation safeguard.
2. Clicking **Import LC** again opens `LcImportModal`; selecting Amendment or Incoming navigates directly with `router.push`, bypassing the safeguard.
3. Avatar “Profile Settings” and breadcrumb navigation also bypass it.

Thus the strongest current route-specific cause is an inconsistent navigation path: the result depends on which header control is used after issuance-LC mounts.

No files were edited.

---

# End-to-end trace

## 1. Header menu data becomes the Import LC menu item

Menu/product data is fetched and cached by:

- `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:183-203`
- `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:208-224`
- `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:250-273`

The raw menu is transformed into an item containing `data.href`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/utils/transformMenu.ts:72-84`
- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/utils/transformMenu.ts:100-110`

`Menu` places everything under `MenuModalProvider`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/Menu.tsx:8-18`
- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/Menu.tsx:22-28`

`MenuMapper` turns a leaf into a click handler that passes the original menu data to `onMenuClick`:

- New skin: `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuMapper.tsx:45-59`
- Old skin: `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuMapper.tsx:76-100`

Its route-derived selected state is recalculated from `asPath`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuMapper.tsx:113-133`

The memoized menu-item structure depends on items, click handler, and language:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuMapper.tsx:135-138`

I found no clear stale-route memo defect here: `selectedKeys` includes `asPath`, and the click-handler structure includes `onItemClick`.

---

## 2. Import LC opens a selector modal instead of navigating immediately

`MenuModalProvider.handleRedirect()` first asks `handleModalOpen()` to consume special menu items:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:123-127`

`href === "/v2/import-lc"` opens `LcImportModal` through an imperative ref:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:59-88`
- Exact Import LC condition: `:73`
- Modal ref rendering: `:165-167`

The modal imperative API is implemented with a toggle:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:51-59`

The choices are:

- Issuance LC: `value: "issuance-lc"` at `:61-67`
- Amendment: `:68-73`
- Incoming: `:74-79`

Authorization filtering uses the bridged or Ordent authority state:

- Auth source selection: `:40-49`
- Memoized authority/menu filtering: `:82-95`
- Choice click wiring: `:154-161`

The shell publishes its auth bridge through the `qc-bridge-sync` window event:

- `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:383-428`

Header-footer consumes and rehydrates that bridge:

- `~/Development/qcash-ui-header-footer/layouts/default/AuthBridge.tsx:84-99`
- `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:7-28`

---

## 3. Selecting Issuance LC

`LcImportModal.handleMenuClick()` builds the destination from `baseFederation` after removing `/v2`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:97-100`

For a different current route, Issuance LC closes the modal and calls:

```ts
router.push(redirectUrl);
```

at:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:108-113`

If issuance is already the current route, it instead reloads through `location.href`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:100-105`

The tests explicitly assert the initial issuance selection uses `router.push`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/__tests__/LcImportModal.test.tsx:103-116`

---

## 4. Shell feature registry and route wrapper

Issuance-LC is registered as a prefix-matched feature:

- `~/Development/qcash-ui/constants/features/registry.ts:516-520`

The shell’s route matcher strips query/hash and performs prefix matching:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:102-135`

Route-to-module resolution:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:160-196`

`ModuleFederationProvider` watches `router.pathname`, determines the active feature, and preloads route-specific styles/modules:

- `~/Development/qcash-ui/components/federation/provider.tsx:114-125`
- `~/Development/qcash-ui/components/federation/provider.tsx:143-158`

The concrete Next page wrapper—not the registry itself—loads both remotes:

- Header/footer remote: `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:10-13`
- Issuance remote: `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:15-18`
- Header/footer layout wrapping issuance: `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:21-27`

Create/detail/edit repeat the same wrapper pattern:

- `~/Development/qcash-ui/pages/import-lc/issuance-lc/create/index.tsx:10-27`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:10-27`
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:10-27`

Remote components are loaded client-side and aborted if their wrapper unmounts:

- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:41-76`
- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:97-100`

The shell’s `/v2` base path is configured at:

- `~/Development/qcash-ui/next.config.js:66-70`

---

## 5. Provider/context structure during mount

The shell wraps pages in:

1. Module federation provider
2. BRI and Ordent auth/global providers
3. Global QUI provider
4. Auth bridge

Relevant locations:

- `~/Development/qcash-ui/pages/_app.tsx:163-197`
- QUI provider specifically: `:69-71`, `:169-172`
- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:43-51`

Because issuance is loaded as an exposed page component, its own `_app.tsx` is not the shell wrapper. It relies on the shared shell providers. Its standalone `_app.tsx` is:

- `~/Development/qcash-ui-issuance-lc/pages/_app.tsx:38-68`

---

## 6. Issuance-LC list mount

The exposed page calls `guard()` once:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:14-28`

It waits for authority state and checks Import LC access:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:30-36`

It then mounts the new-skin list under a query provider:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:38-50`

The list uses QUI components dynamically loaded from the `qui` remote:

- `~/Development/qcash-ui-issuance-lc/containers/new-skin.tsx:1-2`
- `~/Development/qcash-ui-issuance-lc/components/global.tsx:568-590`

The module-level QueryClient is persistent for the lifetime of the loaded remote and is also exposed on `window`:

- `~/Development/qcash-ui-issuance-lc/services/query-client.ts:1-21`

No router-event, history, popstate, or before-history listener exists in the issuance module. Its relevant window listeners are UI scroll/resize/storage listeners, not routing listeners.

---

# What happens on the next header navigation

## A. Ordinary horizontal menu leaf

The click again reaches:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:123-153`

If it is not another modal-opening menu, `handleNavigation()` checks the physical browser pathname:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:91-94`

Every pathname starting with `/v2/import-lc/issuance-lc` is registered for hard navigation:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:37-39`

It then executes:

```ts
globalThis.window.location.href = href;
```

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:95-100`

This forces a complete document unload, thereby destroying the issuance remote and its QUI state before the destination initializes.

## B. Clicking Import LC again and choosing another LC subtype

This takes a different path:

1. Parent Import LC click is consumed by `handleModalOpen()` at `:73`.
2. `handleNavigation()` is never called.
3. `LcImportModal.handleMenuClick()` directly invokes `router.push()` for Amendment or Incoming at `LcImportModal.tsx:108-113`.

Therefore it bypasses the explicit issuance hard-navigation safeguard.

## C. Avatar/header links

The horizontal menu safeguard does not wrap avatar navigation.

“Profile Settings” uses `router.push` directly:

- `~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:81-88`

Other avatar destinations are plain anchors:

- `~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:90-93`
- `~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:115-123`
- `~/Development/qcash-ui-header-footer/layouts/default/container/header/components/NavLink.tsx:52-60`

Plain anchors naturally perform document navigation, but Profile Settings does not.

## D. Breadcrumb navigation

The home icon uses `router.push("/")`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/breadcrumb/components/Breadcrumb.tsx:43-51`

Other crumbs use Next `Link`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/breadcrumb/components/Breadcrumb.tsx:73-80`

These also bypass `MenuModalProvider.handleNavigation()`.

---

# Proven defects / established history

## 1. Proven historical defect: client-side navigation after issuance lost required context

Two committed fixes explicitly identify this behavior:

- Commit `8161f7c`: **“fix: document navigation when leaving issuance LC”**
- Commit `8a86736`: **“adjust navigation issue next route missing context”**

The first commit added the issuance-only `location.href` branch; the second generalized it into `REGISTERED_HARD_NAV_PAGES`.

The current implementation and comment explicitly connect the workaround to issuance’s QUI usage:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:95-98`

This is the strongest evidence that the original failure was not generic menu state—it was failure to fully tear down issuance/QUI state during a Next client transition.

Current status: mitigated for ordinary horizontal-menu leaf clicks.

## 2. Proven control-flow gap: Import LC modal sub-selection bypasses that fix

The hard-nav invariant is in `MenuModalProvider.handleNavigation()`, but modal choices navigate independently:

- Safeguard: `MenuModalProvider.tsx:91-100`
- Modal interception before navigation: `MenuModalProvider.tsx:123-126`
- Direct modal `router.push`: `LcImportModal.tsx:108-113`

This bypass is demonstrable from the code. Whether it reproduces “missing context” for every LC subtype requires runtime confirmation, but it violates the exact route-specific teardown policy introduced by the historical fixes.

## 3. Proven control-flow gap: Profile Settings and breadcrumb home bypass the fix

Direct `router.push` calls remain at:

- `~/Development/qcash-ui-header-footer/layouts/default/container/header/components/ListItem.tsx:83-85`
- `~/Development/qcash-ui-header-footer/layouts/default/container/breadcrumb/components/Breadcrumb.tsx:50`

The bypass is proven; the resulting context failure is a hypothesis pending runtime reproduction.

## 4. Proven listener leak on issuance create page

Several page-3 implementations register a scroll listener during render, with no cleanup:

- `~/Development/qcash-ui-issuance-lc/components/create/page-3.tsx:441-443`
- `~/Development/qcash-ui-issuance-lc/components/create-new-skin/page-3.tsx:1011-1013`
- Equivalent old-skin implementation was also found at `components/create-new-old/page-3.tsx:1004`

Every render can add another listener. A client-side modal sub-route transition can leave callbacks retaining stale component state; hard document navigation hides this defect by destroying the window. This is more relevant when leaving create/edit than when leaving the base issuance list.

## 5. Test coverage does not exercise the important bypass

The hard-nav test only verifies `router.push` was not called and relies on jsdom emitting a `console.error` for location assignment:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/__tests__/MenuModalProvider.test.tsx:56-76`

The LC modal tests affirm that switching subtype uses `router.push`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/__tests__/LcImportModal.test.tsx:103-116`
- Amendment: `:135-149`

There is no integration test covering:

> issuance mounted → Import LC parent clicked → different LC subtype selected → required hard teardown.

---

# Hypotheses and lower-confidence contributors

## Modal state

`LcImportModal.show()` and `.close()` both toggle rather than set an explicit state:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:54-59`

A duplicate `show()` call could close an already-open modal. This is a real state-semantic weakness, but I found no route event that invokes it twice during the normal issuance path.

The modal explicitly sets `show: false` before its normal `router.push`:

- `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:111-113`

So a visibly stuck modal is less likely than incomplete remote/provider teardown.

## Stale effects

The issuance index calls `guard()` in an empty-dependency effect:

- `~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:26-28`

This can capture the first guard function, but it runs only on page mount and is not an obvious cause of the next navigation.

The shell’s remote loader correctly aborts state updates when unmounted:

- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:68-76`

## Storage

No issuance-specific navigation flag was found in local/session storage.

Relevant shell behavior:

- Menu caches: `productMenu`, `productRoles`, `validateMenu`
- Active-feature storage cleanup:
  `~/Development/qcash-ui/components/federation/provider.tsx:114-125`
- Unsaved-leave data can be written to a dynamic session key:
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:130-143`

There is no `setOnLeaveAction` use in the issuance repository, so issuance itself does not appear to drive that confirmation path.

The issuance hard-navigation branch returns before the legacy menu-cache clearing at `MenuModalProvider.tsx:116-120`, but those removals are intentionally attached to non-v2/legacy navigation. I found no evidence they cause this route-specific failure.

## Route events and window events

The shell’s meaningful route listener is session-expiration checking on `routeChangeComplete`:

- `~/Development/qcash-ui/hooks/use-session-expiration.ts:78-106`

It only checks token/modal state; it does not unload remotes.

Auth synchronization uses `qc-bridge-sync` with proper cleanup:

- Producer: `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:423-428`
- Consumer: `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:19-28`

I found no issuance router-event handler capable of cancelling or replacing the next navigation.

---

# Worktree and dirty-state inspection

Only one registered worktree exists for each repository.

## qcash-ui-header-footer

- Path: `~/Development/qcash-ui-header-footer`
- Branch: `dev`
- HEAD: `ef6b18e`
- Tracking: `origin/dev`
- Dirty:
  - Modified: `~/Development/qcash-ui-header-footer/tmp/breadcrumb-diff.json`
  - Untracked: `~/Development/qcash-ui-header-footer/tmp/coverage.txt`

The relevant navigation source files are clean.

## qcash-ui shell

- Path: `~/Development/qcash-ui`
- Branch: `dev`
- HEAD: `4231a1cc`
- Tracking: `origin/dev`
- Clean worktree.

## qcash-ui-issuance-lc

- Path: `~/Development/qcash-ui-issuance-lc`
- Branch: `master`
- HEAD: `28c3cee`
- Tracking: `origin/master`
- Dirty:
  - Modified: `~/Development/qcash-ui-issuance-lc/next.config.js`

The local diff changes the QUI remote from the environment URL to:

```text
qui@http://localhost:49677/static /remoteEntry.js
```

at `~/Development/qcash-ui-issuance-lc/next.config.js:25`.

The embedded space before `/remoteEntry.js` makes this local remote URL suspect and can independently prevent local QUI loading. It is unrelated to the committed header-navigation workaround but materially affects local reproduction.

---

# Likelihood ranking

1. **Highest:** Import LC modal sub-selection after issuance uses `router.push` and bypasses the issuance hard-unmount policy.
2. **High for Profile Settings/breadcrumb reports:** direct `router.push`/Next `Link` bypasses the same policy.
3. **Established root mechanism:** issuance/QUI context survives or is torn down incorrectly during SPA navigation; commit history explicitly calls this “next route missing context.”
4. **Secondary on create/edit:** render-time scroll listeners accumulate and survive client transitions.
5. **Lower:** modal toggle semantics or stale guard/memo state.
6. **Local-only reproduction issue:** dirty malformed QUI remote URL in issuance `next.config.js`.

---

