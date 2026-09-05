package cmd
import(
	"fmt"
	 "github.com/spf13/cobra"

)
var addCmd = &cobra.Command{
	Use : "add",

	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string){
		Name := args[0]
		Login := args[1]
		Password := args[2]
		fmt.Println(Name)
		fmt.Println(Login)
		fmt.Println(Password)
	},
}
