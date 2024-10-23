package template

import (
	"frate-go/cmd"
	"github.com/spf13/cobra"
)

var TemplateInitCMD = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: func(command *cobra.Command, args []string) {
		github, _ := command.Flags().GetString("github")
		if args[0] == "." {
			cmd.RunCommand("git", "clone", github, "temp-dir-for-clone")
			cmd.RunCommand("mv", "temp-dir-for-clone/*", ".")
			cmd.RunCommand("rmdir", "temp-dir-for-clone")
			cmd.RunCommand("rm", "-rf", args[0]+"/.git")
			cmd.RunCommand("git", "init", args[0])
		} else {
			cmd.RunCommand("git", "clone", github, args[0])
			cmd.RunCommand("rm", "-rf", args[0]+"/.git")
			cmd.RunCommand("git", "init", args[0])
		}
	},
}

func init() {
	TemplateInitCMD.Flags().StringP("github", "g", "https://github.com/frate-templates/frate-go.git", "Template to use as a base")
}
