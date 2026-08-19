# Centralized Auth Guard, Module Federation Singleton & Loading Architecture

**Author / Context**: Antigravity Architecture & Pair Programming Session  
**Date**: August 19, 2026  
**Scope**: `qcash-ui` (Host Shell) & 40+ Remote Micro-Frontends (`qcash-ui-account-receivable`, `qcash-ui-fund-transfer`, etc.)

---

## 1. Executive Summary & Root Problem

In a Webpack Module Federation micro-frontend architecture with Next.js, remote feature repositories experienced recurring production and development issues:
1. **`Uncaught Error: Function not implemented`**: Triggered when remote pages executed `guard()` inside `useEffect` while `@ordentco/addons-auth-provider` was unshared (`singleton: false`) or differed in version from the Host.
2. **Infinite Loading / Stale Sessions**: Remotes instantiated isolated React Contexts with `isAuthoritiesReady: false` when not configured with `singleton: true`.
3. **Double Loader Flickering**: Attempting to block rendering at the Host `_app.tsx` root caused flashes between global loading overlays and native page skeletons.
4. **Duplicate API Storms**: Multiple remote components firing `/auth/me` and `/menu/me` on every client-side route navigation.

---

## 2. The Architectural Solution

### Pillar 1: Host-Level Centralized Guard (`HostAuthGate`)
* **Location**: `qcash-ui/components/providers/HostAuthGate.tsx`
* **Responsibilities**:
  - Synchronizes `access-token` from `localStorage` into `briAuth` and `ordentAuth` contexts.
  - Executes throttled `ordentAuth.guard()` (minimum 3-second throttle per route/token) on session or route changes.
  - Automatically handles `401 Unauthorized` responses by purging `access-token` / `refresh-token`, resetting the bridge, and redirecting to `/landing-page`.
  - Directly returns `<>{children}</>` to avoid unmounting page skeletons or introducing layout flashes.

### Pillar 2: Micro-Frontend Context Mirroring (`AuthBridgeSync`)
* **Location**: `qcash-ui/components/providers/AuthBridgeSync.ts`
* **Responsibilities**:
  - Replicates state from `ordentAuth` into `briAuth` on non-BRI routes.
  - Exposes `window.__QCASH_AUTH_BRIDGE__` for legacy or non-singleton remotes.
  - Synchronizes `productMenu`, `productRoles`, and `validateMenu` to `localStorage` on localhost.
  - Resets all bridge caches upon receiving `qc-bridge-sync` logout events.

### Pillar 3: Remote Pages as Pure Data Consumers
* **Location**: `pages/**/*.tsx` across all 40+ remote repositories.
* **Rules**:
  - 🚫 **NEVER call `guard()`** inside remote pages.
  - Remotes only read auth state via `const { userType, companyID, productAuthorities, isAuthoritiesReady } = useAuth()`.
  - Always use optional chaining on authorities (e.g. `productAuthorities?.SWIFT?.dataEntry`).

### Pillar 4: Standalone Development Gate (`StandaloneAuthGate`)
* **Location**: Remote `pages/_app.tsx`
* **Pattern**:
```tsx
function StandaloneAuthGate({ children }: { children: ReactNode }) {
  const { guard, token } = useAuth();

  useEffect(() => {
    if (token && typeof guard === "function") {
      try {
        const res = guard();
        if (res && typeof (res as any).catch === "function") {
          (res as any).catch(() => {});
        }
      } catch {}
    }
  }, [token, guard]);

  return <>{children}</>;
}
```

---

## 3. Module Federation Singleton Matrix & 2-Phase Rollout

| Phase | Remote `next.config.js` | Remote `pages/**/*.tsx` | Status | Notes |
| :---: | :---: | :---: | :---: | :--- |
| **Phase 1 (Hotfix)** | `singleton: true` | Still calls `guard()` | 🟢 **Safe** | Immediately stops `"Function not implemented"` crashes without touching page files. |
| **Phase 2 (Cleanup)** | `singleton: true` | Removed `guard()` | 🚀 **Optimal** | Zero redundant network calls, clean maintainable code. |

### Required `next.config.js` Configuration:
```js
shared: {
  "@ordentco/addons-auth-provider": {
    singleton: true,
    requiredVersion: false,
  },
  "@bri/addons-auth-provider": {
    singleton: true,
    requiredVersion: false,
  },
}
```

---

## 4. Loading Overlay Best Practice

* **Auth Loading**: Handled transparently by the Host Shell in the background.
* **Chunk Download Fallback**: Provided via Host `dynamic(..., { loading: () => <LoadingOverlay /> })`.
* **Business / Table Loading**: Managed locally in the feature container via table skeletons or local spinners.
