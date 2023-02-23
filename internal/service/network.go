package service

import (
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"encoding/json"
	"log"
	"strings"
	"time"
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
	if config.GetIsHttpsEnabled() {
		getCaFilePath(config.GetCaFileDir()) // check ca file first
	}

	gateWayInstance := models.GetInstanceByName(config.GetGateWayInstanceName())
	if gateWayInstance == nil {
		app := GetAppByName("nginx", true)
		if app == nil {
			panic("cant get gateway app nginx")
		}

		var param models.InstanceParam
		param.Name = config.GetGateWayInstanceName()
		param.Summary = "http gateway"
		param.AppName = app.Name
		param.IconUrl = app.IconUrl
		param.NetworkMode = models.BIRDGE_MODE
		param.Version = app.DockerVersions[0].Version
		param.DockerTemplate = app.DockerVersions[0]
		param.DfsVolume[0].Value = config.GetCaFileDir()

		CreateInstance(param, true)
	} else {
		StartInstance(*gateWayInstance)
	}

	config.EnableHttpGateway()
	tryFlushGatewayConfig()
}

func EnableHttps() {
	getCaFilePath(config.GetCaFileDir())
	config.EnableHttps()
	RestartHttpGateway()
}

func DisableHttps() {
	config.DisableHttps()
	tryFlushGatewayConfig()
}

func SetCaFileDir(path string) {
	getCaFilePath(path)
	config.SetCaFileDir(path)
	if config.GetIsHttpsEnabled() {
		utils.RunBackgroundTaskSafe(RestartHttpGateway, time.Second*3)
	}
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

func getCaFilePath(caFileDir string) (string, string) {
	cer, key, msg := utils.GetCaFilePathOnHost(config.GetFullDfsPath(caFileDir), config.GetDomain())
	if msg != "" {
		panic(msg)
	}

	//change to path in nginx container
	fullPath := config.GetFullDfsPath(caFileDir)
	return strings.Replace(cer, fullPath, "/ca", 1), strings.Replace(key, fullPath, "/ca", 1)
}

func updateNginxConfig(instance models.Instance) {
	var param models.InstanceParam
	err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
	if err != nil {
		log.Println("updateNginxConfig error:" + err.Error())
		return
	}
	templateFilePath := config.GetAppMountFilePath(param.Path, "nginx.conf")
	instanceLocalPath := config.GetAppLocalFilePath(instance.Name, "nginx.conf")

	templateData := utils.ReadFile(templateFilePath)

	proxyConfigs := models.GetHttpProxyConfig()
	configStr := `
	    client_max_body_size 102400m;
		map $http_upgrade $connection_upgrade {
			default upgrade;
			''      close;
		}
	`
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
				proxy_pass  http://` + host + `:` + proxyConfig.Port + `;
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

		for _, proxyConfig := range proxyConfigs {
			configStr += `
		server {
			listen 80;
			server_name ` + proxyConfig.HostName + "." + config.GetDomain() + `;
			return 301 https://$server_name$request_uri;
		}
		`
		}
	} else {
		configStr = strings.ReplaceAll(configStr, "PORT_PLACEHOLDER", "80")
		configStr = strings.ReplaceAll(configStr, "SSL_PLACEHOLDER", "")
	}

	utils.WriteFile(instanceLocalPath, strings.Replace(templateData, "#####PLACEHOLDER", configStr, 1))
}
