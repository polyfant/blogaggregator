package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {	
	DBURL          string `json:"db_url"`
	Port           int    `json:"port"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gatorconfig.json"), nil
}

func Read() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If file doesn't exist, create a default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := &Config{
			DBURL: "postgresql://postgres:postgres@localhost:5432/blog?sslmode=disable",
			Port:  0,
		}
		if err := defaultConfig.save(); err != nil {
			return nil, err
		}
		return defaultConfig, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) save() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return c.save()
}
