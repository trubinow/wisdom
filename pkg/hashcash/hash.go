package hashcash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateNonce() ([]byte, error) {
	nonce := make([]byte, 8)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// GenerateHashToken generates a hash-cash token for the given resource and difficulty
func GenerateHashToken(resource string, difficulty int) (string, error) {
	prefix := strings.Repeat("0", difficulty)
	for {
		nonce, err := GenerateNonce()
		if err != nil {
			return "", err
		}

		data := resource + fmt.Sprintf("%x", nonce)
		hash := sha256.Sum256([]byte(data))
		hashString := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashString, prefix) {
			return fmt.Sprintf("%s@%s", hashString, nonce), nil
		}
	}
}

// VerifyHashToken verifies the validity of a hash-cash token
func VerifyHashToken(token string, resource string, difficulty int) bool {
	parts := strings.Split(token, "@")
	if len(parts) != 2 {
		return false
	}

	prefix := strings.Repeat("0", difficulty)
	data := resource + fmt.Sprintf("%x", parts[1])
	hashString := parts[0]

	hash := sha256.Sum256([]byte(data))
	hashStringCalculated := hex.EncodeToString(hash[:])

	if strings.HasPrefix(hashStringCalculated, prefix) && hashString == hashStringCalculated {
		return true
	}

	return false
}
