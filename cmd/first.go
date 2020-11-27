package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// firstCmd represents the first command
var firstCmd = &cobra.Command{
	Use:                   "first",
	Aliases:               []string{"1"},
	DisableFlagsInUseLine: true,
	Short:                 "show the first line of the files to assemble",
	Long: `This command is useful to display the first line of
chapters that usually contains their title.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		configuration.load()
	},
	Run: func(cmd *cobra.Command, args []string) {
		first()
	},
}

func first() {
	for _, file := range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
		fh, err := os.Open(path)
		if err != nil {
			tools.PrintRedf("failed to open file %q\n", path)
			return
		}
		defer fh.Close()
		scanner := bufio.NewScanner(fh)
		scanner.Scan()
		fmt.Printf("%s: %q\n", file, scanner.Text())
		if err := scanner.Err(); err != nil {
			tools.PrintRedf("failed to scan file %q\n", path)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(firstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// firstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// firstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
