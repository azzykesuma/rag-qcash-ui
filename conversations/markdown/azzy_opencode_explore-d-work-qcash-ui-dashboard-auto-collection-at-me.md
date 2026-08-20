# Trace dashboard data flow (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_09097d573ffeODJMvbrnWTU3IQ`
- **Source Tool**: `opencode`
- **Date**: `2026-07-17 16:28:45`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore D:\work\qcash-ui-dashboard-auto-collection at medium thoroughness. Trace exactly how dashboard accounts, summary cards, chart, transaction table, filters, pagination, and downloads currently receive data. Also identify existing HTTP clients and runtime API base URL utilities. Do not edit files. Return relevant files/functions and the smallest integration approach, including any response-shape mismatches with src/services/dashboard.ts.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
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

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 8: Assistant

## Current data flow

All dashboard business data is mocked. No dashboard HTTP request exists.

### Entry/auth

- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx`
  - `DashboardPageView()`
  - Reads `companyID`, `username`, `authorities`, and `guard()` from `useAuth()`.
  - Gates access, then renders `<DashboardPage />`.
  - `companyID` is not passed into the dashboard, so it cannot currently load accounts.

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx`
  - Owns only:
    - `activeType`, default `"payment"`
    - `period`, default `"07"`
  - Selects `mockGraphResponses[activeType]` and `mockTableResponses[activeType]`.

### Accounts/filter header

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx`
  - `accountOptions` is a hard-coded `"all"` option plus ten fake accounts.
  - `DashboardHeader()` keeps `selectedAccounts`, account search, and menu state locally.
  - Account search filters only those hard-coded options.
  - Account selection never leaves this component and therefore does not affect cards, chart, table, pagination, or downloads.
  - Period options are hard-coded to all months of 2026.
  - Only the month string (`"01"`–`"12"`) is sent to the parent.

### Summary cards

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx`
  - Receives `graphResponse.summary`.
  - Uses key mappings from `dashboardContent`.
  - Clicking an amount card changes `activeType`; all six displayed summary values still come from the same response summary object.

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts`
  - Maps payment/hold/release to `total_*` and `count_*`.
  - Formats string amounts with `Number()` and `Intl.NumberFormat`.

### Chart

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx`
  - Receives `mockGraphResponses[activeType]`.
  - Converts `total_amount` strings to numbers.
  - Ignores each item’s `date`; the array index becomes day 1, 2, 3, etc.
  - Period only controls the number of sliced points and labels. Changing month does not fetch or change the mock data, whose dates remain July 2026.
  - Y-axis is fixed at IDR 0–20 billion.

### Transaction table, search, and pagination

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx`
  - Receives `mockTableResponses[activeType]`.
  - Normalizes snake_case and camelCase table fields into QUILA table rows.
  - Adds synthetic `status: "Pending"` and uses payment code as `transactionID`.
  - Search is client-side over every displayed field, including date, amount, and status.
  - `pageIndex` and `pageLimit` only update local state passed to the remote table component.
  - Data is never sliced or refetched. Pagination can change its UI state, but the same ten mock rows remain.
  - `count` is `response.totalData` (`100`) without search, despite only ten rows being available.
  - `key={activeType}` in the parent resets table page/search state when the summary type changes. Period changes do not reset it.

- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx`
  - Defines the local prop contract for the federated QUILA `Table`.
  - The implementation is remote, but callbacks wired locally are `updateIndex`, `updateLimit`, `onSearch`, `downloadFn`, and `downloadFnReceipt`.

### Downloads

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx`
  - Main `downloadFn` is a no-op returning `Promise.resolve()`.
  - Selected-row receipt download:
    1. Receives comma-separated `transactionIDs` from the remote table.
    2. Selects matching rows from the currently loaded response only.
    3. Calls `createDashboardTableCsv()`.
    4. Creates and clicks a browser blob URL.
  - No PDF/XLSX support and no call to `/download`.
  - Cross-page or server-side selected rows cannot be downloaded.

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts`
  - `createDashboardTableCsv()` is the only service-like executable function.
  - Everything else is response types and mock factories.

## Existing HTTP/base URL facilities

- There is no reusable application HTTP client and no `fetch()` call in `src`.
- Axios is only a transitive dependency of `@ordentco/addons-auth-provider`; it is not imported or exposed by this application:
  - `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock`
  - `D:\work\qcash-ui-dashboard-auto-collection\package.json`

- `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts`
  - Exports `{ baseUrl }`, selected at module evaluation from `window.location.hostname`.
  - Mapping:
    - `gym-master`: DEV_ID/DEV_TL/DEV_SG
    - `localhost`: DEV_ID
    - `addons.cms.dev`: PRESTAG
    - `addons.apps`: PREPLOY
    - `demo.addons`: PILOTING
    - `bricams`: PROD_ID/PROD_TL/PROD_SG
    - SSR/unknown hostname: PROD_ID

- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx`
  - Passes `config.baseUrl` only to `AuthProvider`.
  - Dashboard code does not consume it.

- `D:\work\qcash-ui-dashboard-auto-collection\.env`
  - Base URLs already end in `/api`.
  - `NEXT_PUBLIC_API_URL_PILOTING` is referenced by `config.ts` but absent from this file.
  - `NEXT_PUBLIC_API_URL` exists but is unused.

Consequently, documented paths such as `/api/dashboard-auto-collection/data` should be joined as:

```text
${config.baseUrl}/dashboard-auto-collection/data
```

Appending the full documented path would incorrectly produce `/api/api/...`.

## API contract found

`D:\work\qcash-ui-dashboard-auto-collection\confluence-dashboard-tables.json` documents four POST endpoints:

1. `/api/dashboard-auto-collection/accounts`
   - Request: `{ company_id }`
2. `/api/dashboard-auto-collection/summary-graph`
   - Request: `{ account_no, periode, type }`
3. `/api/dashboard-auto-collection/data`
   - Request: `{ account_no, periode, type, filter, limit, page }`
4. `/api/dashboard-auto-collection/download`
   - Request: `{ account_no, periode, type, format }`
   - Response: blob

## Response/request mismatches with `src/services/dashboard.ts`

- No account request/response types exist.
- No download blob type or endpoint function exists.
- No request types or HTTP functions exist for any endpoint.
- `DashboardGraphResponse` matches the documented top-level graph/summary response; it would not match a backend that wraps payload under `data`.
- `DashboardTableResponse` matches `data`, `totalData`, and `totalPage`.
- Table items permit both documented snake_case and camelCase variants, but all identifier/date/name fields are optional, allowing invalid empty items.
- `DashboardTransactionType` uses lowercase values, while requests require `Payment`, `Hold`, or `Release`.
- UI period is only `"07"`; API requires `"07-2026"`.
- Year is hard-coded as 2026 throughout the UI.
- UI account selection is multi-select, while the contract documents one `account_no` or `"All"`. No array/delimiter format is defined.
- API pagination is one-based (`page: 1`); table state starts at zero and must send `pageIndex + 1`.
- Contract prose calls `limit` and `page` strings, but its request example uses numbers.
- API search is limited to entity/payment code, whereas current client search examines every normalized column.
- Chart assumes graph order is complete and chronological instead of using the response `date`.
- The download endpoint supports PDF/XLSX/CSV for a filter scope, but no selected-ID download contract exists.

## Smallest integration approach

1. Keep `src/services/dashboard.ts`; replace mock-only behavior with four small native-`fetch` functions using `config.baseUrl`. No new dependency is needed.
2. Add account and request types there, plus a single lowercase-to-title-case conversion for `type`.
3. Pass `companyID` from `DashboardPageView` into `DashboardPage`; fetch accounts once.
4. Lift selected account state into `DashboardPage` and make `DashboardHeader` controlled.
5. Fetch summary/graph when account, period, or active type changes.
6. Fetch table data when account, period, active type, search, page index, or limit changes; send `pageIndex + 1`.
7. Wire `downloadFn` to the blob endpoint. Keep the current client CSV only for selected rows already loaded.
8. Resolve the multi-account contract before integrating it. The smallest contract-safe behavior is `"All"` or one selected account; aggregating multiple account requests would require correct summary/chart merging and is not a safe UI-only shortcut.

No files were edited.

---

