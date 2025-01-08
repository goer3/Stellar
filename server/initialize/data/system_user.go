package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
	"stellar/pkg/utils"
)

// 用户数据
var systemUsers = []model.SystemUser{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		JobNumber: "S00000001",
		Username:  "super",
		CNName:    "超级管理员",
		ENName:    "Super",
		Password:  utils.CryptoPassword(common.SYSTEM_DEFAULT_PASSWORD),
		Email:     "super@ezops.cn",
		Phone:     "13100000001",
		HidePhone: trans.Uint(0),
		Gender:    trans.Uint(1),
		Avatar:    utils.RandAvatar(1),
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		JobNumber: "S00000002",
		Username:  "admin",
		CNName:    "管理员",
		ENName:    "Administrator",
		Password:  utils.CryptoPassword(common.SYSTEM_DEFAULT_PASSWORD),
		Email:     "admin@ezops.cn",
		Phone:     "13100000002",
		HidePhone: trans.Uint(0),
		Gender:    trans.Uint(2),
		Avatar:    utils.RandAvatar(2),
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{
			Id: 3,
		},
		JobNumber: "S00000003",
		Username:  "devops",
		CNName:    "运维",
		ENName:    "DevOps",
		Password:  utils.CryptoPassword(common.SYSTEM_DEFAULT_PASSWORD),
		Email:     "devops@ezops.cn",
		Phone:     "13100000003",
		HidePhone: trans.Uint(0),
		Gender:    trans.Uint(1),
		Avatar:    utils.RandAvatar(1),
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{
			Id: 4,
		},
		JobNumber: "S00000004",
		Username:  "guest",
		CNName:    "访客",
		ENName:    "Guest",
		Password:  utils.CryptoPassword(common.SYSTEM_DEFAULT_PASSWORD),
		Email:     "guest@ezops.cn",
		Phone:     "13100000004",
		HidePhone: trans.Uint(0),
		Gender:    trans.Uint(0),
		Avatar:    utils.RandAvatar(0),
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
}

// 用户数据迁移
func MigrateSystemUserData() {
	common.DB.Exec("TRUNCATE TABLE system_user")
	common.DB.Exec("TRUNCATE TABLE system_user_department_relation")
	common.DB.Exec("TRUNCATE TABLE system_user_job_position_relation")
	common.DB.Create(&systemUsers)
}
