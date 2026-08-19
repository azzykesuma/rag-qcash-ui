# Map backend integrations (@explore subagent)

- **ID**: `ses_01046c1b7ffefQPoxe6YKbKhgz`
- **Source Tool**: `opencode`
- **Date**: `2026-08-11 14:28:42`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository very thoroughly and map API calls, request/response handling, state management, mutations, downloads/uploads, environment behavior, and integration edge cases needed for a comprehensive manual QA test plan. Include endpoint/method/payload facts where visible and exact file references. Do not edit files.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
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

## Turn 5: Assistant

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
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

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

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

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

# Repository QA map

No files were edited. Findings are based on static inspection; external APIs and federated components were not executed.

## 1. Runtime flow

1. Local page bootstrap wraps the application with the external auth provider, module i18n provider, and QUILA provider:
   - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:15-29`
2. The dashboard calls the external auth `guard()` and waits for:
   - authorities ready
   - non-empty/non-`Guest` username
   - company ID  
   Otherwise it displays the loader indefinitely:
   - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:14-44`
3. Access additionally requires all three exact values:
   - `userType === "cu"`
   - `productAuthorities.DASHBOARD_AUTO_COLLECTION.anyAuthority === true`
   - menu array contains exact string `"Dashboard Auto Collection"`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\utils\auth.ts:1-6`
4. Authorized dashboard initialization:
   1. GET Pelindo company IDs.
   2. POST account list for authenticated company.
   3. Select `"All"` if company is in the Pelindo ID list; otherwise select the first returned account.
   4. POST summary/graph.
   5. Only after graph completion does the transaction table render and request its data.
   - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:52-120`
5. Account, period, transaction type, and refresh changes cancel the previous graph request using `AbortController`.
6. The transaction table is keyed by account/period/type and delegates query lifecycle, pagination, search, selection, and caching behavior to the federated QUILA DataTable:
   - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:114-179`

---

# 2. API inventory

All application API traffic uses Axios except loader assets.

## Shared request behavior

Axios instance:

- `baseURL: config.baseUrl`
- Reads `localStorage["access-token"]` before every request.
- Adds `Authorization: Bearer <token>` when present.
- Adds a UUID as `grpc-metadata-process_id`, unless already supplied.
- Does not set `withCredentials`.
- JSON POST content type is left to Axios defaults.
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:7-22`

### 2.1 Pelindo system key

**Request**

```http
GET {baseUrl}/system/key-secure/acs-company-id-pelindo
Authorization: Bearer <access-token>
grpc-metadata-process_id: <uuid>
```

No payload. An `AbortSignal` is supplied.

**Expected response**

```json
{
  "data": {
    "value": "123467809, 5024345,9876543"
  }
}
```

`value` must be a non-empty string. Missing, empty, or non-string values throw `"Pelindo company ID is unavailable"`.

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\system.ts:3-12`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\system.test.ts:13-24`

The string is interpreted as comma-separated exact company IDs with whitespace trimming:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:19-24`

---

### 2.2 Dashboard accounts

**Request**

```http
POST {dashboardBaseUrl}/dashboard-auto-collection/accounts
Content-Type: application/json
```

```json
{
  "company_id": "5024345"
}
```

**Expected response**

```json
{
  "error": false,
  "code": 200,
  "message": "Success",
  "data": [
    {
      "account_no": "111",
      "account_name": "Operating Account"
    }
  ]
}
```

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:7-17`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:75-76`

The client does not inspect `error`, `code`, or `message` on successful HTTP responses; it directly consumes `data`.

---

### 2.3 Summary and graph

**Request**

```http
POST {dashboardBaseUrl}/dashboard-auto-collection/summary-graph
Content-Type: application/json
```

```json
{
  "account_no": "ALL",
  "periode": "07-2026",
  "type": "payment"
}
```

Rules:

- UI period `YYYY-MM` becomes API period `MM-YYYY`.
- Account `"All"` in any case becomes `"ALL"`.
- Type is lowercase: `payment`, `hold`, or `release`.
- An `AbortSignal` is supplied.

**Expected response**

```json
{
  "error": false,
  "code": 200,
  "message": "Success",
  "graph": [
    {
      "total_amount": "1000000000",
      "date": "01-07-2026"
    }
  ],
  "summary": {
    "total_payment": "1000000000",
    "count_payment": 1,
    "total_hold": "0",
    "count_hold": 0,
    "total_release": "0",
    "count_release": 0
  }
}
```

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:27-47`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:79-80`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:56-65`

---

### 2.4 Paginated transaction data

**Request**

```http
POST {dashboardBaseUrl}/dashboard-auto-collection/data
Content-Type: application/json
```

```json
{
  "account_no": "ALL",
  "periode": "07-2026",
  "type": "payment",
  "filter": "pelindo",
  "limit": 25,
  "page": 3
}
```

Rules:

- `"All"` maps to `"ALL"`.
- Search text is forwarded directly as `filter`.
- Pagination is server-side.
- Sorting supplied by QUILA is ignored.
- Page size options are 10, 20, and 50; initial size is 10.
- An `AbortSignal` is supplied.

**Expected response**

```json
{
  "error": false,
  "code": 200,
  "message": "Success",
  "data": [
    {
      "date": "23-07-2026",
      "entity_name": "Pelindo",
      "payment_code": "PAY-1",
      "amount": "1250.5"
    }
  ],
  "totalData": 1,
  "totalPage": 1
}
```

The client accepts both naming styles:

- `date` or `transactionDate`
- `entity_name` or `entityName`
- `payment_code` or `paymentCode`

Missing text fields become `"-"`. Amount is converted with `Number()`.

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:49-66`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:83-88`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:36-42`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:114-139`

`totalPage` is ignored; only `totalData` is returned to QUILA.

---

### 2.5 Download

**Request**

```http
POST {dashboardBaseUrl}/dashboard-auto-collection/download
Content-Type: application/json
Accept: application/octet-stream
```

Axios uses `responseType: "blob"`.

Normal complete-account example:

```json
{
  "account_no": "111",
  "periode": "07-2026",
  "type": "Payment",
  "format": "xlsx"
}
```

Selected rows:

```json
{
  "account_no": "111",
  "periode": "07-2026",
  "type": "Payment",
  "format": "pdf",
  "ids": ["PAY-1", "PAY-2"]
}
```

Special rules:

- Types become `Payment`, `Hold`, or `Release`.
- UI `xlsx` becomes internal `xls`, then API `xlsx`.
- API formats are `csv`, `pdf`, `xlsx`.
- When selected IDs are present and account is `"All"`, `account_no` is omitted entirely.
- When no IDs are supplied and account is `"All"`, the literal value sent is `"All"`, not `"ALL"`.
- No download cancellation signal or progress callback exists.

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:91-110`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:63-82`

**Expected response**

- Blob body
- Required `Content-Disposition` filename, supporting:
  - `filename="report.xlsx"`
  - `filename*=UTF-8''encoded-name.xlsx`

Missing filename throws `"Download filename is missing from the response"`.

The browser creates an object URL, clicks a detached anchor with its `download` property, then immediately revokes the URL:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:27-34`

For cross-origin APIs, `Content-Disposition` must be exposed through CORS or the download will be treated as missing a filename.

---

### 2.6 Loader assets

Native fetch, not Axios:

```http
GET /v2/lottie/qlola-loader.json
```

If the first response is non-2xx:

```http
GET /qlola-loader.json
```

Any final failure is silently ignored and the CSS spinner remains.

Reference:

- `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\splash-screen\new-qlola.tsx:17-25`

These paths are absolute host-root paths and do not include the application base path.

---

# 3. Response and error handling

## Shared Axios error interceptor

A persistent QUILA danger toast is shown only when all conditions match:

- Request is GET, or URL contains `/dashboard-auto-collection/data`.
- Response has a numeric status.
- Status is not 401.
- URL is not excluded.

The message is `response.data.message` when non-empty, otherwise `"Failed to load data."`.

Process ID lookup order:

1. Response `grpc-metadata-process_id`
2. Request `grpc-metadata-process_id`

The toast:

- stays open indefinitely (`timeout: 0`)
- appears at top-end
- includes `Process ID`
- offers a Copy action
- falls back from `navigator.clipboard.writeText` to hidden textarea plus `document.execCommand("copy")`

References:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:24-38`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:3-59`

No shared toast is produced for:

- 401
- accounts POST failure
- summary/graph POST failure
- download POST failure
- network errors without an HTTP status
- aborted requests

Excluded URL fragments are `/v2/locales/`, `/v2/lottie/`, and `/qlola-loader.json`. Exclusion prevents process ID creation and error toasts, but would not prevent Authorization header injection.

## Local UI errors

| Failure | UI result |
|---|---|
| Pelindo GET | 3-second failure Message; dashboard stops |
| Accounts POST, authorized Pelindo company | Message, then continues with `"All"` |
| Accounts POST, non-Pelindo company | Message plus chart/table error cards; no graph request |
| Graph POST | Message plus chart/table error cards; no table data request |
| Table data POST | Persistent interceptor toast plus table error card |
| Download POST | 3-second Message; parses Blob JSON `message` when possible |
| Remote toast module unavailable | Error state remains, toast failure is swallowed |

Container message:

- right-top
- `qCashFailed`
- duration 3000 ms
- close icon
- a single shared string state, so later errors replace earlier ones

Reference:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:50-50`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:162-162`

Download errors append only the **request** process ID. They do not inspect the response process-ID header:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:68-80`

---

# 4. State and interaction behavior

No custom global store exists. State is divided among:

- external AuthProvider
- i18next provider
- local React state in the dashboard
- local download/table-error state
- opaque state/cache inside federated QUILA DataTable

## Dashboard state

Initial values:

- type: `payment`
- period: current browser month, `YYYY-MM`
- account: empty
- graph: null
- loading: true
- last updated: null
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:31-50`

### Trigger matrix

| User event | Network/state effect |
|---|---|
| Company ID changes | Reset accounts/account/Pelindo ID/graph; reload Pelindo ID and accounts |
| Account changes | Cancel old graph; request graph; remount table after graph settles |
| Period changes | Same as account |
| Payment/Hold/Release card | Same as account |
| Refresh | Increment counter and reload graph; controls disable refresh while loading |
| Table search | Server table-data POST |
| Page/page size | Server table-data POST |
| Row selection | Local QUILA selection only |
| Download | POST download and trigger browser save |

Graph effect cancellation:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:96-120`

Table query key:

```json
["dashboard-auto-collection", "<account>", "<MM-YYYY>", "<type>"]
```

Reference:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:122-135`

Refresh count is not part of this key. The table currently disappears while graph reload state is applied and should remount afterward; verify this with the real federated DataTable because its cache behavior is external.

## Account filter

- Local case-insensitive substring filtering.
- Search input is trimmed.
- Authorized companies receive an `"All"` option before API accounts.
- Clearing/null selection falls back to the first option rather than remaining empty.
- Non-authorized companies default to the first API account.
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:20-26`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:58-99`

## Period filter

- Month picker.
- Current month through 12 months prior are enabled, giving 13 inclusive month choices.
- Future months are disabled.
- Clearing returns to current month.
- API receives `MM-YYYY`.
- Display range is `01/MM/YYYY - last-day/MM/YYYY`, including leap years.
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:27-33`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-header.tsx:103-131`

The period picker uses two potentially different date shapes:

- disabled callback calls `date.year()` and `date.month()`
- change callback accepts `Date` or `{year, month}`

The actual QUILA version must be verified for zero-based versus one-based month behavior.

## Last updated

- Updated only after successful graph response.
- Formatted with browser-local time using `en-GB`, then labeled `"WIB"`.
- It does not explicitly use `Asia/Jakarta`.
- It persists through subsequent failures.
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:23-32`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:105-109`

---

# 5. Display transformations and edge cases

## Currency

Always uses Indonesian number separators, regardless of UI language:

```text
1250.5 -> IDR 1.250,50
```

Invalid numeric values become `IDR 0,00`.

Reference:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\dashboard-config.ts:43-54`

## Chart

- Produces one point per day in selected month.
- Missing days become zero.
- Accepts date shapes resembling `DD-MM-YYYY` or `YYYY-MM-DD`.
- Invalid dates fall back to array index + 1.
- Duplicate days overwrite earlier values.
- Out-of-month points are discarded.
- Invalid amount becomes zero.
- Y-axis minimum maximum is IDR 5 billion.
- Tick step is IDR 5 billion.
- Tooltip does not appear for zero because zero is treated as falsy.
- Tooltip divides by one billion and labels the value `M`.
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:23-35`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx:60-90`

## Table

- Date is displayed unchanged.
- Row identity is payment code.
- Missing payment code becomes `"-"`, potentially causing duplicate row IDs and incorrect selection.
- Duplicate payment codes have the same risk.
- Selected downloads pass payment codes as `ids`.
- Selection scope is current page.
- When DataTable reports selection as `"all"`, only `visibleSelectedKeys` are used.
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:117-139`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:142-170`

Complete download ignores current table search, page, and visible rows. It requests the entire account/period/type dataset.

---

# 6. Mutations, uploads, and downloads

## Mutations

There are no create/update/delete business mutations. All business POST calls are read-like queries or downloads:

- accounts
- summary graph
- table data
- download

## Uploads

No upload implementation exists:

- no file input
- no `FormData`
- no multipart request
- no upload progress or size/type validation

Static upload-instruction assets under `public\img\account-receivable` are unrelated to this dashboard's runtime code.

## Downloads

Test all combinations:

- complete CSV/XLSX/PDF
- selected CSV/XLSX/PDF
- account-specific versus All
- empty selection
- one and many selections
- selections across pages
- duplicate/missing payment codes
- large files
- repeated clicks
- filenames with spaces, Unicode, percent encoding, semicolons, and missing header
- HTTP 400 JSON Blob
- HTTP 401
- HTTP 500 non-JSON Blob
- network timeout/disconnect

---

# 7. Environment behavior

Environment selection occurs at module initialization using `window.location.hostname`:

- `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts:1-21`

| Hostname condition | Base variable |
|---|---|
| SSR/no `window` | `NEXT_PUBLIC_API_URL_PROD_ID` |
| contains `gym-master` and `tl` | `NEXT_PUBLIC_API_URL_DEV_TL` |
| contains `gym-master` and `sg` | `NEXT_PUBLIC_API_URL_DEV_SG` |
| other `gym-master` | `NEXT_PUBLIC_API_URL_DEV_ID` |
| contains `localhost` | `NEXT_PUBLIC_API_URL_DEV_ID` |
| contains `addons.cms.dev` | `NEXT_PUBLIC_API_URL_PRESTAG` |
| contains `addons.apps` | `NEXT_PUBLIC_API_URL_PREPLOY` |
| contains `demo.addons` | `NEXT_PUBLIC_API_URL_PILOTING` |
| contains `bricams` and `tl` | `NEXT_PUBLIC_API_URL_PROD_TL` |
| contains `bricams` and `sg` | `NEXT_PUBLIC_API_URL_PROD_SG` |
| other/unknown host | `NEXT_PUBLIC_API_URL_PROD_ID` |

`dashboardBaseUrl` uses:

```text
NEXT_PUBLIC_DASHBOARD_AUTO_COLLECTION_API_URL ?? baseUrl
```

Important environment findings:

1. `.env` contains DEV ID/TL/SG, PRESTAG, PREPLOY, and PROD ID/TL/SG URLs:
   - `D:\work\qcash-ui-dashboard-auto-collection\.env:1-15`
2. `.env` does not define `NEXT_PUBLIC_API_URL_PILOTING`; a local `demo.addons` run can therefore produce an undefined base URL.
3. `.env` does not define the dashboard-specific override.
4. `NEXT_PUBLIC_API_URL` exists but is not consumed by `src\config.ts`.
5. The dashboard override is not included among Docker build arguments/environment variables:
   - `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:45-64`
6. `.env` files are excluded from Docker context:
   - `D:\work\qcash-ui-dashboard-auto-collection\.dockerignore:10-12`
7. `NEXT_PUBLIC_*` values must be correct at build time; changing container runtime variables generally will not rewrite an already-built client bundle.
8. An empty dashboard override is preserved by `??`, causing dashboard services to throw `"Dashboard API URL is not configured"`.
9. Several checked-in DEV/PREPLOY URLs use HTTP. Serving the UI over HTTPS can result in browser mixed-content blocking.
10. Custom Authorization and process-ID headers require correct CORS preflight allowances. Custom response headers must be exposed for filename/process-ID handling.

---

# 8. Module Federation and host integration

## Application identity and routes

- Base path: `/dashboard-auto-collection-v2`
- Local page: `/dashboard-auto-collection-v2/dashboard`
- Remote entry: `static/chunks/remoteEntry.js`
- Exposed module: `./dashboard`
- Exposed styles: `./globals`
- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:5-7`
- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:33-55`

## Remote dependencies

Build configuration hardcodes HTTP DEV-cluster entries for:

- `qui`
- `qcash-global-component`

Reference:

- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:38-48`

If no existing federation runtime is found, the runtime fallback points to localhost:

- QUILA: `http://localhost:3000/design-system/static/remoteEntry.js`
- global component: `http://localhost:3000/global-component/_next/static/chunks/remoteEntry.js`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:15-66`

Integration tests must cover:

- host-provided runtime versus standalone fallback
- unavailable remote entry
- null/missing module
- wrong shared React version
- duplicate React/auth/i18n instances
- HTTPS host loading HTTP remotes
- production environment accidentally loading DEV remotes
- delayed remotes and temporary blank dynamic components
- host navigation away and back

## Auth-provider integration

When run as a normal Next page, `_app.tsx` supplies AuthProvider. When the exposed `./dashboard` module is mounted directly by a host, `_app.tsx` is not necessarily part of that render tree. The host must provide the shared auth context expected by `useAuth()`.

## Styles

The dashboard attempts to preserve and reappend an existing style element with ID:

```text
qcash-ui-dashboard-auto-collection-globals
```

It can restore CSS only after it has previously captured the style text:

- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:10-30`
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\global.exposes.ts:1-24`

Test first-load behavior when the host does not load `./globals`, navigation that removes styles, duplicate style IDs, and CSS ordering conflicts.

## Local host-compatibility proxy

Recommended local command is `yarn proxy`:

- listens on port 3001 by default
- starts Next on 3002
- strips `origin`, `referer`, and `sec-fetch-*` only for `/_next` and `/__nextjs`
- preserves those headers for other paths
- returns 502 when upstream fails
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\dev-host-compat.mjs:46-103`
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\dev-host-compat.mjs:118-191`
- `D:\work\qcash-ui-dashboard-auto-collection\README.md:19-35`

---

# 9. Internationalization

- Bundled languages: Indonesian and English.
- Default/fallback: English.
- Reads `localStorage["locale"]`.
- Accepts normalized variants such as `ID_id` and `en-US`.
- Unsupported or inaccessible storage falls back to English.
- Resources are compiled into the bundle; runtime locale-file fetch is not required.
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:9-49`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\i18n.ts:51-78`

QA considerations:

- English and Indonesian labels, tooltips, empty states, errors, date-picker locale, and month labels.
- Locale changed before mount versus after mount.
- Blocked local storage.
- Unsupported locale such as `fr`.
- Interceptor fallback errors and Copy/Copied actions remain hardcoded English.
- Indonesian chart/table error titles are also currently English:
  - `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\dashboard-auto-collection.json:63-69`
- Currency remains Indonesian-formatted in both languages.
- Last-updated numeric format is always `en-GB`.

---

# 10. Comprehensive manual QA checklist

## P0 — Authentication and authorization

1. Valid customer user with exact authority/menu reaches dashboard.
2. User type other than `cu` receives Unauthorized.
3. Missing authority, false authority, missing menu, or differently cased menu receives Unauthorized.
4. Missing company ID, empty username, `Guest`, or authorities not ready stays on loader.
5. Verify `guard()` behavior for expired session and redirects.
6. Verify 401 API behavior is owned correctly by AuthProvider, since this application suppresses its own 401 toast.
7. Change/remove `access-token` between requests and verify each new request reads the current value.
8. Block localStorage and verify whether request setup fails cleanly.

## P0 — Successful data flow

1. Verify request order: Pelindo GET → accounts POST → graph POST → data POST.
2. Pelindo company defaults to All.
3. Non-Pelindo company defaults to first account.
4. Account search is local and causes no API call until selection.
5. Verify payment, hold, and release payloads and displayed fields.
6. Verify account/period/type changes cancel stale graph/table requests.
7. Ensure late stale responses cannot overwrite current filters.
8. Verify manual refresh reloads graph and table and disables duplicate refresh.
9. Verify last-updated changes only after successful graph refresh.
10. Verify loading skeleton and no partial stale summary during transitions.

## P0 — Error handling

For each endpoint, test 400, 401, 403, 404, 500, 503, timeout, offline, malformed JSON, and abort where applicable.

1. Pelindo failure: toast plus account-load Message; no account/graph requests.
2. Accounts failure for Pelindo: fallback to All and continue.
3. Accounts failure for non-Pelindo: graph/table error states and no graph request.
4. Graph failure: chart/table error states and no table request.
5. Table failure: persistent toast, process ID Copy, and table alert.
6. Retry table failure via refresh and filter changes.
7. API response message empty/missing should use fallback.
8. Process ID in response header versus only request header.
9. CORS-hidden process-ID header.
10. Toast remote unavailable.
11. Clipboard permission denied and HTTP fallback.
12. Repeated identical error messages while previous Message is still open.

## P0 — Downloads

1. Complete and selected CSV/XLSX/PDF.
2. Verify exact payload casing and `xls`→`xlsx` conversion.
3. Verify All/no-selection sends `"account_no": "All"`.
4. Verify All/selected IDs omits `account_no`.
5. Verify selected account sends both account and IDs.
6. Validate download honors expected scope when the table is searched.
7. Verify selection across pagination uses only current-page IDs.
8. Missing, duplicate, or empty payment codes.
9. Missing `Content-Disposition`.
10. Non-exposed `Content-Disposition`.
11. Quoted, UTF-8, encoded, malformed, and very long filenames.
12. Blob JSON error message and malformed/non-JSON error body.
13. Response process ID versus request process ID.
14. Large files, slow connections, multiple clicks, browser URL revocation timing.
15. Confirm downloaded extension and content match requested format.

## P1 — API response robustness

1. HTTP 200 with `error: true`.
2. HTTP 200 with missing `data`, `summary`, `graph`, or `totalData`.
3. Empty account list for non-Pelindo: current code can remain in loading state because no account is selected.
4. Empty account list for Pelindo: should still request All.
5. Null account entries or duplicate account numbers.
6. String versus numeric company IDs and Pelindo values.
7. Whitespace/case variants of `"All"`.
8. Amounts: zero, negative, decimals, invalid strings, null, extremely large numbers.
9. Both snake_case and camelCase table fields.
10. Duplicate and malformed chart dates.
11. Leap year, 28/29/30/31-day months.
12. Graph points outside selected month.
13. Summary totals inconsistent with graph totals.

## P1 — Interaction/UI

1. Account clear behavior selects first option.
2. Long account numbers/names and many accounts.
3. Current month, exactly 12 months prior, future month, and older month.
4. Date picker month numbering across year boundaries.
5. Browser timezone outside WIB; check mislabeled last-updated time.
6. Module loaded before and after midnight/month rollover.
7. Keyboard operation for cards, account filter, month picker, table, selection, and downloads.
8. Screen-reader labels, pressed card states, loading live region, error alerts, and focus visibility.
9. Small/mobile viewport; summary cards currently remain a horizontal flex row.
10. Tooltip keyboard/touch access and long translated text.
11. No-data table and zero-value chart.

## P1 — Environment and federation

1. Every hostname branch in `src\config.ts`.
2. Unknown hostname defaults to PROD ID.
3. Dashboard-specific override and empty override.
4. Missing PILOTING value.
5. HTTPS host with HTTP API/remotes.
6. CORS preflight for Authorization and process-ID headers.
7. Production build with all build-time variables.
8. Standalone Next route under base path.
9. Loader asset requests under base path and host root.
10. Existing host federation runtime versus localhost fallback.
11. Missing QUILA/global-component remotes.
12. Direct mounting of exposed dashboard with and without host AuthProvider.
13. Host navigation that removes dashboard CSS.
14. Multiple MFE instances on the same page.
15. Local proxy HMR, assets, 502 path, and host cross-origin embedding.

## Known external blind spots

The repository does not contain implementations for AuthProvider, QUILA DataTable/DatePicker, Unauthorized dialog, Message, or other federated components. Their retry, debounce, cache, session-expiry, pagination, and accessibility behavior must be verified in the fully integrated host rather than inferred from this repository.

---

