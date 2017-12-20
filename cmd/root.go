package cmd

import (
	"fmt"
	"os"

	"os/user"
	"path/filepath"

	"github.com/khoiln/sextant/pkg/database"
	"github.com/spf13/cobra"
)

var defaultConfigFileName = ".sextant"
var cfgFile string
var fileDb database.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sextant",
	Short: "Sextant, navigate around your pirate 🛳",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exit(err)
	}
}

func init() {
	cobra.OnInitialize(initDB)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "path to the db file (default is $HOME/.sextant)")
}

func initDB() {
	var err error

	if cfgFile == "" {
		usr, err := user.Current()
		if err != nil {
			exit(err)
		}
		cfgFile = filepath.Join(usr.HomeDir, defaultConfigFileName)
	}

	if fileDb, err = database.New(cfgFile); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
