package service

import (
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"encoding/json"
	"log"
	"time"
)

func GetNetworkInfo() models.NetworkInfo {
	gateway := getGatewayInstance()

	var networkInfo models.NetworkInfo
	networkInfo.IP = utils.GetLocalAddress()
	networkInfo.IPv6 = utils.GetLocalAddressIpv6()
	networkInfo.Domain = config.GetDomain()

	networkInfo.HttpsEnable = config.GetIsHttpsEnabled()
	if gateway != nil && gateway.State != models.STOPPED {
		networkInfo.HttpGatewayEnable = true
		if gateway.State != models.RUNNING {
			networkInfo.HttpGatewayLoading = true
		}
	}

	return networkInfo
}

func GetInstanceHttpPorts() []models.InstanceHttpPorts {
	instanceHttpPortsList := []models.InstanceHttpPorts{}

	instances := models.GetInstance()
	for _, instance := range instances {
		var param models.InstanceParam
		err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
		if err != nil {
			log.Println(err)
			continue
		}
		var instanceHttpPorts models.InstanceHttpPorts
		instanceHttpPorts.InstanceName = instance.Name
		for _, item := range param.PortParams {
			if item.Protocol == "http" {
				if param.NetworkMode == models.HOST_MODE {
					instanceHttpPorts.Ports = append(instanceHttpPorts.Ports, item.Key)
				} else {
					if item.Value == "" || param.NetworkMode == models.NOBUND_MODE {
						instanceHttpPorts.Ports = append(instanceHttpPorts.Ports, item.Key)
					} else {
						instanceHttpPorts.Ports = append(instanceHttpPorts.Ports, item.Value)
					}
				}
			}
		}
		if len(instanceHttpPorts.Ports) > 0 {
			instanceHttpPortsList = append(instanceHttpPortsList, instanceHttpPorts)
		}
	}

	return instanceHttpPortsList
}

func getGatewayInstance() *models.Instance {
	return models.GetInstanceByName(config.GetGateWayInstanceName())
}

func CreateHttpProxyConfig(proxyConfig models.HttpProxyConfig) {
	proxyConfig.CreateTime = time.Now().UnixMilli()
	models.AddHttpProxyConfig(&proxyConfig)
	utils.RunBackgroundTaskSafe(tryFlushGatewayConfig, time.Second*3)
}

func DelHttpProxyConfig(proxyConfig models.HttpProxyConfig) {
	models.DelHttpProxyConfig(&proxyConfig)
	utils.RunBackgroundTaskSafe(tryFlushGatewayConfig, time.Second*3)
}

func SetDomain(domain string) {
	config.SetDomain(domain)
	config.DisableHttps()
	utils.RunBackgroundTaskSafe(tryFlushGatewayConfig, time.Second*3)
}

func RestartHttpGateway() {
	StopHttpGateway()
	EnableHttpGateway()
}

func StopHttpGateway() {
	config.DisableHttpGateway()
	gateway := getGatewayInstance()
	if gateway != nil {
		DeleteInstance(*gateway)
	}
}

func EnableHttpGateway() {
	if config.GetDomain() == "" {
		panic("domain is not set")
	}

	gateWayInstance := models.GetInstanceByName(config.GetGateWayInstanceName())
	if gateWayInstance == nil {
		app := GetAppByName("caddy", true)
		if app == nil {
			panic("cant get gateway app caddy")
		}

		var param models.InstanceParam
		param.Name = config.GetGateWayInstanceName()
		param.Summary = "http gateway"
		param.AppName = app.Name
		param.IconUrl = app.IconUrl
		param.NetworkMode = models.BIRDGE_MODE
		param.Version = app.DockerVersions[0].Version
		param.DockerTemplate = app.DockerVersions[0]

		CreateInstance(param, true)
	} else {
		StartInstance(*gateWayInstance)
	}

	config.EnableHttpGateway()
	tryFlushGatewayConfig()
}

func EnableHttps() {
	config.EnableHttps()
	RestartHttpGateway()
}

func DisableHttps() {
	config.DisableHttps()
	tryFlushGatewayConfig()
}

func tryFlushGatewayConfig() {
	gateway := getGatewayInstance()
	if gateway != nil {
		updateGatewayConfig(*gateway)
		if gateway.State == models.STOPPED {
			StartInstance(*gateway)
		} else {
			RestartInstance(*gateway)
		}
	}
}

func updateGatewayConfig(instance models.Instance) {
	var param models.InstanceParam
	err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
	if err != nil {
		log.Println("updateGatewayConfig error:" + err.Error())
		return
	}

	proxyConfigs := models.GetHttpProxyConfig()
	configStr := ""
	for _, proxyConfig := range proxyConfigs {
		host := "host.docker.internal"
		if proxyConfig.InstanceName != "" {
			proxyedInstance := models.GetInstanceByName(proxyConfig.InstanceName)
			if proxyedInstance == nil {
				continue
			}
			var param models.InstanceParam
			err := json.Unmarshal([]byte(proxyedInstance.InstanceParamStr), &param)
			if err != nil {
				log.Println(err)
			} else {
				if param.NetworkMode == models.NOBUND_MODE {
					host = proxyConfig.InstanceName
				} else if param.NetworkMode != models.HOST_MODE {
					for _, item := range param.PortParams {
						if item.Protocol == "http" {
							if item.Key == proxyConfig.Port && item.Value == "" {
								host = proxyConfig.InstanceName
								break
							}
						}
					}
				}
			}
		}

		directive := ""
		if config.GetIsHttpsEnabled() {
			directive += "https://" + proxyConfig.HostName + "." + config.GetDomain()
		} else {
			directive += "http://" + proxyConfig.HostName + "." + config.GetDomain()
		}

		configStr += `` + directive + ` {
  log
  reverse_proxy ` + host + `:` + proxyConfig.Port + `
}

`
	}

	instanceLocalPath := config.GetAppLocalFilePath(instance.Name, "Caddyfile")
	utils.WriteFile(instanceLocalPath, configStr)
}
