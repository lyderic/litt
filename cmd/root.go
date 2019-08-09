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
	config := "./" + PROGNAME + ".json" // default config file
	rootCmd.PersistentFlags().StringP("config", "c", config, "configuration `file`")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetEnvPrefix(strings.ToUpper(PROGNAME))
	viper.AutomaticEnv() // config file can now be set with envvar 'LITT_CONFIG'
	viper.Set("basedir", getAbsoluteParent(viper.GetString("config")))
}
