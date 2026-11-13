# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-e6aedfe5b8dd67c1eba803767000f1ee6aac5fd29edeb5b31f7d53f01bf251b1`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:31:30`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

In QCash micro-frontends (MFEs), page readiness and authentication gating are handled through a layered approach coordinated between the **Host (Shell)** container and the individual **Remote MFEs**:

---

### 1. Authentication & Authorization Gating

* **Shell-Level Single Source of Truth (`AuthProvider` / `AuthContext`):**
  * The Host application owns session lifecycle management, token storage (JWT/refresh tokens), and base user identity.
  * Authentication state and tokens are shared with remotes either through **Webpack Module Federation shared scope** (as a singleton React context / store) or passed down as mounting props.

* **Route Guards (`ProtectedRoute` / `AuthGuard`):**
  * Top-level routing utilizes wrapper guard components (e.g., `<ProtectedRoute>`, `<AuthGuard>`) before resolving remote page routes.
  * If the session is invalid or expired, navigation is intercepted and redirected to the login flow or SSO gateway.

* **Permission & Role-Based Access Control (RBAC):**
  * Remotes consume permissions via hooks (e.g., `useAuth()`, `usePermissions()`).
  * Gating checks verify required feature flags, permissions, or user roles before rendering sensitive page contents; unauthorized states render an `AccessDenied` or fallback view.

---

### 2. Page Readiness & Lifecycle Management

* **Dynamic Code Splitting & Lazy Loading:**
  * Remote micro-frontend pages are loaded dynamically via `React.lazy()` and dynamic imports (`import('remoteApp/Page')`).

* **Suspense & Fallbacks:**
  * Remote mounts are enclosed within `<React.Suspense fallback={<PageSkeleton />}>` or dedicated loading layouts to provide visual continuity while remote JS chunks are fetched over the network.

* **Data & Bootstrap Readiness:**
  * Pages manage readiness states using initialization hooks or data-fetching libraries (e.g., React Query / SWR / custom `useInit` hooks).
  * A page signals readiness once configuration, mandatory reference data, and user permissions are resolved, transitioning from skeleton/spinner to the active UI.

* **Resilience via Error Boundaries:**
  * Remotes and route pages are wrapped in `<ErrorBoundary>` components to catch chunk-load failures, network timeouts, or remote unhandled exceptions without crashing the shell application.

---

