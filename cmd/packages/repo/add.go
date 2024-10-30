package package_repo

import (
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)

var PackageRepoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new repository [name] [url]",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		url := args[1]
		metadata, err := config.LoadMetadata()
		if err != nil {
			log.Fatal("Error loading metadata: \n\t", err)
			return

		}

		var packageTemplate config.PackageRepo
		packageTemplate.Name = name
		packageTemplate.Url = url
		metadata.Packages.AdditionalRepos = append(metadata.Packages.AdditionalRepos, packageTemplate)
		if err := metadata.SaveMetadata(); err != nil {
			log.Fatal("Error creating metadata: \n\t", err)
		}
	},
}
