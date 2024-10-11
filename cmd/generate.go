package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate/regenerate your project's using CMake config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating the cmake...")
	},
}
