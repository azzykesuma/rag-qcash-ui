# scripts/apply-patches.js usage explained

- **ID**: `opencode-adde4a438729aab85591b04d8fe8197236310549465b6f49e3ba6c0695190bc7`
- **Source Tool**: `opencode`
- **Date**: `2026-09-01 20:51:53`
- **Project**: `qcash-ui-account-receivable`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

can you explain to me the usage of scripts/apply-patches.js and the usage of it?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **read**: Tool: read

</details>

I’ll inspect the script and its package/config references, then summarize how to invoke it and what it changes.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

`scripts/apply-patches.js` applies a local fix directly to an installed dependency under `node_modules`.

**Why It Exists**

`@module-federation/typescript@3.1.3` checks `emitSkipped` after generating TypeScript declaration files:

```js
if (!emitSkipped) {
  return this.tsDefinitionFilesObj;
}
```

With `emitDeclarationOnly`, TypeScript can report `emitSkipped: true` even when declaration generation succeeded. This causes a false failure.

The patch changes the logic to check whether the diagnostics contain actual TypeScript errors:

```js
const hasErrors = diagnostics.some(
  d => d.category === typescript_1.default.DiagnosticCategory.Error,
);

if (!hasErrors) {
  return this.tsDefinitionFilesObj;
}
```

This modification is defined in `scripts/apply-patches.js:9-36`.

**How To Use It**

Normally, no manual action is needed. It runs automatically during dependency installation through the `prepare` script in `package.json:13`:

```json
"prepare": "husky || true && node -e \"require('fs').existsSync('scripts/apply-patches.js')&&require('child_process').execSync('node scripts/apply-patches.js',{stdio:'inherit'})\""
```

Therefore, this should apply it:

```bash
yarn install
```

You can also run it manually from the repository root:

```bash
node scripts/apply-patches.js
```

Expected output when applied:

```text
[apply-patches] Applied: Fix @module-federation/typescript emitSkipped false-positive with emitDeclarationOnly
```

Running it again is safe. The script detects the marker `DiagnosticCategory.Error` and reports:

```text
[apply-patches] Already applied: ...
```

**How It Works**

For every entry in the `patches` array, it:

1. Resolves the target path relative to the project root.
2. Skips the patch if the dependency file does not exist.
3. Checks `appliedMarker` to avoid applying it twice.
4. Searches for the exact original source block in `find`.
5. Replaces it with the block in `replace`.
6. Exits with code `1` if the dependency exists but the expected source cannot be found.

The last case usually means the dependency was upgraded and its implementation changed:

```text
[apply-patches] WARNING: Could not find target ... patch may be outdated
```

The patch would then need to be reviewed and updated.

**Important Notes**

- It modifies `node_modules`, not the dependency source committed to this repository.
- Reinstalling or deleting `node_modules` removes the modification, but `prepare` should reapply it.
- Installation with lifecycle scripts disabled, such as `yarn install --ignore-scripts`, will not apply it.
- When `NODE_ENV=production`, informational output is suppressed, but patching and failure handling still occur.
- The currently installed file in this workspace is not patched, suggesting `prepare` has not run since this script was added, or installation scripts were skipped. Running `node scripts/apply-patches.js` would apply it.

---

