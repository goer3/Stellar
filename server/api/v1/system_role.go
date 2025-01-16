package v1

import (
	"context"
	"encoding/json"
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 获取系统角色授权接口列表
func GetSystemRoleApiListHandler(ctx *gin.Context) {
	// 获取用户的角色关键字
	systemRoleKeyword, err := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色信息失败")
		return
	}

	// 拼接 Redis 中的 key
	key := common.RKP_SYSTEM_ROLE_API_LIST + common.RK_SEPARATOR + systemRoleKeyword
	// 创建 Redis 上下文
	rctx := context.Background()
	// 创建系统Api列表
	var systemApiList []string

	// 查询 Redis 是否有缓存角色Api列表
	result := common.Cache.Get(rctx, key).Val()
	if result != "" {
		// 如果值存在，则需要通过 json 将值转换成 []string
		if err := json.Unmarshal([]byte(result), &systemApiList); err == nil {
			response.SuccessWithData(gin.H{
				"list": systemApiList,
			})
			return
		}
	}

	// Redis 中没有查到数据，则需要查询数据库
	if systemRoleKeyword == common.SystemSuperAdministratorRoleKeyword {
		// 如果角色是超级管理员，则返回所有 Api 列表
		var systemApis []model.SystemApi
		if err := common.DB.Find(&systemApis).Error; err != nil {
			response.FailedWithMessage("获取系统Api列表失败")
			return
		}
		for _, systemApi := range systemApis {
			systemApiList = append(systemApiList, systemApi.Path)
		}
	} else {
		// 不是超级管理员，则需要查询角色对应的 Api 列表
		var systemRole model.SystemRole
		if err := common.DB.Preload("SystemApis").Where("keyword = ?", systemRoleKeyword).First(&systemRole).Error; err != nil {
			response.FailedWithMessage("获取系统角色信息失败")
			return
		}
		for _, systemApi := range systemRole.SystemApis {
			systemApiList = append(systemApiList, systemApi.Path)
		}
	}

	// 将查询到的数据缓存到 Redis 中，但是需要将 []string 转换成 string
	if systemApiListJson, err := json.Marshal(systemApiList); err == nil {
		common.Cache.Set(rctx, key, string(systemApiListJson), common.RKE_SYSTEM_ROLE_API_LIST)
	}

	response.SuccessWithData(gin.H{
		"list": systemApiList,
	})
}

// 获取角色的菜单列表
func GetSystemRoleMenuListHandler(ctx *gin.Context) {
	// 获取用户的角色 ID
	systemRoleId, err := utils.ExtractUintResultFromContext(ctx, "systemRoleId")
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色信息失败")
		return
	}
	// 通过角色查询菜单
	var systemRole model.SystemRole
	if err := common.DB.Preload("SystemMenus").Where("id = ?", systemRoleId).First(&systemRole).Error; err != nil {
		response.FailedWithMessage("获取系统角色的菜单信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": systemRole.SystemMenus,
	})
}
