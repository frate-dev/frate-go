package package_repo

import (
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)

var PackageRepoRemoveCmd = &cobra.Command{
	Use:   "remove",
	Aliases: []string{"rm"}, 
	Short: "build and run your project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		metadata, err := config.LoadMetadata()
		if err != nil {
			log.Fatal("Error loading metadata: \n\t", err)
			return
		}
		var repos []config.PackageRepo
		for _, pkg := range metadata.Packages.AdditionalRepos {
			if pkg.Name != name { 
				repos =   append(repos, pkg)
			}
		}

		metadata.Packages.AdditionalRepos = repos 
		metadata.SaveMetadata() 
	},
}

