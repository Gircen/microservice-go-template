package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func GetConfig() *Config {

	yamlFile, err := os.ReadFile("config/application.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error convert YAML to data: %v", err)
	}
	return &cfg
}
