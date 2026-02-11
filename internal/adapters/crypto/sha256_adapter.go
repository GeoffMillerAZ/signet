package crypto

import (
	"crypto/sha256"
	"fmt"
)

type SHA256Adapter struct{}

func NewSHA256Adapter() *SHA256Adapter {
	return &SHA256Adapter{}
}

func (a *SHA256Adapter) Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func (a *SHA256Adapter) Sign(data []byte) ([]byte, error) {
	// Mock signature
	return []byte("mock-signature"), nil
}

func (a *SHA256Adapter) Verify(data []byte, signature []byte) bool {
	// Mock verification
	return string(signature) == "mock-signature"
}
