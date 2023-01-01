package utils

import (
	"net"
	"time"
)

func IsPortUsed(protocol string, port string) bool {
	if protocol != "udp" {
		protocol = "tcp"
	}
	conn, err := net.DialTimeout(protocol, net.JoinHostPort("localhost", port), time.Second)
	if conn != nil {
		conn.Close()
	}
	if err == nil {
		return true
	}
	return false
}
