# You are a professional software engineer answering technical and ar...

- **ID**: `antigravity-250d4c601b7c74efabe1a253dd72d53ab6ef3e622700c45c345e47c832b70156`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-03 10:34:42`
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

In the context of a BB84 protocol simulation, **quantum entropy** is typically evaluated in one of two ways depending on whether you are tracking the state's density matrix (**Von Neumann Entropy**) or the information leakage based on the Quantum Bit Error Rate (**Binary Shannon Entropy**).

---

### 1. Von Neumann Entropy (State Mixedness / Eve's Impact)

When simulating noisy channels or an intercept-resend attack (Eve), pure quantum states $|\psi\rangle$ degrade into mixed states described by a density matrix $\rho$. 

The Von Neumann entropy $S(\rho)$ is calculated from the eigenvalues $\lambda_i$ of $\rho$:

$$S(\rho) = -\mathrm{Tr}(\rho \log_2 \rho) = -\sum_i \lambda_i \log_2 \lambda_i$$

#### Python Implementation (`numpy` / `scipy`):
```python
import numpy as np

def von_neumann_entropy(rho: np.ndarray) -> float:
    """Calculates Von Neumann entropy S(rho) = -Tr(rho log2 rho)"""
    # Obtain eigenvalues of Hermitian density matrix
    eigenvalues = np.linalg.eigvalsh(rho)
    
    # Filter non-positive eigenvalues to handle log2(0) numerically
    positive_eigenvals = eigenvalues[eigenvalues > 1e-12]
    
    return float(-np.sum(positive_eigenvals * np.log2(positive_eigenvals)))

# Example: Pure state |0><0| has entropy 0; maximally mixed state I/2 has entropy 1
rho_pure = np.array([[1, 0], [0, 0]])
rho_mixed = 0.5 * np.eye(2)

print("Pure State Entropy:", von_neumann_entropy(rho_pure))   # 0.0
print("Mixed State Entropy:", von_neumann_entropy(rho_mixed)) # 1.0
```

---

### 2. Binary Shannon Entropy of QBER (Information Leakage & Key Rate)

In privacy amplification for BB84, entropy is calculated from the **Quantum Bit Error Rate (QBER, $e$)** to bound the mutual information between Alice and Eve.

The binary entropy function $h_2(e)$ is:

$$h_2(e) = -e \log_2(e) - (1 - e) \log_2(1 - e)$$

The asymptotic secret key generation rate $R$ under 1-way classical post-processing is:
$$R \approx 1 - 2h_2(e)$$

#### Python Implementation:
```python
import math

def binary_entropy(e: float) -> float:
    """Calculates Shannon binary entropy h_2(e)."""
    if e <= 0.0 or e >= 1.0:
        return 0.0
    return -e * math.log2(e) - (1.0 - e) * math.log2(1.0 - e)

def calculate_qber_and_key_rate(alice_bits: list[int], bob_bits: list[int]):
    """Calculates QBER from sifted test bits and determines secret key capacity."""
    mismatches = sum(a != b for a, b in zip(alice_bits, bob_bits))
    qber = mismatches / len(alice_bits) if alice_bits else 0.0
    
    h_e = binary_entropy(qber)
    # Shor-Preskill secret key rate lower bound:
    secret_key_rate = max(0.0, 1.0 - 2 * h_e)
    
    return {
        "qber": qber,
        "entropy_leakage": h_e,
        "secret_key_rate": secret_key_rate,
        "secure": qber < 0.11  # BB84 security threshold is ~11%
    }
```

---

