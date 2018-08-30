package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/khoi/compass/pkg/entry"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a  folder from the database",
	Run:   removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		exit(errors.New("Missing folder."))
	}

	path, err := filepath.Abs(args[0])

	if err != nil {
		exit(err)
	}

	entries, err := fileDb.Read()

	if err != nil {
		exit(err)
	}

	newEntries := entry.Entries(entries).Filter(func(e *entry.Entry) bool {
		return e.Path != path
	})

	if err := fileDb.Write(newEntries); err != nil {
		exit(err)
	}

	fmt.Printf("%s removed.\n", path)
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
