package cmd

import (
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion script",
	RunE: func(cmd *cobra.Command, args []string) error {
		return completion()
	},
}

func completion() error {
	return rootCmd.GenBashCompletionFile(PROGNAME + ".completion")
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
