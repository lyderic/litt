package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func sanitizeAllFiles() {
	fmt.Println("Sanitizing")
	var idx int
	var file string
	for idx, file = range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
		tools.Sanitize(path, true)
	}
	n := idx + 1
	fmt.Printf("%s %d file%s processed\n", tools.PROMPT, n, tools.Ternary(n > 1, "s", ""))
}
