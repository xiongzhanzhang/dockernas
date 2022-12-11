package docker

import (
	"context"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func Create(param models.InstanceParam) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	_, err = cli.ImagePull(ctx, param.ImageUrl, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	containerConfig, hostConfig := buildConfig(param)

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, param.Name)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return resp.ID
}

func buildConfig(param models.InstanceParam) (container.Config, container.HostConfig) {
	m := make([]mount.Mount, 0, len(param.DfsVolume)+len(param.LocalVolume))
	for _, item := range param.DfsVolume {
		m = append(m, mount.Mount{Type: mount.TypeBind, Source: item.Key, Target: item.Value})
	}

	var usedVolumeName []string
	for _, item := range param.LocalVolume {
		if item.Name == "" || utils.Contains(usedVolumeName, item.Name) {
			panic("local volume name error:" + item.Name)
		}
		usedVolumeName = append(usedVolumeName, item.Name)
		m = append(m, mount.Mount{
			Type:   mount.TypeBind,
			Source: config.GetLocalVolumePath(param.Name, item.Name),
			Target: item.Key,
		})
	}

	var envs []string
	for _, item := range param.EnvParams {
		envs = append(envs, item.Key+"="+item.Value)
	}

	exports := make(nat.PortSet)
	netPort := make(nat.PortMap)

	for _, item := range param.PortParams {
		natPort, _ := nat.NewPort("tcp", item.Key)
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

	return containerConfig, hostConfig
}
