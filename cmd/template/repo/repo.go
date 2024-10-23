package template_repo 

import "github.com/spf13/cobra"

var TemplateRepoCMD = &cobra.Command{
	Use: "repo",
  Aliases: []string{"t"},
}


func init(){
  TemplateRepoCMD.AddCommand(TemplateRepoAddCMD)
  TemplateRepoCMD.AddCommand(TemplateRepoRemoveCMD)
}
