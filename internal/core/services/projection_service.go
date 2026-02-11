package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/geoffmilleraz/signet/internal/core/domain"
	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type ProjectionService struct {
	ledger ports.LedgerPort // The Read Model (e.g., SQLite CRUD tables)
}

func NewProjectionService(ledger ports.LedgerPort) *ProjectionService {
	return &ProjectionService{ledger: ledger}
}

// Project handles a single event and updates the Read Model
func (s *ProjectionService) Project(ctx context.Context, event domain.Event) error {
	switch event.Type {
	case domain.EventArtifactCreated:
		var data domain.ArtifactCreatedData
		if err := json.Unmarshal(event.Data, &data); err != nil {
			return err
		}
		return s.ledger.SaveArtifact(ctx, data.Artifact)

	case domain.EventSealGenerated:
		var data domain.SealGeneratedData
		if err := json.Unmarshal(event.Data, &data); err != nil {
			return err
		}
		return s.ledger.SaveSeal(ctx, data.Seal)

	case domain.EventEvidenceAdded:
		var data domain.EvidenceAddedData
		if err := json.Unmarshal(event.Data, &data); err != nil {
			return err
		}
		return s.ledger.SaveEvidence(ctx, data.Evidence)

	default:
		return fmt.Errorf("unknown event type: %s", event.Type)
	}
}
