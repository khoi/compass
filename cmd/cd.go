package cmd

import (
	"fmt"
	"sort"
	"strings"

	"errors"

	"github.com/khoiracle/sextant/pkg/entry"
	"github.com/khoiracle/sextant/pkg/path"
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

	for _, e := range entries {
		if strings.Contains(e.Path, query) {
			filtered = append(filtered, e)
		}
	}

	sort.Sort(entry.ByFerecency(filtered))

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
