# Map user-facing flows (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_01046c1dbffeovCr7tkII3VZ3H`
- **Source Tool**: `opencode`
- **Date**: `2026-08-11 14:28:42`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository very thoroughly and map every user-facing app flow for a manual QA test plan. Focus on pages, components, forms, tables, filters, pagination, dialogs, upload/download, validation, navigation, responsive states, and success/error/empty/loading states. Do not edit files. Return concise structured findings with exact source file references and any observable expected behavior.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

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

<details><summary>🔧 Tool Invocations (12)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
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

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

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

## Turn 7: Assistant

# Repository flow map

## 1. Routes and entry states

### R1 — Placeholder root page
- **Route:** `/dashboard-auto-collection-v2`
- Displays only `Welcome to qcash-ui-dashboard-auto-collection`.
- No links, forms, or navigation.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\pages\index.tsx:1-3`
- Base path: `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:5-7`

### R2 — Dashboard
- **Standalone route:** `/dashboard-auto-collection-v2/dashboard`
- Also exposed to a host through Module Federation as `./dashboard`.
- Sources:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:14-45`
  - `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:50-55`
  - `D:\work\qcash-ui-dashboard-auto-collection\README.md:19-35`

### Authentication and authorization matrix

| Condition | Expected result |
|---|---|
| Authorities not ready, missing username/company, or username `Guest` | Qlola loading screen remains visible |
| Auth complete but user is not `cu` | Unauthorized dialog |
| Missing `DASHBOARD_AUTO_COLLECTION.anyAuthority === true` | Unauthorized dialog |
| Menu list omits exact text `Dashboard Auto Collection` | Unauthorized dialog |
| All requirements met | Dashboard renders with authenticated `companyID` |

Sources:
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:15-44`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\auth.ts:1-6`

Loader behavior:
- Initially uses a local spinning circle, then replaces it with Lottie after `/v2/lottie/qlola-loader.json` or fallback `/qlola-loader.json` loads.
- Content loader height is `calc(100vh - 160px)`.
- Lottie has `role="status"` and `aria-label="Loading"`; the fallback spinner has no status text.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:11-65`

The unauthorized dialog is a remote component with no local text or button behavior configured; validate it against the deployed global-component remote:
- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:99-102,173`

---

## 2. Dashboard initialization

### D1 — Normal initial load
1. Dashboard defaults to:
   - Current calendar month.
   - Transaction type `payment`.
   - No account until account initialization completes.
2. It requests the Pelindo company-ID configuration.
3. It requests accounts using authenticated `companyID`.
4. Default account:
   - Authorized Pelindo company: `All`.
   - Other company: first returned account.
5. It requests summary/graph data.
6. On success:
   - Three populated summary cards appear.
   - Trend chart and transaction table appear.
   - “Last updated” changes from `-` to client date/time with a hard-coded `WIB` suffix.

Sources:
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:31-50,52-120,140-162`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\system.ts:7-12`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:19-25,75-88`

API contract:
- Accounts: `POST .../dashboard-auto-collection/accounts`, `{ company_id }`
- Graph: `POST .../dashboard-auto-collection/summary-graph`, `{ account_no, periode: "MM-YYYY", type }`
- `All` is normalized to `ALL` for graph/table requests.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:68-88`

### D2 — Initial loading state
- Header is present.
- Refresh is disabled and its icon spins.
- Three cards are shown as nine shimmer skeletons in a live status region.
- Chart and table are not rendered until graph loading ends.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:140-160`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:13-37`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:39-54`

### D3 — Refresh
- Refresh reruns only the graph request for the current account, month, and type.
- During refresh, prior chart/table content is removed and summary skeletons return.
- Refresh cannot be clicked again until loading ends.
- Successful refresh updates the timestamp.
- Failed refresh retains the previous timestamp but displays zero-value cards and error panels.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:96-120,145-160`

---

## 3. Account and period filters

### F1 — Account selector
- Display format: `account_no - account_name`.
- Authorized Pelindo companies additionally get `All` as the first option.
- Selecting an account reloads summary, chart, and table.
- Clearing/resetting selection chooses the first option:
  - `All` for authorized companies.
  - First account otherwise.
- Popover search:
  - Case-insensitive substring match.
  - Searches the complete number/name label.
  - Trims leading/trailing search input.
  - Local filtering; it does not call an account-search API.
- Popover width is 430px, capped at viewport width minus 40px.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:20-26,57-100`

### F2 — Month selector and validation
- Month-only picker.
- Current month is selected by default.
- Allowed range is inclusive from the same month one year ago through the current month.
- Future months and months older than that are disabled.
- Clearing the picker resets to the current month.
- Selection reloads all dashboard data using `MM-YYYY` API format.
- Locale controls month labels:
  - `id-ID` for Indonesian.
  - `en-US` otherwise.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:27-33,102-131`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:56-65`

No other client-side form validation or form submission exists.

---

## 4. Summary-card flows

Three cards always appear in this order:

1. Paid Funds / Dana Terbayar
2. Hold Funds / Dana Ditahan
3. Released Funds / Dana Dirilis

Expected behavior:
- `payment` is active initially.
- Active card has dark-blue background and white content.
- Clicking another card:
  - Marks it active with `aria-pressed="true"`.
  - Reloads graph and table for that type.
  - Changes chart/table headings and amount-column heading.
- Each card shows:
  - Type title.
  - IDR total with Indonesian separators and two decimals.
  - Item count.
  - Information button with explanatory tooltip.
- Invalid/null amounts display `IDR 0,00`.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:40-129`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:3-53`

There is no separate summary empty state; missing/error summary data becomes three zero-value cards.

---

## 5. Trend chart

### C1 — Successful chart
- Heading and total change with the active card.
- Period is rendered as `01/MM/YYYY - last-day/MM/YYYY`, including leap years.
- One chart point is generated for every day in the selected month.
- Missing dates are filled with zero.
- Supports backend dates in `DD-MM-YYYY` and `YYYY-MM-DD`; malformed dates fall back to array position.
- Y-axis:
  - Right aligned.
  - Billions of IDR.
  - Minimum maximum is 5 billion.
  - Grows in 5-billion increments.
- Hovering a non-zero point shows day/month and `IDR n M`.
- Zero values do not produce a tooltip.
- Empty graph arrays produce a zero-valued chart, not an empty-state message.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:23-35,60-138`

### C2 — Chart error
- Displays a danger alert containing:
  - `Error while fetching chart data`
  - Localized “Failed to load dashboard data.”
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:41-59`

---

## 6. Transaction table

### T1 — Columns and data mapping

| Column | Expected value |
|---|---|
| Transaction Date | `date`, then `transactionDate`, otherwise `-` |
| Entity Name | `entity_name`, then `entityName`, otherwise `-` |
| Payment Code | `payment_code`, then `paymentCode`, otherwise `-` |
| Amount | Right-aligned, bold, Indonesian IDR format |

Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:18-61`

### T2 — Search and pagination
- Search field is enabled with localized “Search in table/Cari di tabel”.
- Initial page size: 10.
- Available page sizes: 10, 20, 50.
- Search text, page number, and page size are delegated to the API.
- `totalData` controls pagination totals; backend `totalPage` is ignored.
- Changing account, period, or type creates a fresh table instance.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:114-140`

Configured request:
```text
{ account_no, periode, type, filter: search, limit: pageSize, page }
```

Sorting is present in the remote table request type but ignored by this application. If sortable headers appear from the remote component, sorting will not alter the backend request:
- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:38-43`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:122-134`

### T3 — Empty and loading states
- Empty/loading UI is owned by the remote QUILA `DataTable`, not implemented locally.
- Local resources supply:
  - `No Data` / `Tidak Ada Data`
  - Rows-per-page and showing-count strings.
- Validate actual visuals against the deployed design-system version.
- Sources:
  - `D:\work\qcash-ui-dashboard-auto-collection\public\locales\en\common.json:5-13`
  - `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\common.json:5-13`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:55-73`

### T4 — Table failure
- A non-aborted request failure replaces the entire table with a danger alert:
  - `Error while fetching transaction table`
  - Localized “Failed to load transaction details.”
- There is no local retry control.
- Manual dashboard refresh does not reset the table’s internal error state when account/month/type remain unchanged.
- Non-401 `/data` failures can additionally show a persistent global danger toast with API message and process ID.
- Sources:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:84-104,122-133`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:24-38`

---

## 7. Selection and download flows

### DL1 — Complete-table download
- Remote table is configured for CSV, XLSX, and PDF.
- Selecting its normal download action requests the entire filtered context without row IDs.
- UI `xlsx` maps to backend format `xlsx` through the internal `xls` type.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:63-82,114-121`

### DL2 — Selected-row download
- Multiple selection is enabled, scoped to the current page.
- Row identity is the payment code.
- Selection action bar defaults to CSV and offers CSV/XLSX/PDF.
- Clicking “Download Table” sends selected payment-code IDs.
- If account is `All` and IDs are supplied, `account_no` is omitted from the download request.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:117-121,137-178`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:91-104`

### DL3 — Successful download
- Browser download uses the backend `Content-Disposition` filename.
- Both plain `filename=` and UTF-8 `filename*=` are supported.
- No success message or progress/disabled state is provided.
- Repeated clicks can initiate repeated requests.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:105-110`

### DL4 — Download failure
- Blob JSON errors use the backend `message`; malformed/missing messages use the localized generic failure.
- Request process ID is appended when available.
- Displays a closable failed message at the upper right for 3 seconds.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:68-81`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:162`

There is no upload flow or file-input component.

---

## 8. Error-state matrix

| Failure | Expected observable result |
|---|---|
| Pelindo ID GET fails | Account failure message; zero cards; chart and table error alerts |
| Pelindo ID value missing/invalid | Same local state, without interceptor toast |
| Accounts fail for authorized Pelindo company | Account failure message, but dashboard continues using `All` |
| Accounts fail for other company | Account failure message plus chart/table errors; no graph request |
| Graph fails | Zero cards, chart/table errors, 3-second dashboard failure message |
| Table data fails | Table danger alert; possible persistent process-ID toast |
| Download fails | 3-second upper-right failed message |
| 401 response | Global interceptor toast is suppressed |
| Optional toast remote fails | Error panel/message remains; toast failure is silently ignored |

Sources:
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:52-120`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:24-38`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:27-59`

Persistent global toast:
- Displays API message.
- Shows process ID.
- “Copy” copies it and changes to disabled “Copied”.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:27-56`

Important edge case:
- A non-Pelindo company receiving a successful but empty accounts array remains on the summary skeleton indefinitely because no account is selected and loading is never completed.
- Source: `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:74-80,96-99`

---

## 9. Responsive QA

Recommended widths: 375px, 639px, 640px, 767px, 768px, and desktop.

Expected source-defined behavior:
- Below `sm`:
  - Title stacks above right-side controls.
  - Account and period controls still remain in one horizontal row.
  - Account field shrinks; period field remains approximately 220px.
- At `sm` and above:
  - Title and controls form a horizontal header.
  - Account field is fixed around 220px.
- Summary cards always remain a single non-wrapping flex row; verify text clipping/overflow on narrow screens.
- Chart is fluid-width with fixed 330px height.
- Chart/table horizontal padding changes from 20px to 40px at `md`.
- Tooltip width is capped to viewport minus 2rem.
- Table responsiveness, overflow, and mobile pagination are delegated to the remote QUILA component.

Sources:
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:35-38,56-57,80,102`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\summary-cards.tsx:43-64,76`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:93-104`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:107-114`

---

## 10. Localization QA

- Supported languages: English and Indonesian.
- Default: English.
- Language is read from local storage key `locale`; variants such as `id-ID` and `en_US` normalize correctly.
- There is no language-switch control in this module.
- Currency always uses Indonesian numeric formatting regardless of language.
- Document `<html lang>` remains `en` even when Indonesian is active.
- Source:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:9-49,71-78`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:43-53`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_document.tsx:3-11`

---

## 11. Explicitly absent or remote-owned flows

- No create/edit forms.
- No field-required validation.
- No upload.
- No internal route navigation, breadcrumbs, back button, or links.
- No confirmation dialog.
- No success toast.
- Only dialog is the remote unauthorized dialog.
- Account, date, download-format popovers and most table states are remote design-system components.
- Column settings are explicitly disabled.
- Sources:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:140`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx:114-173`

No files were edited.

---

