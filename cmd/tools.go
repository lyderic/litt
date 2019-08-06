package cmd

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

func getAbsoluteParent(path string) string {
	relativeParent := filepath.Dir(path)
	absoluteParent, err := filepath.Abs(relativeParent)
	if err != nil {
		log.Fatal(err)
	}
	return absoluteParent
}

func getMontageDir(montage Montage) string {
	montageRelativePath := filepath.Join(viper.GetString("basedir"), montage.Path)
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
		if current.Name == viper.GetString("reference") || strconv.Itoa(current.Id) == viper.GetString("reference") {
			found = true
			montage = current
			break
		}
	}
	if !found {
		log.Fatalf("%s: montage not found! Try: '%s list'\n", viper.GetString("reference"), PROGNAME)
	}
	return
}
