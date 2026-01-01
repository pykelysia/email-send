package config

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/pykelysia/pyketools"
)

func LoadConfig(filePath string) Config {
	file, err := os.Open(filePath)
	if err != nil {
		pyketools.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		pyketools.Fatalf("Failed to decode config file: %v", err)
	}

	return config
}
