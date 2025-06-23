package feature

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecureTokenExam() (string, error) {
	bytes := make([]byte, 12)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
