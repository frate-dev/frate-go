package index

import (
	"fmt"

	"github.com/spf13/cobra"
)

var IndexUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Manage local packages",
  Run: func(cmd *cobra.Command, args []string){
    fmt.Println("updated packaged index")
  },
}

func init(){

}
