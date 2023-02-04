package docker

import (
	"bytes"
	"context"
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func Delete(containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return err
	}

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		log.Println("start docker error")
		return err
	}

	return nil
}

func Stop(containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return err
	}

	timeoutSecond := time.Second * 120
	err = cli.ContainerStop(ctx, containerID, &timeoutSecond)

	if err != nil {
		log.Println("stop docker error")
		return err
	}

	return nil
}

func Start(containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return err
	}

	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Println("start docker error")
		return err
	}

	return nil
}

func Restart(containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return err
	}

	timeoutSecond := time.Second * 180
	err = cli.ContainerRestart(ctx, containerID, &timeoutSecond)
	if err != nil {
		log.Println("restart docker error")
		return err
	}

	return nil
}

func PullImage(imageUrl string) io.ReadCloser {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	reader, err2 := cli.ImagePull(ctx, imageUrl, types.ImagePullOptions{})
	if err2 != nil {
		log.Println("pull image error: " + imageUrl)
		panic(err2)
	}

	return reader
}

func DelImage(imageId string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	_, err2 := cli.ImageRemove(ctx, imageId, types.ImageRemoveOptions{})
	if err2 != nil {
		panic(err2)
	}
}

func ListImage() []models.ImageInfo {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	images, err2 := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err2 != nil {
		panic(err2)
	}

	var infos []models.ImageInfo
	for _, v := range images {
		for _, tag := range v.RepoTags {
			infos = append(infos, models.ImageInfo{
				Id:    v.ID,
				Name:  tag,
				Size:  v.Size,
				State: "100%",
			})
		}
	}

	return infos
}

func ListContainer() []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return containers
}

func GetContainerInspect(containerID string) types.ContainerJSON {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	data, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}

	return data
}

func Create(param *models.InstanceParam) (string, error) {
	containerConfig, hostConfig := buildConfig(param)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return "", err
	}

	// _, err = cli.ImagePull(ctx, param.ImageUrl, types.ImagePullOptions{})
	// if err != nil {
	// 	log.Println("pull image error: " + param.ImageUrl)
	// 	return "", err
	// }

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, param.Name)
	if err != nil {
		log.Println("create container error")
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Println("run container error")
		return resp.ID, err
	}

	return resp.ID, nil
}

func GetContainerStat(id string) (types.ContainerStats, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}
	return cli.ContainerStats(ctx, id, false)
}

func replaceVariable(aStr string, param *models.InstanceParam) string {
	if param.OtherParams == nil {
		return aStr
	}

	for _, v := range param.OtherParams {
		if v.OtherType == models.PLACEHOLDER_PARAM {
			aStr = strings.ReplaceAll(aStr, v.Key, v.Value)
		}
	}

	return aStr
}

func buildConfig(param *models.InstanceParam) (container.Config, container.HostConfig) {
	m := make([]mount.Mount, 0, len(param.DfsVolume)+len(param.LocalVolume))

	var usedVolumeName []string
	for index, item := range param.LocalVolume {
		if item.Name == "" || utils.Contains(usedVolumeName, item.Name) {
			panic("local volume name error:" + item.Name)
		}
		usedVolumeName = append(usedVolumeName, item.Name)

		if item.MountFile {
			config.GetLocalVolumePath(param.Name, "") // create dir if not exit
			templateFilePath := config.GetAppMountFilePath(param.AppName, param.Version, item.Name)
			instanceLocalPath := config.GetAppLocalFilePath(param.Name, item.Name)
			if !utils.IsFileExist(instanceLocalPath) {
				_, err := utils.CopyFile(templateFilePath, instanceLocalPath)
				if err != nil {
					panic(err)
				}
			}
			param.LocalVolume[index].Value = instanceLocalPath
			m = append(m, mount.Mount{
				Type:   mount.TypeBind,
				Source: GetPathOnHost(instanceLocalPath),
				Target: item.Key,
			})
		} else {
			localDir := config.GetLocalVolumePath(param.Name, item.Name)
			param.LocalVolume[index].Value = localDir
			m = append(m, mount.Mount{
				Type:   mount.TypeBind,
				Source: GetPathOnHost(localDir),
				Target: item.Key,
			})
		}
	}

	for index, item := range param.DfsVolume {
		//mount a local dir if dfs dir is empty, let user decide whether or not delete data when delete instance
		if item.Value == "" {
			if item.Name == "" || utils.Contains(usedVolumeName, item.Name) {
				panic("local volume name error:" + item.Name)
			}
			usedVolumeName = append(usedVolumeName, item.Name)
			localDir := config.GetLocalVolumePath(param.Name, item.Name)
			m = append(m, mount.Mount{
				Type:   mount.TypeBind,
				Source: GetPathOnHost(localDir),
				Target: item.Key,
			})
		} else {
			if item.Value[0] != '/' {
				item.Value = "/" + item.Value
				param.DfsVolume[index].Value = item.Value
			}
			dfsPath := config.GetFullDfsPath(item.Value)
			utils.CheckCreateDir(dfsPath)
			m = append(m, mount.Mount{
				Type:   mount.TypeBind,
				Source: GetPathOnHost(dfsPath),
				Target: item.Key,
			})
		}
	}

	var envs []string
	for _, item := range param.EnvParams {
		envs = append(envs, replaceVariable(item.Key, param)+"="+replaceVariable(item.Value, param))
	}

	cmdStr := replaceVariable(param.Cmd, param)
	var cmds []string
	if cmdStr != "" {
		cmds = strings.Split(cmdStr, " ")
	}

	exports := make(nat.PortSet)
	netPort := make(nat.PortMap)

	hostIp := "0.0.0.0"
	if param.HostOnly {
		hostIp = "127.0.0.1"
	}
	for _, item := range param.PortParams {
		proto := "tcp"
		if item.Protocol == "udp" {
			proto = "udp"
		}
		natPort, _ := nat.NewPort(proto, item.Key)
		exports[natPort] = struct{}{}
		portList := make([]nat.PortBinding, 0, 1)
		portList = append(portList, nat.PortBinding{HostIP: hostIp, HostPort: item.Value})
		netPort[natPort] = portList
	}

	containerConfig := container.Config{
		Image:        param.ImageUrl,
		ExposedPorts: exports,
		Env:          envs,
		Cmd:          cmds,
	}
	hostConfig := container.HostConfig{
		PortBindings:  netPort,
		Mounts:        m,
		RestartPolicy: container.RestartPolicy{Name: "always"},
	}

	if param.Privileged {
		hostConfig.Privileged = true
	}

	if utils.GetOperationSystemName() == "linux" {
		// if param.Privileged == false {
		// 	curUser, err := user.Current()
		// 	if err != nil {
		// 		panic("get current user error: " + err.Error())
		// 	}
		// 	containerConfig.User = curUser.Uid
		// }
		hostConfig.ExtraHosts = append(hostConfig.ExtraHosts, "host.docker.internal:host-gateway")
	}

	return containerConfig, hostConfig
}

func GetLog(containerID string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Println("get docker log error")
		panic(err)
	}
	defer out.Close()

	var writer bytes.Buffer
	io.Copy(&writer, out)

	return writer.String()
}

func Exec(container string, columns string) types.HijackedResponse {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	ir, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"sh"},
		Env:          []string{"COLUMNS=" + columns, "TERM=xterm-256color"},
		Tty:          true,
	})
	if err != nil {
		log.Println("exec cmd error")
		panic(err)
	}

	// 附加到上面创建的/bin/bash进程中
	hr, err := cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{Detach: false, Tty: true})
	if err != nil {
		log.Println("attch container error")
		panic(err)
	}
	return hr
}
