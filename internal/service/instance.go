package service

import (
	"bufio"
	"dockernas/internal/backend/docker"
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"
	"time"
)

func checkParamIsValid(param models.InstanceParam) {
	match, _ := regexp.MatchString("[a-zA-Z0-9][a-zA-Z0-9_.-]", param.Name)
	if match == false {
		panic(param.Name + " not match [a-zA-Z0-9][a-zA-Z0-9_.-] to be a container name")
	}

	for _, item := range append(param.EnvParams, param.OtherParams...) {
		if item.Reg != "" {
			match, err := regexp.MatchString(item.Reg, item.Value)
			if err != nil {
				panic("regexp check faild with " + item.Reg + " " + item.Value + ": " + err.Error())
			}
			if match == false {
				panic(item.Value + " is not match " + item.Reg + " on param " + item.Prompt)
			}
		}
	}
}

func GetInstance() []models.Instance {
	instances := models.GetInstance()

	networkInfo := GetNetworkInfo()
	if networkInfo.HttpGatewayEnable {
		for i, instance := range instances {
			proxyConfig := models.GetHttpProxyConfigByInstance(instance.Name, strconv.Itoa(instance.Port))
			if proxyConfig != nil {
				if networkInfo.HttpsEnable {
					instances[i].Url = "https://" + proxyConfig.HostName + "." + networkInfo.Domain
				} else {
					instances[i].Url = "http://" + proxyConfig.HostName + "." + networkInfo.Domain
				}
			}
		}
	}
	return instances
}

func runNewContainer(instance *models.Instance, param models.InstanceParam) {
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
		models.UpdateInstance(instance)
		log.Println(err)
		models.AddEventLog(instance.Id, models.START_EVENT, err.Error())
		panic(err)
	} else {
		models.AddEventLog(instance.Id, models.START_EVENT, "")
	}

	instance.State = models.RUNNING
	models.UpdateInstance(instance)
	SavePortUsed(instance, param)
}

func pullAndRunContainer(instance *models.Instance, param models.InstanceParam, blocking bool) *models.Instance {
	docker.GetBasePathOnHost() //check base path
	log.Println("pull image " + param.ImageUrl)
	reader := docker.PullImage(param.ImageUrl) //if pull image error, break exec here

	instance.State = models.PULL_IMAGE
	models.UpdateInstance(instance)

	if blocking {
		io.Copy(log.Writer(), reader)
		runNewContainer(instance, param)
	} else {
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Println("create instance:", err)
					log.Println(string(debug.Stack()))
				}
				reader.Close()
			}()

			startTime := time.Now().Unix()
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				line := scanner.Text()
				ProcessImagePullMsg(param.ImageUrl, line)
				if (time.Now().Unix() - startTime) >= (60 * 30) { // timeout for 30 minute
					log.Println("pull image " + param.ImageUrl + " time out")
					ReportImagePullStoped(param.ImageUrl)
					instance.State = models.PULL_ERROR
					models.UpdateInstance(instance)
					return
				}
			}
			log.Println("pull image " + param.ImageUrl + " ok")
			ReportImagePullStoped(param.ImageUrl)
			tmp := models.GetInstanceByName(instance.Name)
			if tmp == nil || tmp.Id != instance.Id { //check if instance is deleted
				return
			}
			runNewContainer(instance, param)
		}()
	}

	return instance
}

func GetInstanceByName(name string) models.Instance {
	instance := models.GetInstanceByName(name)
	if instance == nil {
		panic("instance " + name + " not exists")
	}
	if instance.State == models.PULL_IMAGE {
		var param models.InstanceParam
		if utils.GetObjFromJson(instance.InstanceParamStr, &param) != nil {
			instance.ImagePullState = GetImagePullState(param.ImageUrl)
		}
	}
	return *instance
}

func CreateInstance(param models.InstanceParam, blocking bool) *models.Instance {
	checkParamIsValid(param)
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

	pullAndRunContainer(&instance, param, blocking)

	return &instance
}

func EditInstance(instance models.Instance, param models.InstanceParam) {
	checkParamIsValid(param)
	DelInstancePorts(instance)
	CheckIsPortUsed(param)
	log.Println("delete comtainer of instance " + instance.Name)
	err := docker.Delete(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.CONFIG_EVENT, err.Error())
		panic(err)
	} else {
		models.AddEventLog(instance.Id, models.CONFIG_EVENT, "")
	}

	instance.Summary = param.Summary
	instance.State = models.PULL_IMAGE
	instance.AppName = param.AppName
	instance.Version = param.Version
	instance.IconUrl = param.IconUrl
	instance.Port = getFirstHttpPort(param)
	instance.InstanceParamStr = utils.GetJsonFromObj(param)

	models.UpdateInstance(&instance)

	pullAndRunContainer(&instance, param, false)
	models.AddEventLog(instance.Id, models.CONFIG_EVENT, "")
}

func RestartInstance(instance models.Instance) {
	err := docker.Restart(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.RESTART_EVENT, err.Error())
	} else {
		models.AddEventLog(instance.Id, models.RESTART_EVENT, "")
	}
}

func StartInstance(instance models.Instance) {
	if instance.ContainerID == "" {
		var param models.InstanceParam
		err := json.Unmarshal([]byte(instance.InstanceParamStr), &param)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		pullAndRunContainer(&instance, param, false)

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
	// if instance.State == models.RUNNING {
	// 	StopInstance(instance)
	// }

	DelInstancePorts(instance)
	models.DelInstanceStatData(instance.Name)
	models.DelEvents(instance.Id)

	log.Println("delete container of instance " + instance.Name)
	err := docker.Delete(instance.ContainerID)
	if err != nil {
		models.AddEventLog(instance.Id, models.DELETE_EVENT, err.Error())
		panic(err)
	} else {
		models.DeleteInstance(&instance)
		os.RemoveAll(config.GetAppLocalPath(instance.Name))
		// models.AddEventLog(instance.Id, models.DELETE_EVENT, "")
	}
}

func GetInstanceLog(instance models.Instance) string {
	return docker.GetLog(instance.ContainerID)
}
