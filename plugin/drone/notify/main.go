package main

import (
	"log"

	"github.com/hopeio/deploy/plugin/drone/notify/dingtalk"
	"github.com/hopeio/gox/flag"
	"github.com/spf13/cobra"
)

var config dingtalk.Config
var rootCmd = &cobra.Command{
	Use:   "notify",
	Short: "通知",
	Run: func(cmd *cobra.Command, args []string) {
		dingtalk.Notify(&config)
	},
}

func main() {
	pflag := rootCmd.PersistentFlags()
	flag.AddFlag(pflag, &config)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
