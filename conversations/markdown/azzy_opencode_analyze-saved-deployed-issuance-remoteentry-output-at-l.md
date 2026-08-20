# Analyze deployed issuance bundle (@general subagent)

- **Author**: `azzy`
- **ID**: `ses_049add31cffeaXEFLH24dIe74m`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 10:57:46`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Analyze the saved deployed issuance remoteEntry output at ~\.local\share\opencode\tool-output\tool_fb651d403001DQE30uu8hjaqt0. Research only. Find its Module Federation shared dependency registrations, especially @ordentco/addons-auth-provider, next/router/next, and versions; exposed module chunk IDs/URLs; any clues to exact deployed source/version. If useful, fetch only linked deployment chunks from gym-master to search for routeChangeStart, setIsAuthoritiesReady(false), or source strings. Return concise exact evidence.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

I’ll inspect the saved `remoteEntry` first, then follow only its `gym-master` chunk references if they can resolve the requested source/version clues. No workspace files will be changed.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
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

The container identifies itself as `qcash-ui-issuance-lc`. Its federation map exposes nine modules and registers auth provider `0.9.119-mfe` plus Next `13.5.6`; I’m now resolving the runtime’s exact hashed chunk filenames and deployment-origin evidence before deciding which linked chunks to fetch.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>

The deployment path is confirmed by the matching source config’s `basePath: "/issuance-lc-v2"`: the linked hashed chunks at that path return successfully. I’m searching those exact root/detail/create/edit/auth chunks now; the earlier `/issuance-lc` requests were 404 and yielded no evidence.

---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

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

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 20: Assistant

**Federation Registration**
Saved artifact evidence: line 53, columns 58,410-67,705.

- Container: `qcash-ui-issuance-lc`, default share scope.
- `@ordentco/addons-auth-provider`: version `0.9.119-mfe`, singleton, `requiredVersion: false`, provider chunk `6676` -> module `86676`.
- Auth consumer fallback: chunk `4295.2747c0e36842eee2.js` -> the same module `86676`.
- `next/router`: version `13.5.6`, singleton, `requiredVersion: false`, provider chunk `6715` -> module `86715`.
- Router consumer fallback: chunk `9096.30aed0e34ca50a9d.js` -> module `86715`.
- No bare `next` shared registration. Registered Next subpaths are `next/dynamic`, `next/head`, `next/image`, `next/link`, `next/router`, and `next/script`, all `13.5.6`.
- Other provided versions: `react`, `react-dom`, `react-dom/client`, `react/jsx-runtime` `18.3.1`; `i18next` `22.5.1`; `react-i18next` `12.3.1`; `ni18n` `1.1.0`; `styled-jsx` and `styled-jsx/style` `5.1.1`.

**Exposed Modules**
Base URL confirmed by successful chunk fetches:  
`http://internal-service.example.com/issuance-lc-v2/_next/static/chunks/`

| Expose | Chunk -> module | Exact chunk |
|---|---:|---|
| `./noop` | `6579` -> `56579` | `6579.577721810aa13e49.js` |
| `./react` | `6540` -> `96540` | `6540.ecdbced217d478c4.js` |
| `./react-dom` | `961` -> `40961` | `961.e9e42d156edcb4db.js` |
| `./next/router` | `6715` -> `86715` | `6715.3b68cbd3727b8b06.js` |
| `./issuance-lc` | `7244` -> `87244` | `7244.21cd94b56623ce0e.js` |
| `./issuance-lc/detail` | `5991` -> `75991` | `5991.6082ecac1444b126.js` |
| `./issuance-lc/create` | `8606` -> `48606` | `8606.9e66cd8931176ed7.js` |
| `./issuance-lc/edit` | `2741` -> `92741` | `2741.9866a4b657295bff.js` |
| `./globals` | `8965` -> `28965` | `8965.349c9f415b8f8310.js` |

**Route/Auth Evidence**
The four issuance page chunks contain `isAuthoritiesReady` checks but no `routeChangeStart` or literal `setIsAuthoritiesReady`.

The linked auth-provider chunk `6676.9132f8b159f8c5f5.js` contains:

```js
eB=(0,O.useState)(!1),eD=eB[0],eL=eB[1]
```

and:

```js
c.split("?")[0]!==e3.current.split("?")[0]&&(eL(!1),e3.current=c)
m.events.on("routeChangeStart",handleRouteChangeStart)
m.events.off("routeChangeStart",handleRouteChangeStart)
```

Thus `eL(!1)` is the minified `setIsAuthoritiesReady(false)`. It exactly corresponds to the installed `0.9.119-mfe` source at:

- `node_modules/@ordentco/addons-auth-provider/src/auth.tsx:205`
- `node_modules/@ordentco/addons-auth-provider/src/auth.tsx:776-790`

The package lock resolves that version to tarball SHA fragment `ae0510ceac63bb0d16bc8265314d01c6c856b959`.

**Likely Source Revision**
Strongest repository candidate is `qcash-ui-issuance-lc` commit `20b79a5cff9edeb346df2d749f356830dc3d76ee` from 2026-07-20:

- Its config has the exact base path `/issuance-lc-v2`, container name, exposes, and only the deployed `qcash-global-component`/`qui` remotes.
- Its package uses Next `13.5.6` and auth provider `0.9.119-mfe`.
- Its source contains deployed strings including `trade_finance_sr_version` and `[Module Federation] Failed to init remotes:`.
- Its immediate successor `c2fdca7` migrated to Next/Rspack `16.2.9`, which is inconsistent with this deployed artifact.

No build SHA, federation build identifier, or `sourceMappingURL` is embedded, so `20b79a5` is strongly indicated but not cryptographically proven.

---

