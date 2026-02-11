# 👁️ Specification: Overseer Agent (The LLM Logic)

**Status:** Draft | **Version:** 1.0

This specification defines the behavior of the **Overseer**, the AI Agent responsible for semantic code analysis and "Silent Killer" detection.

## 1. The Prompt Strategy
The Overseer operates on a **State-First** model. It does not "chat"; it ingests a context and outputs a JSON/Cuelang Evidence Payload.

### 1.1 System Prompt (The Persona)
> "You are the Overseer, a Principal SRE and Security Architect. Your goal is to analyze code changes for operational reliability risks and architectural violations. You do not care about syntax (linters do that). You care about *Semantic Intent* and *Runtime Safety*."

### 1.2 Analysis Rules (The "Silent Killers")
The Overseer checks for specific patterns defined in the Global Policy:
1.  **Signal Handling:** Does the entry point capture `SIGTERM`/`SIGINT` for graceful shutdown?
2.  **Resource Leaks:** Are file handles, database connections, or goroutines properly closed/cancelled?
3.  **Thundering Herd:** Do retry loops include jitter and exponential backoff?
4.  **Hardcoded Secrets:** Are there high-entropy strings or variable names indicating secrets?
5.  **Schema Compatibility:** (If SQL changed) Is there a corresponding migration file?

## 2. The Evidence Payload
The Overseer must output a structured JSON object.

```json
{
  "analysis_id": "uuid",
  "verdict": "PASS | WARN | FAIL",
  "confidence": 0.98,
  "findings": [
    {
      "type": "RELIABILITY",
      "severity": "HIGH",
      "file": "cmd/server/main.go",
      "line": 42,
      "message": "Missing SIGTERM handler. Server will hard-kill active connections.",
      "remediation": "Implement os.Signal channel notification."
    }
  ],
  "architectural_changes": {
    "persistence": false,
    "networking": true
  }
}

```

## 3. Smart Skip (Inheritance)

To save costs and time:

1. **Merkle Check:** The Engine calculates file hashes for architectural groups.
2. **History Lookup:** If the `networking` group hash matches the previous successful Seal, the Overseer **skips** analysis for that group.
3. **Inheritance:** The Engine injects the *previous* findings for that group into the current Evidence Payload.
