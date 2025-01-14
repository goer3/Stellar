package common

import (
	"embed"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局工具
var (
	FS        embed.FS           // 打包的静态资源
	Config    *Configuration     // 配置文件解析
	SystemLog *zap.SugaredLogger // 系统日志工具
	AccessLog *zap.SugaredLogger // 访问日志工具
	DB        *gorm.DB           // 数据库连接
	Cache     *redis.Client      // Redis 连接
)

// 全局变量，根据实际运行情况进行初始化
var (
	SystemVersion                  string  = ""                 // 系统后端版本
	SystemConfigFilename           string  = ""                 // 系统运行配置
	SystemListenHost               string  = ""                 // 监听地址
	SystemListenPort               string  = ""                 // 监听端口
	SystemRoleWebServer            string  = ""                 // 是否是 Web 后端服务角色，可选值：1、0
	SystemRoleLeaderElection       string  = ""                 // 是否是领导者竞选角色，可选值：1、0
	SystemRoleWorker               string  = ""                 // 是否是工作者角色，可选值：1、0
	SystemNodeId                   *string                      // 系统运行生成的 ID，节点唯一标识
	SystemNodeStartTime            *string                      // 系统启动时间
	SystemSuperAdministratorRoleId uint    = 1                  // 系统超级管理员角色 ID
	SystemAdministratorRoleIds     []uint  = []uint{1, 2}       // 系统管理员角色 ID
	SystemProtectedRoleIds         []uint  = []uint{1, 2, 3, 4} // 系统预设的角色 ID，不允许修改和删除
)

// 全局固定配置
const (
	SYSTEM_NAME                    string = "Stellar"                                      // 系统名称
	SYSTEM_DESCRIPTION             string = "Stellar 是一个集成了系统和业务监控，支持多数据源、告警源、告警通道的运维监控系统" // 系统描述
	SYSTEM_DEFAULT_CONFIG_FILENAME string = "stellar.yaml"                                 // 默认配置文件
	SYSTEM_DEFAULT_CONFIG_FILETYPE string = "yaml"                                         // 默认配置文件类型
	SYSTEM_GO_VERSION              string = "1.23.0"                                       // 系统 Go 版本
	SYSTEM_API_PREFIX              string = "/api/v1"                                      // 系统 Api 前缀
	SYSTEM_DEVELOPER_NAME          string = "DK"                                           // 系统开发者名称
	SYSTEM_DEVELOPER_EMAIL         string = "ezops.cn@gmail.com"                           // 系统开发者邮箱
	SYSTEM_GITHUB_REPOSITORY       string = "https://github.com/goer3/Stellar"             // 系统 GitHub 仓库
	SYSTEM_DEFAULT_PASSWORD        string = "12345678"                                     // 系统默认密码
	SYSTEM_DEFAULT_AVATAR          string = "/image/avatar/default.png"                    // 系统默认头像
	SYSTEM_DEFAULT_CREATOR         string = "default,默认,Default,0"                         // 系统默认创建者
)

// 格式常量
const (
	TIME_FORMAT_MILLISECOND string = "2006-01-02 15:04:05.000"    // 毫秒时间格式化
	TIME_FORMAT_SECOND      string = "2006-01-02 15:04:05"        // 秒时间格式化
	TIME_FORMAT_DATE        string = "2006-01-02"                 // 日期时间格式化
	UPPERCASE_LETTERS       string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 大写字母
	LOWERCASE_LETTERS       string = "abcdefghijklmnopqrstuvwxyz" // 小写字母
	NUMBERS                 string = "0123456789"                 // 数字
	GENDER_UNKNOWN          uint   = 0                            // 性别未知
	GENDER_MALE             uint   = 1                            // 性别男
	GENDER_FEMALE           uint   = 2                            // 性别女
	STATUS_NOTOK            uint   = 0                            // 状态禁用
	STATUS_OK               uint   = 1                            // 状态正常
	BOOLEAN_FALSE           uint   = 0                            // 布尔值 false
	BOOLEAN_TRUE            uint   = 1                            // 布尔值 true
)
