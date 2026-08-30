// Package config provides utilities for reading and writing application configuration.
package config

import (
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func BuildConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullpath := filepath.Join(home, configFileName)
	return fullpath, nil
}

func (c *Config) SetUser(name string) error {
	
	c.CurrentUserName = name
	if err := Write(c); err != nil {
		return err
	}
	return nil
}
