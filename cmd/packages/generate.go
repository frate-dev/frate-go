package packages

import (
	"bufio"
	"encoding/json"
	"fmt"
	"frate-go/config"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Helper function to prompt user for missing data
func promptUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Generate the push.yaml file based on the data from Config
func generatePackage() Package {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	pkg := Package{
		Name:         cfg.ProjectName,
		Version:      cfg.ProjectVersion,
		Description:  promptUserIfEmpty(cfg.ProjectName, "Enter a description for the package"),
		GitURL:       promptUser("Enter the Git URL for the package"),
		License:      promptUser("Enter the license for the package (e.g., MIT, Apache 2.0)"),
		CMakeTarget:  promptUser("Enter the proper way for CMake link to your lib"),
		Dependencies: extractDependencies(cfg.Dependencies),
		Features:     map[string]Feature{},
	}

	// Set the LastModified field to the current time
	pkg.LastModified = time.Now().UTC().Format(time.RFC3339)

	// Add features if necessary
	if addFeatures := promptUser("Do you want to add features? (yes/no)"); strings.ToLower(addFeatures) == "yes" {
		// Add features to the package
		for {
			featureName := promptUser("Enter feature name (or leave empty to finish): ")
			if featureName == "" {
				break
			}
			feature := Feature{
				Description: promptUser("Enter description for the feature"),
				Dependencies: strings.Split(
					promptUser("Enter feature dependencies (comma-separated)"), ",",
				),
			}
			pkg.Features[featureName] = feature
		}
	}

	return pkg
}

func promptUserIfEmpty(fieldValue, prompt string) string {
	if fieldValue != "" {
		return fieldValue
	}
	return promptUser(prompt)
}

func extractDependencies(deps []config.Dep) []string {
	dependencyNames := []string{}
	for _, dep := range deps {
		dependencyNames = append(dependencyNames, dep.Name)
	}
	return dependencyNames
}

func writeYAMLToFile(filename string, pkg Package) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(&pkg); err != nil {
		fmt.Printf("Error encoding YAML: %v\n", err)
	}
}

var PackageGenerateCmd = &cobra.Command{
	Use:  "generate",
	Aliases: []string{"gen"},
	Short: "Generate yaml to allow push to package repo",
	Run: func(cmd *cobra.Command, args []string) {
		pkg := generatePackage()
		writeYAMLToFile("push.yaml", pkg)

		output, _ := json.MarshalIndent(pkg, "", "  ")
		fmt.Println("Generated package:")
		fmt.Println(string(output))
	},
}

func init(){

	PackageCMD.AddCommand(PackageGenerateCmd) 
}
