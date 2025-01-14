package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统公共Api起始ID
var publicApiStartId = publicApiCategoryId

// 系统公共Api
var publicApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: publicApiStartId},
		Name:                "健康检查",
		Method:              "GET",
		Path:                "/health",
		Description:         "健康检查，用于检查服务是否正常",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: publicApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: publicApiStartId + 1},
		Name:                "获取系统信息",
		Method:              "GET",
		Path:                "/information",
		Description:         "获取系统信息，用于获取系统基础信息",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: publicApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: publicApiStartId + 2},
		Name:                "用户登录",
		Method:              "POST",
		Path:                "/login",
		Description:         "用户登录，用于用户登录系统认证",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: publicApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: publicApiStartId + 3},
		Name:                "注销登录",
		Method:              "GET",
		Path:                "/logout",
		Description:         "注销登录，用于用户注销系统登录状态",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: publicApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: publicApiStartId + 4},
		Name:                "校验Token是否过期",
		Method:              "GET",
		Path:                "/token-verification",
		Description:         "校验Token是否过期，用于校验Token是否过期",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: publicApiCategoryId,
	},
}
