# Analyze list approval detail (@general subagent)

- **ID**: `opencode-cd79ff3dfb7ebd1220b420cac42ba55a2eb9765cf67b5171e1a0068ba590e958`
- **Source Tool**: `opencode`
- **Date**: `2026-08-06 13:04:03`
- **Project**: `qcash-ui-dashboard-auto-collection`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only in repository ~\AppData\Local\Temp\opencode\qcash-ui-issuance-lc. Compare tags 1.0.164 and 1.0.260, focusing on issuance list/table/filter behavior, authorization/bulk actions, workflow stepper, detail views, and user-visible UI changes. Use git show/diff. Return concise evidence with file paths and practical behavior impact. Do not edit.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

I’ll inspect the requested repository’s tag range, identify relevant changed files, then read targeted diffs and surrounding code to map changes to user-visible behavior. No files will be modified.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The range is a large UI migration rather than a narrow patch. The relevant path converges on a rewritten ongoing-list screen, new bulk authorization hooks, a shared workflow stepper, and a new detail container; I’m separating those flows from unrelated create-form and generated type changes.

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The evidence shows two distinct stepper changes: the create/edit wizard now prevents jumping into unvisited future steps, while the detail page replaces the old workflow modal/summary treatment with an inline role-aware activity timeline. The standalone `components/workflow-stepper.tsx` is test/demo-only in this range and is not wired into production pages.

---

## Turn 6: Assistant

Compared `1.0.164..1.0.260` directly with `git show` and `git diff`.

**List And Filters**
- The `Ongoing Task` / empty `Done` tab layout became one consolidated issuance list. `components/table/history.tsx` was deleted and tab handling was removed from `containers/new-skin.tsx`. Users no longer see a separate Done tab.
- Filters changed from staged values with `Apply filters` and `Reset` buttons to immediate filter chips. Applicant and beneficiary names apply after 500 ms; country, currency, and issue date apply on selection. Evidence: `components/table/issuance-lc-filter.tsx`, `components/table/autompllete-filter.tsx`, `components/table/textfiled-filter.tsx`.
- Issue Date is now visible; it was commented out in `1.0.164`. Country/currency selectors now include searchable lists, clear controls, deduped/sorted options, and explicit `All Countries` / `All Currencies`. Evidence: `hooks/use-issuance-lc-filter.ts`.
- The table’s general search box was disabled (`isSearchable={false}`), so filtering now depends on the dedicated controls. Evidence: `components/table/ongoing.tsx`.
- Status filtering is split into `Workflow Status` and `Transaction Status`. Selecting one clears the other. Draft is hidden from non-makers. Rejected is now available as a workflow status. Evidence: `components/table/ongoing.tsx`, `hooks/use-ongoing-filters.ts`, `const/index.ts`, `utils/map-status.ts`.
- The table adds a separate `Issue Date` column; `Created At` now uses `createDate` instead of `issuedDate`. Applicant and beneficiary names are limited to two lines with full text available through the HTML title. Evidence: `components/table/ongoing-columns.tsx`.
- Row menus now distinguish `View Detail`, completed-only `Download Document`, and non-draft `Download Report`. Evidence: `components/dropdown.tsx`.
- Table export moved from `/issuance-lc/download` to format-specific `/issuance-lc/downloadIssuanceLC/{format}` and carries the active filters/status. Evidence: `services/index.ts`, `hooks/use-download-issuance.ts`.

**Authorization And Bulk Actions**
- Functional page-scoped bulk actions were added. Checker/signer/releaser users receive Approve, Reject, and Return; makers receive Delete. Previously the ongoing table only tracked selection without these operations. Evidence: `components/table/ongoing.tsx`, `hooks/use-bulk-transaction-actions.ts`.
- Rows are selectable only when eligible for the current role and workflow position. Selection is cleared whenever filters change, and single-row Edit/Approve/menu actions are disabled while bulk selection is active.
- Bulk operations use the new `/issuance-lc/MultiApprovalIssuedLc` endpoint and send all selected transaction IDs with the acting approval position. Evidence: `services/index.ts`.
- Bulk confirmation previews up to ten transactions, summarizes additional rows, and remains open after submission to show per-row success/failure results. Partial failures report processed and failed counts. Evidence: `hooks/use-bulk-transaction-actions.ts`, `utils/ongoing-lc-qr-payload.ts`, `components/table/ongoing-columns.tsx`.
- Bulk deletion accepts maker-owned Draft/Returned rows; invalid selections show an `Invalid Selection` warning. Delete does not require an OTP.
- Single-row authorization now fetches full detail before opening, producing a substantially richer preview with LC parties, amount tolerance, coverage, shipment, documents, correspondence, and attachments. Evidence: `hooks/use-transaction-auth.tsx`, `utils/map-data-issue-to-row.ts`, `components/transaction-preview-lc.tsx`.
- A backend `400` during authorization now shows `SWIFT Format Update Required`, explaining that new Applicant/Beneficiary fields must be completed. Successful actions use explicit result dialogs and refresh the table.
- The transaction authorization component now receives the Bricams ID token rather than the ordinary API token. Evidence: `pages/import-lc/issuance-lc/index.tsx`.

**Workflow Stepper**
- Create/edit users can no longer jump directly to unvisited future steps. `maxUnlockedStep` advances only after Next; completed steps remain navigable and Next is capped at the final step. Evidence: `hooks/use-stepper.ts`, `components/ui/qui-stepper.tsx`.
- Detail workflow moved to an inline Activity Log using the new QUI workflow stepper. It dynamically includes checker and releaser stages only when workflow requirements exist, always includes Created, Signer, and BRI, and displays BRI complete/rework/rejected outcomes. Evidence: `components/detail/activity-log.tsx`.
- Draft details show an empty activity state instead of a workflow: `Rincian aktivitas akan muncul di sini`.
- The activity timeline exposes details and connectors, but is clickable only for a user currently allowed to approve/reject.
- `components/workflow-stepper.tsx` is a static Account/Verification/Approval example covered by tests, but `git grep` shows no production-page usage.

**Detail Views**
- The old monolithic detail page was replaced by `containers/detail/view-new-skin.tsx` and modular sections under `components/detail/`.
- The new layout presents status badges, an issuance-state banner, inline MT preview, amount/tolerance, coverage and exchange-rate expansion, general information, payment terms, shipment, document details, correspondent details, and attachment View/Download actions.
- Detail authorization is role/step-aware in a sticky footer. Pending checker/signer/releaser tasks expose Reject, Return, and Approve; `deleterequest` suppresses Return. The originating user can Edit/Delete only for Draft/Returned/Pending at maker/checker steps, with Delete disabled for Returned. Evidence: `containers/detail/view-new-skin.tsx`.
- Download UI is status-dependent: Draft hides it; completed transactions offer report only; other applicable statuses offer MT preview and report. Reports and previews now use the same detailed MT modal.
- Detail fetching moved to React Query and explicitly renders `UnauthorizedDialog` on HTTP 403 or company mismatch; other failures produce a visible error toast. Evidence: `pages/import-lc/issuance-lc/detail/[transactionEventId].tsx`, `hooks/use-issuance-detail-query.ts`.
- User-visible state banners now distinguish `LC Issuance Draft`, `LC Issuance in Progress`, `LC Issuance was Rejected`, and `LC Issuance was Completed`. Evidence: `components/detail/header-detail.tsx`.

No files were edited.

---

