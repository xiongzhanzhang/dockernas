package docker

import (
	"bytes"
	"context"
	"io"
	"log"
	"time"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"

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

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
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

func PullImage(imageUrl string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	_, err = cli.ImagePull(ctx, imageUrl, types.ImagePullOptions{})
	if err != nil {
		log.Println("pull image error: " + imageUrl)
		panic(err)
	}
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

func buildConfig(param *models.InstanceParam) (container.Config, container.HostConfig) {
	m := make([]mount.Mount, 0, len(param.DfsVolume)+len(param.LocalVolume))
	for _, item := range param.DfsVolume {
		dfsPath := config.GetFullDfsPath(item.Value)
		utils.CheckCreateDir(dfsPath)
		m = append(m, mount.Mount{
			Type:   mount.TypeBind,
			Source: dfsPath,
			Target: item.Key,
		})
	}

	var usedVolumeName []string
	for index, item := range param.LocalVolume {
		if item.Name == "" || utils.Contains(usedVolumeName, item.Name) {
			panic("local volume name error:" + item.Name)
		}
		usedVolumeName = append(usedVolumeName, item.Name)

		localDir := config.GetLocalVolumePath(param.Name, item.Name)
		m = append(m, mount.Mount{
			Type:   mount.TypeBind,
			Source: localDir,
			Target: item.Key,
		})

		param.LocalVolume[index].Value = localDir
	}

	var envs []string
	for _, item := range param.EnvParams {
		envs = append(envs, item.Key+"="+item.Value)
	}

	exports := make(nat.PortSet)
	netPort := make(nat.PortMap)

	for _, item := range param.PortParams {
		proto := "tcp"
		if item.Protocol == "udp" {
			proto = "udp"
		}
		natPort, _ := nat.NewPort(proto, item.Key)
		exports[natPort] = struct{}{}
		portList := make([]nat.PortBinding, 0, 1)
		portList = append(portList, nat.PortBinding{HostIP: "0.0.0.0", HostPort: item.Value})
		netPort[natPort] = portList
	}

	containerConfig := container.Config{
		Image:        param.ImageUrl,
		ExposedPorts: exports,
		Env:          envs,
	}
	hostConfig := container.HostConfig{
		PortBindings:  netPort,
		Mounts:        m,
		RestartPolicy: container.RestartPolicy{Name: "always"},
	}

	if param.Privileged {
		hostConfig.Privileged = true
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
