package main

import (
	"log"

	"github.com/lyderic/litt/cmd"
	"github.com/lyderic/tools"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if !binariesOk() {
		return
	}
	cmd.Execute()
}

func binariesOk() (ok bool) {
	binaries := []string{"pandoc", "pdflatex"}
	if err := tools.CheckBinaries(binaries...); err != nil {
		tools.PrintRedf("These programs must be installed on your system for %s to run:\n", cmd.PROGNAME)
		for _, binary := range binaries {
			tools.PrintRedf(" %s %s\n", tools.BULLET, binary)
		}
		tools.PrintRedln(err)
		return false
	}
	return true
}
