# Inspect header navigation (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_0261875afffeYtLm7Yj6P4Ee1k`
- **Source Tool**: `opencode`
- **Date**: `2026-08-07 08:47:37`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Analyze D:\work\qcash-ui-header-footer and D:\work\qcash-ui-header-footer-bri for menu/navigation code, especially isNewUI behavior. Determine how navigation is performed (Next router, raw history, anchors), whether children are memoized/stale, and whether route event subscriptions or blockers can leave URL changed while content does not. Focus on interaction after entering a federated Issuance LC page. Do not edit. Return exact file/line references and likely or excluded causes. Thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

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
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
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

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

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

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

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

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 14: Assistant

## Executive conclusion

The strongest cause is the loss of a deliberate hard-navigation workaround for the federated Issuance LC page.

In the current `qcash-ui-header-footer` working tree:

- `/v2/import-lc/issuance-lc` is registered as requiring hard navigation.
- The code detects that route.
- But the actual `window.location.href = href` branch is commented out.
- Navigation therefore falls through to `router.push(...)`.

That is exactly the path capable of changing browser history while leaving a federated page mounted if the host router transition does not commit. The committed version of the code explicitly states that document navigation is needed to guarantee unmounting of Issuance LC.

`isNewUI` itself only changes a background CSS class. It does not alter navigation, memoize content, subscribe to router events, or block transitions.

---

## 1. Current Issuance LC interaction flow

### Opening Import LC

The new header/footer receives transformed menu data and gives leaf menu items button callbacks rather than anchors:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:45-73`
- Old-skin equivalent: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:76-101`

The callback is passed through:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\Menu.tsx:8-18`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\Menu.tsx:22-27`

An exact `/v2/import-lc` menu href does not navigate directly. It opens the Import LC modal:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:59-88`
- Specifically: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:72-74`

### Selecting Issuance LC

The modal defines Issuance LC as `issuance-lc`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:61-80`

On selection it:

1. Removes `/v2` from `baseFederation`.
2. Builds an absolute URL such as `https://host/import-lc/issuance-lc`.
3. Calls `router.push(redirectUrl)` for Issuance LC.

References:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:97-103`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:112-119`

The test confirms a different-page Issuance selection is expected to use `router.push`, not an anchor/document navigation:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\LcImportModal.test.tsx:103-116`

Exceptions:

- Selecting the already-current Import LC subpage assigns `globalThis.location.href`, forcing a reload:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:105-109`
- Legacy `amendment` also uses `location.href`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:112-115`

Thus initial entry to Issuance LC normally uses Next routing; it does not necessarily tear down the current document.

---

## 2. Navigation after entering Issuance LC

The route is explicitly registered as a hard-navigation page:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:37-39`

The current pathname is checked:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:91-95`

But in the current working tree, the hard-navigation branch is commented:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:96-101`

The comment explains the intended reason:

> A document navigation guarantees the federated issuance page is unmounted.

Because it is disabled, a normal `/v2/...` menu click from a `/v2/...` page does this:

```ts
const _href = href.split("/").slice(2).join("/");
router.push(`/${_href}`);
```

References:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:103-117`
- Actual relative push: `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:108-116`

A non-v2 target instead performs a real document navigation:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:120-125`

### Git evidence

The hard-navigation behavior was deliberately introduced by:

- Commit `8161f7cc` — `fix: document navigation when leaving issuance LC`
- Follow-up commit `8a867362` — `adjust navigation issue next route missing context`
- Merged by `552a3be4` — `Feature/fix issuance lc navigation`

In committed `HEAD`, the branch is active. The current checkout has uncommitted changes commenting it out. The working-tree diff changes this:

```ts
if (isHardNavPage) {
  window.location.href = href;
  return;
}
```

into comments.

The existing test still expects hard navigation:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\MenuModalProvider.test.tsx:56-79`
- It expects `router.push` not to be called at `:74`.

That test is inconsistent with the current implementation and would statically be expected to fail.

### Assessment

**Very likely cause:** the commented hard-navigation block reintroduces the exact federated-page navigation problem that the July 31 fixes addressed.

---

## 3. URL changes while content remains

There is no direct `history.pushState` in production source. Therefore, if a top-menu click changes the URL without changing content, that URL update comes through `router.push`, not custom history manipulation.

The general menu calls `router.push` without awaiting or inspecting its result:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:108-116`

Issuance modal navigation also does not await/catch the push:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcImportModal.tsx:116-119`

Consequences:

- A route cancellation or failed page-resolution transition is not observed.
- A federated/host router-context mismatch can remain silent at this layer.
- The browser may already have received a history update even though the host never committed replacement content.

The hard-navigation workaround avoids that entire lifecycle by forcing a document reload and guaranteed remote unmount.

The header/footer is itself a federated layout:

- `D:\work\qcash-ui-header-footer\module-federation.config.mjs:2-16`
- Exposed layout: `D:\work\qcash-ui-header-footer\module-federation.config.mjs:9-12`

The commit wording “next route missing context” strongly supports a router/federation boundary issue rather than a menu rendering issue.

---

## 4. `isNewUI` behavior

### New header/footer

`isNewUI` is accepted and passed into `Layouts`:

- `D:\work\qcash-ui-header-footer\layouts\default\index.tsx:4-5`
- `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:24-33`

Its only runtime use is choosing the outer background class:

- `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:82-85`

Behavior:

- `true` → `hf-bg-white`
- false/undefined → `hf-bg-[#F2F5F5]`

It does **not**:

- Select a different router.
- Change `router.push` to `location.href`.
- Key or remount the child.
- Memoize the child.
- Add route event subscriptions.
- Install navigation blockers.
- Control the Issuance LC special case.

The special hard-navigation rule is pathname-based, not `isNewUI`-based:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:37-39`
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:93-94`

The “QUI new component” comment associates the issue with new-UI pages, but the `isNewUI` prop itself has no navigation semantics.

### BRI header/footer

There is no `isNewUI` occurrence in `qcash-ui-header-footer-bri`.

Its exported layout accepts only `children`:

- `D:\work\qcash-ui-header-footer-bri\layouts\default\index.tsx:4-10`
- `D:\work\qcash-ui-header-footer-bri\layouts\default\container\index.tsx:7-13`

**Excluded cause:** `isNewUI` directly causing stale content or blocking navigation.

---

## 5. Are page children memoized or stale?

No page child is memoized by either layout.

### New header/footer

The exported wrapper directly forwards the child:

- `D:\work\qcash-ui-header-footer\layouts\default\index.tsx:4-5`

The layout renders it directly:

- `D:\work\qcash-ui-header-footer\layouts\default\container\index.tsx:90-96`

Neither `ComponentLayoutDefault` nor `Layouts` is wrapped in `React.memo`.

`MenuModalProvider` also renders its child directly:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:161-164`

However, that provider’s child is only `<MenuContent />`, not the host page:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\Menu.tsx:22-27`

### BRI header/footer

Same direct page-child rendering:

- `D:\work\qcash-ui-header-footer-bri\layouts\default\index.tsx:4-10`
- `D:\work\qcash-ui-header-footer-bri\layouts\default\container\index.tsx:18-24`

### What is memoized

Only `MenuMapper` is wrapped in `memo`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:24`
- `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\MenuMapper.tsx:15`

That memoization can preserve menu rendering, but it cannot preserve the federated page child because the page is outside `MenuMapper`.

There is one minor menu-only stale-state possibility in the new mapper:

- New-skin per-item selection uses `asPath`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:45-50`
- But transformed menu items are cached without `asPath`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:135-138`

This may leave per-button new-skin highlight classes stale. The Ant `selectedKeys` value is separately recomputed with `asPath`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:113-133`

It cannot explain old Issuance content remaining on screen.

**Excluded cause:** header/footer memoization of the routed page child.

---

## 6. Route subscriptions and blockers

I found no production use in either repository of:

- `Router.events.on/off`
- `routeChangeStart`, `routeChangeComplete`, or `routeChangeError`
- `beforeHistoryChange`
- `beforePopState`
- `popstate`
- `beforeunload`
- Production `history.pushState` or `history.replaceState`
- A thrown Next-route cancellation

The only raw history operations are in tests, for setting JSDOM URLs:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\__tests__\MenuModalProvider.test.tsx:29,63`
- `D:\work\qcash-ui-header-footer\layouts\default\container\__tests__\index.onboarding-tour.test.tsx:39-40`
- `D:\work\qcash-ui-header-footer\layouts\default\container\__tests__\index.loan-application.test.tsx:45-46`

### `onLeaveAction` guard

There is a header-level unsaved-state guard:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:136-151`
- BRI equivalent:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\Menu.tsx:105-125`

If `onLeaveAction.isLeave` is true, it opens a confirmation modal and returns without navigating. Navigation happens only on confirmation:

- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:137-149`

The confirmation component calls `onConfirm` and then closes:

- `D:\work\qcash-ui-header-footer\components\ui\confirm-modal.tsx:83-89`

Therefore this guard can prevent or delay navigation, but it does not first change the URL and then preserve old content. It is not subscribed to Next route events and does not guard breadcrumb links or raw anchors globally.

**Excluded within these repositories:** a route-event blocker that independently changes the URL while cancelling the page replacement.

A blocker or route-event subscriber in the host application or Issuance LC remote remains possible because those codebases were not part of this search.

---

## 7. Navigation methods by UI area

### Main menu — new header/footer

- Leaf items: `<button>`/Ant Menu callbacks, not anchors:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuMapper.tsx:54-73`
- `/v2` target while currently in `/v2`: relative `router.push`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:108-113`
- `/v2` target from non-v2: full `/v2/...` `router.push`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:113-116`
- Non-v2 target: `window.location.href`:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:120-125`
- Issuance hard-navigation special case: currently disabled:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:93-101`

### Main menu — BRI header/footer

Leaf menu labels have direct `onClick` handlers:

- `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\MenuMapper.tsx:44-57`

Routing:

- Non-v2 → `window.location.href`:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\Menu.tsx:89-93`
- v2-to-v2 → stripped relative `router.push`:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\Menu.tsx:93-102`
- Other v2 → full href `router.push`:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\menu\components\Menu.tsx:102`

BRI has no Issuance LC hard-navigation exception and no Import LC modal. If this older remote is deployed, leaving a federated Issuance page through another v2 menu is always a client-side Next transition.

### Header/avatar links

Most explicit href entries use raw anchors, producing normal browser document navigation:

- New:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\NavLink.tsx:52-59`
- BRI:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\header\components\NavLink.tsx:35-54`

The new profile setting is a button using `router.push`:

- `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\ListItem.tsx:81-88`

### Breadcrumbs

Breadcrumb paths use Next `<Link>`:

- New:
  - `D:\work\qcash-ui-header-footer\layouts\default\container\breadcrumb\components\Breadcrumb.tsx:73-83`
- BRI:
  - `D:\work\qcash-ui-header-footer-bri\layouts\default\container\breadcrumb\components\Breadcrumb.tsx:56-73`

Home icons use `router.push("/")`:

- New: `D:\work\qcash-ui-header-footer\layouts\default\container\breadcrumb\components\Breadcrumb.tsx:43-51`
- BRI: `D:\work\qcash-ui-header-footer-bri\layouts\default\container\breadcrumb\components\Breadcrumb.tsx:31-39`

---

## 8. Ranked causes

### Very likely

1. **Hard-navigation workaround disabled in the current `qcash-ui-header-footer` tree.**  
   The code, test, commit history, and comment all identify Issuance LC as requiring document navigation to unmount the federation page.

2. **Federated router-context/lifecycle mismatch during v2-to-v2 `router.push`.**  
   This matches the fix commit wording “next route missing context” and explains a history update without host content replacement.

3. **If `qcash-ui-header-footer-bri` is active, it never had the Issuance hard-navigation workaround.**  
   It always uses `router.push` for v2-to-v2 top-menu navigation.

### Possible, but external to these repositories

4. **A route event handler or blocker in the host shell or Issuance LC remote cancels/fails the transition after history changes.**  
   No such subscriber exists in these two repositories, but neither the host nor Issuance implementation was searched.

5. **Multiple or mismatched Next router contexts across federation boundaries.**  
   The header/footer consumes `next/router` inside an exposed remote layout. The configuration exposes the component but does not itself show the host’s Next/router sharing setup:
   - `D:\work\qcash-ui-header-footer\module-federation.config.mjs:9-37`
   - `D:\work\qcash-ui-header-footer-bri\next.config.js:5-35`

6. **Unhandled `router.push` failure/cancellation.**  
   General menu pushes and Issuance modal pushes are not awaited or caught.

### Lower probability

7. **String-based v2 detection/path stripping.**  
   The code uses `href.includes("v2")` and `window.location.href.includes("v2")`, then blindly drops the first two slash-separated components:
   - `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx:103-116`
   
   It works for standard `/v2/foo`, but can misclassify absolute URLs, query strings, or unusual paths.

### Excluded by these repositories

- `isNewUI` changing navigation behavior.
- Page children being memoized by the header/footer.
- Main-menu navigation using raw anchors.
- Production code manually calling `history.pushState`.
- A local Next route-event subscription cancelling the transition.
- The menu-data cache itself preserving the old page component.
- The `isHardNavPage` variable currently doing anything; its only functional branch is commented.

No files were edited.

---

