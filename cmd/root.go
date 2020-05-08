package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   PROGNAME,
	Short: "Application to generate books from Markdown files",
	Long: fmt.Sprintf(`%s v.%s (c) Lyderic Landry, London 2019

This program generates PDFs ready to print on KDP or similar
platforms, or manuscripts.

It depends on 'pandoc' and 'pdflatex' being installed on the
computer it is running on.`, PROGNAME, VERSION),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		tools.PrintYellowln(err)
		os.Exit(ROOT_EXECUTE_ERROR)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	config := PROGNAME + ".yaml" // default config file
	debug := false
	rootCmd.PersistentFlags().StringP("config", "c", config, "configuration `file`")
	rootCmd.PersistentFlags().BoolP("debug", "d", debug, "show debugging information")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.SetEnvPrefix(strings.ToUpper(PROGNAME))
	viper.AutomaticEnv() // config file can now be set with envvar 'LITT_CONFIG'
}

func initConfig() {
	viper.Set("basedir", getAbsoluteParent(viper.GetString("config")))
}
