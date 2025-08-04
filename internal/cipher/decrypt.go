package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/arulmozhikumar7/vaultlite/internal/config"
)

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	padLen := int(data[len(data)-1])
	if padLen > len(data) || padLen == 0 {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-padLen], nil
}

func GetAESDecrypted(cipherTextBase64 string) (string, error) {
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
		return "", errors.New("invalid AES key or IV size")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	unpadded, err := pkcs7Unpad(decrypted)
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}
