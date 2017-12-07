package main

import (
	"os"

	"fmt"

	"github.com/khoiln/sextant/cmd"
	"github.com/khoiln/sextant/pkg/database"
	"github.com/urfave/cli"
)

func main() {
	var db database.DB
	var err error
	if db, err = database.NewDefault(); err != nil {
		fmt.Println("Can't open database file")
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = "Sextant, navigate around your pirate 🛳"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:   "add",
			Usage:  "Add new entry",
			Action: cmd.CmdAdd(db),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "ls",
			Usage:  "List the directories sorted by their rank",
			Action: cmd.CmdLs(db),
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
			Action: cmd.CmdReset(db),
			Flags:  []cli.Flag{},
		},
	}

	app.CommandNotFound = CommandNotFound
	app.Run(os.Args)
}
