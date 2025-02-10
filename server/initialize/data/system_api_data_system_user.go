package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统用户Api起始ID
var systemUserApiStartId = (systemUserApiCategoryId - 1) * 100

// 系统用户Api
var systemUserApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 1},
		Name:                "获取用户列表",
		Method:              "GET",
		Path:                "/system/user/list",
		Description:         "获取所有用户列表，支持分页、排序、搜索、筛选",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 2},
		Name:                "获取用户详情",
		Method:              "GET",
		Path:                "/system/user/detail",
		Description:         "获取用户详情，默认获取自己，需要通过传入用户 id 来获取其他用户详情",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 3},
		Name:                "创建用户",
		Method:              "POST",
		Path:                "/system/user/create",
		Description:         "用户创建，支持单用户创建和批量用户创建，默认传递一个用户列表对象",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 4},
		Name:                "批量导入用户",
		Method:              "POST",
		Path:                "/system/user/batch-import",
		Description:         "批量导入用户，通过上传模板文件导入用户",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 5},
		Name:                "更新用户",
		Method:              "PUT",
		Path:                "/system/user/update",
		Description:         "用户更新，通过不同的标识来实现用户的单个字段更新和批量字段更新",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemUserApiStartId + 6},
		Name:                "批量修改用户",
		Method:              "PUT",
		Path:                "/system/user/batch-modify",
		Description:         "批量修改用户，通过不同的标识来实现用户信息修改",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemUserApiCategoryId,
	},
}
