# 🏆 Gold Standard Alignment: Architectural Integrity

**Status:** Draft | **Version:** 1.0

This document explicitly maps the **Signet Platform** architecture to the "Gold Standard" engineering kit defined in `docs/kit`. We strictly adhere to **Hexagonal Architecture**, **Spec-Driven Development**, and **Cuelang-First Configuration**.

## 1. Architecture: Hexagonal Core
*Ref: [ARCHITECTURE.md](../../docs/kit/ARCHITECTURE.md)*

Signet is built as a modular Monolith using strict **Ports & Adapters**.

### The Boundaries
* **Core (`internal/core`)**: Contains the *Pure Domain Logic* for `SignetEngine`, `Overseer`, and `Ledger`.
    * *Rule:* No external imports (no AWS SDK, no Kubernetes client).
* **Ports (`internal/core/ports`)**: Defined interfaces for all interactions.
    * `CryptoPort`: For hashing (SHA-256 / Merkle Trees).
    * `LLMProviderPort`: For the Overseer Agent.
    * `LedgerPort`: For the immutable record.
* **Adapters (`internal/adapters`)**: The implementation details.
    * `QLDBAdapter`: Production ledger.
    * `SQLiteAdapter`: **Local/Demo** ledger (as per `STACK.md`).
    * `CobraAdapter`: The CLI entry point.

## 2. Configuration: Cuelang First
*Ref: [CONFIG.md](../../docs/kit/CONFIG.md)*

We explicitly reject YAML/JSON for internal configuration and policy logic.

* **Schema-Driven:** All internal data structures (Evidence Payloads, Policy Bundles) are defined in `configs/schema/*.cue`.
* **Composition:** We use Cuelang's **Unification** to merge "Global Policy" with "Team Patches," ensuring constraints cannot be overridden.
* **Validation:** Input validation happens at the boundary (CLI/API) using Cuelang constraints, ensuring the Core never receives invalid state.

## 3. Development Methodology
*Ref: [METHODOLOGY.md](../../docs/kit/METHODOLOGY.md)*

* **Spec-Driven (SDD):** No code is written without a `docs/specs/*.md` file defining the prompt and acceptance criteria.
* **Traceability:** Every "Silent Killer" detection rule in the Overseer maps back to a specific `[US-XXX]` user story.
* **ADRs:** Critical decisions (e.g., "Next.js vs Datastar") are recorded in `docs/design/adr/`.

## 4. Quality Assurance
*Ref: [QUALITY.md](../../docs/kit/QUALITY.md)*

* **Table-Driven Tests (TDT):** The Signet Engine's decision logic (Pass/Fail) is tested using extensive TDTs covering edge cases.
* **Hierarchical Fixtures:** Test data for "Compliance Scenarios" is generated via Cuelang composition, not static JSON files.
* **Local-First:** The entire platform runs in `local` mode using `docker-compose`, `SQLite`, and `Ollama` (optional) or mocked LLM adapters.

## 5. Stack Selection
*Ref: [STACK.md](../../docs/kit/STACK.md)*

* **Backend:** Golang (1.23+) for the Engine and CLI.
* **Frontend (Console):**
    * **Decision:** **Next.js (App Router)** for the standalone Console to leverage React Flow (Merkle visualization).
    * **Styling:** Tailwind CSS via **TailwindPlus**.
* **Persistence:**
    * **Local:** SQLite (embedded).
    * **Production:** QLDB (Ledger) + Postgres (Inventory).

## 6. AI Patterns
*Ref: [OUTLINE.md](../../docs/kit/OUTLINE.md) (AI Integration)*

* **LLM as Adapter:** The AI is treated as an infrastructure dependency.
* **State-First:** The Overseer does not "chat." It ingests code/logs and outputs **Structured JSON/Cuelang** that validates against our schema.
