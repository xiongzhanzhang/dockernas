package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type NetworkInfo struct {
	IP                 string `json:"ip"`
	Domain             string `json:"domain"`
	HttpGatewayEnable  bool   `json:"httpGatewayEnable"`
	HttpGatewayLoading bool   `json:"httpGatewayLoading"`
	HttpsEnable        bool   `json:"httpsEnable"`
	SslCertificatePath string `json:"sslCertificatePath"`
}

type HttpProxyConfig struct {
	HostName     string `json:"hostName"  gorm:"unique;not null"`
	InstanceName string `json:"instanceName"`
	Port         string `json:"port"`
	CreateTime   int64  `json:"createTime"`
}

type InstanceHttpPorts struct {
	InstanceName string   `json:"instanceName"`
	Ports        []string `json:"ports"`
}

func AddHttpProxyConfig(proxyConfig *HttpProxyConfig) {
	err := GetDb().Create(proxyConfig).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func DelHttpProxyConfig(proxyConfig *HttpProxyConfig) {
	err := GetDb().Where("host_name = ?", proxyConfig.HostName).Delete(proxyConfig).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func GetHttpProxyConfig() []HttpProxyConfig {
	var configs []HttpProxyConfig
	err := GetDb().Find(&configs).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return configs
}

func GetHttpProxyConfigByHostName(hostName string) HttpProxyConfig {
	var config HttpProxyConfig
	err := GetDb().First(&config, "host_name=?", hostName).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return config
}

func GetHttpProxyConfigByInstance(instanceName string, port string) *HttpProxyConfig {
	var config HttpProxyConfig
	err := GetDb().First(&config, "instance_name=? and port=?", instanceName, port).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Println(err)
		panic(err)
	}

	return &config
}
