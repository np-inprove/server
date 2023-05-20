package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	saltLen = 32
	threads = 1
	time    = 1
	memory  = 64 * 1024
	keyLen  = 32
)

type Encoded struct {
	// Hash will be converted to Base64 by json.Marshal
	Hash []byte `json:"h"`
	// Salt will be converted to Base64 by json.Marshal
	Salt []byte `json:"s"`

	Time   uint32 `json:"t"`
	Memory uint32 `json:"m"`
	KeyLen uint32 `json:"k"`
}

func CreateEncoded(password string) (Encoded, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return Encoded{}, fmt.Errorf("failed to generate salt: %v", err)
	}

	h := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	return Encoded{
		Hash: h, Salt: salt, Time: time, Memory: memory, KeyLen: keyLen,
	}, nil
}

func VerifyEncoded(e Encoded, target string) (bool, error) {
	targetHash := argon2.IDKey([]byte(target), e.Salt, e.Time, e.Memory, threads, e.KeyLen)

	if subtle.ConstantTimeCompare(e.Hash, targetHash) == 1 {
		return true, nil
	}

	return false, nil
}
