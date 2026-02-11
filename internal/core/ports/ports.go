package ports

import (
	"context"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

// LedgerPort defines the interface for immutable storage of seals and evidence
type LedgerPort interface {
	SaveArtifact(ctx context.Context, artifact domain.Artifact) error
	SaveSeal(ctx context.Context, seal domain.Seal) error
	SaveEvidence(ctx context.Context, evidence domain.Evidence) error
	GetSeal(ctx context.Context, sealID string) (domain.Seal, error)
	GetEvidence(ctx context.Context, sealID string) ([]domain.Evidence, error)
}

// LLMProviderPort defines the interface for semantic analysis
type LLMProviderPort interface {
	Analyze(ctx context.Context, diff []byte, policy domain.Policy) (domain.Evidence, error)
}

// GitProviderPort defines the interface for VCS interactions
type GitProviderPort interface {
	GetFile(ctx context.Context, repo, path, ref string) ([]byte, error)
	CreatePR(ctx context.Context, repo, title, body, head, base string) (string, error)
}

// PolicyPort defines the interface for policy unification and validation
type PolicyPort interface {
	Unify(ctx context.Context, globalPolicy, userPatch []byte) (domain.Policy, error)
	Validate(ctx context.Context, data []byte, policy domain.Policy) error
}

// CryptoPort defines the interface for cryptographic operations
type CryptoPort interface {
	Hash(data []byte) string
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, signature []byte) bool
}
