package services

import (
	"context"
	"testing"
	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/policy"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

type MockLLMAdapter struct{}

func (a *MockLLMAdapter) Analyze(ctx context.Context, diff []byte, p domain.Policy) (domain.Evidence, error) {
	if string(diff) == "bad" {
		return domain.Evidence{
			Findings: []domain.Finding{
				{Severity: "HIGH", Message: "Bad code detected"},
			},
		}, nil
	}
	return domain.Evidence{Findings: []domain.Finding{}}, nil
}

func TestValidationService_ValidateDiff(t *testing.T) {
	policyAdapter := policy.NewCUEAdapter()
	llmAdapter := &MockLLMAdapter{}
	cryptoAdapter := crypto.NewSHA256Adapter()
	service := NewValidationService(policyAdapter, llmAdapter, cryptoAdapter)

	globalPolicy := []byte(`#GlobalPolicy: { steps: { scan: { required: true } } }`)
	userPatch := []byte(`#UserPatch: { meta: { name: "test" } }`)

	tests := []struct {
		name    string
		diff    string
		want    domain.Verdict
		wantErr bool
	}{
		{
			name:    "Pass",
			diff:    "good",
			want:    domain.VerdictPass,
			wantErr: false,
		},
		{
			name:    "Fail",
			diff:    "bad",
			want:    domain.VerdictFail,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := service.ValidateDiff(context.Background(), globalPolicy, userPatch, []byte(tt.diff))
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateDiff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
