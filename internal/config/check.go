package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
	Salt string `json:"salt"`
}

func ConfigExistsAndValid() (bool, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return false, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // config does not exist
		}
		return false, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return false, err
	}

	if cfg.Key == "" || cfg.IV == "" {
		return false, errors.New("key or IV is missing in config")
	}

	return true, nil
}

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "vaultlite", "config.json"), nil
}
