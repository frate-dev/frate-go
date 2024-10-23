package config

import (
	"fmt"
	"frate-go/dependency"
  "path/filepath"
	"gopkg.in/yaml.v3"
	"os"
)

var configPath = filepath.Join(os.Getenv("HOME"), ".local", "share", "frate", "metadata.yaml")

type Dep = dependency.Dependency


type Metadata struct {
	Repos struct {
		Default         string   `yaml:"default"`
		AdditionalRepos []string `yaml:"additional_repos"`
	} `yaml:"repos"`
}


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
		return cfg, fmt.Errorf("failed to open config.yaml: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to decode config.yaml: %w", err)
	}
	return cfg, nil
}


func LoadMetadata() (*Metadata, error) {
	var metadata Metadata
	
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("metadata file does not exist: %v", configPath)
	}
	
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening metadata file: %v", err)
	}
	defer file.Close()
	
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&metadata); err != nil {
		return nil, fmt.Errorf("error decoding metadata file: %v", err)
	}

	return &metadata, nil
}

func SaveMetadata(metadata *Metadata) error {
	
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error creating metadata file: %v", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(metadata); err != nil {
		return fmt.Errorf("error encoding metadata: %v", err)
	}

	return nil
}
