# Extract toast action docs (@explore subagent)

- **ID**: `ses_048b638d8ffetgKy8rhinjnXft`
- **Source Tool**: `opencode`
- **Date**: `2026-07-31 15:28:13`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Read/search the saved HTML output file at ~\.local\share\opencode\tool-output\tool_fb749519f0014xQXgnsNrfV2Nq. Extract exact toast code examples involving `actionProps`, `timeout`, `description`, and any copy/clipboard action. Also report the documented `RootProps` action shape if present. Do not edit any files. Return only the relevant snippets and conclusions.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 15: Assistant

```text
Set `timeout: 0` for persistent toasts.
```

```text
Toast Content Value
description | ReactNode | Optional description text
actionProps | RootProps | Action button props
```

```text
Toast Options (common)
description | ReactNode | - | Secondary text
actionProps | RootProps | - | Action button configuration
timeout | number | 4000 | Auto-close timeout in ms (`0` = persistent)
```

**Conclusions**

- No executable toast examples using `actionProps`, `description`, or `timeout` are embedded in the saved HTML.
- No toast-related copy/clipboard action, `navigator.clipboard`, or `writeText` example is present.
- `RootProps` is named as the `actionProps` type, but its action shape/fields are not documented in the file.

---

