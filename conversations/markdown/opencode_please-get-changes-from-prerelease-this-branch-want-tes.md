# Bring prerelease changes into current branch

- **ID**: `ses_000b6dae0ffeg4t5E8xHciAYAf`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 15:00:11`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

please get my changes from prerelease to this branch i want to test it local

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’ll inspect the current branch, worktree state, and available `prerelease` refs first so your local changes are preserved before integrating anything.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The worktree is clean on `dev`, and a local `prerelease` branch exists. I’m checking whether it tracks the remote and how it diverges from `dev` before merging the latest prerelease state.

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

I’ll fetch `origin/prerelease` and merge that fetched commit directly into `dev`; this avoids relying on a potentially stale local prerelease branch. If conflicts occur, I’ll resolve only the merge conflicts and leave unrelated work untouched.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

`dev` and `origin/prerelease` have substantial independent histories, producing conflicts across configuration, providers, localization, and tests. Since the goal is to test prerelease behavior locally, I’ll restart the merge using Git’s prerelease-favoring conflict strategy; non-conflicting `dev` changes remain included, while overlapping hunks take the prerelease version.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The merge completed successfully with prerelease favored in overlapping files and created a local merge commit. I’m doing a final integrity check for unresolved conflicts and confirming the resulting branch state before handing it back for local testing.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

Merged latest `origin/prerelease` into local `dev`.

- Merge commit: `4bd83c0e`
- Conflicts resolved in favor of prerelease
- Worktree is clean with no unresolved conflicts
- Nothing was pushed
- Tests were not run; ready for local testing via `yarn dev`

---

## Turn 11: User

i just want you to cherry-pick only dont take all commit

---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

I’ll remove only the merge commit I just created, restoring `dev` to its exact pre-merge commit. Then I’ll identify the specific prerelease commit rather than importing the full branch history.

---

## Turn 13: User

only my changes

---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The merge is removed and `dev` is back to its original clean state. The newest prerelease-only commit authored with your Git identity is `43097725` (`fix: update auth import not to lazy`), so I’ll cherry-pick only that commit.

---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

Cherry-picked only your latest prerelease commit:

- Source: `43097725`
- Local commit: `f38a50cf`
- `dev` is clean and ahead by 1 commit
- Nothing pushed

---

## Turn 18: User

good the last commit for remove lazy load for auth provider works perfectly and also makes my different issue who will fixed using the last commit please keep that changes on stagged changes

---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

I’ll remove the local cherry-pick commit while keeping its exact file changes staged, so the auth-provider lazy-load fix remains ready for your next commit.

---

## Turn 20: User

on this branch

---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 22: Assistant

The auth-provider changes are staged on the current `dev` branch across 4 files.

Two separate existing changes remain unstaged:

- `components/providers/AuthBridgeSync.ts`
- `components/providers/__tests__/AuthBridgeSync.test.tsx`

---

## Turn 23: User

then please audit the changes what happen for the other remotes while i implement the authprovider no lazy load . especialy who use shared

---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll audit the staged no-lazy-load change as a compatibility review: trace provider initialization, identify every remote consuming the shared auth/provider modules, and check runtime, bundle, and test regressions. I’ll keep the two unstaged `AuthBridgeSync` files outside the reviewed change unless they directly affect the staged behavior.

---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **task**: Tool: task
- **task**: Tool: task
- **task**: Tool: task

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>

The key remote-facing behavior is now clear: both auth packages remain non-eager Federation singletons, but their React providers mount synchronously and every remote runs beneath both contexts. I found one concrete regression already: the existing standalone Ordent provider test suite still assumes `next/dynamic` and now fails against the real package.

---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

Focused verification confirms the new composition tests pass, but the old `OrdentProviders` suite fails 3/3 because it now executes the real provider instead of its obsolete dynamic-import mock. There is also an environment mismatch: source and lockfile require Ordent `0.9.142-mfe`, while local `node_modules` contains `0.9.138-mfe`, so local runtime success does not yet validate the deployment version.

---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 32: Assistant

**Findings**
1. **High: existing Ordent tests are broken.** `OrdentProviders` now imports the package directly at `components/providers/OrdentProviders.tsx:2`, but its tests still mock `next/dynamic` at `components/providers/__tests__/OrdentProviders.test.tsx:52-79`. All 3 tests fail because the real provider accesses `localStorage`.

2. **High: local testing uses the wrong Ordent version.** `package.json:27` and `yarn.lock:534` require `0.9.142-mfe`, but `node_modules/@ordentco/addons-auth-provider/package.json:3` is `0.9.138-mfe`. Current local success does not validate the version that would be deployed.

3. **Medium: shared versions are not compatibility-checked.** Both providers are singleton shares with `requiredVersion: false` at `next.config.js:15-21`. Other remotes can share one provider instance, but an incompatible version will not be rejected. Missing exports, setters, or changed context behavior could fail at runtime.

4. **Medium: direct remote hooks remain risky.** Remotes importing `useAuth()` directly depend on the selected Federation singleton and matching provider context. Migrated remotes should consume `AuthBridgeSync` instead, as documented at `docs/ordent-to-bri-provider-migration.md:120-136` and `:191-193`.

5. **Medium: migration documentation no longer matches `_app`.** The documentation describes route-controlled provider mounting, but `_app` now always mounts both provider trees at `pages/_app.tsx:166-193`. Route selection currently affects which state `AuthBridgeSync` publishes, not which providers are mounted.

**Remote Impact**
- Removing `next/dynamic` does not set Federation `eager: true`. Both packages remain non-eager singleton shares.
- Provider order remains correct: BRI auth, BRI store, Ordent auth, Ordent store, remote subtree.
- Both providers become available immediately when the application subtree mounts.
- The subtree remains behind client-only `ModuleFederationProvider` at `pages/_app.tsx:60-65`, so provider rendering is still client-only. Static package evaluation now happens earlier.
- `AuthBridgeSync` already statically imported both packages, so the auth packages were already part of the startup dependency graph. The incremental bundle/startup impact should be limited.
- Known BRI-preferred remote routes are `/helpdesk`, `/plafon-pegadaian`, and `/systems` from `components/providers/auth-provider-migration-routes.ts:1`.
- `qui/providers` remains startup-critical because it wraps all page content at `pages/_app.tsx:167-192`.
- Exact shared-package consumers cannot be determined from this host repository alone. The runtime monitor derives them from Federation `useIn`/`from` metadata at `components/federation/monitor/auth/provider-versions.ts:32-49`.

**Verification**
- `AuthProviders.test.tsx`: passed
- `BRIProviders.test.tsx`: passed
- `OrdentProviders.test.tsx`: 3 failures
- Staged diff whitespace check: passed
- TypeScript is currently blocked by the separate unstaged `constants/features/registry-dev.ts:7` error
- Your four auth-provider files remain staged; unrelated files remain unstaged

---

## Turn 33: User

the userType and username still inside after logout and in landing page {
    "authSource": "ordentco",
    "username": "BA_MAYAN_M",
    "companyCode": "BRI",
    "companyName": "PT. BANK RAKYAT INDONESIA",
    "userType": "ba",
    "authorities": [],
    "menus": [],
    "productAuthorities": {
        "ALL": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ACCOUNT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "ANNOUNCEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BENEFICIARY_ACCOUNT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BG_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BG_ISSUING": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BG_MAPPING": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BG_MAPPING_DIGITAL": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BG_MONITORING": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CASH_POOLING": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "CBM": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "COMPANY": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "COMPANY_NEW": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "CUT_OFF": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "DEPOSITO": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "DIGI_GUARANTEE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "EXTERNAL_FUND_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "HELPDESK": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": true
        },
        "HOLIDAY": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "IMPORT_LC": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "INTERNAL_FUND_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MANAGE_DEPOSIT_ACCOUNT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MENU_APPEARANCE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MENU_LICENSE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MULTIPAYMENT_CREATE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MULTIPAYMENT_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "NOTIFICATION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "OPEN_NEW_TIME_DEPOSIT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PAYROLL_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PAYROLL_TRANSFER_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PROXY_REGISTRATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ROLE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "GROUP": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SSO_CLIENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SSO_COMPANY": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SSO_USER": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SUBSCRIPTION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SWIFT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SWIFT_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SWIFT_OVERVIEW": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SYSTEM": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "TICKET": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "TIME_DEPOSIT_ACCOUNTS": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "TRADE_IMPORT_ISSUANCE_LC": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "UPLOAD_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "UPLOAD_TRANSFER_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "USER": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "WITHDRAW_DEPOSIT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "WORKFLOW": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BI_FAST": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BULK_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ONLINE_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "KLIRING": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "RTGS": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ACCOUNT_STATEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ACCOUNT_SUMMARY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MPN": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PACKAGE_MANAGEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MASS_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MASS_TRANSFER_CREATE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MASS_TRANSFER_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MASS_TRANSFER_TRANSACTION_ACTION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "USER_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BPDP_SAWIT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "INQUIRY_BILLING_DJBC": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "INCOMING_DOCUMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "REPORT_STATEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BRIVA_INVOICE_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "CREATE_BRIVA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SMART_BILLING": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "LIST_DETAIL_BRIVA_INVOICE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "REPORT_BILLING_MANAGEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_ACCOUNT_STATEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_INTERNAL_FUND_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_EXTERNAL_FUND_TRANSFER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_TRANSACTION_ACTION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_TRANSACTION_INQUIRY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "UPLOAD_CC_PAYMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "UPLOAD_MANDATE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MAKER_CONFIRMATION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "DIRECT_DEBIT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "RECEIVABLE_ACTION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "UPLOAD_DIRECT_DEBIT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "COUNTER_PARTY_VALIDATION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MANDATE_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "RECEIVABLE_OVERVIEW": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "RECEIVABLE_REPORTS": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "DO_PERTAMINA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "NOTIONAL_POOLING": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "DO_PUPUK_INDONESIA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "NOTIONAL_POOLING_CLIENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "TAX_REPORT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "KAI": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BG_CANCELLATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SPB_JASARAHARJA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "MASS_BRIZZI": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SPT_JASARAHARJA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CASH_CARD_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "CASH_CARD": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SWIFT_GLOBAL": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "NOTIONAL_POOLING_GROUP": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "NOTIONAL_POOLING_REPORT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "NOTIONAL_POOLING_GROUP_CLIENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "NOTIONAL_POOLING_REPORT_CLIENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SWIFT_GLOBAL_TRANSACTION_ACTION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SWIFT_GLOBAL_TRANSACTION_INQUIRY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "RDN_TRANSACTION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PLAFON_PEGADAIAN": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_CARD_REQUEST": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_CARD_CONFIRMATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_CARD_RELATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_GENERATE_PIN": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_CARD_BLOCK": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CARD_MERCHANT_HISTORY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CASH_CARD_MANAGEMENT_SUMMARY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "CASH_CARD_SUMMARY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "DO_SEMEN_INDONESIA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "COMPANY_CODE_MAPPING": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BANK_MONITORING_TRANSACTION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "APPROVAL_SIGNATURE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "CASH_CARD_OVERVIEW": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "SPAN": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ACCOUNT_REGISTRATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "BRIVA_INVOICE_OVERVIEW": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "LC_EXPORT_ADVISE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "LC_EXPORT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "LOCAL_TAX_DKI_JAKARTA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "DF_PUPUK_MANAGEMENT": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "DF_PUPUK_INDONESIA": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "PENCAIRAN_BULOG": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "LOCAL_TAXES": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ACCOUNT_VALIDATION": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_R_TIMOR": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "GLOBAL_ACCOUNT_SUMMARY": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "USER_GUIDELINES": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": true,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "UPDATE_USER": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MITRA_ASURANSI": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MITRA_BG": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "AMEND_CANCEL_LC": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "BOND_PRICE": {
            "batch": false,
            "uploadFile": false,
            "template": false,
            "modify": false,
            "delete": false,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": false,
            "viewOnlineActivity": false,
            "dummyauthoritylevel": false,
            "dataEntry": false,
            "transferOverview": false,
            "anyAuthority": false,
            "allAuthority": false
        },
        "ECOSYSTEM_TPM": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "MANUAL_TOKEN_VERIFICATION": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": true
        },
        "HELPDESK_PORTAL": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": true
        },
        "LANGUAGE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "INQUIRY_AND_REPORTS": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "LANDING_PAGE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "SUPPLY_CHAIN_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "HELPDESK_TICKET": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": true
        },
        "RMA_MANAGEMENT": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "REGISTRATION_ONLINE": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        },
        "RECEIVABLE_INQUIRY": {
            "batch": true,
            "uploadFile": true,
            "template": true,
            "modify": true,
            "delete": true,
            "verify": false,
            "approve": false,
            "release": false,
            "downloadReport": true,
            "viewOnlineActivity": true,
            "dummyauthoritylevel": false,
            "dataEntry": true,
            "transferOverview": false,
            "anyAuthority": true,
            "allAuthority": false
        }
    },
    "isAuthoritiesReady": true,
    "packageVersions": {
        "bri": "0.10.6-mfe",
        "ordentco": "0.9.142-mfe",
        "ordent": "0.9.142-mfe"
    },
    "runningVersions": [
        {
            "remote": "qcash-global-component",
            "versions": [
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.140-mfe"
                }
            ]
        },
        {
            "remote": "qcash-ui",
            "versions": [
                {
                    "source": "bri",
                    "pkg": "@bri/addons-auth-provider",
                    "version": "0.10.6-mfe"
                },
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.138-mfe"
                }
            ]
        },
        {
            "remote": "qcash-ui-dashboard",
            "versions": [
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.138-mfe"
                },
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.137-mfe"
                }
            ]
        },
        {
            "remote": "qcash-ui-header-footer",
            "versions": [
                {
                    "source": "bri",
                    "pkg": "@bri/addons-auth-provider",
                    "version": "0.10.6-mfe"
                },
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.138-mfe"
                },
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.142-mfe"
                }
            ]
        },
        {
            "remote": "qcash-ui-live-chat",
            "versions": [
                {
                    "source": "ordentco",
                    "pkg": "@ordentco/addons-auth-provider",
                    "version": "0.9.119-mfe"
                }
            ]
        },
        {
            "remote": "qcash-ui-sistem-management",
            "versions": [
                {
                    "source": "bri",
                    "pkg": "@bri/addons-auth-provider",
                    "version": "0.10.6-mfe"
                }
            ]
        }
    ]
}

---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The payload shows logout clears some arrays but leaves the identity and product-authority snapshot publishable, so `AuthBridgeSync` is rebuilding the bridge from stale Ordent context after navigation to the landing page. I’ll trace the logout event, bridge clearing, and publish conditions, then add the smallest guard that prevents logged-out state from republishing authenticated identity.

---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The failure path is confirmed: `clearBridgeData()` writes a Guest snapshot, but the publishing effect later accepts `isAuthoritiesReady` from stale provider context even when `access-token` is gone. I’ll enforce the security invariant that no authenticated bridge snapshot may exist without the current token, and the logout event will remove the token before clearing so route changes cannot republish the old identity.

---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 41: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 42: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

The new logout/navigation regression passes; one pre-existing unmount test now needs to explicitly model logout because authenticated tests correctly receive a token in setup. I’ll adjust that test to remove the token before unmount, matching the component’s existing cleanup contract.

---

## Turn 43: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 44: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 45: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 46: Assistant

Fixed the stale logout bridge state.

- Logout now removes `access-token` before clearing bridge data.
- Bridge publishing now requires an active token, regardless of stale provider context.
- Landing-page navigation can no longer republish the previous `username`, `userType`, authorities, menus, or product authorities.
- Added regression coverage for stale context after logout and route navigation.
- Focused tests: 17 passed.
- ESLint and diff checks passed.

The original four no-lazy-load files remain staged. `AuthBridgeSync` and its test remain unstaged with the existing work.

---

## Turn 47: User

please check the qcash-ui and qcash-ui-header-footer changes are already cover with unit test and make sure the unit test > 65%

---

## Turn 48: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll inspect the changed files in both repositories, map each behavioral change to tests, then run focused coverage and add only the missing cases needed to exceed 65%. I’ll preserve the current staged/unstaged separation in `qcash-ui`.

---

## Turn 49: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob

</details>



---

## Turn 50: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 51: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>

Both repositories already contain tests for the changed paths. I’m validating whether those tests exercise the logout and bridge behavior rather than only rendering, and I’ll measure coverage specifically for the changed production files so the 65% requirement is meaningful.

---

## Turn 52: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 53: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

Focused changed-file coverage already exceeds the target in both repositories: `qcash-ui` is 92.14% statements / 73.23% branches, and header-footer is 95% statements / 86.2% branches / 66.66% functions. One obsolete `qcash-ui` test suite still mocks `next/dynamic` after the provider became static, so I’ll update that test to mock the provider package directly and ensure the complete provider test set passes.

---

## Turn 54: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 55: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 56: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

All changed-file tests now pass and both repositories exceed 65%. Header-footer’s aggregate function coverage is only narrowly above the threshold because its fallback callbacks are not invoked, so I’ll cover those no-op contracts as well to give the result safe margin rather than relying on 66.66%.

---

## Turn 57: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 58: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 59: Assistant

Both repositories now have unit coverage above 65% for the changed production files.

**qcash-ui**
- 20 tests passed
- Statements: 92.14%
- Branches: 73.23%
- Functions: 93.33%
- Lines: 94.41%
- Updated the obsolete `OrdentProviders` test to mock static package imports correctly.

**qcash-ui-header-footer**
- 7 tests passed
- Statements: 95%
- Branches: 86.2%
- Functions: 91.66%
- Lines: 98.14%
- Covered bridge reset events, logout integration, authority rehydration, fallback data, and fallback callbacks.

`qcash-ui` ESLint passed. Header-footer ESLint could not start because its installed `eslint-plugin-sonarjs` is incompatible with the installed ESLint API, unrelated to these changes. Diff checks passed in both repositories.

---

## Turn 60: User

qcash-ui still had unit test error

---

## Turn 61: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll run the full `qcash-ui` test suite, identify every failure introduced or exposed by these changes, and fix the affected tests without touching unrelated staged work.

---

## Turn 62: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 63: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The full suite exposes 96 failures, all from one root cause: this Node/Vitest environment defines `localStorage` as unavailable, while the shared test setup does not install a storage implementation. I’ll add a standards-compatible in-memory storage shim in `vitest-setup.js`; this fixes the environment centrally instead of patching 96 tests individually.

---

## Turn 64: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 65: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 66: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 67: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The full suite now passes: 132 files and 1,374 tests. I’m running the complete repository coverage command now, not just focused files, to verify the actual project-wide percentage against the 65% target.

---

## Turn 68: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 69: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 70: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 71: Assistant

The full `qcash-ui` unit-test error is fixed.

- Test files: 132 passed
- Tests: 1,374 passed
- Statements: 84.73%
- Branches: 78.83%
- Functions: 86.5%
- Lines: 86.14%
- ESLint passed
- Diff checks passed

Root cause was unavailable `localStorage`/`sessionStorage` in the current Node/Vitest environment. Added a shared in-memory implementation in `vitest-setup.js` and corrected the obsolete `OrdentProviders` dynamic-import mock.

---

## Turn 72: User

is there already provide unit test for the change

---

## Turn 73: Assistant

Yes. Unit tests cover the changes in both repositories.

**qcash-ui**
- `components/providers/__tests__/AuthProviders.test.tsx`
  - Static BRI and Ordent provider composition
  - Provider order and configuration
- `components/providers/__tests__/AuthBridgeSync.test.tsx`
  - Logout bridge reset
  - Stale `username` and `userType` prevention
  - Landing-page navigation after logout
  - Storage cleanup
- `components/providers/__tests__/OrdentProviders.test.tsx`
  - Updated for direct imports

**qcash-ui-header-footer**
- `layouts/default/__tests__/AuthBridge.test.ts`
  - Guest fallback
  - Authority rehydration
  - Logout reset event
- `layouts/default/container/header/components/__tests__/ListItem.test.tsx`
  - Logout dispatches the bridge reset event

All 1,374 `qcash-ui` tests pass, with project coverage above 78% across every metric.

---

## Turn 74: User

please create me summary that explain the no lazy load authentication might be affected what in simple indonesia

---

## Turn 75: Assistant

**Ringkasan Perubahan No Lazy Load Authentication**

Sebelumnya, `AuthProvider` dari BRI dan Ordent dimuat secara bertahap menggunakan lazy loading. Sekarang keduanya langsung dimuat ketika aplikasi mulai berjalan.

**Dampak Positif**
- Authentication provider lebih cepat tersedia.
- Mengurangi jeda saat halaman atau remote module pertama kali dibuka.
- Menghindari kondisi remote module berjalan sebelum authentication provider siap.
- BRI dan Ordent provider tetap memiliki urutan dan konfigurasi yang sama.

**Bagian yang Mungkin Terdampak**
- Waktu awal aplikasi dapat sedikit bertambah karena provider langsung dimuat.
- Semua remote yang memakai shared authentication akan menggunakan provider singleton yang dipilih Module Federation.
- Remote dengan versi authentication provider berbeda mungkin mengalami ketidakcocokan.
- Remote yang langsung menggunakan `useAuth()` harus memastikan provider dan versinya sesuai.
- Remote yang sudah memakai `AuthBridge` lebih aman karena menerima data authentication dari host.

**Perbaikan Logout**
- Setelah logout, data pengguna lama tidak lagi diteruskan ke remote.
- `username`, `userType`, authorities, menu, dan product authorities diubah ke kondisi Guest.
- Data lama tidak akan muncul kembali ketika pengguna masuk ke landing page.
- Cache authentication dan token dibersihkan saat logout.

**Pengujian**
- Provider BRI dan Ordent berhasil dimuat tanpa lazy loading.
- Proses login, logout, bridge synchronization, dan perpindahan ke landing page sudah diuji.
- Seluruh 1.374 unit test `qcash-ui` berhasil.
- Coverage keseluruhan berada di atas 78%.
- Unit test `qcash-ui-header-footer` juga berhasil dengan coverage di atas 86%.

---

## Turn 76: User

please explain to me is the current implementation already handle ssr false defaulty ?

---

## Turn 77: Assistant

Ya, **secara rendering implementasi saat ini masih berjalan seperti `ssr: false`**, tetapi bukan karena `BriProviders` memiliki pengaturan tersebut.

Alasannya:

- `BriProviders` sekarang menggunakan static import di `pages/_app.tsx:19`.
- Namun seluruh provider berada di dalam `ModuleFederationProvider` pada `pages/_app.tsx:161`.
- `ModuleFederationProvider` masih dimuat dengan:

```tsx
const ModuleFederationProvider = dynamic(
  () => import("@/components/federation/provider"),
  { ssr: false }
);
```

Karena parent menggunakan `ssr: false`, seluruh child berikut hanya dirender di browser:

```text
ModuleFederationProvider
└── BriProviders
    ├── BRI AuthProvider
    ├── BRI GlobalStoreProvider
    ├── Ordent AuthProvider
    └── Ordent GlobalStoreProvider
```

**Catatan Penting**
- `BriProviders` sendiri **tidak lagi memiliki `ssr: false`**.
- Package authentication tetap dapat di-import atau dievaluasi saat proses build/server.
- React provider-nya tidak dirender saat SSR karena masih dilindungi oleh parent `ModuleFederationProvider`.
- Jika `BriProviders` nanti dipindahkan keluar dari `ModuleFederationProvider`, provider tersebut dapat ikut dirender saat SSR dan harus diperiksa kembali.

Jadi, implementasi sekarang aman secara rendering client-only, tetapi perlindungannya berasal dari parent, bukan dari konfigurasi default `BriProviders`.

---

## Turn 78: User

please check also issuance-lc that might be still had navigation error because coupling routing. on that repo already use new qui design system that use tanstack router please check is the current fix can also fix that issue ?

---

## Turn 79: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I’ll trace navigation end-to-end across `qcash-ui`, `qcash-ui-issuance-lc`, and the shared `qui` package: router ownership, provider nesting, remote entry mounting, and any direct Next/TanStack coupling. Then I’ll determine whether the no-lazy auth change actually addresses the navigation error or only changes initialization timing.

---

## Turn 80: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **task**: Tool: task
- **task**: Tool: task
- **task**: Tool: task
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 81: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

The first important finding is that the checked-out issuance-LC code does not use TanStack Router; it uses TanStack Query and still navigates through Next Pages Router. I’m checking its current uncommitted `_app`, create/edit, and query-client changes now to ensure this is not a branch-specific migration hidden in the working tree.

---

## Turn 82: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 83: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 84: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 85: Assistant

**Kesimpulan**

Perubahan authentication menjadi no-lazy-load **tidak sepenuhnya memperbaiki masalah navigasi Issuance LC**.

Perubahan tersebut hanya dapat membantu jika error terjadi karena:

- Auth provider belum siap ketika Issuance LC dibuka.
- `useAuth()` masih mendapatkan context Guest.
- Token atau authority terlambat tersedia.
- Lazy auth chunk gagal dimuat.

Namun, perubahan ini tidak memperbaiki coupling routing di dalam repository Issuance LC.

**Temuan Penting**

Issuance LC saat ini **tidak menggunakan TanStack Router**. Dependency yang digunakan adalah:

```json
"@tanstack/react-query": "^5.83.0"
```

Itu adalah TanStack Query untuk pengambilan dan cache data, bukan router.

Navigasi Issuance LC masih menggunakan:

- `next/router`
- `next/link`
- `router.push()`
- `router.replace()`
- `window.location`

Contohnya:

- `containers/form-new-skin.tsx:1244`
- `containers/detail/view-new-skin.tsx:491`
- `components/table/ongoing-columns.tsx:53`
- `components/dropdown.tsx:129`

**Masalah Navigasi yang Masih Ada**

1. **Next Link membungkus QUI Button**

```tsx
<Link href="/import-lc/issuance-lc/create">
  <QuiButton>Create New Issue</QuiButton>
</Link>
```

Ini berpotensi menghasilkan `<a><button /></a>`, yaitu nested interactive element. Perilaku klik dapat berbeda tergantung implementasi QUI.

Referensi: `containers/new-skin.tsx:65-79`.

2. **Dropdown detail terlalu bergantung pada nested Link**

`DropdownQui.Menu` menjalankan `handleAction("detail")`, tetapi tidak ada case `detail` pada switch.

Referensi:

- `components/dropdown.tsx:39-61`
- `components/dropdown.tsx:129-136`

Menu bisa tertutup sebelum nested Link menyelesaikan navigasi.

3. **Host dan remote memiliki base path berbeda**

- Host: `/v2`
- Issuance LC standalone: `/issuance-lc-v2`

Issuance LC menggunakan path tanpa prefix:

```tsx
router.push("/import-lc/issuance-lc");
```

Ini hanya aman jika `next/router` remote benar-benar menggunakan router milik host.

4. **Redirect langsung mengabaikan base path**

```tsx
window.location.replace("/main-page");
window.location.href = "/main-page?logout=true";
```

Referensi: `services/custom-axios.ts:71-127`.

Redirect tersebut bisa keluar dari `/v2`.

5. **Pemeriksaan pathname dengan query tidak pernah benar**

```tsx
pathname === "/main-page?logout=true"
```

`window.location.pathname` tidak berisi query string. Seharusnya query diperiksa melalui `window.location.search`.

6. **Dynamic route bergantung langsung pada host Next Router**

Detail dan edit membaca:

```tsx
router.query.transactionEventId
```

Ini bekerja karena host memiliki route Next dengan nama parameter yang sama. Jika host pindah ke TanStack Router, nilai tersebut tidak otomatis tersedia.

**Tentang QUI**

`qui/providers` sudah dipasang di host dan Issuance LC, tetapi tidak menerima navigation adapter:

```tsx
<QuiProviders>{children}</QuiProviders>
```

Tidak ada prop seperti:

```tsx
<QuiProviders navigate={router.push} />
```

Jadi QUI provider tidak otomatis mengubah navigasi Next menjadi TanStack Router.

**Hasil Akhir**

- No-lazy auth: dapat memperbaiki auth initialization race.
- Logout bridge fix: dapat mencegah stale user setelah logout.
- Navigation coupling Issuance LC: belum diperbaiki.
- TanStack Router: belum digunakan oleh Issuance LC.
- Masalah utama masih pada penggunaan langsung Next Router, hard-coded path, base path, dan nested QUI button di dalam Next Link.

---

## Turn 86: User

how about the Outlet on the qui provider ?

---

## Turn 87: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>

I’ll verify whether the deployed QUI type artifacts expose an `Outlet` or TanStack router provider contract, since the local `qui` checkout does not contain one.

---

## Turn 88: Assistant

`Outlet` tidak ditemukan pada source atau type artifact yang tersedia di:

- `qcash-ui`
- `qcash-ui-issuance-lc`
- checkout lokal `qui`

Tidak ada dependency atau penggunaan:

- `@tanstack/react-router`
- `RouterProvider`
- `Outlet`
- `createRootRoute`
- `useNavigate`

**Jika deployed `qui/providers` menggunakan `<Outlet />`**

Itu berpotensi menjadi sumber navigation error karena `<Outlet />` hanya dapat bekerja di bawah TanStack `RouterProvider` dan route tree yang aktif.

Saat ini host memasang QUI seperti ini:

```tsx
<QuiProviders>
  <AuthBridgeSync />
  <PageComponent />
</QuiProviders>
```

Host masih menggunakan Next Pages Router, bukan TanStack Router. Jadi jika `QuiProviders` hanya merender:

```tsx
<Outlet />
```

maka kemungkinan masalahnya:

- `PageComponent` dari Next tidak dirender.
- Muncul error karena TanStack router context tidak tersedia.
- Outlet tidak menemukan matched route.
- Router TanStack dan Next dapat memiliki URL/history yang berbeda.
- Next `routeChangeComplete` tidak terpanggil.
- Host tidak mengganti remote module setelah navigasi TanStack.
- Base path `/v2` dapat terlewati.

**Kontrak yang aman untuk host saat ini**

`QuiProviders` seharusnya tetap merender `children`:

```tsx
function QuiProviders({ children }) {
  return <DesignSystemProvider>{children}</DesignSystemProvider>;
}
```

Bukan menggantinya dengan:

```tsx
function QuiProviders() {
  return <Outlet />;
}
```

Jika QUI memang membutuhkan TanStack Router, provider harus memiliki mode yang jelas:

```tsx
<QuiProviders mode="next">{children}</QuiProviders>
```

atau menerima navigation adapter dari Next:

```tsx
<QuiProviders
  navigate={(path) => router.push(path)}
  replace={(path) => router.replace(path)}
>
  {children}
</QuiProviders>
```

**Kesimpulan**

No-lazy authentication tidak dapat memperbaiki masalah `Outlet`. Jika deployed `qui/providers` menggunakan TanStack `<Outlet />`, masalah tersebut harus diperbaiki pada kontrak QUI provider atau dengan adapter Next Router. Berdasarkan source yang tersedia, host dan Issuance LC belum menyediakan TanStack `RouterProvider`, jadi penggunaan `<Outlet />` tidak kompatibel dengan implementasi saat ini.

---

