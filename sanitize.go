package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lyderic/tools"
)

func sanitizeAllFiles() {
	fmt.Println("Sanitizing")
	var idx int
	var file string
	for idx, file = range configuration.Files {
		path := filepath.Join(basedir, file)
		sanitizePath(idx, path)
	}
	n := idx + 1
	fmt.Printf("%s %d file%s processed\n", bullet, n, tools.Ternary(n > 1, "s", ""))

}

func sanitizePath(idx int, path string) (err error) {
	start := time.Now()
	var nbytes int
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("cannot sanitize %s: file not found", path)
	}
	if nbytes, err = substitute(path,
		"---", "***",
		"--", "—",
		"...", "…",
		"« ", "« ",
		" »", " »",
		" ?", " ?",
		" !", " !",
		" ;", " ;",
		" :", " :",
		" \n", "\n",
		"  ", " "); err != nil {
		return err
	}
	if verbose {
		fmt.Printf("[%03d] %s: processed %d bytes in %s\n", idx+1, filepath.Base(path), nbytes, time.Since(start))
	}
	return
}

func substitute(file string, replacements ...string) (n int, err error) {
	r := strings.NewReplacer(replacements...)
	var bb []byte
	if bb, err = ioutil.ReadFile(file); err != nil {
		return n, err
	}
	var f *os.File
	if f, err = os.Create(file); err != nil {
		return n, err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	if n, err = r.WriteString(w, string(bb)); err != nil {
		return n, err
	}
	w.Flush()
	return
}
