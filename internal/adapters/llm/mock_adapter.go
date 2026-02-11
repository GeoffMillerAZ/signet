package llm

import (
	"context"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

type MockLLMAdapter struct{}

func NewMockLLMAdapter() *MockLLMAdapter {
	return &MockLLMAdapter{}
}

func (a *MockLLMAdapter) Analyze(ctx context.Context, diff []byte, p domain.Policy) (domain.Evidence, error) {
	return domain.Evidence{
		Findings: []domain.Finding{
			{
				Type:     "RELIABILITY",
				Severity: "HIGH",
				Message:  "Missing SIGTERM handler in main.go (Mock detection)",
				File:     "main.go",
				Line:     42,
			},
		},
	}, nil
}
