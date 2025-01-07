package utils

import (
	"net"
	"strconv"
)

// 验证是否是 IPv4 地址
func IsIPv4(ip string) bool {
	return net.ParseIP(ip) != nil
}

// 验证是否是端口
func IsPort(port string) bool {
	// 将字符串转换为整数
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	// 检查端口号范围是否在 0-65535 之间
	return portNum >= 0 && portNum <= 65535
}
