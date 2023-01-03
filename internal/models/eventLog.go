package models

import (
	"log"
	"time"
)

const (
	CREATE_EVENT  = 0
	STOP_EVENT    = 1
	START_EVENT   = 2
	CONFIG_EVENT  = 3
	DELETE_EVENT  = 4
	RESTART_EVENT = 5
)

type EventLog struct {
	Id         int64  `json:"id"  gorm:"primary_key;auto_increment"`
	InstanceID int    `json:"instanceID"`
	EventType  int    `json:"eventType"`
	Msg        string `json:"msg" gorm:"type:text"`
	CreateTime int64  `json:"createTime"`
}

func AddEventLog(InstanceID int, EventType int, Msg string) error {
	var event EventLog
	event.InstanceID = InstanceID
	event.EventType = EventType
	event.Msg = Msg
	event.CreateTime = time.Now().UnixMilli()

	err := GetDb().Create(&event).Error
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetEvents(InstanceID int) []EventLog {
	var events []EventLog
	err := GetDb().Where("instance_id = ?", InstanceID).Find(&events).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return events
}

func DelEvents(InstanceID int) {
	err := GetDb().Where("instance_id = ?", InstanceID).Delete(&EventLog{}).Error
	if err != nil {
		log.Println(err)
	}
}
