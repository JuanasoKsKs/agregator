package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const(
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}




func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	c.write()
	return nil
}
//Is used by SetUser()
func (c *Config) write() error {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	configPath, err := getConfigFilepaht()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, result, 0644)
	if err != nil {
		return err
	}
	return nil
}
func Read() (Config, error) {
	configPath, err := getConfigFilepaht()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
// Is used by Read()
func getConfigFilepaht() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(home, configFileName)

	return configPath, nil
}
