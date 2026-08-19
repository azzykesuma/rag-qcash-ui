# Duplicate API calls in local development

- **ID**: `ses_034701e76ffeIJ8ovm0p6DTA3P`
- **Source Tool**: `opencode`
- **Date**: `2026-08-04 13:57:12`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

[Image 1]can you explain why i hit this api twice in local developement?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **skill**: Tool: skill

</details>

I’ll trace the request caller and the local React/Next development settings to identify whether this is Strict Mode behavior or duplicate application logic.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **grep**: Tool: grep
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>

The trace confirms one logical caller, not two: `DashboardPage` starts the request in a mount effect, and its cleanup aborts it. The canceled first row followed by a successful second row matches React’s development-only effect remount check.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **context7_query-docs**: Tool: context7_query-docs

</details>



---

## Turn 9: Assistant

This is caused by React Strict Mode in local development:

- `next.config.ts:8` enables `reactStrictMode: true`.
- `src/containers/dashboard/index.tsx:52` calls the API inside `useEffect`.
- React development mode runs `setup → cleanup → setup` to detect unsafe effects.
- Cleanup at `src/containers/dashboard/index.tsx:85` aborts the first request, explaining `(canceled)`.
- The second request is the real one and returns `200`.

Production runs the effect once. Your cleanup is correct, so disabling Strict Mode is unnecessary. Note that cancellation happens client-side after dispatch, so the backend may still briefly receive the first request.

---

