package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统部门Api起始ID
var systemDepartmentApiStartId = (systemDepartmentApiCategoryId - 1) * 100

// 系统部门Api
var systemDepartmentApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: systemDepartmentApiStartId + 1},
		Name:                "获取系统部门列表",
		Method:              "GET",
		Path:                "/system/department/list",
		Description:         "获取系统部门列表",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemDepartmentApiCategoryId,
	},
}
