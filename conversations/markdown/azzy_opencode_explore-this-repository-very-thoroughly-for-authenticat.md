# Map access and locales (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_01046c196ffeH6UxNUtEK0d5HZ`
- **Source Tool**: `opencode`
- **Date**: `2026-08-11 14:28:42`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository very thoroughly for authentication, authorization roles, routing/host integration, localization, accessibility-relevant behavior, browser/responsive considerations, and all configuration/deployment concerns relevant to QA. Also inspect README/docs for stated functionality. Do not edit files. Return structured findings with exact file references.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (16)</summary>

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

<details><summary>🔧 Tool Invocations (13)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (17)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
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

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
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

<details><summary>🔧 Tool Invocations (11)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

# Repository QA findings

Static inspection only; no files were edited and no commands that generate build/test output were run.

## 1. Stated and implemented functionality

- The README only describes the project as a Next.js Module Federation micro-frontend and documents installation, proxy development, build, start, lint, and format commands. It does not document user workflows, authorization requirements, API contracts, supported languages, responsive targets, or deployment variables.  
  `D:\work\qcash-ui-dashboard-auto-collection\README.md:1-42`
- The implemented dashboard provides:
  - Account selection and month filtering.
  - Refresh with a last-updated timestamp.
  - Payment, hold, and release summary cards.
  - Daily amount trend chart.
  - Searchable, paginated transaction table.
  - CSV/XLSX/PDF downloads, including selected-row downloads.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:34-163`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:18-136`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:40-129`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:37-139`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:45-181`
- The root route is only an unauthenticated welcome heading. The functional page is `/dashboard`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\index.tsx:1-3`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:14-45`
- `TRANSLATION_KEYS_MAPPING.md` describes a migration from hardcoded text and backend language-management keys. Much of that migration is now implemented, so portions are historical/stale—especially the claim that month names remain static.  
  `D:\work\qcash-ui-dashboard-auto-collection\TRANSLATION_KEYS_MAPPING.md:3-24`  
  `D:\work\qcash-ui-dashboard-auto-collection\TRANSLATION_KEYS_MAPPING.md:120-134`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:125`
- `qship-mock.md` explicitly says it is temporary benchmark documentation and “Do not merge,” although it is tracked.  
  `D:\work\qcash-ui-dashboard-auto-collection\qship-mock.md:1-5`

## 2. Authentication

### Runtime behavior

- In standalone Next operation, `_app.tsx` wraps pages in `AuthProvider`, passing the hostname-derived API base URL and the Next pages router.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:15-30`
- The dashboard calls `guard()` after mount. Until authorities, username, and company ID are available, it displays a loader. `"Guest"` is also treated as unauthenticated.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:15-38`
- Missing authentication data does not redirect or produce an authentication error. If readiness never arrives, the loader can remain indefinitely. Username comparison is exact and case-sensitive.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:32-38`
- API authentication reads `access-token` from `localStorage` for every Axios request and attaches `Authorization: Bearer …`. No cookie-based or server-side authentication exists in this repository.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:15-21`
- `localStorage.getItem()` is not protected by `try/catch`; browsers/privacy modes that throw on storage access can prevent requests. By contrast, locale storage access is safely caught.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:15-18`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:30-37`
- HTTP 401 responses deliberately do not produce the global error toast. Any logout, refresh, or redirect behavior must come from the external auth provider or host.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:24-37`
- Authentication-provider behavior itself cannot be fully inspected because it comes from `@ordentco/addons-auth-provider`.  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:21-24`

### QA scenarios

- Expired/missing/malformed token.
- Blocked storage access.
- `Guest`, empty username, empty company ID, and authorities that never become ready.
- 401 from each API endpoint and expected host logout/redirect behavior.
- Auth-provider context supplied by a host versus standalone `_app.tsx`.

## 3. Authorization, roles, and company scope

- Dashboard access requires all three:
  1. `userType === "cu"`.
  2. `productAuthorities.DASHBOARD_AUTO_COLLECTION.anyAuthority === true`.
  3. Exact menu string `"Dashboard Auto Collection"`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\auth.ts:1-6`
- Missing permission renders the remote `UnauthorizedDialog`; there is no redirect or local fallback.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:40-42`
- There is no maker/checker role enforcement and no `data_entry:maker` check in the current source. `"maker"` appears only as a test username.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\__tests__\pages\dashboard.test.tsx:57-76`
- `AGENTS.md` is materially stale:
  - It claims a `PELINDO` product-key requirement and maker-role check that do not exist.
  - It claims `AppRouterProvider`, `AppProviders`, and `useAppRouter()` integration that do not exist.
  - It claims there are no tests/Vitest dependency, while both are present.
  - Its Dockerfile warning no longer matches the current Yarn-based Dockerfile and scripts.  
  `D:\work\qcash-ui-dashboard-auto-collection\AGENTS.md:4-7`  
  `D:\work\qcash-ui-dashboard-auto-collection\AGENTS.md:22-31`  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:14-19`  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:44-50`  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:30-40`
- Pelindo company IDs are fetched from `/system/key-secure/acs-company-id-pelindo`. Matching uses a trimmed comma-separated list. Matching companies default to `"All"` and can select all accounts; other companies default to their first returned account.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\system.ts:7-13`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:19-25`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:60-90`
- This company check affects UI/account selection, not the primary page permission. Backend endpoints must independently enforce user, company, account, and download scope; the browser-side checks are not a security boundary.
- API responses’ `error` and `code` fields are not validated before their data is consumed. A nominal HTTP success carrying an application error may be treated as success.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:12-17`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:41-47`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:59-66`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:75-79`

## 4. Routing and host/MFE integration

- The application uses Next.js pages routing with base path `/dashboard-auto-collection-v2`. The standalone dashboard URL is therefore `/dashboard-auto-collection-v2/dashboard`.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:5-7`  
  `D:\work\qcash-ui-dashboard-auto-collection\README.md:29-35`
- Module Federation exposes:
  - `.` and `./root` → welcome page.
  - `./dashboard` → authenticated dashboard page.
  - `./globals` → generated CSS injector.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:33-55`
- Federation is configured only in the client build. All remote UI components also use `next/dynamic(..., { ssr: false })`. Expect client-only rendering and loading gaps for the federated design system.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:15-18`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:114-173`
- Importing `./dashboard` into a host does not execute this repository’s `_app.tsx`. The host must provide compatible shared auth/i18n/runtime context. The dashboard page itself only reloads QUILA globals; `_app.tsx` is responsible for `AuthProvider`, `I18nProvider`, `QuiProvider`, and qcash-global styles in standalone mode.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:18-29`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:14-34`
- The host contract is not documented in README: required providers, shared dependency versions, router expectations, style loading, and remote URLs are omitted.
- Existing Federation runtime is always preferred. If no runtime exists, the fallback points to localhost port 3000 for both external remotes.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:15-33`
- Build-time remote URLs are hardcoded to the HTTP development cluster. On an HTTPS deployment this can be blocked as mixed content. Production/staging remote selection is not configurable here.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:38-41`
- Shared packages are singleton but all use `requiredVersion: false`, so incompatible host/remote package combinations will not be rejected during federation setup.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:42-49`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:34-64`
- The dashboard moves its stylesheet to the end of `<head>` on mount and recreates it from cached text if the host removed it. This intentionally changes cascade order after navigation.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:10-28`
- Exposed Tailwind CSS includes global preflight/reset rules, not only prefixed utilities. Re-appending it can affect host elements outside this MFE.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.css:1-8`  
  `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs:8-33`
- The local proxy:
  - Defaults to port 3001 and starts Next on 3002.
  - Strips `origin`, `referer`, and `sec-fetch-*` only for Next/HMR-style paths.
  - Supports configurable host, ports, target origin, and process spawning.  
  `D:\work\qcash-ui-dashboard-auto-collection\scripts\dev-host-compat.mjs:46-67`  
  `D:\work\qcash-ui-dashboard-auto-collection\scripts\dev-host-compat.mjs:118-191`
- The proxy has no explicit WebSocket `upgrade` forwarding, despite README claiming HMR compatibility. Verify hot reload through the host shell.  
  `D:\work\qcash-ui-dashboard-auto-collection\README.md:27-34`  
  `D:\work\qcash-ui-dashboard-auto-collection\scripts\dev-host-compat.mjs:69-103`

## 5. Localization

- Supported languages are only Indonesian (`id`) and English (`en`); default and fallback are English. Stored locale variants such as `ID_id` and `en-US` are normalized. Unsupported values fall back to English.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:9-27`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:40-49`
- Locale is read from `localStorage["locale"]`. There is no navigator detection, host prop, URL locale, language switch in this MFE, or storage-event synchronization.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:15-37`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\providers\i18n-provider.tsx:6-20`
- Module and common translations are bundled directly from tracked JSON, not fetched from the backend language endpoint mentioned in the mapping document.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:4-21`  
  `D:\work\qcash-ui-dashboard-auto-collection\TRANSLATION_KEYS_MAPPING.md:3-18`
- The document always declares `<html lang="en">`, even when Indonesian is active.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_document.tsx:3-11`
- Missing/inconsistent translations:
  - `dashboardAcs.table.downloadFormat` is used but absent from both module locale files, so its English fallback appears in Indonesian.
  - Indonesian `loadChartFailed` and `loadTableFailed` values remain English.
  - Loader `"Loading"`, toast `"Copy"`/`"Copied"`, `"Process ID"`, `"WIB"`, and the welcome page are hardcoded.
  - Currency always uses Indonesian numeric separators, regardless of selected UI language.
  - Date ranges always use `DD/MM/YYYY`.
  - Chart units remain `IDR` and `M`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:148-174`  
  `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\dashboard-auto-collection.json:63-69`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:58-63`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:37-55`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:43-64`
- Last-updated time is formatted in the browser’s local timezone but always labeled `WIB`; no `Asia/Jakarta` timezone is configured. Test outside WIB.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:23-32`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:145`
- Mandarin is discussed in the mapping document and allowed by the remote DatePicker TypeScript interface, but no Mandarin resources or app support exist.  
  `D:\work\qcash-ui-dashboard-auto-collection\TRANSLATION_KEYS_MAPPING.md:11-18`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:20`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:10`

## 6. Accessibility-relevant behavior

### Positive implementation

- Semantic landmarks and headings are present: `<main>`, sections, `<h1>`, `<h2>`, and summary `<article>` elements.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:131-163`
- Decorative header blocks and icons are marked `aria-hidden`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:134-138`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:49-52`
- Summary loading state uses `role="status"`, `aria-live="polite"`, and hides skeleton decoration.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:13-37`
- Summary cards are native buttons with localized labels, `aria-pressed`, and explicit keyboard focus styling. Info controls are separate labeled native buttons.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:66-101`
- Refresh is a native button and is disabled while requests are active. Account and period controls receive accessible labels.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:43-70`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:103-131`
- The download-format control has an aria-label.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:147-167`

### QA risks

- The fallback spinner shown before/failing Lottie load has no status role or accessible name. The Lottie state has a hardcoded English `aria-label`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:58-64`
- No `prefers-reduced-motion` behavior exists. Lottie loops automatically, spinners rotate continuously, and card/loader transitions remain enabled.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:53-63`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:63`
- The chart has no accessible title, description, table alternative, or daily-value text equivalent. Recharts tooltip interaction should be checked for keyboard and screen-reader access.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:103-135`
- A zero chart value never displays a tooltip because the tooltip returns `null` for falsy values.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:26-34`
- Summary info controls are only 20×20 pixels, below common 44×44 touch-target guidance.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:77-100`
- The account popup search input relies on placeholder text; no explicit label is passed to that nested search field. Search clear-button labeling is delegated to the remote component.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:80-90`
- Alert, toast, unauthorized dialog, autocomplete, date picker, and table semantics are supplied by remote packages and are mocked in local tests. Their focus management, announcements, escape behavior, row-selection labels, and modal semantics require integrated QA.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:114-173`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.test.ts:18-43`
- Error rendering does not move focus to alerts or messages. Refresh/account/type changes do not announce completion beyond the summary skeleton’s initial status.

## 7. Browser and responsive considerations

- TypeScript targets ES2017 with DOM libraries; no Browserslist or explicit supported-browser policy is present.  
  `D:\work\qcash-ui-dashboard-auto-collection\tsconfig.json:2-14`  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:1-52`
- CSS uses `:has()`, requiring reasonably modern browser versions.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.css:18-24`
- Header layout changes to a row at the `sm` breakpoint. The account popup is constrained to viewport width, which is positive.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:35-38`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:80`
- Mobile risks:
  - Account and fixed-width 220px period controls remain in one row.
  - Summary cards and loading cards always use a non-wrapping three-column flex row with 20px gaps and 24px card padding.
  - Table columns total at least 860px; horizontal behavior is delegated to the remote DataTable.
  - Table receives unprefixed `w-full`, while local Tailwind is configured with the `fpl` prefix.
  - Action bar always keeps format control and download button in a single row.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:56-57`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:102`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:16-35`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:43-64`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:50-60`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:139-175`
- Chart width is responsive, but height is fixed at 330px and axis labels use 11px text.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:103-124`
- Loader assets use root-relative URLs, ignoring the configured base path. Under standalone serving, the expected application-owned location would normally include `/dashboard-auto-collection-v2`; under a host, the requests go to the host root.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:17-25`  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:6`
- Blob downloads use `URL.createObjectURL`, synthetic anchor click, and immediate revocation. Verify Safari/iOS behavior and large-file memory use.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:27-34`

## 8. Data and error behavior relevant to QA

- The selectable period includes the current month plus months where the index is not less than `current - 12`, effectively 13 inclusive month values. Confirm whether the requirement is “current plus previous 12” or exactly 12 months.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:107-111`
- Default period is computed once at module import. A session left open across a month boundary will not automatically change its default month.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:31-32`
- Graph date parsing accepts both apparent `DD-MM-YYYY` and `YYYY-MM-DD`; invalid dates fall back to array position. Duplicate days overwrite previous values.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:60-87`
- Table row identity is `paymentCode`. Missing codes become `"-"` and duplicate payment codes will collide during selection/download.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:36-42`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:117`
- Account, graph, and table requests use abort signals. Downloads do not and expose no loading/disabled state, allowing repeated submissions.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:52-120`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:63-82`
- Error presentation differs by endpoint:
  - GET failures and table-data POST failures create persistent remote toasts.
  - Graph/accounts failures create three-second messages.
  - Chart/table also display inline alerts.
  - Downloads create a message and append process ID as plain text.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:24-35`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:27-59`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:162`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:68-101`
- Download filenames are mandatory and parsed from `Content-Disposition`; a successful blob without that header is treated as an error.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:91-110`

## 9. Configuration and deployment

- Runtime API selection is actually build-time-injected public configuration followed by client hostname matching:
  - `gym-master`: DEV ID/TL/SG.
  - `localhost`: DEV ID.
  - `addons.cms.dev`: PRESTAG.
  - `addons.apps`: PREPLOY.
  - `demo.addons`: PILOTING.
  - `bricams`: PROD ID/TL/SG.
  - Unknown/server hostname: PROD ID.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts:1-19`
- `NEXT_PUBLIC_DASHBOARD_AUTO_COLLECTION_API_URL` can override only dashboard endpoints. The AuthProvider and system-key request continue using `baseUrl`.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts:19-21`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:23`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\services\system.ts:7-12`
- The Dockerfile does not declare or export `NEXT_PUBLIC_DASHBOARD_AUTO_COLLECTION_API_URL`, so that tested override cannot be supplied through the existing Docker build interface.  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:45-64`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\config.test.ts:95-100`
- `NEXT_PUBLIC_*` values are compiled into browser assets. Changing container runtime environment variables after build will not switch endpoints; environment-specific images/builds are required.
- Docker:
  - Builds on Node 22 Alpine.
  - Uses Yarn v1 lockfile and `--frozen-lockfile`.
  - Builds Next standalone output.
  - Runs production on port 3000 and `0.0.0.0`.
  - Copies `public`, standalone server, and static output.  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:1-17`  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:30-40`  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:66-90`  
  `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock:1-2`
- Registry/security concerns:
  - Private Nexus credentials are passed as Docker `ARG`s and written into `.npmrc`.
  - Password encoding is base64, not encryption.
  - `strict-ssl false` is configured during build.
  - Cleanup occurs after installation, but credential material may remain in intermediate build layers/cache.  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:4-13`  
  `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:24-40`
- Local Yarn always requires authentication against private Nexus; npm uses private registries for `@bri` and `@ordentco`.  
  `D:\work\qcash-ui-dashboard-auto-collection\.yarnrc:1-2`  
  `D:\work\qcash-ui-dashboard-auto-collection\.npmrc:1-3`
- Local `.env` is ignored and excluded from Docker context. It contains credential-like variables with `NEXT_PUBLIC_` names; such names must never hold genuine secrets because public variables are browser-visible. Their values are not repeated here.  
  `D:\work\qcash-ui-dashboard-auto-collection\.gitignore:39-40`  
  `D:\work\qcash-ui-dashboard-auto-collection\.dockerignore:10-12`  
  `D:\work\qcash-ui-dashboard-auto-collection\.env:10-11`
- Loader JSON files exist locally but are ignored by Git; only favicon and locale files are tracked under `public`. A clean checkout will therefore not contain either loader path unless supplied by the host or another pipeline step.  
  `D:\work\qcash-ui-dashboard-auto-collection\.gitignore:55-59`  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:17-25`
- Generated stylesheet bundle and Federation style module are also ignored, but `predev`/`prebuild` regenerates them through the `style` script. Any deployment that bypasses package scripts must run both style generation steps.  
  `D:\work\qcash-ui-dashboard-auto-collection\.gitignore:62-63`  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:5-17`  
  `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs:37-52`
- No CI pipeline, Kubernetes, Helm, Compose, ingress, or route manifests were found. Deployment configuration is limited to the Dockerfile and Next configuration.
- No Next security headers, CSP, redirects, rewrites, asset prefix, or explicit cross-origin configuration are defined.  
  `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:5-64`

## 10. Automated QA coverage and gaps

- Unit/component tests cover:
  - Authorization predicate and page gate.
  - Host router passed to AuthProvider.
  - Hostname/environment API selection.
  - Locale normalization/resource injection.
  - Account loading and Pelindo behavior.
  - Header account/month interactions.
  - Summary selection and loading status.
  - Graph data mapping.
  - Table queries and downloads.
  - Axios token/process-ID/error behavior.
  - Federation fallback and style generation.
  - Development proxy behavior.  
  `D:\work\qcash-ui-dashboard-auto-collection\vitest.config.ts:5-17`
- Tests heavily mock all remote UI components, so they do not validate actual:
  - Authentication-provider integration.
  - Modal/focus semantics.
  - Table keyboard navigation and selection.
  - Date-picker behavior.
  - Localization supplied by host context.
  - Remote loading failures.
  - Responsive layout.
  - Screen-reader output.
  - Mixed-content/CORS behavior.  
  `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.test.ts:18-43`
- There is no browser E2E framework or visual-regression setup in package dependencies.  
  `D:\work\qcash-ui-dashboard-auto-collection\package.json:37-50`
- Sonar is configured for source, script tests, and LCOV coverage, but no CI invocation is present.  
  `D:\work\qcash-ui-dashboard-auto-collection\sonar-project.properties:1-6`

## Highest-priority integrated QA targets

1. Host-mounted dashboard with real shared auth, i18n, QUILA, qcash-global, router, and styles.
2. HTTPS deployment against hardcoded HTTP remotes.
3. Clean-checkout loader behavior under the configured base path.
4. Authorization combinations and backend account/download scope.
5. Indonesian mode, document language, timezone labeling, and missing translations.
6. Keyboard/screen-reader operation of chart, table, autocomplete, date picker, errors, and unauthorized dialog.
7. 320px–768px layouts, summary-card wrapping, table overflow, and action bar.
8. Expired token/401 and blocked localStorage behavior.
9. Docker build with every hostname environment, especially PILOTING and dashboard API override.

---

