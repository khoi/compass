package main

import (
	"os"

	"github.com/khoiln/sextant/app"
)

func main() {
	a := app.NewApp(Version)
	a.Run(os.Args)
}
