package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port   string `yaml:"port"`
	DBPath string `yaml:"db_path"`
}

func Load(path string) (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		config.Port = ":" + port
	}
	if dbPath, ok := os.LookupEnv("DB_PATH"); ok {
		config.DBPath = dbPath
	}

	return config, nil
}
