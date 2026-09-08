package cmd
import(
	"fmt"
	"github.com/spf13/cobra"
)


var deleteCmd = &cobra.Command{
	Use : "delete",
	Args: cobra.ExactArgs(1),

	Run : func(cmd *cobra.Command, args []string){
		Name := args[0]
		fmt.Println(Name)
	},
}
