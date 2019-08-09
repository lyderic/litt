package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert configuration file from json to yaml",
	Run: func(cmd *cobra.Command, args []string) {
		convert()
	},
}

func convert() {
	configuration.load()
	yamlfile := strings.Replace(viper.GetString("config"), "json", "yaml", 1)
	yamlpath := filepath.Join(viper.GetString("basedir"), yamlfile)
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
