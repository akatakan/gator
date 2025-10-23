package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func SetUser(config Config, username string) error {
	config.Current_user_name = username
	err := write(config)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configFile, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, configFile, 0666)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homePath, configFileName)
	_, err = os.Stat(fullPath)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}
