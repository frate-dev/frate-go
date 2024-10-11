package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config structure for YAML configuration
type Config struct {
	CMakeVersion    string            `yaml:"cmake_version"`
	ProjectName     string            `yaml:"project_name"`
	Language        string            `yaml:"language"`
	LanguageVersion string            `yaml:"languageVersion"`
	IncludeDir      string            `yaml:"IncludeDir"`
	ProjectVersion  string            `yaml:"ProjectVersion"`
	Compiler        string            `yaml:"compiler"`
	BuildDir        string            `yaml:"build_dir"`
	SourceFiles     []string          `yaml:"source_files"`
	Options         map[string]string `yaml:"options"`
}

// GenerateConfig writes the configuration to a YAML file
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
