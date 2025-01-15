package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 系统角色Api起始ID
var systemRoleApiStartId = (systemRoleApiCategoryId - 1) * 100

// 系统角色Api
var systemRoleApis = []model.SystemApi{
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 1},
		Name:                "获取系统角色列表",
		Method:              "GET",
		Path:                "/system/role/list",
		Description:         "获取系统角色列表，支持分页、排序、搜索、筛选",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 2},
		Name:                "获取系统角色详情",
		Method:              "GET",
		Path:                "/system/role/detail",
		Description:         "获取系统角色详情，默认获取自己，需要通过传入角色 id 来获取其他角色详情",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 3},
		Name:                "创建系统角色",
		Method:              "POST",
		Path:                "/system/role/create",
		Description:         "创建系统角色",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 4},
		Name:                "更新系统角色",
		Method:              "PUT",
		Path:                "/system/role/update",
		Description:         "更新系统角色",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 5},
		Name:                "删除系统角色",
		Method:              "DELETE",
		Path:                "/system/role/delete",
		Description:         "删除系统角色",
		Permission:          trans.Uint(common.BOOLEAN_TRUE),
		Disabled:            trans.Uint(common.BOOLEAN_FALSE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
	{
		BaseModel:           model.BaseModel{Id: systemRoleApiStartId + 6},
		Name:                "获取当前用户角色的Api列表",
		Method:              "GET",
		Path:                "/system/role/api-list",
		Description:         "获取当前用户角色的Api列表，超级管理员返回所有Api列表，其他角色返回当前角色拥有的Api列表",
		Permission:          trans.Uint(common.BOOLEAN_FALSE),
		Disabled:            trans.Uint(common.BOOLEAN_TRUE),
		Creator:             common.SYSTEM_DEFAULT_CREATOR,
		SystemApiCategoryId: systemRoleApiCategoryId,
	},
}
