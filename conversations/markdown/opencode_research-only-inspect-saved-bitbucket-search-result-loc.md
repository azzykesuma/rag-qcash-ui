# Find QUILA select API (@explore subagent)

- **ID**: `ses_015319ce1ffeh01qpJRa6e3gi4`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 15:33:42`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only. Inspect the saved Bitbucket search result ~\.local\share\opencode\tool-output\tool_feace24c4001e6mKZg7BE0udxQ using Grep/Read. Determine the qcash-ui-design-system select component's remote module path, compound components, and minimal controlled single-selection example (value/onChange or selectedKey/onSelectionChange). Return exact code/API facts with source file paths. Do not modify files.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **grep**: Tool: grep
- **read**: Tool: read

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

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

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 17: Assistant

- Remote module: `qui/components/select-field` — exposed as `./components/select-field` from `./src/components/select-field.ts`. Source: `qcash-ui-design-system/apps/producer/src/components/component-exposes.ts:48` (saved result: `~\.local\share\opencode\tool-output\tool_feace24c4001e6mKZg7BE0udxQ`).
- It is a React Aria `Select` wrapper (`SelectProps<T>`), so controlled single selection is `selectedKey` + `onSelectionChange`, not `value` + `onChange`. Source: `qcash-ui-design-system/apps/notebook/src/content/docs/forms/select-field.mdx:126,138`.
- Compound parts identified by its implementation files: `Group`, `Trigger`, `Value`, `Indicator`, `Popover`, `ClearButton`, `Prefix`, `Suffix`, `Separator`, `CurrencyItem`, and `SelectionSummary`. Implementation root: `qcash-ui-design-system/packages/components/src/components/select-field/index.ts`.

```tsx
const [selectedKey, setSelectedKey] = useState<string | null>(null);

<SelectField selectedKey={selectedKey} onSelectionChange={setSelectedKey}>
  <SelectField.Group>
    <SelectField.Trigger>
      <SelectField.Value />
      <SelectField.Indicator />
    </SelectField.Trigger>
  </SelectField.Group>

  <SelectField.Popover>
    <ListBox>
      <ListBox.Item id="one">One</ListBox.Item>
      <ListBox.Item id="two">Two</ListBox.Item>
    </ListBox>
  </SelectField.Popover>
</SelectField>
```

`ListBox` is shared rather than a `SelectField` part; the documentation explicitly says the field is built with React Aria `Select`, `Button`, `Popover`, and the shared `ListBox`.

---

