package cmd

import (
	"log"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		edit()
	},
}

func edit() {
	config := viper.GetString("config")
	// The configuration file is *edited* not *created*
	// so we check that it exists
	if !tools.PathExists(config) {
		tools.PrintRedf("%q: not found. Please run 'init' first\n", config)
		return
	}
	err := tools.Vim(config)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(editCmd)
}
