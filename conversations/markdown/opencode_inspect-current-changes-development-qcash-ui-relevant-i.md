# Compare auth navigation timing (@explore subagent)

- **ID**: `ses_ffefda476ffeH4RwuNJ02wZZD6`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 23:02:06`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Inspect current changes in ~/Development/qcash-ui and relevant issuance-lc host pages/registry entries. Analyze whether changing BRI/Ordent auth providers from next/dynamic ssr:false to static imports changes router identity, navigation methods, provider timing, or only auth initialization timing. Identify any issuance-lc navigation failures this can and cannot fix, with exact file/line references. Do not edit. Desired thoroughness: medium.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **bash**: Tool: bash
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

## Scope

The working tree is clean. The relevant change is commit `8e7040a6`, one commit ahead of `origin/dev`, which:

- replaces lazy `next/dynamic({ ssr: false })` auth-provider imports with static imports;
- statically imports `BriProviders` in `_app`;
- separately changes logout handling in `AuthBridgeSync`.

## Router identity and navigation behavior

The import conversion does **not** change router identity or navigation methods.

- `BriProviders` still calls the same Pages Router `useRouter()` and passes that exact object to both auth providers:
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:19`
  - BRI router prop: `:22`
  - Ordent router prop: `:26`
- `OrdentProviders` does the same:
  - `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:12`
  - router prop: `:16`
- The `as AuthRouter` casts are compile-time-only and cannot wrap, clone, bind, or replace the router.
- The added test explicitly verifies reference equality with `mockRouter`:
  - `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:66-75`

Provider configuration and nesting are also preserved:

- BRI Auth → BRI Store → Ordent Auth → Ordent Store:
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:22-33`
- `baseRoute` and `loginRoute` are unchanged:
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:27-28`
  - `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:17-18`
- The auth router contract still exposes `push`, optional `replace`, route fields, and events:
  - `~/Development/qcash-ui/node_modules/@ordentco/addons-auth-provider/src/types/router.ts:1-9`

Therefore, this does not change `push` versus `replace`, destination construction, browser history semantics, router event behavior, or dynamic-route parameter handling.

One nuance: the BRI package itself patches singleton Next router `push` under relogin conditions, but static importing does not introduce a different router. It can only cause package/module initialization to occur earlier:

- `~/Development/qcash-ui/node_modules/@bri/addons-auth-provider/dist/src/global-custom-axios/interceptors/auth-interceptors.js:341-378`

## What timing changes

The meaningful effect is earlier and more deterministic auth initialization:

- `_app` now statically imports `BriProviders`:
  - `~/Development/qcash-ui/pages/_app.tsx:19`
- It mounts the provider around the application:
  - `~/Development/qcash-ui/pages/_app.tsx:166-193`
- BRI and Ordent provider exports are now immediately available instead of waiting for separate dynamic chunks:
  - `~/Development/qcash-ui/components/providers/BRIProviders.tsx:2-9`
  - `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:2`

This removes the old asynchronous BRI wrapper and nested Ordent wrapper boundaries. It can advance provider mount, token hydration, guard execution, and context availability.

It does **not** make auth server-rendered in the effective application tree. `ModuleFederationProvider` remains an outer `ssr:false` boundary:

- `~/Development/qcash-ui/pages/_app.tsx:60-65`
- `~/Development/qcash-ui/pages/_app.tsx:161`

Thus, provider package modules are imported/evaluated earlier, but the provider components still mount in the client-only subtree.

## Issuance-LC host routing

The host-side issuance pages are correctly mapped:

- List:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:15-18`
- Create:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/create/index.tsx:15-18`
- Detail:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/detail/[transactionEventId].tsx:15-18`
- Edit:
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx:15-18`

The registry has the expected feature and base route:

- `~/Development/qcash-ui/constants/features/registry.ts:519-523`

Although only the base path is listed, registry matching defaults to prefix matching, so `/create`, `/detail/...`, and `/edit/...` are covered:

- default prefix type: `~/Development/qcash-ui/services/federation/utils/registry.ts:102-120`
- route filtering: `~/Development/qcash-ui/services/federation/utils/registry.ts:173-183`

Issuance-LC is not a BRI-preferred auth route. Only helpdesk, plafon-pegadaian, and systems use BRI preference:

- `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1`
- `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:184-194`

Consequently, issuance normally consumes the Ordent auth state/bridge once initialized.

## Issuance-LC failures this can fix

The static-import portion can plausibly fix:

1. **Initial deep-link/startup races** where `/import-lc/issuance-lc...` is opened before lazy auth chunks have mounted and initialized.
2. **Lazy auth chunk loading failures or delays** that previously prevented the entire nested application subtree from appearing.
3. **Early guard/token/context availability races**, especially if the issuance remote reads auth immediately on first mount.

The issuance remote remains dynamically loaded client-side:

- `~/Development/qcash-ui/components/federation/remote-dynamic.tsx:13-24`
- load begins in an effect: `:68-76`

Therefore, static auth imports improve ordering and latency, but do not guarantee that asynchronous authority fetching has completed before issuance logic runs.

## Failures it cannot fix

Static auth imports cannot fix:

- Incorrect `router.push`/`replace` calls or malformed destinations inside the issuance remote.
- Back/history-stack behavior.
- Missing or incorrect `transactionEventId`.
- A remote using a different router instance or incompatible `next/router` copy.
- Remote exposure/load failures for `qcash-ui-issuance-lc/...`.
- Disabled `toggleIssuanceLc`; both registry loading and host rendering remain toggle-gated:
  - `~/Development/qcash-ui/constants/features/registry.ts:521-522`
  - `~/Development/qcash-ui/pages/import-lc/issuance-lc/index.tsx:24-26`
- Incorrect registry paths; none are apparent here.
- Navigation failures occurring after providers are already mounted during ordinary in-app issuance transitions.
- Failures caused by the issuance MFE’s own navigation code—the host pages contain no navigation calls, only remote declarations.

## Separate logout fix in the same commit

A logout/relogin issuance failure involving stale identity or authorities may be fixed by the **separate `AuthBridgeSync` changes**, not by static imports:

- logout event clears token and bridge:
  - `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:302-313`
- absent token now always prevents stale provider state from being exposed:
  - `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:407-414`

That can address “navigate to issuance after logout/user switch and receive stale access-denied or previous-user state.” It still does not alter router identity or navigation methods.

**Conclusion:** the static-import change primarily removes auth loading boundaries and advances initialization. It preserves the exact Next router object, provider configuration, provider context packages, and navigation APIs. Any fix to issuance navigation is limited to startup/auth-readiness races or lazy-chunk failures; route construction, history behavior, registry, remote loading, and remote-internal navigation defects remain unaffected.

---

