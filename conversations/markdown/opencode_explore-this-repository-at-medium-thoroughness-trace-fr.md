# Trace API error handling (@explore subagent)

- **ID**: `ses_02a3377c7ffegcgn1rH9KbWken`
- **Source Tool**: `opencode`
- **Date**: `2026-08-06 13:39:39`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore this repository at medium thoroughness. Trace the frontend execution path for GET /api/dashboard-auto-collection/data from request creation through error handling and user-visible message. Find all shared API/error helpers and any existing handling for process ID/request ID/correlation ID, including likely response body fields and headers. Do not edit files. Return exact file paths, relevant line numbers, function names, callers, and the smallest correct edit location.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

## Key finding: method mismatch

There is no frontend `GET /api/dashboard-auto-collection/data` implementation or local Next.js API route. The current working tree sends this request as **POST**:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:80`
- Confirmed by test: `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.test.ts:107-125`

The `/api` prefix comes from `dashboardBaseUrl`; tests configure it with a value ending in `/api` at `dashboard.test.ts:108`. Runtime selection is at:

- `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts:1-19`

There are no files under `src/pages/api` or `src/app/api`, and no Next.js rewrites in `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts`.

## Execution path

1. **Dashboard mounts after authentication**
   - `DashboardPageView()`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx:14-44`
   - Renders `DashboardPage` after auth checks at line 44.

2. **Transaction table is mounted**
   - `DashboardPage()`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx:151-153`
   - Passes `setMessage` as `TransactionTable.onError`.

3. **Remote `DataTable` invokes its supplied query**
   - `TransactionTable()`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:116-139`
   - `queryFn` receives page, page size, search text, and an abort signal.
   - Caller at line 126:
     ```ts
     getDashboardData(accountNo, getApiPeriod(period), activeType, search, pageSize, page, signal)
     ```

4. **Service creates the API payload**
   - `getDashboardData()`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:100-102`
   - Body fields:
     - `account_no`
     - `periode`
     - `type`
     - `filter`
     - `limit`
     - `page`
   - `"All"` is normalized to `"ALL"`.

5. **Shared dashboard request wrapper sends POST**
   - `dashboardRequest<T>()`
   - `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:77-82`
   - Throws `"Dashboard API URL is not configured"` before Axios if no URL.
   - Sends:
     ```text
     POST {dashboardBaseUrl}/dashboard-auto-collection/data
     ```

6. **Shared request interceptor adds headers**
   - `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:15-26`
   - Adds:
     - `Authorization: Bearer <access-token>` when available.
     - `grpc-metadata-process_id: <UUID v4>`, unless explicitly supplied.
   - Also assigns the value to the Axios request config as `processId` at line 22.
   - Exclusions are defined at lines 9-12 for locale/Lottie requests.

7. **Successful response**
   - `dashboardRequest()` returns `response.data` at `dashboard.ts:81`.
   - `TransactionTable.queryFn` maps rows and returns `{ rows, total }` at:
     - `transaction-table.tsx:127-130`

## Error path actually taken by `/data`

### Shared Axios interceptor

`customAxios` has a response interceptor at:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:28-40`

However, line 30 only displays shared errors when:

- Method is `GET`
- A numeric HTTP status exists
- Status is not `401`
- URL is not excluded

Because `/dashboard-auto-collection/data` is currently **POST**, the toast/process-ID branch at lines 31-36 is skipped. The interceptor only re-rejects the error at line 39.

Other cases that skip the shared toast are:

- `401`
- Network errors without `error.response.status`
- Aborted requests
- Excluded URLs

### Component-level handling

The rejected POST reaches:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.tsx:131-137`

For non-aborted requests it:

1. Sets `tableError` to `true`.
2. Calls `onError()` with translated `dashboardAcs.error.loadDetailsFailed`.
3. Rethrows the original error.

This produces two user-visible generic messages:

- **Persistent inline danger alert**
  - `transaction-table.tsx:86-105`
  - Title: `dashboardAcs.error.loadTableFailed`
  - Description: `dashboardAcs.error.loadDetailsFailed`

- **Three-second failed Message**
  - State is set through `onError={setMessage}` at `dashboard/index.tsx:152`
  - Rendered at `dashboard/index.tsx:154`

Translations:

- English: `D:\work\qcash-ui-dashboard-auto-collection\public\locales\en\dashboard-auto-collection.json:63-69`
- Indonesian: `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\dashboard-auto-collection.json:63-69`

Neither component-level message includes the API message or process ID.

## What would happen for an actual GET

A failing GET through `customAxios` would use:

- API message: `error.response.data.message`
- Fallback: `"Failed to load data."`
- Process ID: response header `grpc-metadata-process_id`

Extraction is at:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:31-36`

It calls:

- `showErrorToast()`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:17-39`

The resulting QUI toast is:

- Persistent (`timeout: 0`)
- Positioned `top end`
- Description: `Process ID: <value>`
- Includes a Copy action
- Clipboard fallback: `toast.ts:3-15`

The toast remote is loaded through:

- `loadRemote()`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:69-73`

If the remote cannot load, `showErrorToast` silently catches it; the component error state remains visible.

## Response contract and identifier evidence

### Expected table response

Defined at `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:64-71`:

```text
error: boolean
code: number
message: string
data: DashboardTableItem[]
totalData: number
totalPage: number
```

Table rows may contain:

- `id`
- `date` or `transactionDate`
- `entity_name` or `entityName`
- `payment_code` or `paymentCode`
- `amount`

Evidence fixture:

- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\transaction-table.test.tsx:43-50`

The frontend does not inspect `error` or `code`. A 2xx response with `error: true` is treated as successful if its data shape remains usable.

### Error response assumptions

The only recognized error-body field is:

```text
response.data.message
```

No body-based process ID, request ID, correlation ID, or trace ID is read.

### Headers

Recognized identifier header:

```text
grpc-metadata-process_id
```

It is used both outbound and inbound. Axios-header-object and plain-object access are supported at `customAxios.ts:34`.

No handling was found for:

- `requestId` / `request_id`
- `correlationId` / `correlation_id`
- `traceId` / `trace_id`
- `x-request-id`
- `x-correlation-id`

The generated request process ID is not used as a fallback when the response omits the header. For cross-origin calls, the backend must also allow the request header and expose `grpc-metadata-process_id` through CORS for browser JavaScript to read it.

Tests documenting current behavior:

- Header generation/preservation: `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.test.ts:25-64`
- GET error toast and process ID: `customAxios.test.ts:66-83`
- Missing response header: `customAxios.test.ts:85-91`
- Intentional suppression for POST/401: `customAxios.test.ts:93-103`

## Shared helpers found

| Helper | Location | Purpose |
|---|---|---|
| `dashboardRequest<T>()` | `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:77-82` | Shared graph/table POST wrapper |
| `customAxios` | `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:7-42` | Shared Axios instance, auth/process-ID headers, GET error interception |
| `isExcludedRequest()` | `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:11-13` | Suppresses identifiers/toasts for static requests |
| `showErrorToast()` | `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:17-39` | Persistent danger toast with process-ID copy action |
| `copyToClipboard()` | `D:\work\qcash-ui-dashboard-auto-collection\src\utils\toast.ts:3-15` | Clipboard handling |
| `loadRemote()` | `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:69-73` | Loads remote toast/UI implementation |
| `useDashboardAutoCollectionTranslation()` | `D:\work\qcash-ui-dashboard-auto-collection\src\hooks\useDashboardAutoCollectionTranslation.ts:6-15` | Resolves user-facing localized errors |

No additional generic API-error normalizer or error-response type was found.

## Smallest correct edit location

If the goal is to show the existing API message and process ID for the **currently implemented POST**, the root location is:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\customAxios.ts:30-36`

The method-only `GET` predicate is what prevents the already-implemented `showErrorToast()` path from running. If the response may not echo the header, the process-ID extraction at lines 34-35 is also where the outbound request ID should be used as fallback. `toast.ts` and the table component do not need duplicate logic.

If the backend contract truly requires **GET**, change only `getDashboardData()` at:

- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts:100-102`

Do not globally change `dashboardRequest()`, because `getDashboardGraph()` also uses it and currently expects POST.

The working tree already had uncommitted modifications, including process-ID generation in `customAxios.ts`; no files were edited during this exploration.

---

