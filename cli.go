package main

import (
	"frate-go/cmd"
	"frate-go/cmd/dependency"
	"frate-go/cmd/packages"
	package_repo "frate-go/cmd/packages/repo"
	"frate-go/cmd/template"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frate-go",
	Short: "Frate-go CLI for C/C++ project management",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.BuildCmd)
	rootCmd.AddCommand(cmd.GenerateCmd)
	rootCmd.AddCommand(cmd.RunCmd)
	rootCmd.AddCommand(cmd.WatchCmd) 
	rootCmd.AddCommand(dependency.DependencyCmd)
  rootCmd.AddCommand(template.TemplateCMD)
	rootCmd.AddCommand(packages.PackageCMD) 
	rootCmd.AddCommand(package_repo.PackageRepoCmd)  
}
