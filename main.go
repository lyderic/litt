package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	version = "0.0.7"
	bullet  = "⮞"
)

var (
	jsonPath      string
	basedir       string
	configuration Configuration
	reference     string // name or id of used montage
	nosanitize    bool
	tag           bool
	nocontent     bool
	verbose       bool
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	flag.StringVar(&jsonPath, "f", "./litt.json", "configuration `file`")
	flag.StringVar(&reference, "m", "1", "create `montage`")
	flag.BoolVar(&nocontent, "b", false, "don't build content")
	flag.BoolVar(&nosanitize, "s", false, "don't sanitize files before assembling")
	flag.BoolVar(&tag, "t", false, "tag final PDF with montage name and timestamp")
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.Usage = usage
	flag.Parse()
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		usage()
		log.Fatal(err)
	}
	basedir = getParent(jsonPath)
	configuration.load()
	if len(flag.Args()) == 0 {
		usage()
		fmt.Println("Please provide an action!")
		return
	}
	todo := flag.Args()[0]
	found := false
	var action Action
	for _, current := range actions {
		if current.Name == todo || current.Letter == todo {
			found = true
			action = current
		}
	}
	if !found {
		fmt.Printf("%s: invalid action.\n", todo)
		usage()
		return
	}
	action.Function()
}

func usage() {
	fmt.Printf("\nlitt v.%s - (c) Lyderic Landry, London 2018\n", version)
	fmt.Println("Usage: litt [<option>] <action>")
	fmt.Println("\n Actions:\n")
	for _, action := range actions {
		fmt.Println(action)
	}
	fmt.Println("\n Options:\n")
	flag.PrintDefaults()
	if len(configuration.Montages) > 0 {
		fmt.Println("\n Montages:\n")
		list()
	}
	fmt.Println()
}
