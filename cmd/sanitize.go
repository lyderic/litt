package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/lyderic/tools"
)

func sanitizeAllFiles() {
	fmt.Println("Sanitizing")
	var idx int
	var file string
	for idx, file = range configuration.Files {
		path := filepath.Join(basedir, file)
		tools.Sanitize(path, true)
	}
	n := idx + 1
	fmt.Printf("%s %d file%s processed\n", BULLET, n, tools.Ternary(n > 1, "s", ""))
}
