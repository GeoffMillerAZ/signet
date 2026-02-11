# Signet: Autonomous Governance Platform

**Signet** is a next-generation governance engine that bridges the gap between Developer Velocity and Operational Safety. It combines **Cuelang's deterministic "Law"** with **LLM-powered "Spirit" analysis** to create a cryptographically verifiable audit trail for every deployment.

## 🚀 Key Features

*   **🛡️ The Signet Seal:** A Merkle-hashed cryptographic proof linking Code, Configuration, and Policy.
*   **👁️ The Overseer:** An AI Agent (Gemini 2.0) that detects "Silent Killers" (semantic reliability risks) in your code diffs.
*   **📒 Immutable Ledger:** An event-sourced, append-only log ensuring 100% auditability of every artifact.
*   **🎨 Gold Standard Console:** A high-fidelity Next.js dashboard for visualizing Merkle proofs and policy compliance.

## 🛠️ Installation

```bash
# Clone the repository
git clone https://github.com/GeoffMillerAZ/signet.git
cd signet

# Build the CLI
task build

# Install the binary (optional)
mv bin/signet /usr/local/bin/
```

## 🚦 Quick Start

### 1. Initialize a Repo
Bootstrap a repository with the default Signet structure and Cuelang policies.
```bash
signet init
```

### 2. Run a Local Check
Perform a local governance scan, including Policy Unification and Overseer Dry-Run.
```bash
signet check
```

### 3. Verify Integrity
Ensure your CI workflow hasn't been tampered with.
```bash
signet verify-integrity --workflow .github/workflows/deploy.yml
```

## 🏗️ Architecture

Signet follows the **Hexagonal Architecture** pattern:

*   **Core:** Pure domain logic (Sealing, Promotion, Validation).
*   **Ports:** Interfaces for external interaction (Ledger, LLM, Git).
*   **Adapters:** Implementations (SQLite, Gemini, Go-Github).

The platform uses **Event Sourcing + CQRS** for its persistence layer, ensuring that the history of every artifact is preserved forever.

## 💻 Development

### Prerequisites
*   Go 1.23+
*   Node.js 20+
*   Task

### Commands
*   `task test`: Run the full Go test suite.
*   `task proto`: Regenerate gRPC code from Protobuf definitions.
*   `task dev`: Start the local development environment.

## 📜 License
MIT
