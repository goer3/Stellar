package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统职位Api起始ID
var systemJobPositionApiStartId = (systemJobPositionApiCategoryId - 1) * 100

// 系统职位Api
var systemJobPositionApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: systemJobPositionApiStartId + 1},
		Name:                "获取系统职位列表",
		Method:              "GET",
		Path:                "/system/job-position/list",
		Description:         "获取系统职位列表",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemJobPositionApiCategoryId,
	},
}
