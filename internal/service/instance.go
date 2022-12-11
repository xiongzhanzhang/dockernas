package service

import (
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/models"
)

func CreateInstance(param models.InstanceParam) {
	docker.Create(param)
}
