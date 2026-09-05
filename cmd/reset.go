package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

Var resetCmd = &cobra.Command{
	Use : "Reset",
	Args: cobra.ExactArgs(1)

	Run: func(cmd *cobra.Command, args []string){
		Verify := args[0]

	},
}
