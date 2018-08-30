package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khoi/compass/pkg/entry"
	"github.com/khoi/compass/pkg/path"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the directories along with their ranking",
	Run:   lsRun,
}

var pathOnly bool

func lsRun(cmd *cobra.Command, args []string) {
	var query string
	var entries []*entry.Entry
	var filtered []*entry.Entry
	var filteredPaths []string
	var err error

	if len(args) > 0 {
		query = args[0]
	}

	if entries, err = fileDb.Read(); err != nil {
		exit(err)
	}

	matchedEntries := entry.Entries(entries).Filter(func(e *entry.Entry) bool {
		return strings.Contains(e.Path, query) || strings.Contains(strings.ToLower(e.Path), strings.ToLower(query))
	})

	for _, e := range matchedEntries {
		filtered = append(filtered, e)
		filteredPaths = append(filteredPaths, e.Path)
	}

	sort.Sort(entry.ByFrecency(filtered))

	if common := path.LCP(filteredPaths); common != "" && !pathOnly {
		fmt.Printf("common \t %s\n", common)
	}

	for _, e := range filtered {
		if pathOnly {
			fmt.Printf("%s\n", e.Path)
		} else {
			fmt.Printf("%d \t %s\n", entry.Frecency(e), e.Path)
		}
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVar(&pathOnly, "path-only", false, "Prints out only path without score")
}
