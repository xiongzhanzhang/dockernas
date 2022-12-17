package service

import (
	"log"
	"tinycloud/internal/backend/docker"
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
	instance.InstanceParamStr = utils.GetJsonFromObj(param)
	models.UpdateInstance(&instance)
}
