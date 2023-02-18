package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func GetDockerNasNetworkName() string {
	return "DockerNAS"
}

func IsNetworkExist() (bool, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return false, err
	}

	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Println("create docker client error")
		return false, err
	}

	for _, network := range networks {
		if network.Name == GetDockerNasNetworkName() {
			return true, nil
		}
	}

	return false, nil
}

func CheckNetwork() error {
	isExist, err := IsNetworkExist()
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		return err
	}

	_, err = cli.NetworkCreate(ctx, GetDockerNasNetworkName(), types.NetworkCreate{})
	return err
}
