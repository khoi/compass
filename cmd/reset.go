package cmd

import (
	"github.com/khoiln/sextant/pkg/database"
	"github.com/urfave/cli"
)

func CmdReset(c *cli.Context) error {
	db := c.App.Metadata["db"].(database.DB)
	return db.Truncate()
}
