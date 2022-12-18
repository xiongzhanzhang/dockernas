package service

import (
	"log"
	"os"
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func CreateInstance(param models.InstanceParam) {
	var instance models.Instance
	var err error

	instance.Summary = param.Summary
	instance.Name = param.Name
	instance.State = models.NEW_STATE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	models.AddInstance(&instance)

	instance.InstanceID, err = docker.Create(&param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	if err != nil {
		if instance.InstanceID == "" {
			instance.State = models.CREATE_ERROR
		} else {
			instance.State = models.RUN_ERROR
		}
		models.UpdateInstance(&instance)
		log.Panicln(err)
		panic(err)
	}

	instance.State = models.RUNNING
	models.UpdateInstance(&instance)
}

func StartInstance(instance models.Instance) {
	err := docker.Start(instance.InstanceID)
	if err != nil {
		panic(err)
	}

	instance.State = models.RUNNING
	models.UpdateInstance(&instance)
}

func StopInstance(instance models.Instance) {
	err := docker.Stop(instance.InstanceID)
	if err != nil {
		panic(err)
	}

	instance.State = models.STOPPED
	models.UpdateInstance(&instance)
}

func DeleteInstance(instance models.Instance) {
	err := docker.Delete(instance.InstanceID)
	if err != nil {
		panic(err)
	}

	models.DeleteInstance(&instance)
	os.RemoveAll(config.GetAppLocalPath(instance.Name))
}