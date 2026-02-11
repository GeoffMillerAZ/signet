# 🔐 The Signet Seal: Merkle-Hashing Protocol

**Status:** Draft | **Version:** 1.0

The **Signet Seal** is the fundamental unit of trust in the platform. It is not just a "Pass" flag; it is a cryptographic proof that links Code, Configuration, and Policy.

## 1. The Seal Structure
A Seal is a JSON object that is hashed and signed.

```json
{
  "version": "v1",
  "artifact_sha": "sha256:a1b2c3d4...",
  "config_hash": "sha256:e5f6g7h8...",
  "policy_hash": "sha256:i9j0k1l2...",
  "evidence_tree_root": "sha256:m3n4o5p6...",
  "timestamp": "2026-02-10T14:00:00Z",
  "issuer": "signet-engine-prod-01"
}
```
2. The Merkle Tree of Evidence
To optimize for speed and granular validation, we use a Merkle Tree structure for evidence.

2.1 Architectural Groups
Files are grouped by architectural concern (defined in configs/schema/groups.cue).

Group A (Persistence): db/**/*.sql, models/*.go

Group B (Networking): api/**/*.go, proto/*.proto

2.2 The Hashing Logic
Leaf Nodes: Calculate SHA-256 of every file in a Group.

Group Nodes: Calculate SHA-256 of all Leaf Hashes combined.

Root Node: Calculate SHA-256 of all Group Hashes.

2.3 Smart Validation (The "Short-Circuit")
When a PR is opened:

Signet calculates the new Group Hashes.

It compares them to the Previous Seal.

If Group A hash is unchanged: The LLM analysis for "Persistence" is Inherited (skipped).

If Group B hash changed: The LLM is triggered only for the Networking files.

3. The Truth Stamp (WORM Storage)
Once a Seal is generated:

It is sent to the Ledger Adapter.

The Ledger writes it to immutable storage (QLDB/Object Lock).

The Ledger returns a transaction_id.

This transaction_id is embedded into the Deployment PR as the "Proof of Seal."

4. Verification Flow
At the Admission Controller (Kubernetes) or Deployment Repo:

Extract signet-seal-id from the deployment annotations.

Query the Ledger: "Does this ID exist and is it valid?"

Critical Check: Verify that the artifact_sha in the Seal matches the container image being deployed.

If match: Allow. If mismatch: Deny.
