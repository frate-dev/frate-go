package packages 


import (
	"encoding/json"
	"fmt"
	"frate-go/cmd"
	"log"

	"github.com/spf13/cobra"
)


var PackageSearchCmd = &cobra.Command{
	Use:   "search",
		Aliases: []string{"s", "search"}, 
	Short: "Search available packages",
	Args:  cobra.MinimumNArgs(1), 
	Run: func(command *cobra.Command, args []string) {
		url:= "http://localhost:8000/package?name=" 
		url += args[0] 
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
