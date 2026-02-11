package services

import (
	"context"
	"fmt"

	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type PromotionService struct {
	ledger ports.LedgerPort
	git    ports.GitProviderPort
}

func NewPromotionService(ledger ports.LedgerPort, git ports.GitProviderPort) *PromotionService {
	return &PromotionService{
		ledger: ledger,
		git:    git,
	}
}

func (s *PromotionService) Promote(ctx context.Context, sealID, targetEnv string) (string, error) {
	// 1. Verify Seal exists and passed
	seal, err := s.ledger.GetSeal(ctx, sealID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve seal: %w", err)
	}

	if seal.Verdict != "PASS" && seal.Verdict != "WARN" {
		return "", fmt.Errorf("cannot promote artifact with verdict: %s", seal.Verdict)
	}

	// 2. Clone/Prepare Deployment Repo (Mocked in adapter)
	// 3. Update Manifest (Mocked)
	
	// 4. Create PR
	title := fmt.Sprintf("Promote artifact to %s", targetEnv)
	body := fmt.Sprintf("Automated promotion for Seal ID: %s
Artifact SHA: %s", sealID, seal.ArtifactSHA)
	
	prURL, err := s.git.CreatePR(ctx, "deploy-repo", title, body, "promotion-branch", "main")
	if err != nil {
		return "", fmt.Errorf("failed to create promotion PR: %w", err)
	}

	return prURL, nil
}
