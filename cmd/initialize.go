package cmd

import (
	"fmt"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Initialize a %s new project", PROGNAME),
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.persist()
	},
}

func initialize() {
	configuration.Author = viper.GetString("author")
	configuration.Title = viper.GetString("title")
	fmt.Printf("New project created with config file: %q\n", viper.GetString("config"))
	fmt.Printf("Author: %s\n", viper.GetString("author"))
	fmt.Printf("Title: %s\n", viper.GetString("title"))
	fmt.Printf("%ssing template\n", tools.Ternary(viper.GetBool("template"), "U", "Not u"))
}

func init() {
	rootCmd.AddCommand(initializeCmd)
	initializeCmd.Flags().StringP("author", "a", "Lydéric Landry", "set up author")
	viper.BindPFlag("author", initializeCmd.Flags().Lookup("author"))
	initializeCmd.Flags().StringP("title", "t", "", "(required) set up title")
	initializeCmd.MarkFlagRequired("title")
	viper.BindPFlag("title", initializeCmd.Flags().Lookup("title"))
	initializeCmd.Flags().BoolP("template", "T", false, "use project template")
	viper.BindPFlag("template", initializeCmd.Flags().Lookup("template"))
}
