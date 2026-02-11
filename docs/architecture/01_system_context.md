# 🏗️ System Architecture: Context & Container

**Status:** Draft | **Version:** 1.0

This document defines the high-level architecture of the Signet Platform using the **C4 Model**.

## 1. System Context Diagram (Level 1)
Signet sits between the Developer, the Version Control System (GitHub), and the Cloud Infrastructure (Kubernetes/AWS).

```mermaid
C4Context
    title System Context Diagram for Signet Platform

    Person(dev, "Developer", "Commits code, defines Cuelang pipelines, and promotes artifacts.")
    Person(sre, "SRE / Auditor", "Monitors compliance, audits logs, and manages policy.")

    System(signet, "Signet Platform", "Orchestrates governance, validates configuration, and issues cryptographic seals.")

    System_Ext(github, "GitHub / VCS", "Hosts Artifact and Deployment repositories. Runs CI Actions.")
    System_Ext(k8s, "Kubernetes / Cloud", "Runs the workloads. Enforces admission control based on Signet Seals.")
    System_Ext(llm, "LLM Provider", "Analyzes code for semantic reliability risks (The Spirit).")
    System_Ext(ledger, "WORM Ledger", "Stores immutable audit records (QLDB/S3-Lock).")

    Rel(dev, signet, "Promotes Artifacts via CLI")
    Rel(dev, github, "Pushes Code & Config")
    Rel(signet, github, "Verifies Workflow Integrity & Merges PRs")
    Rel(signet, llm, "Sends Code Diffs for Analysis")
    Rel(signet, ledger, "Writes Truth Stamps")
    Rel(k8s, ledger, "Verifies Seals before Pulling Images")
    Rel(sre, signet, "Audits Compliance via Console")
```

2. Container Diagram (Level 2)
Zooming into the Signet Platform, we see the modular monolith structure following Hexagonal Architecture.

```mermaid
C4Container
    title Container Diagram for Signet Platform

    Container(cli, "Signet CLI", "Go / Cobra", "Developer tool for local validation, initialization, and promotion.")

    Container_Boundary(core, "Signet Core (Hexagonal Monolith)") {
        Component(engine, "Signet Engine", "Go", "Orchestrates the verification workflow and Cuelang unification.")
        Component(overseer, "Overseer Agent", "Go / LLM Adapter", "Performs semantic analysis and reliability scoring.")
        Component(registry, "Registry Service", "Go / SQLite/Postgres", "Tracks tool inventory and maturity scores.")
    }

    Container(console, "Signet Console", "Next.js / React Flow", "Visualizes Merkle Trees and Audit Logs.")

    Rel(cli, engine, "gRPC / HTTPS")
    Rel(engine, overseer, "Internal Function Call")
    Rel(engine, registry, "Internal Function Call")
    Rel(console, registry, "Reads Maturity Data")
```
