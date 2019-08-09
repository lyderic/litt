package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
