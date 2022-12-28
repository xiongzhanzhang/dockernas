package service

import (
	"encoding/json"
	"log"
	"net"
	"time"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
)

func GetNetworkInfo() models.NetworkInfo {
	var networkInfo models.NetworkInfo
	networkInfo.IP = getLocalAddress()
	networkInfo.Domain = config.GetDomain()

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
}

func DelHttpProxyConfig(proxyConfig models.HttpProxyConfig) {
	models.DelHttpProxyConfig(&proxyConfig)
}

func SetDomain(domain string) {
	config.SetDomain(domain)
}
