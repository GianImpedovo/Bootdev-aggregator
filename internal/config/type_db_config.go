package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	file, err := getConfigFile()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(file)
	var config Config
	err = json.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	file, err := getConfigFile()
	if err != nil {
		return err
	}

	err = os.WriteFile(file, data, 0644)

	if err != nil {
		return err
	}
	return nil
}

func getConfigFile() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(path, filaName)
	return fullPath, nil
}
