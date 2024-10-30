package dependency

import (
	"encoding/json"
	"fmt"
	"frate-go/config"
	"frate-go/dependency"
	"frate-go/cmd"
	"log"

	"github.com/spf13/cobra"
)

type Feature struct {
	Description      string   `json:"description"`
	RequiredFeatures []string `json:"required_features,omitempty"`
	Dependencies     []string `json:"dependencies,omitempty"` // Change to []string
}

// Package struct representing the package structure
type Package struct {
	Name         string             `json:"name"`
	Version      string             `json:"version"`
	Description  string             `json:"description"`
	GitURL       string             `json:"git_url"`
	License      string             `json:"license"`
	Supports     string             `json:"supports,omitempty"`
	Stars        int                `json:"stars"`
	LastModified string             `json:"last_modified"`
	Dependencies []string           `json:"dependencies"` // Change to []string
	Features     map[string]Feature `json:"features,omitempty"`
}

var DepAddCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i", "install"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Short: "Build the project using CMake",
	Run: func(command *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()

		github, _ := command.Flags().GetString("github")
		tag, _ := command.Flags().GetString("tag")
		if err != nil {
			log.Fatal(err)
			return
		}

		dep := dependency.Dependency{
			Name: args[0],
		}

		if github != "" {
			dep.GitURL = github
		}

		if tag != "" {
			dep.Tag = tag
		}
		data, err := cmd.Get("http://localhost:8000/packages")
		var packages []Package
		

		err = json.Unmarshal([]byte(data), &packages)
    if err != nil {
      log.Fatal("error unmarshaling data", err)
    }

		AddDependencyRecursively(dep, packages, &cfg)

		err = config.GenerateConfig(cfg)
    if err != nil {
      log.Fatal("error gererating config", err)
    }
	},
}

func dependencyExists(dep dependency.Dependency, cfg config.Config) bool {
	for _, existingDep := range cfg.Dependencies {
		if existingDep.Name == dep.Name {
			return true
		}
	}
	return false
}

func AddDependencyRecursively(dep dependency.Dependency, packages []Package, cfg *config.Config) {
	if dependencyExists(dep, *cfg) {
		fmt.Printf("Dependency %s already exists, skipping to avoid circular dependency.\n", dep.Name)
		return
	}

	for _, pack := range packages {
		if dep.Name == pack.Name {
			dep.GitURL = pack.GitURL
			dep.Tag = pack.Version

			cfg.Dependencies = append(cfg.Dependencies, dep)
			fmt.Printf("Added dependency: %s, GitURL: %s, Version: %s\n", dep.Name, dep.GitURL, dep.Tag)

			for _, depName := range pack.Dependencies {
				newDep := dependency.Dependency{Name: depName}
				AddDependencyRecursively(newDep, packages, cfg)
			}
			break
		}
	}
}

func init() {
	DepAddCmd.Flags().StringP("github", "g", "", "github url")
	DepAddCmd.Flags().StringP("tag", "t", "", "tag for project")
}
