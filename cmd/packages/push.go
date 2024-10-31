package packages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"frate-go/config"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Feature struct {
	Description      string   `json:"description" yaml:"description"`
	RequiredFeatures []string `json:"required_features,omitempty" yaml:"required_features,omitempty"`
	Dependencies     []string `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}

type Package struct {
	Name         string             `json:"name" yaml:"name"`
	Version      string             `json:"version" yaml:"version"`
	Description  string             `json:"description" yaml:"description"`
	GitURL       string             `json:"git_url" yaml:"git_url"`
	License      string             `json:"license,omitempty" yaml:"license,omitempty"`
	Supports     string             `json:"supports,omitempty" yaml:"supports,omitempty"`
	Stars        int                `json:"stars,omitempty" yaml:"stars,omitempty"`
	LastModified string             `json:"last_modified,omitempty" yaml:"last_modified,omitempty"`
	CMakeTarget  string             `json:"cmake_target,omitempty" yaml:"cmake_target,omitempty"`
	Dependencies []string           `json:"dependencies" yaml:"dependencies"`
	Features     map[string]Feature `json:"features,omitempty" yaml:"features,omitempty"`
}

var PackagePushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push package to package repository",
	Run: func(cmd *cobra.Command, args []string) {

		// Load metadata
		meta, err := config.LoadMetadata()
		if err != nil {
			log.Fatalf("Error loading metadata: \n\t%v", err)
		}

		// Determine repository URL
		repo := cmd.Flag("repo").Value.String()
		url := getRepositoryURL(repo, meta)
		if url == "" {
			log.Fatal("Error: No valid repository URL found")
		}

		// Read and validate package from file
		pkg, err := ReadPackage("push.yaml")
		if err != nil {
			log.Fatalf("Error reading package: \n\t%v", err)
		}

		// Send package data
		if err := pushPackage(url, pkg); err != nil {
			log.Fatalf("Error pushing package: \n\t%v", err)
		}

		fmt.Println("Package successfully pushed.")
	},
}

func getRepositoryURL(repo string, meta *config.Metadata) string {
	if repo != "" {
		for _, r := range meta.Packages.AdditionalRepos {
			if r.Name == repo {
				return r.Url
			}
		}
	}
	return meta.Packages.Default.Url + "/packages/create"
}

func ReadPackage(filename string) (Package, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Package{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var pkg Package
	data, err := io.ReadAll(file)
	if err != nil {
		return pkg, fmt.Errorf("error reading file: %w", err)
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return pkg, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	// Validate package data
	if pkg.Name == "" || pkg.Version == "" {
		return pkg, errors.New("package name and version are required")
	}

	return pkg, nil
}

func pushPackage(url string, pkg Package) error {
	// Marshal package to JSON
	pkgData, err := json.Marshal(pkg)
	if err != nil {
		return fmt.Errorf("error marshalling package: %w", err)
	}

	// Send HTTP POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(pkgData))
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to push package, status: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

func init() {
	PackagePushCmd.Flags().StringP("repo", "r", "", "Specify the repository to push to")
}

