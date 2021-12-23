package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	config      *Config
	Projects    = make(map[string]*Project)
	Port        int
	AdminSecret string
)

func LoadConfig() error {
	f, err := os.ReadFile("./config.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return err
	}
	for _, project := range config.Projects {
		Projects[project.ID] = project
	}
	Port = config.Port
	AdminSecret = config.AdminSecret
	return nil
}
