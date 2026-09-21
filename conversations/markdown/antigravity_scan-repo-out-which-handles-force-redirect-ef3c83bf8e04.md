# scan this repo, and find out which file that handles force redirect...

- **ID**: `antigravity-c36d29e661b5eabd4e7b6e300f1e94df9ed9deec3e482326f8f3ef3c83bf8e04`
- **Source Tool**: `antigravity`
- **Date**: `2026-09-07 13:57:22`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

scan this repo, and find out which file that handles force redirection to v2

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I am checking the commit history to see which files were modified for the v2 redirect.

---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I am inspecting commit `86e2adda1b` to see what files were modified for redirecting to v2.

---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I am checking the diff for commit `86e2adda1b`.

---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I am checking the git log for the redirect map setup.

---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I will wait for the git log search to complete.

---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 34: Assistant

Waiting for the search to complete...

---

## Turn 35: Assistant

The primary file that handles the global force redirection logic to **v2** is:

### 1. Primary Redirection Handler
* **[`libs/qcash-ui/src/lib/layouts/index.tsx`](file:///D:/work/addons-web-client/libs/qcash-ui/src/lib/layouts/index.tsx#L13-L44)**
  * Inside the `Layouts` component's `useEffect`, it checks if the app is running in production (`process.env.NODE_ENV !== "development"`).
  * It matches the current `window.location.href` against the migration route patterns.
  * If matched, it prevents rendering and triggers a force redirect via:
    ```typescript
    window.location.replace(`/v2/${finalPath}`);
    ```

### 2. Redirection Route Configuration / Map
* **[`libs/qcash-ui/src/lib/layouts/consts.ts`](file:///D:/work/addons-web-client/libs/qcash-ui/src/lib/layouts/consts.ts#L1-L90)**
  * Defines `ListPage`, the **Routing Redirect Map — v2 Migration Rules** array.
  * Supports prefix matching (entries ending with `*`) and exact route matching.

---

### Other Specific Page-Level Force Redirections to v2
There are also page-level or post-login redirections to v2:

* **Digi Guarantee BG Mapping Pages** (force redirects directly on mount):
  * [`apps/web-client/pages/digi-guarantee/bg-mapping/index.tsx`](file:///D:/work/addons-web-client/apps/web-client/pages/digi-guarantee/bg-mapping/index.tsx#L31-L34)
  * [`apps/web-client/pages/digi-guarantee/bg-mapping/new.tsx`](file:///D:/work/addons-web-client/apps/web-client/pages/digi-guarantee/bg-mapping/new.tsx#L32-L35)
  * [`apps/web-client/pages/digi-guarantee/bg-mapping/[id].tsx`](file:///D:/work/addons-web-client/apps/web-client/pages/digi-guarantee/bg-mapping/[id].tsx#L44-L47)
  * [`apps/web-client/pages/digi-guarantee/bg-mapping/edit/[id].tsx`](file:///D:/work/addons-web-client/apps/web-client/pages/digi-guarantee/bg-mapping/edit/[id].tsx#L36-L39)
* **Login & Landing Page Flow**:
  * [`libs/main-page/src/lib/container/main-page.tsx`](file:///D:/work/addons-web-client/libs/main-page/src/lib/container/main-page.tsx#L209-L230) (redirects to `/v2` or via announcement modal)
  * [`apps/web-client/pages/main-page/index.tsx`](file:///D:/work/addons-web-client/apps/web-client/pages/main-page/index.tsx#L17-L25) (redirects authenticated users via `router.replace("/v2/homepage")`)

---

