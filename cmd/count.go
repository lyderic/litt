package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/lyderic/tools"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var countCmd = &cobra.Command{
	Use:     "count",
	Aliases: []string{"n"},
	Short:   "Count characters, bytes and words",
	Run: func(cmd *cobra.Command, args []string) {
		count()
	},
}

type Item struct {
	Name  string
	Bytes int
	Chars int
	Words int
}

func count() {
	var err error
	var items []Item
	for _, file := range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
		item := Item{}
		item.Name = filepath.Base(path)
		var bb []byte
		if bb, err = ioutil.ReadFile(path); err != nil {
			return
		}
		if !utf8.Valid(bb) {
			log.Fatalf("%s: not a valid UTF8 file", item.Name)
		}
		content := string(bb)
		item.Bytes = len(bb)
		item.Chars = utf8.RuneCount(bb)
		item.Words = len(strings.Fields(content))
		items = append(items, item)
	}
	if err = display(items); err != nil {
		log.Fatal(err)
	}
}

func display(items []Item) (err error) {
	buffer := new(bytes.Buffer)
	lines := [][]string{}
	for _, data := range items {
		n := data.Name
		c := tools.ThousandSeparator(data.Chars)
		w := tools.ThousandSeparator(data.Words)
		b := tools.ThousandSeparator(data.Bytes)
		lines = append(lines, []string{n, c, w, b})
	}
	table := tablewriter.NewWriter(buffer)
	table.SetHeader([]string{"Name", "Chars", "Words", "Bytes"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT})
	table.AppendBulk(lines)
	totals := getTotals(items)
	table.SetFooter([]string{
		totals.Name,
		tools.ThousandSeparator(totals.Chars),
		tools.ThousandSeparator(totals.Words),
		tools.ThousandSeparator(totals.Bytes)})
	table.Render()
	tools.Less(buffer.String())
	return
}

func getTotals(items []Item) Item {
	totals := Item{}
	var c, w, b int
	for _, data := range items {
		c = c + data.Chars
		w = w + data.Words
		b = b + data.Bytes
	}
	totals.Name = "TOTALS"
	totals.Chars = c
	totals.Words = w
	totals.Bytes = b
	return totals
}

func init() {
	rootCmd.AddCommand(countCmd)
}
