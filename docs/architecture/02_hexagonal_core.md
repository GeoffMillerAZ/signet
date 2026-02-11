# ⬡ Hexagonal Core: Ports & Adapters

**Status:** Draft | **Version:** 1.0

Signet strictly adheres to **Hexagonal Architecture**. This ensures that our Core Domain Logic is isolated from infrastructure concerns (GitHub API, LLM Providers, Databases).

## 1. The Domain Core (`internal/core`)
The Core contains the *Business Rules* and *Entities*. It has **zero dependencies** on external libraries.

### Entities
* **Artifact:** Represents a build output (Container Image + SHA).
* **Seal:** A cryptographic object linking an Artifact to a specific Policy Version and Validation Result.
* **PolicyBundle:** A collection of Cuelang constraints.
* **Evidence:** A structured record of a test result or LLM analysis.

### Services
* **PromotionService:** Orchestrates the movement from Artifact Repo to Deployment Repo.
* **ValidationService:** Executes the Cuelang Unification logic.
* **SealService:** Calculates the Merkle Hash and commits to the Ledger.

## 2. Ports (`internal/core/ports`)
Interfaces that define *how* the Core interacts with the outside world.

```go
// Port Definition Example

type LedgerPort interface {
    // WriteSeal commits a seal to the immutable storage.
    WriteSeal(ctx context.Context, seal domain.Seal) error

    // VerifySeal checks if a seal exists and is valid.
    VerifySeal(ctx context.Context, artifactHash string) (bool, error)
}

type LLMProviderPort interface {
    // Analyze returns a structured analysis of the provided code diff.
    Analyze(ctx context.Context, diff string, rules []string) (domain.AnalysisResult, error)
}

type GitProviderPort interface {
    // GetWorkflowContent returns the raw content of the CI workflow file.
    GetWorkflowContent(ctx context.Context, repo, path, ref string) ([]byte, error)

    // CreatePullRequest creates a promotion PR in the target repo.
    CreatePullRequest(ctx context.Context, req domain.PRRequest) error
}
```

3. Adapters (internal/adapters)
Implementations of the Ports. These are injected into the Core at runtime (wiring phase).

Infrastructure Adapters
github_adapter: Uses go-github to talk to GitHub API.

openai_adapter: Connects to OpenAI/Gemini/Anthropic for the Overseer.

qldb_adapter: Connects to AWS QLDB for production ledger.

sqlite_adapter: Connects to a local SQLite file for Local/Demo Mode.

Interface Adapters
cli_adapter: The Cobra CLI commands that drive the Core.

http_adapter: The REST/gRPC handlers for the Console and Webhooks.

4. Wiring (Dependency Injection)
Wiring happens only in cmd/signet/main.go.

```go
func main() {
    // 1. Initialize Adapters
    ledger := sqlite.NewAdapter("signet.db")
    llm := openai.NewAdapter(apiKey)
    git := github.NewAdapter(token)

    // 2. Inject into Services (Core)
    sealer := services.NewSealService(ledger)
    overseer := services.NewOverseerService(llm)

    // 3. Start Application
    app.Run(sealer, overseer, git)
}
```

