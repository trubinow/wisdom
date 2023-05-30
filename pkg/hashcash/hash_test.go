package hashcash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestGenerateNonce(t *testing.T) {
	nonce, err := GenerateNonce()
	if err != nil {
		t.Errorf("GenerateNonce failed: %v", err)
	}

	if len(nonce) != 8 {
		t.Errorf("Unexpected nonce length. Expected 8, got %d", len(nonce))
	}
}

func TestGenerateHashToken(t *testing.T) {
	resource := "example"
	difficulty := 4

	token, err := GenerateHashToken(resource, difficulty)
	if err != nil {
		t.Errorf("GenerateHashToken failed: %v", err)
	}

	parts := strings.Split(token, "@")
	if len(parts) != 2 {
		t.Errorf("Invalid token format. Expected 'hash@nonce', got %s", token)
	}

	hashString := parts[0]
	nonce := parts[1]

	prefix := strings.Repeat("0", difficulty)
	data := resource + nonce
	hash := sha256.Sum256([]byte(data))
	hashStringCalculated := hex.EncodeToString(hash[:])

	if !strings.HasPrefix(hashString, prefix) {
		t.Errorf("Invalid hash prefix. Expected prefix: %s, got: %s", prefix, hashString[:difficulty])
	}

	if hashString != hashStringCalculated {
		t.Errorf("Hash verification failed. Expected: %s, got: %s", hashStringCalculated, hashString)
	}
}

func TestVerifyHashToken(t *testing.T) {
	resource := "example"
	difficulty := 4

	token, err := GenerateHashToken(resource, difficulty)
	if err != nil {
		t.Errorf("GenerateHashToken failed: %v", err)
	}

	valid := VerifyHashToken(token, resource, difficulty)
	if !valid {
		t.Errorf("VerifyHashToken failed. Expected: true, got: false")
	}

	// Test with modified resource
	modifiedResource := "modified"
	valid = VerifyHashToken(token, modifiedResource, difficulty)
	if valid {
		t.Errorf("VerifyHashToken failed. Expected: false, got: true")
	}

	// Test with modified hash
	parts := strings.Split(token, "@")
	modifiedHash := fmt.Sprintf("0%s", parts[0][1:])
	modifiedToken := fmt.Sprintf("%s@%s", modifiedHash, parts[1])
	valid = VerifyHashToken(modifiedToken, resource, difficulty)
	if valid {
		t.Errorf("VerifyHashToken failed. Expected: false, got: true")
	}
}
