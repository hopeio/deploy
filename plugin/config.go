package main

import (
	timei "github.com/hopeio/tailmon/utils/time"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

const TplDir = "/tpl"

type Config struct {
	Repo       string
	DeployDir  string
	CertDir    string
	CommitTag  string
	Group      string
	Name       string
	FullName   string
	DeployKind string
	BuildType  string
	Step       string
	KubeConfig
	DockerConfig
}

type DockerConfig struct {
	DockerfilePath string
	DockerUserName string
	DockerPassword string
	DockerCmds     []string
	ImageTag       string
}

type KubeConfig struct {
	DeployPath            string
	DataDir               string
	ConfDir               string
	Schedule              string
	Cluster               string
	CACRT, DEVCRT, DEVKEY string
}

func GetConfig(c *cli.Context) *Config {
	conf := &Config{
		DeployDir:  c.String("deploy_dir"),
		Repo:       c.String("repo"),
		CommitTag:  c.String("commit_tag"),
		Group:      c.String("group"),
		Name:       c.String("name"),
		DeployKind: c.String("deploy_kind"),
		BuildType:  c.String("build_type"),
		Step:       c.String("step"),
		DockerConfig: DockerConfig{
			DockerUserName: c.String("docker_username"),
			DockerPassword: c.String("docker_password"),
			DockerCmds:     c.StringSlice("docker_cmd"),
			ImageTag:       c.String("image_tag"),
		},
		KubeConfig: KubeConfig{
			DataDir:  c.String("data_dir"),
			ConfDir:  c.String("config_dir"),
			Schedule: c.String("schedule"),
			Cluster:  c.String("cluster"),
			CACRT:    c.String("ca_crt"),
			DEVCRT:   c.String("dev_crt"),
			DEVKEY:   c.String("dev_key"),
		},
	}

	if conf.DeployDir == "" {
		conf.DeployDir = "./deploy"
	}
	conf.DeployDir, _ = strings.CutSuffix(conf.DeployDir, "/")
	conf.CertDir = conf.DeployDir + "/cert"

	if conf.Name == "" {
		conf.FullName = conf.Group
	} else {
		conf.FullName = conf.Group + "-" + conf.Name
	}
	if conf.ImageTag == "" {
		imageTag, found := strings.CutPrefix(conf.CommitTag, conf.FullName+"-v")
		if !found {
			conf.ImageTag = conf.DockerUserName + "/" + conf.FullName + ":" + time.Now().Format(timei.TimeFormatDisplay)
		} else {
			if strings.HasPrefix(conf.CommitTag, "v") {
				conf.ImageTag = conf.DockerUserName + "/" + conf.FullName + ":" + conf.CommitTag[1:]
			} else if tags := strings.Split(conf.CommitTag, "-v"); len(tags) > 1 {
				conf.ImageTag = conf.DockerUserName + "/" + conf.FullName + ":" + tags[1]
			} else {
				conf.ImageTag = conf.DockerUserName + "/" + conf.FullName + ":" + imageTag
			}

		}
	}

	conf.DockerfilePath = TplDir + "/Dockerfile-" + conf.BuildType
	conf.DeployPath = TplDir + "/deploy-" + conf.DeployKind + ".yaml"

	return conf
}
