# there is something wrong with the code. when i build the apk, downl...

- **Author**: `azzy`
- **ID**: `ses_0ba5bc17bfferqT3g5LafRoM9v`
- **Source Tool**: `opencode`
- **Date**: `2026-07-09 13:50:19`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

there is something wrong with the code. when i build the apk, downloaded it to my phone, it doesn't even open the app, it just stuck on the splash screen. can you investigate why this happened?

---

## Turn 2: User

there is something wrong with the code. when i build the apk, downloaded it to my phone, it doesn't even open the app, it just stuck on the splash screen. can you investigate why this happened?

---

## Turn 3: User

there is something wrong with the code. when i build the apk, downloaded it to my phone, it doesn't even open the app, it just stuck on the splash screen. can you investigate why this happened?

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **task**: Tool: task

</details>

I'll investigate the splash screen issue. Let me start by exploring the project structure and understanding the app's

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>

Now let me look at the key files that could cause the splash screen hang

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **bash**: Tool: bash
- **read**: Tool: read

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>

Let me dig deeper into the specific crash points.

---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Let me check the actual splash screen implementation to understand if the manual `preventAutoHideAsync` call could be conflicting with Expo Router's built-in management.

---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **read**: Tool: read

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash

</details>

Now let me check the Android native splash screen implementation to understand the counter mechanism.

---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

Now I have the full picture. Let me read the store code for the final piece of the puzzle.

---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>

I found the root cause. The native Android module reveals it clearly.

**Root cause:** `SplashScreen.preventAutoHideAsync()` in `_layout.tsx` (line 10) sets a flag `userControlledAutoHideEnabled = true` in the native `SplashScreenModule`. Expo Router's `internalMaybeHideAsync()` respects this flag and **refuses to hide the splash screen** when it's true, because it thinks you want to control it manually. But your code never calls `SplashScreen.hideAsync()` — so it stays forever.

---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>

Now let me implement the fixes.

**Fix 1:** Remove the conflicting `SplashScreen.preventAutoHideAsync()` call and add `hideAsync()` after auth is ready.

---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>

Now let me add a safety-net `hideAsync` after auth init and fix the `node-forge` lazy load.

---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>

I see the edit left a duplicate. Let me fix it.

---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>

Good. Now let me lazy-load `node-forge` to prevent potential module evaluation crashes in Hermes.

---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

Expo Router already exports an `ErrorBoundary`. Let me check how to use it.

---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write**: Tool: write

</details>

I'll create a simple Error Boundary for the root layout to prevent silent crashes.

---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>

Now wrap the root layout with the ErrorBoundary.

---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **edit**: Tool: edit

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

Let me verify the TypeScript config supports `require()` and check all final file states.

---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

Let me verify the `password-encryption.ts` compiles with the `require` approach by checking if it's properly typed.

---

## Turn 39: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

Let me check if `node-forge`'s module-level code is safe for Hermes.

---

## Turn 40: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>

Let me verify all the modified files are correct.

---

## Turn 41: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 42: Assistant

## Summary of changes

I identified **the root cause** and made 4 fixes:

### 🔴 Root Cause: `SplashScreen.preventAutoHideAsync()` in `src/app/_layout.tsx:10`

The native `SplashScreenModule` (`SplashScreenModule.kt:25-32`) has a flag `userControlledAutoHideEnabled` that gets set to `true` when you call the public `preventAutoHideAsync()`. Expo Router's internal `_internal_maybeHideAsync()` (`utils/splash.js:45-51`) checks this flag — and **refuses to hide the splash screen** when it's `true`, because it thinks YOU want to control the lifecycle yourself. Since your code never calls `hideAsync()`, the splash screen stays forever.

### Fixes applied:

1. **`src/app/_layout.tsx`** — Removed the manual `SplashScreen.preventAutoHideAsync()` call. Added `SplashScreen.hideAsync()` in `useEffect` as a safety net. Wrapped the tree in an ErrorBoundary.

2. **`src/components/error-boundary.tsx`** (new) — Catches any rendering error, hides the splash screen, and shows a "Something went wrong" fallback with a retry button. Prevents silent crashes.

3. **`src/lib/password-encryption.ts`** — Changed `node-forge` import from module-level (`import forge from 'node-forge'`) to lazy `require()` inside the functions. This prevents potential Hermes module-evaluation crashes during app startup — `node-forge` is only loaded when login/register is actually used.

4. **`.env`** — Added a comment explaining you need to set `EXPO_PUBLIC_API_BASE_URL` to your server's actual address for real devices (otherwise it defaults to the emulator-only `[REDACTED_IP]:8080`).

Rebuild the APK (`eas build --profile preview --platform android`) and the app should no longer get stuck on the splash screen.

---

