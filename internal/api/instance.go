package api

import (
	"encoding/json"
	"log"
	"strconv"
	"tinycloud/internal/models"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func PostInstance(c *gin.Context) {
	var param models.InstanceParam
	c.BindJSON(&param)

	service.CreateInstance(param)

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func GetInstance(c *gin.Context) {
	instances := models.GetInstance()
	c.JSON(200, gin.H{
		"list": instances,
	})
}

func GetInstanceByName(c *gin.Context) {
	name := c.Param("name")
	instance := models.GetInstanceByName(name)
	c.JSON(200, instance)
}

func PatchInstance(c *gin.Context) {
	name := c.Param("name")
	instance := models.GetInstanceByName(name)

	patchMap := map[string]string{}
	c.BindJSON(&patchMap)

	op, _ := patchMap["op"]
	if op == "stop" {
		service.StopInstance(instance)
	}
	if op == "start" {
		service.StartInstance(instance)
	}
	if op == "edit" {
		data, _ := patchMap["data"]
		var param models.InstanceParam
		err := json.Unmarshal([]byte(data), &param)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		service.EditInstance(instance, param)
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func DeleteInstance(c *gin.Context) {
	name := c.Param("name")
	instance := models.GetInstanceByName(name)
	service.DeleteInstance(instance)
	c.JSON(200, gin.H{"msg": "ok"})
}

func GetInstanceLog(c *gin.Context) {
	name := c.Param("name")
	instance := models.GetInstanceByName(name)

	log := ""
	if instance.ContainerID != "" && instance.State == models.RUNNING {
		log = service.GetInstanceLog(instance)
	}

	c.JSON(200, gin.H{"data": log})
}

func GetInstanceEvent(c *gin.Context) {
	name := c.Param("name")
	instance := models.GetInstanceByName(name)
	events := models.GetEvents(instance.Id)
	c.JSON(200, gin.H{"list": events})
}

func GetAllInstanceStats(c *gin.Context) {
	start, err1 := strconv.ParseInt(c.Query("start"), 10, 64)
	end, err2 := strconv.ParseInt(c.Query("end"), 10, 64)

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	c.JSON(200, gin.H{"list": models.GetStatDataByTime(start, end)})
}

func GetInstanceStatsByName(c *gin.Context) {
	start, err1 := strconv.ParseInt(c.Query("start"), 10, 64)
	end, err2 := strconv.ParseInt(c.Query("end"), 10, 64)
	name := c.Param("name")

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	c.JSON(200, gin.H{"list": models.GetStatDataByName(name, start, end)})
}
