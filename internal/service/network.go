package service

import (
	"encoding/json"
	"log"
	"net"
	"strings"
	"time"
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func GetNetworkInfo() models.NetworkInfo {
	gateway := getGatewayInstance()

	var networkInfo models.NetworkInfo
	networkInfo.IP = getLocalAddress()
	networkInfo.Domain = config.GetDomain()
	networkInfo.HttpGatewayEnable = gateway != nil

	return networkInfo
}

func getLocalAddress() string {
	con, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
	}
	defer con.Close()

	localAddress := con.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String()
}

func GetInstanceHttpPorts() []models.InstanceHttpPorts {
	instanceHttpPortsList := []models.InstanceHttpPorts{}

	instances := models.GetInstance()
	for _, instance := range instances {
		var param models.InstanceParam
		err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		var instanceHttpPorts models.InstanceHttpPorts
		instanceHttpPorts.InstanceName = instance.Name
		for _, item := range param.PortParams {
			if item.Protocol == "http" {
				instanceHttpPorts.Ports = append(instanceHttpPorts.Ports, item.Value)
			}
		}
		if len(instanceHttpPorts.Ports) > 0 {
			instanceHttpPortsList = append(instanceHttpPortsList, instanceHttpPorts)
		}
	}

	return instanceHttpPortsList
}

func CreateHttpProxyConfig(proxyConfig models.HttpProxyConfig) {
	proxyConfig.CreateTime = time.Now().UnixMilli()
	models.AddHttpProxyConfig(&proxyConfig)
	tryFlushGatewayConfig()
}

func DelHttpProxyConfig(proxyConfig models.HttpProxyConfig) {
	models.DelHttpProxyConfig(&proxyConfig)
	tryFlushGatewayConfig()
}

func SetDomain(domain string) {
	config.SetDomain(domain)
	tryFlushGatewayConfig()
}

func getGatewayInstance() *models.Instance {
	return models.GetInstanceByName(config.GetGateWayInstanceName())
}

func RestartHttpGateway() {
	gateWayInstance := models.GetInstanceByName(config.GetGateWayInstanceName())
	if gateWayInstance == nil {
		panic("http gateway is not start")
	}

	flushGatewayConfig(*gateWayInstance)
	RestartInstance(*gateWayInstance)
}

func EnableHttpGateway() {
	if config.GetDomain() == "" {
		panic("domain is not set")
	}

	gateWayInstance := models.GetInstanceByName(config.GetGateWayInstanceName())
	if gateWayInstance == nil {
		app := GetAppByName("nginx")
		var param models.InstanceParam
		param.Name = config.GetGateWayInstanceName()
		param.Summary = "http gateway"
		param.AppName = app.Name
		param.IconUrl = app.IconUrl
		param.ImageUrl = app.DockerVersions[0].ImageUrl
		param.Version = app.DockerVersions[0].Version
		param.DfsVolume = app.DockerVersions[0].DfsVolume
		param.LocalVolume = app.DockerVersions[0].LocalVolume
		param.EnvParams = app.DockerVersions[0].EnvParams
		param.PortParams = app.DockerVersions[0].PortParams
		gateWayInstance = CreateInstance(param)
	}

	flushGatewayConfig(*gateWayInstance)
}

func tryFlushGatewayConfig() {
	gateway := getGatewayInstance()
	if gateway != nil {
		flushGatewayConfig(*gateway)
	}
}

func flushGatewayConfig(instance models.Instance) {
	updateNginxConfig(instance)
	docker.Restart(instance.ContainerID)
}

func updateNginxConfig(instance models.Instance) {
	templateFilePath := config.GetAppMountFilePath(instance.AppName, instance.Version, "nginx.conf")
	instanceLocalPath := config.GetAppLocalFilePath(instance.Name, "nginx.conf")

	templateData := utils.ReadFile(templateFilePath)

	proxyConfigs := models.GetHttpProxyConfig()
	configStr := `
		map $http_upgrade $connection_upgrade {
			default upgrade;
			''      close;
		}

	`
	for _, proxyConfig := range proxyConfigs {
		configStr += `
		server {
			listen       80;
			server_name  ` + proxyConfig.HostName + "." + config.GetDomain() + `;
	
			location / {
				proxy_set_header Host      $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
				proxy_http_version 1.1;
				proxy_read_timeout 300s;
                proxy_send_timeout 300s;
				proxy_set_header Upgrade $http_upgrade;
        		proxy_set_header Connection  $connection_upgrade;
				proxy_pass  http://host.docker.internal:` + proxyConfig.Port + `;
			}
		}
		`
	}

	utils.WriteFile(instanceLocalPath, strings.Replace(templateData, "#####PLACEHOLDER", configStr, 1))
}
