package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	config   *Config
	Projects = make(map[string]*Project)
	Port     int
)

func LoadConfig() error {
	if config != nil {
		return errors.New("config already loaded")
	}
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
	return nil
}
