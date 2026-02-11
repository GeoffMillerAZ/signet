# 🛡️ Specification: Engine Integrity (The "Step 0" Check)

**Status:** Draft | **Version:** 1.0

This specification defines the self-validation logic that the Signet Engine must perform *before* executing any pipeline steps. This ensures that the CI workflow file itself has not been tampered with to bypass governance.

## 1. The Threat Model
A developer might attempt to bypass Signet by:
1.  Modifying `.github/workflows/signet.yml` to remove the `seal` step.
2.  Adding `if: always()` or `continue-on-error: true` to critical gates.
3.  Injecting a sidecar container that exfiltrates secrets.

## 2. The Verification Logic (Step 0)
The `signet verify-integrity` command runs as the very first instruction in the CI runner.

### 2.1 Inputs
* `CurrentWorkflowPath`: Path to the executing YAML (e.g., `.github/workflows/deploy.yml`).
* `MasterHashUrl`: URL to the authorized hash in the Signet Registry (or `internal/embed`).

### 2.2 Algorithm
1.  **Read:** Load the content of `CurrentWorkflowPath`.
2.  **Canonicalize:** Parse the YAML into a standard structure (sort keys, remove comments/whitespace) to ensure deterministic hashing.
3.  **Hash:** Calculate `SHA-256(CanonicalYAML)`.
4.  **Fetch:** Retrieve the `AuthorizedHash` for this repo/team from the Registry.
5.  **Compare:**
    * If `CalculatedHash == AuthorizedHash`: **PROCEED**.
    * If `CalculatedHash != AuthorizedHash`: **ABORT**.

### 2.3 Failure Mode
If validation fails, the Engine must:
1.  Log a critical security alert to the Console.
2.  Exit with code `1` (Force Fail).
3.  **Do not** output specific details about *why* it failed to the CI logs (prevent information leakage), just "Integrity Check Failed."

## 3. The "Bootstrap Trap"
To prevent developers from simply deleting the `verify-integrity` step:
* **Repo Policy:** GitHub Branch Protection rules require the "Signet / Integrity" job to pass.
* **LLM Overseer:** The Agent (Spec 002) independently verifies that the standard workflow file exists and is the *only* workflow file triggered by push events.
