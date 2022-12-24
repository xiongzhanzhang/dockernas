package service

import (
	"tinycloud/internal/models"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetHostInfo() models.HostInfo {
	var hostData models.HostInfo

	hostInfo, err := host.Info()
	if err != nil {
		panic(err)
	}
	cpuData, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	memData, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	hostData.HostName = hostInfo.Hostname
	hostData.BootTime = hostInfo.BootTime * 1000
	hostData.Platform = hostInfo.Platform
	hostData.ModelName = cpuData[0].ModelName
	hostData.MemSize = memData.Total

	return hostData
}
