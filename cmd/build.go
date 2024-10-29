package cmd

import (
	"frate-go/config"
	"frate-go/ftemplate"
	"github.com/spf13/cobra"
	"log"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project using CMake",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ReadConfig()
		if err != nil {
			log.Fatal(err)
			return
		}
		ftemplate.GenerateCmake(&config)
		RunCommand("cmake", ".")

	},
}
