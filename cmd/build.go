package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// BuildCmd represents the 'build' command for frate-go
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project using CMake",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building the project...")
		// Add logic to invoke CMake and perform the build
	},
}

