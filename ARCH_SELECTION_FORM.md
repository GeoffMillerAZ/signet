# Initial Architecture & Stack Selection Pass

**Goal:** Standardize the initial architectural choices for the project. 
**Instructions:** Review the recommended options below. They are pre-selected based on the "Gold Standard" for this use case. If you agree, leave as is. If you need to override, change the selection or use the "Free-form" section.

---

## 1. Project Context
*   **Project Name:** Signet Platform
*   **Primary Goal:** Autonomous Governance & Cryptographic Seal Engine
*   **Scale Expectation:** Large/Global (Enterprise Governance)

## 2. Frontend Strategy
*Select all that apply for different sub-domains.*

- [x] **Next.js (App Router):** Recommended for Public SaaS, SEO, and complex web apps.
- [ ] **Datastar / Gonads:** Recommended for real-time dashboards and hypermedia-driven UIs.
- [ ] **Vite + React:** Recommended for local developer tools and embedded admin panels.
- [x] **TailwindPlus MCP:** MANDATORY for all UI/UX. Use for Application UI and Marketing components to ensure high-fidelity design.
- [ ] **Static Export:** Recommended for blogs and brochure sites (99/100 Lighthouse target).

**Frontend Overrides / Notes:**
> Use TailwindPlus exclusively to leverage the premium design package and ensure a "cool factor" for the Signet Console. React Flow will be used for Merkle visualization, styled with TailwindPlus patterns.
---

## 3. Backend & Persistence
*Select the primary pattern for the core domain.*

- [ ] **Standard CRUD (SQLite/Postgres):** Recommended for most applications.
- [x] **Event Sourcing:** Recommended for systems requiring high auditability or complex history.
- [x] **CQRS:** Recommended for performance-divergent read/write requirements.

**Backend Overrides / Notes:**
> We are using Event Sourcing + CQRS to ensure the Ledger is the absolute source of truth. gRPC + Protobuf will be used for all internal and service communication. Google Gen AI (Gemini) will be the engine for the Overseer Agent.

---

## 4. Local Development & Demo
- [x] **SQLite + Seeding:** Recommended for local dev and demo modes.
- [x] **Role Selection Page:** Recommended for skipping login in non-prod environments.

---

## 5. Free-form Architectural Decisions
*Use this section for decisions not covered above or for "Hybrid" architectural plans.*

> 1. **Communication:** gRPC/Protobuf for CLI <-> Engine <-> Overseer.
> 2. **AI Integration:** Google Gen AI SDK with Structured Outputs.
> 3. **Visualization:** React Flow for Merkle Tree navigation.
> 4. **Local LLM:** Support for Ollama as a cost-effective/private local Overseer option.

---

## 6. Real-time Clarity Check
*Are there any specific areas where you need more guidance or where the domain complexity is high?*

> High focus on the Event Sourcing implementation for the Ledger to ensure Merkle proof consistency.
