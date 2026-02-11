package ledger

import (
	"context"
	"errors"
	"sync"

	"github.com/geoffmilleraz/signet/internal/core/domain"
)

type MemoryAdapter struct {
	mu        sync.RWMutex
	artifacts map[string]domain.Artifact
	seals     map[string]domain.Seal
	evidence  map[string][]domain.Evidence
}

func NewMemoryAdapter() *MemoryAdapter {
	return &MemoryAdapter{
		artifacts: make(map[string]domain.Artifact),
		seals:     make(map[string]domain.Seal),
		evidence:  make(map[string][]domain.Evidence),
	}
}

func (a *MemoryAdapter) SaveArtifact(ctx context.Context, artifact domain.Artifact) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.artifacts[artifact.ID] = artifact
	return nil
}

func (a *MemoryAdapter) SaveSeal(ctx context.Context, seal domain.Seal) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.seals[seal.ID] = seal
	return nil
}

func (a *MemoryAdapter) SaveEvidence(ctx context.Context, evidence domain.Evidence) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.evidence[evidence.SealID] = append(a.evidence[evidence.SealID], evidence)
	return nil
}

func (a *MemoryAdapter) GetSeal(ctx context.Context, sealID string) (domain.Seal, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	seal, ok := a.seals[sealID]
	if !ok {
		return domain.Seal{}, errors.New("seal not found")
	}
	return seal, nil
}

func (a *MemoryAdapter) GetEvidence(ctx context.Context, sealID string) ([]domain.Evidence, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.evidence[sealID], nil
}
