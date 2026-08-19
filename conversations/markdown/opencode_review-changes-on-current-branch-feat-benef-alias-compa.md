# Review alias cherry-pick (@general subagent)

- **ID**: `ses_001bc834cffeKrtx39pRc2ipxT`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 10:14:23`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Review the changes on current branch feat-benef-alias compared with prerelease, focusing on the four cherry-picked beneficiary account alias commits (current commits d70f273b, c90db8c2, d915e101, 68430fd0). Read-only review: do not edit files. Identify concrete bugs, behavioral regressions, incorrect alias ordering/mapping, test gaps, or conflict-resolution mistakes. Return findings ordered by severity with exact file and line references; explicitly state if no findings. Also note which focused test files should be run. Thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’ll inspect the specified commits, the aggregate diff against `prerelease`, and the surrounding alias-related implementation/tests. I’ll keep the review read-only and validate suspicious paths against existing data contracts and call sites.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>

The aggregate diff already exposes an important asymmetry: `online-transfer` was changed in the feature commit and then its production change was fully canceled by `68430fd0`, while its tests remain substantially modified. I’m tracing whether that was intentional and checking each form’s option label versus its selected-value label, since those paths currently use different ordering rules.

---

## Turn 5: Assistant

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

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The shared hook returns `alias - name - number`, but the selected-value labels in BIFAST, Kliring, RTGS, beneficiary creation, and SWIFT were moved to `name - number - alias`. Online Transfer and IFT remain alias-first, and some newly added BIFAST assertions still expect alias-first, so this is both a behavioral inconsistency and likely an immediately detectable test failure. I’m running the focused suites and checking whether the cherry-pick also introduced non-alias regressions.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

**Findings**

1. **High: SWIFT drops a valid selected beneficiary when bank name is absent**
   - [`components/form/swift/index.tsx:1200-1213`](components/form/swift/index.tsx#L1200-L1213) only constructs the selected value when `beneficiaryBankName` is truthy, even though the actual identity field is `beneficiaryNumber`.
   - Commit `d70f273b` correctly changed this guard to `value?.beneficiaryNumber`; commit `68430fd0` reverted it to `value?.beneficiaryBankName`, indicating a conflict-resolution/reversion mistake.
   - Drafts, favorites, or migrated records containing account number/name/alias but no bank name render the registered-beneficiary select as empty.
   - The regression is reproduced by [`utils/test/forms/swift-form.test.tsx:2246-2266`](utils/test/forms/swift-form.test.tsx#L2246-L2266): expected `Main | John Doe | 123`, received an empty label.

2. **Medium: alias ordering is inconsistent between dropdown options and selected values**
   - The shared beneficiary hook deliberately produces `alias - name - number` at [`hooks/use-inqury-beneficiary-accounts.ts:35-38`](hooks/use-inqury-beneficiary-accounts.ts#L35-L38).
   - IFT and Online Transfer also remain alias-first at [`components/form/internal-fund-transfer/index.tsx:1781-1783`](components/form/internal-fund-transfer/index.tsx#L1781-L1783) and [`components/form/online-transfer/index.tsx:998-1004`](components/form/online-transfer/index.tsx#L998-L1004).
   - Commit `c90db8c2`, however, moved alias to the end in the selected-value rendering for:
     - [`components/form/bifast/index.tsx:798-804`](components/form/bifast/index.tsx#L798-L804)
     - [`components/form/kliring/index.tsx:657-663`](components/form/kliring/index.tsx#L657-L663)
     - [`components/form/rtgs/index.tsx:767-773`](components/form/rtgs/index.tsx#L767-L773)
     - [`components/beneficiary-favourite/create/form/index.tsx:252-258`](components/beneficiary-favourite/create/form/index.tsx#L252-L258)
   - Consequently, an option appears as `Payroll - Employee - 123` while open, then changes to `Employee - 123 - Payroll` after selection or hydration.
   - This is directly reproduced by:
     - [`utils/test/forms/bifast-form.test.tsx:443-459`](utils/test/forms/bifast-form.test.tsx#L443-L459)
     - [`utils/test/forms/bifast-form.test.tsx:541-555`](utils/test/forms/bifast-form.test.tsx#L541-L555)
     - [`utils/test/forms/kliring-form.test.tsx:318-324`](utils/test/forms/kliring-form.test.tsx#L318-L324)
   - Those tests currently fail with alias-last output.

3. **Medium: SWIFT’s final alias mapping contradicts its own added test and intended ordering**
   - [`hooks/swift/use-swift-system-data.ts:99-102`](hooks/swift/use-swift-system-data.ts#L99-L102) returns `name | number | alias`.
   - [`components/form/swift/index.tsx:1204-1210`](components/form/swift/index.tsx#L1204-L1210) uses the same alias-last order.
   - The feature commit originally made this `alias | name | number`; `d915e101` changed the hook back, and `c90db8c2` changed the component back, but the added expectation remains alias-first at [`utils/test/hooks/use-swift-system-data.test.ts:335-348`](utils/test/hooks/use-swift-system-data.test.ts#L335-L348).
   - The hook suite currently fails: expected `BA | Benef Account | 456`, received `Benef Account | 456 | BA`.
   - This also conflicts with the BIFAST, Kliring, IFT, and Online Transfer expectations that put alias first.

4. **Low: Online Transfer submission coverage was weakened by an unrelated cherry-pick change**
   - [`utils/test/forms/online-transfer-form.test.tsx:765-773`](utils/test/forms/online-transfer-form.test.tsx#L765-L773) no longer verifies that draft submission calls `handleSaveData(..., true)` and approval calls `handleSaveData(..., false)`.
   - After the first click, the final `toHaveBeenCalled()` assertion can pass using the draft call even if the approval click never submits.
   - Commit `68430fd0` weakened these assertions while reverting a temporary production change. The final Online Transfer production code has no net difference from `prerelease`, making these test changes unrelated to the alias feature and a likely cherry-pick cleanup mistake.

5. **Low: RTGS and beneficiary-creation alias ordering lack effective regression coverage**
   - RTGS has the same alias-last defect at [`components/form/rtgs/index.tsx:767-773`](components/form/rtgs/index.tsx#L767-L773), but there is no RTGS form test file under `utils/test/forms/`.
   - The beneficiary creation form has the same issue at [`components/beneficiary-favourite/create/form/index.tsx:252-258`](components/beneficiary-favourite/create/form/index.tsx#L252-L258), but `utils/test/forms/form-benef.test.tsx` does not assert the selected label ordering.
   - These gaps allowed the same `c90db8c2` regression to remain undetected outside BIFAST and Kliring.

**Focused Tests**

Run with `--coverage=false` when executing only these files, since project-wide coverage collection can make otherwise-passing focused runs exit nonzero:

```bash
yarn test \
  utils/test/hooks/use-inqury-beneficiary-accounts.test.ts \
  utils/test/hooks/use-swift-system-data.test.ts \
  utils/test/forms/bifast-form.test.tsx \
  utils/test/forms/internal-fund-transfer.test.tsx \
  utils/test/forms/kliring-form.test.tsx \
  utils/test/forms/online-transfer-form.test.tsx \
  utils/test/forms/swift-form.test.tsx \
  utils/test/forms/form-benef.test.tsx \
  --runInBand --coverage=false
```

Observed results:

- `use-inqury-beneficiary-accounts.test.ts`: passed, 2/2.
- `use-swift-system-data.test.ts`: failed, 1 alias-order assertion.
- `bifast-form.test.tsx`: failed, 2 alias-order assertions.
- `kliring-form.test.tsx`: failed, 1 alias-order assertion.
- `swift-form.test.tsx`: failed, 1 missing-bank-name selection assertion.
- `internal-fund-transfer.test.tsx`: passed, 175/175.
- `online-transfer-form.test.tsx`: passed, 33/33, but contains the weakened submission assertion above.
- `form-benef.test.tsx`: passed, 27/27, but lacks selected-label ordering coverage.
- No focused RTGS form test currently exists.

No files were edited.

---

