package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	fullpath, err := BuildConfigPath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.ReadFile(fullpath)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if err := json.Unmarshal(file, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
