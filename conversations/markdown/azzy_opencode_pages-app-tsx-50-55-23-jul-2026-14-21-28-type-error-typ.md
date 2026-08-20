# AuthProviderProps router type error in _app.tsx

- **Author**: `azzy`
- **ID**: `ses_07212e958ffePUo0mg5uCLjpiQ`
- **Source Tool**: `opencode`
- **Date**: `2026-07-23 14:42:33`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

./pages/_app.tsx:50:55
23-Jul-2026 14:21:28	Type error: Type '{ children: Element; apiUrl: string; router: any; }' is not assignable to type 'IntrinsicAttributes & AuthProviderProps'.
23-Jul-2026 14:21:28	  Property 'router' does not exist on type 'IntrinsicAttributes & AuthProviderProps'.
23-Jul-2026 14:21:28 when i tried building this in my pipeline, i got this error, can you explain what is happening here?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **glob**: Tool: glob

</details>

I’ll inspect `_app.tsx`, the `AuthProvider` definition, and dependency versions to trace why the pipeline’s TypeScript check rejects `router`.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **glob**: Tool: glob
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

The installed contract is the key: this project pins `@ordentco/addons-auth-provider` to `0.9.140-mfe`, while `_app.tsx` still passes a `router` prop. I’m checking the recent version/cast changes to determine whether this is an API-version mismatch or only an invalid workaround.

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch
- **webfetch**: Tool: webfetch

</details>



---

## Turn 8: Assistant

TypeScript is validating the JSX props passed to `AuthProvider` at `pages/_app.tsx:50`:

```tsx
<AuthProvider apiUrl={config.baseUrl as string} router={router}>
```

At the time of the failed pipeline, `yarn.lock` resolved `@ordentco/addons-auth-provider` to `0.9.137`. That version defines `AuthProviderProps` roughly as:

```ts
interface AuthProviderProps {
  children?: React.ReactNode;
  cookieName?: string;
  apiUrl: string;
}
```

Because it has no `router` property, TypeScript rejects the extra JSX prop. `IntrinsicAttributes` is React’s internal JSX typing and is not the underlying problem.

The cast introduced here did not fix the version mismatch:

```tsx
router={router as AuthProviderProps["router"]}
```

A cast can change the type of a value, but it cannot add a missing property to `AuthProviderProps`.

The repository history shows the cause and subsequent fix:

- Pipeline failure: 14:21
- Package was subsequently pinned to `0.9.140-mfe`: 14:34
- That exact package version does define `router: AuthRouter` in `AuthProviderProps`
- The current `package.json` and `yarn.lock` both resolve `0.9.140-mfe`

This was therefore a dependency/API contract mismatch, likely exposed because the pipeline performed a clean installation while a local environment had different or cached package contents. With the current exact version and lockfile committed, a clean pipeline install should recognize the `router` prop.

---

