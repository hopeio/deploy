package main

import (
	"encoding/base64"
	"fmt"
	"github.com/hopeio/deploy/plugin/go/drone/notify/dingtalk"
	execi "github.com/hopeio/utils/os/exec"
	"github.com/hopeio/utils/os/fs"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Deploy(ctx *cli.Context) error {
	c := GetConfig(ctx)

	if fs.NotExist(c.DeployDir) {
		os.Mkdir(c.DeployDir, 0666)
	}

	dockerfile, err := os.ReadFile(TplDir + "/Dockerfile-" + c.BuildType)
	if err != nil {
		return err
	}
	dockerfilepath := c.DeployDir + "/" + c.FullName + "-Dockerfile"
	docker := stringsi.FromBytes(dockerfile)
	docker = strings.ReplaceAll(docker, "${app}", c.FullName)
	docker = strings.ReplaceAll(docker, "${cmd}", strings.Join(c.DockerCmds, `", "`))
	_, err = fs.Write(stringsi.ToBytes(docker), dockerfilepath)
	if err != nil {
		return err
	}

	// docker

	execi.RunGetOutWithLog(fmt.Sprintf(`docker build -f %s -t %s %s`, dockerfilepath, c.ImageTag, c.DeployDir))
	execi.RunGetOutWithLog(fmt.Sprintf(`docker login -u %s -p %s`, c.DockerUserName, c.DockerPassword))
	execi.RunGetOutWithLog(fmt.Sprintf(`docker push %s`, c.ImageTag))

	// kubectl
	deployfile, err := os.ReadFile(TplDir + "/deploy-" + c.DeployKind + ".yaml")
	if err != nil {
		return err
	}

	deploy := stringsi.FromBytes(deployfile)
	deploy = strings.ReplaceAll(deploy, "${app}", c.FullName)
	deploy = strings.ReplaceAll(deploy, "${image}", c.ImageTag)
	deploy = strings.ReplaceAll(deploy, "${group}", c.Group)
	deploy = strings.ReplaceAll(deploy, "${datadir}", c.DataDir)
	deploy = strings.ReplaceAll(deploy, "${confdir}", c.ConfDir)
	if c.DeployKind == "cronjob" {
		deploy = strings.ReplaceAll(deploy, "${schedule}", c.Schedule)
	}
	deploypath := c.DeployDir + "/" + c.FullName + "-" + c.DeployKind + ".yaml"
	_, err = fs.Write(stringsi.ToBytes(deploy), deploypath)
	if err != nil {
		return err
	}

	cacrtpath := c.CertDir + "/" + c.Cluster + "/ca.crt"

	if c.CACRT != "" {
		cacrt, err := base64.StdEncoding.DecodeString(c.CACRT)
		if err != nil {
			return err
		}
		_, err = fs.Write(cacrt, cacrtpath)
		if err != nil {
			return err
		}
	}

	devcrtpath := c.CertDir + "/" + c.Cluster + "/dev.crt"
	if c.DEVCRT != "" {
		devcrt, err := base64.StdEncoding.DecodeString(c.DEVCRT)
		if err != nil {
			return err
		}
		_, err = fs.Write(devcrt, devcrtpath)
		if err != nil {
			return err
		}
	}

	devkeypath := c.CertDir + "/" + c.Cluster + "/dev.key"
	if c.DEVKEY != "" {
		devkey, err := base64.StdEncoding.DecodeString(c.DEVKEY)
		if err != nil {
			return err
		}
		_, err = fs.Write(devkey, devkeypath)
		if err != nil {
			return err
		}
	}

	server := "https://hoper.xyz:6443"
	if c.Cluster == "tot" {
		server = "https://192.168.1.212:6443"
	}
	kubeconfig := `--kubeconfig=/root/.kube/config`

	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-cluster k8s --server=%s --certificate-authority=%s --embed-certs=true %s`, server, cacrtpath, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-credentials dev --client-certificate=%s --client-key=%s --embed-certs=true %s`, devcrtpath, devkeypath, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config set-context dev --cluster=k8s --user=dev %s`, kubeconfig))
	execi.RunGetOutWithLog(fmt.Sprintf(`kubectl config use-context dev %s`, kubeconfig))

	if c.DeployKind == "job" || c.DeployKind == "cronjob" {
		execi.RunGetOutWithLog(fmt.Sprintf("kubectl %s delete --ignore-not-found -f %s", kubeconfig, deployfile))
	}
	execi.RunGetOutWithLog(fmt.Sprintf("kubectl %s apply -f %s", kubeconfig, deploypath))

	// notify

	return dingtalk.Notify(dingtalk.GetConfig(ctx))
}
