package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var assembleCmd = &cobra.Command{
	Use:     "assemble",
	Aliases: []string{"a", "ass"},
	Short:   "Assemble montage",
	Long: `
Assemble montage.

If no montage is specified, the default montage ("1") is used.
If --no-content is given, "content.tex" is not regenerated, unless it doesn't exist.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		configuration.load()
	},
	Run: func(cmd *cobra.Command, args []string) {
		assemble()
	},
}

func assemble() {
	var montage Montage
	var err error
	if montage, err = getSelectedMontage(); err != nil {
		tools.PrintRedf("Montage not found: %v\n", err)
		return
	}
	fmt.Printf("Assembling montage %q\n", montage.Name)
	sanitizeAllFiles()
	buildContent(montage)
	buildpdf(montage)
}

func buildContent(montage Montage) {
	montageDir := getMontageDir(montage)
	contentFile := filepath.Join(montageDir, "content.tex")
	contentExists := true
	if _, err := os.Stat(contentFile); os.IsNotExist(err) {
		contentExists = false
	}
	if viper.GetBool("noc") && contentExists {
		tools.PrintYellowln("content.tex found, not regenerated as requested.")
		return
	}
	if viper.GetBool("noc") && !contentExists {
		tools.PrintYellowln("Not regenerated content as been requested, however content.tex was not found...")
	}
	fmt.Println("Building content")
	fmt.Println(tools.PROMPT, "creating content.tex with pandoc")
	var args []string
	args = append(args, "--from", "markdown-auto_identifiers-space_in_atx_header", "--to", "latex")
	args = append(args, "-o")
	args = append(args, contentFile)
	for _, file := range configuration.Files {
		args = append(args, filepath.Join(viper.GetString("basedir"), file))
	}
	cmd := exec.Command("pandoc", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if viper.GetBool("debug") {
		tools.PrintYellowln(cmd.Args)
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	for _, replacement := range configuration.Replacements {
		fmt.Printf("%s replacing %q -> %q\n", tools.PROMPT, replacement.From, replacement.To)
		input, err := ioutil.ReadFile(contentFile)
		if err != nil {
			log.Fatal(err)
		}
		output := bytes.Replace(input, []byte(replacement.From), []byte(replacement.To), -1)
		if err = ioutil.WriteFile(contentFile, output, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func buildpdf(montage Montage) {
	fmt.Println("Creating PDF")
	montageDir := getMontageDir(montage)
	montageTexName := filepath.Base(montage.Path)
	montagePdfName := strings.Replace(montageTexName, ".tex", ".pdf", 1)
	os.Chdir(montageDir)
	pdflatex(montageTexName)
	if configuration.Double {
		pdflatex(montageTexName)
	}
	pdfName := fmt.Sprintf("%s - %s.pdf", configuration.Title, configuration.Author)
	if viper.GetBool("tag") {
		pdfName = fmt.Sprintf("%s - %s [%s-%s].pdf", configuration.Title, configuration.Author, montage.Name, time.Now().Format("20060102150405"))
	}
	pdfFullPath := filepath.Join(viper.GetString("basedir"), pdfName)
	os.Rename(montagePdfName, pdfFullPath)
	fmt.Printf("%s created %q\n", tools.PROMPT, pdfFullPath)
}

func pdflatex(tex string) {
	fmt.Printf("%s running pdflatex on %q\n", tools.PROMPT, tex)
	cmd := exec.Command("pdflatex", "-halt-on-error", "-shell-escape", tex)
	if viper.GetBool("verbose") {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(assembleCmd)
	assembleCmd.Flags().StringP("montage", "m", "1", "use this montage")
	viper.BindPFlag("reference", assembleCmd.Flags().Lookup("montage"))
	assembleCmd.Flags().BoolP("no-content", "n", false, "don't regenerate content if content.tex already exists")
	viper.BindPFlag("noc", assembleCmd.Flags().Lookup("no-content"))
	assembleCmd.Flags().BoolP("tag", "t", false, "tag final PDF with montage name and timestamp")
	viper.BindPFlag("tag", assembleCmd.Flags().Lookup("tag"))
	assembleCmd.Flags().BoolP("verbose", "v", false, "be verbose: show pdflatex output")
	viper.BindPFlag("verbose", assembleCmd.Flags().Lookup("verbose"))
}
