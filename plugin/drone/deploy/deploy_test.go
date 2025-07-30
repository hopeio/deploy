package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hopeio/gox/eflag"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDeploy(t *testing.T) {
	t.Log(os.Setenv("PLUGIN_DEPLOY_DIR", "./abc"))
	pflag := rootCmd.PersistentFlags()
	eflag.AddFlag(pflag, &config)
	argument := strings.Split(`deploy --repo toolbox --group por --name record_by_id --build_type bin --deploy_kind deployment --step all --docker_username abc --docker_password 123 --docker_cmd por,-c,./conf/local.toml,-t --ding_token abc --ding_secret 123 --commit_tag 1.0.0 --data_dir /data --config_dir /root/config --cluster tx`, " ")
	rootCmd.SetArgs(argument)
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		spew.Dump(config)
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
