package cmd

import (
	"fmt"
	"frate-go/config"
	"frate-go/ftemplate"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate/regenerate your project's using CMake config",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ReadConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		ftemplate.GenerateCmake(config)
	},
}
