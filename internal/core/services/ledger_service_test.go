package services

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/geoffmilleraz/signet/internal/adapters/ledger"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

func TestEventSourcing_Projection(t *testing.T) {
	dbFile := "test_ledger.db"
	defer os.Remove(dbFile)

	adapter, err := ledger.NewSQLiteEventStoreAdapter(dbFile)
	if err != nil {
		t.Fatalf("failed to create adapter: %v", err)
	}

	projectionService := NewProjectionService(adapter)

	t.Run("Append and Project Artifact", func(t *testing.T) {
		artifact := domain.Artifact{
			ID:        "art-123",
			SHA:       "sha256:abc",
			Repo:      "signet",
			Branch:    "main",
			CreatedAt: time.Now(),
		}

		data, _ := json.Marshal(domain.ArtifactCreatedData{Artifact: artifact})
		event := domain.Event{
			ID:        "evt-1",
			Type:      domain.EventArtifactCreated,
			StreamID:  artifact.SHA,
			Data:      data,
			Timestamp: time.Now(),
			Version:   1,
		}

		// 1. Append to Event Store
		if err := adapter.Append(context.Background(), event); err != nil {
			t.Fatalf("failed to append event: %v", err)
		}

		// 2. Project to Read Model
		if err := projectionService.Project(context.Background(), event); err != nil {
			t.Fatalf("failed to project event: %v", err)
		}

		// 3. Verify Read Model (LedgerPort)
		// We can't directly query artifacts from adapter without a GetArtifact method, 
		// but we can test Seal which has a GetSeal method.
	})

	t.Run("Append and Project Seal", func(t *testing.T) {
		seal := domain.Seal{
			ID:               "seal-456",
			ArtifactSHA:      "sha256:abc",
			ConfigHash:       "conf-1",
			PolicyHash:       "pol-1",
			EvidenceRootHash: "root-1",
			Verdict:          domain.VerdictPass,
			Timestamp:        time.Now(),
			Issuer:           "engine-1",
			Signature:        []byte("sig"),
		}

		data, _ := json.Marshal(domain.SealGeneratedData{Seal: seal})
		event := domain.Event{
			ID:        "evt-2",
			Type:      domain.EventSealGenerated,
			StreamID:  seal.ID,
			Data:      data,
			Timestamp: time.Now(),
			Version:   1,
		}

		if err := adapter.Append(context.Background(), event); err != nil {
			t.Fatalf("failed to append event: %v", err)
		}

		if err := projectionService.Project(context.Background(), event); err != nil {
			t.Fatalf("failed to project event: %v", err)
		}

		// Verify via GetSeal
		got, err := adapter.GetSeal(context.Background(), seal.ID)
		if err != nil {
			t.Fatalf("failed to get seal from read model: %v", err)
		}
		if got.ID != seal.ID {
			t.Errorf("expected seal ID %s, got %s", seal.ID, got.ID)
		}
	})
}
