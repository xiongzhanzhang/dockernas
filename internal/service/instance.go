package service

import (
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func CreateInstance(param models.InstanceParam) {
	var instance models.Instance
	instance.Summary = param.Summary
	instance.Name = param.Name
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	models.AddInstance(&instance)

	instance.InstanceID = docker.Create(param)

	models.UpdateInstance(&instance)
}
