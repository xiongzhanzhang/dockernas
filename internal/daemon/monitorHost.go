package daemon

import (
	"time"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func GetHostState() models.ContainerStat {
	var stat models.ContainerStat
	stat.Name = config.GetHostNameInStats()
	stat.CPUPercentage = utils.GetHostCPUPrecent()
	stat.Memory = utils.GetHostMemUsed()
	stat.BlockRead, stat.BlockWrite = utils.GetHostDiskIo()
	stat.NetworkRx, stat.NetworkTx = utils.GetNetIo()
	stat.CreateTime = time.Now().UnixMilli()

	return stat
}
