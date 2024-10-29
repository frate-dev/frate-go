package packages

import (
	"github.com/spf13/cobra"
)

var PackageCMD = &cobra.Command{
	Use:     "package",
	Short:    "Search and Push your project to package repo", 
	Aliases: []string{"t"},
}

func init() {
	PackageCMD.AddCommand(PackageListCmd)
}
