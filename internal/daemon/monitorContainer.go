package daemon

import (
	"dockernas/internal/backend/docker"
	"dockernas/internal/config"
	"dockernas/internal/models"
	"time"
)

var historyStatMap map[string]models.ContainerStat = map[string]models.ContainerStat{}
var newStatMap map[string]models.ContainerStat = map[string]models.ContainerStat{}

func computeSpeed(stat models.ContainerStat, curTime int64) *models.ContainerStat {
	oldStat, ok := historyStatMap[stat.ContainerID]
	historyStatMap[stat.ContainerID] = stat
	newStatMap[stat.ContainerID] = stat
	if !ok {
		return nil
	}

	newStat := stat
	timeGap := float64(stat.CreateTime-oldStat.CreateTime) / 1000
	newStat.NetworkRx = (stat.NetworkRx - oldStat.NetworkRx) / timeGap
	newStat.NetworkTx = (stat.NetworkTx - oldStat.NetworkTx) / timeGap
	newStat.BlockRead = (stat.BlockRead - oldStat.BlockRead) / timeGap
	newStat.BlockWrite = (stat.BlockWrite - oldStat.BlockWrite) / timeGap
	newStat.CreateTime = curTime //reset time make chart on frontend looks good

	if newStat.NetworkRx < 0 || newStat.NetworkTx < 0 || newStat.BlockRead < 0 || newStat.BlockWrite < 0 {
		return nil
	}

	return &newStat
}

func monitorContainer() {
	if !config.IsBasePathSet() {
		return
	}

	newStatMap = map[string]models.ContainerStat{}

	hostStat := computeSpeed(GetHostState(), time.Now().UnixMilli())
	if hostStat != nil {
		models.AddContainerStat([]models.ContainerStat{*hostStat})
	}
	models.DelStatDataByTime(time.Now().UnixMilli() - 7*24*60*60*1000)

	var statsBySpeed []models.ContainerStat
	containerStats := docker.GetContainerStatus()
	curTime := time.Now().UnixMilli()

	for _, stat := range containerStats {
		newStat := computeSpeed(stat, curTime)
		if newStat != nil {
			statsBySpeed = append(statsBySpeed, *newStat)
		}
	}

	historyStatMap = newStatMap
	models.AddContainerStat(statsBySpeed)
	updateInstanceState()
}

func updateInstanceState() {
	instances := models.GetInstance()
	for _, instance := range instances {
		if _, ok := newStatMap[instance.ContainerID]; !ok {
			if instance.State == models.RUNNING {
				_, err := docker.GetContainerStat(instance.ContainerID)
				if err != nil {
					instance.State = models.STOPPED
					models.UpdateInstance(&instance)
				}
			}
		} else {
			if instance.State == models.STOPPED {
				instance.State = models.RUNNING
				models.UpdateInstance(&instance)
			}
		}
	}
}
