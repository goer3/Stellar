package initialize

import (
	"fmt"
	"net"
	"stellar/common"
	"stellar/pkg/trans"
)

// 生成节点唯一标识ID，格式为：系统名称-系统监听地址-系统监听端口-系统启动时间
func NodeId() {
	var ipAddress net.IP
	var ipString string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取网络接口失败，将使用默认地址 127.0.0.1")
		ipString = "127.0.0.1"
	} else {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipv4 := ipNet.IP.To4(); ipv4 != nil {
					ipAddress = ipv4
					break
				}
			}
		}
		if ipAddress == nil {
			fmt.Println("未找到有效的 IP 地址，将使用默认地址 127.0.0.1")
			ipString = "127.0.0.1"
		} else {
			ipString = ipAddress.String()
		}
	}

	// 拼接节点唯一标识ID
	systemNodeId := fmt.Sprintf("%s-%s-%s-%s", common.SYSTEM_NAME, ipString, common.Config.System.Listen.Port, *common.SystemNodeStartTime)
	common.SystemNodeId = trans.String(systemNodeId)
}
