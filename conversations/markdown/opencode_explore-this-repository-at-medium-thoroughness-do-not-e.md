# Map testable source paths (@explore subagent)

- **ID**: `ses_0738bd29affeoSq21i6bLnbX9z`
- **Source Tool**: `opencode`
- **Date**: `2026-07-23 07:50:51`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository at medium thoroughness. Do not edit files. Read every production TypeScript/TSX file under src, identify behavior worth unit-testing, existing test coverage, dependencies that need mocking, and files that should reasonably be excluded as declarative/framework wrappers. Return a concise per-file map and recommend the smallest high-value test additions, especially for files shown by the user: dashboard-config.ts, dashboard-header.tsx, services/dashboard.ts, pages/dashboard/index.tsx, containers/dashboard/index.tsx, components/ui/index.tsx, services/system.ts.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

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

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

## Coverage inventory

There are 20 production TypeScript/TSX files and 26 tests across 5 test files. No Vitest configuration, DOM environment, component-testing dependency, coverage thresholds, or checked-in coverage report exists.

### Per-file map

| Production file | Existing coverage | Behavior worth testing / required mocks |
|---|---|---|
| **`D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts`** | Partial: `getApiPeriod` and leap-year `getPeriodRange` are tested indirectly in `dashboard.test.ts`. | Add direct tests for `formatCurrency`, month-end ranges, and API-period conversion. No mocks. Treat `dashboardContent`, `dashboardTypes`, and `monthNames` as declarative; avoid snapshots of every label. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx`** | None. | Test conditional “All” option, case-insensitive account filtering, clear/fallback selection, period conversion, refresh disabled state, and outside-click closure. Mock `DatePicker`/`Select`; control system time; DOM required. Skip CSS, SVG, and style-function assertions. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts`** | Good partial coverage: endpoint/body contracts, account defaulting, xls→xlsx, missing URL, and basic CSV generation. | Add exact CSV escaping/alias fallback test, AbortSignal forwarding, and one non-payment type case. Mock `customAxios`; mutate/reset `config.dashboardBaseUrl`. Current CSV test only checks row count/content, not escaping or exact output. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx`** | None. | High value: verify `guard()` runs; loading for incomplete/guest auth; unauthorized branch; authorized branch passes `companyID`. Mock `useAuth`, UI wrappers, and dashboard container; DOM/effect renderer required. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx`** | None. | Highest-value component target: bootstrap Pelindo ID/accounts, default account selection, graph/table request arguments, refresh/type/period transitions, loading/error rendering, and abort suppression. Mock both services, translation hook, child components, and `Message`; freeze time before import. Add a failure test: account bootstrap currently appears able to leave the skeleton loading indefinitely. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx`** | None. | Reasonably exclude: declarative Module Federation bindings and TypeScript prop facades. A test mocking `next/dynamic` would mostly duplicate remote strings and `{ssr:false}` declarations. Rely on typecheck plus a host/MFE smoke test. |
| **`D:\work\qcash-ui-dashboard-auto-collection\src\services\system.ts`** | Success path covered. | Add table-driven rejection tests for absent, empty, and non-string values; verify a supplied signal is forwarded. Mock `customAxios.get`. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts` | Strong: all hostname branches, server default, local mock URL, and explicit override are covered by 12 tests. | No immediate additions. Tests correctly need `window`, environment variables, and `vi.resetModules()`. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts` | Strong: token present, absent, and no `localStorage`. | No immediate additions. Existing tests use an Axios adapter and mock only `localStorage`. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\utils\auth.ts` | Core allow/deny matrix and malformed authority values covered. | Optional: malformed/non-array menus. No mocks. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts` | None. | Test reuse of an existing runtime, fallback runtime creation, successful loading, and null-module error. Mock `@module-federation/enhanced/runtime`; package versions and React libraries need not be deeply asserted. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts` | None. | Test stored-locale normalization, unsupported values, storage exceptions, and custom-target resource registration. Mock/reset `window.localStorage`, i18next instances, and modules because browser state is captured at import. Medium priority. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\hooks\useDashboardAutoCollectionTranslation.ts` | None. | Reasonable exclusion as a thin framework hook. Its only effect can be covered through provider/component integration. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\components\providers\i18n-provider.tsx` | None. | Low priority: stored-language change and resource refresh. Mock i18n and `I18nextProvider`. Otherwise exclude as a framework wrapper. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx` | None. | Test API-field alias mapping, search/page/limit requests, stale-request aborting, download format/name, selected-row filtering, and error callbacks. Mock `Table`, dashboard services, `URL`, anchor creation, and `Blob`; DOM required. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx` | None. | One small test for three cards, active `aria-pressed`, formatted totals/counts, and `onSelect`. Mock `Tooltip`/`Skeleton`; skip classes and visual snapshots. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx` | None. | Test month-length truncation, date-to-day fallback, numeric amount conversion, selected summary key, and period label. Mock Recharts to inspect data. A zero-valued tooltip is currently suppressed by a truthiness check and deserves a regression decision. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx` | None. | Reasonably exclude as provider/bootstrap composition. Optional smoke test that global styles remote loads once; mock router, providers, and `loadRemote`. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_document.tsx` | None. | Exclude: standard Next.js document declaration. |
| `D:\work\qcash-ui-dashboard-auto-collection\src\styles\global.exposes.ts` | None. | Exclude generated CSS payload from unit coverage. At most smoke-test stylesheet insertion/update in the generation pipeline, not the CSS text itself. |

## Smallest high-value additions

1. Extend `system.test.ts` with one table-driven invalid-response test and one signal assertion.
2. Extend `dashboard.test.ts` with one exact CSV escaping/alias test and one hold/release plus signal contract test.
3. Add `dashboard-config.test.ts` covering currency and period helpers directly.
4. Add one authorization-branch test file for `pages/dashboard/index.tsx`.
5. Add two container tests: successful orchestration and bootstrap failure/abort behavior.
6. Only then test `dashboard-header.tsx` and `transaction-table.tsx`.

Component tests need a browser environment. The minimum infrastructure addition is `jsdom`; existing `react-dom` can render components while Module Federation UI components are mocked. Exclude `components/ui/index.tsx`, `_document.tsx`, and generated global styles from unit coverage.

---

