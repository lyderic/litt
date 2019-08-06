package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var totalsCmd = &cobra.Command{
	Use:   "totals",
	Short: "Count totals",
	Run: func(cmd *cobra.Command, args []string) {
		totals()
	},
}

func totals() {
	var err error
	bytes, chars, words := 0, 0, 0
	for _, file := range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
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
	fmt.Printf("%s Chars: %9.9s\n", BULLET, tools.ThousandSeparator(chars))
	fmt.Printf("%s Bytes: %9.9s\n", BULLET, tools.ThousandSeparator(bytes))
	fmt.Printf("%s Words: %9.9s\n", BULLET, tools.ThousandSeparator(words))
}

func init() {
	rootCmd.AddCommand(totalsCmd)
}
