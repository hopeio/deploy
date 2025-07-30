package main

import (
	"encoding/base64"
	"fmt"
	"github.com/hopeio/deploy/plugin/drone/notify/dingtalk"
	execi "github.com/hopeio/gox/os/exec"
	"github.com/hopeio/gox/os/fs"
	stringsi "github.com/hopeio/gox/strings"
	"os"
	"strings"
)

func Deploy() error {
	if fs.NotExist(config.DeployDir) {
		os.Mkdir(config.DeployDir, 0666)
	}

	dockerfile, err := os.ReadFile(TplDir + "/Dockerfile-" + config.BuildType)
	if err != nil {
		return err
	}
	dockerfilepath := config.DeployDir + "/" + config.FullName + "-Dockerfile"
	docker := stringsi.FromBytes(dockerfile)
	docker = strings.ReplaceAll(docker, "${app}", config.FullName)
	docker = strings.ReplaceAll(docker, "${cmd}", strings.Join(config.DockerCmds, `", "`))
	_, err = fs.Write(stringsi.ToBytes(docker), dockerfilepath)
	if err != nil {
		return err
	}

	// docker

	execi.RunGetOutWithLog(fmt.Sprintf(`docker build -f %s -t %s %s`, dockerfilepath, config.ImageTag, config.DeployDir))
	execi.RunGetOutWithLog(fmt.Sprintf(`docker login -u %s -p %s`, config.DockerUserName, config.DockerPassword))
	execi.RunGetOutWithLog(fmt.Sprintf(`docker push %s`, config.ImageTag))

	// kubectl
	deployfile, err := os.ReadFile(TplDir + "/deploy-" + config.DeployKind + ".yaml")
	if err != nil {
		return err
	}

	deploy := stringsi.FromBytes(deployfile)
	deploy = strings.ReplaceAll(deploy, "${app}", config.FullName)
	deploy = strings.ReplaceAll(deploy, "${image}", config.ImageTag)
	deploy = strings.ReplaceAll(deploy, "${group}", config.Group)
	deploy = strings.ReplaceAll(deploy, "${datadir}", config.DataDir)
	deploy = strings.ReplaceAll(deploy, "${confdir}", config.ConfDir)
	if config.DeployKind == "cronjob" {
		deploy = strings.ReplaceAll(deploy, "${schedule}", config.Schedule)
	}
	deploypath := config.DeployDir + "/" + config.FullName + "-" + config.DeployKind + ".yaml"
	_, err = fs.Write(stringsi.ToBytes(deploy), deploypath)
	if err != nil {
		return err
	}

	cacrtpath := config.CertDir + "/" + config.Cluster + "/ca.crt"

	if config.CACRT != "" {
		cacrt, err := base64.StdEncoding.DecodeString(config.CACRT)
		if err != nil {
			return err
		}
		_, err = fs.Write(cacrt, cacrtpath)
		if err != nil {
			return err
		}
	}

	devcrtpath := config.CertDir + "/" + config.Cluster + "/dev.crt"
	if config.DEVCRT != "" {
		devcrt, err := base64.StdEncoding.DecodeString(config.DEVCRT)
		if err != nil {
			return err
		}
		_, err = fs.Write(devcrt, devcrtpath)
		if err != nil {
			return err
		}
	}

	devkeypath := config.CertDir + "/" + config.Cluster + "/dev.key"
	if config.DEVKEY != "" {
		devkey, err := base64.StdEncoding.DecodeString(config.DEVKEY)
		if err != nil {
			return err
		}
		_, err = fs.Write(devkey, devkeypath)
		if err != nil {
			return err
		}
	}

	server := "https://hoper.xyz:6443"
	if config.Cluster == "tot" {
		server = "https://192.168.1.212:6443"
	}
	kubeconfig := `--kubeconfig=/root/.kube/config`

	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-cluster k8s --server=%s --certificate-authority=%s --embed-certs=true %s`, server, cacrtpath, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-credentials dev --client-certificate=%s --client-key=%s --embed-certs=true %s`, devcrtpath, devkeypath, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-context dev --cluster=k8s --user=dev %s`, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config use-context dev %s`, kubeconfig))

	if config.DeployKind == "job" || config.DeployKind == "cronjob" {
		execi.RunGetOutWithLog(fmt.Sprintf("kubectl %s delete --ignore-not-found -f %s", kubeconfig, deployfile))
	}
	execi.RunGetOutWithLog(fmt.Sprintf("kubectl %s apply -f %s", kubeconfig, deploypath))

	// notify

	return dingtalk.Notify(&config.Config)
}
