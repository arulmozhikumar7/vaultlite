package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arulmozhikumar7/vaultlite/internal/cipher"
)

const secretsFileName = "secrets.json"

var ErrKeyNotFound = errors.New("key not found")

type SecretEntry struct {
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func getSecretsFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".vaultlite", secretsFileName), nil
}

func loadSecrets() (map[string]SecretEntry, error) {
	secrets := make(map[string]SecretEntry)

	path, err := getSecretsFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return secrets, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &secrets)
	if err != nil {
		return nil, err
	}

	return secrets, nil
}

func saveSecrets(secrets map[string]SecretEntry) error {
	path, err := getSecretsFilePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(secrets, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func AddSecret(key, value string) error {
	secrets, err := loadSecrets()
	if err != nil {
		return err
	}

	encKey, err := encryption.GetAESEncrypted(key)
	if err != nil {
		return err
	}
	encValue, err := encryption.GetAESEncrypted(value)
	if err != nil {
		return err
	}

	now := time.Now().Local().Format(time.RFC1123)

	secrets[encKey] = SecretEntry{
		Value:     encValue,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return saveSecrets(secrets)
}

func UpdateSecret(key, newValue string) error {
	secrets, err := loadSecrets()
	if err != nil {
		return err
	}

	var foundKey string
	for encKey := range secrets {
		decKey, err := encryption.GetAESDecrypted(encKey)
		if err != nil {
			continue
		}
		if decKey == key {
			foundKey = encKey
			break
		}
	}
	if foundKey == "" {
		return ErrKeyNotFound
	}

	encValue, err := encryption.GetAESEncrypted(newValue)
	if err != nil {
		return err
	}

	entry := secrets[foundKey]
	entry.Value = encValue
	entry.UpdatedAt = time.Now().Local().Format(time.RFC1123)

	secrets[foundKey] = entry
	return saveSecrets(secrets)
}

func GetSecret(key string, showMeta bool) (string, error) {
	secrets, err := loadSecrets()
	if err != nil {
		return "", err
	}

	for encKey, entry := range secrets {
		decKey, err := encryption.GetAESDecrypted(encKey)
		if err != nil {
			continue
		}
		if decKey == key {
			decVal, err := encryption.GetAESDecrypted(entry.Value)
			if err != nil {
				return "", err
			}
			if showMeta {
				return fmt.Sprintf("Value: %s\nCreated At: %s\nUpdated At: %s", decVal, entry.CreatedAt, entry.UpdatedAt), nil
			}
			return decVal, nil
		}
	}

	return "", ErrKeyNotFound
}

func RemoveSecret(key string) error {
	secrets, err := loadSecrets()
	if err != nil {
		return err
	}

	var foundKey string
	for encKey := range secrets {
		decKey, err := encryption.GetAESDecrypted(encKey)
		if err != nil {
			continue
		}
		if decKey == key {
			foundKey = encKey
			break
		}
	}
	if foundKey == "" {
		return ErrKeyNotFound
	}

	delete(secrets, foundKey)
	return saveSecrets(secrets)
}

func ListSecrets() ([]string, error) {
	secrets, err := loadSecrets()
	if err != nil {
		return nil, err
	}

	var keys []string

	for encKey := range secrets {
		decKey, err := encryption.GetAESDecrypted(encKey)
		if err != nil {
			continue 
		}
		keys = append(keys, decKey)
	}

	return keys, nil
}
