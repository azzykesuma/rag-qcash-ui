# i got some question regarding the function not implemented error, y...

- **ID**: `antigravity-95a8808fc6f6cd2f1e940c81808cd8516515d1ee926ec70395dfe8875a505719`
- **Source Tool**: `antigravity`
- **Date**: `2026-10-13 10:07:47`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

i got some question regarding the function not implemented error, you can search rag for this question
1. my question here is, there is a company that is using my app that has the most frequent errors for this function not implemented issue, so why is this company always hits this particular issue ? i mean other company do hits this error, but this single company hits the error the most

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

I have initiated a search in the Context Vault for discussions and incident records regarding the "function not implemented" error. I will analyze the findings as soon as the results are returned.

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 9: Assistant

I am retrieving and analyzing the occurrences and investigations of the "function not implemented" error from the vault logs. I will have the findings in a moment.

---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 14: Assistant

I am checking additional incident references across client reports. Let me know if you would like me to focus on specific modules as well.

---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 17: Assistant

Based on the investigations recorded in the Context Vault, here is the technical breakdown of what the **`"Function not implemented"`** error actually is, and **why one specific company is triggering it significantly more frequently than others**.

---

### 1. What the Error Actually Is (Root Cause)

In `@ordentco/addons-auth-provider` (and `@bri/addons-auth-provider`), the default React context defines fallback stub functions:

```ts
const AUTH_INITIAL_VALUES = {
  isAuthoritiesReady: false,
  guard: () => {
    throw new Error("Function not implemented."); // 👈 Thrown ONLY when guard() is executed!
  },
  setToken: () => { throw new Error("Function not implemented."); },
  // ...
};

const AuthContext = createContext(AUTH_INITIAL_VALUES);
```

This error is **action-triggered**:
* It occurs when a remote micro-frontend (MFE) component consumes `useAuth()` and invokes `guard()` (e.g. inside `useEffect(() => { guard(); }, [])`), but **Module Federation failed to connect the remote component to the Shell's active `<AuthProvider>`**.
* When there is a **version drift**, **missing `singleton: true` in `next.config.js`**, or a **dynamic import mount delay**, Webpack instantiates an isolated `AuthContext` in the remote chunk. Without a provider wrapping it, `useAuth()` falls back to `AUTH_INITIAL_VALUES` and crashes with `Uncaught Error: Function not implemented`.

---

### 2. Why Does This Single Company Hit It the Most?

Because all companies share the same codebase, the disparity in error frequency comes down to **module entitlement, payload characteristics, user behavior patterns, and network environments**:

#### A. Entitlement to Un-migrated, Guard-Heavy Remote MFEs (e.g., Mass Transfer / LC / Payroll)
* Not all client companies use the same modules. Standard companies might primarily use Single Transfer, Account Inquiry, or Core Dashboard screens (which may not call `guard()` directly).
* In contrast, large enterprises frequently use modules like **Mass Transfer**, **Issuance LC**, or **Auto Collection**.
* As identified in the stack trace logs, **all exposed pages in `mass-transfer` (Maker Confirmation, Upload Transaction, Transaction Action, etc.) contain unhandled `guard()` calls in passive effects**:
  ```tsx
  useEffect(() => {
    guard(); // ⚠️ Fires on every mount without promise rejection handling
  }, []);
  ```
  If this company's daily business operations heavily revolve around these specific remote MFEs, their users hit the dangerous code path hundreds of times more often than other companies.

#### B. Heavy Maker / Checker / Signer Traffic and Frequent SPA Route Transitions
* A high-transaction company usually has multiple dedicated operators (makers, checkers, signers) working in shifts.
* These users frequently jump between the **Task List / Inbox**, **Dashboard**, and **Remote Feature Screens** to review and approve batches.
* Because every inter-module SPA navigation mounts and unmounts federated micro-frontends, each transition re-triggers the remote chunk's `useEffect(() => guard(), [])`. Higher navigation volume directly correlates with a higher probability of hitting the race window.

#### C. Massive Authority / Menu Payload Size Widens the Race Window
* Large enterprise accounts have complex organizational structures: dozens or hundreds of accounts, multi-tier approvals, and extensive product licenses.
* When their users log in or validate auth, the backend responses for `/menu/me` and `/auth/me` are substantially larger in payload size.
* A heavier payload increases API latency, which **widens the timing window** during which `isAuthoritiesReady` is `false` and the provider is still hydrating. If a remote component mounts before this completes, any early `guard()` invocation immediately explodes.

#### D. Corporate Network Proxies & Aggressive Browser Disk Caching
* Enterprise clients often operate inside locked-down corporate LANs, forward proxies, VPNs, or virtual desktop infrastructures (VDI / Citrix).
* If the company's internal proxy or browser policy aggressively caches JavaScript chunks (`remoteEntry.js` or chunk bundles), their client machines may run **stale remote chunks that still contain older auth guard calls**, while the host shell has already moved forward.
* Mismatched chunk hashes across federation boundaries guarantee context isolation, leading directly to the `Function not implemented` stub fallback.

#### E. Multi-Tab & Concurrent Session Behavior
* High-volume financial operators commonly open multiple tabs simultaneously (e.g., keeping the dashboard open in Tab 1 while processing transactions in Tabs 2 and 3).
* When token expiration or storage sync events occur across tabs, multiple concurrent `guard()` or revalidation requests fire simultaneously, overloading the shared federation state and causing race conditions.

---

### Summary & Permanent Resolution

The issue is disproportionately hitting this company because **their specific workflow (heavy remote MFE usage + frequent maker-checker page transitions + larger account/permission payloads)** repeatedly triggers the unhandled `guard()` call sites.

The established permanent fix across the architecture consists of:
1. **Centralizing the Guard (Passive Remotes):** Removing `useEffect(() => { guard(); }, [])` from federated remote pages and having remotes only passively read auth status (`isAuthoritiesReady`, `username`, `userType`).
2. **Strict Webpack Singleton Enforcement:** Ensuring all remote repositories configure `@ordentco/addons-auth-provider` with `{ singleton: true, requiredVersion: false }`.
3. **Static Provider Imports in the Host:** Importing `<AuthProvider>` statically rather than through `next/dynamic({ ssr: false })` to eliminate tick-0 mounting delays.

---

