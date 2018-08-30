package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge the database.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s purged.\n", cfgFile)
		fileDb.Truncate()
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
