package cmd

import (
	"fmt"
	"os"

	"github.com/khoi/compass/pkg/entry"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove non-existing folder from the database",
	Run: func(cmd *cobra.Command, args []string) {
		if (verbose) {
			fmt.Println("compass cleaning up.")
		}
		
		entries, err := fileDb.Read()

		if err != nil {
			exit(err)
		}

		var valid []*entry.Entry

		for _, e := range entries {
			if _, err := os.Stat(e.Path); err == nil {
				valid = append(valid, e)
				continue
			}
			if (verbose) {
				fmt.Fprintf(os.Stdout, "Removed %s\n", e.Path)
			}
		}

		if err := fileDb.Write(valid); err != nil {
			exit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
