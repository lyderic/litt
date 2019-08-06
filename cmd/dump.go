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
	fmt.Println("*** CONFIGURATION ***")
	fmt.Printf("%#v\n", configuration)
	fmt.Println("*** GLOBALS VARIABLES ***")
	fmt.Println("cfgFile:", cfgFile)
	fmt.Println("jsonPath:", jsonPath)
	fmt.Println("basedir:", basedir)
	fmt.Println("reference:", reference)
	fmt.Println("tag:", tag)
	fmt.Println("nocontent:", nocontent)
	fmt.Println("verbose:", verbose)
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
