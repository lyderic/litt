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

	"github.com/spf13/cobra"
)

var assembleCmd = &cobra.Command{
	Use:   "assemble",
	Short: "Assemble montage",
	Run: func(cmd *cobra.Command, args []string) {
		assemble()
	},
}

func assemble() {
	montage := getSelectedMontage()
	fmt.Printf("Assembling montage %q\n", montage.Name)
	sanitizeAllFiles()
	if !nocontent {
		buildContent(montage)
	}
	buildpdf(montage)
}

func buildContent(montage Montage) {
	fmt.Println("Building content")
	montageDir := getMontageDir(montage)
	contentFile := filepath.Join(montageDir, "content.tex")
	fmt.Println(BULLET, "creating content.tex with pandoc")
	var args []string
	args = append(args, "-o")
	args = append(args, contentFile)
	for _, file := range configuration.Files {
		args = append(args, filepath.Join(basedir, file))
	}
	cmd := exec.Command("pandoc", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	for _, replacement := range configuration.Replacements {
		fmt.Printf("%s replacing %q -> %q\n", BULLET, replacement.From, replacement.To)
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
	if tag {
		pdfName = fmt.Sprintf("%s - %s [%s-%s].pdf", configuration.Title, configuration.Author, montage.Name, time.Now().Format("20060102150405"))
	}
	pdfFullPath := filepath.Join(basedir, pdfName)
	os.Rename(montagePdfName, pdfFullPath)
	fmt.Printf("%s created %q\n", BULLET, pdfFullPath)
}

func pdflatex(tex string) {
	fmt.Printf("%s running pdflatex on %q\n", BULLET, tex)
	cmd := exec.Command("pdflatex", tex)
	if verbose {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(assembleCmd)
	assembleCmd.Flags().StringVarP(&reference, "montage", "m", "1", "create montage")
	assembleCmd.Flags().BoolVarP(&nocontent, "no-content", "n", false, "don't (re)build content")
	assembleCmd.Flags().BoolVarP(&tag, "tag", "t", false, "tag final PDF with montage name and timestamp")
}