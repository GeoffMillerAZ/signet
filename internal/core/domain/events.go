package domain

import (
	"time"
)

// EventType defines the type of event in the ledger
type EventType string

const (
	EventArtifactCreated EventType = "ARTIFACT_CREATED"
	EventSealGenerated   EventType = "SEAL_GENERATED"
	EventEvidenceAdded   EventType = "EVIDENCE_ADDED"
	EventPromotionLogged EventType = "PROMOTION_LOGGED"
)

// Event represents a single immutable record in the event stream
type Event struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	StreamID  string    `json:"stream_id"` // e.g., artifact SHA or Seal ID
	Data      []byte    `json:"data"`      // JSON serialized domain data
	Timestamp time.Time `json:"timestamp"`
	Version   int       `json:"version"`
}

// ArtifactCreatedData is the payload for EventArtifactCreated
type ArtifactCreatedData struct {
	Artifact Artifact `json:"artifact"`
}

// SealGeneratedData is the payload for EventSealGenerated
type SealGeneratedData struct {
	Seal Seal `json:"seal"`
}

// EvidenceAddedData is the payload for EventEvidenceAdded
type EvidenceAddedData struct {
	Evidence Evidence `json:"evidence"`
}
