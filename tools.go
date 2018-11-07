package main

import (
	"log"
	"path/filepath"
	"strconv"
)

func getParent(path string) string {
	relativeParent := filepath.Dir(path)
	absoluteParent, err := filepath.Abs(relativeParent)
	if err != nil {
		log.Fatal(err)
	}
	return absoluteParent
}

func getMontageDir(montage Montage) string {
	montageRelativePath := filepath.Join(basedir, montage.Path)
	montageRelativeDir := filepath.Dir(montageRelativePath)
	var montageAbsoluteDir string
	var err error
	if montageAbsoluteDir, err = filepath.Abs(montageRelativeDir); err != nil {
		log.Fatal(err)
	}
	return montageAbsoluteDir
}

func getSelectedMontage() (montage Montage) {
	found := false
	for _, current := range configuration.Montages {
		if current.Name == reference || strconv.Itoa(current.Id) == reference {
			found = true
			montage = current
			break
		}
	}
	if !found {
		log.Fatalf("%s: montage not found! Try: litt -l\n", reference)
	}
	return
}
