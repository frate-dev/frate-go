package template

import (
  "frate-go/cmd/template/repo"
  "github.com/spf13/cobra"
)

var TemplateCMD = &cobra.Command{
	Use: "template",
  Aliases: []string{"t"},
}


func init(){
  TemplateCMD.AddCommand(TemplateInitCMD)
  TemplateCMD.AddCommand(template_repo.TemplateRepoCMD)
}
