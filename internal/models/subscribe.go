package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

const (
	SUBSCRIBE_INIT        = 0
	SUBSCRIBE_INITERROR   = 1
	SUBSCRIBE_OK          = 2
	SUBSCRIBE_UPDATEING   = 3
	SUBSCRIBE_UPDATE_FAIL = 4
)

type Subscribe struct {
	Id         int    `json:"id"  gorm:"primary_key;auto_increment"`
	Name       string `json:"name"  gorm:"unique;not null"`
	Url        string `json:"url"`
	State      int    `json:"state"`
	UpdateTime int64  `json:"updateTime"`
	CreateTime int64  `json:"createTime"`
}

func AddSubscribe(subscribe *Subscribe) {
	err := GetDb().Create(subscribe).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func UpdateSubscribe(subscribe *Subscribe) {
	err := GetDb().Model(&Subscribe{}).
		Where("id = ? and name = ?", subscribe.Id, subscribe.Name).
		Updates(subscribe).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func DeleteSubscribe(subscribe *Subscribe) {
	err := GetDb().Where("id = ? and name = ?", subscribe.Id, subscribe.Name).Delete(subscribe).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func GetSubscribe() []Subscribe {
	var instances []Subscribe
	err := GetDb().Find(&instances).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return instances
}

func GetOkSubscribe() []Subscribe {
	var instances []Subscribe
	err := GetDb().Where("state != 0 and state != 1").Find(&instances).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return instances
}

func GetSubscribeByName(name string) *Subscribe {
	var subscribe Subscribe
	err := GetDb().First(&subscribe, "name=?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Println(err)
		panic(err)
	}

	return &subscribe
}
