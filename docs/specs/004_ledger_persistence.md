# 📒 Specification: Ledger Persistence Schema

**Status:** Draft | **Version:** 1.0

This specification defines the data models for the WORM (Write-Once-Read-Many) Ledger.

## 1. Storage Strategy
* **Local/Demo:** `SQLite` (embedded).
* **Production:** `AWS QLDB` (Quantum Ledger Database) or `Postgres` (with immutable append-only logs).

## 2. Data Models (Schema)

### 2.1 ArtifactRecord
Tracks the existence of a build artifact.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | UUID | Primary Key. |
| `sha` | String | The container image SHA-256. |
| `repo` | String | The source repository URL. |
| `branch` | String | The git branch (e.g., `main`). |
| `created_at` | Timestamp | When the build finished. |

### 2.2 SealRecord (The Truth Stamp)
Links an Artifact to a Policy Validation.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | UUID | The Seal ID (passed to K8s). |
| `artifact_id` | UUID | Foreign Key to Artifact. |
| `policy_hash` | String | Hash of the Cuelang Policy used. |
| `evidence_hash` | String | Merkle Root of the Evidence Tree. |
| `verdict` | String | `PASS` | `FAIL` | `SHADOW_PASS`. |
| `signer` | String | The Identity of the Engine (e.g., `engine-prod-01`). |
| `signature` | Blob | Cryptographic signature of the row. |

### 2.3 EvidenceRecord
Stores the detailed findings (Audit Log).

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | UUID | Primary Key. |
| `seal_id` | UUID | Foreign Key to Seal. |
| `type` | String | `LLM_ANALYSIS` | `TEST_RESULTS` | `SAST_SCAN`. |
| `payload` | JSON | The raw structured output from the tool. |

## 3. Query Patterns
* **Verification:** `SELECT * FROM SealRecord WHERE id = ?`
* **Audit:** `SELECT * FROM EvidenceRecord WHERE seal_id = ?`
* **Traceability:** `SELECT * FROM SealRecord WHERE artifact_id = (SELECT id FROM ArtifactRecord WHERE sha = ?)`
