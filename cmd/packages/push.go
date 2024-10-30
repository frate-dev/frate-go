package packages

import (
	"bytes"
	"encoding/json"
	"frate-go/config"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Feature struct {
	Description      string   `json:"description"`
	RequiredFeatures []string `json:"required_features,omitempty"`
	Dependencies     []string `json:"dependencies,omitempty"`
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
	Short: "push package to package repository",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		if err != nil {
			log.Fatal("Error reading config: \n\tare you inside your project\n\t", err)
		}
		meta, err := config.LoadMetadata()
		repo := cmd.Flag("repo").Value.String()
		if err != nil {
			log.Fatal("Error reading config: \n\t", err)
		}
		_ = cfg

		url := meta.Packages.Default.Url
		if repo != "" {
			for _, r := range meta.Packages.AdditionalRepos {
				if r.Name == repo {
					url = r.Url
				}
			}
		}
		url += "/packages/create"
		pkg, err := ReadPackage() 
		if err != nil {
			log.Fatal("Error reading package: \n\t", err)
		}

		pkg_data, err:=json.Marshal(pkg) 
		if err != nil {
			log.Fatal("Error marshalling package: \n\t", err)
		}
		http.Post(url, "application/json",  bytes.NewBuffer(pkg_data))

	},
}

func ReadPackage() (Package, error) {
	file, err := os.Open("push.yaml")
	var pkg Package
	if err != nil {
		log.Fatal("Error opening file: \n\t", err)
		return pkg, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file: \n\t", err)
		return pkg, err
	}
	json.Unmarshal(data, &pkg)
	return pkg, nil
}

func init() {
	PackagePushCmd.Flags().StringP("repo", "r", "", "specify the repository to push to")
}
