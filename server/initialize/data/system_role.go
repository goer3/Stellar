package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 角色数据
var systemRoles = []model.SystemRole{
	{
		BaseModel:   model.BaseModel{Id: 1},
		Name:        "超级管理员",
		Keyword:     "SuperAdministrator",
		Description: "系统预设的最高管理员角色，拥有系统所有且最高权限，无需授权，普通用户无法加入该角色",
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel:   model.BaseModel{Id: 2},
		Name:        "管理员",
		Keyword:     "Administrator",
		Description: "系统预设的管理员角色，拥有系统基础的管理权限，拥有用户级管理权限，权限预设且无法调整",
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel:   model.BaseModel{Id: 3},
		Name:        "运维",
		Keyword:     "DevOps",
		Description: "系统初始化角色，拥有系统基础的运维权限，管理员及以上的角色可对其权限进行调整",
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel:   model.BaseModel{Id: 4},
		Name:        "访客",
		Keyword:     "Guest",
		Description: "系统初始化角色，拥有系统的只读权限，只能查看系统的基础信息，无法进行任何修改操作",
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
	},
}

// 角色数据迁移
func MigrateSystemRoleData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行角色数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_role")
	common.DB.Exec("TRUNCATE TABLE system_role_api_relation")
	common.DB.Exec("TRUNCATE TABLE system_role_menu_relation")
	common.DB.Create(&systemRoles)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t角色数据迁移完成")
}