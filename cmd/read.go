package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

Var readCmd = &cobra.Command{
	Use: "Read",

	Args: cobra.ExactArgs(1),
	Run : func(cmd *cobra.Command, args []string){
		Name := args[0]


	},
}
