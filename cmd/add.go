package cmd

import (
	"path/filepath"

	"sort"

	"time"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/khoiln/sextant/pkg/entry"
	"github.com/urfave/cli"
)

func CmdAdd(db database.DB) func(*cli.Context) error {
	return func(c *cli.Context) error {
		path, err := filepath.Abs(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		entries, err := db.Read()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
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

		if err := db.Write(entries); err != nil {
			return err
		}

		return nil
	}
}
