# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-0489d90311b162fbe5bb402a78134769384b2865f2d0b23d6fa667abbf25018c`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:20:06`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

In QCash micro-frontend (MFE) architecture, page readiness and authentication gating are coordinated between the host container (shell) and remote MFEs through a centralized guard and lifecycle pattern:

---

### 1. Authentication & Permission Gating

* **Centralized Auth in Host / Shell:**
  * The host container owns the authentication lifecycle (SSO/OIDC token retrieval, token refresh, and session expiration).
  * Auth state and access tokens are distributed to remote micro-frontends via shared Module Federation singletons (e.g., a shared `AuthContext`, shared state store, or window-level auth bridge).
* **Route Guards (`ProtectedRoute` / `AuthGuard`):**
  * Top-level routing utilizes wrapper components (such as `ProtectedRoute` or `AuthGuard`) in React Router.
  * Checks token validity and user session before mounting remote pages. Unauthenticated sessions are redirected to the login/SSO flow.
* **Role & Permission Gating (`PermissionGuard` / `useAuth`):**
  * Granular access control is handled through permission components (e.g., `PermissionGuard`, `CanAccess`) or custom hooks (`useAuth()`, `usePermissions()`).
  * If a user lacks the necessary permission codes/roles for a micro-frontend view, access is blocked before initiating downstream API calls.

---

### 2. Page Readiness & Remote Loading

* **Asynchronous Chunk Resolution (`React.lazy` + `Suspense`):**
  * Remote micro-frontend page components are dynamically imported using `React.lazy(() => import('mfeRemote/Page'))`.
  * Routed views are wrapped in `React.Suspense` with a shared fallback UI (`PageLoader` or skeleton placeholders) while the remote entry script and page chunks download.
* **Initialization & Data Readiness Hooks:**
  * Remote pages use initialization hooks (e.g., `useInitPage`, `usePageReady`, or local `isLoading`/`isReady` states) to wait for essential prerequisites (dictionaries, master data, user profile) before rendering the interactive view.
  * Prevents flash-of-unauthenticated/empty-content (FOUC) and avoids race conditions in child components that rely on contextual data.
* **Resilience via `ErrorBoundary`:**
  * Remote imports and page trees are encapsulated within `ErrorBoundary` components to catch chunk load failures (e.g., network drops or remote deployment mismatches) and display graceful fallback/retry views without crashing the host container.

---

