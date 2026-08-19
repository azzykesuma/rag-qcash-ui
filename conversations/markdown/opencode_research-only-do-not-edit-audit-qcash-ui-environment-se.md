# Audit qcash environment federation (@explore subagent)

- **ID**: `ses_01142ac8fffe941xP3V63ENVcq`
- **Source Tool**: `opencode`
- **Date**: `2026-08-11 09:53:32`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Audit qcash-ui environment selection and Federation registry for deployment-sensitive errors: ENV_DOMAIN, NEXT_PUBLIC variables, hostname detection, baseFederation, remote URLs, malformed/undefined values, server/client differences, and build-time replacement. Search current code plus cached dev branch/history for relevant changes. Explain whether an environment issue can make qcash-ui's own remoteEntry return 500 versus make a consumed remoteEntry return 500. Return exact references and diagnostic checks.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

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

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 18: Assistant

## Audit result

Current checkout: `v1.630.3-release-com` at `9d4a3e3a`.  
Cached refs inspected: local `dev` at `81bf7427`, cached `origin/dev` at `03ad8635`.

### Main conclusion

An environment-selection error in **qcash-ui** can generate incorrect URLs for remotes that qcash-ui consumes. Those requests may return 500 from the selected upstream, gateway, or—if the malformed URL becomes relative—from qcash-ui’s own origin.

It does **not normally make qcash-ui’s own container** at:

```text
/v2/_next/static/chunks/remoteEntry.js
```

return 500. That file is emitted from static `NextFederationPlugin` configuration with no environment-dependent remotes or exposes. A 500 on that exact resource points primarily to deployment routing, missing/corrupt static assets, proxy behavior, or Next’s error handling—not the runtime registry.

---

## Findings

### 1. `ENV_DOMAIN` is not used by qcash-ui

No `ENV_DOMAIN`, `NEXT_PUBLIC_ENV_DOMAIN`, or equivalent Federation-domain override exists in current qcash-ui or its searched history.

The shell instead derives Federation origin from its API variables:

- `~/Development/qcash-ui/config.ts:5-26`
- `~/Development/qcash-ui/utils/index.ts:16-109`

Sibling MFEs do use `ENV_DOMAIN` conventions, which can cause confusion:

- `~/Development/qcash-ui-complaint/next.config.js:6-15`
- `~/Development/qcash-ui-helpdesk-portal/next.config.js:4-13`
- `~/Development/qcash-ui-issuance-lc/next.config.js:15-25`
- `~/Development/qcash-ui-scm/next.config.js:4-15`

Consequently, an `ENV_DOMAIN` error in a **consumer of qcash-ui** could make that consumer request the wrong qcash-ui URL. It still does not alter qcash-ui’s emitted asset.

---

### 2. Federation origin is coupled to API URL configuration

`baseFederation` is not independently configured:

```ts
baseFederation: getBaseUrlByEnv({...})?.split("/api")[0]
```

Reference:

- `~/Development/qcash-ui/config.ts:16-26`

Risks:

1. An API value without `/api` remains unchanged, producing URLs such as:
   ```text
   https://host/some-api-root/header-footer/...
   ```

2. `/api` is removed using string splitting, not URL/path validation. Values such as `/api-v2` are also truncated.

3. Whitespace is checked with `trim()` but the untrimmed value is returned:
   - `~/Development/qcash-ui/utils/index.ts:102-109`

4. A literal build value `"undefined"` is truthy and passes validation.

5. If the selected environment is empty, code silently falls back to `prod_id`:
   - `~/Development/qcash-ui/utils/index.ts:104-107`

   That can turn a missing dev/staging value into cross-environment production traffic.

6. If both the selected value and `prod_id` are absent, the function returns runtime `undefined` despite its declared `string` return type.

---

### 3. Missing values produce inconsistent malformed URLs

Custom entries interpolate `baseFederation` directly:

```ts
return `${config.baseFederation}/${service}/_next/static/chunks/remoteEntry.js`;
```

- `~/Development/qcash-ui/constants/features/registry.ts:26-32`

With an undefined base, this becomes:

```text
undefined/global-component/_next/static/chunks/remoteEntry.js
```

That is a relative URL. Depending on the current page, the browser can request it from qcash-ui’s own origin.

Standard entries instead default the base to an empty string:

```ts
return buildStandardEntryUrl(name, config.baseFederation ?? "");
```

- `~/Development/qcash-ui/services/federation/utils/registry.ts:81-100`

That becomes:

```text
/main-page-v2/_next/static/chunks/remoteEntry.js
```

Therefore missing configuration has two different failure modes:

- **custom entry:** `undefined/...`, relative to the current location;
- **standard entry:** `/module-v2/...`, absolute on qcash-ui’s current origin.

Either can appear as a qcash-ui-origin 500 even though the request was intended for a consumed MFE.

---

### 4. Hostname detection is client-only and permissive

Environment selection uses `window.location.hostname`:

- `~/Development/qcash-ui/utils/index.ts:16-58`

Rules include:

- `localhost` → dev ID
- `gym-master` plus any occurrence of `tl`/`sg`
- `addons.cms.dev` → prestage
- `addons.apps` → preploy
- `demo.addons` → piloting
- `bricams` plus any occurrence of `tl`/`sg`
- unknown → production ID

Concerns:

- `includes("tl")` and `includes("sg")` inspect the entire hostname rather than a country label.
- New aliases, vanity domains, IP access, preview hosts, and renamed routes silently become production ID.
- The browser hostname—not an ingress forwarding header—is authoritative.
- Unknown hosts do not fail closed; they use production.

The tests confirm this intended fallback:

- `~/Development/qcash-ui/utils/__tests__/index.test.ts:90-100`

---

### 5. Server and client deliberately select different environments

When `window` is unavailable, detection returns `prod_id`:

- `~/Development/qcash-ui/utils/index.ts:16-20`
- `~/Development/qcash-ui/utils/index.ts:29-31`

Thus:

- server-side module evaluation uses production ID;
- browser evaluation uses browser hostname.

Federation initialization itself is browser-gated:

- `~/Development/qcash-ui/pages/_app.tsx:38-49`

Most Federation providers and consumed remotes are also `ssr: false`:

- `~/Development/qcash-ui/pages/_app.tsx:62-78`

This limits server/client Federation disagreement, but `config` is imported by `_app` at:

- `~/Development/qcash-ui/pages/_app.tsx:9-11`

and environment-dependent feature flags can still differ between server rendering and hydration.

---

### 6. `NEXT_PUBLIC_*` values are build-time inputs

The Docker build receives and exports all API URLs before `yarn build`:

- `~/Development/qcash-ui/Dockerfile:81-114`

The runtime image does not re-declare them:

- `~/Development/qcash-ui/Dockerfile:124-136`

For client bundles, Next replaces `NEXT_PUBLIC_*` references during build. Changing only Kubernetes/OpenShift runtime environment variables after the image is built will not change browser remote URLs.

Relevant local values:

- `~/Development/qcash-ui/.env:1-8`

Notably, the checked-in `.env` has no `NEXT_PUBLIC_API_URL_PILOTING`, although both code and Docker expect it:

- `~/Development/qcash-ui/config.ts:11,22`
- `~/Development/qcash-ui/Dockerfile:92-93`

A piloting build therefore depends on CI supplying that argument. Otherwise it silently falls back to production ID.

---

### 7. Registry construction registers all features

Federation initialization calls `getAllFeatures()` and constructs entries for every registry item:

- `~/Development/qcash-ui/services/federation/init.ts:37-50`
- `~/Development/qcash-ui/services/federation/utils/registry.ts:11-13`

Feature toggles are not applied during remote registration. They affect route selection later. Consequently, malformed values can be present in the runtime registry even for disabled features.

Global remotes are particularly important because global styles/providers are loaded immediately:

- `~/Development/qcash-ui/services/federation/init.ts:72-84`
- `~/Development/qcash-ui/constants/features/registry.ts:37-52`

Current global entry patterns are:

```text
${baseFederation}/global-component/_next/static/chunks/remoteEntry.js
${baseFederation}/header-footer/_next/static/chunks/remoteEntry.js
${baseFederation}/design-system/static/remoteEntry.js
```

---

## qcash-ui’s own `remoteEntry`

Its own container is configured independently of the runtime registry:

- `~/Development/qcash-ui/next.config.js:7-13`
- `~/Development/qcash-ui/next.config.js:154`

Configuration:

```js
name: "qcash-ui",
filename: "static/chunks/remoteEntry.js",
remotes: {},
exposes: {},
```

The application has `basePath: "/v2"`:

- `~/Development/qcash-ui/next.config.js:66-69`
- `~/Development/qcash-ui/.next/routes-manifest.json:1`

Therefore its browser URL is expected to be:

```text
https://<qcash-host>/v2/_next/static/chunks/remoteEntry.js
```

The cached build contains:

- `~/Development/qcash-ui/.next/static/chunks/remoteEntry.js`
- `~/Development/qcash-ui/.next/static/chunks/remoteEntry.js.gz`
- `~/Development/qcash-ui/.next/static/ssr/remoteEntry.js`
- `~/Development/qcash-ui/.next/server/chunks/remoteEntry.js`

The standalone directory itself does not contain the browser entry, but Docker explicitly copies `.next/static` into the final image:

- `~/Development/qcash-ui/Dockerfile:132-135`

### Can environment selection make this exact endpoint return 500?

**Directly: no.**

The endpoint’s generation and URL do not reference `baseFederation`, hostname detection, `ENV_DOMAIN`, or API URL variables. Furthermore, an HTTP 500 is decided before any downloaded JavaScript executes; a runtime Federation exception cannot retroactively turn a successful static response into HTTP 500.

Possible indirect cases:

- `/v2` is omitted and ingress forwards the request into a failing Next 404/error renderer.
- `.next/static` was not copied into the deployed image.
- an ingress rewrite sends the static path to another upstream.
- compressed asset negotiation/proxy handling fails.
- the container is starting or crashing while serving requests.
- a consumer’s bad domain points “qcash-ui remoteEntry” at a different application that returns 500.

---

## Relevant history

- `ea7a69d4` — introduced hostname-driven API selection and Docker build arguments.
- `fe867d07` — added piloting selection.
- `36c989b7` — tightened prestage hostname matching from `addons.cms` to `addons.cms.dev`.
- `bec5526a` — refactored selection into current SSR-safe environment detection.
- `07856e41` — introduced object-form URL arguments and current fallback validation.
- `f0945ed4` — introduced the `remoteEntry(service)` registry helper.
- `a5f574e5` — moved dynamic/local URL construction into `module-entry.ts`.
- `9d4a3e3a` — current release commit; disables the timestamp-stripping plugin.

The cached dev implementation has the same fundamental hostname selection, API-derived `baseFederation`, and custom/standard URL behavior. No cached dev change fixes runtime/build-time environment promotion.

---

## Diagnostic checks

### 1. Test the exact shell asset and base path

```bash
curl -vkI https://HOST/v2/_next/static/chunks/remoteEntry.js
curl -vk  https://HOST/v2/_next/static/chunks/remoteEntry.js
curl -vkI https://HOST/_next/static/chunks/remoteEntry.js
```

Expected:

- first URL: `200`, JavaScript content type;
- no-base-path URL: normally 404/redirect, not the canonical resource.

Record response headers, body, ingress request ID, `Server`, `Via`, and upstream name. A branded gateway error indicates the 500 is upstream of Next.

### 2. Verify the deployed image, not merely the builder

Inside the running container:

```bash
test -s /usr/src/app/addons-build/.next/static/chunks/remoteEntry.js
test -s /usr/src/app/addons-build/.next/BUILD_ID
```

Also verify that the running image was built through the current Dockerfile’s final copy steps.

### 3. Inspect actual browser-selected remotes

In DevTools:

```js
location.hostname
performance
  .getEntriesByType("resource")
  .map(x => x.name)
  .filter(x => x.includes("remoteEntry"))
```

Check for:

```text
undefined/
null/
%22
/api/
double slashes after the hostname
wrong environment hostname
missing /v2 on qcash-ui itself
```

Use the Network panel’s full **Request URL**, response body, and remote IP—not only the displayed filename.

### 4. Validate build arguments before building

Require every expected value to be nonempty and a valid absolute URL, especially:

```text
NEXT_PUBLIC_API_URL_PILOTING
NEXT_PUBLIC_API_URL_PROD_ID
the environment matching the deployment hostname
```

Do not pass literal strings `undefined` or `null`. Confirm the API values terminate at the intended `/api` boundary.

### 5. Verify build-time replacement

Inspect the deployed client chunks for known environment hosts and malformed literals. Runtime pod environment alone is insufficient because public variables were inlined during `yarn build`.

Rebuild the image after correcting any `NEXT_PUBLIC_*` value; restarting the old image will not update browser URLs.

### 6. Separate shell failure from consumed-remote failure

- If `/v2/_next/static/chunks/remoteEntry.js` itself is 500: inspect qcash-ui ingress/static packaging/server logs.
- If a URL such as `/header-footer/...`, `/giro-v2/...`, or `undefined/...` is 500: inspect qcash-ui’s generated registry URL and the selected remote upstream.
- Curl the failing consumed URL directly from both outside and inside the qcash-ui pod. If it returns 500 independently, qcash-ui is only the caller.

---

