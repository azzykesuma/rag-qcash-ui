# Trace shell module loading (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_07d1f3812ffeOzbe0x0lgPY8UK`
- **Source Tool**: `opencode`
- **Date**: `2026-07-21 11:13:17`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Inspect D:\work\qcash-ui thoroughly enough to explain the exact integration contract for displaying a new microfrontend module at a URL like /v2/dashboard-auto-collection/dashboard. Find route handling, feature/module registry, runtime federation discovery, remote URL construction, API/config dependencies, error fallback trigger, and deployment-related configuration. Do not edit files. Return concise findings with absolute paths and line numbers, including the required module name, exposed module key, remote entry/assets expectations, and likely causes of the shown 'Feature temporarily unavailable' page.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (13)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (16)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (13)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (14)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

## Exact integration contract

- **Browser URL:** `/v2/dashboard-auto-collection/dashboard`. `/v2` is the shell’s Next.js `basePath`; the page route itself is `/dashboard-auto-collection/dashboard`.  
  `D:\work\qcash-ui\next.config.js:66-69`  
  `D:\work\qcash-ui\pages\dashboard-auto-collection\dashboard.tsx:20-35`

- **Registry identity:** `qcash-ui-dashboard-auto-collection`, always enabled because it has no `toggleKey`. Its route defaults to prefix matching.  
  `D:\work\qcash-ui\constants\features\registry.ts:680-686`  
  `D:\work\qcash-ui\services\federation\utils\registry.ts:104-120,160-187`

- **Required component expose:** host calls  
  `qcash-ui-dashboard-auto-collection/dashboard`, so the remote must expose **`./dashboard`**, returning a React component/default module compatible with `next/dynamic`. The shell passes one prop: Next’s `router`.  
  `D:\work\qcash-ui\pages\dashboard-auto-collection\dashboard.tsx:14-24`

- **Required/expected style expose:** route preloading calls  
  `qcash-ui-dashboard-auto-collection/globals`, so the remote is expected to expose **`./globals`**. Missing globals are replaced with empty CSS rather than the full-page fallback.  
  `D:\work\qcash-ui\services\federation\utils\loader.ts:19-37`  
  `D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:21-39`

- **Layout dependency:** page content is wrapped by `qcash-ui-header-footer/default`. A failure here uses a special header/footer indicator, not normally the shown generic page.  
  `D:\work\qcash-ui\pages\dashboard-auto-collection\dashboard.tsx:8-12,27-33`  
  `D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:41-47`

## Discovery and URLs

There is no manifest/API-based remote discovery. The shell statically enumerates the registry, calculates every entry URL, and registers them with `@module-federation/runtime` under the registry name.  
`D:\work\qcash-ui\services\federation\init.ts:32-50`

For this module, the deployed entry is:

```text
{baseFederation}/dashboard-auto-collection-v2/_next/static/chunks/remoteEntry.js
```

`baseFederation` is the selected `NEXT_PUBLIC_API_URL_*` with `/api` removed.  
`D:\work\qcash-ui\constants\features\registry.ts:26-32,680-686`  
`D:\work\qcash-ui\config.ts:16-26`

The deployment must therefore serve not only `remoteEntry.js`, but all chunk/static URLs referenced by it under the **`/dashboard-auto-collection-v2/_next/static/...`** service prefix. Incorrect remote `basePath`/`publicPath` is a likely chunk-load failure.

For local development, the hard-coded mapping rewrites the entry to:

```text
http://localhost:3001/dashboard-auto-collection-v2/_next/static/chunks/remoteEntry.js
```

`D:\work\qcash-ui\constants\features\registry-dev.ts:6-10`  
`D:\work\qcash-ui\services\federation\utils\module-entry.ts:5-25`

## Configuration/API dependencies

- Environment selection is hostname-based; unknown hostnames default to production Indonesia. Missing or wrongly baked `NEXT_PUBLIC_API_URL_*` can therefore point federation at the wrong origin.  
  `D:\work\qcash-ui\utils\index.ts:24-59,65-109`
- These public URLs are build arguments in the standalone Docker build, so changing runtime container environment alone will not reliably change the browser bundle.  
  `D:\work\qcash-ui\Dockerfile:81-114,124-136`
- Namespace `dashboard-auto-collection` is requested, but no corresponding shell locale JSON exists and it is not in the migrated-language set. Loading falls through among shell `/v2/locales/...`, the language API, and remote `/dashboard-auto-collection-v2/locales/...`; locale failures are swallowed and should not cause the shown page.  
  `D:\work\qcash-ui\constants\features\registry.ts:683-685`  
  `D:\work\qcash-ui\ni18n.config.ts:18-52,80-97`  
  `D:\work\qcash-ui\services\federation\utils\i18n.ts:38-57,117-142,162-219`
- Production routing must expose the shell under `/v2` and the remote under the separate `/dashboard-auto-collection-v2` prefix.  
  `D:\work\qcash-ui\next.config.js:68-80`

## Why “Feature temporarily unavailable” appears

That exact page is returned by the runtime `errorLoadRemote` hook for a **non-CSS regular remote request**.  
`D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:89-127`

Most likely causes:

1. Local remote is not running on port `3001`, or it serves `remoteEntry.js` at a different path.
2. Deployed entry returns 404/HTML, is unreachable, or the selected environment points at the wrong host.
3. Remote federation name/build is incompatible, or **`./dashboard` is absent/misnamed**.
4. `remoteEntry.js` loads, but one of its referenced JS chunks uses the wrong service prefix/public path or returns 404.
5. TLS, CSP, proxy/ingress, authentication redirect, or stale/cached entry serves invalid JavaScript.
6. Less likely, another regular global remote such as `qui/providers` fails and produces the same generic fallback.

The runtime config retries script loading three times before fallback handling. The console logs the precise failing module ID via `FallbackPlugin: Providing fallback for module ...`.  
`D:\work\qcash-ui\services\federation\init.ts:48-65`  
`D:\work\qcash-ui\services\federation\fallbackPlugin.tsx:17-20`

A missing `./globals` or translation bundle alone should not generate this exact page; federation initialization timeout or a component render exception uses a different error boundary UI.  
`D:\work\qcash-ui\components\federation\provider.tsx:143-158,190-198`  
`D:\work\qcash-ui\components\ui\RemoteErrorBoundary.tsx:64-102`

---

