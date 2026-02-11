package services

import (
	"context"
	"github.com/geoffmilleraz/signet/internal/core/domain"
	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type ValidationService struct {
	policy ports.PolicyPort
	llm    ports.LLMProviderPort
	crypto ports.CryptoPort
}

func NewValidationService(policy ports.PolicyPort, llm ports.LLMProviderPort, crypto ports.CryptoPort) *ValidationService {
	return &ValidationService{
		policy: policy,
		llm:    llm,
		crypto: crypto,
	}
}

func (s *ValidationService) ValidateDiff(ctx context.Context, globalPolicy, userPatch, diff []byte) (domain.Verdict, []domain.Finding, error) {
	// 1. Unify Policy
	p, err := s.policy.Unify(ctx, globalPolicy, userPatch)
	if err != nil {
		return domain.VerdictFail, nil, err
	}

	// 2. Run LLM Analysis (Overseer)
	evidence, err := s.llm.Analyze(ctx, diff, p)
	if err != nil {
		return domain.VerdictFail, nil, err
	}

	// 3. Determine Verdict
	verdict := domain.VerdictPass
	for _, f := range evidence.Findings {
		if f.Severity == "HIGH" {
			verdict = domain.VerdictFail
			break
		} else if f.Severity == "MEDIUM" {
			verdict = domain.VerdictWarn
		}
	}

	return verdict, evidence.Findings, nil
}
