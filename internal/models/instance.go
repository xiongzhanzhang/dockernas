package models

import "log"

const (
	NEW_STATE    = 0
	CREATE_ERROR = 1
	RUN_ERROR    = 2
	RUNNING      = 3
)

type Instance struct {
	InstanceID       string `json:"instanceID"`
	Summary          string `json:"summary"`
	State            int    `json:"state"`
	IconUrl          string `json:"iconUrl"`
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

func GetInstance() []Instance {
	var instances []Instance
	err := GetDb().Find(&instances).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return instances
}
