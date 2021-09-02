package utils

import "net"

func IsIPv4(s string) bool {
	ip := net.ParseIP(s)
	return len(ip.To4()) == net.IPv4len
}
