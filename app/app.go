package app

import (
	"github.com/khoiln/sextant/cmd"
	"github.com/khoiln/sextant/pkg/database"
	"github.com/urfave/cli"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = "Sextant, navigate around your pirate 🛳"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Usage: "Path to the db file (Default: ~/.sextant)",
		},
	}

	app.Before = func(c *cli.Context) error {
		var db database.DB
		var err error
		dbPath := c.String("db")
		if dbPath == "" {
			db, err = database.NewDefault()
		} else {
			db, err = database.New(dbPath)
		}
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		c.App.Metadata["db"] = db
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:   "add",
			Usage:  "Add new entry",
			Action: cmd.CmdAdd,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "ls",
			Usage:  "List the directories sorted by their rank",
			Action: cmd.CmdLs,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "shell",
			Usage:  "Prints out the shell integration scripts.",
			Action: cmd.CmdShell,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type",
					Usage: "type of your shell (bash | zsh)",
					Value: "sh",
				},
			},
		},
		{
			Name:   "reset",
			Usage:  "When you need a new beginning.",
			Action: cmd.CmdReset,
			Flags:  []cli.Flag{},
		},
	}

	app.CommandNotFound = CommandNotFound
	return app
}
