package cmd

import (
	"fmt"
	"frate-go/config"
	"frate-go/ftemplate"
	"log"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "build and run your project",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		if err != nil {
			log.Fatal(err)
			return
		}
		ftemplate.GenerateCmake(&cfg)
		config.GenerateConfig(cfg)
		RunCommand("cmake", ".")
		RunCommand("make")
    fmt.Println()
		RunCommand("./" + cfg.BuildDir + "/" + cfg.ProjectName)

	},
}
