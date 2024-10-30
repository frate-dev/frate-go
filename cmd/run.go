package cmd

import (
	"fmt"
	"frate-go/config"
	"frate-go/ftemplate"
	"log"
	"path"

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
		if cfg.BuildCmd == "" {
			command := path.Join(".", cfg.BuildDir, cfg.ProjectName)
			RunCommand(command)
			return
		}
		RunCommand("bash", "-c", cfg.BuildCmd)

	},
}
