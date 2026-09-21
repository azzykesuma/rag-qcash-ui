# Compare modal routes (@explore subagent)

- **ID**: `opencode-e4ee64a7bb2e74dba22d3e15011889f0776f6f5a7353296e7caa64d76832d772`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 08:39:29`
- **Project**: `Development`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Search qcash*/addons* repositories for header navigation items that open a selection modal before routing. Compare issuance-LC behavior with at least one route that works correctly. Focus on menu metadata, href manipulation, product/authority selection, modal callbacks, router method used, path/basePath handling, and state reset after route change. Return the smallest behavioral difference likely to explain why only issuance-LC causes the next header navigation to be stale, with exact references.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 22: Assistant

## Finding

The smallest issuance-only difference is **how the next header navigation leaves the page**:

- Normal routes use `router.push(...)`, preserving the mounted federated runtime.
- When the current page is issuance-LC, the current header-footer code explicitly uses `window.location.href`, forcing a document reload and unmounting the issuance/QUI runtime.

This targeted exception was added in commits `8161f7c` (“document navigation when leaving issuance LC”) and `8a86736` (“next route missing context”), strongly identifying retained federation/router context—not menu metadata or authority selection—as the stale-navigation cause.

### Exact references

- Issuance is the only registered hard-navigation page:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:37-39`
- Leaving it performs document navigation:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:91-100`
- Other V2 header items use soft navigation through `router.push`:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:102-113`
- The comment explicitly associates the exception with keeping the federated issuance page/QUI component mounted:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:96-97`
- The test asserts that issuance must not call `router.push`:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/__tests__/MenuModalProvider.test.tsx:56-79`

## Comparison with working selection-modal routes

### Menu metadata

All modal-opening headers originate from the same transformed metadata:

- `href` comes from `rootPath` or product `BaseHreff`:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/utils/transformMenu.ts:70-78`
- The resulting click data contains `href`, `productID`, SSO, and transfer metadata:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/utils/transformMenu.ts:100-107`
- Import LC opens its modal for `/v2/import-lc`; LC Export does so for `/v2/lc-export`:  
  `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:72-73`

Nothing in this metadata handling is issuance-specific.

### Product/authority selection

Import-LC options are distinguished only by values and authority names:

- Issuance: `value: "issuance-lc"`, authority `"Import LC"`  
- Amendment: `"amend-cancel-lc"`/`"amendment"`, authority `"Amend Cancel LC"`
- Incoming: `"incoming"`, authority `"Incoming Document"`

References:  
`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:61-80`

Filtering is shared across all three options:  
`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:82-95`

The issuance page repeats the expected `"Import LC"` menu/authority guard:  
`~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/index.tsx:15-24,30-35`

Therefore authority selection can affect visibility/access, but does not explain stale navigation after successfully entering issuance.

### Modal callback and path handling

Issuance and working Incoming Document use the **same callback branch**:

```ts
setModalState({ method: menuValue, show: false });
router.push(redirectUrl);
```

Reference:  
`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:97-113`

The callback:

1. Removes `"/v2"` from `config.baseFederation`.
2. Builds an absolute `${basePath}/import-lc/${menuValue}` URL.
3. Passes that URL to `router.push`.

References:  
`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:98-100,111-112`

By contrast, working LC Export builds a relative path and lets the host router apply its base path:

- Relative `/lc-export/advise-lc`
- Reset selection
- `await router.push(path)`

Reference:  
`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcExportModal.tsx:51-54`

The host shell has `basePath: "/v2"`:  
`~/Development/qcash-ui/next.config.js:65-70`

The older addons implementation also routed Import LC with an explicit V2-relative path:

```ts
const redirectUrl = `/v2/import-lc/${menuValue}`;
router.push(redirectUrl);
```

Reference:  
`~/Development/addons-web-client/libs/qcash-ui/src/lib/layouts/menu/components/LcImportModal.tsx:60-67`

The current absolute/base-path manipulation is suspicious, but it applies equally to issuance and incoming, so it does not by itself explain an issuance-only stale **next** click.

### State reset

Working LC Export fully resets modal state before routing:

`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcExportModal.tsx:51-53`

Import LC instead leaves `method` set to `"issuance-lc"` or `"incoming"`:

`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:111-112`

Close callbacks do reset it:

`~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:121,132,150`

However, `modalState.method` is never subsequently read by this component, and Incoming uses the identical non-resetting branch. Thus this is unlikely to cause an issuance-only stale header route.

## Conclusion

The strongest and smallest explanatory difference is:

> **A soft `router.push` away from issuance-LC can leave its new QUI/federated page mounted with stale router context; working routes unmount correctly. A full `window.location.href` navigation resets that runtime and header state.**

The current source already contains that narrowly targeted mitigation. The absolute URL/basePath construction and incomplete modal-state reset are inconsistencies worth noting, but neither is issuance-only; the explicit hard-navigation guard and its commit history are the direct evidence.

---

