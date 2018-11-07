package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/lyderic/tools"
)

var extensionsToClean = []string{".aux", ".log", ".out", ".html", ".bbl", ".blg", ".css", ".dvi", ".idv", ".lg", ".tmp", ".toc", ".xref", ".4ct", ".4tc", ".rtf", ".pdf"}

func clean() {
	montage := getSelectedMontage()
	fmt.Printf("Cleaning montage %q\n", montage.Name)
	a := cleanDir(getMontageDir(montage))
	b := cleanDir(basedir)
	n := a + b
	fmt.Printf("%d file%s removed\n", n, tools.Ternary(n > 1, "s", ""))
}

func cleanDir(dir string) (n int) {
	fmt.Println("[directory]", dir)
	var listing []os.FileInfo
	var err error
	if listing, err = ioutil.ReadDir(dir); err != nil {
		log.Fatal(err)
	}
	for _, fifo := range listing {
		name := fifo.Name()
		path := filepath.Join(dir, name)
		for _, extension := range extensionsToClean {
			if strings.HasSuffix(name, extension) {
				fmt.Println("⮞", path)
				os.Remove(path)
				n++
			}
		}
	}
	return
}
