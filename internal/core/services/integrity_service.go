package services

import (
	"context"
	"errors"
	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type IntegrityService struct {
	crypto ports.CryptoPort
	git    ports.GitProviderPort
}

func NewIntegrityService(crypto ports.CryptoPort, git ports.GitProviderPort) *IntegrityService {
	return &IntegrityService{
		crypto: crypto,
		git:    git,
	}
}

func (s *IntegrityService) VerifyWorkflow(ctx context.Context, path string, content []byte, authorizedHash string) error {
	// 1. Canonicalize and Hash (Simplified for now)
	calculatedHash := s.crypto.Hash(content)

	// 2. Compare
	if calculatedHash != authorizedHash {
		// Log security alert (omitted for now)
		return errors.New("Integrity Check Failed")
	}

	return nil
}
