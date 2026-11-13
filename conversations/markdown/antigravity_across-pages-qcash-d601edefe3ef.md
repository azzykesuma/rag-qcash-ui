# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-c6c409e7b28feca590d0ec039eda726357ab49c61188355e1dbbd601edefe3ef`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:21:19`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

How is page readiness and authentication gating handled across pages in QCash micro-frontends?

---

## Turn 2: Assistant

In QCash micro-frontend (MFE) architecture, page readiness and authentication gating are typically handled through a coordinated shell-and-remote pattern using Webpack/Vite Module Federation:

---

### 1. Centralized Authentication Gating (Shell & Shared Context)

* **Shell-Level Gatekeeping:**  
  Primary authentication validation (SSO, JWT/OAuth tokens, session cookies) is handled by the Shell (host) application before or during route resolution. Unauthenticated users are redirected to the identity provider/login page before protected remote bundles are loaded.
* **Shared `AuthProvider` Singleton:**  
  The authentication state is maintained via a shared `AuthContext`/`AuthProvider` configured in `module-federation.config.js` as a shared singleton (`singleton: true`). This ensures remote MFEs consume the exact same React context and token state as the shell, preventing duplicated auth instances or desynchronized session states.
* **Route Guards (`ProtectedRoute` / `AuthGuard`):**  
  Routes are wrapped in guard components (e.g., `<ProtectedRoute>`, `<AuthGuard>`, or role-based permission wrappers). The guard evaluates:
  1. `isAuthenticated`: Redirects unauthenticated requests to login.
  2. `permissions` / `roles` (RBAC): Evaluates feature flags or permissions before rendering the target page; unauthorized access routes to a `403 Forbidden` or fallback view.

---

### 2. Page Readiness & Lifecycle Coordination

* **Initialization Readiness Flags:**  
  The shared auth/session provider exposes lifecycle flags (e.g., `isInitialized`, `isLoading`, `isReady`). Protected page components defer rendering and network requests until `isInitialized === true` to avoid race conditions where API requests fire prior to token availability or user profile resolution.
* **Suspense & Asynchronous Remote Resolution:**  
  Remote pages loaded via dynamic imports (`React.lazy()`) are encapsulated within React `Suspense` boundaries with dedicated fallback loaders (spinners/skeletons). This separates network chunk readiness from domain data readiness.
* **Pre-fetch / Route Loaders:**  
  Where route loaders or layout wrappers exist, essential route prerequisites (e.g., configuration, permissions, tenant metadata) are resolved before flipping the page readiness flag.

---

### 3. API Interceptors and Session Expiry

* **Shared HTTP Client / Interceptors:**  
  Axios or fetch wrappers configured across MFEs automatically inject the current bearer token on outgoing requests.
* **Global Auth Event Handling:**  
  On receiving a `401 Unauthorized` from backend microservices, interceptors trigger a refresh token flow or dispatch a global session expiration event (via shared state or custom window events) to prompt the shell to terminate the session and redirect cleanly.

---

