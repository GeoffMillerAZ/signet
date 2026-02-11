package ledger

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/geoffmilleraz/signet/internal/core/domain"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteEventStoreAdapter struct {
	db *sql.DB
}

func NewSQLiteEventStoreAdapter(dsn string) (*SQLiteEventStoreAdapter, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite: %w", err)
	}

	// Create events table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		type TEXT NOT NULL,
		stream_id TEXT NOT NULL,
		data BLOB NOT NULL,
		timestamp DATETIME NOT NULL,
		version INTEGER NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_stream_id ON events(stream_id);

	-- Read Model Tables (CQRS Projections)
	CREATE TABLE IF NOT EXISTS artifacts (
		id TEXT PRIMARY KEY,
		sha TEXT NOT NULL,
		repo TEXT NOT NULL,
		branch TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);
	CREATE TABLE IF NOT EXISTS seals (
		id TEXT PRIMARY KEY,
		artifact_sha TEXT NOT NULL,
		config_hash TEXT NOT NULL,
		policy_hash TEXT NOT NULL,
		evidence_root_hash TEXT NOT NULL,
		verdict TEXT NOT NULL,
		timestamp DATETIME NOT NULL,
		issuer TEXT NOT NULL,
		signature BLOB NOT NULL
	);
	CREATE TABLE IF NOT EXISTS evidence (
		id TEXT PRIMARY KEY,
		seal_id TEXT NOT NULL,
		type TEXT NOT NULL,
		payload BLOB NOT NULL,
		created_at DATETIME NOT NULL
	);
	`
	if _, err := db.Exec(query); err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	return &SQLiteEventStoreAdapter{db: db}, nil
}

// LedgerPort Implementation (The Read Model)

func (a *SQLiteEventStoreAdapter) SaveArtifact(ctx context.Context, artifact domain.Artifact) error {
	query := `INSERT OR REPLACE INTO artifacts (id, sha, repo, branch, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := a.db.ExecContext(ctx, query, artifact.ID, artifact.SHA, artifact.Repo, artifact.Branch, artifact.CreatedAt)
	return err
}

func (a *SQLiteEventStoreAdapter) SaveSeal(ctx context.Context, seal domain.Seal) error {
	query := `INSERT OR REPLACE INTO seals (id, artifact_sha, config_hash, policy_hash, evidence_root_hash, verdict, timestamp, issuer, signature) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.db.ExecContext(ctx, query, seal.ID, seal.ArtifactSHA, seal.ConfigHash, seal.PolicyHash, seal.EvidenceRootHash, string(seal.Verdict), seal.Timestamp, seal.Issuer, seal.Signature)
	return err
}

func (a *SQLiteEventStoreAdapter) SaveEvidence(ctx context.Context, e domain.Evidence) error {
	query := `INSERT OR REPLACE INTO evidence (id, seal_id, type, payload, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := a.db.ExecContext(ctx, query, e.ID, e.SealID, e.Type, e.Payload, e.CreatedAt)
	return err
}

func (a *SQLiteEventStoreAdapter) GetSeal(ctx context.Context, sealID string) (domain.Seal, error) {
	query := `SELECT id, artifact_sha, config_hash, policy_hash, evidence_root_hash, verdict, timestamp, issuer, signature FROM seals WHERE id = ?`
	var s domain.Seal
	var verdict string
	err := a.db.QueryRowContext(ctx, query, sealID).Scan(&s.ID, &s.ArtifactSHA, &s.ConfigHash, &s.PolicyHash, &s.EvidenceRootHash, &verdict, &s.Timestamp, &s.Issuer, &s.Signature)
	if err != nil {
		return domain.Seal{}, err
	}
	s.Verdict = domain.Verdict(verdict)
	return s, nil
}

func (a *SQLiteEventStoreAdapter) GetEvidence(ctx context.Context, sealID string) ([]domain.Evidence, error) {
	query := `SELECT id, seal_id, type, payload, created_at FROM evidence WHERE seal_id = ?`
	rows, err := a.db.QueryContext(ctx, query, sealID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var evidence []domain.Evidence
	for rows.Next() {
		var e domain.Evidence
		if err := rows.Scan(&e.ID, &e.SealID, &e.Type, &e.Payload, &e.CreatedAt); err != nil {
			return nil, err
		}
		evidence = append(evidence, e)
	}
	return evidence, nil
}

// EventStorePort Implementation

func (a *SQLiteEventStoreAdapter) Append(ctx context.Context, event domain.Event) error {
	query := `INSERT INTO events (id, type, stream_id, data, timestamp, version) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := a.db.ExecContext(ctx, query, event.ID, event.Type, event.StreamID, event.Data, event.Timestamp, event.Version)
	if err != nil {
		return fmt.Errorf("failed to append event: %w", err)
	}
	return nil
}

func (a *SQLiteEventStoreAdapter) GetStream(ctx context.Context, streamID string) ([]domain.Event, error) {
	query := `SELECT id, type, stream_id, data, timestamp, version FROM events WHERE stream_id = ? ORDER BY version ASC`
	rows, err := a.db.QueryContext(ctx, query, streamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var e domain.Event
		if err := rows.Scan(&e.ID, &e.Type, &e.StreamID, &e.Data, &e.Timestamp, &e.Version); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
