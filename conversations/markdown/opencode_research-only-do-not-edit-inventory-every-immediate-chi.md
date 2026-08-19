# Audit branch repo health (@general subagent)

- **ID**: `ses_016c9599bffeEXdA3AYMoBeWf5`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 08:08:20`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Inventory every immediate child repository under ~/Development (detect .git). For each, determine current branch, whether local dev branch exists, whether origin/dev exists if refs are available (do not fetch), dirty status, primary stack from manifests, and broad architecture category. Analyze dev branch without checking out or altering worktrees: use git show/dev refs where needed. Identify repos with no dev branch, detached/non-dev current branches, dirty worktrees, stale/untracked patterns, and branch-level debt indicators (TODO/FIXME, disabled lint/type checks, missing tests, committed build/generated artifacts) with representative exact references. Return compact inventory/counts and highest-risk findings. Do not modify or fetch anything.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read

</details>

I’ll enumerate only immediate child directories containing `.git`, then inspect repository state and available refs without checkout or fetch. I’ll also compare/analyze `dev` blobs and trees directly where available, preserving every worktree.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The workspace contains a large, mostly frontend repository fleet. Initial state inspection already shows many repositories intentionally parked on `master` without a local `dev`, plus several very large dirty worktrees; I’m tightening the status accounting and then correlating each result to the selected analysis ref (`dev`, otherwise `origin/dev`, otherwise `HEAD`).

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Ref analysis confirms 86 repositories: 84 are TypeScript/React frontend assets or related UI packages, one Go service, and one macOS/Swift app; `qlola-nav-logger` is documentation-only. I’m now separating real source debt from fleet-wide generated federation typings and checking exact committed-ref evidence, so generated noise does not dominate the risk ranking.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The strongest fleet-wide pattern is build-gate suppression rather than isolated TODO volume: several active `dev` refs explicitly ignore ESLint and/or TypeScript build errors, while 34 analyzed refs contain no test files. Dirty-state risk is concentrated in generated federation typings, but a few worktrees also contain tracked `node_modules`, `.DS_Store`, coverage, and build outputs.

---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

**Scope**
Read-only inspection on 2026-08-10. No fetch, checkout, ref update, or worktree modification occurred.

Analysis ref selection: local `dev`, otherwise cached `origin/dev`, otherwise `HEAD`. Cached remote refs may not reflect the server.

**Summary**
- Immediate child repositories: **86**
- Local `dev`: **28**
- Cached `origin/dev`: **54**
- No local or cached remote `dev`: **32**
- Currently on `dev`: **26**
- Currently on another branch: **60**
- Detached HEAD: **0**
- Dirty worktrees: **23**
- Tracked changes: **21 repos**
- Untracked content: **12 repos**
- Testless analyzed refs: **34**
- Local `dev` divergence: `qcash-ui-account-receivable` is 9 behind cached `origin/dev`; `qcash-ui-account-statement` is 2 behind. The other 26 local `dev` refs match cached `origin/dev`.
- Committed federation typings: **47,393 files across 69 analyzed refs**

Legend: `L/O` indicates local `dev` / cached `origin/dev`; `dirty` is tracked porcelain entries/untracked top-level entries; `tests` is test-file count on the analyzed ref.

**Inventory**
```text
Repository                              Current branch                    L/O   Dirty    Stack / architecture          Tests
addons-auth-provider                    feat/second-attempt-redirected    -/-   2/1      TypeScript npm library         0
addons-auth-service                     dev                               Y/Y   0/0      Go backend service              0
addons-web-client                       dev                               Y/Y   0/0      Next/React monorepo             244
qcash-global-component                  dev                               Y/Y   0/0      Next shared-component MFE       315
qcash-ui-account-management             v1.575.1-release                  Y/Y   0/0      Next/React business MFE         94
qcash-ui-account-receivable             dev                               Y/Y   5/1      Next/React business MFE         157
qcash-ui-account-registration           dev                               Y/Y   146/65   Next/React business MFE         30
qcash-ui-account-statement              master                            Y/Y   142/64   Next/React business MFE         19
qcash-ui-account-summary                master                            -/-   0/0      Next/React business MFE         15
qcash-ui-account-validation             master                            -/-   0/0      Next/React business MFE         0
qcash-ui-amend-cancel-lc                master                            -/-   0/0      Next/React business MFE         49
qcash-ui-announcement-management        master                            -/Y   0/0      Next/React business MFE         11
qcash-ui-approval-signature             master                            -/-   0/0      Next/React business MFE         0
qcash-ui-bank-monitoring-transaction    master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-base                           master                            -/-   0/0      Next/React MFE template         0
qcash-ui-beneficiary-account-management dev                               Y/Y   0/0      Next/React business MFE         49
qcash-ui-bifast                         master                            -/-   0/0      Next/React business MFE         0
qcash-ui-cashcard                       master                            -/-   0/0      Next/React business MFE         49
qcash-ui-company-management             dev                               Y/Y   0/0      Next/React business MFE         68
qcash-ui-complaint                      master                            -/Y   0/0      Next/React business MFE         34
qcash-ui-content-management-system      master                            -/Y   0/0      Next/React business MFE         36
qcash-ui-credit-card                    dev                               Y/Y   1/0      Next/React business MFE         0
qcash-ui-cut-off-management             master                            -/-   0/0      Next/React business MFE         0
qcash-ui-dashboard-auto-collection      dev                               Y/Y   0/0      Next/React business MFE         23
qcash-ui-dashboard-dhe                  master                            -/Y   0/0      Next/React business MFE         16
qcash-ui-dashboard                      dev                               Y/Y   1/0      Next/React business MFE         68
qcash-ui-design-system                  dev                               Y/Y   11/0     React/Turborepo design system   84
qcash-ui-df-pupuk-indonesia             master                            -/Y   0/0      Next/React business MFE         22
qcash-ui-digi-guarantee                 master                            -/Y   2/0      Next/React business MFE         85
qcash-ui-do-pertamina                   master                            -/Y   0/0      Next/React business MFE         14
qcash-ui-do-pupuk-indonesia             master                            -/Y   0/0      Next/React business MFE         18
qcash-ui-do-semen-indonesia             master                            -/-   0/0      Next/React business MFE         0
qcash-ui-employee-data                  master                            -/Y   0/0      Next/React business MFE         27
qcash-ui-ewallet-topup                  master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-forex                          dev                               Y/Y   0/1      Next/React business MFE         0
qcash-ui-fund-transfer                  dev                               Y/Y   32/19    Next/React business MFE         55
qcash-ui-global-account-statement       master                            -/-   0/0      Next/React business MFE         0
qcash-ui-global-account-summary         master                            -/-   0/0      Next/React business MFE         0
qcash-ui-global-external-transfer       master                            -/Y   0/0      Next/React business MFE         42
qcash-ui-global-internal-transfer       master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-group-management               dev                               Y/Y   0/0      Next/React business MFE         2
qcash-ui-header-footer-bri              master                            -/-   7/1      Next/React shared-layout MFE    0
qcash-ui-header-footer                  dev                               Y/Y   1/1      Next/React shared-layout MFE    84
qcash-ui-helpdesk-portal                dev                               Y/Y   3/1      Next/React business MFE         18
qcash-ui-holiday-management             master                            -/-   0/0      Next/React business MFE         0
qcash-ui-issuance-lc                    master                            -/-   1/0      Next/React business MFE         68
qcash-ui-landing-page                   dev                               Y/Y   1/1      Next/React business MFE         59
qcash-ui-language-management            master                            -/Y   0/0      Next/React business MFE         67
qcash-ui-liquidities                    master                            -/-   0/0      Next/React business MFE         40
qcash-ui-local-tax-dki-jakarta          master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-main-page                      master                            -/-   3/0      Next/React business MFE         0
qcash-ui-mass-brizzi                    master                            -/Y   0/0      Next/React business MFE         17
qcash-ui-mass-transfer                  master                            -/-   0/0      Next/React business MFE         57
qcash-ui-menu-management                dev                               Y/Y   0/0      Next/React business MFE         29
qcash-ui-mitra-asuransi                 master                            -/-   0/0      Next/React business MFE         0
qcash-ui-mpn                            master                            -/Y   0/0      Next/React business MFE         66
qcash-ui-onboarding-tour                dev                               Y/Y   0/0      Next/React business MFE         36
qcash-ui-partners-integrations          master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-payroll-transfer               master                            -/Y   0/0      Next/React business MFE         46
qcash-ui-pencairan-bulog                master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-plafon-pegadaian               dev                               Y/Y   1/0      Next/React business MFE         5
qcash-ui-rdn-summary                    master                            -/-   0/0      Next/React business MFE         16
qcash-ui-rdn-transaction                master                            -/Y   0/0      Next/React business MFE         0
qcash-ui-registration-giro              master                            -/Y   4/5      Next/React business MFE         3
qcash-ui-registration-online            dev                               Y/Y   0/0      Next/React business MFE         0
qcash-ui-regol                          master                            -/-   0/0      Next/React business MFE         0
qcash-ui-report-statement               master                            -/-   2/0      Next/React business MFE         1
qcash-ui-rma-management                 master                            -/Y   0/0      Next/React business MFE         15
qcash-ui-role-management                master                            -/Y   1/0      Next/React business MFE         156
qcash-ui-rtgs                           master                            -/-   0/0      Next/React business MFE         0
qcash-ui-scm-cu                         master                            -/Y   0/0      Next/React business MFE         92
qcash-ui-scm                            master                            -/-   0/0      Next/React business MFE         87
qcash-ui-sistem-management              dev                               Y/Y   0/1      Next/React business MFE         13
qcash-ui-span                           master                            -/-   0/0      Next/React business MFE         0
qcash-ui-subscription-management        master                            -/-   0/0      Next/React business MFE         0
qcash-ui-swift                          master                            -/-   0/0      Next/React business MFE         0
qcash-ui-update-user                    master                            -/-   0/0      Next/React business MFE         0
qcash-ui-user-basic-management          master                            -/Y   0/0      Next/React business MFE         45
qcash-ui-user-manual                    dev                               Y/Y   6/0      Next/React business MFE         23
qcash-ui-user-registration-online       dev                               Y/Y   0/0      Next/React business MFE         22
qcash-ui-user                           master                            -/-   0/0      Next/React business MFE         0
qcash-ui-workflow                       dev                               Y/Y   0/0      Next/React business MFE         105
qcash-ui                                dev                               Y/Y   0/0      Next/React host application     131
qlola-nav-logger                        main                              -/-   1/0      JavaScript browser extension    0
qui                                     master                            -/-   0/0      React/Turborepo design system   0
voice_typer                             main                              -/-   0/0      Python desktop automation       0
```

**Highest Risk**
1. **Build validation is explicitly bypassed on active development refs.**
   - `qcash-global-component` `refs/heads/dev:next.config.js:11`: `ignoreDuringBuilds: true`
   - `qcash-global-component` `refs/heads/dev:next.config.js:14`: `ignoreBuildErrors: true`
   - `qcash-ui` `refs/heads/dev:next.config.js:78`: `eslint: { ignoreDuringBuilds: true }`
   - `qcash-ui` `refs/heads/dev:next.config.js:79`: `typescript: { ignoreBuildErrors: true }`
   - `qcash-ui-company-management` `refs/heads/dev:next.config.js:50`: `ignoreBuildErrors: true`
   - `addons-web-client` has `"strict": false` at `refs/heads/dev:apps/web-client/tsconfig.json:9`, plus the same setting in `apps/design-system`, `libs/abonnements`, `libs/menus`, and `libs/systems`.
   - Active suppressions also occur in `qcash-ui-account-management`, `qcash-ui-account-receivable`, `qcash-ui-content-management-system`, `qcash-ui-group-management`, `qcash-ui-header-footer`, `qcash-ui-language-management`, `qcash-ui-user-basic-management`, and `qcash-ui-workflow`.

2. **Three worktrees are dominated by generated federation-type churn.**
   - `qcash-ui-account-registration`: 146 tracked changes and 65 untracked paths, overwhelmingly under `@mf-types/qcash-global-component/`, plus modified `package.json` and `yarn.lock`.
   - `qcash-ui-account-statement`: 142 tracked changes and 64 untracked paths under the same generated tree.
   - `qcash-ui-fund-transfer`: 32 tracked changes and 19 untracked paths under the generated tree.
   - This is amplified at branch level: 69 analyzed refs commit 47,393 `@mf-types` files.

3. **Several refs commit dependency/build products.**
   - `qcash-ui-amend-cancel-lc`: 1,230 committed files under `@mf-types/qui/node_modules/`; representative path `@mf-types/qui/node_modules/@qui/components/.rslib/declarations/esm/charts/bars/chart.d.ts`.
   - The same embedded dependency pattern exists in `qcash-ui-dashboard-dhe`, `qcash-ui-dashboard`, `qcash-ui-forex`, `qcash-ui-issuance-lc`, `qcash-ui-landing-page`, and `qcash-ui-scm-cu`.
   - `addons-auth-provider` commits 105 `dist` files, including `HEAD:dist/src/auth.js`; its source TODO is duplicated into generated output.
   - `qcash-ui-design-system` commits coverage output at `[BITBUCKET_PAT_REDACTED].info`.
   - `qcash-global-component` commits `refs/heads/dev:dist/middleware-manifest.json`.
   - `qcash-ui-registration-giro` current `master` tracks `node_modules/@next/swc-darwin-arm64/next-swc.darwin-arm64.node`, presently deleted locally, and contains seven committed `.DS_Store` files.

4. **Large active source TODO/FIXME inventories exist.**
   - `addons-web-client`: 116 non-generated matches; representative `refs/heads/dev:apps/web-client/components/layouts/default-backup.tsx:1`: `// TODO: remove unused file`
   - `qcash-global-component`: 23 source matches, including:
     - `refs/heads/dev:components/activity-log/download-button.tsx:1`
     - `refs/heads/dev:services/kliring-transfer/index.ts:178`: `TODO: Move to real api`
     - `refs/heads/dev:utils/task/get-description.tsx:5`: `FIXME: get currentStep from WORKFLOW`
   - `qcash-ui-fund-transfer` has unresolved transfer-path markers such as `[BITBUCKET_PAT_REDACTED]-fund-transfer/index.tsx:1414`.
   - `qcash-ui-digi-guarantee` `refs/remotes/origin/dev:containers/bg-issuing/view/view-bg-issuing.tsx:57`: `TODO: Implement preview functionality`
   - `qcash-ui-partners-integrations` `[BITBUCKET_PAT_REDACTED].tsx:201`: unresolved entity-ID handling.
   - `qcash-ui-swift` `HEAD:components/swift-filter-page.tsx:181`: payload fields still awaiting implementation.

5. **Test coverage is absent from 34 analyzed refs.**
   - `addons-auth-provider`, `addons-auth-service`, `qcash-ui-account-validation`, `qcash-ui-approval-signature`, `qcash-ui-bank-monitoring-transaction`, `qcash-ui-base`, `qcash-ui-bifast`, `qcash-ui-credit-card`, `qcash-ui-cut-off-management`, `qcash-ui-do-semen-indonesia`, `qcash-ui-ewallet-topup`, `qcash-ui-forex`, `qcash-ui-global-account-statement`, `qcash-ui-global-account-summary`, `qcash-ui-global-internal-transfer`, `qcash-ui-header-footer-bri`, `qcash-ui-holiday-management`, `qcash-ui-local-tax-dki-jakarta`, `qcash-ui-main-page`, `qcash-ui-mitra-asuransi`, `qcash-ui-partners-integrations`, `qcash-ui-pencairan-bulog`, `qcash-ui-rdn-transaction`, `qcash-ui-registration-online`, `qcash-ui-regol`, `qcash-ui-rtgs`, `qcash-ui-span`, `qcash-ui-subscription-management`, `qcash-ui-swift`, `qcash-ui-update-user`, `qcash-ui-user`, `qlola-nav-logger`, `qui`, and `voice_typer`.
   - `qcash-ui-group-management` technically has two test files, but `refs/heads/dev:hooks/__tests__/use-debounce.test.tsx:1` says `TODO: add runtime tests`.

**Branch And Staleness**
- Non-`dev` current branches: 56 `master`, 2 `main`, one feature branch, and one release branch.
- Particularly notable:
  - `qcash-ui-account-management` is on `v1.575.1-release` despite having an up-to-date local `dev`.
  - `qcash-ui-account-statement` is on dirty `master`; local `dev` is two commits behind cached `origin/dev`.
  - `addons-auth-provider` is dirty on `feat/second-attempt-redirected` and has no `dev` ref.
- Cached analyzed refs older than one year include:
  - `qcash-ui-partners-integrations` `origin/dev`: 2024-04-24
  - `qcash-ui-rdn-transaction` `origin/dev`: 2024-12-24
  - `qcash-ui-pencairan-bulog` `origin/dev`: 2024-12-30
  - `qcash-ui-global-internal-transfer` `origin/dev`: 2025-06-10
  - `qcash-ui-ewallet-topup` `origin/dev`: 2025-07-22
  - Ten additional no-`dev` repositories have analyzed `HEAD` dates before 2025-08-10.
- Because no fetch was performed, these dates establish local ref staleness only, not remote-server inactivity.

**Untracked Patterns**
- Generated typings: account registration, account statement, account receivable, and fund transfer.
- Local dependency/config files: `qcash-ui-forex:yarn.lock`, `qcash-ui-header-footer-bri:.npmrc`.
- Temporary reports: `qcash-ui-header-footer:tmp/coverage.txt`.
- New source/tests: `qcash-ui-helpdesk-portal:containers/tickets/mount-integration.test.tsx`, `qcash-ui-landing-page:utils/useUntranslatableTitle.ts`.
- Tool state: `qcash-ui-sistem-management:.agent-memory/`.
- Machine artifacts: `.DS_Store` in `addons-auth-provider` and several `node_modules` locations in `qcash-ui-registration-giro`.

---

