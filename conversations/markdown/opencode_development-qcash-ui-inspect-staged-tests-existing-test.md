# Assess provider tests (@explore subagent)

- **ID**: `ses_0009aad6affeivF4JgBEQzxhG8`
- **Source Tool**: `opencode`
- **Date**: `2026-08-14 15:30:58`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

In ~/Development/qcash-ui, inspect staged tests and existing tests/config for BRIProviders, OrdentProviders, _app provider wiring, and federation provider mocks. Determine what regressions are covered or missing and suggest exact focused commands/tests. Do not edit files. Return findings with file/line references. Desired thoroughness: medium.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **bash**: Tool: bash

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

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

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **glob**: Tool: glob

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 15: Assistant

## Summary

The staged test covers the intended static provider composition well, but `_app`’s client-only-to-static wiring change and Module Federation singleton/runtime behavior remain untested. One existing provider suite is now stale and fails.

## Covered

- The staged test verifies the full provider order:
  `BRI Auth → BRI Store → Ordent Auth → Ordent Store → children`  
  `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:50-64`
- It verifies `apiUrl`, router identity, Ordent base route, and `/main-page` login route for `BriProviders`.  
  `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:66-76`
- It covers `/landing-page` for standalone `OrdentProviders`.  
  `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:79-96`
- Existing BRI tests also cover children, URL, provider presence, nesting, router-hook use, and multiple children.  
  `~/Development/qcash-ui/components/providers/__tests__/BRIProviders.test.tsx:78-165`
- TypeScript accepts the staged static imports and props. `npx tsc --noEmit --pretty false` passed.
- Both auth packages remain configured as Federation singletons.  
  `~/Development/qcash-ui/next.config.js:15-21`

## Missing or regressed

### 1. Existing `OrdentProviders` tests are broken by the static-import conversion

The suite still mocks `next/dynamic` and expects that mock to manufacture the providers:

- Dynamic mock construction:  
  `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:52-62`
- `next/dynamic` mock:  
  `~/Development/qcash-ui/components/providers/__tests__/OrdentProviders.test.tsx:76-79`

The implementation now imports the real package directly:

- `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:2`
- `~/Development/qcash-ui/components/providers/OrdentProviders.tsx:14-20`

Observed result: all three existing Ordent tests fail because the real auth provider executes and reads `localStorage`. The suite needs package mocks analogous to the staged test, not a `next/dynamic` mock.

Also, no production caller of standalone `OrdentProviders` was found; references are implementation plus tests only.

### 2. `_app` provider wiring has no direct test

The staged change makes `BriProviders` a static import:

- `~/Development/qcash-ui/pages/_app.tsx:19`

It is mounted inside the dynamically loaded Federation provider and around the remote `QuiProviders`:

- Federation provider dynamic load: `~/Development/qcash-ui/pages/_app.tsx:60-65`
- Remote provider load: `~/Development/qcash-ui/pages/_app.tsx:66-68`
- Runtime tree: `~/Development/qcash-ui/pages/_app.tsx:160-194`

No test imports or renders `pages/_app.tsx`. The only `_app`-named test reimplements one wrapper locally and therefore cannot detect provider-order or import-mode regressions:

- `~/Development/qcash-ui/components/ui/__tests__/_app-mfa-hide.test.tsx:12-18`

Missing focused test: mock `ModuleFederationProvider`, `BriProviders`, `QuiProviders`, and ancillary providers, render the actual app, and assert:

`ModuleFederationProvider > BriProviders > QuiProviders > AuthBridgeSync > page`.

### 3. SSR/client-only behavior is no longer protected

Previously `BriProviders` had `ssr: false`; it will now participate in `_app` server rendering. The staged unit test mocks both auth packages:

- `~/Development/qcash-ui/components/providers/__tests__/AuthProviders.test.tsx:26-38`

Consequently, it cannot detect import-time browser-global access, hydration differences, or a Federation shared-module initialization problem in the real packages. A production Next build and hard-refresh smoke test are the useful checks here.

### 4. Federation provider tests do not cover auth-provider sharing or `_app`

The Federation suite mocks routing, module-loading hooks, and all `next/dynamic` components:

- `~/Development/qcash-ui/components/federation/__tests__/provider.test.tsx:12-35`

It only verifies children, monitor visibility, and readiness polling:

- `~/Development/qcash-ui/components/federation/__tests__/provider.test.tsx:84-123`

It does not exercise `_app`, `loadRemote("qui/providers")`, either auth singleton, or context identity between host and remote.

The suite currently also fails in this environment because `localStorage` is unavailable when the provider initializes:

- Access: `~/Development/qcash-ui/components/federation/provider.tsx:81-84`
- Vitest setup contains only jest-dom and no storage shim:  
  `~/Development/qcash-ui/vitest-setup.js:1`  
  `~/Development/qcash-ui/vitest.config.mts:8-12`

### 5. Route-toggle matrix is incomplete

Missing assertions:

- `BriProviders` with `toggleRedirectLandingPage=true`.
- Standalone `OrdentProviders` with the false toggle and `/main-page`.
- Exact props/no-extra-props for both global-store providers.

The migration documentation additionally calls for hard refresh, legacy-to-migrated navigation, relogin, and authority-readiness testing:

- `~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md:160-167`

## Focused commands

```bash
# Passing staged/static composition coverage
npm test -- components/providers/__tests__/AuthProviders.test.tsx components/providers/__tests__/BRIProviders.test.tsx

# Currently exposes the stale next/dynamic-based Ordent tests
npm test -- components/providers/__tests__/OrdentProviders.test.tsx

# Currently exposes the Federation test storage/config problem
npm test -- components/federation/__tests__/provider.test.tsx

# Static type check; observed passing
npx tsc --noEmit --pretty false

# Best check for the newly exposed SSR/Federation import path
npm run build
```

Observed focused run: `AuthProviders` and `BRIProviders` passed; `OrdentProviders` failed 3/3; Federation provider failed 2/3. No files were edited.

---

