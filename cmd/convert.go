package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var convertCmd = &cobra.Command{
	Use:                   "convert",
	Short:                 "Convert configuration file from/to json to/from yaml",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		convert()
	},
}

func convert() {
	configuration.load()
	config := viper.GetString("config")
	var err error
	var data []byte
	var path string
	switch filepath.Ext(config) {
	case ".json":
		path = strings.Replace(viper.GetString("config"), "json", "yaml", 1)
		data, err = yaml.Marshal(&configuration)
		fmt.Printf("Creating YAML file: %q\n", path)
	case ".yaml":
		path := strings.Replace(viper.GetString("config"), "yaml", "json", 1)
		data, err = json.MarshalIndent(&configuration, "", "  ")
		fmt.Printf("Creating JSON file: %q\n", path)
	default:
		abortIfInvalidConfigurationFormat()
	}
	if err != nil {
		tools.PrintRedln("Configuration marshaling failed!")
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(">OK")
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
