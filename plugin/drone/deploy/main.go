package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "deploy",
		Usage: "发布",
		Action: func(c *cli.Context) error {
			return Deploy(c)
		},
		Flags: flags,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "group",
		Usage:   "deploy group",
		EnvVars: []string{"PLUGIN_GROUP"},
	},
	&cli.StringFlag{
		Name:    "name",
		Usage:   "deploy name",
		EnvVars: []string{"PLUGIN_NAME"},
	},
	&cli.StringFlag{
		Name:    "deploy_dir",
		Usage:   "deploy dir",
		EnvVars: []string{"PLUGIN_DEPLOY_DIR"},
	},
	&cli.StringFlag{
		Name:    "build_type",
		Usage:   "build type",
		EnvVars: []string{"PLUGIN_BUILD_TYPE"},
	},
	&cli.StringFlag{
		Name:    "deploy_kind",
		Usage:   "deploy kind",
		EnvVars: []string{"PLUGIN_DEPLOY_KIND"},
	},
	&cli.StringFlag{
		Name:    "step",
		Usage:   "deploy step",
		EnvVars: []string{"PLUGIN_STEP"},
	},
	&cli.StringFlag{
		Name:    "docker_username",
		Usage:   "docker username",
		EnvVars: []string{"PLUGIN_DOCKER_USERNAME"},
	},
	&cli.StringFlag{
		Name:    "docker_password",
		Usage:   "docker password",
		EnvVars: []string{"PLUGIN_DOCKER_PASSWORD"},
	},
	&cli.StringSliceFlag{
		Name:    "docker_cmd",
		Usage:   "docker cmd",
		EnvVars: []string{"PLUGIN_DOCKER_CMD"},
	},
	&cli.StringFlag{
		Name:    "image_tag",
		Usage:   "image tag",
		EnvVars: []string{"PLUGIN_IMAGE_TAG"},
	},
	&cli.StringFlag{
		Name:    "ding_token",
		Usage:   "dingding webhook url",
		EnvVars: []string{"PLUGIN_DING_TOKEN"},
	},
	&cli.StringFlag{
		Name:    "ding_secret",
		Usage:   "dingding secret",
		EnvVars: []string{"PLUGIN_DING_SECRET"},
	},
	&cli.StringFlag{
		Name:    "data_dir",
		Usage:   "volume data directory",
		EnvVars: []string{"PLUGIN_DATA_DIR"},
	},
	&cli.StringFlag{
		Name:    "config_dir",
		Usage:   "volume config directory",
		EnvVars: []string{"PLUGIN_CONFIG_DIR"},
	},
	&cli.StringFlag{
		Name:    "schedule",
		Usage:   "cronjob schedule",
		EnvVars: []string{"PLUGIN_SCHEDULE"},
	},
	&cli.StringFlag{
		Name:    "cluster",
		Usage:   "deploy cluster",
		EnvVars: []string{"PLUGIN_CLUSTER"},
	},
	&cli.StringFlag{
		Name:    "ca_crt",
		Usage:   "cluster ca.crt",
		EnvVars: []string{"PLUGIN_CA_CRT"},
	},
	&cli.StringFlag{
		Name:    "dev_crt",
		Usage:   "cluster dev.crt",
		EnvVars: []string{"PLUGIN_DEV_CRT"},
	},
	&cli.StringFlag{
		Name:    "dev_key",
		Usage:   "cluster dev.key",
		EnvVars: []string{"PLUGIN_DEV_KEY"},
	},
	&cli.StringFlag{
		Name:    "repo",
		Usage:   "git repo",
		EnvVars: []string{"DRONE_REPO"},
	},
	&cli.StringFlag{
		Name:    "commit",
		Usage:   "git commit",
		EnvVars: []string{"DRONE_COMMIT"},
	},
	&cli.StringFlag{
		Name:    "commit_author_name",
		Usage:   "git commit author name",
		EnvVars: []string{"DRONE_COMMIT_AUTHOR_NAME"},
	},
	&cli.StringFlag{
		Name:    "commit_author",
		Usage:   "git commit author",
		EnvVars: []string{"DRONE_COMMIT_AUTHOR"},
	},
	&cli.StringFlag{
		Name:    "commit_link",
		Usage:   "git commit link",
		EnvVars: []string{"DRONE_COMMIT_LINK"},
	},
	&cli.StringFlag{
		Name:    "commit_ref",
		Usage:   "git commit ref",
		EnvVars: []string{"DRONE_COMMIT_REF"},
	},
	&cli.StringFlag{
		Name:    "commit_message",
		Usage:   "git commit message",
		EnvVars: []string{"DRONE_COMMIT_MESSAGE"},
	},
	&cli.StringFlag{
		Name:    "commit_branch",
		Usage:   "git commit branch",
		EnvVars: []string{"DRONE_COMMIT_BRANCH"},
	},
	&cli.StringFlag{
		Name:    "commit_tag",
		Usage:   "git commit tag",
		EnvVars: []string{"DRONE_TAG"},
	},
	&cli.StringFlag{
		Name:    "drone_build_link",
		Usage:   "drone build link",
		EnvVars: []string{"DRONE_BUILD_LINK"},
	},
}
