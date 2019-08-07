package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

type Montage struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (montage Montage) String() string {
	return fmt.Sprintf("[%02d] %s [%s]", montage.Id, montage.Name, montage.Path)
}

type Replacement struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Configuration struct {
	Author       string        `json:"author"`
	Title        string        `json:"title"`
	Montages     []Montage     `json:"montages"`
	Files        []string      `json:"files"`
	Replacements []Replacement `json:"replacements"`
	Double       bool          `json:"double"` // when double compilation is required
}

func (configuration *Configuration) load() {
	if !tools.PathExists(config) {
		tools.PrintRedf("%q: configuration file not found!\n", config)
		os.Exit(CONFIG_FILE_NOT_FOUND)
	}
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		tools.PrintRedf("%q: cannot load configuration!\n", config)
		os.Exit(CONFIG_FILE_NOT_LOADABLE)
	}
	viper.Set("basedir", getAbsoluteParent(viper.ConfigFileUsed()))
	err := viper.Unmarshal(&configuration)
	if err != nil {
		tools.PrintRedf("%q: invalid configuration file!\n", config)
		os.Exit(INVALID_CONFIG_FILE)
	}
	configuration.check()
}

func (configuration *Configuration) check() {
	for _, file := range configuration.Files {
		path := filepath.Join(viper.GetString("basedir"), file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			tools.PrintRedf("Error in configuration file: %q\nListed file not found on disk: %q\n", viper.ConfigFileUsed(), path)
			os.Exit(LISTED_FILE_NOT_FOUND)
		}
	}
	for _, montage := range configuration.Montages {
		path := filepath.Join(viper.GetString("basedir"), montage.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			tools.PrintRedf("Error in configuration file: %q\nListed montage not found on disk: %+v\nFile not found: %q\n", viper.ConfigFileUsed(), montage, path)
			os.Exit(LISTED_MONTAGE_NOT_FOUND)
		}
	}
}
