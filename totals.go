package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/lyderic/tools"
)

func totals() {
	var err error
	bytes, chars, words := 0, 0, 0
	for _, file := range configuration.Files {
		path := filepath.Join(basedir, file)
		var bb []byte
		if bb, err = ioutil.ReadFile(path); err != nil {
			return
		}
		if !utf8.Valid(bb) {
			log.Fatalf("%s: not a valid UTF8 file", path)
		}
		content := string(bb)
		bytes = bytes + len(bb)
		chars = chars + utf8.RuneCount(bb)
		words = words + len(strings.Fields(content))
	}
	fmt.Println("Count Totals:")
	fmt.Printf("%s Chars: %9.9s\n", bullet, tools.ThousandSeparator(chars))
	fmt.Printf("%s Bytes: %9.9s\n", bullet, tools.ThousandSeparator(bytes))
	fmt.Printf("%s Words: %9.9s\n", bullet, tools.ThousandSeparator(words))
}
