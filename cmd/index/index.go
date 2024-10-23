package index

import (
	"github.com/spf13/cobra"
)

var IndexCmd = &cobra.Command{
	Use:   "index",
	Short: "Manage local packages",
}

func init(){
  IndexCmd.AddCommand(IndexUpdateCmd)
}
