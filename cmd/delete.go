package cmd
import(
	"fmt"
	"github.com/spf13/cobra"
)


Var deleteCmd = &cobra.Command{
	Use : "Delete",
	Args: cobra.ExactArgs(1)

	Run : func(cmd *cobra.Command, args []string){
		Name := args[0]
	}
}
