# 🧠 LLM Context Vault: Benchmark Report

**Date:** 2026-10-15 11:53:01  
**Runner Backend:** `OpenCode (deepseek/deepseek-flash)`  
**Tasks Evaluated:** 5 (Successful: 4)  

---

## 📊 Executive Scorecard

| Metric | Baseline (Arm A) | Vault-Assisted (Arm B) | Improvement / Impact |
| :--- | :--- | :--- | :--- |
| **Average Tokens / Task** | `369` tokens | `1720` tokens | **+366.3%** tokens |
| **Average Wall-Clock Time** | `31112 ms` | `7734 ms` | **4.02x speedup** |
| **Fact / Keyword Recall** | `36.2%` | `71.2%` | **+96.6% accuracy** |

> [!NOTE]
> **Exploration Savings Context:** In interactive agent workflows, a baseline agent typically requires 5–15 roundtrip tool calls (`grep`, `view_file`, `read`) consuming **15,000–50,000+ tokens** to locate institutional knowledge. Injecting a bounded vault excerpt upfront eliminates these exploratory turns.

---

## 📋 Per-Task Comparison Table

| Task ID | Project | FTS Latency | Baseline Tokens | Vault Tokens | Baseline Acc | Vault Acc | Δ Score |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `auth-gate-architecture` ✅ | `qcash-ui` | `79 ms` | `385` | `2078` | `0%` | `50%` | `+50%` |
| `vault-scan-optimization` ✅ | `llm-context-vault` | `69 ms` | `395` | `2286` | `20%` | `60%` | `+40%` |
| `sanitization-path-policy` ✅ | `llm-context-vault` | `127 ms` | `264` | `2031` | `25%` | `75%` | `+50%` |
| `overengineering-dead-assets` ⚠️ Error | `qcash-ui-dashboard-auto-collection` | `84 ms` | `0` | `1932` | `0%` | `100%` | `+100%` |
| `negative-control-novel-task` ✅ | `general` | `63 ms` | `432` | `487` | `100%` | `100%` | `+0%` |

---

## 🔍 Detailed Task Breakdown

### Task 1: `auth-gate-architecture`

**Query:** *"How is page readiness and authentication gating handled across pages in QCash micro-frontends?"*  
**Ground Truth:** `HostAuthGate owns /auth/me and /menu/me. Pages should control rendering readiness through isAuthoritiesReady rather than stale local cache or local login strings.`  

- **Vault SQLite FTS Retrieval:** `79 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `385` (duration `51234 ms`), Vault-Assisted = `2078` (duration `4113 ms`)
- **Matched Keywords:** Baseline = `` | Vault = `HostAuthGate, isAuthoritiesReady`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
QCash centralizes both concerns rather than re-implementing them in every remote page.

- **Auth gating**: The shell owns auth via `AuthProvider`/`useAuth`, and routes are wrapped with `AuthGate`/`ProtectedRoute`. It waits for auth resolution (`isLoading`), then either renders the page or redirects unauthenticated users to `/login` with a return URL. Public pages opt out with `requiresAuth={false}`.
- **Page readiness**: The shell also provides `PageReadyProvider`/`usePageReady`. Each page uses the shared `Page`/`PageReadyBoundary` wrapper, which resets readiness on route change and marks the page ready after mount/data settle. The shell shows `PageLoader`/`PageSkeleton` until readiness is true.
- **Across MFEs**: Remote pages consume the same shared auth and readiness context instead of owning guards. The shared `Page` component composes both concerns, e.g. `AuthGate` around `PageReadyBoundary`, so `AccountsPage`, `TransferPage`, `CardsPage`, etc. behave uniformly. Standalone MFE dev uses the same providers with a mock auth adapter.

Key pieces: `apps/shell/src/App.tsx` (providers/router), `ProtectedRoute`/`AuthGate` (auth), `PageReadyProvider`/`usePageReady` (readiness), and the shared `Page` wrapper (per-page composition).
```

#### Vault-Assisted Response (Arm B):
```
## Page readiness and auth gating in QCash micro-frontends

The flow is centralized rather than spread across pages:

**Entry point — `pages/_app.tsx`**
- `HostAuthGate` is mounted **inside both auth providers (Ordent and BRI) and above federated children**, so every page — host or federated — passes through it.

**`components/providers/HostAuthGate.tsx`** (the centralized gate)
- **Readiness blocking:** protected children are not rendered until the required auth provider(s) report ready (e.g. `isAuthoritiesReady`).
- **Public-route bypass:** public routes skip gating entirely.
- **Guard execution:** runs the Ordent `guard()` centrally, plus the BRI guard only on BRI-owned routes. Direct executable `guard()` calls were removed from `components/onboarding-tour/notif-onboarding-tour.tsx` and `hooks/use-modal-session-expired.tsx`.
- **Per-token validation:** validation is tracked per token, with handling for token replacement during an in-flight validation, so a rotated token doesn't inherit a stale result.
- **Failure handling:** shows a retryable auth error, and on 401 clears session/provider tokens and redirects.
- **Sync awareness:** listens for same-tab `"qc-bridge-sync"` custom events and cross-tab `storage` events (event-driven, no polling).

**`components/providers/AuthBridgeSync.ts`**
- Complements the gate by propagating token changes across the federation boundary via `"qc-bridge-sync"`. It is intentionally left untouched.

**Why it's centralized:** federated modules 
[...truncated]
```
</details>

### Task 2: `vault-scan-optimization`

**Query:** *"How does llm-context-vault optimize repeat scans and avoid re-processing unchanged sessions?"*  
**Ground Truth:** `Uses an incremental cache to skip unchanged file-based sessions, limits changed sessions per batch (default 10), and defers the repository-wide privacy audit for bounded scans.`  

- **Vault SQLite FTS Retrieval:** `69 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `395` (duration `25960 ms`), Vault-Assisted = `2286` (duration `4660 ms`)
- **Matched Keywords:** Baseline = `unchanged` | Vault = `unchanged, audit, bounded`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
llm-context-vault is incremental by default. Its scanner keeps a persisted scan manifest/state store (keyed by session identity) rather than re-parsing every source on each run.

Typical flow:

1. **Cheap stat pass** — it discovers source session files and compares each file’s `mtime` and `size` against the cached entry. If both match, the session is skipped before any parsing.
2. **Content-hash fallback** — if metadata changed, it computes a content hash (e.g. SHA-256). If the hash still matches the manifest, it refreshes metadata but still avoids re-processing.
3. **Append-only cursor** — for JSONL/transcript-style session logs, it stores a byte offset or line cursor, so growing sessions only need the newly appended tail read and parsed.
4. **Stable identity + dedupe** — sessions are tracked by a stable provider/session ID, so renames/moves don’t force reprocessing and duplicate exports can be skipped.
5. **Deletion reconciliation** — missing source files are pruned from the manifest/index instead of triggering full rescans.
6. **Explicit full rebuild** — `--force`/`--full` (or equivalent) bypasses the cache when a complete reprocess is required.

Net effect: unchanged sessions never reach the parser/indexer/embedder; only new or genuinely changed content is processed.
```

#### Vault-Assisted Response (Arm B):
```
Short answer: in the material I can see, that optimization is **proposed, not confirmed-implemented**. The reference session is an analysis where an assistant inspected the code and recommended changes; there's no evidence in the excerpt that the incremental-scan design was actually merged. So I'd frame this as "the intended design," and you should verify against the current `pkg/vault` code before relying on it.

## What the analysis found as the current bottleneck

- Every conversation import triggers a full read/parse/rewrite of `dataset.jsonl` (`pkg/exporter/jsonl.go:55`), so N sessions means N passes over the master dataset.
- Repeat scans re-extract, re-sanitize, re-audit, and re-export unchanged sessions via the scan pipeline in `pkg/vault/vault.go:358`.
- OpenCode extraction issues a separate parts query per message (`pkg/extractor/opencode.go:124`).
- Antigravity assigns `time.Now()` on every extraction, causing output churn even when content is unchanged.

## The proposed incremental-scan mechanism (section B of the recommendation)

The design centers on a **local, Git-ignored scan-state cache** that records:

1. **Source fingerprints** — for file-based sessions, source size + modification time, escalating to content hashes when needed; for OpenCode, verified database revision fields including message/part changes.
2. **Processing versions** — so a change to sanitizer settings, custom redaction words, or extractor versions forces reprocessing.
3. **Corresponding exp
[...truncated]
```
</details>

### Task 3: `sanitization-path-policy`

**Query:** *"What is the vault sanitization policy for Windows and Unix machine file paths before publishing?"*  
**Ground Truth:** `Strips Windows paths (C:\Users\username\...) and Unix paths (/home/username/...) replacing them with ~/...`  

- **Vault SQLite FTS Retrieval:** `127 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `264` (duration `31355 ms`), Vault-Assisted = `2031` (duration `6801 ms`)
- **Matched Keywords:** Baseline = `path` | Vault = `C:\Users, redact, path`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
The `Vault` class explicitly **excludes machine file paths from sanitization**. In `Vault.Add` (e.g., `src/Agent.Worker/Vault.cs` / `src/Runner.Common/Vault.cs`), before a value is registered with the secret masker, it is checked by `IsMachineFilepath`:

- **Windows:** drive-letter absolute paths such as `C:\...`, `D:\...` (case-insensitive).
- **Unix:** absolute paths beginning with `/`, such as `/usr/bin/...`, `/home/...`.

If the value matches either pattern, it is **not added to the vault**, so it is **not replaced/masked by `Vault.Sanitize` when logs are published**. Only non-machine-path secrets (and values of length ≥ 3) are registered and sanitized. The rationale is to avoid masking common, non-sensitive machine paths and making logs unreadable.
```

#### Vault-Assisted Response (Arm B):
```
I don't have the sanitizer source in this context, so I can only summarize what the cited review findings establish — treat the exact regex/placeholder names as unverified.

**Where the policy lives**
- Redaction is a regex-based pass applied to conversation and message content **and** metadata at ingestion, before the vault writes exports. Source IDs are derived *after* sanitization (`pkg/vault/vault.go:245-262`), so they're stable against the sanitized text.
- Existing dataset records are canonicalized through the same pipeline in `pkg/vault/naming.go:91-110`.

**Path rules (per the audit's regression-fixture requirement)**
The sanitizer is expected to redact machine-specific absolute paths for both platforms, alongside credentials, emails, IPs, and custom terms:
- **Windows:** drive-letter paths (`C:\Users\…\…`) and backslash/UNC-style paths.
- **Unix:** absolute home/system paths (`/home/<user>/…`, `/Users/<user>/…`, `/var/…`).
- Matches are replaced with a placeholder token rather than removed, so the surrounding context survives.

**Important caveat flagged in the review**
That policy is *not* retroactively enforced on legacy state. If scan state is absent (first scan after upgrade, or `.vault/` deleted), existing records are canonicalized **without** the current sanitizer, and without a prior ownership entry retirement is never scheduled (`pkg/vault/scan.go:353-357`). Result: stale unredacted exports remain alongside new sanitized ones when `--redact-words` (or added p
[...truncated]
```
</details>

### Task 4: `overengineering-dead-assets`

**Query:** *"What dead scripts and unnecessary assets were identified for removal during the overengineering audit?"*  
**Ground Truth:** `scripts/analyze-bundle.js, scripts/update-version-file.js, unused asset directories (files, flags, fonts, icons, images, img), and version file.`  

- **Vault SQLite FTS Retrieval:** `84 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `0` (duration `60000 ms`), Vault-Assisted = `1932` (duration `2468 ms`)
- **Matched Keywords:** Baseline = `` | Vault = `analyze-bundle.js, update-version-file.js, assets, audit`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```

```

#### Vault-Assisted Response (Arm B):
```
The audit flagged these as **ignored/untracked workspace debris** (so they don't affect Git repository size, only local workspace cleanup):

## Dead scripts
- **`scripts/analyze-bundle.js`** — 37 lines. No package script, no analyzer dependency, no matching Next analyzer configuration.
- **`scripts/update-version-file.js`** — 9 lines. No caller, and its write operation is commented out.

## Stale tooling/config files
- **`.eslintrc.json`** — 3 lines. Biome is used and ESLint is not installed.
- **`todo.md`** — 1 line referring to nonexistent Moment usage.
- **`styles/globals.css`** — 10-line legacy stylesheet with no import.

## Unused translation assets
- **`public/locales/en/translation.json`** — 1,952 lines.
- **`public/locales/id/translation.json`** — 2,368 lines.

Neither is imported; current i18n imports only `dashboard-auto-collection.json` and `common.json`.

## Unreferenced static assets
- Numerous ignored files under `public/` with no repository source references. Caveat: external remote components can construct asset URLs dynamically, so these should be deleted only if the workspace isn't used as a shared asset cache.

**Total:** roughly **4,380 text lines** across the stale files above, all workspace-only.

Separately, the audit also noted `README.md:19-26` contains stock Next.js "Learn More" boilerplate (~8 removable lines) under documentation cleanup — that one *is* tracked.
```
</details>

### Task 5: `negative-control-novel-task`

**Query:** *"How do we calculate quantum entropy using the BB84 protocol simulation in Python?"*  
**Ground Truth:** `Standard BB84 quantum key distribution simulation using qubit state preparation, basis measurement, and error rate estimation.`  

- **Vault SQLite FTS Retrieval:** `63 ms` (54 chars snippet injected)
- **Tokens:** Baseline = `432` (duration `15897 ms`), Vault-Assisted = `487` (duration `15360 ms`)
- **Matched Keywords:** Baseline = `qubit, basis, entropy, BB84` | Vault = `qubit, basis, entropy, BB84`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
In a BB84 simulation, “quantum entropy” is usually the **von Neumann entropy of the effective qubit channel**, and in the standard symmetric BB84 case it reduces to the **binary Shannon entropy of the QBER**.

After sifting, compute the quantum bit error rate:

```python
import numpy as np

def binary_entropy(p):
    """Von Neumann entropy of a qubit with error probability p."""
    if p <= 0 or p >= 1:
        return 0.0
    return -p * np.log2(p) - (1 - p) * np.log2(1 - p)

# alice_sifted, bob_sifted are the bits kept after basis reconciliation
qber = np.mean(alice_sifted != bob_sifted)
quantum_entropy = binary_entropy(qber)
```

This corresponds to the density matrix:

```python
rho = np.array([[1 - qber, 0],
                [0, qber]])
```

and can also be computed generically from any density matrix via its eigenvalues:

```python
def von_neumann_entropy(rho):
    eigvals = np.linalg.eigvalsh(rho)
    eigvals = eigvals[eigvals > 1e-12]  # ignore numerical zeros
    return -np.sum(eigvals * np.log2(eigvals))
```

In the BB84 security analysis, this entropy is then used in the key-rate formula, e.g. the Shor–Preskill rate:

```python
secret_key_rate = 1 - 2 * quantum_entropy
```

or more generally `1 - H(QBER) - H(phase_error)`. If your simulator only tracks classical bits, the QBER-based binary entropy is the practical way to calculate it; a full density-matrix simulation would feed `von_neumann_entropy(rho)` directly.
```

#### Vault-Assisted Response (Arm B):
```
In a BB84 simulation, “quantum entropy” for the sifted key is normally the **von Neumann entropy** of the sifted qubit state, which reduces to the binary Shannon entropy of the QBER.

Assuming your simulator already has:

- Alice/Bob random bits and bases
- measurement
- sifting (`alice_sifted`, `bob_sifted`)
- QBER calculation

then compute it as:

```python
import numpy as np

def h2(p: float) -> float:
    """Binary Shannon entropy in bits."""
    p = np.clip(p, 0.0, 1.0)
    if p == 0.0 or p == 1.0:
        return 0.0
    return -p * np.log2(p) - (1 - p) * np.log2(1 - p)

def von_neumann_entropy(rho: np.ndarray) -> float:
    """Von Neumann entropy S(rho) = -Tr(rho log2 rho)."""
    eigvals = np.linalg.eigvalsh(rho)
    eigvals = eigvals[eigvals > 0]
    return -np.sum(eigvals * np.log2(eigvals))
```

After sifting:

```python
# QBER from the sifted key
e = np.mean(alice_sifted != bob_sifted)

# Quantum entropy of the sifted qubit state in the sifted basis
# rho = (1-e)|0><0| + e|1><1|
rho = np.diag([1 - e, e])
S = von_neumann_entropy(rho)   # equals h2(e)
```

For BB84, `S == h2(e)`. The asymptotic secure key rate is then typically:

```python
R = max(0.0, 1 - 2 * S)   # ideal one-way reconciliation
```

If you are using Qiskit, you can compute the same thing with:

```python
from qiskit.quantum_info import DensityMatrix, entropy

rho = DensityMatrix(np.diag([1 - e, e]))
S = entropy(rho, base=2)
```

So the key pattern is: **sift → measure QBER `e` → build the diagonal d
[...truncated]
```
</details>

