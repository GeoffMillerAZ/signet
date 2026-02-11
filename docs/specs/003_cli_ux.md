# 💻 Specification: Signet CLI (Command Reference)

**Status:** Draft | **Version:** 1.0

This specification defines the user experience for the `signet` command-line interface.

## 1. Global Flags
* `--verbose, -v`: Enable debug logging.
* `--config`: Path to custom config file (default: `.signet/config.cue`).
* `--mode`: `local` (uses SQLite/Mock LLM) or `connected` (uses Registry/QLDB).

## 2. Core Commands

### `signet init`
Bootstraps a repository for Signet governance.
* **Action:**
    1.  Creates `.signet/` directory structure.
    2.  Writes default Cuelang schemas (`core/`, `app/`).
    3.  Injects `.github/workflows/signet.yml`.
    4.  Configures `CODEOWNERS`.
* **Flags:** `--template [go|node|python]`

### `signet check`
Runs a local governance scan (Pre-Push hook).
* **Action:**
    1.  Validates Cuelang unification (Global Policy + Local Patch).
    2.  Runs a "Dry Run" of the Overseer (using Local/Mock LLM).
    3.  Reports violations to stdout.
* **Output:**
    ```text
    [PASS] Schema Validation
    [FAIL] Overseer Analysis:
       - main.go:45: Missing SIGTERM handler (HIGH)
    ```

### `signet promote`
Promotes a sealed artifact to a target environment.
* **Syntax:** `signet promote --env <target> --artifact <tag>`
* **Action:**
    1.  Verifies the artifact has a valid **Seal** in the Ledger.
    2.  Clones the **Deployment Repo** (if separate).
    3.  Updates the `values.cue` or `kustomization.yaml` with the new tag and Seal ID.
    4.  Creates a Pull Request with the changes.
* **Flags:**
    * `--auto-merge`: If policy allows, auto-merge the PR.
    * `--break-glass`: Bypass gates (triggers Incident Incident).

## 3. Developer Experience (DX)
* **Error Messages:** Must be actionable. Link to the specific Policy ID in the documentation.
* **Speed:** Local checks must complete in < 2s (excluding LLM latency).
* **Updates:** The CLI should auto-check for updates against the Registry.

