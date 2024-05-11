package main

import (
	"fmt"
	"math/rand"
	"net"
)

func genIPv4() string {
	// It's known that IPv4 has the length of 32 bits, divided into 4 octets
	ip := make([]byte, 4)
	for i := 0; i < 4; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To4().String()
}

func genIPv6() string {
	// While IPv6 has the length of 128 bits
	// https://www.loganalyzer.net/log-analyzer/ipv6-log-analysis.html
	buf := make([]byte, 16)
	for i := 0; i < 16; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return net.IP(buf).To16().String()
}

func processFileSize(filesize int64) string {
	// 1048576 = 1mb
	if filesize < 1048576 {
		sizeInKB := float64(filesize) / 1024
		return fmt.Sprintf("%.2f KB", sizeInKB)
	} else {
		sizeInMB := float64(filesize) / 1024 / 1024
		return fmt.Sprintf("%.2f MB", sizeInMB)
	}
}
