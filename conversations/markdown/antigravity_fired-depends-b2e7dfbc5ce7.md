# http://internal-service.example.com/api/bg/v2/issuing/task?pag...

- **ID**: `antigravity-d1128be7c68932b373804d58aa3eb024b530d5477ef0d768cc6db2e7dfbc5ce7`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 21:46:47`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

http://internal-service.example.com/api/bg/v2/issuing/task?page=1&limit=10&query=&filter=&filterOr=data.issuingPortal:%40%3E%5B%7B%22status%22:%22Draft%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Approval%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Posting%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Confirmation+BRISpot%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Draft+BRISpot%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Review%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Approval%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Hold+Account%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Posting%22%7D%5D%7Cdata.issuingPortal:%40%3E%5B%7B%22status%22:%22Insurance+Covering+Submission%22%7D%5D%7C(status:1,workflow_doc.workflow.currentStep:checker)%7C(status:1,workflow_doc.workflow.currentStep:signer)%7C(status:1,workflow_doc.workflow.currentStep:releaser)%7Cstatus:3%7Cstatus:2%7C(status:1,workflow_doc.workflow.currentStep:maker)&sort=status&dir=DESC&tab=on+going can you investigate this url endpoing, when this is being fired and depends on what

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 29: Assistant

### Endpoint Overview

This request is a `GET` call to the **Bank Guarantee (BG) Issuing Task V2** API:
* **Service Method:** [`getDataBGIssuing`](file:///C:/Work/qcash-ui-digi-guarantee/services/digi-guarantee/bg-issuing-operations.ts#L33-L35) in [`services/digi-guarantee/bg-issuing-operations.ts`](file:///C:/Work/qcash-ui-digi-guarantee/services/digi-guarantee/bg-issuing-operations.ts).
* **Base URL:** Configured via [`config.ts`](file:///C:/Work/qcash-ui-digi-guarantee/config.ts#L4-L15) resolving in the DEV environment to `http://internal-service.example.com/api`.
* **Path:** `/bg/v2/issuing/task`

---

### When It Is Fired

This call is triggered from the **BG Issuing History / List** page (`/digi-guarantee/bg-issuing/history` via [`pages/bg-issuing/history/index.tsx`](file:///C:/Work/qcash-ui-digi-guarantee/pages/bg-issuing/history/index.tsx#L52)) through the [`usePagination`](file:///C:/Work/qcash-ui-digi-guarantee/hooks/use-pagination.ts#L95-L121) hook in [`containers/bg-issuing/index.tsx`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L265-L273).

Specifically, it fires:

1. **On Initial Page Mount:**
   When navigating to the BG Issuing page, `usePagination` mounts with default params for the **"Transaksi Berjalan" (Ongoing)** tab and immediately executes `fetch()`.
2. **On Tab Switching / Status Filter:**
   When the user clicks the "Transaksi Berjalan" tab header or selects the "All" status pill inside the Ongoing view ([`handleFilter`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L309-L349)).
3. **On Table Pagination or Sorting:**
   When changing page numbers (`setPageIndex`), page sizes (`setPageLimit`), or clicking table headers to re-sort (`sortChanged`).
4. **On Search & Filter Application:**
   When typing into the search field ([`handleSearch`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L292-L307)) or applying date/type filters ([`handleSetFilter`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L371-L463)).
5. **On Manual Refresh / Action Reload:**
   When clicking the **Refresh** button or after completing actions like task deletion or checking status in [`containers/bg-issuing/index.tsx`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L620-L633).

---

### What It Depends On

#### 1. Preconditions & Feature Toggles
* **V2 Toggle:** [`ENV.HIDE_V2`](file:///C:/Work/qcash-ui-digi-guarantee/pages/bg-issuing/history/index.tsx#L10) must not be `"true"`. If `HIDE_V2` is active, it falls back to the V1 endpoint `/bg/issuing/task`.
* **User Authorization:** 
  * `productAuthorities.BG_ISSUING.anyAuthority` must be `true` ([`pages/bg-issuing/history/index.tsx`](file:///C:/Work/qcash-ui-digi-guarantee/pages/bg-issuing/history/index.tsx#L37)).
  * An `access-token` must exist in `localStorage`, injected by the [`customAxios`](file:///C:/Work/qcash-ui-digi-guarantee/services/custom-axios.ts#L16-L24) request interceptor.

#### 2. Query Parameters Breakdown

| Parameter | Value in URL | Source / Dependency |
| :--- | :--- | :--- |
| `page` | `1` | `pageIndex + 1` in [`usePagination`](file:///C:/Work/qcash-ui-digi-guarantee/hooks/use-pagination.ts#L47) (defaults to page index `0`). |
| `limit` | `10` | `pageLimit` state in [`usePagination`](file:///C:/Work/qcash-ui-digi-guarantee/hooks/use-pagination.ts#L32) (default is `10`). |
| `query` | `""` | `search` state from table search input ([`handleSearch`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L292)). |
| `filter` | `""` | `filter` state updated by the custom date / BG type filter modal ([`handleSetFilter`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L371)). |
| `sort` | `status` | Derived from the default sort string `"status:DESC"` passed into `usePagination`. |
| `dir` | `DESC` | Direction part of the default sort string `"status:DESC"`. |
| `tab` | `on going` | Evaluated from `tabValue === "ongoing" ? "on going" : "history"` ([`containers/bg-issuing/index.tsx`](file:///C:/Work/qcash-ui-digi-guarantee/containers/bg-issuing/index.tsx#L272)). |
| `filterOr` | *(Complex OR string)* | [`IssuingFilterTabs.ONGOING`](file:///C:/Work/qcash-ui-digi-guarantee/const/issuing.ts#L2). |

#### 3. What `filterOr` Represents
The `filterOr` parameter filters for active/in-progress transactions across two systems:
* **Issuing Portal Statuses (`data.issuingPortal`):** `Draft`, `Approval`, `Posting`, `Confirmation BRISpot`, `Draft BRISpot`, `Review`, `Hold Account`, `Insurance Covering Submission`.
* **Workflow / Step Statuses (`workflow_doc.workflow.currentStep`):**
  * `(status:1, currentStep:checker)` — Pending Checker
  * `(status:1, currentStep:signer)` — Pending Signer
  * `(status:1, currentStep:releaser)` — Pending Releaser
  * `(status:1, currentStep:maker)` — Pending Maker
  * `status:2` — Draft
  * `status:3` — Returned

---

