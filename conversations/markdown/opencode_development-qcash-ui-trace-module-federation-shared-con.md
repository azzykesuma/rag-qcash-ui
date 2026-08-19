# Trace shared remote impact (@explore subagent)

- **ID**: `ses_0009aad75ffeISrCMVPqziOkDO`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 15:30:58`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

In ~/Development/qcash-ui, trace Module Federation/shared configuration and all local registry/config references relevant to the staged Auth/Ordent/BRI provider no-lazy-load change. Focus especially on remote applications/modules that consume shared auth providers, version constraints, singleton/eager behavior, customEntry handling, and whether eager imports can break host startup when a remote is missing or incompatible. Do not edit files. Return a map of impacted remotes and concrete review findings with exact file/line references. Desired thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

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
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **bash**: Tool: bash
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 12: Assistant

## Scope and staged delta

The staged change affects four files:

- `~/Development/qcash-ui/components/providers/BRIProviders.tsx`
- `~/Development/qcash-ui/components/providers/OrdentProviders.tsx`
- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx`
- `~/Development/qcash-ui/pages/_app.tsx`

It removes all `next/dynamic(..., { ssr: false })` boundaries around the BRI and Ordent provider implementations and statically imports `BriProviders` into `_app`.

Important distinction: this is “no lazy load” at the Next/JavaScript import level. It does **not** set Module Federation `eager: true`.

---

# Concrete review findings

## 1. High: the local installation/build is not testing the declared Ordent version

The repository declares and locks Ordent `0.9.142-mfe`:

- `~/Development/qcash-ui/package.json:27`
- `~/Development/qcash-ui/yarn.lock:534-542`

But the installed package identifies itself as `0.9.138-mfe`:

- `~/Development/qcash-ui/node_modules/@ordentco/addons-auth-provider/package.json:2-5`

The existing generated federation output also advertises `0.9.138-mfe`:

- `~/Development/qcash-ui/.next/server/mf-manifest.json:50-64`
- `~/Development/qcash-ui/.next/static/chunks/remoteEntry.js:5999-6000`
- `~/Development/qcash-ui/.next/static/chunks/remoteEntry.js:6059-6061`

Consequences:

- A local build/test against this workspace does not prove that the newly static imports work with the declared deployment version `0.9.142-mfe`.
- Export shape, initialization behavior, SSR safety, or context implementation differences between `0.9.138-mfe` and `0.9.142-mfe` remain unvalidated.
- The `.next` federation manifest is stale relative to source configuration and should not be treated as deployment truth.

BRI is consistent at `0.10.6-mfe`:

- `~/Development/qcash-ui/package.json:20`
- `~/Development/qcash-ui/yarn.lock:98-106`
- `~/Development/qcash-ui/node_modules/@bri/addons-auth-provider/package.json:2-5`

---

## 2. High: `requiredVersion: false` prevents version rejection, not API incompatibility

Both auth packages are configured as singleton shares with no required version:

- `~/Development/qcash-ui/next.config.js:15-21`

Specifically:

```js
"@bri/addons-auth-provider": { requiredVersion: false, singleton: true }
"@ordentco/addons-auth-provider": { requiredVersion: false, singleton: true }
```

Therefore:

- Only one runtime instance of each package should be selected in the default share scope.
- The host does not require a compatible semver range.
- `strictVersion` is not set and defaults to false.
- A differently versioned remote provider will generally not be rejected merely because its version differs.
- This avoids many hard “unsatisfied version” failures, but it permits an ABI/API-incompatible singleton to be selected.

The staged code assumes these exports exist immediately:

- BRI `AuthProvider` and `GlobalStoreProvider`:  
  `~/Development/qcash-ui/components/providers/BRIProviders.tsx:2-5`
- Ordent `AuthProvider` and `GlobalStoreProvider`:  
  `~/Development/qcash-ui/components/providers/BRIProviders.tsx:6-9`
- Standalone Ordent exports:  
  `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:2`

The bridge additionally assumes numerous hook values and mutation functions:

- Hooks: `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:44-46`
- Hook consumption: `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:249-252`
- BRI setter lookup: `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:322-326`

An incompatible singleton can therefore fail through missing exports, invalid hook/context identity, or missing setter methods even though federation reports no semver violation.

The generated client runtime confirms the effective share settings were historically:

```json
{
  "eager": false,
  "requiredVersion": false,
  "strictVersion": false,
  "singleton": true
}
```

at:

- `~/Development/qcash-ui/.next/static/chunks/webpack.js:6163-6174`

That generated output is stale, but it corroborates the source configuration’s intended behavior.

---

## 3. High: the staged test does not exercise federation startup, package resolution, or SSR/module evaluation

The new test completely mocks both auth packages:

- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:26-38`

It verifies only nesting and props:

- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:50-77`
- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:79-97`

It cannot detect:

- “Shared module is not available for eager consumption”
- share-scope initialization order
- a missing export in the real `0.9.142-mfe` package
- a remote selecting an incompatible singleton
- browser-global access during package evaluation
- behavior when `qui/providers` is unavailable
- server evaluation/build failures

There is also an existing Ordent provider test still built around mocking `next/dynamic`:

- Dynamic mock construction:  
  `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:8-62`
- `next/dynamic` mock:  
  `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:76-79`

It does **not** mock `@ordentco/addons-auth-provider`. After the staged implementation stops calling `next/dynamic`, its test components and spies at lines 36-55 are no longer connected to `OrdentProviders`. Assertions for those mock test IDs and spies at lines 90-130 are therefore stale and are likely to fail or cease testing the intended behavior.

---

## 4. Medium: the static import is not actually the first static auth-package dependency in `_app`

The staged `_app` change broadens the direct initial import graph:

- `~/Development/qcash-ui/pages/_app.tsx:19`
- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:2-10`

However, both auth packages were already reachable from `_app` through the statically imported `AuthBridgeSync`:

- `_app` import: `~/Development/qcash-ui/pages/_app.tsx:26`
- BRI and Ordent package imports:  
  `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:44-46`

Additional statically imported host components consume Ordent directly:

- `~/Development/qcash-ui/components/providers/OnboardingTourProvider.tsx:3-4`
- `~/Development/qcash-ui/components/onboarding-tour/notif-onboarding-tour.tsx:4`
- `~/Development/qcash-ui/components/ui/MFAErrorModal.tsx:3`

Thus the staged change does not fully transform the packages from “never in startup graph” to “in startup graph”; they were already present. Its clearest runtime change is removal of the provider-specific asynchronous render boundaries.

This reduces the likelihood that the change alone introduces a new federation eager-consumption error, but it does not eliminate that risk because the initial graph still synchronously consumes non-eager shared modules.

---

## 5. Medium: shared `eager` remains false, so synchronous source imports still rely on federation bootstrap handling

No auth share has `eager: true`:

- `~/Development/qcash-ui/next.config.js:16`
- `~/Development/qcash-ui/next.config.js:21`

Only lodash and moment explicitly write `eager: false`, at:

- `~/Development/qcash-ui/next.config.js:38-39`

For the auth packages, omitted `eager` also resolves to false. The generated runtime registers their local factories through asynchronous chunks:

- `~/Development/qcash-ui/.next/static/chunks/webpack.js:6110-6111`
- `~/Development/qcash-ui/.next/static/chunks/webpack.js:6163-6174`

Therefore a source-level static import does not make the shared implementation itself an eager federation provision. It creates a synchronous consumer of a share whose local factory may load asynchronously.

Webpack/Next Federation can normally bridge this when the entry module is asynchronous, and `_app` already has top-level federation initialization:

- `~/Development/qcash-ui/pages/_app.tsx:40-51`

Nevertheless, if the deployed build’s bootstrap/share-scope ordering differs, the runtime contains an explicit failure path for:

> Shared module is not available for eager consumption

at:

- `~/Development/qcash-ui/.next/static/chunks/webpack.js:5381`

A production build from a clean install is the meaningful validation; the unit tests do not cover this path.

---

## 6. Medium: removal of `ssr: false` has less SSR effect than it first appears, but changes module evaluation boundaries

The providers are now normal components:

- `~/Development/qcash-ui/components/providers/BRIProviders.tsx:18-34`
- `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:11-22`

But the entire provider subtree remains under `ModuleFederationProvider`, which is still dynamically loaded with `ssr: false`:

- Definition: `~/Development/qcash-ui/pages/_app.tsx:60-65`
- Tree root: `~/Development/qcash-ui/pages/_app.tsx:160-194`

As currently nested, React should not execute `BriProviders` during SSR because its `ssr:false` ancestor does not render the child subtree on the server.

However:

- The module containing `BriProviders` is now statically evaluated as part of `_app`.
- Its package imports are resolved on both build/server and client compilation paths.
- Future movement of `BriProviders` outside that outer client-only boundary would immediately expose the providers to SSR rendering.

The installed package implementations contain many browser-global operations, although most auth initialization paths appear guarded with `typeof window !== "undefined"`. For example:

- BRI token initialization guard:  
  `~/Development/qcash-ui/node_modules/@bri/addons-auth-provider/dist/src/auth.js:236-237`
- Ordent token initialization guard:  
  `~/Development/qcash-ui/node_modules/@ordentco/addons-auth-provider/dist/src/auth.js:232-233`

Because the installed Ordent package is not the declared version, this is not sufficient evidence for `0.9.142-mfe`.

---

## 7. Medium: local migration documentation and feature-toggle architecture are out of sync with actual `_app`

The migration documentation says `_app` chooses between `OrdentProviders` and `BriProviders` based on:

- `toggleAddonsAuthProviderMigration`
- `BRI_PROVIDER_ROUTE_PREFIXES`

References:

- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:10-34`
- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:148-154`

Actual `_app` now always renders `BriProviders`:

- `~/Development/qcash-ui/pages/_app.tsx:166-193`

No `_app` reference to `OrdentProviders`, `toggleAddonsAuthProviderMigration`, or `isBriProviderRoute` remains.

The environment toggle is still defined:

- `~/Development/qcash-ui/config.ts:472-478`

But its only codebase matches are configuration and documentation; it no longer controls provider mounting.

The route prefixes still affect which context `AuthBridgeSync` publishes, rather than whether both providers are mounted:

- Prefix list:  
  `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1`
- Prefix matching:  
  `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:13-14`
- Selection use:  
  `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:184-194`
- Runtime selection:  
  `~/Development/qcash-ui/components/providers/AuthBridgeSync.ts:314-320`

This documentation drift predates or extends beyond the staged change, but it is directly relevant to reviewing which remotes are expected to consume BRI versus Ordent state.

---

## 8. Medium: a missing feature remote does not supply the host auth providers, but `qui/providers` remains a startup-critical remote

The auth providers are local npm dependencies. No registry remote is used to fetch `AuthProvider` or `GlobalStoreProvider`.

Federation initialization builds remote descriptors from the registry:

- `~/Development/qcash-ui/services/federation/init.ts:37-50`

This generally registers names and entries; it does not eagerly import every remote module. Global-style failures are caught independently:

- `~/Development/qcash-ui/services/federation/init.ts:72-84`

Consequently, an unavailable route-specific remote should not prevent the local BRI/Ordent providers from loading.

However, `qui/providers` is loaded on every application render:

- Loader declaration: `~/Development/qcash-ui/pages/_app.tsx:66-68`
- It wraps `AuthBridgeSync` and all page content:  
  `~/Development/qcash-ui/pages/_app.tsx:166-193`

`qui` is a global custom-entry remote:

- `~/Development/qcash-ui/constants/features/registry.ts:49-52`

If that remote is missing, its failure occurs before `AuthBridgeSync`, onboarding providers, and the page can mount. The fallback plugin returns a generic error component:

- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:41-42`
- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:89-113`
- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:123-127`

The generic fallback does not render its `children`, so a missing `qui/providers` can effectively replace the full application subtree with an error view. That is an existing startup dependency, not introduced by the auth no-lazy-load change.

---

# Impacted remote map

## Tier 1: explicitly BRI-preferred migration remotes

### `qcash-ui-helpdesk-portal`

- BRI-preferred path: `/helpdesk`  
  `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1`
- Registry entry and route:  
  `~/Development/qcash-ui/constants/features/registry.ts:501-505`
- Standard entry resolution:  
  `${baseFederation}/helpdesk-portal-v2/_next/static/chunks/remoteEntry.js`
- Page consumption:
  - `~/Development/qcash-ui/pages/helpdesk/portal/index.tsx:10-13`
  - `~/Development/qcash-ui/pages/helpdesk/portal/tickets/index.tsx:11`
  - `~/Development/qcash-ui/pages/helpdesk/portal/tickets/view/[userID]/[id].tsx:11`
- Layout remote selection:
  - BRI header/footer when enabled:  
    `~/Development/qcash-ui/pages/helpdesk/portal/index.tsx:15-21`
- Relevant feature toggles:
  - Helpdesk enabled all environments:  
    `~/Development/qcash-ui/config.ts:458-464`
  - BRI header/footer disabled all environments:  
    `~/Development/qcash-ui/config.ts:479-485`

Review implication: this is the most direct migration consumer. It must tolerate both provider packages being present while using the bridge-selected BRI state. The local migration contract explicitly warns against direct remote `useAuth()` binding:

- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:120-130`
- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:191-193`

### `qcash-ui-plafon-pegadaian`

- BRI-preferred path: `/plafon-pegadaian`  
  `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1`
- Registry entry:  
  `~/Development/qcash-ui/constants/features/registry.ts:270-274`
- Standard entry resolution:  
  `${baseFederation}/plafon-pegadaian-v2/_next/static/chunks/remoteEntry.js`
- Remote loads:
  - `~/Development/qcash-ui/pages/plafon-pegadaian/index.tsx:9-17`
  - `~/Development/qcash-ui/pages/plafon-pegadaian/create/index.tsx:11`
  - `~/Development/qcash-ui/pages/plafon-pegadaian/edit/[id].tsx:11`
  - `~/Development/qcash-ui/pages/plafon-pegadaian/view/[id].tsx:11`
- Toggle enabled all environments:  
  `~/Development/qcash-ui/config.ts:161-167`

It also consumes `qcash-ui-header-footer/default`, making the auth share interaction at least a three-party scope: host, feature remote, and layout remote.

### `qcash-ui-sistem-management`

- BRI-preferred path: `/systems`  
  `~/Development/qcash-ui/components/providers/auth-provider-migration-routes.ts:1`
- Registry entry and custom entry:  
  `~/Development/qcash-ui/constants/features/registry.ts:122-127`
- Custom entry resolves to:  
  `${baseFederation}/system-v2/_next/static/chunks/remoteEntry.js`
- Remote loads:
  - `~/Development/qcash-ui/pages/systems/index.tsx:14-17`
  - `~/Development/qcash-ui/pages/systems/edit/[id].tsx:8`
  - `~/Development/qcash-ui/pages/systems/view/[id].tsx:8`
- It also loads `qcash-ui-header-footer/default`:  
  `~/Development/qcash-ui/pages/systems/index.tsx:8-11`

This custom entry is important: without it, standard name stripping would target `sistem-management-v2`, not `system-v2`.

---

## Tier 2: global/root remotes affected by provider availability and startup ordering

### `qui`

- Registry: `~/Development/qcash-ui/constants/features/registry.ts:49-52`
- Custom entry: `${baseFederation}/design-system/static/remoteEntry.js`
- Loaded as `qui/providers`: `~/Development/qcash-ui/pages/_app.tsx:66-68`
- Wraps auth bridge/page tree: `~/Development/qcash-ui/pages/_app.tsx:166-193`
- Special ETag handling designation:  
  `~/Development/qcash-ui/services/federation/init.ts:10`
- Development CSS proxy handling:  
  `~/Development/qcash-ui/next.config.js:120-143`

There is no local evidence that `qui/providers` imports either auth package, but it is startup-critical and runs inside the BRI/Ordent provider tree.

### `qcash-ui-header-footer`

- Global registry entry:  
  `~/Development/qcash-ui/constants/features/registry.ts:43-47`
- Custom entry: `${baseFederation}/header-footer/_next/static/chunks/remoteEntry.js`
- It is loaded as a layout by many pages, including all three explicitly migrated route families shown above.
- Legacy remotes are documented as expecting Ordent provider hooks:  
  `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:46-58`

This is a high-priority compatibility remote because it is reused across both legacy and BRI-preferred routes.

### `qcash-ui-header-footer-bri`

- Registry/custom entry:  
  `~/Development/qcash-ui/constants/features/registry.ts:532-537`
- Entry: `${baseFederation}/header-footer-bri/_next/static/chunks/remoteEntry.js`
- Used conditionally on helpdesk:  
  `~/Development/qcash-ui/pages/helpdesk/portal/index.tsx:15-21`
- Disabled in all configured environments:  
  `~/Development/qcash-ui/config.ts:479-485`

It is locally configured but currently dormant unless runtime/environment behavior overrides the checked-in toggle.

### `qcash-global-component`

- Global registry/custom entry:  
  `~/Development/qcash-ui/constants/features/registry.ts:38-42`
- Entry: `${baseFederation}/global-component/_next/static/chunks/remoteEntry.js`

No local static evidence identifies its auth provider dependency, but it is globally registered and therefore belongs in runtime share-scope verification.

---

## Tier 3: all route remotes are potential singleton auth consumers

All features are registered with federation at startup:

- Registry aggregation:  
  `~/Development/qcash-ui/services/federation/utils/registry.ts:11-13`
- Runtime registration:  
  `~/Development/qcash-ui/services/federation/init.ts:37-50`

The local registry does not declare which auth package each remote consumes. That information only becomes visible after remote containers join the share scope. The repository’s diagnostics derive it from `useIn` or `from`:

- Share-scope traversal:  
  `~/Development/qcash-ui/components/federation/monitor/auth/provider-versions.ts:8-25`
- Per-remote provider/version map:  
  `~/Development/qcash-ui/components/federation/monitor/auth/provider-versions.ts:32-49`
- Monitor explanation:  
  `~/Development/qcash-ui/components/federation/monitor/auth/index.tsx:239-259`

Therefore the three Tier 1 remotes are the only statically identifiable BRI-preferred consumers. Exact Ordent/BRI dependencies for every other remote require either their repositories or a live share-scope snapshot.

---

# `customEntry` handling map

`customEntry` takes precedence over standard name-based entry resolution:

- `~/Development/qcash-ui/services/federation/utils/registry.ts:87-100`

Standard entries use:

```text
${baseFederation}/${stripModuleName(name)}-v2/_next/static/chunks/remoteEntry.js
```

- `~/Development/qcash-ui/services/federation/utils/registry.ts:81-84`
- Prefix stripping: `~/Development/qcash-ui/services/federation/utils/name.ts:1-4`

Development port overrides preserve the custom-entry path:

- `~/Development/qcash-ui/services/federation/utils/module-entry.ts:5-25`

Custom-entry remotes relevant to runtime/share-scope review are:

- Global component: registry `41`
- Header/footer: `46`
- `qui`: `51`
- Notification: `113`
- System management: `123`
- Partner integration: `130`
- Company: `313`
- Workflows: `366`
- History: `527`
- Header/footer BRI: `534`
- CMS: `588`
- Landing page: `594`
- Amend/cancel LC: `620`
- Ecosystem/BGN management: `669`
- Foreign exchange: `680`
- Dashboard auto collection: `688`
- Giro: `697`
- Dashboard BGN: `705`
- Loan application: `717`
- Loan payment: `739`
- Loan portfolio: `751`
- Auto-collection registration: `762`

All references are in:

`~/Development/qcash-ui/constants/features/registry.ts`

Changing from standard to custom entry changes which deployed container contributes versions to the shared scope. It does not change the host’s auth-provider imports or sharing rules.

---

# Missing/incompatible remote startup assessment

## Missing route-specific remote

Normally does **not** prevent host auth provider startup:

1. The providers come from local npm packages.
2. `init()` registers all remote descriptors.
3. Route components invoke `loadRemote()` later.
4. The fallback plugin can substitute an unavailable route component.

Relevant references:

- `~/Development/qcash-ui/services/federation/init.ts:37-66`
- `~/Development/qcash-ui/services/federation/fallbackPlugin.tsx:123-127`

## Missing global `qui/providers`

Can prevent the effective application subtree from mounting because it wraps `AuthBridgeSync` and all page content:

- `~/Development/qcash-ui/pages/_app.tsx:66-68`
- `~/Development/qcash-ui/pages/_app.tsx:166-193`

## Incompatible remote auth version

Host `requiredVersion: false` means a differing remote version will not be rejected on host semver grounds. Singleton selection can still expose incompatible APIs. The risk is highest when:

- a remote initializes the share before the host provider is consumed;
- the remote declares its own stricter singleton requirement;
- the selected implementation lacks the provider exports or bridge-required setters;
- the remote bundles its own non-shared copy and binds to a different React context.

The migration documentation explicitly identifies direct remote provider-hook binding as a failure mode:

- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:191-193`

## Bottom line

The staged static imports do not directly depend on a remote being present. The more concrete risks are:

1. unvalidated Ordent `0.9.142-mfe` behavior because the workspace has `0.9.138-mfe`;
2. non-eager shared modules being consumed from the initial graph;
3. no semantic compatibility protection due to `requiredVersion: false`;
4. stale tests that do not exercise federation or real package initialization;
5. `qui/providers` remaining an independent startup-critical remote;
6. all remotes now consistently executing beneath both provider contexts, regardless of the obsolete provider-tree toggle documentation.

---

