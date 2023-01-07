package service

import (
	"log"
	"os"
	"path/filepath"
	"time"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"

	"github.com/go-git/go-git/v5"
)

func AddSubscribe(subscribe models.Subscribe) {
	subscribe.CreateTime = time.Now().UnixMilli()
	subscribe.UpdateTime = time.Now().UnixMilli()
	subscribe.State = models.SUBSCRIBE_INIT

	repositoryPath := getSubscribeRepositoryPath(subscribe.Name)
	if utils.IsFileExist(repositoryPath) {
		panic(subscribe.Name + " is used on app store path:" + repositoryPath)
	}
	os.MkdirAll(repositoryPath, os.ModePerm)
	models.AddSubscribe(&subscribe)

	pullSubscribe(subscribe)
}

func DelSubscribe(name string) {
	models.DeleteSubscribe(models.GetSubscribeByName(name))
	os.RemoveAll(getSubscribeRepositoryPath(name))
}

func UpdateSubscribe() {
	subscribes := models.GetOkSubscribe()
	for _, v := range subscribes {
		v.State = models.SUBSCRIBE_UPDATEING
		models.UpdateSubscribe(&v)
		updateSubscribe(v)
	}
}

func getSubscribeRepositoryPath(name string) string {
	return filepath.Join(config.GetExtraAppPath(), name)
}

func pullSubscribe(subscribe models.Subscribe) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println("pull subscribe error:", err)
			}
		}()
		log.Println("start pull " + subscribe.Url)
		_, err := git.PlainClone(
			getSubscribeRepositoryPath(subscribe.Name),
			false,
			&git.CloneOptions{
				URL:      subscribe.Url,
				Progress: log.Writer(),
			},
		)

		if err == nil {
			subscribe.State = models.SUBSCRIBE_OK
			log.Println("pull ok " + subscribe.Url)
		} else {
			subscribe.State = models.SUBSCRIBE_INITERROR
			log.Println("pull error " + subscribe.Url)
			log.Println(err)
		}
		subscribe.UpdateTime = time.Now().UnixMilli()
		models.UpdateSubscribe(&subscribe)
	}()
}

func updateSubscribe(subscribe models.Subscribe) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println("update subscribe error:", err)
			}
		}()
		log.Println("start update " + subscribe.Url)
		repository, err := git.PlainOpen(getSubscribeRepositoryPath(subscribe.Name))
		if err != nil {
			err = repository.Fetch(&git.FetchOptions{})
		}

		if err == nil {
			subscribe.State = models.SUBSCRIBE_OK
			log.Println("update ok " + subscribe.Url)
		} else {
			subscribe.State = models.SUBSCRIBE_UPDATE_FAIL
			log.Println("update error " + subscribe.Url)
			log.Println(err)
		}
		subscribe.UpdateTime = time.Now().UnixMilli()
		models.UpdateSubscribe(&subscribe)
	}()
}
