package common

import "time"

// Redis 通用
const (
	RK_SEPARATOR = ":" // 分隔符
)

// Redis 键，无需要拼接
const (
	RK_LEADER = "Leader" // 领导节点
)

// Redis 键前缀，需要拼接
const (
	RKP_HEARTBEAT             = "Heartbeat"          // 心跳上报，格式：Heartbeat:ClientId
	RKP_WORKER                = "Worker"             // 工作节点，格式：Worker:ClientId
	RKP_WEB_SERVER            = "WebServer"          // Web 后端服务，格式：WebServer:ClientId
	RKP_LOGIN_TOKEN           = "LoginToken"         // 登录 Token，格式：LoginToken:Username
	RKP_LOGIN_ERROR_TIMES     = "LoginErrorTimes"    // 登录错误次数，格式：LoginErrorTimes:Account:IP
	RKP_SYSTEM_ROLE_API_LIST  = "SystemRoleApiList"  // 系统角色Api列表，[]string 使用 JSON 序列化，格式：SystemRoleApiList:RoleKeyword
	RKP_SYSTEM_ROLE_MENU_LIST = "SystemRoleMenuList" // 系统角色菜单列表，[]model.SystemMenu 使用 GOB 序列化，格式：SystemRoleMenuList:RoleKeyword
)

// Redis 键过期时间
const (
	RKE_HEARTBEAT             = time.Second * 15 // 心跳上报过期时间
	RKE_LEADER                = time.Second * 15 // 领导节点过期时间
	RKE_WORKER                = time.Second * 15 // 工作节点过期时间
	RKE_WEB_SERVER            = time.Second * 15 // Web 后端服务过期时间
	RKE_SYSTEM_ROLE_API_LIST  = time.Second * 60 // 系统角色Api列表过期时间
	RKE_SYSTEM_ROLE_MENU_LIST = time.Second * 60 // 系统角色菜单列表过期时间
)

// 任务间隔时间，不能等于过期时间
const (
	RKE_INTERVAL_HEARTBEAT  = time.Second * 10 // 心跳上报间隔时间
	RKE_INTERVAL_LEADER     = time.Second * 10 // 领导节点选举间隔时间
	RKE_INTERVAL_WORKER     = time.Second * 10 // 工作节点选举间隔时间
	RKE_INTERVAL_WEB_SERVER = time.Second * 10 // Web 后端服务选举间隔时间
)
