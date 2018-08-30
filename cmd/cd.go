package cmd

import (
	"fmt"
	"sort"
	"strings"

	"errors"

	"github.com/khoi/compass/pkg/entry"
	"github.com/khoi/compass/pkg/path"
	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Print the top match for search terms",
	Run:   cdRun,
}

func cdRun(cmd *cobra.Command, args []string) {
	var query string

	if len(args) > 0 {
		query = args[0]
	}

	entries, err := fileDb.Read()

	if err != nil {
		exit(err)
	}

	var filtered []*entry.Entry
	var filteredPaths []string

	matchedEntries := entry.Entries(entries).Filter(func(e *entry.Entry) bool {
		return strings.Contains(e.Path, query) || strings.Contains(strings.ToLower(e.Path), strings.ToLower(query))
	})

	for _, e := range matchedEntries {
		filtered = append(filtered, e)
		filteredPaths = append(filteredPaths, e.Path)
	}

	sort.Sort(entry.ByFrecency(filtered))

	for _, e := range filtered {
		filteredPaths = append(filteredPaths, e.Path)
	}

	if len(filteredPaths) == 0 {
		exit(errors.New(""))
	}

	output := filteredPaths[len(filteredPaths)-1]
	if common := path.LCP(filteredPaths); common != "" {
		for _, p := range filteredPaths {
			if p == common {
				output = common
				break
			}
		}
	}

	fmt.Printf("%s\n", output)
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
