# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-93075effebfcc522edf09c505dcd4595d8944991151fe73dd27ed105adfef10c`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-02 09:19:16`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

Across QCash micro-frontends (MFEs), **page readiness** and **authentication gating** are handled hierarchically between the host container shell and the federated remote pages using standardized guards, contexts, and lifecycle hooks:

### 1. Authentication Gating & Route Protection
* **Host Shell / Root Auth Guard**: Authentication verification typically originates at the container shell via a top-level route guard (e.g., `AuthGuard`, `ProtectedRoute`, or `PrivateRoute`). Unauthenticated requests are intercepted and redirected to the login/SSO flow before remote micro-frontend bundles are mounted.
* **Shared Auth Context / Session State**: The authentication token and user session (profile, permissions, entitlements) are maintained globally in a shared context or session provider (often exposed as a shared React context or distributed singleton module via Module Federation). Remote pages consume this context to enforce route- or page-level authorization.
* **Role / Permission Guards (`PermissionGuard` / `RoleGuard`)**: Remote pages wrap sub-routes or sensitive action components with permission checks, verifying against the user's entitlements payload before rendering the underlying views.

### 2. Page Readiness & Lifecycle Synchronization
* **Suspense & Dynamic Remote Chunk Loading**: Federated pages are lazy-loaded via `React.lazy` and wrapped in `Suspense` with standardized skeleton or fallback spinners (e.g., `PageSkeleton` or `LoadingFallback`), preventing blank screens while chunks resolve over the network.
* **Initial Data & Readiness Hooks**: Remote pages typically employ a standardized initialization hook (such as `usePageInit` or component-level lifecycle hooks) that orchestrates:
  1. Verifying that shared session/auth context is populated.
  2. Fetching page-level bootstrap/lookup data.
  3. Setting an internal `isReady` / `loading` flag.
* **Error Boundaries**: Remote pages and federated entry points are isolated within `ErrorBoundary` components to ensure runtime readiness failures (e.g., remote chunk load timeouts or 401/403 authorization failures) gracefully show dedicated fallback or unauthorized screens without crashing the host container.

---

