package repo


import (
	"fmt"
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)

var SetupCMD = &cobra.Command{
	Use:   "setup",
	Short: "build and run your project",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ReadConfig()
		if err != nil {
			log.Fatal("error", err)
		}
		fmt.Println(config)
	},
}
