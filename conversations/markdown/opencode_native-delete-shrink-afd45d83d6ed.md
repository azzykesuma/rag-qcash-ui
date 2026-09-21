# Audit source complexity (@explore subagent)

- **ID**: `opencode-ad60c38d4a3c2e590c63bf38cb8357847c68a22b255c8d170c16afd45d83d6ed`
- **Source Tool**: `opencode`
- **Date**: `2026-07-15 15:44:09`
- **Project**: `qcash-ui-pelindo`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Perform a very thorough whole-tree audit of D:\work\qcash-ui-pelindo focused ONLY on over-engineering, not correctness/security/performance. Inspect all current application source under src/, generated/type material only enough to classify it, and identify concrete opportunities tagged delete/stdlib/native/yagni/shrink. Verify callers/usages before calling anything dead or single-use. Return ranked findings with exact paths, concise rationale, replacement, and a defensible approximate line count removable per finding. Do not edit files. Ignore deleted files shown by git status and audit the current filesystem tree.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

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
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
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

<details><summary>🔧 Tool Invocations (8)</summary>

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

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

## Audit scope

Inspected all 17 current files under `D:\work\qcash-ui-pelindo\src` (735 physical lines), plus configuration and generated declarations only to establish usage. Searches covered the current whole tree; Git-deleted files were ignored.

## Ranked findings

1. **[delete][yagni] Remove 25 unused QUILA component wrappers**
   - Path: `D:\work\qcash-ui-pelindo\src\components\ui\index.tsx:9-129`
   - Only `QuiProvider` is imported anywhere. `Button` through `Modal`, including every compound subcomponent, have zero callers.
   - Replacement: retain only the `QuiProvider` dynamic export; add wrappers when first consumed.
   - **Approximately 121 lines removable.**

2. **[native][delete][shrink] Replace the custom router facade with existing routers**
   - Paths:
     - `D:\work\qcash-ui-pelindo\src\hooks\useRouter.tsx`
     - `D:\work\qcash-ui-pelindo\src\pages\_app.tsx:5,17`
     - `D:\work\qcash-ui-pelindo\src\pages\dashboard\index.tsx:6,10-16,32,36`
   - Verified callers: `useAppRouter` is used only by `_app`; `AppRouterProvider` only wraps the dashboard, but no dashboard descendant calls `useAppRouter`.
   - Replacement: use `next/compat/router` directly in `_app`, falling back to the installed auth provider’s exported `noopRouter`. Remove the unused dashboard router context and prop.
   - **Approximately 80–86 net lines removable.**

3. **[delete][yagni] Remove the unused application provider stack**
   - Paths:
     - `D:\work\qcash-ui-pelindo\src\components\providers\app-provider.tsx`
     - `D:\work\qcash-ui-pelindo\src\pages\dashboard\index.tsx:4,33,35`
   - `QueryClientProvider`, its injectable client, and `GlobalStoreProvider` have no consumers under `src`; no query hooks or global-store hooks exist. `AppProviders` has exactly one caller, which never supplies `queryClient`.
   - Replacement: render `DashboardPage` directly. Reintroduce the relevant provider when a consumer appears.
   - **Approximately 34–36 lines removable.**

4. **[shrink][yagni] Collapse defensive i18n resource replication**
   - Paths:
     - `D:\work\qcash-ui-pelindo\src\utils\i18n.ts:55-79`
     - `D:\work\qcash-ui-pelindo\src\hooks\usePelindoTranslation.ts`
     - `D:\work\qcash-ui-pelindo\src\components\providers\i18n-provider.tsx`
   - The resources are already supplied when the private i18next instance initializes. They are then redundantly copied into that instance, the global singleton, and the context instance during module load, initialization, provider mount, and every translation-hook mount.
   - Replacement: initialize the private instance once, provide it once, and make `usePelindoTranslation` a direct `useTranslation("pelindo")` call.
   - **Approximately 33–40 lines removable.**

5. **[delete][yagni] Remove the currently empty generated global-style exposure**
   - Path: `D:\work\qcash-ui-pelindo\src\styles\global.exposes.ts`
   - `CSS_TEXT` is empty, so the exposed module only inserts an empty `<style>`. Its generator is not wired into package scripts, and `src/styles/globals.bundle.css` does not exist.
   - Replacement: remove this source module and its federation exposure until a host actually needs a generated CSS payload.
   - **Approximately 34 source lines removable.**
   - Caveat: external host imports cannot be verified from this repository, although the current payload remains empty even if imported.

6. **[shrink] Use the auth provider’s declared authority type**
   - Path: `D:\work\qcash-ui-pelindo\src\utils\auth.ts`
   - Both exported functions are called only with `useAuth().authorities`, declared by the installed package as `Map<string, string[]>`. The local structural interface, runtime map-shape predicate, array validation, and `.has()` before `.get()` duplicate that contract.
   - Replacement: accept `Map<string, string[]>`; use `authorities.has("PELINDO")` and `authorities.get("PELINDO")?.includes(...)`.
   - **Approximately 18–22 lines removable.**

7. **[stdlib][shrink] Replace nine positional environment arguments and repeated flags**
   - Paths:
     - `D:\work\qcash-ui-pelindo\src\utils\environment.ts`
     - `D:\work\qcash-ui-pelindo\src\config.ts`
   - `getBaseUrlByEnv` has exactly one caller. It receives nine same-typed positional values and repeatedly reads the same hostname into named booleans used once.
   - Replacement: read `location.hostname` once and use an ordered array of `[predicate, environment variable]` entries with `Array.find`, or inline the ordered selection into `src/config.ts`.
   - **Approximately 18–25 net lines removable.**

8. **[delete][shrink] Remove redundant eager federation initialization**
   - Paths:
     - `D:\work\qcash-ui-pelindo\src\pages\_app.tsx:1,7,19-21`
     - `D:\work\qcash-ui-pelindo\src\utils\federation.ts:70-72`
   - `initializeQuiFederationRuntime` has one caller and only invokes the same resolver that every `loadQuiRemote` call already invokes.
   - Replacement: let the first remote load initialize or retrieve the runtime.
   - **Approximately 7–8 lines removable.**

9. **[shrink] Remove unused theme declarations**
   - Path: `D:\work\qcash-ui-pelindo\src\styles\globals.css:4-5,8-10,13,16-17`
   - Whole-tree class search found no use of the Geist font aliases or the `primary-light`, `primary-link`, `accent`, `warning`, `pending-bg`, and `gray-bg` colors.
   - Replacement: keep only tokens referenced by current markup; add tokens alongside their first use.
   - **8 lines removable.**

10. **[delete][shrink] Remove the unused public language helper**
    - Path: `D:\work\qcash-ui-pelindo\src\utils\i18n.ts:38-40`
    - `getSupportedLanguage` has zero callers. `SupportedLanguage` is used only internally despite being exported.
    - Replacement: delete the function and stop exporting the internal type.
    - **3 lines removable**, plus two unnecessary export keywords.
    - This overlaps finding 4.

11. **[shrink][yagni] Merge the single-use dashboard container**
    - Paths:
      - `D:\work\qcash-ui-pelindo\src\containers\dashboard\index.tsx`
      - `D:\work\qcash-ui-pelindo\src\pages\dashboard\index.tsx`
    - The container has exactly one caller and adds only a prop interface and component boundary; it has no independent state, reuse, or framework role.
    - Replacement: place its markup in the dashboard page after the authorization gates.
    - **Approximately 5–7 net lines removable.**

12. **[shrink] Remove the unnecessary fragment around one child**
    - Path: `D:\work\qcash-ui-pelindo\src\components\providers\app-provider.tsx:25-29`
    - The fragment contains only `{children}`.
    - Replacement: `{isQuiGlobalsReady ? children : null}`.
    - **2 lines removable.**
    - Superseded if finding 3 is taken.

## Larger consolidation option

**[delete][yagni] Remove all consumer-side QUILA machinery until a QUILA component is used.**

Affected paths:

- `D:\work\qcash-ui-pelindo\src\components\ui\index.tsx`
- `D:\work\qcash-ui-pelindo\src\utils\federation.ts`
- `D:\work\qcash-ui-pelindo\src\hooks\useQuiGlobalsReady.ts`
- `D:\work\qcash-ui-pelindo\src\components\providers\app-provider.tsx`
- corresponding wrappers in `_app.tsx` and `pages/dashboard/index.tsx`

The current dashboard renders only native elements. This would remove roughly **300–310 source lines**, but overlaps findings 1, 3, 5, and 8 and should not be added to their totals.

## Generated/type classification

`D:\work\qcash-ui-pelindo\@mf-types\qcash-global-component` is a large generated remote-declaration bundle, not application source. No current `src` file imports it. It should be regenerated or removed as a unit rather than manually simplified. The current source instead references `@mf-types/qui/apis`, which is not part of that bundle; no over-engineering finding was assigned to generated declarations themselves.

---

