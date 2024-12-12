package config

import (
	"encoding/json"
	"os"
)

type Config struct {	
	DBURL          string `json:"db_url"`
	Port           int    `json:"port"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	configPath := ".gatorconfig.json"
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


func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(".gatorconfig.json", data, 0644)
}
