package main

import (
	"os"

	"github.com/khoiln/sextant/app"
)

func main() {
	a := app.NewApp()
	a.Run(os.Args)
}
