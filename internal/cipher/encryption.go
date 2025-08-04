package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/arulmozhikumar7/vaultlite/internal/config"
)

func pkcs7Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func GetAESEncrypted(plaintext string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(cfg.Key)
	if err != nil {
		return "", errors.New("failed to decode AES key from config")
	}

	iv, err := base64.StdEncoding.DecodeString(cfg.IV)
	if err != nil {
		return "", errors.New("failed to decode IV from config")
	}

	if len(key) != 32 || len(iv) != aes.BlockSize {
		return "", errors.New("invalid AES key or IV length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	padded := pkcs7Pad([]byte(plaintext), aes.BlockSize)
	ciphertext := make([]byte, len(padded))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padded)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func GetSHA256Hash(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("cannot hash empty or whitespace-only input")
	}

	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:]), nil
}
