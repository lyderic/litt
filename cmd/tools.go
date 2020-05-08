package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lyderic/tools"
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

func getSelectedMontage() (montage Montage, err error) {
	if len(configuration.Montages) == 0 {
		err = fmt.Errorf("No montage defined in configuration file %q", viper.GetString("config"))
		return
	}
	found := false
	for _, current := range configuration.Montages {
		if current.Name == viper.GetString("reference") || strconv.Itoa(current.Id) == viper.GetString("reference") {
			found = true
			montage = current
			break
		}
	}
	if !found {
		err = fmt.Errorf("%s: montage not found! Try: '%s list'\n", viper.GetString("reference"), PROGNAME)
		return
	}
	return
}

func sanitizeAllFiles() {
	fmt.Println("Sanitizing")
	var idx int
	var file string
	for idx, file = range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
		sanitize(path, true)
	}
	n := idx + 1
	fmt.Printf("%s %d file%s processed\n", tools.PROMPT, n, tools.Ternary(n > 1, "s", ""))
}

func abortIfInvalidConfigurationFormat() {
	tools.PrintRedf("Invalid configuration format: %q. Only json or yaml are valid.\n", filepath.Ext(viper.GetString("config")))
	os.Exit(INVALID_CONFIGURATION_FORMAT)
}
