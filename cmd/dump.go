package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump configuration (for debugging)",
	Run: func(cmd *cobra.Command, args []string) {
		dump()
	},
}

func dump() {
	fmt.Println("*** VIPER ***")
	fmt.Printf("%#v\n", viper.AllSettings())
	fmt.Println("*** GLOBALS VARIABLES ***")
	fmt.Printf("cfgFile: %#v\n", cfgFile)
	fmt.Printf("%#v\n", configuration)
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
