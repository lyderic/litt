package main

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
)

func assemble() {
	montage := getSelectedMontage()
	fmt.Printf("Assembling montage %q\n", montage.Name)
	if !nosanitize {
		sanitizeAllFiles()
	}
	if !nocontent {
		buildContent(montage)
	}
	pdf(montage)
}

func buildContent(montage Montage) {
	fmt.Println("> pandoc: building content")
	montageDir := getMontageDir(montage)
	contentFile := filepath.Join(montageDir, "content.tex")
	var args []string
	args = append(args, "-o")
	args = append(args, contentFile)
	for _, file := range configuration.Files {
		args = append(args, filepath.Join(basedir, file))
	}
	cmd := exec.Command("pandoc", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	//fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("> replacing \\section to \\chapter*")
	input, err := ioutil.ReadFile(contentFile)
	if err != nil {
		log.Fatal(err)
	}
	output := bytes.Replace(input, []byte("\\section"), []byte("\\chapter*"), -1)
	if err = ioutil.WriteFile(contentFile, output, 0644); err != nil {
		log.Fatal(err)
	}
}

func pdf(montage Montage) {
	fmt.Println("pdflatex", montage.Path)
	montageDir := getMontageDir(montage)
	montageTexName := filepath.Base(montage.Path)
	montagePdfName := strings.Replace(montageTexName, ".tex", ".pdf", 1)
	os.Chdir(montageDir)
	cmd := exec.Command("pdflatex", montageTexName)
	if verbose {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	pdfName := fmt.Sprintf("%s - %s.pdf", configuration.Title, configuration.Autor)
	if tag {
		pdfName = fmt.Sprintf("%s - %s [%s-%s].pdf", configuration.Title, configuration.Autor, montage.Name, time.Now().Format("20060102150405"))
	}
	fmt.Printf("moving PDF to %s\n", pdfName)
	os.Rename(montagePdfName, filepath.Join(basedir, pdfName))
}
