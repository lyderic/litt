package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:                   "list",
	Aliases:               []string{"ls", "l"},
	DisableFlagsInUseLine: true,
	Short:                 "List montages",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func list() error {
	if len(configuration.Montages) == 0 {
		return fmt.Errorf("No montage found!")
	}
	// what is the longest name?
	ln := 0
	for _, montage := range configuration.Montages {
		i := len(montage.Name)
		if i > ln {
			ln = i
		}
	}
	format := fmt.Sprintf("  [%%02d] %%-%d.%ds [%%s]", ln, ln)
	for _, montage := range configuration.Montages {
		fmt.Printf(format, montage.Id, montage.Name, montage.Path)
		if montage.Id == 1 {
			fmt.Println(" (default)")
		} else {
			fmt.Println()
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
