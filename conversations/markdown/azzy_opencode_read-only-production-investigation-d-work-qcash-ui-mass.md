# Inspect mass transfer release (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_00205ca37ffeKVHzjv2VriObH2`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 08:54:21`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read-only production investigation. In D:\work\qcash-ui-mass-transfer inspect exactly git ref refs/remotes/origin/v1.629.0-release-com (commit 265beed1), do not edit/checkout. Trace auth dependency version, _app provider, every useAuth/guard call, qcash-global-component consumption, Module Federation shared settings/remotes/exposes, runtime loading, and whether this MFE is rendered standalone or as a remote beneath qcash-ui. Identify exactly where an unmatched AuthContext can call the default `Function not implemented.` guard and whether calls catch rejected promises. Return exact ref:path:line references and a concrete failure path.

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

I’ll inspect the release tree exclusively through read-only `git show`, `git grep`, and metadata commands against commit `265beed1`, without changing either repository’s worktree. I’ll then correlate its federation contract with the host `qcash-ui` integration and report only references tied to the inspected ref (clearly distinguishing any host-side ref).

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

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The release tree exposes route components directly, while its own `_app` is the only place that mounts `AuthProvider`. That makes the host’s federation share scope decisive: when an exposed page is imported by `qcash-ui`, the remote `_app` is not part of that render path, so I’m now verifying the exact resolved auth package and the matching host release configuration.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The lockfile resolves the remote’s declared `^0.9.129-mfe` to stable `0.9.137`, not to `0.9.129-mfe`. The available host release refs use both a newer `@bri` auth package and an older pinned `@ordentco` package, so I’m narrowing this to the actual provider nesting and mass-transfer route wrappers rather than treating version difference alone as proof of the context mismatch.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

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
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **webfetch**: Tool: webfetch

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

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_close**: Tool: playwright_browser_close

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 19: Assistant

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

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_navigate**: Tool: playwright_browser_navigate

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_evaluate**: Tool: playwright_browser_evaluate

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **playwright_browser_close**: Tool: playwright_browser_close

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

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

The host release confirms the production branch chooses the legacy `0.9.136-mfe` provider, while the mass-transfer build contributes `0.9.137`; the host does not mark that package singleton, but the remote does. I’m checking load order now, because whether the remote is registered/preloaded before the host provider decides if this is merely a configuration risk or a directly reachable split-context path.

---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 30: Assistant

**Conclusion**
The source contains a concrete split-`AuthContext` failure path. It is not guaranteed on every load because Module Federation may select one shared auth instance, but the release does not guarantee identity:

- Mass transfer declares `^0.9.129-mfe` but locks stable `0.9.137`: `265beed1:package.json:18`, `265beed1:yarn.lock:666-669`.
- The examined `qcash-ui` COM host locks `0.9.136-mfe`: `9b7e1d39:package.json:27`, `9b7e1d39:yarn.lock:696-699`.
- Mass transfer marks auth `singleton: true`; the host does not: `265beed1:next.config.js:26-30`, `9b7e1d39:next.config.js:15-21`.
- When rendered by the host, mass transfer’s `_app` and its `AuthProvider` are bypassed. The exposed page must therefore consume the exact same auth module instance as the host provider.
- All ten exposed pages call `guard()` without error handling.

`refs/remotes/origin/v1.629.0-release-com` resolves exactly to `265beed1770f860d9516e9257abd840bd9a3cc40`. For host evidence, `qcash-ui` has no corresponding `v1.629.0-release-com` ref; the nearest COM host ref containing all routes is `refs/remotes/origin/v1.628.9-release-com`, commit `9b7e1d39791991d654b4579c20e13f10d2265edc`.

**Auth Provider**
The remote’s standalone `_app` imports and mounts both providers:

- Import: `265beed1:pages/_app.tsx:6`
- `AuthProvider`: `265beed1:pages/_app.tsx:62-64`
- `GlobalStoreProvider`: `265beed1:pages/_app.tsx:63`
- API URL source: `265beed1:config.ts:5-15`

The exact error text is not present in commit `265beed1`; it comes from the locked npm artifact. In `@ordentco/addons-auth-provider@0.9.137`:

- Default `guard` throws synchronously: `npm:@ordentco/addons-auth-provider@0.9.137:dist/src/auth.js:118-120`
- Context is created from those defaults: `...:dist/src/auth.js:183`
- `useAuth` reads that context: `...:dist/src/auth.js:921`
- The real provider guard is asynchronous: `...:dist/src/auth.js:222`
- Its `validateMenu` and refresh awaits can reject: `...:dist/src/auth.js:249`, `:262`
- Its declaration incorrectly exposes `guard: () => void`: `...:dist/src/auth.d.ts:27`

On the host, production selects the legacy Ordent provider because the production migration toggle is false:

- Toggle: `9b7e1d39:config.ts:472-477`
- Provider selection: `9b7e1d39:pages/_app.tsx:94`
- Provider wraps the route component: `9b7e1d39:pages/_app.tsx:99-129`
- Legacy provider import and mount: `9b7e1d39:components/providers/OrdentProviders.tsx:8-20`, `:28-38`

Even migration mode nests an Ordent provider inside the BRI provider, so the issue is module/context identity, not complete absence of an Ordent provider: `9b7e1d39:components/providers/BRIProviders.tsx:21-28`, `:42-46`.

**Every Hook Site**

| Mass-transfer site | `useAuth` | `guard()` |
|---|---:|---:|
| Maker confirmation list | `265beed1:pages/mass-transfer/maker-confirmation/index.tsx:19` | `:23-25` |
| Maker confirmation batch | `265beed1:pages/mass-transfer/maker-confirmation/view/[batch].tsx:17` | `:20-22` |
| Maker confirmation transaction | `265beed1:pages/mass-transfer/maker-confirmation/view/detail/[batch]/[cureff].tsx:18` | `:30-32` |
| Transaction action list | `265beed1:pages/mass-transfer/transaction-action/index.tsx:17` | `:20-22` |
| Transaction action batch | `265beed1:pages/mass-transfer/transaction-action/view/[batch].tsx:18` | `:21-23` |
| Transaction action transaction | `265beed1:pages/mass-transfer/transaction-action/view/detail/[batch]/[cureff].tsx:17` | `:29-31` |
| Transfer inquiry list | `265beed1:pages/mass-transfer/transfer-inquiry/index.tsx:17` | `:19-21` |
| Transfer inquiry batch | `265beed1:pages/mass-transfer/transfer-inquiry/view/[batch].tsx:17` | `:20-22` |
| Transfer inquiry transaction | `265beed1:pages/mass-transfer/transfer-inquiry/view/detail/[batch]/[cureff].tsx:17` | `:29-31` |
| Upload transaction | `265beed1:pages/mass-transfer/upload-transaction/index.tsx:15` | `:19-21` |
| Nested transfer inquiry container | `265beed1:containers/transfer-inquiry/detail.tsx:51` | No guard; default `roleIDs: null` can reach the request at `:489-493` |

Every guard call is a bare call inside `useEffect(() => { guard(); }, [])`. None is awaited, returned, wrapped in `try/catch`, or followed by `.catch()`.

The default guard throws synchronously, so it is not a rejected Promise. The real async guard can reject, and those rejections become unhandled because every caller discards the Promise.

**Federation And Global Components**
Mass-transfer federation configuration:

- Container name and entry: `265beed1:next.config.js:5-7`
- Only remote is `qcash-global-component`; build configuration hard-codes the development entry: `265beed1:next.config.js:8-10`
- Ten route modules plus `./globals` CSS are exposed: `265beed1:next.config.js:12-23`
- Explicit shared modules are auth, `ni18n`, `i18next`, and `react-i18next`, all singleton with version checking disabled: `265beed1:next.config.js:26-43`
- `react` and `react-dom` are not explicitly listed.
- Plugins are installed at `265beed1:next.config.js:56-57`.
- Standalone base path is `/mass-transfer-v2`: `265beed1:next.config.js:48`.

Resolved federation versions:

- `nextjs-mf` 8.2.5: `265beed1:yarn.lock:433-436`
- Direct runtime 0.1.21: `265beed1:yarn.lock:475-478`
- A separate transitive runtime 0.0.17 also exists: `265beed1:yarn.lock:468-473`
- Federation TypeScript plugin 3.1.3: `265beed1:yarn.lock:492-495`

Standalone runtime loading occurs in `_app`:

- Runtime initialization: `265beed1:pages/_app.tsx:25-34`
- Environment-derived global-component entry: `265beed1:pages/_app.tsx:30-32`
- Fire-and-forget global CSS load: `265beed1:pages/_app.tsx:35`
- `baseFederation` removes `/api` from the selected API URL: `265beed1:config.ts:16-26`

The UI is heavily composed from `qcash-global-component`:

- Forty-nine client-only dynamic component wrappers: `265beed1:components/global.ts:39-182`
- Twenty named icon loaders from its `icons` expose: `265beed1:components/icons.ts:1-43`
- Barrel re-export: `265beed1:components/index.ts:1-4`
- Checked-in remote types: `265beed1:tsconfig.json:20-26`

Under `qcash-ui`, the host performs runtime registration instead:

- Host initializes federation before rendering: `9b7e1d39:pages/_app.tsx:37-48`
- Global component is always registered: `9b7e1d39:constants/features/registry.ts:37-42`
- Mass transfer is registered for `/mass-transfer`: `9b7e1d39:constants/features/registry.ts:254-258`
- Runtime builds and registers all remote entries: `9b7e1d39:services/federation/init.ts:32-50`
- Standard mass-transfer URL becomes `${baseFederation}/mass-transfer-v2/_next/static/chunks/remoteEntry.js`: `9b7e1d39:services/federation/utils/name.ts:1-3`, `9b7e1d39:services/federation/utils/registry.ts:81-99`
- Host eagerly loads global CSS with retry/error handling: `9b7e1d39:services/federation/init.ts:72-84`, `9b7e1d39:services/federation/utils/loader.ts:19-76`

**Render Mode**
The service can run as a standalone Next application, evidenced by its normal `next start`, independent Docker process, base path, and `_app`: `265beed1:package.json:6-10`, `265beed1:Dockerfile:74-79`, `265beed1:next.config.js:48`.

The user-facing host integration renders it as a React remote beneath `qcash-ui`, not as an iframe or standalone navigation:

- Host dynamically loads each exposed mass-transfer page; for example upload transaction: `9b7e1d39:pages/mass-transfer/upload-transaction/index.tsx:13-16`
- Host wraps it in the remote header/footer layout: `9b7e1d39:pages/mass-transfer/upload-transaction/index.tsx:18-27`
- Equivalent loaders exist for all ten host routes: `9b7e1d39:pages/mass-transfer/**`
- Host base path is `/v2`: `9b7e1d39:next.config.js:68`

Therefore `/v2/mass-transfer/...` renders an exposed page component under the host `_app`. The mass-transfer `_app` at `265beed1:pages/_app.tsx` is not executed in that path.

**Concrete Failure Path**
1. A user opens `/v2/mass-transfer/upload-transaction`.
2. The host registers the mass-transfer remote and loads `qcash-ui-mass-transfer/upload-transaction`: `9b7e1d39:services/federation/init.ts:37-50`, `9b7e1d39:pages/mass-transfer/upload-transaction/index.tsx:13`.
3. Production wraps it with host `@ordentco/addons-auth-provider@0.9.136-mfe`: `9b7e1d39:config.ts:472-477`, `9b7e1d39:components/providers/OrdentProviders.tsx:31-38`.
4. The exposed module is the page itself, not the remote `_app`: `265beed1:next.config.js:13`.
5. If share negotiation supplies the page’s `useAuth` from its separate `0.9.137` module while the provider uses the host’s `0.9.136-mfe` module, the two `AuthContext` objects are different.
6. `useAuth()` returns `AUTH_INITIAL_VALUES`; `username` is `"Guest"` and `guard` is the throwing default.
7. Upload transaction renders its loading state at `265beed1:pages/mass-transfer/upload-transaction/index.tsx:23-24`.
8. Its passive effect calls the default guard at `265beed1:pages/mass-transfer/upload-transaction/index.tsx:19-21`.
9. The default throws `Error("Function not implemented.")` synchronously. The host’s surrounding remote error boundary can replace the subtree with its production fallback: `9b7e1d39:components/federation/provider.tsx:171-176`, `9b7e1d39:components/ui/RemoteErrorBoundary.tsx:98-113`.
10. If a matching real provider is found but its async guard rejects instead, no page catches that rejection.

Static source proves this path is reachable but cannot prove which auth entry a particular production browser selected. Runtime confirmation requires inspecting the host federation `shareScopeMap`; the host already contains diagnostics for that map at `9b7e1d39:components/federation/monitor/auth/provider-versions.ts:15-20`.

No files or refs were changed; both worktrees remained unchanged.

---

