package template_repo

import (
	"fmt"
	"frate-go/config"
	"github.com/spf13/cobra"
)


var TemplateRepoRemoveCMD = &cobra.Command{
	Use:   "remove [repo-url]",
	Short: "Remove a template repository",
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]

		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}

		var updatedRepos []string
		for _, repo := range metadata.Repos.AdditionalRepos {
			if repo != repoURL {
				updatedRepos = append(updatedRepos, repo)
			}
		}
		metadata.Repos.AdditionalRepos = updatedRepos

		if err := config.SaveMetadata(metadata); err != nil {
			fmt.Printf("Error saving metadata: %v\n", err)
			return
		}

		fmt.Printf("Repository %s removed successfully!\n", repoURL)
	},
}

func init() {
	TemplateRepoCMD.AddCommand(TemplateRepoRemoveCMD)
}

