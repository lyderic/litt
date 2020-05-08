package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func sanitize(path string, verbose bool) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("cannot sanitize %s: file not found", path)
	}
	return substitute(path, verbose,
		"---", "* * *",
		"***", "* * *",
		"--", "—",
		"...", "…",
		"« ", "« ",
		" »", " »",
		" ?", " ?",
		" !", " !",
		" ;", " ;",
		" :", " :",
		" \n", "\n",
		"  ", " ")
}

func substitute(path string, verbose bool, replacements ...string) (err error) {
	start := time.Now()
	replacer := strings.NewReplacer(replacements...)
	// get content of the file
	var buffer []byte
	if buffer, err = ioutil.ReadFile(path); err != nil {
		return
	}
	content := string(buffer)
	// Apply the replacements
	replaced := replacer.Replace(content)
	// If content and replaced differ, write file and notify (if verbose)
	if replaced != content {
		ioutil.WriteFile(path, []byte(replaced), 0644)
		if verbose {
			fmt.Printf("%s: processed in %s\n", filepath.Base(path), time.Since(start))
		}
	}
	return
}
