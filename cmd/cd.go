package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/khoiln/sextant/pkg/entry"
	"github.com/khoiln/sextant/pkg/path"
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
			if strings.Contains(e.Path, query) {
				filtered = append(filtered, e)
			}
		}
	}

	sort.Sort(entry.ByFerecency(filtered))

	var paths []string

	for _, e := range filtered {
		paths = append(paths, e.Path)
	}

	if len(paths) == 0 {
		return cli.NewExitError("", 1)
	}

	output := paths[len(paths)-1]
	if common := path.LCP(paths); common != "" {
		for _, p := range paths {
			if p == common {
				output = common
				break
			}
		}
	}

	fmt.Fprintf(c.App.Writer, "%s\n", output)
	return nil
}


