package cmd

import (
	"log"

	"github.com/spf13/cobra"
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

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", config, "configuration `file`")
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
