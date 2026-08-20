# Audit repository scaffolding (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_081be808fffeR9ZUDmAdlVijND`
- **Source Tool**: `opencode`
- **Date**: `2026-07-20 13:40:49`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a VERY THOROUGH over-engineering audit of D:\work\qcash-ui-dashboard-auto-collection outside src: root configs, scripts, public/locales, styles, tests, package.json dependencies, Docker/build files, README, hidden project config where relevant. This is NOT correctness/security/performance review and do not edit. Exclude .git, node_modules, .next, .scannerwork, generated @mf-types, IDE metadata, generated style bundles, and user design/reference artifacts. Hunt dead scripts/config/dependencies, speculative abstractions, wrappers, duplicated tooling, hand-rolled standard/native behavior. Verify dependency usage repo-wide before claiming removable. Return ranked findings with exact paths and estimated line/dependency cuts, tags delete/stdlib/native/yagni/shrink. Also state areas checked with no finding.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (20)</summary>

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

<details><summary>🔧 Tool Invocations (17)</summary>

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
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

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

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

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

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

# Over-engineering audit

No files were edited. I searched references repo-wide, checked tracked/ignored status, inspected all root tooling, scripts, public assets/locales, tests, Docker files, and used `yarn why` for questionable dependencies.

## Ranked findings

### 1. Legacy `public/` dump is entirely unused
**Tags:** `delete`, `yagni`  
**Confidence:** High

Only `favicon.ico` and four locale files are tracked. Source imports exactly:

- `public/locales/{en,id}/dashboard-auto-collection.json`
- `public/locales/{en,id}/common.json`

No references were found to the following ignored files:

- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\en\translation.json` — 1,952 lines
- `D:\work\qcash-ui-dashboard-auto-collection\public\locales\id\translation.json` — 2,368 lines
- `D:\work\qcash-ui-dashboard-auto-collection\public\bricams.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\cbm.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\qlola.png`
- `D:\work\qcash-ui-dashboard-auto-collection\public\vercel.svg`
- `D:\work\qcash-ui-dashboard-auto-collection\public\files\`
- `D:\work\qcash-ui-dashboard-auto-collection\public\flags\`
- `D:\work\qcash-ui-dashboard-auto-collection\public\fonts\`
- `D:\work\qcash-ui-dashboard-auto-collection\public\icons\`
- `D:\work\qcash-ui-dashboard-auto-collection\public\images\`
- `D:\work\qcash-ui-dashboard-auto-collection\public\img\`

The asset directories contain approximately 113 files copied from a much larger application.

**Estimated cut:** 4,320 JSON lines plus 113 binary/SVG/document files; 115 files total.

---

### 2. Three abandoned scripts/artifacts implement work the project does not perform
**Tags:** `delete`, `native`, `yagni`  
**Confidence:** High

- `D:\work\qcash-ui-dashboard-auto-collection\scripts\analyze-bundle.js` — 37 lines  
  No `analyze-bundle` package script, no analyzer dependency, and no `ANALYZE` handling in `next.config.ts`. Its Moment/date-fns advice also targets dependencies absent from the repository.
- `D:\work\qcash-ui-dashboard-auto-collection\scripts\update-version-file.js` — 9 lines  
  Not called anywhere; the actual write is commented out.
- `D:\work\qcash-ui-dashboard-auto-collection\version` — 1 line  
  Disagrees with `package.json` and is read nowhere. `package.json` is already the native package-version source.
- `D:\work\qcash-ui-dashboard-auto-collection\todo.md` — 1 line  
  Requests migration of nonexistent “momment” usage to a nonexistent date utility.

All are explicitly ignored by `.gitignore`, confirming they are known leftovers rather than active project files.

**Estimated cut:** 48 lines, four files.

---

### 3. ESLint and Prettier configs duplicate the active Biome toolchain
**Tags:** `delete`, `shrink`  
**Confidence:** High

- `D:\work\qcash-ui-dashboard-auto-collection\.eslintrc.json` — 3 lines
- `D:\work\qcash-ui-dashboard-auto-collection\.prettierrc` — 14 lines

There are no ESLint or Prettier dependencies or scripts. `package.json` exclusively invokes Biome, and both files are ignored as legacy leftovers.

**Estimated cut:** 17 lines, two files.

---

### 4. Stale Husky runtime exists without Husky or any hook
**Tags:** `delete`, `yagni`  
**Confidence:** High

`D:\work\qcash-ui-dashboard-auto-collection\.husky\_` contains 17 generated files, but:

- No `husky` dependency
- No `prepare` script
- No project-level hook such as `.husky/pre-commit`
- `.husky` is explicitly ignored

Only Husky’s generated dispatcher wrappers remain.

**Estimated cut:** 17 files and approximately 59 shell-script lines.

---

### 5. React Compiler is an optional tooling layer with two dedicated dependencies
**Tags:** `delete`, `yagni`  
**Confidence:** Medium; conditional on the compiler not being an explicit product requirement

Configuration:

- `D:\work\qcash-ui-dashboard-auto-collection\next.config.ts:9-11`

Dedicated packages:

- `D:\work\qcash-ui-dashboard-auto-collection\package.json:26` — `react-compiler-runtime`
- `D:\work\qcash-ui-dashboard-auto-collection\package.json:39` — `babel-plugin-react-compiler`

`yarn why` shows both are present solely because the root package declares them. There is no documentation, benchmark, or code dependency requiring compiler transforms. The plugin accounts for roughly 6.9 MB installed with its unique/transitive dependencies.

If compiler behavior is intentional for the React 18 federation contract, retain it.

**Estimated cut:** two direct dependencies, five source/config lines, and roughly 9–15 lockfile lines.

---

### 6. `.gitignore` acts as a registry of files that should be deleted
**Tags:** `shrink`, `delete`  
**Confidence:** High after findings 1–3 are applied

`D:\work\qcash-ui-dashboard-auto-collection\.gitignore:49-61` individually suppresses:

- Dead root styles, todo, and version files
- Dead ESLint/Prettier configs
- The legacy public tree and translation dumps
- Two dead scripts

Deleting the underlying debris makes this whole exception-heavy block unnecessary. The generated style ignores at lines 62–63 should remain.

**Estimated cut:** 13 lines.

---

### 7. Second, unused global stylesheet duplicates Tailwind’s active reset
**Tags:** `delete`, `native`  
**Confidence:** High

- `D:\work\qcash-ui-dashboard-auto-collection\styles\globals.css` — 10 lines

The app imports `D:\work\qcash-ui-dashboard-auto-collection\src\styles\globals.css`, not this file. Its box-sizing/margin reset and sans-serif declaration are already covered by the active Tailwind base/theme output. The entire root `styles` directory is ignored.

**Estimated cut:** 10 lines, one file.

---

### 8. Docker/local registry configuration contains an unused package scope and dead stage cleanup
**Tags:** `delete`, `shrink`  
**Confidence:** High

- `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:34` configures `@bri:registry`.
- `D:\work\qcash-ui-dashboard-auto-collection\.npmrc:1` repeats the `@bri` scope.

No `@bri/*` package exists in `package.json`, `yarn.lock`, or source. The `@ordentco` registry remains necessary.

Additionally, `D:\work\qcash-ui-dashboard-auto-collection\Dockerfile:69-72` clears proxy variables immediately before the build stage ends. Nothing executes or inherits from that stage afterward.

**Estimated cut:** six configuration lines.

---

### 9. Ignored local `.env` carries six unrelated legacy settings
**Tags:** `delete`, `shrink`  
**Confidence:** Medium because it is user-local configuration

`D:\work\qcash-ui-dashboard-auto-collection\.env:1,10-14` has no static references in source, scripts, Next config, Docker, or package scripts. The dashboard uses the environment-specific API variables instead.

The framework-specific variable on line 15 was not classified as dead because Next/Rspack may consume it without a source reference.

**Estimated cut:** six local configuration lines; no tracked-project cut.

---

### 10. `allowJs` is inert under the TypeScript include set
**Tags:** `delete`, `shrink`  
**Confidence:** High

- `D:\work\qcash-ui-dashboard-auto-collection\tsconfig.json:5`

`allowJs` is enabled, but `include` contains only TypeScript/MTS patterns. The JavaScript/MJS scripts are not included in the TypeScript project.

The explicit `.next` and `next-env` include entries are also covered by `**/*.ts`, but retaining them as framework documentation is reasonable.

**Estimated cut:** one line.

---

### 11. Sonar coverage property has no repository-side producer
**Tags:** `delete`, `yagni`  
**Confidence:** Medium because an external CI pipeline could supply coverage

- `D:\work\qcash-ui-dashboard-auto-collection\sonar-project.properties:4`

The repository has no coverage script, Vitest coverage provider, Vitest config, or other reference to `coverage/lcov.info`. Remove this property unless an external pipeline generates that exact report.

The remaining Sonar source/test settings correspond to the two existing tests.

**Estimated cut:** one line.

---

### 12. README retains generic Next.js scaffold material
**Tags:** `delete`, `yagni`, `shrink`  
**Confidence:** High, low priority

- `D:\work\qcash-ui-dashboard-auto-collection\README.md:19-26`

The generic “Learn More” section and framework links do not document this micro frontend. The setup and local URL sections are useful and should remain.

**Estimated cut:** eight lines.

## Dependency verification

No unquestionably dead dependency was found.

Verified as used or contractually required:

- `@module-federation/enhanced` — runtime remote loading
- `@ordentco/addons-auth-provider` — app and dashboard auth
- `i18next`, `react-i18next`, `ni18n` — local and federation i18n sharing
- `next`, `next-rspack`, `react`, `react-dom` — application/runtime
- `recharts` — the area chart
- `react-is` — required peer dependency of Recharts, confirmed in Recharts metadata
- `@biomejs/biome` — lint/format scripts
- `tailwindcss`, `@tailwindcss/cli`, `@tailwindcss/postcss` — application CSS plus federated style generation
- `typescript` and React/Node type packages — TypeScript build
- `vitest` — two current test files

Only the two React Compiler packages are conditional removal candidates.

## Checked with no finding

- Active `package.json` lifecycle scripts: style generation, development, build, start, lint, format, and tests are connected.
- `postcss.config.mjs`: required for the stylesheet imported directly by the Next app.
- `scripts/build-globals-expose.mjs`: custom, but it serves the explicit Module Federation stylesheet exposure; not safe to replace without validating host CSS loading.
- `next.config.ts`: both `qui` and `qcash-global-component` remotes are referenced dynamically.
- Docker standalone runner copies match `output: "standalone"`.
- Active dashboard and common locale files are small. `common` has no direct static key references, but dynamically loaded global components can consume its namespace, so it was not marked removable.
- Existing Vitest tests are small and require no extra test framework abstraction.
- `.editorconfig`, `.yarnrc`, the `@ordentco` npm scope, and Docker Nexus authentication are justified.
- `public/favicon.ico` uses the browser’s native favicon convention.
- Biome configuration has several defensive ignores/rule overrides, but none offered a high-confidence simplification.
- Generated `next-env.d.ts`, `tsconfig.tsbuildinfo`, generated style bundles, and `@mf-types` were not treated as authored over-engineering.

Excluded as requested: Git internals, dependencies/build outputs, scanner output, generated federation types/styles, IDE/AI metadata, and design/reference artifacts.

---

