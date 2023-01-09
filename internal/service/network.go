package service

import (
	"encoding/json"
	"log"
	"strings"
	"time"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func GetNetworkInfo() models.NetworkInfo {
	gateway := getGatewayInstance()

	var networkInfo models.NetworkInfo
	networkInfo.IP = utils.GetLocalAddress()
	networkInfo.IPv6 = utils.GetLocalAddressIpv6()
	networkInfo.Domain = config.GetDomain()

	networkInfo.HttpsEnable = config.GetIsHttpsEnabled()
	networkInfo.SslCertificatePath = config.GetCaFileDir()
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
	tryFlushGatewayConfig()
}

func StopHttpGateway() {
	gateway := getGatewayInstance()
	if gateway != nil {
		StopInstance(*gateway)
	}
}

func EnableHttpGateway() {
	if config.GetDomain() == "" {
		panic("domain is not set")
	}
	if config.GetIsHttpsEnabled() {
		getCaFilePath(config.GetCaFileDir()) // check ca file first
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
		param.DfsVolume = []models.ParamItem{}
		param.LocalVolume = app.DockerVersions[0].LocalVolume
		param.EnvParams = app.DockerVersions[0].EnvParams
		param.PortParams = app.DockerVersions[0].PortParams

		if config.GetIsHttpsEnabled() {
			param.DfsVolume = app.DockerVersions[0].DfsVolume
			param.DfsVolume[0].Value = config.GetCaFileDir()
		}

		CreateInstance(param, true)
	}

	tryFlushGatewayConfig()
}

func EnableHttps() {
	getCaFilePath(config.GetCaFileDir())
	config.EnableHttps()
	tryFlushGatewayConfig()
}

func DisableHttps() {
	config.DisableHttps()
	tryFlushGatewayConfig()
}

func SetCaFileDir(path string) {
	getCaFilePath(path)
	config.SetCaFileDir(path)

	tryFlushGatewayConfig()
}

func getCaFilePath(caFileDir string) (string, string) {
	// caFileDir := config.GetCaFileDir()
	if caFileDir == "" {
		panic("ca file dir is not set")
	}
	domain := config.GetDomain()
	if domain == "" {
		panic("domain is not set")
	}

	fullPath := config.GetFullDfsPath(caFileDir)

	cer := fullPath + "/" + domain + ".cer"
	key := fullPath + "/" + domain + ".key"
	if !utils.IsFileExist(cer) || !utils.IsFileExist(key) {
		cer = fullPath + "/" + domain + "/" + domain + ".cer"
		key = fullPath + "/" + domain + "/" + domain + ".key"
		if !utils.IsFileExist(cer) || !utils.IsFileExist(key) {
			panic("can't find ca file under " + caFileDir)
		}
	}

	//change to path in nginx container
	return strings.Replace(cer, fullPath, "/ca", 1), strings.Replace(key, fullPath, "/ca", 1)
}

func tryFlushGatewayConfig() {
	gateway := getGatewayInstance()
	if gateway != nil {
		updateNginxConfig(*gateway)
		if gateway.State == models.STOPPED {
			StartInstance(*gateway)
		} else {
			RestartInstance(*gateway)
		}
	}
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
			listen       PORT_PLACEHOLDER;
			server_name  ` + proxyConfig.HostName + "." + config.GetDomain() + `;

			SSL_PLACEHOLDER
	
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

	if config.GetIsHttpsEnabled() {
		cer, key := getCaFilePath(config.GetCaFileDir())
		sslCOnfig := `
		ssl_certificate "` + cer + `";
		ssl_certificate_key "` + key + `";
		ssl_session_timeout 5m;
		ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
		ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
		ssl_prefer_server_ciphers on;
		`
		configStr = strings.ReplaceAll(configStr, "PORT_PLACEHOLDER", "443 ssl http2")
		configStr = strings.ReplaceAll(configStr, "SSL_PLACEHOLDER", sslCOnfig)
	} else {
		configStr = strings.ReplaceAll(configStr, "PORT_PLACEHOLDER", "80")
		configStr = strings.ReplaceAll(configStr, "SSL_PLACEHOLDER", "")
	}

	utils.WriteFile(instanceLocalPath, strings.Replace(templateData, "#####PLACEHOLDER", configStr, 1))
}
