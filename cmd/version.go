package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	fmt.Println(VERSION)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
