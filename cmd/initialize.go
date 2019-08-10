package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/packr"
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
}

func initialize() {
	if len(viper.GetString("title")) == 0 {
		tools.PrintRedln("Title is a required option! Please use '--title' switch")
		return
	}
	configuration.Author = viper.GetString("author")
	configuration.Title = viper.GetString("title")
	fmt.Printf("Creating new project with config file: %q\n", viper.GetString("config"))
	fmt.Printf("Author: %s\n", viper.GetString("author"))
	fmt.Printf("Title: %s\n", viper.GetString("title"))
	// basedir has to be empty AND exist
	basedir := viper.GetString("basedir")
	listing, err := ioutil.ReadDir(basedir)
	if err != nil {
		log.Fatal(err)
	}
	if len(listing) > 0 || !tools.PathExists(basedir) {
		tools.PrintRed("FAILED! ")
		fmt.Println("Project directory has to exist and has to be empty!")
		return
	}
	box := packr.NewBox("../template")
	templateListing := box.List()
	for idx, file := range templateListing {
		fmt.Println("Processing", file)
		var path, dir, content string
		var err error
		var tmpl *template.Template
		var outputh *os.File
		path = filepath.Join(basedir, file)
		dir = filepath.Dir(path)
		if !tools.PathExists(dir) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				panic(err)
			}
		}
		content, err = box.FindString(file)
		if err != nil {
			log.Fatal(err)
		}
		tmpl, err = template.New(fmt.Sprintf("%02d", idx+1)).Parse(content)
		if err != nil {
			log.Fatal(err)
		}
		outputh, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer outputh.Close()
		err = tmpl.Execute(outputh, configuration)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print("> ")
	tools.PrintGreenln("Ok")

}

func init() {
	rootCmd.AddCommand(initializeCmd)
	initializeCmd.Flags().StringP("author", "a", "Lydéric Landry", "set up author")
	viper.BindPFlag("author", initializeCmd.Flags().Lookup("author"))
	initializeCmd.Flags().StringP("title", "t", "", "(required) set up title")
	viper.BindPFlag("title", initializeCmd.Flags().Lookup("title"))
}
