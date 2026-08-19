# Design lean issuance fix (@explore subagent)

- **ID**: `ses_0152a3567ffeB24qJPgHtHGU8X`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 15:41:47`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. In ~/Development/qcash-ui-issuance-lc and qcash-ui shell, design the smallest robust code change to allow soft navigation away from hosted issuance-LC without header hard reload. Focus on removing duplicate Federation init/providers, preserving standalone mode, query-client lifecycle, and leaked event listeners. Inspect current tests and identify exact tests to modify/add. Explicitly assess whether any NEW singleton config is actually required after the lean fix. Return exact files/lines and proposed minimal diff.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

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

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
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
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **bash**: Tool: bash
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

## Conclusion

The lean fix belongs entirely in `qcash-ui-issuance-lc`. No shell code change and no new Module Federation singleton configuration are required.

The hosted shell already owns Federation initialization and global providers. Exposed remote pages do not execute the remote’s `_app.tsx`, so `_app.tsx` is the correct standalone-only boundary.

The current uncommitted implementation is not safe:

- Create can never leave its skeleton because `isFederationReady` is still required while its initializing effect is commented out.
- `services/query-client.ts` calls `useState` at module scope and accesses `window` unguarded, causing invalid-hook and SSR failures.
- `_app.tsx` renders `QuiProviders` while its standalone Federation initialization is commented out.
- Three page-3 components register a scroll listener during every render and never remove it.

---

# Proposed minimal diff

## 1. Preserve standalone Federation initialization and standalone providers

### File

`~/Development/qcash-ui-issuance-lc/pages/_app.tsx`

Current relevant lines: 11, 14, 26–37, 64–70.

Keep Federation initialization here because this file runs in standalone mode but is not included when the page is consumed as a remote.

```diff
 import { init, loadRemote } from "@module-federation/runtime";
 import { QuiProviders, QuiToast } from "@/components/global";

-if (typeof window !== "undefined") {
-  // currently commented out
+if (typeof window !== "undefined") {
   init({
     name: "qcash-ui-issuance-lc",
     remotes: [
       {
         name: "qcash-global-component",
         entry: `${config.baseFederation}/global-component/_next/static/chunks/remoteEntry.js`,
       },
+      {
+        name: "qui",
+        entry: `${config.baseFederation}/design-system/static/remoteEntry.js`,
+      },
     ],
   });
-  loadRemote("qcash-global-component/globals");
+
+  void Promise.all([
+    loadRemote("qcash-global-component/globals"),
+    loadRemote("qui/globals"),
+  ]);
 }

 ...
 <I18nextProvider i18n={i18nInstance}>
   <AuthProvider apiUrl={config.baseUrl as string}>
-    <GlobalStoreProvider>{getLayout(<Component {...pageProps} />)}</GlobalStoreProvider>
+    <QuiProviders>
+      <QuiToast.Provider />
+      <GlobalStoreProvider>{getLayout(<Component {...pageProps} />)}</GlobalStoreProvider>
+    </QuiProviders>
   </AuthProvider>
 </I18nextProvider>
```

This gives standalone mode its own:

- Federation initialization
- Auth/global-store providers
- QUI provider and toast
- Global remote styles

It does not duplicate any of those in hosted mode because the remote `_app.tsx` is not exposed.

---

## 2. Make the create page a plain exposed component

### File

`~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/create/index.tsx`

Current relevant lines: 1–4, 28–60, 83–84, 112–162, 189–205.

Remove all page-level Federation initialization, warming, provider loading, and readiness state.

```diff
-import { loadRemote, init } from "@module-federation/runtime";
-import config from "@/config";
 import { useAuth } from "@ordentco/addons-auth-provider";
-import { ComponentProps, ComponentType, ReactNode, useEffect, useRef, useState } from "react";
-import { QuiToast, UnauthorizedDialog } from "@/components/global";
+import { useEffect, useRef, useState } from "react";
+import { UnauthorizedDialog } from "@/components/global";
 ...
-import { LATEST_ISSUANCE_VERSION } from "@/constants";

-type FormIssuanceLcNewProps = ComponentProps<typeof FormIssuanceLcNew>;
-type FormIssuanceLcOldProps = ComponentProps<typeof FormIssuanceLcOld>;
-type ReadyComponents = {
-  FormOld: ComponentType<FormIssuanceLcOldProps>;
-  Form: ComponentType<FormIssuanceLcNewProps>;
-  Providers: ComponentType<{ children?: ReactNode }>;
-};
-
-const warmUpCreatePageRemoteModules = async () => {
-  ...
-};

 ...
-const [isFederationReady, setIsFederationReady] = useState(false);
-const [readyComponents, setReadyComponents] = useState<ReadyComponents | null>(null);

-// Entire commented/previous Federation useEffect
-...

-if (!isFederationReady || !readyComponents || !isSystemParamReady) {
+if (!isSystemParamReady) {
   return <IssuanceLcFormSkeleton />;
 }

-const { FormOld, Form, Providers } = readyComponents;
 const isVersion = srVersion;

 if (isVersion !== "2026") {
-  return <FormOld {...formProps} />;
+  return <FormIssuanceLcOld {...formProps} />;
 }

-return (
-  <Providers>
-    <QuiToast.Provider />
-    <Form {...formProps} />
-  </Providers>
-);
+return <FormIssuanceLcNew {...formProps} />;
```

The shell has already initialized both `qcash-global-component` and `qui` before loading this exposure. Reinitializing the runtime from this page is especially risky because:

- Shell runtime: `@module-federation/runtime` 0.18.1
- Remote runtime: `@module-federation/runtime` `^0.1.1`

---

## 3. Remove duplicate Federation initialization/providers from edit

### File

`~/Development/qcash-ui-issuance-lc/pages/import-lc/issuance-lc/edit/[transactionEventId].tsx`

Current relevant lines: 2–5, 31, 90–121, 211–220.

```diff
 import { useCallback, useEffect, useRef, useState } from "react";
-import { loadRemote, init } from "@module-federation/runtime";
-import config from "@/config";
 import { useAuth } from "@ordentco/addons-auth-provider";
-import { QuiProviders, QuiToast, UnauthorizedDialog } from "@/components/global";
+import { UnauthorizedDialog } from "@/components/global";

 ...
-const [isFederationReady, setIsFederationReady] = useState(false);

-useEffect(() => {
-  if (typeof window === "undefined") return;
-
-  const initFederation = async () => {
-    ...
-  };
-
-  initFederation();
-}, []);

 ...
-if (!isFederationReady) {
-  return <IssuanceLcFormSkeleton />;
-}

 ...
-return (
-  <QuiProviders>
-    <QuiToast.Provider />
-    <FormIssuanceLcNew {...formProps} />
-  </QuiProviders>
-);
+return <FormIssuanceLcNew {...formProps} />;
```

The old-form path should remain unchanged except that it no longer waits for redundant Federation initialization.

---

## 4. Restore the module-scoped QueryClient

### File

`~/Development/qcash-ui-issuance-lc/services/query-client.ts`

Current lines 1–30 should be replaced with the original, valid implementation:

```diff
 import { QueryClient } from "@tanstack/react-query";
-import { useState } from "react";

-export const queryClient = new QueryClient();
-window.__TANSTACK_QUERY_CLIENT__ = queryClient;
-
-const [queryClient] = useState(
-  () =>
-    new QueryClient({
-      ...
-    }),
-);
+export const queryClient = new QueryClient({
+  defaultOptions: {
+    queries: {
+      refetchOnWindowFocus: false,
+      retry: 1,
+      staleTime: 3000,
+    },
+  },
+});

 declare global {
-    interface Window {
-        __TANSTACK_QUERY_CLIENT__: import("@tanstack/query-core").QueryClient;
-    }
+  interface Window {
+    __TANSTACK_QUERY_CLIENT__: import("@tanstack/query-core").QueryClient;
+  }
 }

 if (globalThis.window) {
-    globalThis.window.__TANSTACK_QUERY_CLIENT__ = queryClient;
+  globalThis.window.__TANSTACK_QUERY_CLIENT__ = queryClient;
 }
```

Why this lifecycle is correct:

- One client is created per loaded remote module instance.
- Soft route changes do not recreate it because the JS module remains cached.
- `QueryClientProvider` can mount/unmount without discarding the client.
- The list page still needs its provider because `use-issuance-lc-filter.ts` uses contextual `useQuery`.
- Detail hooks explicitly pass this client at  
  `~/Development/qcash-ui-issuance-lc/hooks/use-issuance-detail-query.ts:105–125`.

Do not place a React hook at module scope.

---

## 5. Clean leaked scroll listeners

The page-4, currency and calendar listeners already have cleanup. The actual leaks are these three render-time registrations.

### Files

- `~/Development/qcash-ui-issuance-lc/components/create/page-3.tsx:441–443`
- `~/Development/qcash-ui-issuance-lc/components/create-new-old/page-3.tsx:1003–1005`
- `~/Development/qcash-ui-issuance-lc/components/create-new-skin/page-3.tsx:1011–1013`

Apply the same replacement in each:

```diff
-if (typeof window !== "undefined") {
+useEffect(() => {
   window.addEventListener("scroll", toggleVisible);
-}
+
+  return () => {
+    window.removeEventListener("scroll", toggleVisible);
+  };
+}, []);
```

Registering in render currently adds another callback after every state update. Those callbacks survive remote unmount and continue dispatching state updates during later header/navigation activity.

---

# Shell assessment

## No production change required

Relevant existing shell files:

- `~/Development/qcash-ui/pages/_app.tsx:39–50`  
  Initializes the host Federation runtime once.
- `~/Development/qcash-ui/pages/_app.tsx:163–197`  
  Owns `ModuleFederationProvider`, BRI/Auth bridge and QUI providers.
- `~/Development/qcash-ui/constants/features/registry.ts:516–520`  
  Registers issuance-LC for all `/import-lc/issuance-lc` subroutes.
- `~/Development/qcash-ui/pages/import-lc/issuance-lc/*.tsx:10–27`  
  The shell owns the header/footer layout and loads issuance as a child remote.

The hosted remote should therefore be passive: no runtime `init()` and no nested QUI provider. That leaves the shell header/router tree alive during client-side navigation.

Changing the four shell wrappers to a new layout mechanism would be broader and would not address the remote runtime corruption or leaked listeners. It is unnecessary for this fix.

---

# Exact test changes

## Modify create page test

### File

`~/Development/qcash-ui-issuance-lc/__tests__/pages/import-lc/issuance-lc/create/index.test.tsx`

At lines 106–117, replace the Federation-initialization expectation:

```diff
-it("initializes federation and renders the new form", async () => {
+it("renders as a hosted component without reinitializing federation", async () => {
   await renderPage();
   expect(screen.getByTestId("new-form")).toBeInTheDocument();
-  expect(init).toHaveBeenCalled();
+  expect(init).not.toHaveBeenCalled();
+  expect(loadRemote).not.toHaveBeenCalled();
   ...
 });
```

Replace lines 119–128, which currently test a behavior the page should no longer own, with a pending-system-parameter test:

```ts
it("keeps the skeleton visible while the system parameter is pending", async () => {
  mockGetDataSystemAll.mockReturnValueOnce(new Promise(() => {}));
  await renderPage();
  expect(screen.getByText("Loading form")).toBeInTheDocument();
});
```

Also use a mocked `QuiProviders` function and assert it was not rendered by the exposed page.

## Modify edit page test

### File

`~/Development/qcash-ui-issuance-lc/__tests__/pages/import-lc/issuance-lc/edit/transactionEventId-new-skin.test.tsx`

At lines 111–132:

```diff
-it("loads detail data and renders the new form after federation initializes", async () => {
+it("loads detail data without reinitializing host federation", async () => {
 ...
-  expect(init).toHaveBeenCalledWith(...);
-  expect(loadRemote).toHaveBeenCalledWith("qui/globals");
+  expect(init).not.toHaveBeenCalled();
+  expect(loadRemote).not.toHaveBeenCalled();
 ...
});
```

Add an assertion that the page-local `QuiProviders` and `QuiToast.Provider` mocks were not called.

## Add standalone `_app` coverage

### New test

`~/Development/qcash-ui-issuance-lc/__tests__/pages/_app.test.tsx`

Mock `appWithI18Next` as identity and verify:

1. Standalone module initialization registers both:
   - `qcash-global-component`
   - `qui`
2. Globals for both remotes are requested.
3. A rendered page is inside `AuthProvider`, `GlobalStoreProvider`, `QuiProviders` and `QuiToast.Provider`.

This protects the distinction between standalone bootstrap and hosted exposure.

## Add QueryClient lifecycle test

### New test

`~/Development/qcash-ui-issuance-lc/services/__tests__/query-client.test.ts`

Assert:

- `queryClient.getDefaultOptions().queries` contains the expected options.
- `window.__TANSTACK_QUERY_CLIENT__ === queryClient`.
- Requiring/importing the module again returns the same client from the module cache.

## Add listener cleanup tests

Modify:

- `~/Development/qcash-ui-issuance-lc/components/create-new-old/__tests__/page-3.test.tsx`
- `~/Development/qcash-ui-issuance-lc/components/create-new-skin/__tests__/page-3.test.tsx`

For each:

1. Spy on `window.addEventListener` and `window.removeEventListener`.
2. Render the page.
3. Capture the callback passed for `"scroll"`.
4. Unmount or call Testing Library `cleanup()`.
5. Assert the same callback is removed.

Also add equivalent coverage for legacy:

- New: `~/Development/qcash-ui-issuance-lc/components/create/__tests__/page-3-listener.test.tsx`

---

# Singleton configuration decision

**No new singleton is required after this fix.**

Specifically:

- Do not add `@module-federation/runtime` as shared. The shell should be the only hosted runtime initializer, and the two repositories currently use substantially different runtime versions.
- Do not add `@tanstack/react-query` to the shell. The shell does not depend on it, and the remote owns both its hooks and its module-scoped client.
- Do not add a global “hosted mode” singleton/config flag. The `_app.tsx` boundary already distinguishes standalone from exposed execution.
- React, React DOM and Next router internals are handled by `NextFederationPlugin`; no additional entry is needed.
- Existing auth/i18n singleton declarations in  
  `~/Development/qcash-ui-issuance-lc/next.config.js:41–58` should remain unchanged.

The smallest robust solution is therefore: standalone setup in `_app.tsx`, passive exposed pages, one valid remote-owned QueryClient, and listener cleanup.

---

