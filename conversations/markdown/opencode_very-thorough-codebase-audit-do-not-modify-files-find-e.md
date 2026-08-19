# Audit UI process display (@explore subagent)

- **ID**: `ses_048f6e8a7ffetTjTU60JkVOnHz`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 14:17:34`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Very thorough codebase audit. Do not modify files. Find every consumer of API/service errors and every error UI (ResponseModal, Message, ModalPage, etc.). Determine which request flows display `processId` and which drop it at the UI. Return endpoint/service to consumer mapping where possible, exact file:line references, and categorize definite missing vs conditional. Focus UI handling only; do not duplicate transport interceptor/service propagation analysis.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

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

- **glob**: Tool: glob
- **glob**: Tool: glob
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

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

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

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

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

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 17: Assistant

# Audit result

Scope: UI consumers and error presentation only. Transport/interceptor and service-level propagation mechanics are intentionally not re-analyzed.

`ResponseModal` and `Message` are the only locally declared error primitives with a `processId` prop:

- `D:\work\qcash-ui-menu-management\@mf-types\qcash-global-component\_types\components\response-modal\index.d.ts:6-24`
- `D:\work\qcash-ui-menu-management\@mf-types\qcash-global-component\_types\components\message\message.d.ts:5-14`

Their implementations are federated remotes, so this audit can prove that the UI passes `processId`, but not inspect the remote’s final rendering:

- `D:\work\qcash-ui-menu-management\components\global.ts:57-59`
- `D:\work\qcash-ui-menu-management\components\global.ts:71-74`

`ModalPage` has no `processId` property and therefore cannot display one:

- `D:\work\qcash-ui-menu-management\components\modal\index.tsx:18-26`
- `D:\work\qcash-ui-menu-management\components\modal\index.tsx:28-72`

## Executive findings

### Flows that pass `processId` to error UI

1. Main license pending-list and paginated-list failures.
2. License/appearance activity-log failures.
3. Appearance “my pending menus” failures.
4. Appearance table-list failures.
5. License default-fee lookup failures.
6. License configuration submission failures.
7. Appearance draft/approval request rejections.
8. Approve/reject/rework actions routed through `UseOptionModal`.

### Definite UI drops

1. License country, company, and filter-option lookups show a `Message` but omit `processId`.
2. License direct approve/delete flow has its request outside the `try`; rejection bypasses all error UI.
3. License and appearance page-load failures become generic `UnauthorizedDialog`/`ErrorDialog`.
4. Appearance changelog pagination supplies no `openMessage`, so errors produce no UI.
5. Multiple initialization/background lookups reject without any catch or UI.
6. Host-to-host option requests are silent or console-only.
7. Direct changelog downloads are console-only.
8. All `ModalPage` error branches drop `processId`.
9. Several legacy action handlers swallow errors or reduce them to “Network Error”.

### Conditional or unverifiable

1. Table downloads and bulk actions execute inside the federated `Table`; the local code supplies no error modal callback, but the remote implementation is unavailable.
2. Several old reject/error handlers appear unreachable because their modal-opening state is never set.
3. Body-level `code !== 200` responses in appearance save flows use `ModalPage`; transport rejections use process-aware `ResponseModal`.

---

# 1. Process-aware request flows

| Endpoint/service | Consumer and UI | Result |
|---|---|---|
| `POST menu/v3/license/task` — `postMenuLicenseTasksV3`, `D:\work\qcash-ui-menu-management\services\index.tsx:573-593` | Pending cards: request at `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:69-80`; catch passes PID at `:83-89`; `Message` mounted at `:286` | Passes `processId` |
| Same endpoint | Main license table: `usePagination` at `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:206-208`; shared catch passes PID at `D:\work\qcash-ui-menu-management\hooks\use-pagination.ts:119-129`; `Message` at license index `:286` | Passes `processId` |
| `POST /task/activity-logs/Menu:License` — `getActivityLicenseLogs`, service `:227-251` | `ActivityLog` at `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:261-263`; pagination receives `openMessage` at `D:\work\qcash-ui-menu-management\components\activity-log.tsx:17-18` | Passes `processId` |
| `POST /task/activity-logs/Menu:Appearance` — `getActivityAppearanceLogs`, service `:182-204` | `D:\work\qcash-ui-menu-management\containers\menus\appearance\index.tsx:255-258` through `ActivityLog`/`usePagination` | Passes `processId` |
| `GET menu/appearance/task/me` — `getMyMenu`, service `:126-137` | Catch and `Message` at `D:\work\qcash-ui-menu-management\containers\menus\appearance\index.tsx:174-204`, mounted at `:258` | Passes `processId` |
| `GET menu/appearance/task` — `getMenuTaskAppearance`, service `:139-152` | Appearance table pagination at `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:234-248`; `Message` at `:487` | Passes `processId` |
| `GET system/data` — `getDefaultFee`, service `:379-404` | Selected-item fee lookup catch at `D:\work\qcash-ui-menu-management\containers\menus\license\edit\index.tsx:77-100`; `Message` at `:756` | Passes `processId` |
| `POST menu/license/task/edit/{taskID}` — `putMenuTaskLicense`, service `:618-637` | Submit catch at `D:\work\qcash-ui-menu-management\containers\menus\license\edit\index.tsx:206-227`; `ResponseModal` at `:755` | Passes `processId` |
| `POST menu/appearance/task/edit/{id}?mockAppearanceStatus=502` — `postApproval`, service `:536-552` | Rejection catch at `D:\work\qcash-ui-menu-management\containers\menus\appearance\edit\index.tsx:383-412`; `ResponseModal` at `:587` | Passes `processId` on rejected promise |
| `POST menu/appearancesasa/task/edit/{id}?mockAppearanceStatus=502` — `postDraft`, service `:555-570` | Rejection catch at `D:\work\qcash-ui-menu-management\containers\menus\appearance\edit\index.tsx:341-370`; `ResponseModal` at `:587` | Passes `processId` on rejected promise |
| `POST /task/{taskId}` — `putTaskMenuLicense`, service `:723-742` | License actions using `UseOptionModal`, wired at `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:188-204`; hook catches at `D:\work\qcash-ui-menu-management\hooks\use-modal-option.ts:383-388` and `:408-413`; modal at license index `:285` | Passes `processId` |
| `POST task/{taskId}` — `putTaskLicense`, service `:682-700` | License detail actions wired at `D:\work\qcash-ui-menu-management\containers\menus\license\detail\index.tsx:234-247`; `ResponseModal` at `:568` | Passes `processId` |
| `POST menu/apperance/v2/task/action/{taskId}` — `putTaskApperance`, service `:661-679` | Appearance table actions wired at `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:219-232`; modal at `:488` | Passes `processId` |
| Same endpoint | Appearance detail actions wired at `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:126-139`; modal at `:420` | Passes `processId` |
| Same endpoint | Appearance pending-card actions wired at `D:\work\qcash-ui-menu-management\components\table\menu-index-container-pending.tsx:79-92`; modal at `:494` | Passes `processId` |

All generic `UseOptionModal` catches are process-aware:

- Company action: `D:\work\qcash-ui-menu-management\hooks\use-modal-option.ts:238-243`
- User action: `:284-289`
- Reset password: `:318-323`
- Reject/rework/delete-style action: `:383-388`
- Approve action: `:408-413`
- Workflow action: `:456-461`
- Delete action: `:541-546`

Only the `methodUpdate` branches are instantiated by this repository’s current menu screens.

---

# 2. Definite missing `processId`

## 2.1 Error message is shown, but PID is omitted

| Endpoint/service | Consumer | Error UI | Finding |
|---|---|---|---|
| `GET /system/master-data/MdCountry` — `getDataMdCountry`, service `D:\work\qcash-ui-menu-management\services\index.tsx:357-376` | `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:93-131` | `Message` opened at `:126-130` | Shows description only; no `processId` |
| `GET /company/v2/data` — `getCompaniesV2`, service `:319-337` | `D:\work\qcash-ui-menu-management\components\table\filter-license.tsx:46-73` | `Message` at `:68-72` | PID dropped |
| `POST menu/v3/license/distinct` — `postMenuLicenseDistinctV3`, service `:596-615` | Created/reviewed-by options at `D:\work\qcash-ui-menu-management\components\table\filter-license.tsx:87-120` | `Message` at `:115-119` | PID dropped |
| `POST menu/appearance/task/edit/{id}` — `postApproval` | Fulfilled response where `res.data.code !== 200`: `D:\work\qcash-ui-menu-management\containers\menus\appearance\edit\index.tsx:388-397` | `ModalPage` at `:556-569` | Message shown, PID impossible |
| Draft endpoint — `postDraft` | Fulfilled response where `res.data.code !== 200`: `D:\work\qcash-ui-menu-management\containers\menus\appearance\edit\index.tsx:346-355` | Same `ModalPage` | Message shown, PID impossible |
| Appearance action, legacy direct handler | Catch saves only message at `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:94-117` | `ModalPage` at `:370-384` | PID dropped |
| Generic task action, legacy direct handler | Catch saves only message at `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:71-95` | `ModalPage` at `:443-459` | PID dropped |

## 2.2 Request failure bypasses or receives no error UI

### License direct approve/delete

`putTaskMenuLicense` is awaited before the `try`:

- Request: `D:\work\qcash-ui-menu-management\containers\menus\license\index.tsx:135-142`
- `try` begins only at `:143`
- Intended error `ResponseModal` is at `:154-165`

Therefore a service rejection never reaches that modal. This affects calls from:

- Table approve: `D:\work\qcash-ui-menu-management\components\table\table-license.tsx:184-190`
- Table draft delete: `:268-275`
- Pending-card draft delete: `D:\work\qcash-ui-menu-management\components\pending-task.tsx:81-90`
- Pending-card delete-request approve: `:112-124`

Result: no local error UI and no PID.

### License page loaders

`GET menu/license/task/{taskID}` — `getMenuTaskLicenseById`, service `D:\work\qcash-ui-menu-management\services\index.tsx:457-474`.

Consumers discard the error object entirely:

- View page catch: `D:\work\qcash-ui-menu-management\pages\menus\license\view\[taskId].tsx:52-63`
- Generic `UnauthorizedDialog`: `:80-85`
- Edit page catch: `D:\work\qcash-ui-menu-management\pages\menus\license\edit\[taskId].tsx:44-58`
- Generic `UnauthorizedDialog`: `:76-82`

Result: request failures are conflated with authorization/state failures; message and PID are both lost.

### Appearance page loaders

`getActiveMenu` through `fetchMenuAppearance`:

- Service endpoint: `D:\work\qcash-ui-menu-management\services\index.tsx:102-124`
- Utility request: `D:\work\qcash-ui-menu-management\utils\appearance-menus.ts:36-45`
- Utility converts every exception to `isDataExist: false`: `:59-64`

Consumers render generic dialogs:

- View page request: `D:\work\qcash-ui-menu-management\pages\menus\appearances\view\[taskId].tsx:48-54`
- `ErrorDialog`: `:78-80`
- Edit page request: `D:\work\qcash-ui-menu-management\pages\menus\appearances\edit\[taskId].tsx:45-50`
- `ErrorDialog`: `:81-83`

Result: PID and server message are both dropped.

### Appearance changelog pagination

`POST task/activity-logs/{type}` — `getChangeLogs`, service `D:\work\qcash-ui-menu-management\services\index.tsx:292-317`.

The consumer omits `openMessage`:

- `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:36-49`

`usePagination` only presents an error through optional `openMessage`:

- `D:\work\qcash-ui-menu-management\hooks\use-pagination.ts:119-130`

Result: the catch executes but displays nothing.

### Direct changelog downloads

Endpoint constructed as `${NEXT_PUBLIC_API_URL}/task/activity-logs/${action}/download/{format}`:

- Base URL consumer: `D:\work\qcash-ui-menu-management\components\change-log\menu-change-log.tsx:134-140`
- PDF catch: `D:\work\qcash-ui-menu-management\components\change-log\download-button.tsx:85-114`
- XLS catch: `:125-154`
- CSV catch: `:165-194`

All three only call `console.error`; no user-facing error or PID.

---

# 3. Initialization/background requests with no UI handling

These functions are called from effects without surrounding error handling. Once such a request rejects, its PID cannot reach any mounted error primitive.

## License edit initialization

`fetchDataLicense` at `D:\work\qcash-ui-menu-management\containers\menus\license\edit\index.tsx:106-141`, invoked without catch at `:608-612`:

| Service / endpoint | Call |
|---|---|
| `getProducts` → `GET /product` | `:107` |
| `getCUMenuAppearance` → `GET menu/appearance/data` | `:109-112` |
| `getCompanyLimit` → `GET /company/limit` | `:114-116` |
| `getDefaultTransactionLimit` → `GET system/data` | `:139-140` |

Additional unhandled effects:

- Per-currency `getDefaultTLCustom` → `GET system/data`: `:143-154`, invoked at `:614-616`
- `getAccountData` → `GET /account/data`: `:579-592`, invoked at `:623-628`
- `getUsersAuthority` → `POST role/authority/user`: `:594-606`, invoked at `:623-628`

The screen does mount `Message` and `ResponseModal` at `:755-756`, but none of these handlers opens them.

## License detail initialization

`fetchTaskLicense` at `D:\work\qcash-ui-menu-management\containers\menus\license\detail\index.tsx:131-168`, invoked without catch at `:175-179`:

- `getProducts`: `:132`
- `getCUMenuAppearance`: `:134-137`
- `getCompanyLimit`: `:138-140`
- `getDefaultFee`: `:162-164`
- `getDefaultTransactionLimit`: `:166-167`

The `ResponseModal` at `:568` is only connected to action errors, not these initialization requests.

## Appearance configuration preview

`LayoutAppearance.fetchMenu` performs three `getActiveMenu` calls through `Promise.all`:

- `D:\work\qcash-ui-menu-management\containers\menus\appearance\index.tsx:46-57`
- Effect invocation without catch: `:59-61`

Result: all three preview requests are silent on failure.

## Appearance edit “add module” lookup

`getAllDataAppearance` → `GET menu/appearance/data`, service `D:\work\qcash-ui-menu-management\services\index.tsx:255-269`.

Consumers await it without a catch:

- Non-transaction module action: `D:\work\qcash-ui-menu-management\containers\menus\appearance\edit\index.tsx:1271-1277`
- Transaction module action: `:1310-1317`
- Service is passed into the child at `:503-504` and `:1881-1882`

Result: event promise rejection, no error modal and no PID.

---

# 4. Conditional product-specific missing flows

These are reachable only when the relevant Host-to-Host configuration is rendered. `HostToHost` is mounted for eligible transactional products at:

- `D:\work\qcash-ui-menu-management\components\detail\module-slide-over.tsx:1653-1711`

| Endpoint/service | Handling | Result |
|---|---|---|
| `GET /account/data` — `getAccountData` | Initial effect uses `.then` with no catch: `D:\work\qcash-ui-menu-management\components\detail\host-to-host.tsx:263-279` | No UI/PID |
| Same | Menu-close reload uses `.then` with no catch: `:290-302` | No UI/PID |
| Same through `getAccountNumberOptions` | Debounced search catches and ignores: `:280-289`; utility rejects at `D:\work\qcash-ui-menu-management\utils\accounts.ts:3-24` | Silent |
| `GET /mass-transfer/mapping/file-format` — `getFileFormatGroupMapping` | Catch logs only: `D:\work\qcash-ui-menu-management\components\detail\host-to-host.tsx:329-337` | Console-only |
| `POST /approval-signature/list/names` — `getApprovalNameList` | Catch logs and returns `undefined`: `:343-361` | Console-only |

These are definite PID drops once the conditional feature path executes.

---

# 5. Dormant or apparently unreachable legacy error flows

## Pending appearance custom reject flow

The custom request swallows errors:

- Request/catch: `D:\work\qcash-ui-menu-management\components\table\menu-index-container-pending.tsx:94-119`
- Catch body simply returns: `:113-115`

Its `ModalPage`:

- `:363-388`
- Has an empty description at `:372`
- `openError` is declared at `:30`, but no code sets it to `true`

Also, `rejectModal` is never set to `true`; only reset to false. The visible current actions instead use `UseOptionModal`, which is process-aware.

## License detail custom reject flow

The legacy request catches without retaining the error:

- `D:\work\qcash-ui-menu-management\containers\menus\license\detail\index.tsx:208-229`

It then turns failure into a generic “Network Error” raw modal:

- `:436-454`
- Raw modal rendered at `:499-519`

No code opens `confirmationReasonModal`; only close/validation updates exist. Current buttons use `UseOptionModal`.

## Appearance detail custom reject flow

- Request catch opens `ModalPage` with message only: `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:94-117`
- Caller then dereferences `result.code`: `:151-159`, even though the catch returns `undefined`
- Error UI: `:370-384`

No code sets `modalOpened` to true, so the legacy reject modal appears dormant. Current reject actions use `UseOptionModal`.

## Appearance table custom reject flow

- Generic `putTask` handler: `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:71-95`
- Legacy caller/error handling: `:351-390`
- PID-incapable `ModalPage`: `:443-459`

`taskID` is never set and `modalOpened` is never set true in this component. Current row actions use `UseOptionModal`.

## Unused standalone components

- `MenuAppearanceFilter` calls `getMenuTaskAppearance` with no catch at `D:\work\qcash-ui-menu-management\components\table\filter-appearance.tsx:35-48`, but there is no repository consumer of this component.
- `PendingMenu` catches `getMenuTaskAppearance` with `console.error` only at `D:\work\qcash-ui-menu-management\components\summary\index.tsx:85-100`, but there is no repository consumer.

---

# 6. Federated-table flows: handling cannot be verified

## Downloads

- Appearance table passes `downloadTableMenuAppearance` at `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:280-298`
- License table passes `downloadTableMenuLicense` at `D:\work\qcash-ui-menu-management\components\table\table-license.tsx:340-418`
- Endpoints are `/menu/appearance/task/file/{format}` and `/menu/v2/task/file/{format}` at `D:\work\qcash-ui-menu-management\services\index.tsx:94-98`

The federated table declaration accepts only `downloadFn`; it exposes no local error callback:

- `D:\work\qcash-ui-menu-management\@mf-types\qcash-global-component\_types\components\table\table.d.ts:41-62`
- `D:\work\qcash-ui-menu-management\@mf-types\qcash-global-component\_types\components\table\download-button.d.ts:5-18`

Classification: **conditional/unverifiable**. There is no local `Message` or `ResponseModal` wiring for these download errors.

## License bulk actions

`putTaskMenuLicense` is passed into the federated table as `actionFn`:

- `D:\work\qcash-ui-menu-management\components\table\table-license.tsx:390-399`

The remote bulk-action declaration has no `processId` field in its return model:

- `D:\work\qcash-ui-menu-management\@mf-types\qcash-global-component\_types\components\table\bulk-action-modal.d.ts:5-29`

Classification: **conditional/unverifiable**. Any remote generic error modal cannot be audited from this repository.

---

# 7. Error UI inventory

| UI | API error use | PID capability |
|---|---|---|
| `ResponseModal` | License/appearance submissions and `UseOptionModal` action failures | Yes, `processId?: string` |
| `Message` | List/pagination/activity failures, filter lookups, fee lookup | Yes, but several callers omit it |
| `ModalPage variant="error"` | Appearance legacy/direct failures and body-level save failures | No PID prop |
| `ErrorDialog` | Appearance loader failure | No props passed; generic |
| `UnauthorizedDialog` | License loader failure as well as genuine authorization failure | Generic; no PID |
| Raw federated `Modal` | License legacy action failure | Message-only |
| `ComponentsApplicationModal` | Current repository calls use it for success at `D:\work\qcash-ui-menu-management\components\table\table-appearance.tsx:405-441` and `D:\work\qcash-ui-menu-management\containers\menus\appearance\detail\index.tsx:386-419` | Supports an error variant internally, but not used for current API failures |
| `MenuModal` | Confirmation/reason input, not final error output | Not process-aware |
| `ConfirmModal` | Confirmation only | Not an API error consumer |

Formik `ErrorMessage` instances in Host-to-Host and Decryption are field-validation UI, not API/service error handling. Decryption’s failed `Message` calls at `D:\work\qcash-ui-menu-management\components\detail\decryption.tsx:147-177` concern clipboard/download browser operations, not service requests.

## Bottom line

The reliable process-aware path is:

`service rejection → usePagination / UseOptionModal / explicit catch → Message or ResponseModal → processId prop`.

The principal gaps are direct async effects, generic page fallbacks, `ModalPage` error branches, the misplaced license action `try`, Host-to-Host option requests, changelog errors/downloads, and federated-table operations whose error behavior is not visible locally.

---

