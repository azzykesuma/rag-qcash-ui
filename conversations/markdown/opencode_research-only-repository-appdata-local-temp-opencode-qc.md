# Analyze reporting quality stack (@general subagent)

- **ID**: `ses_02a540ee3ffee8SWkRamcOka1n`
- **Source Tool**: `opencode`
- **Date**: `2026-08-06 13:04:03`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only in repository ~\AppData\Local\Temp\opencode\qcash-ui-issuance-lc. Compare tags 1.0.164 and 1.0.260, focusing on PDF/MT/report generation, dependencies/configuration, tests/coverage/Sonar, and generated Module Federation QUI types. Quantify where practical and return evidence with paths. Do not edit.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **skill**: Tool: skill

</details>

I’ll inspect the requested repository and compare the two tags directly, then trace the relevant commits and file-level diffs for each requested area. No files will be changed.

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



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The range is substantial: 243 files, +41,397/-6,297. The requested areas dominate the change, including four new PDF/report generators, extensive report/MT UI rewrites, a large test expansion, and 63 changed generated QUI declaration/package files; I’m narrowing this to behavioral and configuration evidence now.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Tests grew from 30 files/4,848 lines to 65 files/17,962 lines, a net +35 files and +13,114 lines. The tags do not contain checked-in coverage reports; Sonar evidence is configuration plus commit history, so I’ll distinguish configured collection from measured coverage rather than infer a percentage.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The implementation evolved from a 1.0.164 screen-only MT preview into two client-side output paths in 1.0.260: `react-to-print` for DOM printing and native `jsPDF` generators for downloaded MT/report PDFs. Two HTML/canvas helpers are present but have no production callers, which is relevant to both dependency and maintenance scope.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

**Scope**
- `1.0.164` (`e4485cb`, 2026-06-17) to `1.0.260` (`c77dc44`, 2026-07-31).
- 473 commits, including 257 non-merge commits.
- 243 files changed, `+41,397/-6,297`.
- Repository remained clean; no files edited.

**PDF / MT / Reports**

| Measure | `1.0.164` | `1.0.260` |
|---|---:|---:|
| MT/report modal size | 275 lines | 1,906 lines |
| Dedicated generation utilities | 0 | 5 files, 2,725 lines |
| Generation-focused tests | 0 | 7 files, 1,174 lines |

- `1.0.164` only provided an on-screen MT preview. It built SWIFT rows in React, sliced them at a fixed 36 rows per page, and had zoom controls but no print/download handler: `components/form-skin/page-five/preview-mt-modal.tsx`.
- A generic server-side table export already accepted PDF in `1.0.164`; this was separate from MT/report rendering: `hooks/use-download-issuance.ts:31`, `services/index.ts:1038-1160`.
- `1.0.260` adds printing through `react-to-print` for both preview and report DOMs: `components/form-skin/page-five/preview-mt-modal.tsx:1428-1527`.
- Downloaded MT PDFs use native vector `jsPDF`, dynamic row-height pagination, watermarks, disclaimers, and optional logos: `utils/generate-mt-preview-pdf.ts:1-318`. This replaces rasterized HTML for the production MT path.
- Downloaded LC reports use a separate native A4 `jsPDF` renderer with automatic pages, status/header/footer, overview fields, tables, activity logs, coverage data, and `srVersion`-specific applicant/beneficiary formatting: `utils/generate-report-pdf.ts:347-405`, `:992-1381`.
- Report/preview downloads are selected in `handleDownload`: `components/form-skin/page-five/preview-mt-modal.tsx:1530-1595`.
- Report actions are available from transaction detail and the ongoing table: `containers/detail/view-new-skin.tsx:533-558`, `components/table/ongoing.tsx:409-421`, `hooks/use-report-modal.ts`.
- Branding added two PNG assets: `public/icons/qlola-new-bri.png` and `public/icons/watermark-bri.png`. Watermark fallback generation is in `utils/bri-watermark.ts`.
- Two new implementations have no production callers at `1.0.260`: `utils/generate-pdf-from-html.ts` and `utils/build-report-download-html.ts`. Repository references are limited to their own tests, while production calls `generateMtPreviewPdf` and `generateReportPdf`.
- Relevant implementation history includes `bd0933f` (“generate pdf using jsPDF”), `0129734` (“download preview mt, download pdf”), `c605dd7` (“fix generate pdf”), and `be7bd76` (`srVersion` report formatting).

**Dependencies / Configuration**
- Three direct runtime dependencies were added in `package.json`:
  - `html2canvas ^1.4.1`
  - `jspdf ^4.2.1`
  - `react-to-print ^3.3.0`
- `yarn.lock` gained 24 package stanzas and removed none, `+155` lines. These include jsPDF’s image/compression/SVG stack such as `canvg`, `dompurify`, `fast-png`, `fflate`, `pako`, and `svg-pathdata`.
- `html2canvas` currently supports only the production-unreferenced `utils/generate-pdf-from-html.ts`; native jsPDF handles active downloads.
- TypeScript output target changed from `es5` to `es2016`: `tsconfig.json:3`.
- Jest setup gained `TextEncoder`, `TextDecoder`, and jsdom `Blob.text()`/`Blob.arrayBuffer()` polyfills, primarily enabling PDF/blob tests: `jest.setup.ts:1-32`.
- Module Federation configuration only gained the comment `// dts: true`: `next.config.js:37`. No active DTS setting was enabled.
- Both tags have the `FederatedTypesPlugin` import commented out while attempting to construct it inside an empty `try/catch`: `next.config.js:13`, `:89-96`. Therefore no effective generated-type behavior changed through configuration.

**Tests / Coverage / Sonar**

| Measure | `1.0.164` | `1.0.260` | Change |
|---|---:|---:|---:|
| Test files | 30 | 65 | +35 / +116.7% |
| Test source lines | 4,848 | 17,962 | +13,114 / +270.5% |
| Lexical `it`/`test` call sites | 112 | 535 | +423 |
| Changed test files in range |  | 56 | `+13,366/-252` |

- Ten page-level suites alone add 6,360 lines under `components/create-new-old/__tests__/` and `components/create-new-skin/__tests__/`.
- Other large additions include `components/table/__tests__/ongoing.test.tsx` (1,197 lines), `containers/detail/__tests__/view-new-skin.test.tsx` (695), and `components/__tests__/template-list.test.tsx` (562).
- Generation coverage adds 1,174 lines across:
  - `components/form-skin/page-five/__tests__/preview-mt-modal.test.tsx`
  - `utils/__tests__/bri-watermark.test.ts`
  - `utils/__tests__/build-report-download-html.test.ts`
  - `utils/__tests__/generate-mt-preview-pdf.test.ts`
  - `utils/__tests__/generate-pdf-from-html.test.ts`
  - `utils/__tests__/generate-pdf-from-html.ssr.test.ts`
  - `utils/__tests__/generate-report-pdf.test.ts`
- `jest.config.js` retains the same `collectCoverageFrom` scope, but removes explicit `coverageProvider: "v8"` and `coverageReporters: ["text", "lcov"]`. LCOV generation therefore relies on Jest defaults in `1.0.260`.
- `test:coverage` remains `jest --coverage`, and Sonar still consumes `coverage/lcov.info`: `package.json:12`, `sonar-project.properties:10`.
- `sonar-project.properties` is byte-for-byte unchanged. It excludes tests/declarations from source analysis and excludes generated/form-heavy areas from coverage or duplication as configured there.
- The history contains repeated Sonar/coverage work, notably `318f8fd`, `30ee9c2`, `5481d3b`, and `72e3c86`.
- Neither tag contains a checked-in LCOV report or measured coverage percentage, so test growth cannot be translated into a verified coverage increase from repository evidence alone.

**Generated QUI Types**
- `@mf-types/qui/**`: 62 changed files, `+1,886/-1,052`.
- Change composition: 22 added, 39 modified, one 51%-similar rename.
- Snapshot size grew from 1,348 files/3,070,648 bytes to 1,370 files/3,108,633 bytes: +22 files and +37,985 bytes (+1.24%).
- Six top-level compiled facades changed under `@mf-types/qui/compiled-types/src/components/`.
- Data Grid gained selection-control APIs: `DataGridSelectionControlProps` and `DataGridSelectionControlVariant` in `compiled-types/src/components/data-grid.d.ts`.
- Data Table downloads now expose visible/order/pinning column metadata and configurable formats: `node_modules/@qui/components/.rslib/declarations/esm/components/data-table/types.d.ts:54-129`.
- Date and date-range pickers gained English, Indonesian, and Chinese locale types/configuration: `.../date-picker/shared.d.ts`.
- Stepper declarations were reorganized into `stepper/base/` and `stepper/workflow/`, adding locale-aware workflow generation. The top-level compiled facade now exports only the compound default `Stepper`: `compiled-types/src/components/stepper.d.ts`.
- Transaction authorization adds normalized locales, `otp`/`otpType`, richer reason fields, and structured submit feedback/error results while deprecating `plainOtp` and `encryptedOtp`: `.../transaction-authorization/types.d.ts`, `.../transaction-authorization/translate.d.ts`.
- Icons gained lazy country exports plus `Saldo` and `Transfer`: `@mf-types/qui/node_modules/@qui/icons/dist/`.
- Much of the type update landed in commit `5b46678` (“push dts”). Both embedded QUI package versions remain `0.0.0`, so package metadata does not identify the remote release used to generate the snapshots.
- One generated file, `data-table/context.d.ts`, accounts for `724/724` changed lines, indicating substantial generated churn rather than purely additive API change.

---

