package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   PROGNAME,
	Short: "Application to generate books from Markdown files",
	Long: `This program generates PDFs ready to print on KDP or similar
platforms, or manuscripts.

It depends on 'pandoc' and 'pdflatex' being installed on the
computer it is running on.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	config = "./" + PROGNAME + ".json" // default config file
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", config, "configuration `file`")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	viper.Set("basedir", getAbsoluteParent(viper.ConfigFileUsed()))
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
}
