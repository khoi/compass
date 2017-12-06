package command

import (
	"github.com/khoiln/sextant/database"
	"github.com/urfave/cli"
)

func CmdReset(db database.DB) func(*cli.Context) error {
	return func(c *cli.Context) error {
		return db.Truncate()
	}
}
