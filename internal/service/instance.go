package service

import (
	"log"
	"os"
	"time"
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func runNewContainer(instance models.Instance, param models.InstanceParam) {
	var err error

	instance.ContainerID, err = docker.Create(&param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	if err != nil {
		if instance.ContainerID == "" {
			instance.State = models.CREATE_ERROR
		} else {
			instance.State = models.RUN_ERROR
		}
		models.UpdateInstance(&instance)
		log.Panicln(err)
		models.AddEventLog(instance.Id, models.START_EVENT, err.Error())
		panic(err)
	}

	instance.State = models.RUNNING
	models.UpdateInstance(&instance)
}

func CreateInstance(param models.InstanceParam) {
	var instance models.Instance

	instance.Name = param.Name
	instance.Summary = param.Summary
	instance.State = models.NEW_STATE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.InstanceParamStr = utils.GetJsonFromObj(param)
	instance.CreateTime = time.Now().UnixMilli()

	models.AddInstance(&instance)

	runNewContainer(instance, param)
}

func EditInstance(instance models.Instance, param models.InstanceParam) {
	err := docker.Delete(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.CONFIG_EVENT, err.Error())
		panic(err)
	}

	instance.Summary = param.Summary
	instance.State = models.NEW_STATE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	models.UpdateInstance(&instance)

	runNewContainer(instance, param)
	models.AddEventLog(instance.Id, models.CONFIG_EVENT, "")
}

func StartInstance(instance models.Instance) {
	err := docker.Start(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.START_EVENT, err.Error())
		panic(err)
	}

	instance.State = models.RUNNING
	models.UpdateInstance(&instance)
	models.AddEventLog(instance.Id, models.START_EVENT, "")
}

func StopInstance(instance models.Instance) {
	err := docker.Stop(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.STOP_EVENT, err.Error())
		panic(err)
	}

	instance.State = models.STOPPED
	models.UpdateInstance(&instance)
	models.AddEventLog(instance.Id, models.STOP_EVENT, "")
}

func DeleteInstance(instance models.Instance) {
	err := docker.Delete(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.DELETE_EVENT, err.Error())
		log.Println(err)
	}

	models.DeleteInstance(&instance)
	os.RemoveAll(config.GetAppLocalPath(instance.Name))
	models.AddEventLog(instance.Id, models.DELETE_EVENT, "")
}

func GetInstanceLog(instance models.Instance) string {
	return docker.GetLog(instance.ContainerID)
}
