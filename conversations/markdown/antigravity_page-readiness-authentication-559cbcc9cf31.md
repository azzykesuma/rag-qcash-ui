# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-f156f27df921d6b2944c738c0c194e60fb9de7bd52ae536f7fc9559cbcc9cf31`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:27:50`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

In QCash micro-frontends (MFEs), page readiness and authentication gating are structured through a coordinated shell-and-remote architecture using the following patterns:

---

### 1. Authentication Gating
* **Host/Shell-Level Auth Control**: 
  * The root container (shell) manages primary session lifecycle, token storage (JWT/refresh tokens), and SSO state.
  * Auth state is passed to remote MFEs via a shared context/singleton store (e.g., shared React Context or Zustand/Redux instance via Module Federation shared scope) or custom cross-application events (`window.dispatchEvent`).
* **Route Guards (`AuthGuard` / `ProtectedRoute`)**:
  * Pages or MFE sub-routes are wrapped in authentication guards (e.g., `<ProtectedRoute>` or `<AuthGuard>`).
  * The guard verifies token existence, expiration, and user role/permission claims before allowing the child component tree to mount.
  * If unauthenticated or unauthorized, the guard redirects the user to the central login route or triggers a shell-level redirect.
* **HTTP / Interceptor-Level Gating**:
  * Shared API clients (e.g., configured Axios instances) attach authentication headers and intercept `401 Unauthorized` / `403 Forbidden` responses to invalidate the session, trigger token refresh flows, or broadcast logout events to all mounted MFEs.

---

### 2. Page Readiness & Lifecycle
* **Module Federation Suspense & Lazy Loading**:
  * Remote MFE entry points and page chunks are dynamically imported using `React.lazy()` and wrapped in `<Suspense>` boundaries.
  * Fallback placeholders (skeleton screens or loaders) are displayed while the remote bundle is fetched and parsed.
* **Initialization & Pre-flight Hooks (`usePageReady` / `useAuth`)**:
  * Pages use initialization hooks or state checks to verify that:
    1. Remote dependencies and shared singletons (e.g., React, routing, store) are resolved.
    2. User permissions and session context have populated.
    3. Mandatory initial API data (e.g., via React Query / RTK Query) has completed its first fetch without critical errors.
* **Error Boundaries**:
  * Remote boundaries (`<ErrorBoundary>`) catch bundle loading failures (e.g., network timeout fetching remote entry points) or runtime render crashes, providing retry mechanisms or fallback UI without crashing the host container.

---

