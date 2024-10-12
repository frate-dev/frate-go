package main

import (
	"github.com/spf13/cobra"
	"frate-go/cmd"
  "frate-go/cmd/dependency"
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
	rootCmd.AddCommand(dependency.DependencyCmd)
}

