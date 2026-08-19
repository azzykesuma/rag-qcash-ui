# Audit dependencies tooling (@explore subagent)

- **ID**: `ses_090c08885ffe6qBnahTJbNZJlw`
- **Source Tool**: `opencode`
- **Date**: `2026-07-17 15:44:18`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a very thorough repository-wide audit of D:\work\qcash-ui-dashboard-auto-collection focused only on removable dependencies and over-engineered tooling/configuration. Scan package.json/yarn.lock imports across the whole tree, Next/Module Federation/i18n/auth/config/build scripts, Docker, configs, generated artifacts, and documentation. Do not report correctness bugs. Identify dependencies with no runtime/source use, redundant tooling or generated tracked files, one-use wrappers/config layers, and native/stdlib replacements. Do not edit. Return ranked candidate findings with exact paths and conservative estimated lines/dependencies removable; verify actual imports/usages before reporting.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **bash**: Tool: bash
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

<details><summary>🔧 Tool Invocations (10)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
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
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

# Repository audit: removable dependencies and tooling

Audited the current working tree, including tracked, modified, untracked, and ignored files. No files were edited. Findings exclude correctness defects.

## Ranked candidates

### 1. Remove unused TanStack Query/status-query scaffolding
**Confidence: High**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json`
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\services\dashboard.test.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\api\dashboard-status.ts`

**Verification**
- `QueryClientProvider` wraps the dashboard, but no production component calls a TanStack query hook.
- `dashboardStatusQuery` is referenced only by its test.
- The status API route exists solely for that otherwise-unused query.
- `yarn why` confirms `@tanstack/react-query` is present only because it is a direct dependency.

**Conservative removal**
- 1 direct dependency: `@tanstack/react-query`
- 1 transitive package: `@tanstack/query-core`
- Approximately 35–42 source/config lines
- Installed footprint reported by Yarn: approximately 4.69 MB including unique dependencies

Keep it only when the dashboard begins fetching/cache-managing actual data.

---

### 2. Remove inactive Husky/lint-staged setup
**Confidence: High**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json`
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock`
- `D:\work\qcash-ui-dashboard-auto-collection\.gitignore`
- `D:\work\qcash-ui-dashboard-auto-collection\.husky\` — ignored workspace output

**Verification**
- No tracked Husky hooks exist.
- `.husky` is ignored wholesale.
- The only files under `.husky` are Husky-generated internal shim files.
- `lint-staged` has no invocation except its manual package script; no hook calls it.
- `yarn why` reports both packages exist solely as direct dev dependencies.

**Conservative removal**
- 2 direct dev dependencies: `husky`, `lint-staged`
- Approximately 7–9 manifest/config lines
- Ignored generated `.husky\_` directory
- Yarn reports approximately 3.92 MB for `lint-staged` and 28 KB for Husky, before accounting for shared packages

Add them back only with a tracked, enforced hook.

---

### 3. Remove completely unused `react-is`
**Confidence: High**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:32`
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock:2198`

**Verification**
- No source, config, script, or documentation import/reference exists.
- The lockfile contains no other dependency reference to this installation.
- `yarn why react-is` says it exists only because it is directly declared.

**Conservative removal**
- 1 direct runtime dependency
- 1 manifest line plus its generated lock stanza
- Approximately 32 KB installed

---

### 4. Stop tracking generated CSS and generated TypeScript-with-embedded-CSS
**Confidence: High**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.bundle.css`
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\global.exposes.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\build-globals-expose.mjs`
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:6-17`
- `D:\work\qcash-ui-dashboard-auto-collection\.gitignore`

**Verification**
- `predev` and `prebuild` always run the style generator.
- `globals.bundle.css` is generated by Tailwind CLI.
- `global.exposes.ts` is then generated from that bundle and embeds the same minified CSS as a string.
- Both outputs are tracked despite deterministic regeneration.

**Conservative removal**
- Approximately 17.9 KB of duplicated tracked generated output in the checked-in versions
- 26 physical lines, including two very large generated lines
- 0 dependencies required to untrack them
- Add two ignore entries; retain the generator and build scripts

This removes repository churn, not runtime output.

---

### 5. Remove legacy `ni18n` layer, subject to remote-contract confirmation
**Confidence: Medium**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:27`
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock:2067`
- `D:\work\qcash-ui-dashboard-auto-collection\ni18n.config.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:42`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:5`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:10`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:45-49`

**Verification**
- Application i18n uses `i18next` and `react-i18next` directly.
- `ni18n.config.ts` is never imported.
- `ni18n` is used only to manufacture a Module Federation shared entry.
- `yarn why ni18n` reports only the direct declaration as its reason for installation.

**Conservative removal**
- 1 direct runtime dependency
- Approximately 13–16 source/config lines
- Entire 5-line `ni18n.config.ts`
- Yarn reports approximately 1.89 MB with transitive dependencies

**Caveat:** confirm that `qui` and `qcash-global-component` do not require the host to supply `ni18n` as a singleton. This is the only reason not to remove it immediately.

---

### 6. Remove package-manager fields unsupported by the documented Yarn 1 workflow
**Confidence: Medium-high**

**Path**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:53-60`

**Verification**
- The repository uses Yarn Classic (`yarn.lock`; installed Yarn reports 1.22.22).
- `trustedDependencies` is not a Yarn Classic installation control.
- The package-level `ignoreScripts` array is likewise not the Yarn Classic `--ignore-scripts`/configuration mechanism.
- The same packages appear in both arrays, but neither block drives the documented install/build commands.

**Conservative removal**
- 8 configuration lines
- 0 dependencies

Retain only if a separate Bun-compatible corporate build consumer demonstrably reads these fields.

---

### 7. Collapse the unused injected-router context/provider layer
**Confidence: Medium**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\src\hooks\useRouter.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\dashboard\index.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx`

**Verification**
- `AppRouterProvider` is instantiated only around `DashboardPage`.
- No descendant calls `useAppRouter`.
- The only `useAppRouter` call is in `_app.tsx`, above that provider.
- Therefore the context/injected-router branch has no current consumer.
- The fallback router remains useful for `AuthProvider`; the context/provider portion does not.

**Conservative removal**
- Approximately 18–25 lines
- The dashboard `router` prop and provider wrapper
- 0 dependencies

**Caveat:** the exposed dashboard’s optional `router` prop may be treated as an external Module Federation API even though it currently has no effect internally.

---

### 8. Consolidate duplicate Module Federation loaders
**Confidence: High, low impact**

**Path**
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\federation.ts:69-81`
- Callers in:
  - `D:\work\qcash-ui-dashboard-auto-collection\src\components\ui\index.tsx`
  - `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx`

**Verification**
- `loadQuiRemote` and `loadGlobalComponentRemote` have identical implementations.
- Their only difference is the module string supplied by callers.

**Conservative removal**
- Approximately 6–7 net lines by exposing one `loadRemote`
- 0 dependencies

---

### 9. Collapse the one-use environment/config chain
**Confidence: Medium**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\src\config.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\utils\environment.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx`

**Verification**
- `config.ts` has one consumer.
- `getBaseUrlByEnv` has one caller: `config.ts`.
- The wrapper exists only to pass nine environment values positionally into another one-use function.

**Conservative removal**
- One entire wrapper file
- Approximately 10–18 net lines after replacing the positional chain with a direct hostname-to-environment lookup
- 0 dependencies

The environment selection itself is required; only the two-layer/9-argument structure is removable.

---

### 10. Remove the one-use translation hook wrapper
**Confidence: Medium**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\src\hooks\useDashboardAutoCollectionTranslation.ts`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\index.tsx`

**Verification**
- The hook has exactly one caller.
- Only one application translation call exists: the dashboard title.
- Resource installation is already performed during i18n initialization and again by `I18nProvider`.

**Conservative removal**
- Approximately 5–8 net lines after inlining the namespace call
- 1 file
- 0 dependencies

Keep the wrapper only if more module components are expected to use it soon.

---

### 11. Optionally remove the PostCSS Tailwind path and consume the precompiled bundle
**Confidence: Medium-low**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\postcss.config.mjs`
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:38`
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_app.tsx:13`
- `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.bundle.css`

**Verification**
- Tailwind is currently compiled twice through separate paths:
  1. Tailwind CLI generates `globals.bundle.css`.
  2. Next/PostCSS processes `globals.css` for the local app.
- The CLI path is already mandatory for the Module Federation globals expose.

**Conservative removal**
- 1 direct dev dependency: `@tailwindcss/postcss`
- Entire 7-line `postcss.config.mjs`
- 1 manifest line
- Yarn reports approximately 8.18 MB including shared/transitive packages; actual lockfile reduction will be smaller because the CLI shares many packages.

**Trade-off:** importing the precompiled bundle locally removes Tailwind-aware development rebuild/HMR unless the CLI is put into watch mode. Therefore this is not an immediate recommendation.

---

### 12. Delete the custom document if its single body utility is unnecessary
**Confidence: Low-medium**

**Path**
- `D:\work\qcash-ui-dashboard-auto-collection\src\pages\_document.tsx`

**Verification**
- It otherwise reproduces the standard Next.js document structure.
- Its only project-specific behavior is `lang="en"` and `fpl:antialiased` on `<body>`.

**Conservative removal**
- Approximately 10–13 lines
- 1 file
- 0 dependencies

Use Next’s native default document only if the body class and fixed language can be dropped or expressed in global CSS.

---

### 13. Replace Recharts with native SVG only if bundle size is a measured priority
**Confidence: Low; conditional**

**Paths**
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:33`
- `D:\work\qcash-ui-dashboard-auto-collection\src\containers\dashboard\components\trend-chart.tsx`
- `D:\work\qcash-ui-dashboard-auto-collection\yarn.lock`

**Verification**
- Recharts is genuinely used for one area chart.
- `yarn why recharts` reports approximately 38.98 MB installed with transitive dependencies.
- The chart uses responsive sizing, axes, grid, gradient, interpolation, active dots, and tooltips.

**Potential removal**
- 1 direct dependency and its charting dependency tree
- Likely no source-line reduction: a native responsive SVG implementation would add roughly 10–40 net lines and require maintaining interaction/accessibility behavior.

This is not dead weight; replace only after production bundle measurement justifies owning the chart implementation.

---

## Ignored/untracked workspace debris

These do not affect Git repository size because they are already ignored, but they are removable from the local workspace.

### Clearly stale text/tooling files
- `D:\work\qcash-ui-dashboard-auto-collection\.eslintrc.json` — 3 lines; Biome is used and ESLint is not installed.
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\analyze-bundle.js` — 37 lines; no package script, no analyzer dependency, no matching Next analyzer configuration.
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\update-version-file.js` — 9 lines; no caller and its write operation is commented out.
- `D:\work\qcash-ui-dashboard-auto-collection\todo.md` — 1 line referring to nonexistent Moment usage.
- `D:\work\qcash-ui-dashboard-auto-collection\styles\globals.css` — 10-line legacy stylesheet; no import.

### Large ignored translation leftovers
- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\en\translation.json` — 1,952 lines
- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\translation.json` — 2,368 lines

Neither is imported. Current i18n imports only:
- `dashboard-auto-collection.json`
- `common.json`

**Workspace-only cleanup:** approximately 4,380 text lines across the stale files above.

### Ignored binary/static asset tree
Numerous ignored files exist below `D:\work\qcash-ui-dashboard-auto-collection\public\` with no repository source references. Because external remote components can construct asset URLs dynamically, they should be deleted only if this workspace is not being used as a shared asset cache.

---

## Documentation/config cleanup

- `D:\work\qcash-ui-dashboard-auto-collection\README.md:19-26` is stock Next.js “Learn More” boilerplate: approximately 8 removable documentation lines.
- `D:\work\qcash-ui-dashboard-auto-collection\sonar-project.properties` has no in-repository invocation, but Sonar conventionally discovers this file externally. Do not remove without checking CI configuration.

## Dependencies verified as actually used or implicitly required

Not removal candidates based on this audit:

- `@module-federation/enhanced` — imported by federation runtime.
- `@ordentco/addons-auth-provider` — used by `_app`, dashboard auth, and federation sharing.
- `i18next`, `react-i18next` — used by provider/hooks and remote component resources.
- `recharts` — used by the trend chart.
- `vitest` — used by `src/services/dashboard.test.ts`.
- `next-rspack` — imported by `next.config.ts`.
- Tailwind CLI/Tailwind — required by the globals exposure generator.
- `react-compiler-runtime` and `babel-plugin-react-compiler` — implicit support for the configured React Compiler target 18.
- Biome and TypeScript/type packages — active scripts/configuration.

The strongest low-risk sequence is: remove TanStack scaffolding, remove Husky/lint-staged, remove `react-is`, and stop tracking generated style outputs.

---

