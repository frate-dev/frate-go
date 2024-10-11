package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project using CMake",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building the project...")
	},
}

