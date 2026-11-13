# i want to fix this issue once and for all

- **ID**: `antigravity-ed26f474a8789dcc7d2ebc4f35f51c408597635a0ce5bf8ef63b8d8d523b781a`
- **Source Tool**: `antigravity`
- **Date**: `2026-10-30 15:46:40`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

i want to fix this issue once and for all
  Running yarn dev without NEXT_PRIVATE_LOCAL_WEBPACK=true failed with:                                                                    
                                                                                                                                           
    Error: Cannot find module 'C:\Work\qcash-ui-user\node_modules\next\dist\compiled\lib\util\makeSerializable'                            
                                                                                                                                           
  Because @module-federation/nextjs-mf in next.config.js requires the local Webpack package instead of Next.js's precompiled Webpack       
  bundle, I have relaunched yarn dev in the background with NEXT_PRIVATE_LOCAL_WEBPACK=true enabled                                        
  ($env:NEXT_PRIVATE_LOCAL_WEBPACK="true"; yarn dev). Let me know if you would also like me to add this environment variable to            
  package.json:6.   
whenever i tried to run locally, this happen. 
it seems there is something wrong with the yarn dev command, or the webpack
how do we fix this

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 11: Assistant

### Why This Keeps Happening

1. **Why Next.js crashes without `NEXT_PRIVATE_LOCAL_WEBPACK=true`:**
   By default, Next.js 13 uses its own precompiled, stripped-down Webpack bundle (`next/dist/compiled/webpack/bundle5.js`), which omits internal Webpack utilities like `webpack/lib/util/makeSerializable` (checked in [webpack.js](file:///C:/Work/qcash-ui/node_modules/next/dist/compiled/webpack/webpack.js#L5-L26)).
   However, `@module-federation/nextjs-mf` ([NextFederationPlugin](file:///C:/Work/qcash-ui/node_modules/@module-federation/nextjs-mf/dist/src/plugins/NextFederationPlugin/index.js#L41-L79)) requires the full `webpack` package from `node_modules/webpack` and resolves `makeSerializable` relative to the active Webpack path. When `NEXT_PRIVATE_LOCAL_WEBPACK=true` is not set before `next dev` starts, it looks inside `next/dist/compiled/lib/util/makeSerializable` and crashes.

2. **Why `yarn dev` works in `qcash-ui` but fails in `qcash-ui-user` (and ~45 other local QCash repos):**
   - In [qcash-ui/package.json](file:///C:/Work/qcash-ui/package.json#L9), `"dev"` is properly configured for Windows using `cross-env`:
     ```json
     "dev": "cross-env NEXT_PRIVATE_LOCAL_WEBPACK=true next dev"
     ```
   - In [qcash-ui-user/package.json](file:///C:/Work/qcash-ui-user/package.json#L6-L7) (and ~45 other `C:\Work\qcash-ui-*` repos), `"dev"` is missing the variable entirely:
     ```json
     "dev": "yarn build-style && next dev",
     "build": "yarn build-style && NEXT_PRIVATE_LOCAL_WEBPACK=true next build"
     ```
   - Additionally, some repos use Unix-only inline syntax (`NEXT_PRIVATE_LOCAL_WEBPACK=true next dev` or `export NEXT_PRIVATE_LOCAL_WEBPACK=true && ...`) without `cross-env`. Because `yarn` runs scripts via `cmd.exe` on Windows, inline Unix env syntax only works on macOS/Linux/Docker and fails on Windows.

---

### How to Fix It Once and For All

Combine both fixes below so you never hit this again:

#### 1. Set `NEXT_PRIVATE_LOCAL_WEBPACK=true` Globally in Windows (Fixes ~45 repos immediately)
By setting `NEXT_PRIVATE_LOCAL_WEBPACK=true` in your **Windows User Environment Variables**, every terminal and background task on your machine will inherit it automatically—even in repos where `"dev"` is just `"yarn build-style && next dev"` (like [qcash-ui-user/package.json](file:///C:/Work/qcash-ui-user/package.json#L6)):

```powershell
[Environment]::SetEnvironmentVariable("NEXT_PRIVATE_LOCAL_WEBPACK", "true", "User")
```

#### 2. Fix Broken Unix-Style Scripts in `package.json`
For any repo that hardcodes Unix-only syntax in `"dev"` (like `NEXT_PRIVATE_LOCAL_WEBPACK=true next dev` or `export NEXT_PRIVATE_LOCAL_WEBPACK=true && ...` without `cross-env`), or if you want [qcash-ui-user/package.json](file:///C:/Work/qcash-ui-user/package.json#L6-L7) fixed in git for other Windows devs, use `cross-env` (or Remove the inline Unix prefix once the global env var is set).

---

Would you like me to:
1. **Set `NEXT_PRIVATE_LOCAL_WEBPACK=true` in your Windows User Environment Variables** right now so all `"yarn build-style && next dev"` repos work out of the box?
2. **Also check/fix any `C:\Work\qcash-ui-*` repos** that have invalid Windows syntax (like `export NEXT_PRIVATE_LOCAL_WEBPACK=true` or inline `NEXT_PRIVATE_LOCAL_WEBPACK=true` without `cross-env`)?

---

## Turn 12: User

okay it works now on the qcash ui user. i need the cross env to run the dev, and add   "dev": "cross-env NEXT_PRIVATE_LOCAL_WEBPACK=true next dev", 
explain to me why this works

---

## Turn 13: Assistant

`"dev": "cross-env NEXT_PRIVATE_LOCAL_WEBPACK=true next dev"` solves two separate problems at once: **how Next.js loads Webpack** and **how Windows handles environment variables in `package.json`**.

---

### 1. What `NEXT_PRIVATE_LOCAL_WEBPACK=true` does (Next.js + Module Federation)

When `next dev` starts, Next.js runs [`init()`](file:///C:/Work/qcash-ui/node_modules/next/dist/compiled/webpack/webpack.js#L5-L26) in `node_modules/next/dist/compiled/webpack/webpack.js`:

```js
exports.init = function () {
  if (process.env.NEXT_PRIVATE_LOCAL_WEBPACK) {
    Object.assign(exports, {
      // Loads the full webpack package from node_modules/webpack
      webpack: require('webpack'),
      ...
    })
  } else {
    // Loads Next.js's stripped-down precompiled webpack bundle
    Object.assign(exports, require('./bundle5')())
  }
}
```

- **Without `NEXT_PRIVATE_LOCAL_WEBPACK=true`:**
  Next.js uses its own precompiled bundle (`next/dist/compiled/webpack/bundle5.js`). To reduce package size, Next.js bundles all of Webpack into a single file and omits individual internal files like `lib/util/makeSerializable.js`.
  When `@module-federation/nextjs-mf` ([`NextFederationPlugin`](file:///C:/Work/qcash-ui/node_modules/@module-federation/nextjs-mf/dist/src/plugins/NextFederationPlugin/index.js#L41-L45)) inspects the active Webpack path via [`getWebpackPath`](file:///C:/Work/qcash-ui/node_modules/@module-federation/sdk/dist/normalize-webpack-path.cjs.js#L11-L48) and tries to require `lib/util/makeSerializable` relative to it, Node looks for `node_modules/next/dist/compiled/lib/util/makeSerializable`—which does not exist—and crashes.

- **With `NEXT_PRIVATE_LOCAL_WEBPACK=true`:**
  Next.js skips its internal bundle and loads the full `webpack` package from `node_modules/webpack`, where `node_modules/webpack/lib/util/makeSerializable.js` exists as a real file.

---

### 2. Why `cross-env` is required on Windows

When you run `yarn dev`, Yarn executes the script string from `package.json` using the operating system's default shell—on Windows, that is **`cmd.exe`**.

Every shell has different syntax for setting an inline environment variable:
- **Linux / macOS (`bash` / `sh`):** `NEXT_PRIVATE_LOCAL_WEBPACK=true next dev`
- **Windows (`cmd.exe`):** `set NEXT_PRIVATE_LOCAL_WEBPACK=true&& next dev`
- **Windows (PowerShell):** `$env:NEXT_PRIVATE_LOCAL_WEBPACK="true"; next dev`

If you put `NEXT_PRIVATE_LOCAL_WEBPACK=true next dev` directly into `package.json` without `cross-env`, `cmd.exe` on Windows thinks `NEXT_PRIVATE_LOCAL_WEBPACK=true` is the name of an `.exe` program and fails.

**`cross-env`** is a small Node CLI that:
1. Reads `NEXT_PRIVATE_LOCAL_WEBPACK=true` as an argument,
2. Sets `process.env.NEXT_PRIVATE_LOCAL_WEBPACK = "true"` inside Node, and
3. Spawns `next dev` with that environment variable inherited—so the exact same script works identically on Windows, macOS, and Linux.

---

## Turn 14: User

okay i want you to check the core architecture for the qcash ui. please investigate the issues that exists in the current architecture and group them by severity

---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read_file**: 
- **read_file**: 
- **run_command**: 
- **run_command**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **run_command**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **manage_task**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **run_command**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 
- **read_file**: 

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_file**: 

</details>



---

## Turn 27: Assistant

I have completed the investigation of the `qcash-ui` core architecture and compiled the full report—grouped into **Critical**, **High**, **Medium**, and **Low / Dead Code** severity tiers—in [qcash_ui_architecture_audit.md](file:///C:~/.gemini/antigravity-cli/brain/4ba1cfaa-5d07-45c4-b38e-1ea584a41021/qcash_ui_architecture_audit.md).

### Key Decisions / Next Steps
Which area would you like to prioritize fixing first?
1. **🔴 Critical Security & Bootstrap Fixes**: Removing the secret key console log in [`services/auth.ts`](file:///C:/Work/qcash-ui/services/auth.ts#L9), validating `auth_redirect_path` and removing the 500ms polling loop in [`pages/_app.tsx`](file:///C:/Work/qcash-ui/pages/_app.tsx#L101-L161), and unblocking the top-level `await initFederation()`.
2. **🟠 High-Severity Runtime & Network Fixes**: Fixing the `customAxios` queue memory leak and `return error` bug in [`services/custom-axios.ts`](file:///C:/Work/qcash-ui/services/custom-axios.ts#L71-L155), fixing the `createRemoteDynamic` double-mount pattern in [`remote-dynamic.tsx`](file:///C:/Work/qcash-ui/components/federation/remote-dynamic.tsx#L31-L103), and removing the `JSON.stringify(featureRegistry)` overhead in [`registry.ts`](file:///C:/Work/qcash-ui/services/federation/utils/registry.ts#L160-L165).
3. **🟡 Medium Registry & Config Fixes**: Fixing the route/toggle bugs in [`constants/features/registry.ts`](file:///C:/Work/qcash-ui/constants/features/registry.ts#L323) (`qcash-ui-advise-lc` missing `/` and `qcash-ui-global-account-statement` wrong toggle key) and cleaning up [`next.config.js`](file:///C:/Work/qcash-ui/next.config.js#L15-L62).
4. **🟢 Dead Code Elimination**: Deleting ~2,200+ lines of unused code (`urshifuPlugin.tsx`, unused `SecurityValidator`/`RateLimiter` classes, legacy `retryException`, and unused `SessionExpired` components).

---

## Turn 28: User

make it on the table please

---

## Turn 29: Assistant

Here is the complete `qcash-ui` core architecture audit organized in a table, grouped by severity:

| Severity | # | Category | Issue | Location | Impact | Recommended Fix |
| :--- | :---: | :--- | :--- | :--- | :--- | :--- |
| 🔴 **Critical** | 1 | **Security / Auth** | **Client-Side HMAC Secrets & Console Leak**: `CreateSignatureForAuth` logs `NEXT_PUBLIC_SECRET_KEY_LOGIN_V2` to browser console (`console.log("⚠️Secret Key:", secret)`), and `NEXT_PUBLIC_SIGNATURE` + hardcoded fallbacks (`"addons"`, `"dmljdG9yIHdhcyBoZXJlIQ=="`, `"BRIC@MS4dd0ns"`) are bundled into client JS. | [`services/auth.ts`](file:///C:/Work/qcash-ui/services/auth.ts#L5-L12)<br/>[`custom-axios.ts`](file:///C:/Work/qcash-ui/services/custom-axios.ts#L168-L185)<br/>[`encryptPassword.tsx`](file:///C:/Work/qcash-ui/utils/encryptPassword.tsx#L25-L28)<br/>[`transfer-signed-post.ts`](file:///C:/Work/qcash-ui/services/server/transfer-signed-post.ts#L81-L82) | Anyone with browser DevTools can read the login secret and HMAC signing keys to forge signed backend requests. | Remove `console.log` immediately; move all HMAC signing to server-only Next.js API routes using non-`NEXT_PUBLIC_` env vars (like [`multipayment-signed-post.ts`](file:///C:/Work/qcash-ui/services/server/multipayment-signed-post.ts#L139)). |
| 🔴 **Critical** | 2 | **Security / Routing** | **Unvalidated Redirect, `basePath` Bug & 500ms Polling**: `sessionStorage.getItem("auth_redirect_path")` is passed directly to `router.replace()` without validation; `isAuthPage` checks `["/v2/landing-page"]` against `router.pathname` (which strips `/v2`, so it's always `false`); polls `localStorage` every 500ms via `setInterval`. | [`pages/_app.tsx`](file:///C:/Work/qcash-ui/pages/_app.tsx#L101-L161) | Open redirect / unexpected navigation if `auth_redirect_path` is poisoned or set on `/landing-page`; unnecessary 500ms timer loop. | Validate `redirectPath` starts with `/` (not `//`), include `"/landing-page"` in `isAuthPage`, and listen to auth events instead of polling. |
| 🔴 **Critical** | 3 | **Build & CI/CD** | **Disabled Build Safety & Insecure Dockerfile**: `ignoreBuildErrors: true` (TS) and `ignoreDuringBuilds: true` (ESLint) in `next.config.js`. `Dockerfile` runs as `root`, runs `chmod -R 777`, disables SSL (`strict-ssl false`), writes Nexus credentials into layer history, and has duplicate `ARG NO_PROXY` that overwrites itself. | [`next.config.js`](file:///C:/Work/qcash-ui/next.config.js#L77-L78)<br/>[`Dockerfile`](file:///C:/Work/qcash-ui/Dockerfile#L20-L78) | Type/lint bugs ship to production silently; container runs with root privileges and world-writable files. | Enable TS/ESLint checks in CI, run `STAGE 3` as `USER node`, remove `chmod 777`, and deduplicate `ARG` lines in `Dockerfile`. |
| 🔴 **Critical** | 4 | **MFE Bootstrap** | **Top-Level `await` Blocking App Startup**: `_app.tsx` runs `await initFederation()` at module evaluation time, which resolves ~78 remotes and awaits `dynamicLoadStyles()` (with 5s retry delays) for all 3 global remotes before React can even mount. | [`pages/_app.tsx`](file:///C:/Work/qcash-ui/pages/_app.tsx#L49-L70)<br/>[`init.ts`](file:///C:/Work/qcash-ui/services/federation/init.ts#L32-L87)<br/>[`loader.ts`](file:///C:/Work/qcash-ui/services/federation/utils/loader.ts#L54) | If any global remote (`qui`, `header-footer`, `global-component`) is slow or down, the entire app shows a blank screen for 15–20+ seconds. | Remove blocking top-level `await` on global style loading and reduce `RetryPlugin` / `loader.ts` retry delays from `5000ms` to exponential backoff (`500ms`–`1000ms`). |
| 🟠 **High** | 5 | **SSR / Config** | **SSR Environment Detection Defaults to `prod_id`**: `detectEnvironment()` checks `window.location.hostname` and returns `"prod_id"` whenever `window` is undefined (SSR). `config.ts` evaluates once at import time. | [`utils/index.ts`](file:///C:/Work/qcash-ui/utils/index.ts#L16-L59)<br/>[`config.ts`](file:///C:/Work/qcash-ui/config.ts#L4-L769) | On the server in `dev`/`prestage`/`preploy`, `config` uses production URLs and production feature flags, mismatching client hydration. | Add an env var check (e.g., `process.env.NEXT_PUBLIC_APP_ENV`) in `detectEnvironment()` before falling back to `"prod_id"`. |
| 🟠 **High** | 6 | **Networking** | **`customAxios` Queue Memory Leak & Interceptor Bugs**: Calls `enqueue()` on every HTTP error without consuming the promise; `processQueue` never clears `queue`; `pathname === "/main-page?logout=true"` is always `false` (`pathname` has no query string); line 92 does `return error` instead of `Promise.reject(error)`; hard redirects omit `/v2`. | [`services/custom-axios.ts`](file:///C:/Work/qcash-ui/services/custom-axios.ts#L58-L155) | Unbounded memory leak in `queue`, stale promises firing on relogin, failed `/login` calls resolving as success, and broken `/landing-page` redirects. | Remove stray `enqueue()`, reset `queue = []` in `processQueue`, fix query check via `location.search`, return `Promise.reject(error)`, and include `/v2` in `location.href`. |
| 🟠 **High** | 7 | **MFE Loading** | **`createRemoteDynamic` Recreates `next/dynamic` in `useState`**: Calls `dynamic(() => loadRemote(remote))` inside `useEffect` on every mount and stores the component constructor in `useState`. | [`remote-dynamic.tsx`](file:///C:/Work/qcash-ui/components/federation/remote-dynamic.tsx#L9-L103) | Defeats module caching, triggers double loading spinners on every mount, and remounts remote component trees unnecessarily. | Create the `next/dynamic` component once outside the render lifecycle (at `createRemoteDynamic()` call time). |
| 🟠 **High** | 8 | **Performance** | **`JSON.stringify(featureRegistry)` on Every Route Check**: `getModulesForRoute()` serializes the entire 845-line `featureRegistry` object on every call just to compute `configHash` for a `Map` lookup. | [`registry.ts`](file:///C:/Work/qcash-ui/services/federation/utils/registry.ts#L160-L165) | Hashing a static constant via `JSON.stringify` on every route check is slower than the array `.filter()` it is caching. | Remove `configHash = JSON.stringify(featureRegistry)` since `featureRegistry` is a static module constant. |
| 🟡 **Medium** | 9 | **Routing & Pages** | **3 Competing Page Patterns, Layout Remounting & Registry Typos**: 350+ pages each re-declare `ComponentLayoutDefault = dynamic(...)` locally; 3 different remote-loading patterns exist; `qcash-ui-advise-lc` is missing leading `/` in `"lc-export/new-advise-lc"`; `qcash-ui-global-account-statement` uses wrong toggle `toggleGlobalExternalTransfer`. | [`pages/index.tsx`](file:///C:/Work/qcash-ui/pages/index.tsx#L11-L36)<br/>[`SecureFeatureWrapper.tsx`](file:///C:/Work/qcash-ui/components/ui/SecureFeatureWrapper.tsx#L1-L61)<br/>[`constants/features/registry.ts`](file:///C:/Work/qcash-ui/constants/features/registry.ts#L323) | Header/Footer unmounts and remounts on every page change; `/lc-export/new-advise-lc` fails route matching; wrong feature toggle controls Global Account Statement. | Centralize `ComponentLayoutDefault` in `_app.tsx` or `SecureFeatureWrapper`, add `/` to `"/lc-export/new-advise-lc"`, and change toggle to `toggleGlobalAccountStatement`. |
| 🟡 **Medium** | 10 | **Session & State** | **Fragile `window` Globals, `login.split("-")` & Duplicate Relogin Code**: Session prefill splits `localStorage.getItem("login")` by `"-"`; auth relies on 5+ untyped `window.__*` globals and `document` events; `performLogin` duplicates 45 lines across `if/else` and forgets `refresh-token` in `else`. | [`use-modal-session-expired.tsx`](file:///C:/Work/qcash-ui/hooks/use-modal-session-expired.tsx#L91-L173)<br/>[`HostAuthGate.tsx`](file:///C:/Work/qcash-ui/components/providers/HostAuthGate.tsx#L46-L105)<br/>[`use-session-expiration.ts`](file:///C:/Work/qcash-ui/hooks/use-session-expiration.ts#L26-L69) | Any username, company code, or ID with a `-` breaks session-expired prefill; difficult-to-trace race conditions between host and remotes. | Store `login` metadata as JSON instead of hyphen-joined string, deduplicate `performLogin`, and type the shared `window` bridge contract. |
| 🟡 **Medium** | 11 | **Webpack & CSS** | **Polluted Webpack `shared` Config & Uncleaned Remote `<style>` Tags**: `next.config.js` shares build-time packages (`cross-env`, `sass`) and ~30 uninstalled packages; `dynamicLoadStyles` injects global CSS for visited remotes without cleanup. | [`next.config.js`](file:///C:/Work/qcash-ui/next.config.js#L15-L62)<br/>[`loader.ts`](file:///C:/Work/qcash-ui/services/federation/utils/loader.ts#L19-L38) | Bloated federation share scope and progressive CSS specificity conflicts across MFE navigations. | Prune `federationConfig.shared` to actual runtime shared dependencies (`@bri/addons-auth-provider`, `ni18n`, `i18next`, `react-i18next`). |
| 🟢 **Low** | 12 | **Dead Code** | **~2,200+ Lines of Unused Plugins, Validators & Legacy Components**: `urshifuPlugin.tsx` (+ tests), `SecurityValidator` & `RateLimiter` in `security.ts` (+ tests), `stripTimestampPlugin`, `retryException` (v1), `SessionExpired.tsx`, `SessionExpiredLayout.tsx`, and `preload.ts` are unused in production. | [`urshifuPlugin.tsx`](file:///C:/Work/qcash-ui/services/federation/urshifuPlugin.tsx#L1-L657)<br/>[`security.ts`](file:///C:/Work/qcash-ui/services/federation/utils/security.ts#L22-L193)<br/>[`custom-axios.ts`](file:///C:/Work/qcash-ui/services/custom-axios.ts#L157-L190)<br/>[`SessionExpired.tsx`](file:///C:/Work/qcash-ui/components/ui/SessionExpired.tsx) | Unnecessary maintenance burden, slower test suite runs, and misleading abstractions that appear active but are commented out or unreferenced. | Delete the unreferenced files, dead exports, and their corresponding dead unit tests. |
| 🟢 **Low** | 13 | **Config Bloat** | **769-Line Repetitive Feature Toggle File**: `config.ts` repeats 7-line `getValueByEnv(true, true, true, true, true)` blocks ~100 times. | [`config.ts`](file:///C:/Work/qcash-ui/config.ts#L1-L769) | High visual noise and error-prone manual editing for 100+ boolean flags. | Consolidate common toggle presets (`ALL_ENVS = true`, `DEV_PRESTAGE_ONLY = getValueByEnv(true, true, false, false, false)`), cutting ~650 lines. |

Would you like me to start fixing any of these—for example, starting with the **🔴 Critical** and **🟠 High** runtime issues, or cleaning up the **🟢 Dead Code**?

---

