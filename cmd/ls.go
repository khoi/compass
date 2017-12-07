package cmd

import (
	"sort"

	"fmt"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/khoiln/sextant/pkg/entry"
	"github.com/khoiln/sextant/pkg/fuzzy"
	"github.com/urfave/cli"
)

func CmdLs(db database.DB) func(*cli.Context) error {
	return func(c *cli.Context) error {
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
				if fuzzy.MatchFold(query, e.Path) {
					filtered = append(filtered, e)
				}
			}
		}

		sort.Sort(entry.ByRank(filtered))

		for i := len(filtered) - 1; i >= 0; i -= 1 {
			fmt.Fprintf(c.App.Writer, "%s\n", filtered[i].Path)
		}

		return nil
	}
}
