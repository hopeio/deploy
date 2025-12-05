package main

import (
	"log"

	"github.com/hopeio/gox/flag"
	"github.com/spf13/cobra"
)

var config Config
var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "发布",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.AfterInject()
	},
	Run: func(cmd *cobra.Command, args []string) {
		Deploy()
	},
}

func main() {
	pflag := rootCmd.PersistentFlags()
	flag.AddFlag(pflag, &config)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
