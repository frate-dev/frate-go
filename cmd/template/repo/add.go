package template_repo

import (
	"fmt"
	"frate-go/config"
	"github.com/spf13/cobra"
)

var TemplateRepoAddCMD = &cobra.Command{
	Use:   "add [name] [repo-url]",
	Short: "Add a new template repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		repoURL := args[1]
		repo := config.TemplateRepo{Name: name, Url: repoURL}
		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}

		metadata.Templates.AdditionalRepos = append(metadata.Templates.AdditionalRepos, repo)

		if err := metadata.SaveMetadata(); err != nil {
			fmt.Printf("Error saving metadata: %v\n", err)
			return
		}

		fmt.Printf("Repository %s added successfully!\n", name)
	},
}

func init() {
	TemplateRepoCMD.AddCommand(TemplateRepoAddCMD)
}
