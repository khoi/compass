package cmd

import (
	"errors"
	"path/filepath"
	"sort"
	"time"

	"github.com/khoiln/sextant/pkg/entry"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a folder to the database",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
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

	sort.Sort(entry.ByPath(entries))
	idx := sort.Search(len(entries), func(i int) bool {
		return entries[i].Path >= path
	})

	if idx < len(entries) && entries[idx].Path == path { // Entry exists, update the score
		entries[idx].VisitedCount += 1
		entries[idx].LastVisited = int(time.Now().Unix())
	} else { // Create a new entry
		entries = append(entries, nil)
		copy(entries[idx+1:], entries[idx:])
		entries[idx] = &entry.Entry{
			Path:         path,
			VisitedCount: 1,
			LastVisited:  int(time.Now().Unix()),
		}
	}

	if err := fileDb.Write(entries); err != nil {
		exit(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
