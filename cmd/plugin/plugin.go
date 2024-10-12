package plugin 

import (
	"github.com/spf13/cobra"
)

var PluginCmd = &cobra.Command{
	Use:     "plugin",
	Aliases: []string{"d", "dep"},
	Short:   "Build the project using CMake",
}

func init() {
	PluginCmd.AddCommand(PluginAddCmd)
}
