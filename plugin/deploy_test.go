package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"strings"
	"testing"
)

func TestDeploy(t *testing.T) {
	app := &cli.App{
		Name:  "deploy",
		Usage: "发布",
		Action: func(c *cli.Context) error {
			return Deploy(c)
		},
		Flags: flags,
	}

	argument := strings.Split(`deploy --repo toolbox --group por --name record_by_id --build_type bin --deploy_kind deployment --step all --docker_username abc --docker_password 123 --docker_cmd por,-c,./conf/local.toml,-t --ding_token abc --ding_secret 123 --commit_tag 1.0.0 --data_dir /data --config_dir /root/config --cluster tx`, " ")

	if err := app.Run(argument); err != nil {
		log.Fatal(err)
	}
}
