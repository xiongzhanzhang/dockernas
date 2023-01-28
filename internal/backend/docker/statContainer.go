package docker

import (
	"context"
	"dockernas/internal/config"
	"dockernas/internal/models"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerStatus() []models.ContainerStat {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("create docker client error")
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var stats []models.ContainerStat
	for _, container := range containers {
		if len(container.Names) > 0 && container.Names[0] == "/"+config.GetHostNameInStats() {
			continue
		}
		stat := collect(ctx, container, cli)
		if stat != nil {
			stats = append(stats, *stat)
		}
	}

	return stats
}

func collect(ctx context.Context, container types.Container, cli client.APIClient) *models.ContainerStat {
	var (
		previousCPU    uint64
		previousSystem uint64
	)

	response, err := cli.ContainerStats(ctx, container.ID, false)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	var (
		v                 *types.StatsJSON
		cpuPercent        = 0.0
		blkRead, blkWrite uint64 // Only used on Linux
		mem               = 0.0
	)

	if err := dec.Decode(&v); err != nil {
		log.Println(err)
		return nil
	}

	if response.OSType != "windows" {
		previousCPU = v.PreCPUStats.CPUUsage.TotalUsage
		previousSystem = v.PreCPUStats.SystemUsage
		cpuPercent = calculateCPUPercentUnix(previousCPU, previousSystem, v)
		blkRead, blkWrite = calculateBlockIO(v.BlkioStats)
		mem = float64(v.MemoryStats.Usage)
	} else {
		cpuPercent = calculateCPUPercentWindows(v)
		blkRead = v.StorageStats.ReadSizeBytes
		blkWrite = v.StorageStats.WriteSizeBytes
		mem = float64(v.MemoryStats.PrivateWorkingSet)
	}
	netRx, netTx := calculateNetwork(v.Networks)

	return &models.ContainerStat{
		Name:          strings.TrimPrefix(v.Name, "/"),
		ContainerID:   v.ID,
		CPUPercentage: cpuPercent,
		Memory:        mem,
		NetworkRx:     netRx,
		NetworkTx:     netTx,
		BlockRead:     float64(blkRead),
		BlockWrite:    float64(blkWrite),
		CreateTime:    time.Now().UnixMilli(),
	}
}

func calculateCPUPercentUnix(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(v.CPUStats.SystemUsage) - float64(previousSystem)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(v.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}

func calculateCPUPercentWindows(v *types.StatsJSON) float64 {
	// Max number of 100ns intervals between the previous time read and now
	possIntervals := uint64(v.Read.Sub(v.PreRead).Nanoseconds()) // Start with number of ns intervals
	possIntervals /= 100                                         // Convert to number of 100ns intervals
	possIntervals *= uint64(v.NumProcs)                          // Multiple by the number of processors

	// Intervals used
	intervalsUsed := v.CPUStats.CPUUsage.TotalUsage - v.PreCPUStats.CPUUsage.TotalUsage

	// Percentage avoiding divide-by-zero
	if possIntervals > 0 {
		return float64(intervalsUsed) / float64(possIntervals) * 100.0
	}
	return 0.00
}

func calculateBlockIO(blkio types.BlkioStats) (blkRead uint64, blkWrite uint64) {
	for _, bioEntry := range blkio.IoServiceBytesRecursive {
		switch strings.ToLower(bioEntry.Op) {
		case "read":
			blkRead = blkRead + bioEntry.Value
		case "write":
			blkWrite = blkWrite + bioEntry.Value
		}
	}
	return
}

func calculateNetwork(network map[string]types.NetworkStats) (float64, float64) {
	var rx, tx float64

	for _, v := range network {
		rx += float64(v.RxBytes)
		tx += float64(v.TxBytes)
	}
	return rx, tx
}
