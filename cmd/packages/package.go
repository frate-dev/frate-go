package packages

import (
	"github.com/spf13/cobra"
	"frate-go/cmd/packages/repo" 
)

var PackageCMD = &cobra.Command{
	Use:     "packages",
	Short:    "Search and Push your project to package repo", 
	Aliases: []string{"package"},
}

func init() {
	PackageCMD.AddCommand(PackageListCmd)
	PackageCMD.AddCommand(PackagePushCmd)
	PackageCMD.AddCommand(PackageSearchCmd)
	PackageCMD.AddCommand(package_repo.PackageRepoCmd) 
}
