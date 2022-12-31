package daemon

import (
	"log"
	"time"
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
)

var historyStatMap map[string]models.ContainerStat = map[string]models.ContainerStat{}

func monitorContainer() {
	if !config.IsBasePathSet() {
		return
	}

	var statsBySpeed []models.ContainerStat
	var newStatMap map[string]models.ContainerStat = map[string]models.ContainerStat{}
	containerStats := docker.GetContainerStatus()
	curTime := time.Now().UnixMilli()

	for _, stat := range containerStats {
		newStatMap[stat.ContainerID] = stat
		oldStat, ok := historyStatMap[stat.ContainerID]
		if !ok {
			continue
		}

		//compute io speed
		newStat := stat
		timeGap := float64(stat.CreateTime-oldStat.CreateTime) / 1000
		newStat.NetworkRx = (stat.NetworkRx - oldStat.NetworkRx) / timeGap
		newStat.NetworkTx = (stat.NetworkTx - oldStat.NetworkTx) / timeGap
		newStat.BlockRead = (stat.BlockRead - oldStat.BlockRead) / timeGap
		newStat.BlockWrite = (stat.BlockWrite - oldStat.BlockWrite) / timeGap
		newStat.CreateTime = curTime //reset time make chart on frontend looks good

		if newStat.NetworkRx < 0 || newStat.NetworkTx < 0 || newStat.BlockRead < 0 || newStat.BlockWrite < 0 {
			continue
		}

		log.Println(newStat)
		statsBySpeed = append(statsBySpeed, newStat)
	}

	historyStatMap = newStatMap
	models.AddContainerStat(statsBySpeed)
	models.DelStatDataByTime(time.Now().UnixMilli() - 7*24*60*60*1000)
}
