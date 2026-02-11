package services

import (
	"sort"
	"strings"

	"github.com/geoffmilleraz/signet/internal/core/ports"
)

type MerkleService struct {
	crypto ports.CryptoPort
}

func NewMerkleService(crypto ports.CryptoPort) *MerkleService {
	return &MerkleService{crypto: crypto}
}

// CalculateGroupHash calculates a single hash for a collection of file hashes
func (s *MerkleService) CalculateGroupHash(fileHashes map[string]string) string {
	if len(fileHashes) == 0 {
		return ""
	}

	// 1. Sort keys for deterministic hashing
	keys := make([]string, 0, len(fileHashes))
	for k := range fileHashes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 2. Concatenate hashes
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fileHashes[k])
	}

	// 3. Hash the result
	return s.crypto.Hash([]byte(sb.String()))
}

// CalculateRootHash calculates the root of the Merkle Tree from group hashes
func (s *MerkleService) CalculateRootHash(groupHashes map[string]string) string {
	return s.CalculateGroupHash(groupHashes) // Same logic for groups to root
}

// IdentifyGroups (Mock implementation for now)
// In production, this would use the CUE groups definition to map file paths to groups.
func (s *MerkleService) IdentifyGroups(files []string, groups map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for _, file := range files {
		assigned := false
		for groupName, patterns := range groups {
			for _, pattern := range patterns {
				// Simple suffix match for now (e.g., .go, .sql)
				if strings.HasSuffix(file, pattern) {
					result[groupName] = append(result[groupName], file)
					assigned = true
					break
				}
			}
			if assigned {
				break
			}
		}
		if !assigned {
			result["other"] = append(result["other"], file)
		}
	}
	return result
}
