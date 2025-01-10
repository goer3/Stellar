package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统API公共数据
var systemPublicApiStartId = systemPublicApiCategoryId
var systemPublicApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: systemPublicApiStartId},
		Name:                "健康检查",
		Method:              "GET",
		Path:                "/health",
		Description:         "健康检查，用于检查服务是否正常",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemPublicApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemPublicApiStartId + 1},
		Name:                "获取系统信息",
		Method:              "GET",
		Path:                "/info",
		Description:         "获取系统信息，用于获取系统基础信息",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemPublicApiCategoryId,
	},
}
