package template

import (
	"encoding/json"
	"fmt"
	"frate-go/cmd"
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)





var TemplateListCMD = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List available templates", 
	Run: func(command *cobra.Command, args []string) {
		repoflag, _ := command.Flags().GetString("repo") 
		metadata, err := config.LoadMetadata()
		if err != nil {
			log.Fatal("Error loading metadata: \n\t", err)
			return
		}
		url := metadata.Templates.Default.Url + "/list-templates"
		fmt.Println(metadata)
		fmt.Println(url)
		if repo := repoflag; repo != "" { 
			for _, repo := range metadata.Templates.AdditionalRepos { 
				if repo.Name == repoflag { 
					url = repo.Url + "/list-templates" 
				}
			}
		}

		if len(args) != 0 {
			for _, repo := range metadata.Templates.AdditionalRepos {
				if args[0] == repo.Name{
					url = repo.Url + "/list-templates"
				}
			}
		}

		data, err := cmd.Get(url) 
		var prettyData interface{} 
		json.Unmarshal([]byte(data), &prettyData) 	
		if err != nil {
			log.Fatal("Error getting templates: \n\t", err)
		}
		prettyJson, err:=json.MarshalIndent(prettyData, "", "  ")
		fmt.Println( string(prettyJson))
	},
}

func init() {
	TemplateListCMD.Flags().StringP("repo", "r", "", "Name of the repo to list templates from") 
	TemplateCMD.AddCommand(TemplateListCMD)
}

