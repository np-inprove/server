package auth

import (
	"crypto/rand"
	"encoding/hex
	"fmt"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 1, 32)
	return hex.EncodeToString(hash), nil
}
