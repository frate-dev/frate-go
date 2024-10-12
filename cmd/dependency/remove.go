package dependency

import (
	"frate-go/config"
	"frate-go/utils"
	"log"

	"github.com/spf13/cobra"
)

var DepRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Short: "Remove Dependency from your project",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		if err != nil {
			log.Fatal(err)
		}
		for index, _ := range cfg.Dependencies {
			for _, arg := range args {
				if arg == cfg.Dependencies[index].Name {
					cfg.Dependencies = utils.RemoveIndex(cfg.Dependencies, index)
				}
			}

		}
		config.GenerateConfig(cfg)
	},
}

func init() {
}
