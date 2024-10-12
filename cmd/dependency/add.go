package dependency

import (
	"frate-go/config"
	"frate-go/dependency"
	"log"

	"github.com/spf13/cobra"
)

var DepAddCmd = &cobra.Command{
	Use:   "install",
  Aliases: []string{"i", "install"},
  Args: func(cmd *cobra.Command, args []string) error {
    if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
        return err
    }

    return nil 
  },
	Short: "Build the project using CMake",
  Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		github, _ := cmd.Flags().GetString("github")
    if err != nil {
      log.Fatal(err)
    }
    dep := dependency.Dependency{
      Name: args[0],
    } 
    if github != ""{
      dep.GitURL = github
    }
    cfg.Dependencies = append(cfg.Dependencies, dep)
		config.GenerateConfig(cfg)
	},
}

func init() {
	DepAddCmd.Flags().StringP("github", "g", "", "github url")
}


