package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "litt",
	Short: "Application to generate books from Markdown files",
	Long: `litt generates PDFs ready to print on KDP or similar
platforms, or manuscripts.

It depends on 'pandoc' and 'pdflatex' being installed on the
computer it is running on.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "litt.json", "configuration `file`")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be verbose")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("litt")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	viper.Set("basedir", getAbsoluteParent(viper.ConfigFileUsed()))
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
}
