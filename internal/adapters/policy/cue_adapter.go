package policy

import (
	"context"
	"fmt"
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/geoffmilleraz/signet/internal/core/domain"
)

type CUEAdapter struct {
	ctx *cue.Context
}

func NewCUEAdapter() *CUEAdapter {
	return &CUEAdapter{
		ctx: cuecontext.New(),
	}
}

func (a *CUEAdapter) Unify(ctx context.Context, globalPolicy, userPatch []byte) (domain.Policy, error) {
	gp := a.ctx.CompileBytes(globalPolicy)
	if gp.Err() != nil {
		return domain.Policy{}, fmt.Errorf("failed to compile global policy: %w", gp.Err())
	}

	up := a.ctx.CompileBytes(userPatch)
	if up.Err() != nil {
		return domain.Policy{}, fmt.Errorf("failed to compile user patch: %w", up.Err())
	}

	merged := gp.Unify(up)
	if merged.Err() != nil {
		return domain.Policy{}, fmt.Errorf("policy unification failed: %w", merged.Err())
	}

	if err := merged.Validate(); err != nil {
		return domain.Policy{}, fmt.Errorf("merged policy validation failed: %w", err)
	}

	// For now, return a simplified domain.Policy
	return domain.Policy{
		Version: "v1", // Extract from CUE if possible
	}, nil
}

func (a *CUEAdapter) Validate(ctx context.Context, data []byte, policy domain.Policy) error {
	// Implementation for validating data against a compiled CUE value
	return nil
}
