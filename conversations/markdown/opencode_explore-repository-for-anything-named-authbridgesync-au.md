# Trace auth bridge sync (@explore subagent)

- **ID**: `ses_0059c1dbdffej40FwVbMcxZt7l`
- **Source Tool**: `opencode`
- **Date**: `2026-08-13 16:11:17`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore the repository for anything named authbridgesync, AuthBridgeSync, auth bridge sync, or close variants. Assess the implementation's correctness by tracing definitions, callers, data flow, lifecycle/concurrency behavior, and tests. Do not edit files. Return concise findings ordered by severity, with exact file paths and line numbers, plus any testing gaps. Thoroughness: very thorough.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **grep**: Tool: grep
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 8: Assistant

## Findings

### High

1. **Logout/null-token synchronization preserves stale credentials**
   - `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:8-13`
   - `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:192`
   - The bridge writes non-null tokens but deliberately does nothing when the host publishes `token: null`. `useMenu` then falls back to the old `access-token`, retaining the previous user’s menu and making requests with stale credentials.
   - The claimed interceptor cleanup is incomplete: `~/Development/qcash-ui-header-footer/services/custom-axios.ts:50-55,74-89` only clears tokens for particular error/navigation paths, not ordinary host logout or expiry events.

2. **A present but partial bridge is accepted as valid and can crash consumers**
   - `~/Development/qcash-ui-header-footer/layouts/default/AuthBridge.tsx:84-99`
   - `~/Development/qcash-ui-header-footer/layouts/default/container/index.tsx:37-39`
   - `getAuthBridge()` neither validates nor merges host data with `FALLBACK`; any truthy value, including `{}`, activates bridge mode. Required properties can therefore be undefined.
   - Immediate crash examples:
     - `~/Development/qcash-ui-header-footer/hooks/useLocalTaxMenu.tsx:17,25` calls `authorities.size` and `.get`.
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcExportModal.tsx:18,38,61`
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LocalTaxesModal.tsx:14,51`
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/GlobalExternalTransferModal.tsx:29,104,116`
     - These index `productAuthorities` without guarding the object itself.

### Medium

3. **Bridge updates can be lost between render and effect subscription**
   - `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:5,7-27`
   - State is read during render, but the event listener is installed later. If the host replaces the bridge and emits `qc-bridge-sync` in that interval, the mount effect reads the new bridge only to copy its token; it does not update `data`. Consumers remain stale until another event occurs.
   - This can also leave bridge selection stale at `~/Development/qcash-ui-header-footer/layouts/default/container/index.tsx:37-39`.

4. **Advertised legacy-auth fallback is applied inconsistently**
   - Layout and LC Import explicitly fall back to Ordent auth:
     - `~/Development/qcash-ui-header-footer/layouts/default/container/index.tsx:34-39`
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:41-48`
   - Other consumers always use bridge fallback data:
     - `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:184-192`
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/MenuModalProvider.tsx:38`
     - The other authority modals listed above.
   - With no bridge, `userType` remains null, menu loading can remain true (`useMenu.tsx:305`), authority-driven options are disabled, and leave-state actions become fallback no-ops despite a working legacy provider.

5. **Shallow snapshots do not safely support in-place host mutations**
   - `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:19-23`
   - The outer object is copied, but nested `authorities`, `products`, and similar data retain their identity. Consumers memoize by those references:
     - `~/Development/qcash-ui-header-footer/layouts/default/container/menu/components/LcImportModal.tsx:82-95`
     - `~/Development/qcash-ui-header-footer/hooks/useMenu.tsx:251-273`
   - If the host mutates nested bridge collections in place and emits the event, components rerender but these memoized computations can remain stale. Correctness currently depends on an undocumented immutable-update contract.

6. **Each consumer creates an independent external-store subscription**
   - `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:4-28`
   - There are nine production callers. Each maintains separate state and listener, causing repeated token writes and snapshots for every event. React event batching reduces—but does not formally eliminate—cross-consumer tearing and mounting races. A single provider or `useSyncExternalStore` subscription would better model this global mutable source.

### Low

7. **Event constant is unused, allowing protocol-name drift**
   - Defined at `~/Development/qcash-ui-header-footer/layouts/default/AuthBridge.tsx:4`
   - Listener hard-codes the same string at `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:26-27`.

8. **Possible SSR hydration mismatch**
   - `~/Development/qcash-ui-header-footer/layouts/default/AuthBridge.tsx:85`
   - `~/Development/qcash-ui-header-footer/hooks/useAuthBridge.ts:5`
   - SSR always renders fallback data, while the hydration initializer can immediately read host data. If this federated layout is server-rendered and the bridge exists before hydration, header/menu output can disagree with server markup.

## Testing gaps

- No direct test file exists for `AuthBridge.tsx` or `useAuthBridge.ts`; coverage marks the token-write path uncovered.
- Existing consumer tests mostly mock `useAuthBridge`, so they do not exercise actual listener registration, event delivery, rehydration, cleanup, or localStorage flow.
- Missing cases:
  - Initial bridge read and array-to-`Map` conversion.
  - Event-driven updates and listener removal on unmount.
  - Null-token logout and token rotation.
  - Event occurring between render and effect installation.
  - Partial/malformed bridge objects.
  - In-place versus immutable nested updates.
  - Multiple simultaneous subscribers and Strict Mode mount/unmount cycles.
  - Real no-bridge legacy-provider behavior.
  - SSR/hydration behavior.
  - End-to-end proof that requests use the newly bridged token and never a prior user’s token.

---

