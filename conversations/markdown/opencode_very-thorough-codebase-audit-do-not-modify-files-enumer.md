# Audit endpoint process IDs (@explore subagent)

- **ID**: `ses_048f6e929ffeFapKgTMhoVWhy0`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 14:17:34`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Very thorough codebase audit. Do not modify files. Enumerate every API endpoint/request in this repository (customAxios, axios, fetch, helpers that wrap them). For each unique endpoint or endpoint pattern, determine whether an HTTP error ultimately exposed to its service caller has a correct direct `processId` property. Trace request interceptor generation, response interceptor status branches, service catch/rethrow logic, and note unsafe accesses that can throw or lose IDs. Return a concise categorized report: correct, incorrect/missing, conditional/uncertain, with exact file:line references and endpoint patterns. Focus only transport/service propagation, not UI rendering.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

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

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 11: Assistant

# Audit result

**Criterion:** “correct direct property” means the value is available as `error.processId`, not only as `error.config.processId` or embedded in `error.message`.

Path abbreviations below are absolute:

- **AX:** `D:\work\qcash-ui-menu-management\services\custom-axios.ts`
- **SVC:** `D:\work\qcash-ui-menu-management\services\index.tsx`
- **EX:** `D:\work\qcash-ui-menu-management\services\example.ts`
- **DL:** `D:\work\qcash-ui-menu-management\components\change-log\download-button.tsx`
- **MDL:** `D:\work\qcash-ui-menu-management\components\change-log\menu-change-log.tsx`
- **CFG:** `D:\work\qcash-ui-menu-management\config.ts`
- **ENV:** `D:\work\qcash-ui-menu-management\utils\index.ts`

## Summary

- **Unconditionally correct:** none.
- **Conditional:** 22 customAxios endpoint patterns. They expose the generated request ID directly only for ordinary Axios HTTP rejections when process-ID generation is enabled and the interceptor does not enter a special status branch.
- **Incorrect/missing:** `/product`, `/some-endpoint`, and the three direct-axios activity-log downloads.
- No native Fetch API request exists. `use-pagination.ts:95-102,164,168` calls a locally named `fetch` callback, not `window.fetch`.
- Commented-out axios calls in `components\download-button\download-button.tsx` are inactive and excluded.

---

# Transport/interceptor trace

## Request generation

- `customAxios` uses the environment-selected API base: **AX:7-12**, **CFG:5-15**.
- When `config.toggleProcessId` is true, every request receives:
  - a fresh UUID in `config.processId`;
  - `Grpc-Metadata-Process_Id` request header.  
  **AX:27-34**
- The toggle is true only for dev/prestage and false for preploy/piloting/prod: **CFG:83-89**, selected by hostname at **ENV:67-87**. It is also undefined during SSR.
- A retry through `customAxios(originalRequest)` reruns the request interceptor and therefore generates a new process ID: **AX:120-123**.

## Response handling

- The interceptor computes the preferred ID as response header first, request config second: **AX:86**.
- For ordinary HTTP errors, that computed value is discarded; the original Axios error is rejected unchanged: **AX:150-156**. Thus it has only nested `error.config.processId`, not direct `error.processId`.
- Consequently, service catches are solely responsible for adding the direct property.

### Special status branches

- **401 on `/homepage` or `/`:** clears authentication, redirects, and returns `undefined`, converting rejection to fulfillment: **AX:88-94**.
- **`auth/v2/login` 401:** no matching request exists in this repository. It dispatches a dialog:
  - toggle on: rejects `new Error(JSON.stringify(...processId...))`; the ID is only in the message, not `error.processId`;
  - toggle off: remains pending until external code settles it.  
  **AX:96-109**
- **401/499/408 elsewhere:**
  - queued refresh rejection becomes a new `Error` with no config or direct ID: **AX:112-126**;
  - first refresh attempt sets `isRefreshing` permanently—there is no reset—and only immediately rejects when the toggle is on: **AX:128-141**;
  - that rejection object has a direct ID, but service wrappers generally discard it and reread `error.config.processId`;
  - a successful retry gets a newly generated ID.
- **“Another Login Detected”:** redirects and returns `undefined`, swallowing the rejection: **AX:143-147**.
- **No response/network failure:** rejects a spread object with nested config but does not explicitly add direct `processId`: **AX:150-154**.

### Unsafe interceptor accesses

- `error.response?.headers[...]` does not guard a missing `headers`, and `(error.config as any).processId` does not guard missing config: **AX:86**.
- `originalRequest.url` and `originalRequest._retry` assume `error.config` exists: **AX:96,112**.
- The current test expects an ordinary 502 to have direct `processId` at  
  `D:\work\qcash-ui-menu-management\__tests__\services\custom-axios.test.ts:55-74`, but production code rejects it unchanged at **AX:156**.

---

# Correct

**None.** Every customAxios endpoint depends on the environment toggle and status branch. No endpoint consistently preserves a server-returned process ID as a direct property in all configured environments.

---

# Incorrect or missing

| Endpoint/request | Result | References |
|---|---|---|
| `GET /product` | Both variants rethrow the raw Axios error. Generated ID remains only at `error.config.processId`; there is no direct property. Locally generated non-200 errors have no ID at all. | **SVC:480-515** |
| `POST /some-endpoint` | Catches and swallows every HTTP error; caller receives resolved `undefined`, with no error or process ID. | **EX:3-10** |
| `POST ${NEXT_PUBLIC_API_URL}/task/activity-logs/{action}/download/pdf` | Uses plain axios, so no process-ID interceptor. Error is only logged and swallowed. Actual current action is `menu:appearance`. | **MDL:134-139**; `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:50-61`; **DL:87-114** |
| Same download pattern ending `/xls` | Same defect. | **DL:127-154** |
| Same download pattern ending `/csv` | Same defect. | **DL:167-194** |

---

# Conditional / uncertain

For every endpoint below:

- **Correct only when** process-ID generation is enabled, the failure is an ordinary Axios HTTP rejection, and `error.config.processId` survives.
- The service then constructs a new `Error` and assigns that nested value to direct `err.processId`.
- **Missing or potentially wrong when** the toggle is off, only the response header contains the ID, the response-header ID differs from the request ID, or a special interceptor branch transforms/swallows the error.

| Endpoint pattern | Request and propagation references |
|---|---|
| `GET /menu/appearance/task/file/{format}?{filter}[&sort][&dir]` | Shared request/catch **SVC:59-91**; binding **SVC:94-95** |
| `GET /menu/v2/task/file/{format}?{filter}[&sort][&dir]` | Shared request/catch **SVC:59-91**; binding **SVC:97-98** |
| `GET menu/appearance/task` | Helper request **SVC:38-46**; calls **SVC:102-107,139-140**; error conversion **SVC:29-35,53-55** |
| `GET menu/appearance/task/me` | Request **SVC:126-127**; error conversion **SVC:29-35,53-55** |
| `GET menu/appearance/data` | Helper call **SVC:154-165**, conversion **SVC:29-35,53-55**; second service request/catch **SVC:255-269** |
| `POST /task/activity-logs/{type}` including `Menu:Appearance` and `Menu:License` | Appearance **SVC:182-204**; License **SVC:227-251**; dynamic type **SVC:292-316** |
| `GET /account/data` | **SVC:207-223** |
| `POST /approval-signature/list/names` | **SVC:272-289** |
| `GET /company/v2/data` | **SVC:319-337** |
| `GET /company/limit` | **SVC:340-353** |
| `GET /system/master-data/MdCountry` | **SVC:357-376** |
| `GET system/data?system.key={key}` | Shared helper **SVC:379-399**; fixed `company-limit-default` request **SVC:411-431** |
| `GET /mass-transfer/mapping/file-format?companyID={companyID}` | **SVC:435-453** |
| `GET menu/license/task/{taskID}` | **SVC:457-474** |
| `POST role/authority/user` | **SVC:518-533** |
| `POST menu/appearance/task/edit/{id}?mockAppearanceStatus=502` | **SVC:536-552**; local mock adapter generation **AX:36-64** |
| `POST menu/appearancesasa/task/edit/{id}?mockAppearanceStatus=502` | **SVC:555-570**. Exact source contains `appearancesasa`; it does not match the local mock regex at **AX:39-42**. |
| `POST menu/v3/license/task` | **SVC:573-593** |
| `POST menu/v3/license/distinct` | **SVC:596-615** |
| `POST menu/license/task/edit/{taskID}` | **SVC:618-637** |
| `POST task/{taskId}` / `/task/{taskId}` | Implementations and catches **SVC:640-658,682-700,703-720,723-742** |
| `POST menu/apperance/v2/task/action/{taskId}` | **SVC:661-679**. Exact source uses misspelled `apperance`. |

---

# Service-level loss and unsafe access

## Response-header IDs are universally lost

Although **AX:86** prioritizes `grpc-metadata-process_id`, no ordinary rejection receives that value. All converting services read only `error.config.processId`; none reads the response header. Therefore:

- with the toggle off, a valid server response-header ID is lost;
- if the server header differs from the client UUID, the service exposes the client UUID rather than the interceptor’s preferred server ID.

## Unsafe `error.config.processId`

These accesses can throw a second `TypeError` and replace the original error whenever an interceptor returns `undefined`, refresh handling returns a new `Error`, or service validation throws a plain `Error`:

**SVC:86,198,218,246,264,284,311,332,348,371,426,448,469,528,547,565,588,610,632,653,674,695,715,737**

Only the shared menu handler and `getSystemData` use optional config access:

- **SVC:29-34**
- **SVC:392-398**

They still expose `processId: undefined` when no config exists.

## Locally synthesized errors

Many services throw plain `Error` objects after a fulfilled response reports `data.error` or an unexpected status. These errors have no config. Unsafe catches then throw a `TypeError` instead of the intended message/process ID. Representative validation and catch pairs include:

- **SVC:213-221**
- **SVC:258-267**
- **SVC:277-287**
- **SVC:323-335**
- **SVC:342-351**
- **SVC:363-374**
- **SVC:420-429**
- **SVC:445-451**
- **SVC:461-472**
- **SVC:522-531**
- **SVC:579-591**
- **SVC:602-613**
- **SVC:627-635**
- **SVC:648-656**
- **SVC:669-677**
- **SVC:690-698**
- **SVC:710-718**
- **SVC:731-740**

The post-helper validations at **SVC:109-110,128-129,141-142,167-168** also throw direct plain errors outside the helper’s catch, with no process ID. These are application/validation failures rather than Axios HTTP rejections, but they are relevant ID-loss paths.

---

