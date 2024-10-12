package config

import (
	"fmt"
	"frate-go/dependency"
	"gopkg.in/yaml.v3"
	"os"
)

type Dep = dependency.Dependency

type Config struct {
	CMakeVersion    string            `yaml:"cmake_version"`
	ProjectName     string            `yaml:"project_name"`
	Language        string            `yaml:"language"`
	LanguageVersion string            `yaml:"language_version"`
	IncludeDir      string            `yaml:"include_directory"`
	ProjectVersion  string            `yaml:"project_version"`
	Compiler        string            `yaml:"compiler"`
	BuildDir        string            `yaml:"build_directory"`
	SourceFiles     []string          `yaml:"source_files"`
	Options         map[string]string `yaml:"options"`
	Dependencies    []Dep             `yaml:"dependencies"`
}

func GenerateConfig(cfg Config) error {
	file, err := os.Create("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to create config.yaml: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("failed to encode config.yaml: %w", err)
	}
	return nil
}

func ReadConfig() (Config, error ){
  cfg := Config{}
	file, err := os.Open("config.yaml")
	if err != nil {
		return cfg,fmt.Errorf("failed to open config.yaml: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to decode config.yaml: %w", err)
	}
	return cfg, nil
}
