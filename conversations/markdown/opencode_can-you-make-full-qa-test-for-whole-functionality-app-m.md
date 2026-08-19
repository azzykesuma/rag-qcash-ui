# Full app QA test plan in .md

- **ID**: `ses_010473ab6ffe8erBGYkmJanbm4`
- **Source Tool**: `opencode`
- **Date**: `2026-08-11 14:28:11`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

can you make a full qa test for the whole functionality of the app? make a .md

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll map the app’s user flows, permissions, API behavior, validation, and failure states, then turn them into an executable QA checklist with preconditions and expected results.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **task**: Tool: task
- **task**: Tool: task
- **task**: Tool: task
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

## Objective
- Create a comprehensive Markdown manual QA test plan covering the app’s full functionality.

## Important Details
- QA scope: user flows, authorization, APIs, validation, loading/error/empty states, filters, pagination, downloads, localization, accessibility, responsive behavior, and host/MFE integration.
- Functional route: `/dashboard-auto-collection-v2/dashboard`; root route only shows a welcome heading.
- Access requires `userType === "cu"`, `DASHBOARD_AUTO_COLLECTION.anyAuthority === true`, and menu text `"Dashboard Auto Collection"`.
- Main features: account/month filters, refresh timestamp, summary cards, trend chart, searchable paginated transaction table, row selection, and CSV/XLSX/PDF downloads.
- Recommended local command: `yarn proxy`; URL: `http://localhost:3001/dashboard-auto-collection-v2/dashboard`.
- Static inspection only; external APIs and federated components have not been executed.

## Work State
### Completed
- Inventoried user-facing routes, states, permissions, controls, and dashboard flows.
- Mapped backend integrations, request lifecycle, cancellation, table behavior, and downloads.
- Reviewed localization, accessibility, configuration, deployment, and README guidance.
- Confirmed available scripts: `yarn build`, `yarn lint`, `yarn test`, and `yarn test:coverage`.

### Active
- Converting repository findings into executable QA scenarios with preconditions, steps, expected results, and test data.
- Selecting and creating the final `.md` QA document.

### Blocked
- None.

## Next Move
1. Draft the comprehensive QA matrix covering happy paths, permissions, failures, edge cases, accessibility, responsive behavior, localization, APIs, and downloads.
2. Create the Markdown file and verify coverage against all discovered functionality.

## Relevant Files
- `README.md`: setup, `yarn proxy`, scripts, and local dashboard URL.
- `package.json`: build, lint, test, coverage, and runtime scripts.
- `src/pages/dashboard/index.tsx`: authentication gate and dashboard route.
- `src/pages/index.tsx`: placeholder root route.
- `src/containers/dashboard/index.tsx`: dashboard initialization and filter/request flow.
- `src/containers/dashboard/components/dashboard-header.tsx`: account, month, and refresh controls.
- `src/containers/dashboard/components/summary-cards.tsx`: payment, hold, and release summaries.
- `src/containers/dashboard/components/trend-chart.tsx`: daily amount chart.
- `src/containers/dashboard/components/transaction-table.tsx`: search, pagination, selection, and downloads.
- `src/services/dashboard.ts`: dashboard API integrations.
- `src/services/customAxios.ts`: shared Axios behavior.
- `src/utils/auth.ts`: authorization requirements.
- `src/hooks/useDashboardAutoCollectionTranslation.ts`: localization behavior.
- `src/components/ui/splash-screen/new-qlola.tsx`: loading UI and fallback behavior.
- `TRANSLATION_KEYS_MAPPING.md`: translation-key history and localization context.

---

## Turn 5: User

where's the list of the test case?

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The test-case list was not created; the previous response only summarized planned work. I’m creating `QA_TEST_PLAN.md` now with executable cases rather than another status report.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

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

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>

The source confirms three transaction modes, account and 13-month period filters, server-side table search/pagination, page-scoped selection, and three download formats. I’m encoding these plus authorization, API failures, localization, accessibility, responsive, and MFE checks into one QA matrix.

---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

I’m writing the document at the repository root. The expected results describe user-visible acceptance criteria and include API payload checks where behavior depends on account scope, period conversion, or selected-row downloads.

---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 15: Assistant

The full list is now in [`QA_TEST_PLAN.md`](QA_TEST_PLAN.md).

It contains 89 test cases covering authorization, filters, charts, tables, downloads, localization, accessibility, responsive behavior, API resilience, security, and MFE integration. `git diff --check` passed.

---

