package cmd

import (
 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
 Use:   "EPROM",
 Short: "Tui For passmanager",
 Long:  "Passmanager using EEPROM circuit.",
}

func Execute() {
 cobra.CheckErr(rootCmd.Execute())
}

func init() {
 rootCmd.AddCommand(addCmd)
 rootCmd.AddCommand(readCmd)
 rootCmd.AddCommand(resetCmd)
 rootCmd.AddCommand(deleteCmd)
}
