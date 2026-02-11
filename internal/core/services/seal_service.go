package services

import (
	"context"
	"fmt"
	"time"

	"github.com/geoffmilleraz/signet/internal/core/domain"
	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type SealService struct {
	ledger ports.LedgerPort
	crypto ports.CryptoPort
	merkle *MerkleService
}

func NewSealService(ledger ports.LedgerPort, crypto ports.CryptoPort, merkle *MerkleService) *SealService {
	return &SealService{
		ledger: ledger,
		crypto: crypto,
		merkle: merkle,
	}
}

func (s *SealService) CreateSeal(ctx context.Context, artifact domain.Artifact, evidence []domain.Evidence, policy domain.Policy) (domain.Seal, error) {
	// 1. Calculate Evidence Merkle Root
	groupHashes := make(map[string]string)
	for _, e := range evidence {
		groupHashes[e.Type] = s.crypto.Hash(e.Payload)
	}
	rootHash := s.merkle.CalculateRootHash(groupHashes)

	// 2. Create Seal object
	seal := domain.Seal{
		ID:               fmt.Sprintf("seal-%d", time.Now().UnixNano()),
		ArtifactSHA:      artifact.SHA,
		EvidenceRootHash: rootHash,
		Verdict:          domain.VerdictPass, // Logic to aggregate verdicts from evidence
		Timestamp:        time.Now(),
		Issuer:           "signet-engine-01",
	}

	// 3. Sign the Seal (Mock)
	signature, err := s.crypto.Sign([]byte(seal.ID + seal.ArtifactSHA + seal.EvidenceRootHash))
	if err != nil {
		return domain.Seal{}, err
	}
	seal.Signature = signature

	// 4. Save to Ledger
	if err := s.ledger.SaveSeal(ctx, seal); err != nil {
		return domain.Seal{}, err
	}

	return seal, nil
}
