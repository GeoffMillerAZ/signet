package services

import (
	"context"
	"testing"

	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/git"
)

func TestIntegrityService_VerifyWorkflow(t *testing.T) {
	cryptoAdapter := crypto.NewSHA256Adapter()
	gitAdapter := git.NewFileAdapter()
	service := NewIntegrityService(cryptoAdapter, gitAdapter)

	tests := []struct {
		name           string
		content        string
		authorizedHash string
		wantErr        bool
	}{
		{
			name:           "Valid Integrity",
			content:        "workflow: valid",
			authorizedHash: cryptoAdapter.Hash([]byte("workflow: valid")),
			wantErr:        false,
		},
		{
			name:           "Tampered Workflow",
			content:        "workflow: tampered",
			authorizedHash: cryptoAdapter.Hash([]byte("workflow: valid")),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.VerifyWorkflow(context.Background(), "test.yml", []byte(tt.content), tt.authorizedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
