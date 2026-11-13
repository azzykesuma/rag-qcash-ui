# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-78f557df61f7006c2088137dd96d8fdedbe733411e09eea54e10c7936e1d4b48`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 09:28:55`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

You are a professional software engineer answering technical and architectural questions about this codebase. Answer concisely and accurately, citing specific components, files, or patterns when known.

Do not use tools, inspect files, or search the web. Answer using only the question and any supplied vault excerpt.

No relevant previous conversations found in the vault.

User Question:
How do we calculate quantum entropy using the BB84 protocol simulation in Python?

---

## Turn 2: Assistant

In a BB84 protocol simulation, **quantum entropy** is typically evaluated in one of two contexts:

1. **Shannon / Binary Entropy of QBER ($H_2(e)$)**: Measures information leakage to an eavesdropper (Eve) and bounds the distillable secret key rate:
   $$R \ge 1 - 2H_2(e), \quad \text{where } H_2(e) = -e \log_2(e) - (1 - e) \log_2(1 - e)$$
2. **Von Neumann Entropy ($S(\rho)$)**: Measures the entanglement or state mixture of the density matrix $\rho$:
   $$S(\rho) = -\mathrm{Tr}(\rho \log_2 \rho) = -\sum_i \lambda_i \log_2 \lambda_i$$

---

### Python Implementation

```python
import numpy as np

# -------------------------------------------------------------
# 1. Binary Entropy of QBER (Information Leakage / Secret Key Rate)
# -------------------------------------------------------------
def binary_entropy(p: float) -> float:
    """Computes Shannon entropy H_2(p) for error probability p."""
    if p <= 0.0 or p >= 1.0:
        return 0.0
    return -p * np.log2(p) - (1 - p) * np.log2(1 - p)

def secret_key_rate(qber: float) -> float:
    """Asymptotic secret key rate under one-way BB84."""
    h = binary_entropy(qber)
    rate = 1.0 - 2.0 * h
    return max(0.0, rate)


# -------------------------------------------------------------
# 2. Von Neumann Entropy of a Density Matrix
# -------------------------------------------------------------
def von_neumann_entropy(rho: np.ndarray) -> float:
    """
    Computes S(rho) = -Tr(rho log2(rho)) from density matrix eigenvalues.
    """
    eigenvalues = np.linalg.eigvalsh(rho)
    # Filter out numerical zeros/negatives
    pos_eigvals = eigenvalues[eigenvalues > 1e-12]
    return float(-np.sum(pos_eigvals * np.log2(pos_eigvals)))


# -------------------------------------------------------------
# 3. Example BB84 State & QBER Simulation
# -------------------------------------------------------------
if __name__ == "__main__":
    # BB84 ensemble density matrix (equal mix of |0>, |1>, |+>, |->):
    # rho = 1/4 * (|0><0| + |1><1| + |+><+| + |-><-|) = I / 2
    rho_ensemble = 0.5 * np.eye(2, dtype=complex)
    s_rho = von_neumann_entropy(rho_ensemble)
    print(f"Von Neumann Entropy S(rho): {s_rho:.4f} bits")  # Expected: 1.0 bit

    # Simulated QBER (e.g., 5% channel error / Eve intercept)
    qber = 0.05
    leakage = binary_entropy(qber)
    asymptotic_rate = secret_key_rate(qber)
    print(f"QBER: {qber * 100:.1f}%")
    print(f"Binary Entropy H_2(e): {leakage:.4f} bits")
    print(f"Distillable Key Fraction: {asymptotic_rate:.4f}")
```

### Key Considerations
- **Threshold Limit**: If the QBER exceeds $\approx 11\%$ ($H_2(0.11) \approx 0.5$), $1 - 2H_2(e) \le 0$, meaning no secure key can be distilled via standard one-way post-processing.
- **Finite-Key Effects**: In finite-length key simulations, replace asymptotic formulas with smooth min-entropy $H_{\min}^\varepsilon(A|E)$ bounds via the Leftover Hash Lemma.

---

