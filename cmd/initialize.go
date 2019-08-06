package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project, creating a minimal litt.json",
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func initialize() {
	var c Configuration
	c.Author = "Lydéric Landry"
	c.Title = "Unset Title"
	c.Montages = []Montage{
		Montage{
			Id:   1,
			Name: "kdp",
			Path: "meta/montage_kdp.tex",
		},
	}
	c.Files = []string{"chapitre01.lkl"}
	c.Replacements = []Replacement{
		Replacement{
			From: "\\section{",
			To:   "\\chapter*{\\centering ",
		},
	}
	c.Double = false
	fmt.Printf("%#v\n", c)
}

func init() {
	rootCmd.AddCommand(initializeCmd)
}
