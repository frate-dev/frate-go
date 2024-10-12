package plugin 

import (

	"github.com/spf13/cobra"
)

var PluginAddCmd = &cobra.Command{
	Use:   "add",
  Aliases: []string{"d", "dep"},
  Args: func(cmd *cobra.Command, args []string) error {
    if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
        return err
    }

    return nil 
  },
	Short: "Build the project using CMake",
  Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}


