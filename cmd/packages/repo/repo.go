package package_repo

import (
	"github.com/spf13/cobra"
)

var PackageRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Add, Remove, or List package repositories",
}

func init(){
	PackageRepoCmd.AddCommand(PackageRepoAddCmd)
	PackageRepoCmd.AddCommand(PackageRepoListCmd)
	PackageRepoCmd.AddCommand(PackageRepoRemoveCmd)
}
