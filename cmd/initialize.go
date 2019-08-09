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
	configuration.Montages = []Montage{
		Montage{
			Id:   1,
			Name: "kdp",
			Path: "meta/montage_kdp.tex",
		},
	}
	configuration.Files = []string{"chapitre01.lkl"}
	configuration.Replacements = []Replacement{
		Replacement{
			From: "\\section{",
			To:   "\\chapter*{\\centering ",
		},
	}
	configuration.Double = false
	fmt.Printf("%#v\n", configuration)
}

func init() {
	rootCmd.AddCommand(initializeCmd)
}
