package service

import (
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"log"
	"strconv"
)

func DelInstancePorts(instance models.Instance) {
	err := models.DelPortByInstanceName(instance.Name)
	if err != nil {
		panic(err)
	}
}

func CheckIsPortUsed(param models.InstanceParam) {
	for _, item := range param.PortParams {
		if item.Value == "" {
			continue
		}
		port, err := models.GetInstancePort(item.Protocol, item.Value)
		if err != nil {
			panic(err)
		}
		if port != nil {
			if port.InstanceName == param.Name {
				continue
			} else {
				panic("port " + port.Port + " with " + port.Protocol + " protocol is used by " + port.InstanceName)
			}
		}
		portValue, err := strconv.Atoi(item.Value)
		if err != nil {
			panic(item.Value + " is not a valide port value")
		}
		if portValue >= 65535 {
			panic(item.Value + " is greater than max port value 65535")
		}
		if utils.IsPortUsed(getCheckHost(), item.Protocol, item.Value) {
			panic("port " + item.Value + " with " + item.Protocol + " protocol is used")
		}
	}
}

func getCheckHost() string {
	host := "localhost"
	if config.IsRunInConainer() {
		host = "host.docker.internal"
	}

	return host
}

func SavePortUsed(instance *models.Instance, param models.InstanceParam) {
	for _, item := range param.PortParams {
		models.AddInstancePort(instance.Id,
			instance.Name,
			instance.AppName,
			item.Protocol,
			item.Value,
		)
	}
}

func getFirstHttpPort(param models.InstanceParam) int {
	for _, item := range param.PortParams {
		if item.Protocol == "http" {
			port, err := strconv.Atoi(item.Value)
			if err != nil {
				log.Println(err)
			}
			return port
		}
	}

	return 0
}
