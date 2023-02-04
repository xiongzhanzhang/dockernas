package utils

import (
	"log"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
)

func GetOperationSystemName() string {
	return runtime.GOOS
}

func GetHostMemUsed() float64 {
	memData, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	return float64(memData.Used)
}

func GetHostCPUPrecent() float64 {
	cpuData, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		panic(err)
	}

	return cpuData[0]
}

func GetHostDiskIo() (float64, float64) {
	ReadBytes := 0.0
	WriteBytes := 0.0
	mapStat, _ := disk.IOCounters()
	for _, v := range mapStat {
		ReadBytes += float64(v.ReadBytes)
		WriteBytes += float64(v.WriteBytes)
	}

	return ReadBytes, WriteBytes
}

func GetNetIo() (float64, float64) {
	ReadBytes := 0.0
	WriteBytes := 0.0
	mapStat, _ := net2.IOCounters(false)
	for _, v := range mapStat {
		ReadBytes += float64(v.BytesRecv)
		WriteBytes += float64(v.BytesSent)
	}

	return ReadBytes, WriteBytes
}

func IsPortUsed(host string, protocol string, port string) bool {
	if protocol != "tcp" && protocol != "http" {
		return false //only check tcp port
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Millisecond*200)
	if conn != nil {
		conn.Close()
	}
	if err == nil {
		return true
	}
	return false
}

func GetLocalAddress() string {
	con, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Println(error)
		return ""
	}
	defer con.Close()
	localAddress := con.LocalAddr().(*net.UDPAddr)
	return localAddress.IP.String()
}

func GetLocalAddressIpv6() string {
	con, error := net.Dial("udp", "[2001:4860:4860::8888]:80")
	if error != nil {
		log.Println(error)
		return ""
	}
	defer con.Close()
	localAddress := con.LocalAddr().(*net.UDPAddr)
	return localAddress.IP.String()
}

func GetFileSystemTypeByCmd(mountPoint string) string {
	cmd := exec.Command("mount")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		log.Println(line)
		if strings.Contains(line, mountPoint) {
			data := strings.Split(DeleteExtraSpace(line), " ")
			return data[4]
		}
	}

	return ""
}

func GetDeviceSizeByCmd(mountPoint string) (string, string, int64, int64) {
	cmd := exec.Command("df")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		log.Println(line)
		if strings.Contains(line, mountPoint) {
			data := strings.Split(DeleteExtraSpace(line), " ")
			Capacity, err := strconv.ParseInt(data[1], 10, 64)
			if err != nil {
				panic(err)
			}
			FreeSize, err := strconv.ParseInt(data[3], 10, 64)
			if err != nil {
				panic(err)
			}
			return data[0], GetFileSystemTypeByCmd(mountPoint), Capacity * 1024, FreeSize * 1024
		}
	}

	panic("can't find device on mountpoint:" + mountPoint)
}
