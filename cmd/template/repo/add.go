package template_repo 

import (
	"fmt"
	"frate-go/config"
	"github.com/spf13/cobra"
)

var TemplateRepoAddCMD = &cobra.Command{
	Use:   "add [repo-url]",
	Short: "Add a new template repository",
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]
		
		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}
		
		metadata.Repos.AdditionalRepos = append(metadata.Repos.AdditionalRepos, repoURL)
		
		if err := config.SaveMetadata(metadata); err != nil {
			fmt.Printf("Error saving metadata: %v\n", err)
			return
		}

		fmt.Printf("Repository %s added successfully!\n", repoURL)
	},
}

func init() {
	TemplateRepoCMD.AddCommand(TemplateRepoAddCMD)
}

