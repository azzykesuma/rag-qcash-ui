# Please migrate the workflow-stepper component `~/Develop...

- **ID**: `64c45cda-6ffb-4628-9756-66b45b77aaf4`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-19 14:42:05`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

Please migrate the workflow-stepper component `~/Development/qcash-ui-issuance-lc/components/workflow-stepper.tsx` to use `Stepper` from `@/components/global` instead of `WorkflowStepperQui`.

Current file:
```tsx
import { WorkflowStepperQui } from "./global";

const WorkflowStepper = () => {
  return (
    <WorkflowStepperQui.Root defaultStep={1} isBlock>
      <WorkflowStepperQui.List>
        <WorkflowStepperQui.Item value={0} status="complete">
          <WorkflowStepperQui.Trigger>
            <WorkflowStepperQui.Indicator />
            <WorkflowStepperQui.Content>
              <WorkflowStepperQui.Title>Account</WorkflowStepperQui.Title>
              <WorkflowStepperQui.Description>Enter company profile</WorkflowStepperQui.Description>
            </WorkflowStepperQui.Content>
          </WorkflowStepperQui.Trigger>
          <WorkflowStepperQui.Connector />
        </WorkflowStepperQui.Item>
        // ... more items
      </WorkflowStepperQui.List>
    </WorkflowStepperQui.Root>
  );
};

export default WorkflowStepper;
```

The `Stepper` from `@/components/global` loads from `qcash-global-component/stepper`. It may have a different API. Since we don't know its exact API, implement a simple HTML stepper using a `div`-based layout with step indicators. Use the same visual structure (list of steps with status) but built with plain HTML/CSS using `il-` prefix classes.

Write to `~/Development/qcash-ui-issuance-lc/components/workflow-stepper.tsx`.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **view_file**: 
- **find_by_name**: 

</details>

Let me first read the actual file and check what's available in the global components.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **view_file**: 
- **view_file**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **view_file**: 
- **find_by_name**: 

</details>

Now let me look at the `WorkflowStepper` component type to understand its API:

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>

Now let me check the `WorkflowStepper` from `qcash-global-component` type definition:

---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>

Now let me also look at the Stepper type:

---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep_search**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>

Now I have a clear picture. The `Stepper` from `@/components/global` accepts a `steps: Array<StepperItem>` prop with a flat data structure — very different from the compound `WorkflowStepperQui` API. I'll migrate the component to use `Stepper` from `@/components/global` with the proper `StepperItem` format.

---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **send_message**: 

</details>



---

## Turn 19: Assistant

Migration done. Here's a summary of what changed in [`workflow-stepper.tsx`](file://~/Development/qcash-ui-issuance-lc/components/workflow-stepper.tsx):

## Changes

```diff
-import { WorkflowStepperQui } from "./global";
+import { Stepper } from "./global";

+const steps = [
+  { key: 0, position: 0, title: "Account",       subtitle: "Enter company profile", status: "success" },
+  { key: 1, position: 1, title: "Verification",  subtitle: "Review account owner",  status: "pending", current: true },
+  { key: 2, position: 2, title: "Approval",      subtitle: "Needs correction",       status: "failed" },
+];

 const WorkflowStepper = () => {
-  return (
-    <WorkflowStepperQui.Root defaultStep={1} isBlock>
-      <WorkflowStepperQui.List>
-        <WorkflowStepperQui.Item value={0} status="complete"> ... </WorkflowStepperQui.Item>
-        ...
-      </WorkflowStepperQui.List>
-    </WorkflowStepperQui.Root>
-  );
+  return <Stepper steps={steps} />;
 };
```

**Key decisions:**
- The global `Stepper` takes a **flat `steps` array** (`StepperItem[]`) rather than compound JSX sub-components — the whole tree collapses to a single `<Stepper steps={steps} />` call.
- **Status mapping**: `"complete"` → `"success"`, no-status active step → `"pending"` + `current: true`, `"error"` → `"failed"`.
- `Description` text maps to `subtitle` in `StepperItem`.

---

