# 🧪 Collaboration Session 1: Architectural Refinement

This document outlines the outstanding architectural decisions for the Signet Platform. Based on your preferences for **gRPC**, **Google Gen AI**, and **Local-First** development, here are my evaluations and recommendations.

## 1. Persistence Pattern: Event Sourcing + CQRS vs. Standard CRUD
How should we store the Signet Seals and Audit Evidence?

| Pattern | Trade-offs | Recommendation |
| :--- | :--- | :--- |
| **Standard CRUD** | Simple implementation; fast development; easy queries. However, audit logs are "side-effects" and can drift from actual state. | Not recommended for the core Ledger. |
| **Event Sourcing + CQRS** | **Pros:** The "Truth" is the event log; 100% auditability; impossible to delete history. **Cons:** Higher complexity; requires projection logic for the UI. | **Recommended Solution.** |

**Strategic Choice:** **Event Sourcing (ES).**
Signet is a governance tool. In ES, the "Seal" isn't a row in a DB; it's the result of a `SealGenerated` event. This aligns perfectly with the WORM (Write-Once-Read-Many) requirement. We will use **CQRS** to project these events into a local SQLite "Read Model" for the Console to display.

---

## 2. Agentic Framework: Google Gen AI (Vertex/Gemini)
How should the Overseer Agent be structured?

*   **Evaluation:** Using the **Google Gen AI SDK** provides the best integration with Gemini 2.0's 2M context window (essential for analyzing large codebases/diffs).
*   **Recommendation:** Implement the Overseer using **Structured Outputs** (JSON Schema) to ensure the agent returns valid Cuelang/JSON evidence payloads every time.
*   **Local Strategy:** Use the **Vertex AI Local Runtime** or a mock adapter that simulates Gemini responses using local Cuelang fixtures.

---

## 3. Communication: gRPC + Protobuf
*Decision: Finalized by User.*
*   **Implementation:** All internal service communication (CLI -> Engine, Engine -> Overseer) will use gRPC.
*   **Benefit:** Strict contract sharing between the Go backend components and the Console.

---

## 4. Local Mocking & Future Cloud Alignment
How do we build now while preparing for AWS/GCP?

| Component | Local Solution (Now) | Future Cloud Path |
| :--- | :--- | :--- |
| **Ledger Storage** | **EventStoreDB (Docker)** or **File-based Append Log** | **AWS QLDB** or **GCP Ledger** |
| **Read Models** | **SQLite** | **Postgres (Cloud SQL/RDS)** |
| **Object Storage** | **Local Filesystem** | **GCS / S3** |
| **Secrets** | **.env / Local CUE** | **GCP Secret Manager / AWS Secrets Manager** |

**Recommendation:** Use **SQLite** as the primary Read Model even in "Connected" mode for the local CLI, only pushing to the Cloud DB during the final "Seal Write" phase.

---

## 5. UI/UX Strategy: TailwindPlus "Gold Standard" Design
*   **Mandate:** **EVERY** UI component for the Signet Console and CLI outputs must be sourced or inspired by the **TailwindPlus MCP Server**.
*   **Goal:** Leverage the "Gold Standard" design package to ensure the Signet Console has a premium, high-fidelity look with a significant "cool factor."
*   **Workflow:** Before building any UI, we will query `tailwindplus` for "Application UI" and "Marketing" components to maintain a consistent, professional aesthetic that justifies the platform's governance authority.

## 6. UI Visualization: React Flow + TailwindPlus
*   **Recommendation:** Use **React Flow** in the Next.js Console to visualize the Merkle Tree of evidence.
*   **Design Integration:** We will use TailwindPlus components (Overlays, Slide-overs, and Cards) to display the detailed metadata for each node in the tree, ensuring the transition from the Merkle graph to the data view is seamless and visually stunning.

---

### 📝 Your Selections
Please indicate your preference for the **Persistence Pattern**:
1.  **[x] Option A:** Event Sourcing + CQRS (High-fidelity audit, higher complexity)
2.  **[ ] Option B:** Standard CRUD with Audit Tables (Simpler, lower audit rigor)

*I recommend Option A for a "Gold Standard" governance tool.*
