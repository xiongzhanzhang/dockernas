package models

import "log"

type ContainerStat struct {
	Name          string  `json:"name"`
	ContainerID   string  `json:"containerID"`
	CPUPercentage float64 `json:"CPUPercentage"`
	Memory        float64 `json:"memory"`
	NetworkRx     float64 `json:"networkRx"`
	NetworkTx     float64 `json:"networkTx"`
	BlockRead     float64 `json:"blockRead"`
	BlockWrite    float64 `json:"blockWrite"`
	CreateTime    int64   `json:"createTime"`
}

func AddContainerStat(stats []ContainerStat) {
	for _, stat := range stats {
		err := GetDb().Create(&stat).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func DelInstanceStatData(instanceName string) {
	err := GetDb().Where("name = ?", instanceName).Delete(&ContainerStat{}).Error
	if err != nil {
		log.Println(err)
	}
}

func DelStatDataByTime(time int64) {
	err := GetDb().Where("create_time <= ?", time).Delete(&ContainerStat{}).Error
	if err != nil {
		log.Println(err)
	}
}

func GetStatDataByTime(start int64, end int64) []ContainerStat {
	var stats []ContainerStat
	err := GetDb().Where("create_time >= ? and create_time <= ? and name!=?", start, end, "physic_host").Find(&stats).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return stats
}

func GetStatDataByName(name string, start int64, end int64) []ContainerStat {
	var stats []ContainerStat
	err := GetDb().Where("create_time >= ? and create_time <= ? and name=?", start, end, name).Find(&stats).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return stats
}
