package common

import (
	"embed"
)

// 全局工具
var (
	FS     embed.FS       // 打包的静态资源
	Config *Configuration // 配置文件解析
)

// 全局变量，根据实际运行情况进行初始化
var (
	SystemVersion            = ""                                              // 系统后端版本
	SystemConfigFilename     = ""                                              // 系统运行配置
	SystemListenHost         = ""                                              // 监听地址
	SystemListenPort         = ""                                              // 监听端口
	SystemRoleWebServer      = ""                                              // 是否是 Web 后端服务角色，可选值：1、0
	SystemRoleLeaderElection = ""                                              // 是否是领导者竞选角色，可选值：1、0
	SystemRoleWorker         = ""                                              // 是否是工作者角色，可选值：1、0
	SystemAdminRoleKeywords  = []string{"SuperAdministrator", "Administrator"} // 系统管理员角色关键字
	SystemNodeId             *string                                           // 系统运行生成的 ID，节点唯一标识
	SystemNodeStartTime      *string                                           // 系统启动时间
)

// 全局固定配置
const (
	SYSTEM_NAME                    = "Stellar"                                                    // 系统名称
	SYSTEM_DESCRIPTION             = "Stellar 是一个集成了系统监控和业务监控，支持多数据源、多告警源、多告警通知方式、多告警处理方式的运维监控系统" // 系统描述
	SYSTEM_DEFAULT_CONFIG_FILENAME = "stellar.yaml"                                               // 默认配置文件
	SYSTEM_DEFAULT_CONFIG_FILETYPE = "yaml"                                                       // 默认配置文件类型
	SYSTEM_GO_VERSION              = "1.23.0"                                                     // 系统 Go 版本
	SYSTEM_API_PREFIX              = "/api/v1"                                                    // 系统 API 前缀
	SYSTEM_DEVELOPER_NAME          = "DK"                                                         // 系统开发者名称
	SYSTEM_DEVELOPER_EMAIL         = "ezops.cn@gmail.com"                                         // 系统开发者邮箱
	SYSTEM_GITHUB_REPOSITORY       = "https://github.com/goer3/Stellar"                           // 系统 GitHub 仓库
)

// 格式常量
const (
	TIME_FORMAT_MILLISECOND = "2006-01-02 15:04:05.000"    // 毫秒时间格式化
	TIME_FORMAT_SECOND      = "2006-01-02 15:04:05"        // 秒时间格式化
	TIME_FORMAT_DATE        = "2006-01-02"                 // 日期时间格式化
	UPPERCASE_LETTERS       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 大写字母
	LOWERCASE_LETTERS       = "abcdefghijklmnopqrstuvwxyz" // 小写字母
	NUMBERS                 = "0123456789"                 // 数字
)
