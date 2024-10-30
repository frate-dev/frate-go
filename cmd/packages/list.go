package packages

import (
	"encoding/json"
	"fmt"
	"frate-go/cmd"
	"log"

	"github.com/spf13/cobra"
)


var PackageListCmd = &cobra.Command{
	Use:   "list",
	Aliases:  []string{"ls"},
	Short: "List available packages",
	Run: func(command *cobra.Command, args []string) {
		url:= "http://localhost:8000/packages" 
		data, err := cmd.Get(url) 
		var prettyData interface{} 
		json.Unmarshal([]byte(data), &prettyData) 	
		if err != nil {
			log.Fatal("Error getting templates: \n\t", err)
		}
		prettyJson, err:=json.MarshalIndent(prettyData, "", "  ")
		fmt.Println( string(prettyJson))
	},
}


func init(){

}
