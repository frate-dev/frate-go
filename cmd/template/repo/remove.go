package template_repo

import (
	"fmt"
	"frate-go/config"
	"github.com/spf13/cobra"
)


var TemplateRepoRemoveCMD = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a template repository",
	Aliases: []string{"remove", "rm"},
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]

		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}

		var updatedRepos []config.TemplateRepo 
		for _, repo := range metadata.AdditionalRepos {
			if repo.Name != repoName {
				updatedRepos = append(updatedRepos, repo)
			}
		}
		metadata.AdditionalRepos = updatedRepos

		if err := config.SaveMetadata(metadata); err != nil {
			fmt.Printf("Error saving metadata: %v\n", err)
			return
		}

		fmt.Printf("Repository %s removed successfully!\n", repoName)
	},
}

func init() {
	TemplateRepoCMD.AddCommand(TemplateRepoRemoveCMD)
}

