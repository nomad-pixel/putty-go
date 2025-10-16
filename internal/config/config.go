package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Theme        string `json:"theme"`
	Language     string `json:"language"`
	DefaultPort  int    `json:"default_port"`
	SavePassword bool   `json:"save_password"`
}

func DefaultConfig() *Config {
	return &Config{
		Theme:        "light",
		Language:     "ru",
		DefaultPort:  22,
		SavePassword: false,
	}
}

func (c *Config) Load() error {
	configPath := c.getConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return c.Save()
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

func (c *Config) Save() error {
	configPath := c.getConfigPath()

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (c *Config) getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".putty-go", "config.json")
}
