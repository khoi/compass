package cmd

import (
	"os"

	"github.com/khoiracle/sextant/pkg/entry"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove non-existing folder from the database",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := fileDb.Read()

		if err != nil {
			exit(err)
		}

		var valid []*entry.Entry

		for _, e := range entries {
			if _, err := os.Stat(e.Path); err == nil {
				valid = append(valid, e)
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
