package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Initialize a new project", PROGNAME),
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.persist()
	},
}

func initialize() {
	configuration.Author = "Lydéric Landry"
	configuration.Title = "Unset Title"
}

func init() {
	rootCmd.AddCommand(initializeCmd)
}
