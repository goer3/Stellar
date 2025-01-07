package initialize

import (
	"bytes"
	"fmt"
	"os"
	"stellar/common"
	"stellar/pkg/utils"

	"github.com/spf13/viper"
)

// 配置初始化
func Config() {
	var bs []byte
	var err error

	// 初始化 viper，设置配置文件类型
	v := viper.New()
	v.SetConfigType(common.SYSTEM_DEFAULT_CONFIG_FILETYPE)

	// 读取配置文件内容
	if common.SystemConfigFilename != "" {
		// 如果启动参数中有指定配置文件，则使用指定配置文件
		filename := common.SystemConfigFilename
		bs, err = os.ReadFile(filename)
	} else {
		// 否则使用默认配置文件，读取方式不一样
		filename := common.SYSTEM_DEFAULT_CONFIG_FILENAME
		bs, err = common.FS.ReadFile(filename)
		// 设置配置文件，方便后面使用
		common.SystemConfigFilename = filename
	}

	// 读取配置文件失败
	if err != nil {
		panic("读取配置文件失败：" + err.Error())
	}

	// viper 解析配置文件
	err = v.ReadConfig(bytes.NewReader(bs))
	if err != nil {
		panic("解析配置文件失败：" + err.Error())
	}

	// 反序列化配置
	err = v.Unmarshal(&common.Config)
	if err != nil {
		panic("反序列化配置文件失败：" + err.Error())
	}

	// 命令行参数解析覆盖配置文件中的设置
	// 监听地址
	if common.SystemListenHost != "" {
		if !utils.IsIPv4(common.SystemListenHost) {
			panic("监听地址格式错误：" + common.SystemListenHost)
		}
		common.Config.System.Listen.Host = common.SystemListenHost
	}

	// 监听端口
	if common.SystemListenPort != "" {
		if !utils.IsPort(common.SystemListenPort) {
			panic("监听端口格式错误：" + common.SystemListenPort)
		}
		common.Config.System.Listen.Port = common.SystemListenPort
	}

	// WebServer 角色
	if common.SystemRoleWebServer != "" {
		if common.SystemRoleWebServer != "1" && common.SystemRoleWebServer != "0" {
			panic("角色参数 --web-server 设置错误，仅支持 1 或者 0，当前：" + common.SystemRoleWebServer)
		}
		if common.SystemRoleWebServer == "1" {
			common.Config.System.Role.WebServer = true
		} else {
			common.Config.System.Role.WebServer = false
		}
	}

	// LeaderElection 角色
	if common.SystemRoleLeaderElection != "" {
		if common.SystemRoleLeaderElection != "1" && common.SystemRoleLeaderElection != "0" {
			panic("角色参数 --leader-election 设置错误，仅支持 1 或者 0，当前：" + common.SystemRoleLeaderElection)
		}
		if common.SystemRoleLeaderElection == "1" {
			common.Config.System.Role.LeaderElection = true
		} else {
			common.Config.System.Role.LeaderElection = false
		}
	}

	// Worker 角色
	if common.SystemRoleWorker != "" {
		if common.SystemRoleWorker != "1" && common.SystemRoleWorker != "0" {
			panic("角色参数 --worker 设置错误，仅支持 1 或者 0，当前：" + common.SystemRoleWorker)
		}
		if common.SystemRoleWorker == "1" {
			common.Config.System.Role.Worker = true
		} else {
			common.Config.System.Role.Worker = false
		}
	}

	// 如果所有角色都是 False，则退出
	if !common.Config.System.Role.WebServer && !common.Config.System.Role.LeaderElection && !common.Config.System.Role.Worker {
		// 正常退出，不使用 panic
		fmt.Println("所有角色都未启用，无法启动服务，服务即将退出...")
		os.Exit(0)
	}
}
