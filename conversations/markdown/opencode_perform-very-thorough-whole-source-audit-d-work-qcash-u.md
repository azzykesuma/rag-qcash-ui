# Audit source complexity (@explore subagent)

- **ID**: `ses_090c0898cffedwxPa7Gqt0nO1P`
- **Source Tool**: `opencode`
- **Date**: `2026-07-17 15:44:17`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a very thorough whole-source audit of D:\work\qcash-ui-dashboard-auto-collection for over-engineering only, not correctness/security/performance/style. Scan all src/, public/, scripts/, styles/, tests, generated-looking tracked files, and usages/callers. Identify dead code/speculative features, wrappers that only delegate, one-implementation abstractions, duplicated or hand-rolled stdlib/platform functionality, and code that can be materially shrunk. Do not edit. Return ranked candidate findings with exact paths, concise rationale, replacement, and conservative estimated lines removable. Explicitly validate whether each candidate is actually used before reporting.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

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

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (18)</summary>

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
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (16)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
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
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

# Over-engineering audit

Audited all source, callers/imports, public assets, scripts, styles, the single test file, ignored legacy files, generated outputs, and tracked-file status. No files were edited.

## Ranked candidates

### 1. Inert account-filter feature

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx`

**Validation:** `selectedAccounts` only changes the selector’s label/value. It is never passed to the dashboard, mocks, graph, table, query, or service. All account-search/menu machinery is locally consumed but has no downstream effect.

**Rationale:** Roughly half of the 304-line header implements a speculative filter: custom menu, search, checkbox option, select-all logic, controlled menu state, outside-click handling, styles, refs, and extra remote Select typings.

**Replacement:** Remove the account selector until account selection affects the data request; retain only the period selector.

**Conservative removal:** **120–145 lines**

---

### 2. Router abstraction with no consuming descendants

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\hooks\useRouter.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx`

**Validation:** `useAppRouter()` is called only in `_app.tsx`. `AppRouterProvider` wraps `DashboardPage`, but no descendant calls `useAppRouter` or reads `AppRouterContext`. The injected `router` prop therefore enters a context with zero consumers.

**Rationale:** A 77-line synthetic router implements every `NextRouter` member, warnings and no-op events for a single `_app` call where the native pages-router hook is available.

**Replacement:** Use `useRouter` from `next/router` in `_app.tsx`; remove `AppRouterProvider`, the unused page prop, and `useRouter.tsx`. Confirm that the exported remote’s unused `router` prop is not a promised external API before dropping it.

**Conservative removal:** **75–85 lines**

---

### 3. Unused health-status/TanStack Query vertical slice

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\api\dashboard-status.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts` lines 117–129
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.test.ts` lines 1–17
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx` lines 2, 16–20, 40–42
- `D:\work\qcash-ui-dashboard-auto-collection\package.json` line 24
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock`

**Validation:** `dashboardStatusQuery` is used only by its test. No runtime component fetches or displays status. The otherwise-empty `QueryClientProvider` has no `useQuery` consumers anywhere. TanStack Query has no other source usage.

**Rationale:** API endpoint, query descriptor, provider, client configuration, dependency, and test exist solely to test one another.

**Replacement:** Remove the complete status slice and `@tanstack/react-query`; reintroduce Query when the dashboard has a real query consumer.

**Conservative removal:** **35–45 source/lock lines**

---

### 4. Large ignored legacy corpus present in the workspace

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\public\files\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\flags\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\fonts\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\icons\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\images\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\img\**`
- `D:\work\qcash-ui-dashboard-auto-collection\public\bricams.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\cbm.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\qlola.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\vercel.svg`
- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\en\translation.json`
- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\translation.json`
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\analyze-bundle.js`
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\update-version-file.js`
- `D:\work\qcash-ui-dashboard-auto-collection\styles\globals.css`

**Validation:** No source, CSS, configuration, package script, or README reference points to these assets/files. `analyze-bundle.js` has no package script and requires analyzer configuration/dependencies that are absent. `update-version-file.js` has its write operation commented out and is uncalled. The root stylesheet is not imported. `.gitignore` explicitly labels these paths “legacy/generated leftovers.”

Excluded from this finding because they are used:

- `public/favicon.ico`: Next public convention.
- Dashboard locale JSON: statically imported.
- Both `common.json` files: statically imported and potentially consumed by the remote table.

**Rationale:** **115 unused files / 10,890,320 bytes** remain physically present despite already being ignored. This includes 4,320 lines of unrelated translation dictionaries and 56 lines of dead scripts/styles.

**Replacement:** Delete the ignored local corpus after checking for direct-URL consumers outside this repository.

**Conservative removal:** **4,376 textual lines plus 113 binary files (~10.9 MB)**

---

### 5. Generated CSS is committed twice

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.bundle.css`
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\global.exposes.ts`
- Generator: `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs`

**Validation:** Both files are tracked. `globals.bundle.css` is generated by `style:bundle`; its entire contents are then embedded into the generated `global.exposes.ts`. Both `predev` and `prebuild` regenerate them.

**Rationale:** The same generated CSS is committed as a 14,160-byte CSS file and again inside a 15,146-byte TypeScript file, creating review churn and duplicate repository material.

**Replacement:** Treat both as build artifacts and generate them through the existing pre-scripts. The style-exposure mechanism itself is used by Module Federation and should remain.

**Conservative removal:** **27 tracked generated lines / ~29 KB**

---

### 6. Hand-rolled tooltip dismissal duplicates the Popover platform API

**Path**

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\info-tooltip.tsx`

**Validation:** Used only by `SummaryCards`, three instances per row. Its React state, refs, effect, pointer listener, keyboard listener, and cleanup exist solely to provide toggle, outside-click dismissal, and Escape dismissal.

**Rationale:** These are native light-dismiss behaviors of HTML `popover`.

**Replacement:** Use a button with `popovertarget` and a `popover` element; retain the current content and styling.

**Conservative removal:** **18–22 lines**

---

### 7. Mock response models carry unused API compatibility layers

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx`

**Validation:**

- `error`, `code`, and `message` are never read.
- `totalPage` is never read.
- Mocks produce only `date`, `entity_name`, and `payment_code`.
- `transactionDate`, `entityName`, and `paymentCode` aliases are accepted but never produced by any current source.

**Rationale:** The local mock layer models speculative response envelopes and two naming conventions despite having no real dashboard API client.

**Replacement:** Use one normalized internal item shape and retain only `graph`, `summary`, `data`, and `totalData`. Normalize once when a real API is introduced.

**Conservative removal:** **15–20 lines**

---

### 8. Two identical remote-loader wrappers

**Path**

- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts` lines 69–81

**Validation:** Both exported loaders are used, but their implementations are identical; only the caller-supplied module string differs. `loadQuiRemote` has one caller and `loadGlobalComponentRemote` has three.

**Rationale:** The wrappers do not encode remote-specific behavior or constraints.

**Replacement:** Export one `loadRemote<T>(module)` function and use it for both prefixes.

**Conservative removal:** **6–8 lines**

---

### 9. Dead exported utility branches

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\auth.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts`

**Validation:**

- `canCreateDashboardAutoCollectionTransaction` has no caller.
- `PELINDO_DATA_ENTRY_MAKER_AUTHORITY` exists only for that dead function.
- `getSupportedLanguage` has no caller.
- `hasDashboardAutoCollectionProduct` and the remaining i18n helpers are used and are not included.

**Rationale:** Exported speculative capabilities increase the apparent public surface without runtime use.

**Replacement:** Delete the two unused exports and the maker constant. Add them when a create action actually exists.

**Conservative removal:** **7–9 lines**

---

### 10. Unused Tailwind theme tokens

**Path**

- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.css`

**Validation:** Whole-source class search found use of `neutral` and `danger`, but no use of `font-mono`, `primary`, `primary-light`, `primary-link`, `accent`, `success`, `warning`, `secondary`, `pending-bg`, `gray-bg`, or `border-light`.

**Rationale:** Eleven custom tokens are exposed and included in generated CSS for hypothetical consumers while dashboard code primarily uses literal design colors.

**Replacement:** Keep only tokens used by this remote; add named tokens when replacing literals or when they become part of the documented federation stylesheet contract.

**Conservative removal:** **11 lines**

---

### 11. Dead ni18n configuration file

**Path**

- `D:\work\qcash-ui-dashboard-auto-collection\ni18n.config.ts`

**Validation:** `ni18nConfig` has no importer or script reference. The application initializes i18next directly through `src/utils/i18n.ts`. `ni18n` itself remains used as a Module Federation shared module, so only the config is dead.

**Replacement:** Delete the unused configuration file.

**Conservative removal:** **5 lines**

---

### 12. Direct dependency with no source or transitive lock use

**Paths**

- `D:\work\qcash-ui-dashboard-auto-collection\package.json` line 32
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock` lines 2198–2201

**Validation:** `react-is` has no import or require call anywhere. Its lock entry appears only for the direct `^19.2.7` declaration; no other lock package references it.

**Replacement:** Remove the direct dependency.

**Conservative removal:** **5 manifest/lock lines**

## Explicitly reviewed but not reported

- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx`: wrapper is actively used and provides the necessary client-only dynamic Module Federation boundary.
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts`: repeated resource registration looks heavy, but it supports both the standalone provider and an injected host i18n instance when consumed as a remote.
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs`: actively invoked by `predev`/`prebuild`.
- All dashboard component files are reachable from the exposed page.
- Graph/table mock factories and CSV generation are runtime-used and test-covered.
- No `tests/` directory exists; the only test is `src/services/dashboard.test.ts`.

---

