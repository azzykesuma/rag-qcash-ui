# please do research is different version ordentco provider will had ...

- **ID**: `e0636c26-4fd9-4d1d-9155-86d1bef0048b`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-19 14:42:06`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

please do research is different version ordentco provider will had not same data/token and its need to be a TC ? is the singleton makes the provider version on each mfe domain is same or not ?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 16: Assistant

Based on research into the `qcash-ui` codebase, the Module Federation configuration, and browser runtime behaviors, here is a detailed breakdown of how the **Ordent provider** (and the new **BRI provider**) handles data/tokens, the singleton behavior, and what happens across MFE domains.

---

### Part 1: Will different versions of `@ordentco/addons-auth-provider` have different data/tokens, and does it need to be a "TC"?

#### 1. Why different versions might have different data/tokens
* **In-Memory Instantiation (Multiple Contexts):** If different versions are loaded without a successful singleton resolution, Webpack/Module Federation will bundle and instantiate multiple separate copies of the package. This creates multiple React context instances. The host will write to one context, and the remote MFE will read from another. They will **not** share token/auth state because they point to different memory references.
* **Storage Keys / Internal State:** Different versions of `@ordentco/addons-auth-provider` might use different internal keys or storage mechanisms to persist the token (e.g., in `localStorage`, `sessionStorage`, or cookies). If one version uses `accessToken` and another uses `access-token`, they will not see each other's data.
* **Current Sync Mechanism:** In [AuthBridgeSync.ts](file://~/Development/qcash-ui/components/providers/AuthBridgeSync.ts#L107-L116), the project synchronizes tokens from local storage using:
  ```ts
  const accessToken = localStorage.getItem("access-token");
  ```
  If MFEs are running on the **same domain/origin**, they can successfully share this local storage token. If they are on different domains, `localStorage` is isolated by the browser's Same-Origin Policy.

#### 2. Does it need to be a "TC"?
Depending on what **"TC"** refers to in your team's terminology:
* **TC = Token Cookie:** 
  * **Yes, it does.** If MFEs run on different domains (e.g. `internal-service.example.com` and `internal-service.example.com`), they cannot share `localStorage`. To share session tokens, the token must be saved in a **Shared Cookie** (Token Cookie) with the domain attribute set to the parent domain (e.g., `domain=.bri.co.id`). This allows the browser to automatically send the cookie across all subdomains.
* **TC = Test Case (or Testing Scenario):** 
  * **Yes.** A dedicated Test Case should be written (similar to the ones in [AuthBridgeSync.test.tsx](file://~/Development/qcash-ui/components/providers/__tests__/AuthBridgeSync.test.tsx)) to guarantee that when different provider versions or sources are mounted, the auth bridge correctly maps the fallback logic and syncs credentials without leaking state.
* **TC = Technical Clarification:**
  * **Yes.** A Technical Clarification is recommended to align all remote MFE teams to consume auth data only through the shared **`AuthBridge`** (reading `window.__QCASH_AUTH_BRIDGE__` and listening to the `qc-bridge-sync` event) rather than importing provider hooks like `useAuth()` directly.

---

### Part 2: Does the `singleton` setting make the provider version on each MFE domain the same?

Here is how Module Federation resolves [next.config.js](file://~/Development/qcash-ui/next.config.js#L16-L21) where both providers are configured with `singleton: true`:
```javascript
"@bri/addons-auth-provider": { requiredVersion: false, singleton: true },
"@ordentco/addons-auth-provider": { requiredVersion: false, singleton: true },
```

#### 1. Inside the Same Page / Single Tab (Same window context)
* **Yes.** The `singleton: true` configuration guarantees that the Module Federation runtime will load **only one** instance of the package in the browser's memory for that page session.
* Even if host and remote MFEs specify different versions in their `package.json`, the runtime resolves them to a single shared instance (usually selecting the highest compatible version or the host's version, since `strictVersion` is not enabled). 
* Since they share the same JS instance, they share the same in-memory React Context and state (tokens, user info, etc.).

#### 2. Across Different Domains / Subdomains (Or Page Hard-Reloads)
* **No.** A JavaScript singleton is only a singleton within the **runtime browser memory of a single window context**.
* If MFEs are hosted on different domains, or if navigating between them causes a full page reload, the browser instantiates a completely separate JavaScript environment. 
* In this case, `singleton: true` has no effect across the different domains. Each domain will load its own instance of the provider in its own memory. Sharing token/data must then rely on browser-level storage (e.g., shared cookies) rather than React Context memory.

---

## Turn 17: User

TC i mean is test case

---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 21: Assistant

Here is the research on whether this scenario needs to be covered by a **Test Case (TC)** and how the `singleton` option behaves across MFE domains:

---

### 1. Does different version data/token mismatch need to be a Test Case (TC)?

**Yes, it absolutely should be covered by a Test Case.** 

To prevent integration regressions and token sync failures between host applications and MFEs running different provider versions, you should verify these scenarios. 

The codebase already implements some of these checks in [AuthBridgeSync.test.tsx](file://~/Development/qcash-ui/components/providers/__tests__/AuthBridgeSync.test.tsx) and [AuthMonitor.test.tsx](file://~/Development/qcash-ui/components/federation/monitor/auth/__tests__/AuthMonitor.test.tsx), but you should ensure the following test cases are defined:

* **TC 1: Storage-Based Token Synchronization**
  * **Goal:** Verify that if the host and remote MFEs load at different times or with different provider instances, they synchronize their token values through `localStorage.getItem("access-token")`.
  * **Reference:** Already partially verified in [AuthBridgeSync.test.tsx:L196-L217](file://~/Development/qcash-ui/components/providers/__tests__/AuthBridgeSync.test.tsx#L196-L217) (*"on a BRI-preferred path uses 'bri' source even if briAuth is guest but has token in storage"*).
* **TC 2: Unified Bridge Contract Exposure**
  * **Goal:** Verify that the host extracts the state of the active provider and publishes it to `window.__QCASH_AUTH_BRIDGE__` so remote MFEs read a single unified payload instead of querying their own provider versions directly.
* **TC 3: Session Expiry / Logout Cleanup**
  * **Goal:** Verify that when the user logs out or the session expires, both provider states and the global window bridge are completely cleared to prevent session leakage.
  * **Reference:** Verified in [AuthBridgeSync.test.tsx:L168-L194](file://~/Development/qcash-ui/components/providers/__tests__/AuthBridgeSync.test.tsx#L168-L194) (*"clears stale bridge data and localStorage when both providers are in guest/logged-out state"*).

---

### 2. Does the `singleton` configuration make the provider version on each MFE domain the same?

It depends on whether you are running them inside the same page context or across different domains:

#### **A. Within the Same Page / Tab (Single Execution Context)**
* **Yes.** In [next.config.js](file://~/Development/qcash-ui/next.config.js#L16-L21), both `@bri/addons-auth-provider` and `@ordentco/addons-auth-provider` are declared as `singleton: true`.
* During execution, Module Federation will resolve version mismatches (e.g., if one MFE requests `v0.9.0` and another requests `v0.10.0`) to a **single shared runtime version in memory**.
* All MFEs loaded on that page will point to the exact same React context, memory address, and token variables.

#### **B. Across Different MFE Domains or Full Page Reloads**
* **No.** A JavaScript singleton is strictly an **in-memory object** tied to the browser’s active `window` context.
* If your MFEs are loaded across different domains (e.g., `internal-service.example.com` and `internal-service.example.com`) or if you perform a hard refresh, the browser creates a new, isolated execution context.
* The Module Federation singleton mechanism cannot bridge across different window/tab boundaries. Each domain will load its own version and initialize its own provider memory.
* To sync tokens/data across different domains, you must use **Token Cookies** configured with a shared parent domain (e.g., `domain=.bri.co.id`) rather than relying on the JS singleton.

---

## Turn 22: User

please check is the singleton make sure the runtime shared data is right ? or it be breaking changes for the other repo who not use same version as qcash-ui ?

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

### 1. Does `singleton: true` guarantee the runtime shared data is correct?

**No, not automatically.** 

While `singleton: true` guarantees that **only one instance of the code is executed in memory** (preventing multiple isolated contexts), it does not guarantee that the shared data is *correct* or compatible. 

* **The Selection Mechanism:** By default (with `requiredVersion: false`), Module Federation resolves a singleton by loading the **highest version** available among the host and the loaded remotes.
* **The Risk:** If the highest resolved version behaves differently or expects a different configuration than what the host initialized, the runtime shared data could be corrupted, incomplete, or structured in an unexpected format.

---

### 2. Can it cause breaking changes for other repositories that use a different version?

**Yes, it can easily cause runtime breaking changes.** 

Since all repositories (host and MFEs) are forced to use the single resolved version of `@ordentco/addons-auth-provider` or `@bri/addons-auth-provider` at runtime, any version mismatch can break other repos in the following ways:

#### A. API and Export Mismatches (Crashes)
If the version loaded by `qcash-ui` has renamed hook exports, altered hook signatures, or removed functions that an older MFE relies on, the MFE will crash at runtime with errors like:
```text
TypeError: (0 , auth_provider.useAuth) is not a function
```
Because TypeScript only type-checks at **build-time**, the MFE might compile perfectly in its own repository, but crash in production when loaded by the host.

#### B. Data Structure Mismatches (Silent Failures)
If a newer version of the provider changes the shape of the user session object (e.g., changing the structure of `menuData`, `authorities`, or `userType`), an older MFE expecting the legacy structure will fail to render menus, block access to pages, or throw exceptions when trying to read nested fields.

#### C. Behavioral Changes
If the newer provider introduces different side effects (e.g., changes to storage keys, token refresh schedules, redirect paths on session expiry), it will override the expected behavior of the older MFE.

---

### 3. How `qcash-ui` Mitigates This: The `AuthBridge` Pattern

To prevent these exact version-mismatch breaking changes, the codebase uses the **`AuthBridge` pattern** (described in [ordent-to-bri-provider-migration.md](file://~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md#L120-L130)):

1. **Avoid Direct Context Calls:** Remotes are discouraged from calling `useAuth()` directly from the package, as this binds them directly to the provider instance.
2. **The Sync Bridge:** The host ([AuthBridgeSync.ts](file://~/Development/qcash-ui/components/providers/AuthBridgeSync.ts)) reads the data from the provider contexts and publishes a reduced, flat JS object to `window.__QCASH_AUTH_BRIDGE__`.
3. **Event Subscriptions:** The remote MFEs consume the plain data object and listen for the `"qc-bridge-sync"` event instead of interacting directly with the provider library.

---

### 4. Recommendations for Safe Version Upgrades

If other repositories are not on the same provider version as `qcash-ui`, enforce the following safety rules:

1. **Use the AuthBridge Adapter:** Ensure all remotes migrate away from direct `@ordentco/addons-auth-provider` context hooks and use the `window.__QCASH_AUTH_BRIDGE__` adapter instead.
2. **Backwards Compatibility:** When updating `@ordentco/` or `@bri/` packages, do not remove legacy properties or rename exports unless all dependent repositories have been updated first.
3. **Define SemVer Ranges:** In `next.config.js`, you can specify `requiredVersion` (e.g., `^0.9.0`) instead of `false`. This allows Module Federation to emit a warning or fall back to loading separate instances if version compatibility is broken.

---

## Turn 26: User

cool. is there cool when i remove the singleton ?

---

## Turn 27: Assistant

If you remove the `singleton: true` configuration, Webpack Module Federation will allow the host and each remote MFE to load and run **their own separate instances** of the auth provider library. 

Generally, **it is not recommended (not "cool") to remove the singleton** unless you have fully migrated every single remote MFE to be 100% decoupled from the provider. 

Here is a detailed look at what will happen if you remove it:

---

### 1. The Bad Consequences of Removing the Singleton

#### ❌ Hook State Desynchronization (Broken Contexts)
If a remote MFE still imports hooks directly (like `const { token } = useAuth()`), it will bind to its **own local instance** of the React Context. 
* When the Host logs in or refreshes the token, the Host's context updates.
* The remote MFE's context **will not update**, causing the remote MFE to think the user is still logged out or has an expired token.

#### ❌ Redundant API Calls (Duplicate Hydration)
Each instance of the auth provider will run its own lifecycle logic. If a remote MFE mounts its own provider instance, it may independently call the `guard()` function or fetch user credentials, causing **duplicate API calls** (like `/auth/me` or `/validate-menu`) to your backend.

#### ❌ Bundle Size and Memory Overhead
Instead of downloading and executing the library once, the browser will download, parse, and keep multiple copies of the same library in memory (one for the host and one for each remote MFE).

---

### 2. When would it be "cool" to remove the singleton?

It is only safe to remove `singleton: true` if:
1. **100% Bridge Integration:** All remote MFEs have migrated entirely to the [AuthBridge](file://~/Development/qcash-ui/docs/ordent-to-bri-provider-migration.md#L120-L130) model and never call `useAuth()` or `useGlobalStore()` from `@ordentco/addons-auth-provider`.
2. **No Providers Mounted in Remotes:** Remote MFEs only export simple page components and do not wrap their components in their own `<AuthProvider>` trees.

---

### Summary Recommendation

Keep `singleton: true` active. If you are worried about version differences breaking other repositories, the safest approach is:

1. Keep `{ singleton: true, requiredVersion: false }` to guarantee a single instance is shared.
2. Ensure all remote repositories consume auth state using the **`window.__QCASH_AUTH_BRIDGE__`** data object rather than React hooks.

---

## Turn 28: User

oke thanks. how about to test in different version things for QA ?

---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 30: Assistant

For QA to verify different version scenarios, they can use the built-in diagnostic tools, local environments, and specific test cases.

Here is a guide on how QA can test version mismatches and verify auth compatibility:

---

### 1. Using the Built-in MFE Monitor UI

The project has a built-in diagnostics monitor designed specifically to help developers and QA track Module Federation status.

* **How to open it:** Press `Ctrl+Shift+Z` on the keyboard (in development or staging environments).
* **Where to look:** Go to the **Auth** tab (`AuthMonitor`).
* **What to verify:** 
  * Check the **Provider Versions By Repo** section. It queries the runtime share scope and displays exactly which version of the auth provider is bundled by each remote MFE (e.g., `@bri v0.10.5` vs `@ordentco v0.9.140`).
  * Verify that the **Active Runtime** card shows the correct provider source matching the active route (e.g., `BRI Provider` for `/v2/helpdesk` and `Ordentco Provider` for legacy pages).

---

### 2. QA Test Cases to Execute

When QA is validating a release containing a new provider version or a new MFE deployment, they should execute the following test paths:

#### Test Path A: Client-Side Navigation (SPA Flow)
1. Log in from the host login page.
2. Navigate dynamically (via sidebar/links) to a remote MFE that runs a different provider version.
3. **Assert:** The user session remains intact, username/company info is displayed correctly, and no page crashes or "Access Denied" screens occur.

#### Test Path B: Deep-Linking & Hard Refresh (Hydration Flow)
1. Navigate directly to a deep link inside a remote MFE (e.g., `/v2/helpdesk`).
2. Hard refresh the page (`Cmd+Shift+R` or `Ctrl+F5`).
3. **Assert:** The page successfully reads the token from `localStorage`, hydrates the correct authorities, and does not flash blank or crash.

#### Test Path C: Session Expiration & Logout
1. Open a page belonging to a remote MFE.
2. Click **Logout** or trigger session expiry.
3. **Assert:** The app redirects to the landing page, and all sensitive data (like `localStorage` tokens, `validateMenu`, and `window.__QCASH_AUTH_BRIDGE__` values) is wiped out.

---

### 3. Automated testing with Playwright / Cypress

QA automation engineers can test the state of the unified bridge directly by asserting on the global `window` object in their test scripts:

```javascript
// Cypress Example
describe('Auth Provider Integration', () => {
  it('should expose the correct unified auth bridge', () => {
    cy.visit('/v2/helpdesk');
    
    cy.window().then((win) => {
      // Assert that the bridge exists
      expect(win.__QCASH_AUTH_BRIDGE__).to.exist;
      
      // Assert that authorities are ready and active source matches
      expect(win.__QCASH_AUTH_BRIDGE__.authSource).to.equal('bri');
      expect(win.__QCASH_AUTH_BRIDGE__.isAuthoritiesReady).to.be.true;
      expect(win.__QCASH_AUTH_BRIDGE__.username).to.not.equal('Guest');
    });
  });
});
```

Using this approach, QA does not need to know the inner workings of different package versions—they only assert that the final unified contract is correct and matches what the remote components require.

---

## Turn 31: User

not between bri and ordent but between ordent v140 and other version

---

## Turn 32: Assistant

To specifically test version mismatches within the **Ordent provider** (e.g., `qcash-ui` running `@ordentco/addons-auth-provider` version `0.9.140` while a remote MFE is running version `0.9.139` or `0.9.141`), QA can use the following steps:

---

### Step 1: Open the MFE Monitor to Identify the "Winner"
Since Module Federation's `singleton: true` resolves different versions to a single instance at runtime, QA must check which version of the Ordent provider actually "won" the negotiation and is executing in the browser.

1. Press `Ctrl+Shift+Z` on the page to open the diagnostics monitor.
2. Go to the **Auth** tab.
3. Look at the **Provider Versions By Repo** list:
   * It will show the versions requested by each repository (e.g., `qcash-ui: @ordentco v0.9.140` and `remote-mfe: @ordentco v0.9.139`).
   * The version badge with a **highlighted border/ring** is the active singleton version running in memory.

---

### Step 2: Check the Browser Console for Version Warnings
When there is a version mismatch, Module Federation's runtime checks compatibility. Open the browser's developer console (`F12`) and check for warnings like:
```text
[Module Federation] Unsatisfied version of shared singleton @ordentco/addons-auth-provider...
```
* If you see this warning, it means a version mismatch occurred, but Module Federation resolved it using its fallback rules (loading the highest version available).
* If the app is functional despite the warning, the version resolution worked.

---

### Step 3: Run the Ordent Version Compatibility Checklist
If the host has resolved to a different version of the Ordent provider than the remote MFE, QA should run these checks:

| Check Item | Target | What to verify |
| :--- | :--- | :--- |
| **Auth Session Sharing** | Legacy / Ordent Routes | Ensure navigating from the host to the remote MFE preserves the active user session (no automatic logouts). |
| **Menu Permissions** | MFE Pages | Verify that the MFE is able to load the user's menus and sidebars. If the version mismatch changed the `menuData` structure, menus will fail to render. |
| **Role-based Actions** | MFE Buttons/Forms | Test buttons governed by roles (e.g. Approve, Reject, Maker/Checker actions) to ensure permissions/authorities are parsed correctly. |
| **Login Redirects** | Main Flow | Log out and log back in from the remote MFE route to ensure the new version redirects to the correct login/landing pages. |

---

### Step 4: Simulating Different Versions Locally (For Dev/QA Automation)
To test this in a controlled dev environment before pushing to staging:
1. In a local remote MFE repository, change the version of `@ordentco/addons-auth-provider` in `package.json` (e.g., change `0.9.140` to `0.9.139`).
2. Run the remote MFE dev server locally.
3. Load `qcash-ui` locally, press `Ctrl+Shift+Z`, go to **Settings**, and point the remote MFE URL to your local dev port.
4. Verify if the host or MFE throws any Javascript crashes (`TypeError`, `undefined is not a function`) when accessing auth hooks.

---

## Turn 33: User

please adjust the devtool to show which ordent or bri version runtime was active

---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **replace_file_content**: 

</details>



---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 41: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 42: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **multi_replace_file_content**: 

</details>



---

## Turn 43: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 44: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 45: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 46: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 47: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **command_status**: 

</details>



---

## Turn 48: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 49: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **multi_replace_file_content**: 

</details>



---

## Turn 50: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 51: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 52: Assistant

I have updated the **Auth Provider Diagnostics devtool** to display the active version of the `@ordentco/addons-auth-provider` or `@bri/addons-auth-provider` package running in the browser memory.

### Summary of Changes

1. **Active Version Resolver ([provider-versions.ts](file://~/Development/qcash-ui/components/federation/monitor/auth/provider-versions.ts)):**
   * Added a helper function `getActiveVersionForPkg(pkg: string): string | null` that checks the Module Federation runtime (`shareScopeMap`) to detect which specific version of the shared provider singleton has been resolved and loaded (`shared.loaded === true`).

2. **Diagnostics UI ([index.tsx](file://~/Development/qcash-ui/components/federation/monitor/auth/index.tsx)):**
   * Integrated the helper function into the `AuthMonitor` component.
   * Updated the **Active Runtime** diagnostic card to show the `Active Version` alongside the active `Auth Source` and `Module Name`.

3. **Unit & Integration Tests ([AuthMonitor.test.tsx](file://~/Development/qcash-ui/components/federation/monitor/auth/__tests__/AuthMonitor.test.tsx)):**
   * Mocked `getActiveVersionForPkg` in the test suite.
   * Added a new test case verifying that the active version element renders correctly when a version is resolved (e.g. `v0.10.5` or `v0.9.140`).
   * Ran the tests and verified they all pass successfully.

---

### How to Use This in QA
When QA opens the monitor panel (via `Ctrl+Shift+Z`), they will now see:

* **Auth Source:** `Ordentco Provider (Legacy)` or `BRI Provider`
* **Module Name:** `@ordentco/addons-auth-provider` or `@bri/addons-auth-provider`
* **Active Version:** `v0.9.140` (or whichever version won the singleton negotiation in memory)

This allows QA to immediately confirm which exact package version is running the user's active session, making it easy to identify version compatibility issues.

---

