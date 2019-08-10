package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var convertCmd = &cobra.Command{
	Use:                   "convert",
	Short:                 "Convert configuration file from json to yaml",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		convert()
	},
}

func convert() {
	configuration.load()
	config := viper.GetString("config")
	switch filepath.Ext(config) {
	case ".json":
		yamlpath := strings.Replace(viper.GetString("config"), "json", "yaml", 1)
		var err error
		var data []byte
		if data, err = yaml.Marshal(&configuration); err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(yamlpath, data, 0644); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("YAML file created: %q\n", yamlpath)
	case ".yaml":
		jsonpath := strings.Replace(viper.GetString("config"), "yaml", "json", 1)
		var err error
		var data []byte
		if data, err = json.MarshalIndent(&configuration, "", "  "); err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(jsonpath, data, 0644); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("JSON file created: %q\n", jsonpath)
	default:
		tools.PrintRedf("Invalid configuration format: %s. Only json or yaml are valid.\n", filepath.Ext(config))
		os.Exit(INVALID_CONFIGURATION_FORMAT)
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
