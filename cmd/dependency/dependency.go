package dependency 

import (
	"github.com/spf13/cobra"
)

var DependencyCmd = &cobra.Command{
	Use:     "dependency",
	Aliases: []string{"d", "dep"},
	Short:   "Manage your dependencies",
}

func init() {
	DependencyCmd.AddCommand(DepAddCmd)
	DependencyCmd.AddCommand(DepRemoveCmd)
}
