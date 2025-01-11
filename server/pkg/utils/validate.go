package utils

import (
	"net"
	"regexp"
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

// 验证用户名
func IsUsername(username string) bool {
	regex := `^[a-z][a-z0-9]{1,30}$`
	return regexp.MustCompile(regex).MatchString(username)
}

// 验证是否是邮箱
func IsEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

// 验证是否是手机号
func IsPhone(phone string) bool {
	regex := `1[3-9]\d{9}$`
	return regexp.MustCompile(regex).MatchString(phone)
}
