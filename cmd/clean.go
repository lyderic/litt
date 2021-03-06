package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cleanCmd = &cobra.Command{
	Use:                   "clean",
	Aliases:               []string{"c", "clear"},
	DisableFlagsInUseLine: true,
	Short:                 "Clean configuration",
	PreRun: func(cmd *cobra.Command, args []string) {
		configuration.load()
	},
	Run: func(cmd *cobra.Command, args []string) {
		clean()
	},
}

var extensionsToClean = []string{".aux", ".log", ".out", ".html",
	".bbl", ".blg", ".dvi", ".idv", ".lg", ".tmp", ".toc", ".xref",
	".4ct", ".4tc", ".rtf", ".pdf", ".epub"}

func clean() {
	cleanMontages()
	//cleanAll()
}

func cleanAll() {
	err := filepath.Walk(viper.GetString("basedir"), func(path string, finfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if finfo.IsDir() {
			return nil
		}
		fmt.Println(">>>>>", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func cleanMontages() {
	for _, montage := range configuration.Montages {
		fmt.Printf("Cleaning montage %q\n", montage.Name)
		a := cleanDir(getMontageDir(montage))
		b := cleanDir(viper.GetString("basedir"))
		n := a + b
		fmt.Printf("%d file%s removed\n", n, tools.Ternary(n > 1, "s", ""))
	}
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
		if filepath.Base(path) == "content.tex" {
			fmt.Println(tools.PROMPT, path)
			os.Remove(path)
			n++
		}
		for _, extension := range extensionsToClean {
			if strings.HasSuffix(name, extension) {
				fmt.Println(tools.PROMPT, path)
				os.Remove(path)
				n++
			}
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
