# Analyze form architecture (@general subagent)

- **ID**: `opencode-142d62f40c66815604b48cb09f96948ae3bb8ec4126748c35c78626900ee5401`
- **Source Tool**: `opencode`
- **Date**: `2026-08-06 13:04:02`
- **Project**: `qcash-ui-dashboard-auto-collection`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only in repository ~\AppData\Local\Temp\opencode\qcash-ui-issuance-lc. Compare git tags 1.0.164 and 1.0.260, focusing on create/edit form architecture, old/new skin/version branching, form pages, service endpoints, and business validation. Use git show/diff and inspect relevant files. Return 5-10 evidence-based main differences with file paths and commit/tag context. Do not edit anything.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **skill**: Tool: skill

</details>

I’ll inspect the requested repository and compare both tags, then trace the form pages, branching, services, and validation through the relevant snapshots. No files will be changed.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



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

The range is a major new-skin rollout: 243 files changed and about 41k insertions, with the decisive architecture commits clustered around `7213d9a` (old-format create), `a9d3700` (system version and unified save), and subsequent version-condition fixes. I’m narrowing the evidence to entry pages, form containers, page components, services, and validation utilities rather than generated federation types and broad UI churn.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
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

At `1.0.164`, “new skin” was an environment toggle and create/edit were inconsistent: create could fall back to the legacy form, while edit always mounted the new-skin form. By `1.0.260`, both entry pages wait for `trade_finance_sr_version` and consistently choose separate old/new form stacks, with compatibility logic inside the forms for persisted transaction versions.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
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

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

**Scope**
Compared `1.0.164` (`e4485cb`, 2026-06-17) with `1.0.260` (`c77dc44`, 2026-07-31). The range contains 243 changed files, approximately 41k insertions.

1. **Skin selection changed from environment-based to server-controlled versioning.**  
At `1.0.164`, create selected the new skin only when both `NEXT_PUBLIC_DOMAIN === "dev"` and `NEXT_PUBLIC_VARIANT === "new-skin"`; otherwise it rendered the legacy form (`pages/import-lc/issuance-lc/create/index.tsx:129-153`). Edit was inconsistent and always rendered `FormIssuanceLcNew` after authorization (`pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:169-199`). At `1.0.260`, both pages query the existing `/system/data` service for `trade_finance_sr_version`, wait for that request, and select old format unless the value is `"2026"` (`create/index.tsx:86-100,198-215`; `edit/[transactionEventId].tsx:32-50,200-216`; endpoint in `services/system.ts:479`). Create routing arrived in `a9d3700`; edit parity followed in `fd50862`.

2. **The form architecture became two explicit versioned stacks.**  
At `1.0.164`, the stacks were `containers/form.tsx` using `components/create/page-{1..5}.tsx` and `containers/form-new-skin.tsx` using `components/create-new-skin/page-{1..5}.tsx`. At `1.0.260`, active create/edit routing uses `containers/form-new-old.tsx` with `components/create-new-old/page-{1..5}.tsx`, or `containers/form-new-skin.tsx` with the new-skin pages (`containers/form-new-old.tsx:3-8`; `containers/form-new-skin.tsx:3-8`). The old `containers/form.tsx` remains but is no longer imported by either entry page. Commit `7213d9a` added the old-format container, all five page components, tests, and roughly 12k lines.

3. **Old and 2026 forms now enforce different field requirements rather than merely looking different.**  
At `1.0.260`, the 2026 Page 2 requires applicant country, applicant town/city/state, beneficiary country, and beneficiary town/city/state (`components/create-new-skin/page-2.tsx:23-57`). The old-format schema omits applicant country and both town/city/state requirements (`components/create-new-old/page-2.tsx:23-54`). Page 1 is effectively shared, while Pages 2-5 materially diverge. This preserves the smaller old-format business contract instead of forcing 2026 fields onto legacy transactions.

4. **Version compatibility now extends through edit, template loading, preview, and submission.**  
`FormIssuanceLcOld` compares the current system version with persisted `dataFetch.srVersion`, treating blank/`2023` data as old format and gating address, country, HS-code, and other mappings (`containers/form-new-old.tsx:170-175,287-352`). The 2026 form only populates version-sensitive applicant/beneficiary values from records marked `"2026"` (`containers/form-new-skin.tsx:273-283,885-950`). `srVersion` is also passed through Page 5 into summaries and confirmation (`components/create-new-skin/page-5.tsx:54-87,244,445,473`). Relevant corrections landed in `894e80e` and `ecc91be`.

5. **Create/draft/submit moved to a unified backend endpoint.**  
The shared `getSaveIssuedLc` method changed from `POST /issuance-lc/SaveIssuedLc` at `1.0.164` to `POST /issuance-lc/SaveIssuedLcUnified` at `1.0.260`, retaining the signed, URL-encoded request format (`services/index.ts:404-418`). Because drafts in both containers and final submission in `components/modal/modal-confirm/index.tsx:343` use this method, the endpoint switch covers both create and edit flows. This was part of `a9d3700`.

6. **The 2026 payload contract was renamed and made version-conditional.**  
At `1.0.164`, final submission sent `beneficiaryPostCode`, `applicantPostCode`, `incoterms`, and `incotermsDetail` unconditionally (`components/modal/modal-confirm/index.tsx:242-248`). At `1.0.260`, it sends `beneficiaryPostalCode`, `applicantPostalCode`, `incotermsCode`, `incotermsNarrative`, HS code, country codes, and town/city/state only when `srVersion === "2026"` (`components/modal/modal-confirm/index.tsx:138-154,264-270`). Draft and template payloads use the same renamed fields (`containers/form-new-skin.tsx:731-742,1193-1208`). The backend-contract correction is documented in `708a068`; version gating was refined in `894e80e`.

7. **Saved applicant/beneficiary data gained richer, conditional metadata.**  
The endpoint remains `POST /issuance-lc/saveBenefApplicant`, but its service contract expanded from type/name/address/country to also accept `idBenefApplicant`, `countryCode`, `townCityState`, and `postalCode` (`services/index.ts:1290-1310`). Final submission includes these details only for selected favorite records and, for the additional location fields, only for 2026 format (`components/modal/modal-confirm/index.tsx:291-337`). This originated in `cb22e24`.

8. **Text and SWIFT validation became harder to bypass.**  
At `1.0.164`, required description/document fields used plain Yup `.required()` (`components/create-new-skin/page-3.tsx:103-104`), so whitespace handling and pasted multiline limits depended mainly on controls. At `1.0.260`, required free-text fields are trimmed before validation (`components/create-new-skin/page-3.tsx:116-117`), and `utils/swift-field-validation.ts` centrally rejects overlong rows, excessive row counts, and excessive total characters. Page 3 invokes it for additional conditions and other multiline fields (`components/create-new-skin/page-3.tsx:896,2411`). The mandatory-document and paste/row fixes are in `d9b557a`; the reusable SWIFT validator was added by `0226dac`.

9. **Coverage validation changed from indirect button state to explicit submit-time business checks.**  
At `1.0.164`, Page 4 primarily derived an `isCalculate` flag for selected account types before allowing calculation. At `1.0.260`, submission explicitly blocks empty min/max tolerances, coverage amounts above available balance, and backend calculation results containing `"insufficient balance"` (`components/create-new-skin/page-4.tsx:165-181,847-880`). `utils/coverage-account.ts` handles formatted numeric balances, exempts Import Line from cash-balance limits, and recognizes backend insufficiency text. Initial min/max and insufficient-coverage blocking landed in `708a068`; broader account-balance handling and tests landed in `2ded9a8`.

10. **Template application and upload payloads now discard stale state.**  
At `1.0.164`, raw `fileUpload` objects were sent and applying a template did not centrally clear prior coverage calculations (`containers/form-new-skin.tsx:721,1152`). At `1.0.260`, `filterUploadedFiles` removes placeholder entries without real filenames, while template application resets coverage details and triggers Page 4 calculation reset (`utils/issuance-lc.ts:1-6`; `containers/form-new-skin.tsx:744,953,1210,1528`; corresponding old-form locations `772,988,1249,1544`). Upload filtering originated in `c0d81d2`; coverage/template reset behavior was consolidated in `55b40df`.

No repository files or worktree state were modified.

---

