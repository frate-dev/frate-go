package template_repo

import (
	"fmt"
	"frate-go/config"
	"github.com/spf13/cobra"
)

var TemplateRepoListCMD = &cobra.Command{
	Use:     "list",
	Short:   "Add a new template repository",
	Aliases: []string{"list", "ls"},
	Run: func(cmd *cobra.Command, args []string) {

		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}

		fmt.Println(metadata)
	},
}

func init() {
	TemplateRepoCMD.AddCommand(TemplateRepoListCMD)
}
