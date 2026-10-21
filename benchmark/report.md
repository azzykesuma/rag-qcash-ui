# 🧠 LLM Context Vault: Benchmark Report

**Date:** 2026-10-15 15:17:33  
**Runner Backend:** `OpenCode (deepseek/deepseek-flash)`  
**Tasks Evaluated:** 5 (Successful: 3)  

---

## 📊 Executive Scorecard

| Metric | Baseline (Arm A) | Vault-Assisted (Arm B) | Improvement / Impact |
| :--- | :--- | :--- | :--- |
| **Average Tokens / Task** | `361` tokens | `1672` tokens | **+362.7%** tokens |
| **Average Wall-Clock Time** | `20105 ms` | `16336 ms` | **1.23x speedup** |
| **Fact / Keyword Recall** | `31.7%` | `61.7%` | **+94.7% accuracy** |

> [!NOTE]
> **Exploration Savings Context:** In interactive agent workflows, a baseline agent typically requires 5–15 roundtrip tool calls (`grep`, `view_file`, `read`) consuming **15,000–50,000+ tokens** to locate institutional knowledge. Injecting a bounded vault excerpt upfront eliminates these exploratory turns.

---

## 📋 Per-Task Comparison Table

| Task ID | Project | FTS Latency | Baseline Tokens | Vault Tokens | Baseline Acc | Vault Acc | Δ Score |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `auth-gate-architecture` ✅ | `qcash-ui` | `1455 ms` | `113` | `2251` | `0%` | `50%` | `+50%` |
| `vault-scan-optimization` ✅ | `llm-context-vault` | `98 ms` | `467` | `2094` | `20%` | `60%` | `+40%` |
| `sanitization-path-policy` ⚠️ Error | `llm-context-vault` | `1595 ms` | `0` | `2088` | `0%` | `75%` | `+75%` |
| `overengineering-dead-assets` ⚠️ Error | `qcash-ui-dashboard-auto-collection` | `885 ms` | `0` | `2028` | `0%` | `100%` | `+100%` |
| `negative-control-novel-task` ✅ | `general` | `270 ms` | `504` | `671` | `75%` | `75%` | `+0%` |

---

## 🔍 Detailed Task Breakdown

### Task 1: `auth-gate-architecture`

**Query:** *"How is page readiness and authentication gating handled across pages in QCash micro-frontends?"*  
**Ground Truth:** `HostAuthGate owns /auth/me and /menu/me. Pages should control rendering readiness through isAuthoritiesReady rather than stale local cache or local login strings.`  

- **Vault SQLite FTS Retrieval:** `1455 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `113` (duration `944 ms`), Vault-Assisted = `2251` (duration `5202 ms`)
- **Matched Keywords:** Baseline = `` | Vault = `HostAuthGate, isAuthoritiesReady`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
<｜｜DSML｜｜ calls>
<｜｜DSML｜｜ invoke name="bash">
<｜｜DSML｜｜ parameter name="command" string="true">ls</｜｜DSML｜｜ parameter>
</｜｜DSML｜｜ invoke>
</｜｜DSML｜｜ calls>
```

#### Vault-Assisted Response (Arm B):
```
# Page Readiness & Auth Gating in QCash Micro-Frontends

Gating is centralized in a single host component rather than delegated to each federated module.

## Core component
- **`components/providers/HostAuthGate.tsx`** — the single gate that owns all auth validation for federated children.
- **`pages/_app.tsx`** — places `HostAuthGate` *inside* both the Ordent and BRI auth providers, and *above* every federated child route. This ordering is what lets one gate wrap all micro-frontends regardless of which provider a page belongs to.

## What the gate does
1. **Readiness gating** — blocks protected children until the required auth providers (`isAuthoritiesReady`) have resolved.
2. **Public-route bypass** — routes marked public skip validation entirely.
3. **Central guard execution** — runs the Ordent `guard()` for Ordent routes and the BRI `guard()` for BRI-owned routes, instead of letting each federated module call its own copy.
4. **Per-token validation tracking** — validation is keyed by token so it runs once per token, including handling token replacement while a validation is in flight.
5. **Failure handling** — retryable auth errors render a retry UI; 401s clear session/provider tokens and redirect.
6. **Cross-tab/storage sync** — subscribes to the custom `"qc-bridge-sync"` event (same-tab) and native `storage` events (cross-tab). `AuthBridgeSync.ts` remains the sync owner and is intentionally untouched.

## Why centralization matters here
The original production bug was t
[...truncated]
```
</details>

### Task 2: `vault-scan-optimization`

**Query:** *"How does llm-context-vault optimize repeat scans and avoid re-processing unchanged sessions?"*  
**Ground Truth:** `Uses an incremental cache to skip unchanged file-based sessions, limits changed sessions per batch (default 10), and defers the repository-wide privacy audit for bounded scans.`  

- **Vault SQLite FTS Retrieval:** `98 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `467` (duration `39717 ms`), Vault-Assisted = `2094` (duration `14551 ms`)
- **Matched Keywords:** Baseline = `unchanged` | Vault = `unchanged, audit, bounded`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
llm-context-vault treats a repeat scan as an incremental sync rather than a full rebuild. The scanner keeps a persistent scan cache/state keyed by source path or session identity, storing metadata such as:

- source `mtimeMs`
- source `size`
- a content fingerprint/hash of the session file or parsed session
- for append-only logs, a checkpoint such as last processed byte offset, message UUID, or message count

On each scan it:

1. **Stats first:** if `mtimeMs` and `size` match the cached entry, the session is skipped entirely — no file read, parsing, rendering, chunking, or embedding.
2. **Hashes only when metadata changes:** if `mtime`/`size` differ, it computes the content hash. If the hash is unchanged (e.g. a `touch`), it only refreshes the cached metadata and still skips re-processing.
3. **Parses only changed sessions:** if the hash differs, only that session is re-parsed and upserted. For append-only LLM transcripts, it can parse from the stored offset/UUID checkpoint instead of rereading the whole file; truncation/rewrite triggers a full reparse of just that session.
4. **Scopes derived updates:** messages, chunks, embeddings, and vault output are updated only for the changed session; unchanged sessions reuse existing data.
5. **Prunes deleted sources:** files no longer present are removed from the cache/index rather than forcing a full rescan.
6. **Allows bypass:** a `--force`/`--full`-style option bypasses the cache when a complete re-scan is needed.

This makes rep
[...truncated]
```

#### Vault-Assisted Response (Arm B):
```
Based on the referenced session, **llm-context-vault did not yet optimize repeat scans**; it reprocessed unchanged sessions. The assistant identified the bottleneck and proposed an incremental-scan design.  

In the captured code:

- `pkg/vault/vault.go:358` extracts, sanitizes, audits, and exports **every session on every scan**, even unchanged ones.
- `pkg/exporter/jsonl.go:55` reads, parses, and rewrites the **entire `dataset.jsonl` for each imported conversation**.

The proposed optimization — **incremental scanning** — would avoid re-processing by:

- Keeping a local, Git-ignored **scan-state cache** containing source fingerprints, processing versions, and corresponding export IDs.
- For file-based sessions: comparing **size + modification time**, falling back to **content hashes** when needed.
- For OpenCode: tracking **verified database revision fields**, including message/part changes.
- Reprocessing only when:
  - the source fingerprint changes,
  - sanitizer settings, custom redaction words, or extractor versions change,
  - a previous run failed or the source was actively changing.
- Recording completion **only after successful persistence**, retrying failures, recovering missing outputs, and offering a `--full` flag to force reprocessing.

Additional proposed scan optimizations:

- **Batch dataset updates**: collect imports into one scan-wide batch, merge with `dataset.jsonl` once, write via temp file + safe replacement, and skip rewriting identical Markdown/Share
[...truncated]
```
</details>

### Task 3: `sanitization-path-policy`

**Query:** *"What is the vault sanitization policy for Windows and Unix machine file paths before publishing?"*  
**Ground Truth:** `Strips Windows paths (C:\Users\username\...) and Unix paths (/home/username/...) replacing them with ~/...`  

- **Vault SQLite FTS Retrieval:** `1595 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `0` (duration `60000 ms`), Vault-Assisted = `2088` (duration `9185 ms`)
- **Matched Keywords:** Baseline = `` | Vault = `C:\Users, redact, path`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```

```

#### Vault-Assisted Response (Arm B):
```
Based on the reference material available, I can't cite the exact regex/rule set for path redaction — the excerpts describe the *policy intent* and test coverage but not the sanitizer implementation itself. What the vault design establishes:

- **Paths are treated as sensitive, first-class redaction targets.** The stated sanitization policy is to "treat all persisted fields as sensitive" and sanitize both `Conversation.Metadata` and `Message.Metadata`, with regression fixtures explicitly covering **credentials, Windows/Unix paths, emails, and custom terms**. So Windows machine paths (e.g. `C:\Users\...`) and Unix paths (e.g. `/home/...`, `/Users/...`) are redacted before anything is published, alongside secrets.

- **Redaction is applied at write/scan time, not at read time.** Sanitization happens in the scan/import pipeline (`pkg/vault/naming.go` derives source IDs *after* sanitization; `pkg/vault/vault.go` canonicalizes records), and ownership/retirement of stale unredacted exports is handled in `pkg/vault/scan.go`. The audit command (`vault audit`) is the pre-publish gate that verifies no unsanitized content remains.

- **Custom terms extend the base rules.** `--redact-words` lets users add terms on top of the built-in credential/email/path patterns.

Caveats worth verifying against the current tree, since these excerpts are historical and known to be incomplete:
1. The policy depends on sanitization running *before* filenames/IDs are derived — a prior review flagged that 
[...truncated]
```
</details>

### Task 4: `overengineering-dead-assets`

**Query:** *"What dead scripts and unnecessary assets were identified for removal during the overengineering audit?"*  
**Ground Truth:** `scripts/analyze-bundle.js, scripts/update-version-file.js, unused asset directories (files, flags, fonts, icons, images, img), and version file.`  

- **Vault SQLite FTS Retrieval:** `885 ms` (5998 chars snippet injected)
- **Tokens:** Baseline = `0` (duration `60000 ms`), Vault-Assisted = `2028` (duration `4164 ms`)
- **Matched Keywords:** Baseline = `` | Vault = `analyze-bundle.js, update-version-file.js, assets, audit`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```

```

#### Vault-Assisted Response (Arm B):
```
During the overengineering audit, the following **dead scripts and unnecessary assets** were flagged as workspace debris (already Git-ignored, so they don't affect repository size but are removable locally):

**Dead/stale scripts & tooling files**
- `scripts/analyze-bundle.js` — 37 lines; no package script, no analyzer dependency, and no matching Next analyzer configuration.
- `scripts/update-version-file.js` — 9 lines; no caller, and its write operation is commented out.
- `.eslintrc.json` — 3 lines; the project uses Biome, and ESLint isn't installed.

**Unused text/config leftovers**
- `todo.md` — 1 line referencing a nonexistent Moment usage.
- `styles/globals.css` — 10-line legacy stylesheet; never imported.
- `README.md:19-26` — stock Next.js "Learn More" boilerplate (~8 lines).

**Large unused translation leftovers**
- `public/locales/en/translation.json` — 1,952 lines
- `public/locales/id/translation.json` — 2,368 lines

Neither is imported; current i18n uses only `dashboard-auto-collection.json` and `common.json`. Together with the stale files above, this is roughly **4,380 text lines** of workspace-only cleanup.

**Ignored binary/static asset tree**
- Numerous ignored files under `public/` with no repository source references. **Caveat:** because external remote (Module Federation) components can construct asset URLs dynamically, these should only be deleted if the workspace isn't serving as a shared asset cache.

Note: the audit distinguished these from *actual over
[...truncated]
```
</details>

### Task 5: `negative-control-novel-task`

**Query:** *"How do we calculate quantum entropy using the BB84 protocol simulation in Python?"*  
**Ground Truth:** `Standard BB84 quantum key distribution simulation using qubit state preparation, basis measurement, and error rate estimation.`  

- **Vault SQLite FTS Retrieval:** `270 ms` (54 chars snippet injected)
- **Tokens:** Baseline = `504` (duration `19654 ms`), Vault-Assisted = `671` (duration `29255 ms`)
- **Matched Keywords:** Baseline = `basis, entropy, BB84` | Vault = `qubit, entropy, BB84`

<details><summary><b>View Model Responses (Arm A vs Arm B)</b></summary>

#### Baseline Response (Arm A):
```
In a BB84 simulation, “quantum entropy” usually means one of two things:

1. **Shannon entropy of the sifted-key error rate** — used for privacy amplification.
2. **von Neumann entropy of the quantum state** — `S(ρ) = -Tr(ρ log₂ ρ)`.

For the BB84 key-rate calculation, use the QBER after basis sifting:

```python
import numpy as np

def shannon_entropy(p):
    p = np.asarray(p, dtype=float)
    p = p[p > 0]
    return -np.sum(p * np.log2(p))

# after sifting: keep indices where alice_bases == bob_bases
alice_sifted = alice_bits[matching]
bob_sifted   = bob_bits[matching]

qber = np.mean(alice_sifted != bob_sifted)
H_qber = shannon_entropy([qber, 1 - qber])

# typical BB84 secret-key rate estimate
key_rate = 1 - 2 * H_qber
```

For the actual quantum/von Neumann entropy, reconstruct the density matrix and diagonalize it:

```python
def von_neumann_entropy(rho):
    eigvals = np.linalg.eigvalsh(rho)
    eigvals = eigvals[eigvals > 1e-12]
    return -np.sum(eigvals * np.log2(eigvals))
```

If you only have the sifted classical bits, the density matrix is diagonal:

```python
p0 = np.mean(np.array(alice_sifted) == 0)
p1 = 1 - p0
rho = np.diag([p0, p1])
S = von_neumann_entropy(rho)   # equals shannon_entropy([p0, p1])
```

If you want the full BB84 state entropy, build `ρ` from the four BB84 states `|0⟩, |1⟩, |+⟩, |−⟩` with their probabilities, then call `von_neumann_entropy(rho)`. With uniformly random bits and bases, this gives the maximally mixed state `I/2`, so `S = 1` bit.  

[...truncated]
```

#### Vault-Assisted Response (Arm B):
```
In a BB84 simulation, “quantum entropy” usually means either the **von Neumann entropy** of the qubit/channel state or the **Shannon entropy** of the sifted key. They are different, so compute the one your security analysis needs.

Assuming your simulation produces `alice_key`, `bob_key`, and optionally density matrices:

```python
import numpy as np

def h2(p):
    """Binary Shannon/von Neumann entropy for a bit-flip channel."""
    if p <= 0.0 or p >= 1.0:
        return 0.0
    return -p * np.log2(p) - (1 - p) * np.log2(1 - p)

def shannon_entropy(bits):
    """Classical entropy of the sifted key bits."""
    bits = np.asarray(bits)
    p1 = np.mean(bits)
    p0 = 1.0 - p1
    probs = np.array([p0, p1])
    probs = probs[probs > 0]
    return -np.sum(probs * np.log2(probs))

def von_neumann_entropy(rho):
    """Quantum entropy S(rho) = -Tr(rho log2 rho)."""
    eigvals = np.linalg.eigvalsh(rho)
    eigvals = eigvals[eigvals > 1e-12]
    return -np.sum(eigvals * np.log2(eigvals))

# After BB84 sifting and parameter estimation:
qber = sum(a != b for a, b in zip(alice_key, bob_key)) / len(alice_key)

# 1. Quantum entropy of a bit-flip/depolarizing channel with QBER e:
S_quantum = h2(qber)

# 2. Classical entropy of the sifted key:
H_key = shannon_entropy(bob_key)
```

For BB84, a common security-rate calculation uses the binary entropy of the QBER:

```python
secure_key_rate = 1 - 2 * h2(qber)
# or, with explicit error correction leakage:
secure_key_rate = 1 - h2(qber) - leak
[...truncated]
```
</details>

