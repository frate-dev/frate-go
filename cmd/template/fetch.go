package template

import (
	"frate-go/cmd"
	"github.com/spf13/cobra"
)

var TemplateFetchCMD = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: func(command *cobra.Command, args []string) {
		cmd.RunCommand("git", "clone", args[0])
	},
}

func init() {
}
