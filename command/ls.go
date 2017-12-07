package command

import (
	"sort"

	"fmt"

	"github.com/khoiln/sextant/database"
	"github.com/khoiln/sextant/search"
	"github.com/urfave/cli"
)

func CmdLs(db database.DB) func(*cli.Context) error {
	return func(c *cli.Context) error {
		query := c.Args().First()
		entries, err := db.Read()

		if err != nil {
			return err
		}

		var filtered []*search.Entry

		if query == "" {
			filtered = entries
		} else {
			for _, e := range entries {
				if search.MatchFold(query, e.Path) {
					filtered = append(filtered, e)
				}
			}
		}

		sort.Sort(search.ByRank(filtered))

		for i := len(filtered) - 1; i >= 0; i -= 1 {
			fmt.Fprintf(c.App.Writer, "%s\n", filtered[i].Path)
		}

		return nil
	}
}
