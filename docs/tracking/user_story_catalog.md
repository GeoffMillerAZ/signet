# 📋 User Story Catalog: Signet Platform

**Status:** Draft | **Version:** 1.0
**Project:** Signet

---

## 🏗️ Phase 1: Foundation (Integrity & CLI)

| ID | User Story | Acceptance Criteria | Status |
| :--- | :--- | :--- | :--- |
| **[US-101]** | **Secure Bootstrap**<br>As a Platform Engineer, I want the Signet Engine to verify its own workflow file hash before running, so that developers cannot bypass governance steps. | 1. Engine canonicalizes and hashes its YAML.<br>2. Compares against authorized registry hash.<br>3. Hard fails if mismatch detected. | 📝 Draft |
| **[US-102]** | **Local Integrity Check**<br>As a Developer, I want to run `signet check` locally before pushing, so that I can catch architectural violations early. | 1. CLI runs local CUE unification.<br>2. CLI runs dry-run of Overseer.<br>3. Actionable errors returned to stdout. | 📝 Draft |

## 👁️ Phase 2: The Overseer (AI & Merkle)

| ID | User Story | Acceptance Criteria | Status |
| :--- | :--- | :--- | :--- |
| **[US-201]** | **Semantic Reliability Scan**<br>As a SRE, I want the Overseer Agent to detect "Silent Killers" (e.g., missing SIGTERM handlers), so that we reduce production outages. | 1. LLM analyzes code diffs using SRE persona.<br>2. Outputs structured JSON Evidence Payload.<br>3. Findings map to specific line numbers. | 📝 Draft |
| **[US-202]** | **Smart Validation Skip**<br>As a Developer, I want Signet to skip AI analysis for unchanged architectural groups, so that my CI pipelines remain fast. | 1. Merkle Group hashes compared to previous Seal.<br>2. Analysis inherited if hashes match.<br>3. New Seal includes inherited evidence. | 📝 Draft |

## 📒 Phase 3: The Ledger (Persistence)

| ID | User Story | Acceptance Criteria | Status |
| :--- | :--- | :--- | :--- |
| **[US-301]** | **Immutable Truth Stamp**<br>As an Auditor, I want every valid deployment to have a cryptographically signed Seal in an immutable ledger, so that I have a tamper-proof audit trail. | 1. Seal record includes Merkle root of evidence.<br>2. Record signed by Signet Engine.<br>3. Entry is append-only and immutable. | 📝 Draft |
| **[US-302]** | **Local-First Demo Mode**<br>As a Stakeholder, I want to run the entire Signet flow on my laptop using SQLite, so that I can evaluate the platform without cloud dependencies. | 1. System uses SQLiteAdapter for Ledger.<br>2. Pre-seeded with architectural fixtures.<br>3. No loss of logic fidelity in local mode. | 📝 Draft |

## 🎨 Phase 4: The Console (UI/UX)

| ID | User Story | Acceptance Criteria | Status |
| :--- | :--- | :--- | :--- |
| **[US-401]** | **Merkle Proof Visualization**<br>As an SRE, I want to visually navigate the Seal's Merkle Tree in the Console, so that I can quickly understand which evidence led to a promotion. | 1. React Flow used for tree visualization.<br>2. Nodes expandable to show LLM findings.<br>3. TailwindPlus styling applied for "Gold Standard" look. | 📝 Draft |
| **[US-402]** | **Cross-Repo Promotion**<br>As a Developer, I want to promote a sealed artifact to production via a single CLI command, so that I don't have to manually update deployment manifests. | 1. CLI verifies Seal in Ledger before promotion.<br>2. CLI updates deployment repo and opens PR.<br>3. PR includes the Ledger Transaction ID. | 📝 Draft |