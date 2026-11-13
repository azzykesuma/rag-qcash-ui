# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-1547af43ab56f3801db11fe85141089f2ffe0ba79929b513d4ee2ffa7a09a3ee`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 09:29:41`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

In QCash micro-frontends (MFEs), **page readiness** and **authentication gating** follow a layered architecture distributed between the host/shell container and the federated remote micro-frontends:

---

### 1. Authentication Gating

* **Shell-Level Auth Provider & Session Management:**
  * Authentication (tokens, user profile, and RBAC permissions) is managed centrally in the **Host/Shell container** (e.g., via an `AuthProvider` or root store).
  * Remote MFEs consume auth context either as a federated shared module (single instance) or passed down via shared context/props to prevent multiple conflicting session states.
* **Route Protection (`ProtectedRoute` / `AuthGuard`):**
  * Top-level and sub-routes are wrapped in a route guard (e.g., `AuthGuard` or `ProtectedRoute`).
  * Gating checks for active authentication tokens before rendering child routes. If unauthenticated or expired, it redirects to the central login/SSO portal.
* **Permission / RBAC Gating:**
  * Pages or specific sub-views check user permissions/roles (e.g., `usePermissions` or `<PermissionGate permissions={[...]} />`).
  * If the user lacks the required permission, a standardized `403 Forbidden` / `Unauthorized` fallback is rendered.

---

### 2. Page Readiness & Remote Lifecycle

* **Federated Dynamic Imports & Suspense:**
  * Remote micro-frontend entry points are loaded lazily using `React.lazy(() => import('mfe_name/Page'))`.
  * Page loading readiness is handled by wrapping dynamic remote routes in `<Suspense fallback={<PageSkeleton />}>` (or shared `LoadingSpinner`).
* **Error Boundaries (`RemoteErrorBoundary`):**
  * To prevent one failing remote from breaking the entire application shell, remote page chunks are wrapped in dedicated `ErrorBoundary` components that catch chunk loading errors, network failures, or remote runtime exceptions.
* **Data Initialization & Readiness Checks:**
  * Remote pages typically manage their internal readiness state using standard queries (e.g., TanStack/React Query) and an `isReady` or `isLoading` flag:
    * Configuration, user session, and initial reference data are validated before transitioning out of the initial skeleton state.
* **Shared Singleton Dependencies:**
  * Core libraries (`react`, `react-dom`, routing libraries, and query clients) are configured as **shared singletons** in the Module Federation configuration to ensure synchronization of auth state and avoid multiple runtime contexts across federated boundaries.

---

