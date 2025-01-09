package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 职位数据
var systemJobPositions = []model.SystemJobPosition{
	{
		BaseModel: model.BaseModel{Id: 1},
		Name:      "首席执行官（CEO）",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[0],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 2},
		Name:      "产品总监（CPO）",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[0],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 3},
		Name:      "研发总监（CTO）",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 4},
		Name:      "高级运维开发工程师",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[2],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 5},
		Name:      "测试工程师",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[3],
		},
	},
}

// 迁移职位数据
func MigrateSystemJobPositionData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行职位数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_job_position")
	common.DB.Create(&systemJobPositions)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t职位数据迁移完成")
}
