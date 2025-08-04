package encryption

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKeyAndIV(passphrase string, salt []byte) ([]byte, []byte) {
	// 48 bytes: 32 for AES key, 16 for IV
	keyIV := pbkdf2.Key([]byte(passphrase), salt, 100_000, 48, sha256.New)
	return keyIV[:32], keyIV[32:] // AES-256 key, AES IV
}
