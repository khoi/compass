package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khoiln/sextant/pkg/entry"
	"github.com/khoiln/sextant/pkg/path"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the directories along with their ranking",
	Run:   lsRun,
}

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

	for _, e := range entries {
		if strings.Contains(e.Path, query) {
			filtered = append(filtered, e)
			filteredPaths = append(filteredPaths, e.Path)
		}
	}

	sort.Sort(entry.ByFerecency(filtered))

	if common := path.LCP(filteredPaths); common != "" {
		fmt.Printf("common \t %s\n", common)
	}

	for _, e := range filtered {
		fmt.Printf("%d \t %s\n", entry.Frecency(e), e.Path)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
