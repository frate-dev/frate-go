package package_repo

import (
	"fmt"
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)

var PackageRepoListCmd = &cobra.Command{
	Use:   "list",
	Aliases:  []string{"ls"},
	Short: "List all package repositories", 
	Run: func(cmd *cobra.Command, args []string) {


		metadata, err := config.LoadMetadata()
		if err != nil {
			log.Fatal("Error loading metadata: \n\t", err)
			return

		}

		fmt.Println("Default Repository:") 
		fmt.Println(metadata.Packages.Default)  

		fmt.Println("Additional Repositories:") 
		for _, repo := range metadata.Packages.AdditionalRepos { 
			fmt.Println(repo) 
		}
	},
}
