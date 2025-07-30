package main

import (
	"github.com/hopeio/deploy/plugin/drone/notify/dingtalk"
	timei "github.com/hopeio/gox/time"
	"strings"
	"time"
)

const TplDir = "/tpl"

type Config struct {
	DeployDir  string `flag:"name:deploy_dir;usage:deploy dir;env:PLUGIN_DEPLOY_DIR"`
	CertDir    string `flag:"name:cert_dir;usage:cert_dir;env:PLUGIN_CERT_DIR"`
	Group      string `flag:"name:group;usage:deploy group;env:PLUGIN_GROUP"`
	Name       string `flag:"name:name;usage:deploy name;env:PLUGIN_NAME"`
	FullName   string `flag:"name:full_name;usage:deploy full name;env:PLUGIN_FULL_NAME"`
	DeployKind string `flag:"name:deploy_kind;usage:deploy kind;env:PLUGIN_DEPLOY_KIND"`
	BuildType  string `flag:"name:build_type;usage:build_type;env:PLUGIN_BUILD_TYPE"`
	Step       string `flag:"name:step;usage:deploy step;env:PLUGIN_STEP"`
	KubeConfig
	DockerConfig
	dingtalk.Config
}
type DockerConfig struct {
	DockerfilePath string   `flag:"name:dockerfile_path;usage:dockerfile_path;env:PLUGIN_DOCKERFILE_PATH"`
	DockerUserName string   `flag:"name:docker_username;usage:docker_username;env:PLUGIN_DOCKER_USERNAME"`
	DockerPassword string   `flag:"name:docker_password;usage:docker_password;env:PLUGIN_DOCKER_PASSWORD"`
	DockerCmds     []string `flag:"name:docker_cmd;usage:docker_cmd;env:PLUGIN_DOCKER_CMD"`
	ImageTag       string   `flag:"name:image_tag;usage:image_tag;env:PLUGIN_IMAGE_TAG"`
}

type KubeConfig struct {
	DeployPath string `flag:"name:deploy_path;usage:deploy path;env:PLUGIN_DEPLOY_PATH"`
	DataDir    string `flag:"name:data_dir;usage:volume data directory;env:PLUGIN_DATA_DIR"`
	ConfDir    string `flag:"name:config_dir;usage:volume config directory;env:PLUGIN_CONFIG_DIR"`
	Schedule   string `flag:"name:schedule;usage:cronjob schedule;env:PLUGIN_SCHEDULE"`
	Cluster    string `flag:"name:cluster;usage:deploy cluster;env:PLUGIN_CLUSTER"`
	CACRT      string `flag:"name:ca_crt;usage:ca.crt;env:PLUGIN_CA_CRT"`
	DEVCRT     string `flag:"name:dev_crt;usage:dev.crt;env:PLUGIN_DEV_CRT"`
	DEVKEY     string `flag:"name:dev_key;usage:dev.key;env:PLUGIN_DEV_KEY"`
}

func (conf *Config) AfterInject() {

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
			conf.ImageTag = conf.DockerUserName + "/" + conf.FullName + ":" + time.Now().Format(timei.LayoutDateTime)
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
}
