# Inspect global release (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_00205ca98ffeXadIw7dECVsdAu`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 08:54:21`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only production investigation. In D:\work\qcash-global-component inspect exactly git ref refs/heads/v1.626.5-release-com (commit 5e51ff23), not unrelated working changes; do not edit/checkout. Trace its @ordentco auth version, Module Federation shared config, exposed auth-using components, local _app provider versus federated consumption, and any guard/default-callback calls. Check whether the package is singleton and whether load order can make its exposed component bind a separate context. Return exact ref:path:line references and failure scenarios relevant to qcash-ui shell + header/footer + mass-transfer.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

I’ll query Git objects directly with `git show`, `git grep`, and `git ls-tree` at `refs/heads/v1.626.5-release-com`, first verifying it resolves to `5e51ff23`. I will not checkout the ref or read the repository working tree.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

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

The ref resolves exactly to `5e51ff23aaee07b7ae429b50dfc4eea4593c3976`. Initial evidence shows npm metadata pins auth `0.9.140-mfe`, while `yarn.lock` still contains `0.9.126-mfe`; I’m now separating build-effective configuration from stale lockfile evidence and mapping each auth-consuming exposed module.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

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

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The shell evidence establishes the key asymmetry: at its inspected Git object, `qcash-ui` shares `@ordentco/addons-auth-provider` without `singleton`, while the global-component remote declares it singleton. I’m checking the sibling header/footer and mass-transfer Git refs next so the load-order scenarios are grounded in all three containers rather than inferred from the global remote alone.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

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

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The local `_app` providers are not part of an exposed module’s render path: global-component’s `_app` wraps only its own Next pages, and header/footer’s provider is commented out. The federated header/footer therefore depends entirely on the shell’s provider plus a compatible shared module identity; mass-transfer additionally calls `guard()` from whatever auth instance its exposed page resolves.

---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

One production-significant ambiguity remains in the target ref: `package.json`/`package-lock.json` say `0.9.140-mfe`, but the committed `yarn.lock` says `0.9.126-mfe`, and the build script invokes Yarn. I’m tracing the image/install path now because the effective bundled auth version depends on whether CI used npm, mutable Yarn resolution, or frozen Yarn.

---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **webfetch**: Tool: webfetch

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 41: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 42: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 43: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 44: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 45: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 46: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 47: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 48: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 49: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 50: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 51: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 52: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 53: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 54: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 55: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 56: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 57: Assistant

**Conclusion**
`refs/heads/v1.626.5-release-com` resolves exactly to `5e51ff23aaee07b7ae429b50dfc4eea4593c3976`.

- The target remote intends and clean-builds with `@ordentco/addons-auth-provider@0.9.140-mfe`.
- The target config marks auth `singleton: true`, but the wider shell graph is not consistently singleton.
- The target `_app` provider does not wrap federated exposures.
- In one correctly initialized default share scope, load order should change which auth version wins, not give the target’s singleton consumers a second context.
- A separate context remains possible through an isolated/nested runtime share scope or a non-singleton consumer resolving another factory. That would produce the exact mass-transfer/header failures described below.

**Version And Build**
- Exact dependency: `qcash-global-component@refs/heads/v1.626.5-release-com:package.json:27`
- npm lock confirms `0.9.140-mfe` and its tarball/integrity: `qcash-global-component@refs/heads/v1.626.5-release-com:package-lock.json:16`, `:8983-8992`
- Committed `yarn.lock` is stale at `0.9.126-mfe`: `qcash-global-component@refs/heads/v1.626.5-release-com:yarn.lock:3163-3171`
- The image copies only `package.json` and `yarn.lock`, then runs non-frozen `yarn install`: `qcash-global-component@refs/heads/v1.626.5-release-com:Dockerfile:28-34`
- It subsequently runs `yarn build`: `qcash-global-component@refs/heads/v1.626.5-release-com:Dockerfile:94-95`

Therefore, a successful clean Docker build resolves the exact manifest request `0.9.140-mfe` and rewrites the stale lock inside the build layer. The deployed image digest remains the final byte-level proof.

**Federation Configuration**
- Auth is explicitly `singleton: true, requiredVersion: false`: `qcash-global-component@refs/heads/v1.626.5-release-com:federation.config.js:5-8`
- React and React DOM are client-side singletons: `qcash-global-component@refs/heads/v1.626.5-release-com:next.config.js:47-58`
- The shared config is applied to both builds, while React additions are client-only: `qcash-global-component@refs/heads/v1.626.5-release-com:next.config.js:49-57`
- No `strictVersion`, concrete `requiredVersion`, `eager`, or custom `shareScope` is set for auth.

The package is therefore singleton for consumers compiled by this remote, but not intrinsically singleton across every container.

**Auth-Using Exposes**

| Exposed modules | Auth path |
|---|---|
| `./activity-log-new`, `./activity-log-v2`, `./activity-log` | Exposed at `federation.config.js:39-41`; new/v2 delegate through `activity-log-self.tsx:15,37` to `use-activity-log.ts:4,18`; legacy reads token at `activity-log.tsx:2,44` |
| `./confirm-modal` | Exposed at `federation.config.js:56`; reads auth/global store at `components/confirm-modal/index.tsx:1,63-64` |
| `./external-transfer-modal` | Exposed at `federation.config.js:65`; reads authorities at `external-transfer-modal.tsx:7,40` and company/menu/store at `use-external-transfer-methods.ts:3,28-29` |
| `./header` | Exposed at `federation.config.js:71`; avatar/list children read auth at `ListItem.tsx:3,14` and `avatar-dropdown.tsx:3,15` |
| `./layouts` | Exposed at `federation.config.js:80`; renders the auth-using header at `components/layouts/index.tsx:15` |
| `./menu` | Exposed at `federation.config.js:85`; reads leave state at `components/layouts/menu/components/Menu.tsx:1,32` and includes several authority-dependent modals |
| `./modal-single-action` | Exposed at `federation.config.js:92`; reads `bricamsUser` at `components/advanced-table/modal-single-action.tsx:18,86` |
| `./table` | Exposed at `federation.config.js:131`; imports/renders `BulkActionModal` at `components/table/table.tsx:4,52-81`, which reads auth/store at `bulk-action-modal.tsx:13,170-171` |
| `./transaction-modal` | Exposed at `federation.config.js:139`; reads auth/store at `components/transaction-modal/index.tsx:12,151,169` |

`./footer` is exposed at `federation.config.js:68` and is auth-free: `components/layouts/footer/footer.tsx:1-17`.

`PendingTaskList` uses auth and calls `canIApprove`, but is not exposed: `components/pending-task-list/pending-task-list.tsx:7,29,91-107`.

**Provider Boundary**
The target’s standalone Next app wraps local pages:

- Imports providers: `qcash-global-component@refs/heads/v1.626.5-release-com:pages/_app.tsx:6-9`
- Wraps its local component: `qcash-global-component@refs/heads/v1.626.5-release-com:pages/_app.tsx:49-53`

Federated consumers load modules listed in `federation.config.js`; they do not execute or render that `_app`. Thus the effective production tree is:

```text
qcash-ui shell AuthProvider / GlobalStoreProvider
  header-footer/default
    mass-transfer exposed page
      global-component exposed controls/modals
```

Cross-repo evidence:

- Shell provider dynamically imports both Ordent providers and wraps children: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:components/providers/OrdentProviders.tsx:8-21,28-38`
- Shell wraps the routed component with that provider: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:pages/_app.tsx:97,101-124`
- Even the BRI migration path nests an Ordent provider: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:components/providers/BRIProviders.tsx:21-28`
- Header/footer’s own providers are commented out: `qcash-ui-header-footer@refs/remotes/origin/v1.626.5-release-com:pages/_app.tsx:62-66`
- Mass-transfer’s `_app` provider is active only when mass-transfer runs standalone: `qcash-ui-mass-transfer@refs/remotes/origin/v1.629.0-release-com:pages/_app.tsx:62-64`

**Singleton Matrix**

| Container/ref | Installed auth | Shared policy |
|---|---:|---|
| Global component target | `0.9.140-mfe` | Singleton: `federation.config.js:8` |
| Shell COM anchor | `0.9.136-mfe` | Not singleton: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:next.config.js:21`; version at `package.json:24` |
| Header/footer COM anchor | `0.9.136-mfe` | Not singleton, lazy: `qcash-ui-header-footer@refs/remotes/origin/v1.626.5-release-com:module-federation.config.mjs:16-18`; version at `package.json:25` |
| Mass-transfer COM anchor | lock resolves `0.9.137` | Singleton: `qcash-ui-mass-transfer@refs/remotes/origin/v1.629.0-release-com:next.config.js:26-30`; lock at `yarn.lock:666-669` |

The shell pre-registers and loads `globals` from global-component and header/footer before rendering its provider:

- Global remotes: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:constants/features/registry.ts:11-21`
- Preload loop: `services/federation/init.ts:72-84`
- Actual `${moduleName}/globals` load: `services/federation/utils/loader.ts:19-37`

Under standard Webpack default-scope resolution, this makes `0.9.140-mfe` the likely highest registered version before the shell imports its non-singleton provider. Consequently, the global remote can effectively supply the shell’s provider implementation. That is a version takeover, but still one context.

**Guard And Default Calls**
There is no production `guard()` or literal `defaultCallback` call in the target global-component tree.

There are nevertheless fallback-callback hazards:

- Target header calls `logout("/main-page")`: `qcash-global-component@refs/heads/v1.626.5-release-com:components/layouts/header/components/ListItem.tsx:104`
- Header/footer calls `logout(href)`: `qcash-ui-header-footer@refs/remotes/origin/v1.626.5-release-com:layouts/default/container/header/components/ListItem.tsx:70`
- Every mass-transfer exposure calls zero-argument `guard()` in a mount-only effect:
  - Maker: `maker-confirmation/index.tsx:24`, `view/[batch].tsx:21`, `view/detail/[batch]/[cureff].tsx:31`
  - Transaction action: `transaction-action/index.tsx:21`, `view/[batch].tsx:22`, `view/detail/[batch]/[cureff].tsx:30`
  - Inquiry: `transfer-inquiry/index.tsx:20`, `view/[batch].tsx:21`, `view/detail/[batch]/[cureff].tsx:30`
  - Upload: `upload-transaction/index.tsx:20`
  - All are exposed at `qcash-ui-mass-transfer@refs/remotes/origin/v1.629.0-release-com:next.config.js:12-23`

The exact `0.9.140-mfe` package fallback context uses `username: "Guest"` and callbacks that throw `"Function not implemented."`:

- `@ordentco/addons-auth-provider@0.9.140-mfe:src/auth.tsx:130-223`
- Default `guard`: `:152-154`
- Default `logout`: `:161-163`
- Context creation: `:223`
- Real provider guard: `:298-342`
- Default global store is empty with a no-op dispatch: `src/global-store/index.tsx:26-29`, `src/global-store/initial-state.ts:3-33`

**Failure Scenarios**
1. **Mass-transfer permanent loader plus unhandled rejection.** If its `useAuth` resolves a different context object from the shell provider, it receives `Guest`, `isAuthoritiesReady=false`, and the fallback `guard`. The mount effect rejects with `Function not implemented`, while page checks keep returning the loading overlay. Example checks: `transaction-action/index.tsx:17-25` and `upload-transaction/index.tsx:15-24`.

2. **Header/footer shows Guest or empty menus.** Its layout reads username/company/user type/readiness at `layouts/default/container/index.tsx:8,20-21`, then passes those values to the header at `:64`. A separate context yields Guest/blank company/loading, while logout can invoke the throwing fallback.

3. **EFT V2 reports “Company ID Not Found.”** Header/footer dynamically consumes the target’s `external-transfer-modal`: `components/component-global/index.tsx:5-10`. The feature switch mounts it at `MenuModalProvider.tsx:164-168`. A separate target context supplies `companyID=null`; the target explicitly throws that error at `use-external-transfer-methods.ts:68-74` and catches it into modal error state at `:117-120`.

4. **Mass-transfer global modals use empty identity/authority state.** Mass-transfer renders target `ConfirmModal` at `containers/upload-transaction/index.tsx:700` and `TransactionModal` at `containers/transfer-inquiry/detail.tsx:656`. If the target context is separate, these receive null user type/company level and empty `bricamsUser`, altering authorization and confirmation behavior.

5. **Single context, unexpected auth behavior.** The nominal preload order can make target `0.9.140-mfe` supply the shell provider even though shell/header were built against `0.9.136-mfe`. This avoids context separation but makes the target release an implicit shell-wide auth upgrade.

6. **True context split indicates scope/runtime isolation, not version mismatch alone.** In one initialized default Webpack share scope, the target’s singleton consumer follows the selected/loaded singleton. A split points to a different share-scope object, failed/late container initialization, cached incompatible remoteEntry, or nested runtime isolation. Header/footer and mass-transfer bundle their own `@module-federation/runtime` and perform nested global-component loads, so that is the first production runtime condition to inspect.

Only the global-component ref was supplied as authoritative. The shell/header/mass references above are exact local COM anchors, not proof that those three image commits were co-deployed. No checkout or repository files were changed.

---

## Turn 58: Assistant

## Objective
- Perform a read-only production investigation of `D:\work\qcash-global-component` at exactly `refs/heads/v1.626.5-release-com` (`5e51ff23`).
- Trace `@ordentco/addons-auth-provider`, Module Federation sharing/exposes, provider boundaries, guard/default callbacks, and load-order failure scenarios involving `qcash-ui`, header/footer, and mass-transfer.

## Important Details
- Do not edit, checkout, or rely on unrelated working-tree changes; use Git-object commands such as `git show`, `git grep`, `git blame`, and `git ls-tree`.
- Target ref verified as `5e51ff23aaee07b7ae429b50dfc4eea4593c3976`.
- Final response must provide exact `ref:path:line` references and distinguish verified target facts from companion-repository evidence.
- Companion refs inspected:
  - Header/footer: `refs/remotes/origin/v1.626.5-release-com` → `ee409cab7d8193db987ae6b0f3f27322ccc4a6b5`.
  - Shell has no `v1.626.5` ref; inspected `refs/remotes/origin/v1.627.1-release-com` → `5ffb35ba29c109a92c3a2626ee65014018dea4f5`.
  - Mass-transfer has no `v1.626.x` ref; inspected `refs/remotes/origin/v1.629.0-release-com` → `265beed1770f860d9516e9257abd840bd9a3cc40`.
- Public artifact corroboration:
  - `https://unpkg.com/@ordentco/addons-auth-provider@0.9.140-mfe/src/auth.tsx`
  - Default context has `username: "Guest"` at line 135, throwing `guard` at 152-153, throwing `logout` at 161-162, `createContext(AUTH_INITIAL_VALUES)` at 223, and `useAuth` at 1188.
  - A consumer bound to an unprovided package instance therefore receives defaults; `guard()`/`logout()` throw `"Function not implemented."`.

## Work State
### Completed
- Verified target auth metadata conflict:
  - `refs/heads/v1.626.5-release-com:package-lock.json:16` pins `0.9.140-mfe`.
  - `refs/heads/v1.626.5-release-com:yarn.lock:3163-3166` contains only `0.9.126-mfe`.
  - `package.json` also declares `0.9.140-mfe`; exact line still needs a narrow capture.
- Verified target sharing:
  - `refs/heads/v1.626.5-release-com:federation.config.js:5-8` declares `@ordentco/addons-auth-provider` with `{ singleton: true, requiredVersion: false }`.
  - `next.config.js` imports and installs this federation config.
- Verified target local provider boundary:
  - `refs/heads/v1.626.5-release-com:pages/_app.tsx:6-9` imports `AuthProvider` and `GlobalStoreProvider`.
  - These wrappers apply to local Next pages, not directly to an exposed module rendered inside another host.
- Verified target auth consumers, including:
  - `components/activity-log/activity-log.tsx:2,44`
  - `components/activity-log/use-activity-log.ts:4,18`
  - `components/advanced-table/modal-single-action.tsx:18,86`
  - `components/confirm-modal/index.tsx:1,63-64`
  - `components/external-transfer-modal/external-transfer-modal.tsx:7,40`
  - `components/external-transfer-modal/use-external-transfer-methods.ts:28-29`
  - `components/table/bulk-action-modal.tsx:170-171`
  - `components/transaction-modal/index.tsx:151,169`
  - `components/layouts/menu/components/Menu.tsx:1`
- Verified target auth-method calls:
  - `components/layouts/header/components/ListItem.tsx:104` calls `logout("/main-page")`.
  - `components/layouts/menu/components/Menu.tsx:57` calls `setOnLeaveAction({})`.
  - `components/pending-task-list/pending-task-list.tsx:91,99,107` calls `canIApprove(...)`.
  - No production `guard()` or literal `defaultCallback` call was found in target source; matching “default callbacks” were tests/unrelated utility callbacks.
- Verified header/footer topology:
  - `module-federation.config.mjs:9-17` exposes `./default` and shares auth with `{ requiredVersion:false, eager:false }`, but **not** `singleton:true`.
  - `package.json:25` uses `0.9.136-mfe`.
  - `pages/_app.tsx:62-66` has its local auth provider commented out.
  - `layouts/default/index.tsx:9-10` routes exposed `./default` into `Layouts`.
  - `layouts/default/container/index.tsx:2-5` imports Header, Menu, Breadcrumb, and Footer.
  - `hooks/useMenu.tsx:2` imports `useAuth`/`useGlobalStore`.
  - `layouts/default/container/header/components/ListItem.tsx:70,160` calls `logout`.
  - `components/component-global/index.tsx:5-14` nested-loads target `external-transfer-modal` and `response-modal`.
- Verified shell topology at the stated companion ref:
  - `next.config.js:15+` shares auth without `singleton`; exact auth line still needs a narrow capture.
  - `package.json:24` uses `0.9.136-mfe`.
  - `pages/_app.tsx:97` selects `BriProviders` or `OrdentProviders` by feature toggle.
  - `pages/_app.tsx:102-108` places the selected auth provider beneath `ModuleFederationProvider`.
  - `components/providers/OrdentProviders.tsx:8-19` dynamically imports Ordent `AuthProvider` and `GlobalStoreProvider`.
  - `constants/features/registry.ts:11-20` registers global-component before header/footer.
  - `services/federation/init.ts:37-50` resolves all remotes concurrently before runtime `init`.
  - Mass-transfer shell pages independently load mass-transfer and header/footer, e.g. `pages/mass-transfer/transaction-action/index.tsx:13-21`.
- Verified mass-transfer companion behavior:
  - `next.config.js:12-28` exposes auth-using pages and marks auth singleton.
  - `package.json:18` requests `^0.9.129-mfe`; `yarn.lock:666-668` resolves `0.9.137`.
  - `pages/_app.tsx:6,62-64` has local providers, but those do not wrap pages when consumed as exposed modules.
  - All ten exposed page flows call `guard()`:
    - `maker-confirmation/index.tsx:24`
    - `maker-confirmation/view/[batch].tsx:21`
    - `maker-confirmation/view/detail/[batch]/[cureff].tsx:31`
    - `transaction-action/index.tsx:21`
    - `transaction-action/view/[batch].tsx:22`
    - `transaction-action/view/detail/[batch]/[cureff].tsx:30`
    - `transfer-inquiry/index.tsx:20`
    - `transfer-inquiry/view/[batch].tsx:21`
    - `transfer-inquiry/view/detail/[batch]/[cureff].tsx:30`
    - `upload-transaction/index.tsx:20`
- Established likely failure mechanism:
  - Target and mass-transfer compile auth as singleton consumers; shell and header/footer do not.
  - `requiredVersion:false` permits `0.9.136`, `0.9.137`, and `0.9.140` to coexist.
  - Webpack 5.91 singleton selection prefers an already-loaded singleton (`ConsumeSharedRuntimeModule.js:111-118,127-138`), while non-singleton loading selects a version independently (`103-110,215-239`).
  - Registration/consumption order can therefore make shell provider, header/footer, and a nested target component resolve different auth factories/React contexts.
  - Resulting symptoms include empty menus/authorities, `"Guest"`/null auth state, hidden transfer methods, no-op global-store dispatch, unauthorized/loading states, and mass-transfer `guard()` rejecting with `"Function not implemented."`.

### Active
- Finish mapping every target `federation.config.js` exposure to direct or transitive auth usage, especially `header`, `advanced-table`, `table`, `confirm-modal`, `external-transfer-modal`, `transaction-modal`, `pending-task-list`, and activity-log variants.
- Capture exact target `package.json`, `package-lock.json`, Docker install, `_app` wrapper, expose, and shell shared-config line ranges.
- Determine whether Docker/CI uses npm or Yarn so the effective target build version can be stated conclusively.
- Refine load-order scenarios without overstating deployment-ref assumptions.

### Blocked
- `rg` is unavailable: `"The term 'rg' is not recognized as a name of a cmdlet, function, script file, or executable program."`; use `git grep`, `findstr`, and narrow `git blame`.
- Several broad command outputs were truncated; rerun narrow line ranges.
- No exact matching shell or mass-transfer Git ref was found, so production deployment mapping remains unknown and must be disclosed.

## Next Move
1. Run narrow read-only inspections for `federation.config.js`, `package.json`, `package-lock.json`, `Dockerfile`, `pages/_app.tsx`, and each exposed module to complete exact line mappings and effective-version evidence.
2. Produce the final concise finding: target-level singleton versus topology-wide non-singleton, provider-boundary explanation, verified guard/default behavior, and ordered shell/header/footer/mass-transfer failure scenarios with ref-qualified citations.

## Relevant Files
- `D:\work\qcash-global-component\federation.config.js`: target shared configuration and exposes.
- `D:\work\qcash-global-component\next.config.js`: installs `NextFederationPlugin`.
- `D:\work\qcash-global-component\package.json`: declared auth version.
- `D:\work\qcash-global-component\package-lock.json`: npm-resolved `0.9.140-mfe`.
- `D:\work\qcash-global-component\yarn.lock`: stale `0.9.126-mfe`.
- `D:\work\qcash-global-component\Dockerfile`: determines effective package manager/version.
- `D:\work\qcash-global-component\pages\_app.tsx`: local-only auth providers.
- `D:\work\qcash-global-component\components\layouts\`: exposed layout/header/menu auth usage.
- `D:\work\qcash-global-component\components\external-transfer-modal\`: nested exposed auth consumer.
- `D:\work\qcash-global-component\components\advanced-table\modal-single-action.tsx`: global-store consumer.
- `D:\work\qcash-global-component\components\confirm-modal\index.tsx`: auth/global-store consumer.
- `D:\work\qcash-global-component\components\table\bulk-action-modal.tsx`: auth/global-store consumer.
- `D:\work\qcash-global-component\components\transaction-modal\index.tsx`: auth/global-store consumer.
- `D:\work\qcash-ui\next.config.js`: shell non-singleton shared configuration.
- `D:\work\qcash-ui\pages\_app.tsx`: provider selection and federation/provider order.
- `D:\work\qcash-ui\components\providers\OrdentProviders.tsx`: shell Ordent provider factory.
- `D:\work\qcash-ui\constants\features\registry.ts`: global/header remote ordering.
- `D:\work\qcash-ui\services\federation\init.ts`: concurrent remote registration.
- `D:\work\qcash-ui-header-footer\module-federation.config.mjs`: exposed default and non-singleton auth share.
- `D:\work\qcash-ui-header-footer\pages\_app.tsx`: commented local provider.
- `D:\work\qcash-ui-header-footer\layouts\default\`: exposed auth-consuming layout.
- `D:\work\qcash-ui-mass-transfer\next.config.js`: singleton share and exposed pages.
- `D:\work\qcash-ui-mass-transfer\pages\_app.tsx`: local provider not used during federated consumption.
- `D:\work\qcash-ui-mass-transfer\pages\mass-transfer\`: exposed `guard()` call sites.

---

## Turn 59: User

Continue if you have next steps, or stop and ask for clarification if you are unsure how to proceed.

---

## Turn 60: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

I’m tightening the remaining uncertainty around runtime initialization and share-scope identity. I’ll verify whether the nested loaders reuse Webpack’s default scope or can create isolated scopes, then separate proven failure paths from deployment-dependent hypotheses.

---

## Turn 61: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **context7_query-docs**: Tool: context7_query-docs
- **context7_query-docs**: Tool: context7_query-docs
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 62: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The cross-repo loaders do not all use the same runtime package: the shell resolves `0.18.1`, header/footer `0.2.8`, and mass-transfer `0.1.21`, while their Next federation plugin graph also contains `0.0.17`. I’m checking those exact runtime releases now, because whether `loadRemote()` delegates to one global instance or a package-local instance decides whether context isolation is a concrete path rather than speculation.

---

## Turn 63: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 64: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 65: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 66: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The exact runtime source changes the diagnosis: each package release keeps a module-local `FederationInstance`, and `loadRemote()` throws unless that same package copy has executed `init()`. Shell initialization of `0.18.1` does not initialize header/footer’s `0.2.8` variable. I’m tracing which exposed entrypoints call `init()` themselves; this may be a more direct production failure than auth-context negotiation.

---

## Turn 67: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 68: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 69: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 70: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The static path is now concrete: header/footer and mass-transfer initialize their own runtime only in local `_app`, which is bypassed when their modules are exposed. Their nested `loadRemote()` calls use different package copies and have no matching `init()` in the exposed path. I’m checking committed production endpoint/config metadata next so this can be verified against the deployed remote entries rather than left as a source-only conclusion.

---

## Turn 71: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 72: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 73: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **context7_query-docs**: Tool: context7_query-docs
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

One source-level assumption needs validation before I call this a defect: header/footer renders a nested global `ResponseModal` on the normal layout path, so an actually uninitialized runtime would fail immediately and would likely have been obvious. I’m checking `NextFederationPlugin@8.2.5` for automatic runtime sharing/bootstrap behavior that may initialize or substitute the remote’s direct runtime import.

---

## Turn 74: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 75: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>

The plugin source resolves the contradiction: `NextFederationPlugin@8.2.5` force-aliases every exact `@module-federation/runtime` import to its own nested runtime (`0.0.17`). The apparent shell/header/mass runtime versions are not the browser implementation for these imports, so the “uninitialized 0.2.8/0.1.21 copy” path is not valid. I’m now confirming all selected refs use that plugin release and tracing its container-init bridge to establish whether nested remotes receive the host’s same scope object.

---

## Turn 76: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 77: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 78: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 79: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 80: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 81: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 82: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>

The normal auth topology is now much stronger than the earlier hypothesis: container initialization explicitly rebinds each remote’s `default` share scope to the host’s exact scope object, and the runtime plugin forces `loaded-first`. I’m comparing auth `0.9.136`, `0.9.137`, and `0.9.140` next, because the credible residual risk is version takeover or fallback to an older implementation, not duplicate React context in the standard shell path.

---

## Turn 83: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 84: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 85: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 86: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 87: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 88: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 89: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 90: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 91: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 92: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 93: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 94: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 95: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans

</details>

Static resolution is complete enough to rule out the normal duplicate-context theory. The remaining deployment-level question is whether production actually built and paired these refs; I’m checking Bamboo metadata for the target branch before asking for a production URL or image digest.

---

## Turn 96: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans

</details>



---

## Turn 97: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bamboo-build-doctor_bamboo_diagnose_build**: Tool: bamboo-build-doctor_bamboo_diagnose_build
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds

</details>



---

## Turn 98: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_diagnose_build**: Tool: bamboo-build-doctor_bamboo_diagnose_build
- **bamboo-build-doctor_bamboo_diagnose_build**: Tool: bamboo-build-doctor_bamboo_diagnose_build

</details>



---

## Turn 99: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans
- **bamboo-build-doctor_bamboo_search_plans**: Tool: bamboo-build-doctor_bamboo_search_plans

</details>



---

## Turn 100: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 101: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 102: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 103: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 104: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 105: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

A more direct provider-boundary issue surfaced while validating load order: `_app` statically imports and calls `useSessionExpiration()` in `MyApp` before returning the `<AuthProvider>` that wraps only its children. If that hook calls Ordent `useAuth`, it is outside the provider by React tree position and will always receive the default context even when sharing is correct. I’m tracing that path now; it may explain the reported default `guard()` behavior without any duplicate context.

---

## Turn 106: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The apparent outer-hook issue is not present: `useSessionExpiration()` itself does not consume auth, and the modal hook that does consume auth is rendered inside the provider. However, `_app` statically imports `NotifOnboardingTour`, which calls `useAuth`; that can synchronously consume shell `0.9.136` during module evaluation before the top-level federation preload. I’m validating those initial imports because they may pin `0.9.136`, reversing which version actually wins while still preserving one context.

---

## Turn 107: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 108: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>

The effective version is now determinable: the shell’s static onboarding imports consume auth before `_app` reaches federation initialization, so `loaded-first` pins shell `0.9.136`; target `0.9.140` and mass `0.9.137` then reuse that same factory. I’m checking the real `0.9.136` guard error path now, because mass-transfer fires it without awaiting or catching it and gates every page on `isAuthoritiesReady`.

---

## Turn 109: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 110: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds

</details>

A closer deployment-time shell ref exists: `qcash-ui` has `v1.628.8-release-com` at `b332cb12`, committed August 4, one day after the exact global/header images. I’m switching the companion shell anchor to that ref and checking only the relevant auth/federation files; this removes the earlier July-ref ambiguity.

---

## Turn 111: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 112: Assistant

## Objective
- Complete a read-only production root-cause investigation of `D:\work\qcash-global-component` at exactly `refs/heads/v1.626.5-release-com` (`5e51ff23aaee07b7ae429b50dfc4eea4593c3976`).
- Produce a corrected final report on auth version selection, Module Federation scope identity, provider boundaries, mass-transfer failures, and deployment evidence.

## Important Details
- Do not edit, checkout, or trust unrelated working-tree state; use `git show`, `git grep`, `git blame`, and `git ls-tree`.
- Only the global-component ref is authoritative. Companion refs are evidence, not proof of co-deployment:
  - Header/footer: `refs/remotes/origin/v1.626.5-release-com` → `ee409cab7d8193db987ae6b0f3f27322ccc4a6b5`.
  - Mass-transfer: `refs/remotes/origin/v1.629.0-release-com` → `265beed1770f860d9516e9257abd840bd9a3cc40`.
  - Initial shell anchor: `refs/remotes/origin/v1.627.1-release-com` → `5ffb35ba29c109a92c3a2626ee65014018dea4f5`.
  - Closer shell anchor now under review: `refs/remotes/origin/v1.628.8-release-com` → `b332cb126b875ff785e42375372855cb97f09d45`.
- Corrected conclusion: ordinary version mismatch does **not** explain duplicate auth contexts in the inspected standard shell path. The enhanced runtime attaches remotes to the host’s exact default share-scope object and uses `loaded-first`.
- Superseded claim: target `0.9.140-mfe` is not expected to take over the inspected shell. Static shell imports appear to load `0.9.136-mfe` before federation preload, pinning it for later singleton and non-singleton consumers.
- A true default context still indicates a missing provider, isolated/different share scope, failed/late initialization, stale remoteEntry, or nested runtime isolation.
- No repository files were changed.

## Work State
### Completed
- Verified target dependency/build metadata:
  - `qcash-global-component@refs/heads/v1.626.5-release-com:package.json:27` declares `0.9.140-mfe`.
  - `package-lock.json:16,8983-8992` resolves `0.9.140-mfe`.
  - `yarn.lock:3163-3171` is stale at `0.9.126-mfe`.
  - `Dockerfile:28-34,94-95` copies `package.json`/`yarn.lock`, runs non-frozen `yarn install`, then `yarn build`; a clean successful build resolves the manifest’s `0.9.140-mfe`.
- Verified target federation policy:
  - `federation.config.js:5-8` marks auth `{ singleton: true, requiredVersion: false }`.
  - `next.config.js:47-58` adds client React/React DOM singletons.
  - No concrete `requiredVersion`, `strictVersion`, `eager`, or custom `shareScope`.
  - Commit `847e112bb9bebdde5a8c706891e096fc39bf72f2` changed target auth from `singleton:false` to `singleton:true`.
- Verified target provider boundary:
  - `pages/_app.tsx:6-9,49-53` wraps standalone pages.
  - Federated exposes do not execute that `_app`; they depend on the shell provider.
- Completed target exposure/auth mapping:
  - Activity logs: `federation.config.js:39-41`; `activity-log-self.tsx:15,37`, `use-activity-log.ts:4,18`, `activity-log.tsx:2,44`.
  - Confirm modal: expose `:56`; `components/confirm-modal/index.tsx:1,63-64`.
  - External transfer: expose `:65`; `external-transfer-modal.tsx:7,40`, `use-external-transfer-methods.ts:3,28-29`.
  - Header/layout/menu: exposes `:71,80,85`; auth usage in `ListItem.tsx:3,14,104`, `avatar-dropdown.tsx:3,15`, and `Menu.tsx:1,32,57`.
  - Modal-single/table/transaction-modal: exposes `:92,131,139`; auth usage at `modal-single-action.tsx:18,86`, `bulk-action-modal.tsx:13,170-171`, and `transaction-modal/index.tsx:12,151,169`.
  - `./footer` at `:68` is auth-free.
  - `PendingTaskList` uses auth and `canIApprove` but is not exposed.
- Verified no production target `guard()` or literal `defaultCallback` call. Target header does call `logout("/main-page")` at `components/layouts/header/components/ListItem.tsx:104`.
- Verified companion provider/sharing topology at the initial shell anchor:
  - Shell auth `0.9.136-mfe`, shared non-singleton: `qcash-ui@refs/remotes/origin/v1.627.1-release-com:package.json:24`, `next.config.js:21`.
  - Header auth `0.9.136-mfe`, non-singleton/lazy: `qcash-ui-header-footer@refs/remotes/origin/v1.626.5-release-com:package.json:25`, `module-federation.config.mjs:16-18`.
  - Mass lock resolves `0.9.137`, shared singleton: `qcash-ui-mass-transfer@refs/remotes/origin/v1.629.0-release-com:yarn.lock:666-669`, `next.config.js:26-30`.
  - Shell provider wraps routed content: `components/providers/OrdentProviders.tsx:8-21,28-38`, `pages/_app.tsx:97,101-124`.
  - Header provider is commented out: `pages/_app.tsx:62-66`.
  - Mass provider applies only standalone: `pages/_app.tsx:62-64`.
- Verified every mass-transfer exposure calls zero-argument `guard()` in a mount-only effect; expose list is `next.config.js:12-23`. Representative gates are `transaction-action/index.tsx:17-25` and `upload-transaction/index.tsx:15-24`.
- Corrected the runtime-version false lead:
  - Direct declarations resolve shell runtime `0.18.1`, header `0.2.8`, and mass `0.1.21`, but these are not the exact browser implementation.
  - Shell uses `@module-federation/nextjs-mf@8.2.6`; target/header/mass use `8.2.5`.
  - Both plugin releases alias exact `@module-federation/runtime$` imports to their nested runtime `0.0.17`: `NextFederationPlugin/index.js:54,58`.
  - Therefore the proposed “header’s uninitialized `0.2.8` or mass’s uninitialized `0.1.21` module-local instance” defect is invalid.
- Verified enhanced runtime scope behavior:
  - Next internal runtime plugin sets shared strategy to `loaded-first`: `runtimePlugin.js:39`.
  - Bundler runtime attaches `webpackRequire.S` to `federation.instance.shareScopeMap`: `webpack-bundler-runtime@0.0.17/dist/index.esm.js:5-15`.
  - Container init calls `initShareScopeMap(name, shareScope)` and reattaches it: `:264-276`.
  - Loaded-first prefers an already loaded version: `runtime@0.0.17/dist/share.esm.js:727-780`.
- Established expected version winner for inspected shell `v1.627.1`:
  - `_app` statically imports auth-consuming onboarding modules before top-level federation initialization: `pages/_app.tsx:20-21,37-48`.
  - `NotifOnboardingTour` imports `useAuth`: `components/onboarding-tour/notif-onboarding-tour.tsx:4`.
  - `OnboardingTourProvider` imports `useAuth`: `components/providers/OnboardingTourProvider.tsx:3`.
  - The actual provider is dynamically imported later: `components/providers/OrdentProviders.tsx:8-20`.
  - Thus shell `0.9.136-mfe` should be loaded first and reused by target `0.9.140` and mass `0.9.137` consumers in the normal default scope.
- Resolved an apparent outer-provider issue:
  - `useSessionExpiration()` itself does not consume auth.
  - `useModalSessionExpired()` does consume auth, but its component is rendered inside the provider.
- Compared auth package behavior:
  - `0.9.140-mfe` default context has `username:"Guest"` and throwing `guard`/`logout`: `src/auth.tsx:130-223`.
  - `0.9.136`/`0.9.137` use default login/base routes `/main-page` and `/`; `0.9.140` defaults to `/landing-page` and `/homepage`.
  - `0.9.136` always permits session-freshness caching; `0.9.140` only uses it when `guard(true)` is called. Mass calls `guard()` without arguments.
- Found a concrete same-context mass-transfer failure:
  - Real `0.9.136` guard sets `isAuthoritiesReady(false)` at `src/auth.tsx:313`.
  - `validateToken()` is caught at `:317-336`, but `validateMenu(token)` at `:337` is outside that catch.
  - If menu validation rejects, mass pages neither await nor catch `guard()`, and readiness remains false, producing an unhandled rejection plus permanent loading overlay without any duplicate context.
- Verified Bamboo branch builds:
  - Target: `BABMF-QCASHGLOBALCOMPONENT450-1`, successful, 2026-08-03.
  - Header/footer: `BABMF-QCASHUIHEADERFOOTER238-1`, successful, 2026-08-03.
  - Mass-transfer: `BABMF-QCASHUIMASSTRANSFER34-2`, successful, 2026-08-06.
  - Shell `v1.628.8-release-com` exists as Bamboo branch `BABMF-QCASHUI557`; latest visible result is `BABMF-QCASHUI557-5`.

### Active
- Narrowly compare `qcash-ui@refs/remotes/origin/v1.628.8-release-com` against `v1.627.1` for:
  - `package.json`, `yarn.lock`, `next.config.js`.
  - `pages/_app.tsx`, onboarding static imports, `OrdentProviders.tsx`.
  - Federation initialization, registry, and preload behavior.
- Inspect `BABMF-QCASHUI557-5` state and `vcsRevisionKey`; correlate all successful builds with exact Git commits.
- Finish checking the full `0.9.136` guard tail and corresponding `0.9.140` error handling.
- Verify whether nested header/mass `loadRemote("qcash-global-component/...")` uses the shell-registered remote or the compiled hardcoded `gym-master` entry.
- Prepare the corrected final report and a production runtime diagnostic for share-scope identity, loaded auth version, provider presence, and remote-entry selection.

### Blocked
- No production URL, image digest, deployment manifest, or pod image ID is available; Bamboo proves builds, not co-deployment.
- Successful Bamboo diagnosis returned `logs: []`, so image tags/digests are not yet visible.
- `rg` is unavailable: `"The term 'rg' is not recognized as a name of a cmdlet, function, script file, or executable program."`
- Broad Git/Bamboo outputs are truncated; use narrow path/range queries.

## Next Move
1. Inspect the narrow `v1.628.8` shell files and diagnose `BABMF-QCASHUI557-5` to lock the companion shell commit and effective auth winner.
2. Finalize the corrected findings; request only a production URL/image digest if deployment-level proof is required.

## Relevant Files
- `D:\work\qcash-global-component\package.json`, `package-lock.json`, `yarn.lock`, `Dockerfile`: declared/resolved/build auth version.
- `D:\work\qcash-global-component\federation.config.js`, `next.config.js`: target sharing and exposes.
- `D:\work\qcash-global-component\pages\_app.tsx`: standalone-only providers.
- `D:\work\qcash-global-component\components\layouts`, `components\external-transfer-modal`, `components\confirm-modal`, `components\table`, `components\transaction-modal`: exposed auth consumers.
- `D:\work\qcash-ui\pages\_app.tsx`: static imports, top-level federation initialization, and provider ordering.
- `D:\work\qcash-ui\components\providers\OrdentProviders.tsx`: dynamically loaded shell provider.
- `D:\work\qcash-ui\components\onboarding-tour\notif-onboarding-tour.tsx`: early static auth consumer.
- `D:\work\qcash-ui\components\providers\OnboardingTourProvider.tsx`: early static auth consumer.
- `D:\work\qcash-ui\services\federation\init.ts`, `services\federation\utils\loader.ts`, `constants\features\registry.ts`: runtime registration and preload.
- `D:\work\qcash-ui-header-footer\module-federation.config.mjs`, `pages\_app.tsx`, `components\component-global\index.tsx`: sharing, bypassed provider, nested remote loads.
- `D:\work\qcash-ui-mass-transfer\next.config.js`, `pages\_app.tsx`, `pages\mass-transfer`, `components\global.ts`: exposes, guard calls, and nested remote loads.
- `https://unpkg.com/@ordentco/addons-auth-provider@0.9.136-mfe/src/auth.tsx`: expected shell auth implementation.
- `https://unpkg.com/@ordentco/addons-auth-provider@0.9.140-mfe/src/auth.tsx`: target implementation/default context.
- `https://unpkg.com/@module-federation/nextjs-mf@8.2.5/dist/src/plugins/NextFederationPlugin/index.js`: runtime alias.
- `https://unpkg.com/@module-federation/nextjs-mf@8.2.5/dist/src/plugins/container/runtimePlugin.js`: `loaded-first` strategy.
- `https://unpkg.com/@module-federation/webpack-bundler-runtime@0.0.17/dist/index.esm.js`: share-scope attachment and container initialization.

---

## Turn 113: User

Continue if you have next steps, or stop and ask for clarification if you are unsure how to proceed.

---

## Turn 114: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_diagnose_build**: Tool: bamboo-build-doctor_bamboo_diagnose_build

</details>

The closer `v1.628.8` shell keeps the relevant federation configuration and early auth-import ordering unchanged; its visible changes are unrelated registry, feature-toggle, CSP, and test tooling updates. I’m now locking the exact auth resolution and Bamboo revision, then validating the guard failure path before issuing the corrected conclusion.

---

## Turn 115: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **bamboo-build-doctor_bamboo_get_builds**: Tool: bamboo-build-doctor_bamboo_get_builds
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>

The exact shell build is now confirmed: `BABMF-QCASHUI557-5` successfully built commit `b332cb126b875ff785e42375372855cb97f09d45`, and that commit still pins Ordent auth `0.9.136-mfe` with the same early consumers. One remaining high-impact check is the production migration toggle, because `_app` can choose a different `@bri` provider while those consumers and remotes still import the Ordent context.

---

## Turn 116: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>

Production on the confirmed shell commit selects `OrdentProviders` (`toggleAddonsAuthProviderMigration` is `false` for prod), so the mixed `@bri` provider path is ruled out. The build-to-commit chain is now exact for shell, target, header, and mass-transfer; I’m narrowing the remaining uncertainty to nested remote resolution and separating the default-context symptom from the independently reproducible guard hang.

---

## Turn 117: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 118: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 119: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 120: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 121: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 122: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 123: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 124: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 125: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 126: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 127: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 128: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>

The remaining runtime trace is conclusive: exposed header and mass-transfer code executes against each remote’s own auto-initialized federation instance. Container initialization shares the shell’s auth scope, but it does not replace that instance’s `remotes`; runtime `formatOptions` merges and preserves them. Consequently, nested global-component loads retain the committed `http://gym-master...` development endpoint, which is a separate production defect even though it does not by itself create a second auth context.

---

## Turn 129: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 130: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 131: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 132: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 133: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 134: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_tabs**: Tool: playwright_browser_tabs

</details>



---

## Turn 135: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 136: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 137: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_network_requests**: Tool: playwright_browser_network_requests

</details>



---

## Turn 138: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_network_request**: Tool: playwright_browser_network_request

</details>



---

## Turn 139: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_tabs**: Tool: playwright_browser_tabs

</details>



---

## Turn 140: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_tabs**: Tool: playwright_browser_tabs

</details>



---

## Turn 141: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_tabs**: Tool: playwright_browser_tabs

</details>



---

## Turn 142: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 143: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 144: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 145: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 146: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 147: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 148: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 149: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The build timestamps reveal a likely coordinated release batch: target, header, and shell build `BABMF-QCASHUI557-3` all started within three minutes on Aug 3 and were run by the same user. I’m checking that shell revision (`d5e9e7…`) as well as the branch-tip build, because it is the stronger temporal deployment candidate.

---

## Turn 150: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 151: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 152: Assistant

**Final Findings**
1. **Permanent mass-transfer loading has a concrete same-context cause.** The shell supplies auth `0.9.136-mfe`. Its `guard()` sets `isAuthoritiesReady(false)` at `auth.tsx:313`, but only `validateToken()` is inside the `try/catch`. `validateMenu(token)` at `:337` can reject, while readiness is restored only at `:475`.
   
   Mass pages call `guard()` without awaiting or catching it, then gate rendering on `isAuthoritiesReady`, for example `transaction-action/index.tsx:20-25` and `maker-confirmation/index.tsx:23-28`. A menu API failure therefore produces an unhandled rejection and a permanent loader.

2. **The guard cache is effectively stale.** `setSessionLastValidatedAt()` is called at `auth.tsx:477`, but the callback dependencies at `:478` are only `[token, authService]`. The callback normally retains the old `null` timestamp, causing repeated token/menu validation across mass-transfer routes. Versions `0.9.137` and `0.9.140` retain the same dependency and error-handling problem.

3. **Header and mass-transfer contain a confirmed production remote-resolution defect.**
   - Header: `module-federation.config.mjs:5-7`
   - Mass-transfer: `next.config.js:8-10`

   Both embed:

   ```text
   http://internal-service.example.com/global-component/_next/static/chunks/remoteEntry.js
   ```

   Their exposed modules use their own auto-initialized federation instances. Container initialization adopts the shell’s share scope but preserves each remote’s existing `remotes` configuration. Consequently, nested `loadRemote("qcash-global-component/...")` calls retain this development endpoint; their `_app` runtime initialization is bypassed when federated.

   The endpoint currently returns `200`, serves a global component with auth `0.9.140-mfe`, and was last modified August 13. This means production can load mutable development code. From an HTTPS shell, the HTTP script can also be blocked as active mixed content.

4. **Normal auth version skew is not the expected source of a second React context.** Both candidate shell builds statically import auth consumers before federation initialization at `pages/_app.tsx:20-21,37-48`. The shell’s `0.9.136-mfe` factory is therefore loaded first. Enhanced runtime `loaded-first` selection and container share-scope rebinding should make target `0.9.140` and mass `0.9.137` consumers reuse that factory.

   Production also selects `OrdentProviders`: `config.ts:472-477` sets the migration toggle to `false` for prod.

5. **`Function not implemented.` remains a distinct diagnostic.** That literal exists only in the auth default context, such as `0.9.136-mfe:auth.tsx:150-152`. If observed, the component did not receive the real Ordent provider context. Given the standard topology above, inspect actual deployment divergence, failed container initialization, stale remote entries, or rendering outside the provider rather than assuming ordinary version mismatch.

**Build Evidence**
- Target `5e51ff23...`: successful `BABMF-QCASHGLOBALCOMPONENT450-1`.
- Header `ee409cab...`: successful `BABMF-QCASHUIHEADERFOOTER238-1`.
- Shell `d5e9e7de...`: successful `BABMF-QCASHUI557-3`; it started within three minutes of target/header under the same operator, making it the strongest coordinated-release candidate.
- Shell branch tip `b332cb12...`: successful `BABMF-QCASHUI557-5`.
- Relevant auth/federation files are unchanged between those two shell revisions.
- Mass-transfer `265beed1...`: successful `BABMF-QCASHUIMASSTRANSFER34-2`.

**Fix Order**
1. Wrap the entire auth validation flow, including `validateMenu` and response processing, in explicit error handling. Expose an error state instead of leaving `isAuthoritiesReady=false`.
2. Centralize mass-transfer guarding and handle the promise with an error/retry UI.
3. Fix guard freshness using a ref or complete dependencies.
4. Remove the hardcoded development remote. Runtime registration must execute for federated exposes, not only each app’s `_app`.
5. Align auth singleton declarations and versions as hardening, although this is not the primary failure found.

**Runtime Check**
Run this in the production console to verify the remaining deployment-dependent facts:

```js
(() => {
  const instances = globalThis.__FEDERATION__?.__INSTANCES__ ?? [];
  const shell = instances.find((x) => x.options?.name === "qcash-ui");
  const authKey = "@ordentco/addons-auth-provider";

  return instances.map((instance) => {
    const scope = instance.shareScopeMap?.default;
    const auth = scope?.[authKey] ?? {};

    return {
      name: instance.options?.name,
      sameScopeAsShell: scope === shell?.shareScopeMap?.default,
      remotes: (instance.options?.remotes ?? []).map(
        ({ name, entry }) => `${name}=${entry}`
      ),
      auth: Object.fromEntries(
        Object.entries(auth).map(([version, value]) => [
          version,
          {
            from: value.from,
            loaded: value.loaded,
            useIn: value.useIn,
          },
        ])
      ),
    };
  });
})();
```

Expected normal result: all loaded remotes report the same shell scope and auth `0.9.136-mfe` is loaded. Any `gym-master` entry confirms the nested production misconfiguration.

No files were changed. The only unresolved deployment fact is the actual production image digest or remote URL currently serving each application.

---

