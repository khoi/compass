package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/khoiln/sextant/pkg/entry"
	"github.com/urfave/cli"
)

func CmdCd(c *cli.Context) error {
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

	if len(filtered) > 0 {
		fmt.Fprintf(c.App.Writer, "%s\n", filtered[len(filtered)-1].Path)
	}

	return nil
}
