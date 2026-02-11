package domain

import (
	"time"
)

// Artifact represents a build output (e.g., container image)
type Artifact struct {
	ID        string
	SHA       string
	Repo      string
	Branch    string
	CreatedAt time.Time
}

// Verdict represents the outcome of a validation
type Verdict string

const (
	VerdictPass   Verdict = "PASS"
	VerdictWarn   Verdict = "WARN"
	VerdictFail   Verdict = "FAIL"
	VerdictShadow Verdict = "SHADOW_PASS"
)

// Finding represents a single semantic or syntactic violation
type Finding struct {
	Type        string
	Severity    string
	File        string
	Line        int
	Message     string
	Remediation string
}

// Evidence is a structured record of a validation tool's output
type Evidence struct {
	ID        string
	SealID    string
	Type      string
	Payload   []byte // JSON or CUE
	Findings  []Finding
	CreatedAt time.Time
}

// Seal is the cryptographic proof of validation
type Seal struct {
	ID               string
	ArtifactSHA      string
	ConfigHash       string
	PolicyHash       string
	EvidenceRootHash string
	Verdict          Verdict
	Timestamp        time.Time
	Issuer           string
	Signature        []byte
}

// Policy defines the constraints for a project
type Policy struct {
	Version string
	Groups  map[string][]string
	Rules   map[string]bool
}
