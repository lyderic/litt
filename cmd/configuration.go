package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Montage struct {
	Id   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Path string `json:"path" yaml:"path"`
}

func (montage Montage) String() string {
	return fmt.Sprintf("[%02d] %s [%s]", montage.Id, montage.Name, montage.Path)
}

type Replacement struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

type Configuration struct {
	Author       string        `json:"author,omitempty" yaml:"author,omitempty"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	Montages     []Montage     `json:"montages,omitempty" yaml:"montages,omitempty"`
	Files        []string      `json:"files,omitempty" yaml:"files,omitempty"`
	Replacements []Replacement `json:"replacements,omitempty" yaml:"replacements,omitempty"`
	Double       bool          `json:"double,omitempty" yaml:"double,omitempty"` // when double compilation is required
}

func (configuration *Configuration) load() {
	config := viper.GetString("config")
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(config); err != nil {
		tools.PrintRedln("Cannot load configuration")
		fmt.Printf("%v\nTry: '%s init' or '%s -c litt.json <command>'\n", err, PROGNAME, PROGNAME)
		os.Exit(CONFIG_FILE_NOT_LOADABLE)
	}
	switch filepath.Ext(config) {
	case ".json":
		err = json.Unmarshal(content, &configuration)
	case ".yaml":
		err = yaml.Unmarshal(content, &configuration)
	default:
		tools.PrintRedf("Invalid configuration format: %s. Only json or yaml are valid.\n", filepath.Ext(config))
		os.Exit(INVALID_CONFIGURATION_FORMAT)
	}
	if err != nil {
		tools.PrintRedf("%q: invalid configuration file!\n%s %v\n", config, tools.PROMPT, err)
		os.Exit(UNMARSHALING_FAILED)
	}
	configuration.check()
}

func (configuration *Configuration) check() {
	if len(configuration.Files) == 0 {
		tools.PrintYellowf("No markdown file defined in %q\n", viper.GetString("config"))
	} else {
		for _, file := range configuration.Files {
			path := filepath.Join(viper.GetString("basedir"), file)
			finfo, err := os.Stat(path)
			if os.IsNotExist(err) {
				tools.PrintRedf("Error in configuration file: %q\nListed file not found on disk: %q\n", viper.GetString("config"), path)
				os.Exit(LISTED_FILE_NOT_FOUND)
			}
			if finfo.Size() == 0 {
				tools.PrintRedf("This file is empty: %s\n", path)
				os.Exit(EMPTY_FILE)
			}
		}
	}
	if len(configuration.Montages) == 0 {
		tools.PrintYellowf("No montage defined in %q\n", viper.GetString("config"))
	} else {
		for _, montage := range configuration.Montages {
			path := filepath.Join(viper.GetString("basedir"), montage.Path)
			finfo, err := os.Stat(path)
			if os.IsNotExist(err) {
				tools.PrintRedf("Error in configuration file: %q\nListed montage not found on disk: %+v\nFile not found: %q\n", viper.GetString("config"), montage, path)
				os.Exit(LISTED_MONTAGE_NOT_FOUND)
			}
			if finfo.Size() == 0 {
				tools.PrintRedf("This file is empty: %s\n", path)
				os.Exit(EMPTY_FILE)
			}
		}
	}
}

func (configuration *Configuration) persist() {
	config := viper.GetString("config")
	var data []byte
	var err error
	switch filepath.Ext(config) {
	case ".json":
		data, err = json.MarshalIndent(configuration, "", "  ")
	case ".yaml":
		data, err = yaml.Marshal(configuration)
	default:
		tools.PrintRedf("Invalid configuration format: %s. Only json or yaml are valid.\n", filepath.Ext(config))
		os.Exit(INVALID_CONFIGURATION_FORMAT)
	}
	if err != nil {
		tools.PrintRedf("JSON marshaling failed! %v\n", err)
		os.Exit(MARSHALING_FAILED)
	}
	if ioutil.WriteFile(config, data, 0644); err != nil {
		tools.PrintRedf("Persisting configuration failed! %v\n", err)
		os.Exit(WRITE_FILE_FAILED)
	}
}
