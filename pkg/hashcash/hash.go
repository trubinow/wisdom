package hashcash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

// formatNonceBytes formats the nonce as []byte, avoiding string conversions.
// It iteratively divides the nonce value by 10 and sets the corresponding byte value as the ASCII representation of the digit
func formatNonceBytes(nonceBytes []byte, nonce int) []byte {
	for i := len(nonceBytes) - 1; i >= 0; i-- {
		nonceBytes[i] = '0' + byte(nonce%10)
		nonce /= 10
		if nonce == 0 {
			break
		}
	}
	return nonceBytes
}

func GenerateHashToken(resource []byte, difficulty int) (string, error) {
	for {
		nonce, err := GenerateNonce()
		if err != nil {
			return "", err
		}
		data := append(resource, nonce...)
		hash := sha256.Sum256(data)

		if hasPrefixZeroes(hash[:], difficulty) {
			return fmt.Sprintf("%x@%x", hash[:], nonce), nil
		}
	}

	return "", errors.New("hash not found")
}

func hasPrefixZeroes(hash []byte, difficulty int) bool {
	for i := 0; i < difficulty; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

// VerifyHashToken verifies the validity of a hash-cash token
func VerifyHashToken(token string, resource string, difficulty int) bool {
	parts := strings.Split(token, "@")
	if len(parts) != 2 {
		return false
	}

	data := resource + fmt.Sprintf("%x", parts[1])
	hashString := parts[0]

	hash := sha256.Sum256([]byte(data))
	hashStringCalculated := hex.EncodeToString(hash[:])

	if hasPrefixZeroes(hash[:], difficulty) && hashString == hashStringCalculated {
		return true
	}

	return false
}
