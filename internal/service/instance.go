package service

import (
	"bufio"
	"encoding/json"
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

	log.Println("create instance " + instance.Name)
	instance.ContainerID, err = docker.Create(&param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	if err != nil {
		if instance.ContainerID == "" {
			instance.State = models.CREATE_ERROR
		} else {
			instance.State = models.RUN_ERROR
		}
		models.UpdateInstance(&instance)
		log.Println(err)
		models.AddEventLog(instance.Id, models.START_EVENT, err.Error())
		panic(err)
	}

	instance.State = models.RUNNING
	models.UpdateInstance(&instance)
	SavePortUsed(instance, param)
}

func CreateInstance(param models.InstanceParam) *models.Instance {
	log.Println("pull image " + param.ImageUrl)
	reader := docker.PullImage(param.ImageUrl) //if pull image error, break exec here
	CheckIsPortUsed(param)

	var instance models.Instance

	instance.Name = param.Name
	instance.Summary = param.Summary
	instance.State = models.PULL_IMAGE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.Port = getFirstHttpPort(param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)
	instance.CreateTime = time.Now().UnixMilli()

	models.AddInstance(&instance)

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println("create instance:", err)
			}
			reader.Close()
		}()

		startTime := time.Now().Unix()
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			log.Println(line)
			if (time.Now().Unix() - startTime) >= (60 * 30) { // timeout for 30 minute
				log.Println("pull image " + param.ImageUrl + " time out")
				instance.State = models.PULL_ERROR
				models.UpdateInstance(&instance)
				return
			}
		}
		log.Println("pull image " + param.ImageUrl + " ok")
		tmp := models.GetInstanceByName(instance.Name)
		if tmp == nil || tmp.Id != instance.Id { //check if instance is deleted
			return
		}
		runNewContainer(instance, param)
	}()

	return &instance
}

func EditInstance(instance models.Instance, param models.InstanceParam) {
	DelInstancePorts(instance)
	CheckIsPortUsed(param)
	log.Println("delete comtainer of instance " + instance.Name)
	err := docker.Delete(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.CONFIG_EVENT, err.Error())
		panic(err)
	}

	instance.Summary = param.Summary
	instance.State = models.PULL_IMAGE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.Port = getFirstHttpPort(param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	models.UpdateInstance(&instance)

	runNewContainer(instance, param)
	models.AddEventLog(instance.Id, models.CONFIG_EVENT, "")
}

func RestartInstance(instance models.Instance) {
	StopInstance(instance)
	StartInstance(instance)
}

func StartInstance(instance models.Instance) {
	if instance.ContainerID == "" {
		var param models.InstanceParam
		err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		runNewContainer(instance, param)
	} else {
		log.Println("start comtainer of instance " + instance.Name)
		err := docker.Start(instance.ContainerID)
		if err != nil {
			models.AddEventLog(instance.Id, models.START_EVENT, err.Error())
			panic(err)
		}
		instance.State = models.RUNNING
		models.UpdateInstance(&instance)
		models.AddEventLog(instance.Id, models.START_EVENT, "")
	}
}

func StopInstance(instance models.Instance) {
	log.Println("stop comtainer of instance " + instance.Name)
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
	if instance.State == models.RUNNING {
		StopInstance(instance)
	}

	DelInstancePorts(instance)
	models.DelInstanceStatData(instance.Name)

	log.Println("delete comtainer of instance " + instance.Name)
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
