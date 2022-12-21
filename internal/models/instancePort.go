package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type InstancePort struct {
	InstanceId   int    `json:"instanceId"`
	InstanceName string `json:"instanceName"`
	AppName      string `json:"appName"`
	Protocol     string `json:"protocol"`
	Port         string `json:"port"`
}

func AddInstancePort(InstanceID int, InstanceName string, AppName string, Protocol string, Port string) error {
	var instancePort = InstancePort{
		InstanceId:   InstanceID,
		InstanceName: InstanceName,
		AppName:      AppName,
		Protocol:     Protocol,
		Port:         Port,
	}

	err := GetDb().Create(&instancePort).Error
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetInstancePort(Protocol string, Port string) (*InstancePort, error) {
	var port InstancePort
	err := GetDb().Where("protocol = ? AND port = ?", Protocol, Port).First(&port).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return &port, nil
}

func DelPortByInstanceName(InstanceName string) error {
	err := db.Where("instance_name = ?", InstanceName).Delete(&InstancePort{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
