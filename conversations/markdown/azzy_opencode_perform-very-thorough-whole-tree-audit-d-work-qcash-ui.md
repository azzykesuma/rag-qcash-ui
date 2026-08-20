# Audit config dependencies (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_09b0d6385ffeeZKDq8CmTJlXky`
- **Source Tool**: `opencode`
- **Date**: `2026-07-15 15:44:09`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a very thorough whole-tree audit of D:\work\qcash-ui-pelindo focused ONLY on over-engineering in package dependencies, root configs, scripts, public assets/locales, Docker, opencode files, and non-src directories. Inspect package.json and verify dependency imports/usages across the current filesystem before labeling deps removable. Identify dead scripts/assets/config, native/stdlib replacements, speculative layers, and shrink opportunities. Return ranked findings tagged delete/stdlib/native/yagni/shrink with exact paths and defensible approximate removable lines/dependency counts. Do not edit. Ignore correctness issues and git-deleted historical files; audit what currently exists.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **bash**: Tool: bash
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (15)</summary>

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

<details><summary>🔧 Tool Invocations (16)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **glob**: Tool: glob
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

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

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
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (11)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

## Ranked findings

1. **[delete] Stale generated Module Federation type corpus**
   - `D:\work\qcash-ui-pelindo\@mf-types\qcash-global-component\`
   - **529 files, ~13,025 lines removable.**
   - Current source references only `@mf-types/qui/apis`; nothing references `qcash-global-component`. This corpus is also explicitly included by `tsconfig.json`, increasing TypeScript scan surface.
   - Follow-up shrink in `D:\work\qcash-ui-pelindo\tsconfig.json`: remove the broad `"*": ["./@mf-types/*"]` path and obsolete explicit include if these types are not regenerated.

2. **[delete] Legacy public asset payload**
   - `D:\work\qcash-ui-pelindo\public\files\`
   - `D:\work\qcash-ui-pelindo\public\fonts\`
   - `D:\work\qcash-ui-pelindo\public\flags\`
   - `D:\work\qcash-ui-pelindo\public\icons\`
   - `D:\work\qcash-ui-pelindo\public\images\`
   - `D:\work\qcash-ui-pelindo\public\img\`
   - `D:\work\qcash-ui-pelindo\public\bricams.png`
   - `D:\work\qcash-ui-pelindo\public\cbm.png`
   - `D:\work\qcash-ui-pelindo\public\qlola.png`
   - `D:\work\qcash-ui-pelindo\public\vercel.svg`
   - **Approximately 113 files removable.**
   - Whole-current-tree searches found no references to these paths or filenames. Preserve `favicon.ico` unless intentionally dropping the browser convention. Runtime-generated asset URLs are the only unobservable caveat.

3. **[delete] Two entire unused legacy locale dictionaries**
   - `D:\work\qcash-ui-pelindo\public\locales\en\translation.json` — **1,952 lines**
   - `D:\work\qcash-ui-pelindo\public\locales\id\translation.json` — **2,368 lines**
   - **4,320 lines removable.**
   - Current i18n code directly imports only the two `pelindo.json` files, and current source has one translation key: `dashboard.title`.

4. **[delete] Unused direct application dependencies**
   - `D:\work\qcash-ui-pelindo\package.json`
   - No current-filesystem imports/usages were found for:
     - `@bufbuild/protobuf`
     - `@connectrpc/connect`
     - `@connectrpc/connect-query`
     - `@connectrpc/connect-web`
     - `@tanstack/react-form-nextjs`
     - `decimal.js`
     - `es-toolkit`
     - `file-type`
     - `i18next-http-backend`
     - `valibot`
     - `zustand`
     - `react-aria-components`
     - `@tailwindcss/cli`
   - **13 direct dependencies removable.**
   - `axios` becomes a **14th** removable dependency after finding #1: its only current references are imports inside the stale declarations.
   - Retained as demonstrably used or build-integrated: Module Federation, auth provider, React Query, i18next/ni18n/react-i18next, Next/React, Tailwind/PostCSS, Rspack, React Compiler packages, Biome and TypeScript packages.

5. **[delete] Dead scripts not exposed by `package.json`**
   - `D:\work\qcash-ui-pelindo\scripts\dev-host-compat.mjs` — **181 lines**
   - `D:\work\qcash-ui-pelindo\scripts\build-globals-expose.mjs` — **53 lines**
   - `D:\work\qcash-ui-pelindo\scripts\analyze-bundle.js` — **37 lines**
   - `D:\work\qcash-ui-pelindo\scripts\update-version-file.js` — **9 lines**
   - **280 lines removable.**
   - None has a package script or active caller. The bundle script expects analyzer integration that is absent; the globals script expects a nonexistent `globals.bundle.css`; the version script has its write operation commented out.

6. **[delete] Legacy test scripts that cannot run**
   - `D:\work\qcash-ui-pelindo\package.json:14-16`
   - `test`, `test:coverage`, and duplicate `coverage` invoke Vitest, but Vitest is not installed and no tests/config exist.
   - **3 script entries removable.**
   - Add one test command only when the repository acquires an actual test runner and test files.

7. **[delete] Husky/lint-staged layer with no project hook**
   - `D:\work\qcash-ui-pelindo\package.json:12-13,51-52,58-60`
   - `D:\work\qcash-ui-pelindo\.husky\_`
   - There is no project-level `.husky/pre-commit` or other configured hook; `_` contains generated Husky shims only.
   - **2 dependencies, ~7 package lines, 17 generated files removable.**

8. **[delete] Obsolete parallel lint/format configuration**
   - `D:\work\qcash-ui-pelindo\.eslintrc.json` — **3 lines**
   - `D:\work\qcash-ui-pelindo\.prettierrc` — **14 lines**
   - `D:\work\qcash-ui-pelindo\.vscode\extensions.json` — remove obsolete ESLint/Prettier recommendations, **2 entries**
   - `D:\work\qcash-ui-pelindo\.idea\inspectionProfiles\Project_Default.xml`
   - `D:\work\qcash-ui-pelindo\.idea\prettier.xml`
   - Biome is the installed and scripted formatter/linter; ESLint and Prettier are not dependencies.
   - **At least 24 lines removable**, excluding broader IDE metadata deletion.

9. **[delete] Dead duplicate root configuration/style files**
   - `D:\work\qcash-ui-pelindo\config.ts` — **5 lines**
   - `D:\work\qcash-ui-pelindo\styles\globals.css` — **10 lines**
   - `D:\work\qcash-ui-pelindo\ni18n.config.ts` — **5 lines**
   - No current caller imports them. Active equivalents are `src/config.ts`, `src/styles/globals.css`, and `src/utils/i18n.ts`.
   - **20 lines removable.**

10. **[delete] Generated/local analysis artifacts**
    - `D:\work\qcash-ui-pelindo\.scannerwork\` — **34 generated files**
    - `D:\work\qcash-ui-pelindo\.playwright-mcp\page-2026-07-15T07-49-35-836Z.yml` — **56 lines**
    - The Playwright file is an unrelated Google homepage snapshot; `.scannerwork` is disposable Sonar output.
    - Add both directories to ignore rules if these tools are run repeatedly.

11. **[yagni] Repository-local OpenCode integrations unrelated to the application**
    - `D:\work\qcash-ui-pelindo\.opencode\skills\9router-chat\SKILL.md` — **75 lines**
    - `D:\work\qcash-ui-pelindo\.opencode\opencode.json` — **30 lines**
    - The skill has no repository caller. The config provisions three external MCP services, including `npx ...@latest`, independently of build/development scripts.
    - Delete the skill outright unless this repository specifically uses 9Router. Move personal MCP configuration to user-level OpenCode config if it is not team-required.
    - `D:\work\qcash-ui-pelindo\.opencode\package.json`, lockfile and `node_modules` are generated local tooling state: **1 local dependency/cache removable from the project tree**, recreatable by OpenCode.

12. **[shrink] Docker carries unused build-time environment plumbing**
    - `D:\work\qcash-ui-pelindo\Dockerfile:55-60,71-76`
    - No current application/config usage was found for:
      - `NEXT_PUBLIC_SIGNATURE`
      - `NEXT_PUBLIC_LIMIT_MEGABITPERSEC`
      - `NEXT_PUBLIC_SECRET_KEY_LOGIN_V2`
      - `NEXT_PUBLIC_DATAVIZ`
      - `NEXT_PUBLIC_MAX_RECORD`
      - `NEXT_PRIVATE_LOCAL_WEBPACK`
    - Remove their ARG/ENV pairs: **approximately 12 lines**.
    - `Dockerfile:81-84` clears proxy variables only after the final build command in a stage that is then discarded: **4 more lines removable**.

13. **[shrink] Docker copies the entire repository despite standalone output**
    - `D:\work\qcash-ui-pelindo\Dockerfile:43`
    - `D:\work\qcash-ui-pelindo\.dockerignore`
    - `COPY . ./` includes IDE, OpenCode, Playwright and scanner artifacts unless individually ignored. Expand `.dockerignore` for those directories, or copy only build inputs.
    - Also, deleting findings #1–3 materially reduces build context and final `public` image payload without Docker-specific machinery.

14. **[shrink] Overbroad TypeScript configuration**
    - `D:\work\qcash-ui-pelindo\tsconfig.json:5,15-18,22`
    - `allowJs` has no demonstrated TypeScript input need; `incremental` produces disposable `tsconfig.tsbuildinfo`; the catch-all `"*"` alias routes arbitrary imports toward generated MF types.
    - **Approximately 3–4 config lines removable**, plus a simpler include/path set.

15. **[shrink] Duplicate package-manager registry configuration**
    - `D:\work\qcash-ui-pelindo\.npmrc` — **3 lines**
    - `D:\work\qcash-ui-pelindo\.yarnrc` — **2 lines**
    - The repository documents and locks Yarn, while `.yarnrc` already routes the entire registry through Nexus. `.npmrc` is therefore a second policy surface with differing HTTP/HTTPS registry choices.
    - Prefer the Yarn configuration alone: **3 lines/file removable**, unless external CI explicitly invokes npm.

16. **[shrink] `dayjs` is shared despite no local consumer**
    - `D:\work\qcash-ui-pelindo\next.config.ts:44`
    - `D:\work\qcash-ui-pelindo\package.json:28`
    - No current source imports `dayjs`; its only use is declaring it as an eager federation singleton.
    - Removing the unused shared entry makes `dayjs` a **15th removable application dependency**. Retain only if the host contract explicitly requires this remote to provide it.

17. **[delete] Stale version/todo residue**
    - `D:\work\qcash-ui-pelindo\version` — **1 line**, disagrees with `package.json` and has no reader.
    - `D:\work\qcash-ui-pelindo\todo.md` — **1 obsolete line** referring to Moment.js, which is neither installed nor used.
    - Combined with the inert version-update script: **11 lines/files removable**.

## Defensible reduction total

- **Direct dependencies:** 16 immediately removable  
  - 14 unused application dependencies including `axios` after stale types are removed
  - 2 unused hook-tool dependencies
- **Conditional additional dependency:** `dayjs` — 1
- **Files:** approximately **810+ removable files**
  - 529 stale declarations
  - ~115 dead public files/locales
  - generated Husky/Sonar/tool artifacts and dead scripts/configs
- **Text:** approximately **17,800+ lines removable**
  - dominated by 13,025 declaration lines and 4,320 legacy locale lines

No files were edited. Git-deleted historical application files were excluded; usage conclusions are based on the current filesystem.

---

