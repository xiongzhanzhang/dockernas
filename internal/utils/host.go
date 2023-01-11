package utils

import (
	"log"
	"net"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
)

func GetOperationSystemName()string{
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

func IsPortUsed(protocol string, port string) bool {
	if protocol != "udp" {
		protocol = "tcp"
	}
	conn, err := net.DialTimeout(protocol, net.JoinHostPort("localhost", port), time.Millisecond*200)
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
