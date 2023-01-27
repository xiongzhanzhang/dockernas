package docker

import (
	"dockernas/internal/config"
	"os"
	"strings"
)

var basePathOnHost = ""

func GetBasePathOnHost() string {
	if config.IsRunInConainer() {
		if basePathOnHost != "" {
			return basePathOnHost
		}

		hostName, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		containers := ListContainer()
		for _, container := range containers {
			data := GetContainerInspect(container.ID)
			if data.Config.Hostname == hostName {
				for _, mount := range data.Mounts {
					if mount.Destination == config.GetBasePath() {
						basePathOnHost = mount.Source
						return mount.Source
					}
				}
			}
		}

		panic("can't get base path on host")
	}
	return config.GetBasePath()
}

func GetPathOnHost(aPath string) string {
	if config.IsRunInConainer() {
		return strings.Replace(aPath, config.GetBasePath(), GetBasePathOnHost(), 1)
	}
	return aPath
}
