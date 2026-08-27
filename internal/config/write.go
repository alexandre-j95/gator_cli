package config

import (
	"encoding/json"
	"os"
)

func Write(conf *Config) error {
	fullpath, err := BuildConfigPath()
	if err != nil {
		return err
	}

	jsondata, err := json.MarshalIndent(conf, "", "	")
	if err != nil {
		return err
	}

	if err = os.WriteFile(fullpath, jsondata, 0o644); err != nil {
		return err
	}
	return nil
}
