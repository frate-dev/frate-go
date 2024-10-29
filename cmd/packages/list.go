package packages

import (
	"frate-go/config"
	"log"

	"github.com/spf13/cobra"
)


var PackageListCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate/regenerate your project's using CMake config",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ReadConfig()
		if err != nil {
			log.Fatal("error loading config", err)
		}
		_ = config
	},
}


func init(){

}
