# Inspect header release (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_00205c9d8ffejpp9rSTDdczEQF`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 08:54:21`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only production investigation. In D:\work\qcash-ui-header-footer inspect exactly git ref refs/remotes/origin/v1.630.3-release-com (commit ee409cab), not checked-out dev and do not edit/checkout. Trace auth dependency version(s), provider/guard/useAuth sites, Module Federation config shared singleton behavior, exposures consumed by qcash-ui, qcash-global-component consumption, and error handling. Determine whether header/footer can get a different Ordent AuthContext from the shell and whether it invokes default callbacks that crash. Return exact ref:path:line references and concrete production failure paths.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’ll verify the target commit without changing either worktree, then inspect its tree with `git show`/`git grep` and correlate the relevant shell code in `qcash-ui`. All conclusions will distinguish committed ref content from any checked-out files.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>

The target ref resolves exactly to `ee409cab7d8193db987ae6b0f3f27322ccc4a6b5`; the checked-out header/footer `dev` has unrelated dirty files and remains untouched. A matching `origin/v1.630.3-release-com` also exists in the shell repo, so I’ll use that committed shell ref rather than its checked-out prerelease branch for the production correlation.

---

## Turn 5: Assistant

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

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
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
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **context7_query-docs**: Tool: context7_query-docs
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **context7_query-docs**: Tool: context7_query-docs

</details>

The production shell pins the same Ordent package (`0.9.136-mfe`) but neither side marks it `singleton`; `requiredVersion: false` only removes compatibility enforcement. I’m now checking the matching global-component release and the shell’s boundaries because a split context is possible only under particular share-scope/load conditions, and load failures and callback failures are handled differently.

---

## Turn 13: Assistant

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

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

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

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 23: Assistant

## Objective
- Perform a read-only production investigation of `D:\work\qcash-ui-header-footer` at `refs/remotes/origin/v1.630.3-release-com`, tracing auth context identity, Module Federation sharing, shell/global-component consumption, error handling, and concrete crash paths.

## Important Details
- Header/footer target is exactly `ee409cab7d8193db987ae6b0f3f27322ccc4a6b5`; no checkout or edits permitted.
- Header/footer worktree remains on dirty `dev`; unrelated modified files:
  - `layouts/default/container/menu/components/LcImportModal.tsx`
  - `layouts/default/container/menu/components/MenuModalProvider.tsx`
- Shell correlation uses `D:\work\qcash-ui` at `refs/remotes/origin/v1.630.3-release-com` commit `b6ce2a6bf3cf6564b0aa3f56dd396e4fe84b5a1b`, not checked-out `prerelease-fix-route`.
- `D:\work\qcash-global-component` has no `origin/v1.630.3-release-com`; nearest inspected ref is `origin/v1.626.5-release-com` commit `5e51ff23aaee07b7ae429b50dfc4eea4593c3976`. Do not present it as proven deployed production content.
- Shell provider selection via `config.toggleAddonsAuthProviderMigration` is unresolved and likely decisive: if production selects `@bri/addons-auth-provider`, header/footer’s `@ordentco/addons-auth-provider` has no matching provider regardless of sharing.
- Current shell `node_modules` contains auth `0.9.140-mfe`, so it must not be treated as target-ref evidence. Header/footer `node_modules` contains matching locked artifact `0.9.136-mfe`.
- `rg` is unavailable (`The term 'rg' is not recognized...`); use `git grep`, `git show`, or PowerShell `Select-String`.

## Work State
### Completed
- Verified header/footer ref and commit without changing worktrees.
- Verified auth lock versions:
  - Header/footer `refs/remotes/origin/v1.630.3-release-com:yarn.lock:1071-1074` → `0.9.136-mfe`.
  - Shell `refs/remotes/origin/v1.630.3-release-com:yarn.lock:696-700` → same `0.9.136-mfe` tarball hash.
  - Nearest global-component ref `yarn.lock:3163` starts `0.9.126-mfe`.
- Header/footer has no production `AuthProvider` or `guard()` site; it relies on its host provider.
- Known header/footer `useAuth` sites:
  - `hooks/useLocalTaxMenu.tsx:3,17`
  - `hooks/useMenu.tsx:2,21`
  - `layouts/default/container/header/components/ListItem.tsx:2,49`
  - `layouts/default/container/menu/components/LcExportModal.tsx:2`
- Header/footer federation configuration:
  - `module-federation.config.mjs:3-4` names remote and entry.
  - `:6-7` hardcodes global-component dev URL `http://internal-service.example.com/global-component/_next/static/chunks/remoteEntry.js`.
  - `:9-12` exposes only `./default` and `./globals`.
  - `:17` shares `@ordentco/addons-auth-provider` as `{ requiredVersion: false, eager: false }`; `singleton` is omitted.
- Shell also shares Ordent auth without `singleton`; exact shell line still needs extraction.
- Runtime evidence confirms omitted `singleton` defaults to `false`:
  - `D:\work\qcash-ui\node_modules\@module-federation\runtime-core\dist\index.esm.js:755-760`
  - Next federation runtime plugin changes unspecified share strategies to `loaded-first` at `...\runtimePlugin.js:31-41`.
  - `requiredVersion: false` accepts the selected share entry at header runtime `share.esm.js:797-825`; it does not itself enforce singleton identity.
- Header/footer’s exact `0.9.136-mfe` package artifact defines:
  - defaults at `src/auth.tsx:129-215`
  - throwing default `guard` at `:150-152`
  - throwing default async `logout` at `:159-161`
  - `AuthContext` at `:221`
  - provider at `:1110-1112`
  - `useAuth` at `:1117`
  - runtime equivalents at `dist/src/auth.js:95-208,1038-1042`
  - error string: `Function not implemented.`
- Therefore, if header/footer resolves a different Ordent module instance or is hosted only under the BRI provider, `useAuth()` returns defaults. A logout click can call the default async `logout()` and reject with `Function not implemented.`; exact `ListItem` awaiting/catch behavior remains to verify.
- Shell provider candidates found in `components/providers/BRIProviders.tsx`:
  - BRI dynamic import/provider at `:8-15`
  - Ordent dynamic import/provider begins `:21-26`
- Shell guard handling found:
  - `components/onboarding-tour/notif-onboarding-tour.tsx:86` awaits `guard()` inside catch at `:93-95`.
  - `hooks/use-modal-session-expired.tsx:243-248` has caught and uncaught branches.
  - `config.ts:676-681` sets `toggleUpdateLoginService` true through production.
- Shell registers global-component and header/footer as global remotes:
  - `constants/features/registry.ts:37-47`
- Shell canonical header/footer consumer:
  - `components/ui/SecureFeatureWrapper.tsx:7-10` loads `qcash-ui-header-footer/default`.
  - Literal `qcash-ui-header-footer/default` occurs in 422 shell files.
- Header/footer consumes global-component through:
  - `components/component-global/index.tsx:1,6,12`
  - `qcash-global-component/external-transfer-modal`
  - `qcash-global-component/response-modal`
  - `MenuModalProvider.tsx:16,165`
  - `LcExportModal.tsx:13,109`
  - `LcImportModal.tsx:1,132`
- `config.ts:28` disables `modalEFTV2` for preploy, piloting, and production, so the global external-transfer modal path is not normally active there.
- Header/footer standalone `_app` registers global-component using environment federation URL at `pages/_app.tsx:20-21` and loads globals at `:26`.

### Active
- Determine the production value and branch direction of `toggleAddonsAuthProviderMigration` in `components/providers/BRIProviders.tsx` and `config.ts:472+`.
- Complete exhaustive non-test `useAuth`, callback invocation, provider nesting, and shell shared-config line inventories.
- Verify `ListItem` logout sequencing, awaiting, rejection handling, and whether React/remote boundaries can catch the async failure.
- Distinguish normal same-scope behavior—where shell/header both use loaded `0.9.136-mfe`—from concrete split paths involving BRI provider selection, separate runtime/share scopes, or global-component `0.9.126-mfe`.
- Trace shell loader/retry/error-boundary behavior and whether header/footer’s hardcoded global remote is used when consumed as an exposed remote.

### Blocked
- No matching `qcash-global-component` `origin/v1.630.3-release-com`, so exact deployed global-component implementation/version cannot be proven locally.
- Runtime deployment state, actual production remote entries, share-scope contents, and load order have not been observed.
- Several command outputs were truncated; exact line ranges still need targeted `git show`/`git grep -e "^"` extraction.

## Next Move
1. Extract complete line-numbered content for shell `components/providers/BRIProviders.tsx`, `pages/_app.tsx`, `config.ts:465-485`, and the Ordent/BRI entries in `next.config.js`; resolve which provider wraps production children.
2. Extract full header/footer `ListItem.tsx`, all non-test auth imports/calls, shell `SecureFeatureWrapper`/`RemoteErrorBoundary`/loader files, then formulate exact normal and failure paths with `ref:path:line` citations.

## Relevant Files
- `D:\work\qcash-ui-header-footer\module-federation.config.mjs`: remote, exposures, shared auth, hardcoded global URL.
- `D:\work\qcash-ui-header-footer\package.json`: auth/runtime dependencies.
- `D:\work\qcash-ui-header-footer\yarn.lock`: exact auth and federation package versions.
- `D:\work\qcash-ui-header-footer\pages\_app.tsx`: standalone runtime registration and global styles loading.
- `D:\work\qcash-ui-header-footer\layouts\default\index.tsx`: exposed default layout entry.
- `D:\work\qcash-ui-header-footer\layouts\default\container\header\components\ListItem.tsx`: `useAuth`, logout invocation, likely default-callback failure.
- `D:\work\qcash-ui-header-footer\hooks\useMenu.tsx`: auth-derived menu behavior.
- `D:\work\qcash-ui-header-footer\hooks\useLocalTaxMenu.tsx`: authorities consumption.
- `D:\work\qcash-ui-header-footer\components\component-global\index.tsx`: nested global-component remote loads.
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\MenuModalProvider.tsx`: global external-transfer modal consumer.
- `D:\work\qcash-ui-header-footer\layouts\default\container\menu\components\LcExportModal.tsx`: auth and response-modal consumer.
- `D:\work\qcash-ui-header-footer\node_modules\@ordentco\addons-auth-provider\src\auth.tsx`: matching `0.9.136-mfe` AuthContext/default callback implementation.
- `D:\work\qcash-ui\next.config.js`: shell federation shared configuration.
- `D:\work\qcash-ui\components\providers\BRIProviders.tsx`: BRI-versus-Ordent provider selection.
- `D:\work\qcash-ui\pages\_app.tsx`: shell provider nesting and federation bootstrap.
- `D:\work\qcash-ui\config.ts`: production feature/provider toggles.
- `D:\work\qcash-ui\constants\features\registry.ts`: global remote URLs and registration.
- `D:\work\qcash-ui\components\ui\SecureFeatureWrapper.tsx`: canonical header/footer loader and boundary.
- `D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx`: synchronous remote render error handling.
- `D:\work\qcash-ui\services\federation\utils\loader.ts`: retries, timeout, and load-error handling.
- `D:\work\qcash-ui\components\onboarding-tour\notif-onboarding-tour.tsx`: caught shell `guard()` invocation.
- `D:\work\qcash-ui\hooks\use-modal-session-expired.tsx`: conditional guard error handling.
- `D:\work\qcash-global-component\yarn.lock`: nearest-ref auth version evidence only.

---

## Turn 24: User

Continue if you have next steps, or stop and ask for clarification if you are unsure how to proceed.

---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The production shell selects the Ordent provider, not the BRI provider: `toggleAddonsAuthProviderMigration` is `false` in production. I’m now closing the remaining gaps around logout rejection handling, provider/render ordering, nested global-component sharing, and remote error boundaries.

---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The logout path is concrete: the header calls `logout(href)` inside a timer without `await`, `void`, or `.catch`, so a default-context logout becomes an unhandled rejected promise and bypasses the React error boundary. The production provider is rendered before all page/layout children; I’m checking whether every header callback default is safe and whether federation initialization can still create an isolated share scope.

---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 41: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 42: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 43: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 44: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 45: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The normal production composition resolves to one context: the shell’s Ordent provider loads before the header subtree, both exact refs request the same `0.9.136-mfe`, and the Next federation runtime applies `loaded-first` to their shared entries. The configuration still lacks a singleton guarantee, and I found a separate production defect: nested global-component loads originate from the header’s own federation instance, whose static remote points to the dev cluster.

---

## Turn 46: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 47: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 48: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 49: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 50: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 51: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 52: Assistant

**Verdict**
In the expected `qcash-ui` production composition, header/footer should receive the **same Ordent `AuthContext`** as the shell. A different context is possible in standalone or isolated-host scenarios because auth is not configured as a singleton.

The header’s default `logout()` does not synchronously crash React. It creates an **unhandled rejected promise** with `Function not implemented.` after the logout click.

**Key Findings**
1. Shell and header/footer use the identical auth artifact:
   - Header `refs/remotes/origin/v1.630.3-release-com:package.json:25`
   - Shell `refs/remotes/origin/v1.630.3-release-com:package.json:27`
   - Both resolve `0.9.136-mfe`, tarball hash `0520ac2...`.

2. Production selects the Ordent provider:
   - Toggle is `false` for production at shell `config.ts:472-477`.
   - Selection occurs at `pages/_app.tsx:99`.
   - The provider wraps all page/layout children at `pages/_app.tsx:109-135`.
   - Provider implementation is loaded at `components/providers/OrdentProviders.tsx:8-38`.
   - Even the BRI migration provider contains a nested Ordent provider at `components/providers/BRIProviders.tsx:42-45`.

3. Header/footer has no provider:
   - Its provider is commented out at header `pages/_app.tsx:62-66`.
   - The exposed layout directly renders `Layouts` at `layouts/default/index.tsx:9-10`.
   - Therefore it relies completely on the consuming host.

4. Module Federation preserves identity in the expected shell path:
   - Shell shares auth at `next.config.js:21`.
   - Header shares auth at `module-federation.config.mjs:17`.
   - Neither declares `singleton: true`.
   - Next Federation applies `loaded-first` to unspecified share strategies.
   - The shell provider consumes auth before the header subtree is rendered.
   - Remote container initialization attaches the shell share-scope object to the header runtime, so later header `useAuth()` calls select the already-loaded provider module.

5. The setup does not strictly guarantee identity:
   - Header standalone mode has no provider and therefore gets defaults.
   - Another host that does not share auth can load the header fallback copy.
   - A separately initialized share scope can produce another module instance.
   - Deployed assets differing from these refs can change version/load ordering.

**Default Failure Path**
The exact `0.9.136-mfe` artifact defines:

- `username: "Guest"`
- `userType: null`
- empty authorities and menus
- `isAuthoritiesReady: false`
- default async `logout()` throwing `Function not implemented.`

Evidence: `@ordentco/addons-auth-provider@0.9.136-mfe/src/auth.tsx:129-221`.

If header/footer receives that default context:

- Header remains in loading state through `layouts/default/container/index.tsx:21,64`.
- Menu fetching never starts because `userType` is empty at `hooks/useMenu.tsx:87-101`.
- Clicking logout reaches `ListItem.tsx:160`.
- After 500 ms, it calls `logout(href)` at `ListItem.tsx:60-73`.
- The returned rejected promise is neither awaited nor caught at line 70.
- `RemoteErrorBoundary` at shell `SecureFeatureWrapper.tsx:52-59` cannot catch timer/event-handler promise rejections.
- No `unhandledrejection` handler exists in either inspected ref.

Header/footer does not call default `guard()`. Its other callback, `setOnLeaveAction`, is called only when `onLeaveAction?.isLeave` is true at `MenuModalProvider.tsx:132-142`; both values are null in the default context, so that path is safe.

**Global Component**
The exact matching global-component release ref was unavailable. The nearest local ref, `origin/v1.626.5-release-com`, shows:

- Auth declared as `0.9.140-mfe` at `package.json:27`.
- `package-lock.json:8983-8985` also resolves `0.9.140-mfe`.
- `yarn.lock:3163-3165` is stale at `0.9.126-mfe`.
- Docker runs non-frozen `yarn install` at `Dockerfile:29-34`.
- Auth is shared as a singleton at `federation.config.js:8`.

The shell preloads global-component and header/footer remotes before rendering the auth provider at shell `services/federation/init.ts:74-86`. Consequently, a deployed global component newer than `0.9.136-mfe` can influence which auth version is selected. It should still be the **same context instance**, but potentially a different implementation version than the shell manifest suggests.

Header/footer also hardcodes the global-component dev remote at `module-federation.config.mjs:5-7`. The shell normally masks this by preloading the production global container first. If that preload fails, or header/footer runs standalone, nested loads from `components/component-global/index.tsx:5-14` can fall back to the dev-cluster URL.

**Recommended Fixes**
1. Align the auth version across shell, header/footer, and global-component, including all lockfiles.
2. Configure auth as `singleton: true` with an explicit compatible `requiredVersion` everywhere.
3. Remove the hardcoded dev global-component URL from the header production configuration.
4. Await and catch the header logout promise.
5. Change the auth library context default to `undefined` and make `useAuth()` throw a clear synchronous “missing AuthProvider” error.

No files, branches, or refs were modified. The header/footer worktree still contains only the two pre-existing unrelated modifications.

---

