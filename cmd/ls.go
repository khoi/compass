package cmd

import (
	"sort"
	"strings"

	"fmt"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/khoiln/sextant/pkg/entry"
	"github.com/khoiln/sextant/pkg/path"
	"github.com/urfave/cli"
)

func CmdLs(c *cli.Context) error {
	db := c.App.Metadata["db"].(database.DB)
	query := c.Args().First()
	entries, err := db.Read()

	if err != nil {
		return err
	}

	var filtered []*entry.Entry

	if query == "" {
		filtered = entries
	} else {
		for _, e := range entries {
			if strings.Contains(strings.ToLower(e.Path), query) {
				filtered = append(filtered, e)
			}
		}
	}

	sort.Sort(entry.ByFerecency(filtered))

	var paths []string
	for _, e := range filtered {
		paths = append(paths, e.Path)
	}

	if common := path.LCP(paths); common != "" {
		fmt.Fprintf(c.App.Writer, "common \t %s\n", common)
	}

	for _, e := range filtered {
		fmt.Fprintf(c.App.Writer, "%d \t %s\n", entry.Frecency(e), e.Path)
	}

	return nil
}
