package models

import "log"

type Instance struct {
	InstanceID       string `json:"instanceID"`
	Summary          string `json:"summary"`
	State            int    `json:"state"`
	Port             int    `json:"port"`
	Name             string `json:"name"  gorm:"unique;not null"`
	AppName          string `json:"appName"`
	Version          string `json:"version"`
	InstanceParamStr string `json:"instanceParamStr" gorm:"type:varchar(1024)"` //store json str
}

func AddInstance(instance *Instance) {
	err := GetDb().Create(instance).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func UpdateInstance(instance *Instance) {
	err := GetDb().Model(&Instance{}).Where("name = ?", instance.Name).Updates(instance).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
