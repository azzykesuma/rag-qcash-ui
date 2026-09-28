# Summarize auth related PRs (@explore subagent)

- **ID**: `opencode-ba9025ef11c475a8aa66f2e27dfac8146f4fec1b64c3ed5c038c2cbbaccbf4b2`
- **Source Tool**: `opencode`
- **Date**: `2026-09-28 09:43:11`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only. Inspect these three saved Bitbucket pull-request result files and extract only PRs relevant to qcash-ui-header-footer auth-provider migration, BRI vs Ordent, logout, stale session/menu, or centralized auth. Return PR id/title/state/source/target/updated date and any especially relevant description. Deduplicate repeated PRs. Files: ~\.local\share\opencode\tool-output\tool_0e5e4c0c2001JtqbjVyIAptb3M ; ~\.local\share\opencode\tool-output\tool_0e5e4c15a0016rJ5VXMpxf2tWB ; ~\.local\share\opencode\tool-output\tool_0e5e4c120001AInjC2753CJMLn. Do not edit anything.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

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

- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

Deduplicated by PR ID; release-train PRs that merely repeated these commits were excluded.

### Auth-provider migration / centralized auth

| PR | Title | State | Source → Target | Updated (UTC) |
|---:|---|---|---|---|
| 701 | fix: revert auth provider version [NQLA-71406] | MERGED | `feat-revert-auth` → `prerelease` | 2026-09-28 |
| 707 | fix: revert auth provider version [NQLA-71406] | MERGED | `feat-revert-auth-dev` → `dev` | 2026-09-28 |
| 669 | refactor(auth): remove useAuthBridge and migrate to @bri/addons-auth-provider | MERGED | `refactor/remove-useauthbridge` → `dev` | 2026-09-17 |
| 660 | feat(auth): centralize auth guard provider and remove legacy guard routines | OPEN | `auth-guard-centralized-v1.648.4-release` → `v1.648.4-release` | 2026-09-16 |
| 659 | feat(auth): centralize auth guard provider and remove legacy guard routines | OPEN | `auth-guard-centralized-v1.671.1-release` → `v1.671.1-release` | 2026-09-16 |
| 657 | chore(deps): upgrade @bri/addons-auth-provider to ^0.10.7-mfe and log productAuthority.TEST_PRODUCT | MERGED | `chore/upgrade-addons-auth-provider-0.10.7-prerelease` → `prerelease` | 2026-09-10 |
| 656 | chore(deps): upgrade @bri/addons-auth-provider to ^0.10.7-mfe | MERGED | `chore/upgrade-addons-auth-provider-0.10.7-mfe` → `dev` | 2026-09-10 |
| 651 | feat: updated yarn lock file (bri auth provider) | MERGED | `feat/merchant-registration-prerelease` → `prerelease` | 2026-09-09 |
| 638 | fix: add @next/bundle-analyzer devDependency and update auth bridge | MERGED | `fix/remove-auth-bridge-v001` → `v.0.0.1-release` | 2026-09-05 |
| 634 | chore: merge v.0.0.1-release into prerelease | MERGED | `v.0.0.1-release` → `prerelease` | 2026-09-05 |
| 637 | refactor(auth): remove useAuthBridge and migrate to bri auth provider (v.0.0.1) | MERGED | `fix/remove-auth-bridge-v001` → `v.0.0.1-release` | 2026-09-05 |
| 636 | refactor(auth): remove useAuthBridge and migrate to bri auth provider | OPEN | `fix-bri-auth-provider-only` → `dev` | 2026-09-05 |
| 628 | fix: use BRI auth provider in container | DECLINED | `fix-bri-auth-provider-only` → `dev` | 2026-09-02 |
| 627 | feat: centralize auth guard with bri addons auth provider | DECLINED | `feat/auth-guard-centralization` → `dev` | 2026-08-31 |
| 476 | fix: update bri provider version | MERGED | `rxn-auth-test` → `dev` | 2026-07-15 |
| 165 | chore: update @ordentco/addons-auth-provider to version 0.9.126-mfe | MERGED | `feature/update-version-auth-provider-126-mfe` → `master` | 2025-08-01 |
| 127 | chore: update auth provider | MERGED | `update-package-json` → `1.0.147` | 2025-05-17 |
| 118 | fix: remove @bri/addons-auth-provider package | MERGED | `fix/remove-bri-auth-provider` → `master` | 2025-04-25 |
| 111 | fix: migrate all functions to use default package @ordentco/addons-auth-provider except _app.tsx | MERGED | `fix/migrate-addons-auth-provider` → `master` | 2025-04-09 |
| 108 | Fix/migrate addons auth provider | MERGED | `fix/migrate-addons-auth-provider` → `master` | 2025-03-24 |

Especially relevant descriptions:

- **#669:** Removed legacy `useAuthBridge`/`AuthBridge`, removed Ordent from module-federation sharing, and migrated menu hooks, modals, and layouts directly to BRI.
- **#659/#660:** Centralize auth with `@bri/addons-auth-provider ^0.10.7-mfe`, wrap `_app.tsx` with `AuthProvider` and `GlobalStoreProvider`, replace Ordent imports, and remove legacy guard routines.
- **#638:** Refactored the auth bridge in favor of BRI while also fixing the CI bundle-analyzer dependency.
- **#634:** Cherry-picked guard removal and BRI-provider migration into `prerelease`.
- **#637:** Replaced `useAuthBridge` and legacy `AuthBridge` with the singleton BRI provider.
- **#636:** Migrates header/footer auth consumers to BRI and removes obsolete bridge components.
- **#628:** Uses BRI directly in the default layout container and removes the legacy bridge/fallback path.
- **#627:** Centralizes the guard with BRI and removes Ordent references across layouts, hooks, and modals.
- **#108:** Applied a toggle for the then-new auth-provider implementation.

### Logout / stale menu or auth state

| PR | Title | State | Source → Target | Updated (UTC) |
|---:|---|---|---|---|
| 595 | fix: auth bridge reset while logout | DECLINED | `rxn-bridge-fix` → `dev` | 2026-09-16 |
| 437 | Feature/victor NQLA-48946 cherry pick | MERGED | `feature/victor-NQLA-48946-cherry-pick` → `v1.624.0-release` | 2026-07-01 |
| 414 | fix: invalidate menu cache on logout | MERGED | `feature/victor-NQLA-48946-prerelease-sync` → `prerelease` | 2026-06-23 |
| 413 | fix: invalidate menu cache on logout | MERGED | `feature/victor-NQLA-48946-dev-sync` → `dev` | 2026-06-23 |
| 246 | fix: add / in logout destination | MERGED | `fix/bug` → `master` | 2025-11-07 |
| 201 | fix: bug logout di level submenut | MERGED | `feat/redirect-to-landing-page` → `master` | 2025-09-23 |
| 188 | refactor: change logout flow to non-blocking by removing await on logoutLog | MERGED | `bugfix/handle-button-disabled-logout` → `master` | 2025-09-10 |
| 184 | refactor logout | DECLINED | `refactor-logout` → `master` | 2025-08-21 |
| 185 | refactor: enhance logout process with loading state and disable button | MERGED | `bugfix/handle-button-disabled-logout` → `master` | 2025-08-21 |
| 102 | CMSO-85372 (feat: remove toggle logoutLog, default to prod) | MERGED | `feature/config` → `master` | 2025-03-18 |
| 101 | CMSO-85372 (feat: toggle logoutLog to prod) | MERGED | `feature/config` → `master` | 2025-03-18 |
| 100 | CMSO-85372 (fix: toggle logout logging to preploy) | MERGED | `feature/config` → `master` | 2025-03-13 |
| 99 | Feature/config | MERGED | `feature/config` → `master` | 2025-03-04 |

Especially relevant descriptions:

- **#437:** Includes menu stale-time/fetch-once changes and menu-cache invalidation on logout.
- **#188:** Calls `logoutLog`, briefly delays so the token remains available, clears local/session storage, redirects immediately without awaiting logging, then lets logging complete in the background.
- **#184:** Refactored the logout mechanism; declined.
- **#99:** Added the logout logging API and its configuration toggle.

No PR explicitly used the phrase **“stale session”**. The closest matches are **#595** for resetting bridged auth state during logout and **#413/#414/#437** for preventing stale menu cache after logout.

Inspected:

- `~\.local\share\opencode\tool-output\tool_0e5e4c0c2001JtqbjVyIAptb3M`
- `~\.local\share\opencode\tool-output\tool_0e5e4c15a0016rJ5VXMpxf2tWB`
- `~\.local\share\opencode\tool-output\tool_0e5e4c120001AInjC2753CJMLn`

---

