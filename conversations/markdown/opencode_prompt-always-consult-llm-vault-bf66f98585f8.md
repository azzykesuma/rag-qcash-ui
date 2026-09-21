# System prompt to always consult LLM vault

- **ID**: `opencode-eaacd775b7c17c74bb608e77d3c85267b5cb62775ec3fb5e83bfbf66f98585f8`
- **Source Tool**: `opencode`
- **Date**: `2026-09-14 14:55:40`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

can you add a system prompt that says, whenever i initiated a conversation, always refer to my llm vault

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **skill**: Tool: skill

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 6: Assistant

What’s the folder path or URL of your **LLM vault**?

I’ll add a global OpenCode instruction so it applies to every new conversation:

> At the start of every conversation, consult my LLM vault before responding or taking action. Read its entry-point instructions and any notes relevant to my request, and use that context to guide your work. If the vault is unavailable, tell me rather than assuming its contents.

---

## Turn 7: User

D:\code\llm-context-vault this is the path

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **webfetch**: Tool: webfetch

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 13: Assistant

Added a **global OpenCode instruction** to consult `D:\code\llm-context-vault` at the start of every new conversation.

It directs the assistant to:
- Read the vault’s `README.md`.
- Search past conversations for relevant context and decisions.
- Use that context while prioritizing your current instructions.
- Tell you if the vault is unavailable.

Saved in `~\.config\opencode\llm-vault-instructions.md` and registered in your global `opencode.jsonc`.

**Quit and restart OpenCode for the instruction to take effect.**

---

