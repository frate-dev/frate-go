package config

import (
	"fmt"
	"frate-go/dependency"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

var configPath = filepath.Join(os.Getenv("HOME"), ".local", "share", "frate", "metadata.yaml")

type Dep = dependency.Dependency

type TemplateRepo struct {
	Name string `yaml:"default"`
	Url  string `yaml:"url"`
}

type Metadata struct {
	Default         TemplateRepo   `yaml:"default"`
	AdditionalRepos []TemplateRepo `yaml:"additional_repos"`
}

type Config struct {
	CMakeVersion    string            `yaml:"cmake_version"`
	ProjectName     string            `yaml:"project_name"`
	Language        string            `yaml:"language"`
	LanguageVersion string            `yaml:"language_version"`
	IncludeDir      string            `yaml:"include_directory"`
	SourceDir       string            `yaml:"source_directory"`
	ProjectVersion  string            `yaml:"project_version"`
	Template        string            `yaml:"template"`
	Compiler        string            `yaml:"compiler"`
	BuildDir        string            `yaml:"build_directory"`
	SourceFiles     []string          `yaml:"source_files"`
	Options         map[string]string `yaml:"options"`
	Dependencies    []Dep             `yaml:"dependencies"`
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	// Check if the struct field is a slice
	if structFieldType.Kind() == reflect.Slice {
		// Ensure that the provided value is also a slice
		if val.Kind() != reflect.Slice {
			return fmt.Errorf("Cannot assign non-slice value to slice field %s", name)
		}

		// Convert []interface{} to []string if necessary
		if structFieldType.Elem().Kind() == reflect.String {
			stringSlice := make([]string, val.Len())
			for i := 0; i < val.Len(); i++ {
				elem := val.Index(i)
				if str, ok := elem.Interface().(string); ok {
					stringSlice[i] = str
				} else {
					return fmt.Errorf("Cannot assign non-string element to string slice field %s", name)
				}
			}
			structFieldValue.Set(reflect.ValueOf(stringSlice))
			return nil
		}
	}

	// Handle non-slice fields normally
	if structFieldType != val.Type() {
		return fmt.Errorf("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func (s *Config) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		fmt.Printf("Setting field: %s with value: %v\n", k, v) // Debugging output
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
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

func ReadConfig() (Config, error) {
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

func CreateMetadata() error { 
	metadata := Metadata{
		Default: TemplateRepo{
			Name: "default",
			Url:  "http://localhost:8080",
		},
	}
	if err := SaveMetadata(&metadata); err != nil { 
		log.Fatal("Error creating metadata: \n\t", err) 
	}
	return nil
} 


func LoadMetadata() (*Metadata, error) {
	var metadata Metadata

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err :=CreateMetadata()
		if err != nil { 
			log.Fatal("Error creating metadata: \n\t", err) 
		}
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
