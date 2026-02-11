package services

import (
	"context"
	"testing"

	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/ledger"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

func TestSealService_CreateSeal(t *testing.T) {
	cryptoAdapter := crypto.NewSHA256Adapter()
	ledgerAdapter := ledger.NewMemoryAdapter()
	merkleService := NewMerkleService(cryptoAdapter)
	service := NewSealService(ledgerAdapter, cryptoAdapter, merkleService)

	artifact := domain.Artifact{
		ID:  "art-1",
		SHA: "sha256:artifact",
	}
	evidence := []domain.Evidence{
		{Type: "LLM_ANALYSIS", Payload: []byte("finding-1")},
		{Type: "TEST_RESULTS", Payload: []byte("pass")},
	}
	policy := domain.Policy{}

	t.Run("Create and Save Seal", func(t *testing.T) {
		seal, err := service.CreateSeal(context.Background(), artifact, evidence, policy)
		if err != nil {
			t.Fatalf("CreateSeal failed: %v", err)
		}

		if seal.ArtifactSHA != artifact.SHA {
			t.Errorf("expected artifact SHA %s, got %s", artifact.SHA, seal.ArtifactSHA)
		}

		if seal.Signature == nil {
			t.Error("expected signature, got nil")
		}

		// Verify it was saved to the ledger
		savedSeal, err := ledgerAdapter.GetSeal(context.Background(), seal.ID)
		if err != nil {
			t.Fatalf("failed to retrieve saved seal: %v", err)
		}
		if savedSeal.ID != seal.ID {
			t.Errorf("saved seal ID mismatch: %s != %s", savedSeal.ID, seal.ID)
		}
	})
}
