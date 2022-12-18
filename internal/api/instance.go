package api

import (
	"encoding/json"
	"log"
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
