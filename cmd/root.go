package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   PROGNAME,
	Short: "Application to generate books from Markdown files",
	Long: fmt.Sprintf(`%s (c) Lyderic Landry, London 2019

This program generates PDFs ready to print on KDP or similar
platforms, or manuscripts.

It depends on 'pandoc' and 'pdflatex' being installed on the
computer it is running on.`, PROGNAME),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", config, "configuration `file`")
	/*
		BETTER: no need for global 'config' variable + 'config' should not be persistent anyway

		rootCmd.Flags().StringVarP("config", "c", config, "configuration `file`")
		viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
		viper.SetEnvPrefix(strings.ToUpper(PROGNAME))
		viper.AutomaticEnv() // config can now be set with envvar 'LITT_CONFIG'
	*/
	commandsNeedingConfiguration := []*cobra.Command{
		assembleCmd,
		cleanCmd,
		countCmd,
		listCmd,
		totalsCmd,
	}
	for _, command := range commandsNeedingConfiguration {
		command.PreRun = func(cmd *cobra.Command, args []string) {
			configuration.load()
		}
	}
}

/*
func initConfig() {
	if !tools.PathExists(config) {
		tools.PrintRedf("%s: configuration file not found!", config)
		return
	}
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		tools.PrintRedln("Configuration file not readable!")
		log.Fatalf("Aborting")
	}
	viper.Set("basedir", getAbsoluteParent(viper.ConfigFileUsed()))
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
}
*/
