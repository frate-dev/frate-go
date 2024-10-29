package template

import (
  "frate-go/cmd/template/repo"
  "github.com/spf13/cobra"
)

var TemplateCMD = &cobra.Command{
	Use: "template",
	Short:    "Manage templates, create, list, delete templates and manage template repos", 
  Aliases: []string{"t"},
}

func init(){
  TemplateCMD.AddCommand(TemplateInitCMD)
  TemplateCMD.AddCommand(template_repo.TemplateRepoCMD)
}
