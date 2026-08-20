# Audit application complexity (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_081be81a4ffeQyrxKtTgjiSZ01`
- **Source Tool**: `opencode`
- **Date**: `2026-07-20 13:40:49`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a VERY THOROUGH whole-tree over-engineering audit of D:\work\qcash-ui-dashboard-auto-collection\src only. This is NOT a correctness/security/performance review and do not edit. Hunt only dead code, speculative features, wrappers that only delegate, one-implementation abstractions, repeated logic reducible substantially, hand-rolled stdlib/native behavior, unused exports/types/props/functions. Trace references repo-wide before claiming dead code. Exclude generated src/styles/globals.bundle.css. Return ranked candidate findings with exact paths, what can be cut, replacement, and defensible estimated net lines removed. Use tags delete/stdlib/native/yagni/shrink. Also explicitly list areas checked with no finding so coverage is clear.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (13)</summary>

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
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (14)</summary>

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
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

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

- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

# Whole-tree over-engineering audit

Audited all 23 non-excluded files under `src`; `src/styles/globals.bundle.css` was excluded as requested. References were traced through tests, configuration, scripts, generated federation declarations, locales, and documentation. `tsc --noUnusedLocals --noUnusedParameters` reported no compiler-detectable locals. No files were edited.

Net-line estimates are source lines, independent unless noted.

## Ranked candidates

### 1. Replace the hand-built account search/menu state with Select’s built-in behavior
**Tags:** `shrink` `native`  
**Estimated net removal:** **65–75 lines**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:26-36,88-117,199-236,271-301`

**Cut:**
- `AccountMenuListProps`
- `AccountMenuList` and its separate `<input type="search">`
- `accountSearch`, `accountMenuOpen`, `accountFilterRef`, `keepAccountMenuOpen`
- Manual outside-click listener
- `visibleAccountOptions`, `closeAccountMenu`
- `menuSearchValue` and `onMenuSearch*` plumbing
- Multi-select configuration that is reduced back to one selected account with `options.at(-1)`

**Replacement:**
Use the existing remote Select as a normal searchable single-select:

- `isSearchable`
- `options={accountOptions}`
- `value={selectedAccount}`
- direct single-option `onChange`

The generated declaration confirms that the remote Select extends react-select’s `StateManagerProps`:

- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\_types\components\form\select\select.d.ts:9-24`

The custom `AccountOption` renderer and style overrides can remain if their visuals are required.

---

### 2. Delete hand-maintained mirrors of generated remote component types
**Tags:** `delete` `shrink`  
**Estimated net removal:** **68–78 lines**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:6-86`

**Cut:**
- `SelectOption`
- `GlobalSelectProps`
- `GlobalTableProps`
- `MessageProps`
- `LoadingOverlayProps`
- `UnauthorizedDialogProps`

These duplicate generated declarations already present at:

- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\select.d.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\table.d.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\message.d.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\loading-overlay.d.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\@mf-types\qcash-global-component\unauthorized.d.ts`

**Replacement:**
Type each loader from its generated module, for example with `typeof import("qcash-global-component/table")`, or import the generated prop type. Retain a small intersection only for runtime props missing from stale generated declarations.

Keep the local `SkeletonProps` and provider typing: no corresponding `qui` declarations were found.

This overlaps candidate 1 by roughly five custom Select-property lines.

---

### 3. Register i18n resources once per actual instance, not repeatedly everywhere
**Tags:** `delete` `shrink`  
**Estimated net removal:** **17–21 lines**

**Paths:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:53-78`
- `D:\work\qcash-ui-dashboard-auto-collection\src\components\providers\i18n-provider.tsx:1-14`

**Current duplication:**
- The local instance is initialized with complete `resources`.
- It is then populated again immediately.
- It is populated again from the initialization promise.
- The provider populates it again on mount.
- Language synchronization repeats the population.
- `ensureDashboardAutoCollectionResources` always touches the local instance, global singleton, and optional target.

**Replacement:**
- Let `i18nInstance.init({ resources: ... })` initialize the standalone instance.
- Make `ensureDashboardAutoCollectionResources(target)` populate only its required target.
- Keep the dashboard translation hook’s target registration because the Module Federation page can be consumed without local `_app.tsx`.
- Make `I18nProvider` a plain `I18nextProvider`; its stored language was already selected during browser-side instance initialization.

The hook itself should remain—it handles the externally supplied host i18n context.

---

### 4. Remove unused response metadata and one-use DTO abstractions
**Tags:** `delete` `yagni` `shrink`  
**Estimated net removal:** **14–17 lines**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:12-17,27-30,41-66`

**Cut:**
- Unread `error`, `code`, and `message` fields from all three response interfaces
- Unread `totalPage`
- One-use `DashboardAccountsResponse`; inline `{ data: DashboardAccount[] }`
- One-use `DashboardGraphItem`; inline its two-field graph item shape
- Unneeded exports from `DashboardTableItem` and remaining internal-only response types

TypeScript response types do not need to enumerate wire fields no consumer reads. Repo-wide search found no reads of these fields.

Also remove the unused `export` from:

- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:13` — `SupportedLanguage` is internal only.

That export-surface reduction removes no physical line by itself.

---

### 5. Use a native select for the fixed twelve-month period
**Tags:** `native` `shrink`  
**Estimated net removal:** **10–14 lines**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:58-61,189-196,207,303-314`

**Cut:**
- Object-based `periodOptions`
- `PeriodSingleValue`
- `selectedPeriod`
- Remote Select configuration for twelve static choices

**Replacement:**
A controlled `<select value={period}>` with twelve `<option>` elements generated from `monthNames`. This retains the fixed month selection without a federated component plus a custom value renderer.

---

### 6. Delete proven-unused dashboard content fields
**Tags:** `delete` `yagni`  
**Estimated net removal:** **6 lines**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:8-9,19-20,30-31`

**Cut:**
- `amountTitle`
- `countTitle`

Both occur once per transaction type and have zero references repo-wide. The used fields are `cardTitle`, `totalTitle`, `detailTitle`, `amountHeader`, keys, and tooltip.

---

### 7. Reduce the exposed stylesheet module to its side-effect contract
**Tags:** `delete` `yagni` `shrink`  
**Estimated net removal:** **6–8 lines**  
**Confidence:** Conditional on confirming the external Module Federation contract.

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\global.exposes.ts:4-25`

**Cut candidates:**
- `styleId`
- `cssText`
- Named `ensureGlobalsStylesheet` export
- Default export
- Returned CSS text
- Reassigning identical CSS when the style ID already exists

**Replacement:**
Keep a local injection function plus its module-level invocation; consumers load `./globals` for its side effect.

Repo-wide there are no consumers of the named/default exports. The module is externally exposed at:

- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:49`

The project’s own loading convention ignores the returned module:

- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:18-20`

Because external host repositories cannot be searched here, this is not an unconditional dead-code claim. The generator must also be adjusted to prevent regeneration:

- `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs:10-34`

---

### 8. Remove the unused Tailwind theme color
**Tags:** `delete`  
**Estimated net removal:** **1 line**

**Path:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.css:6`

**Cut:** `--color-danger`

No `danger` utility or variable reference exists under `src`. `--font-sans` and `--color-neutral` are both used.

---

### 9. Collapse the one-line constant module into its only owning service
**Tags:** `shrink`  
**Estimated net removal:** **1 line and one file**

**Paths:**
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\constant.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:1,19-24`

**Cut:** The dedicated `constant.ts` module.

**

---

