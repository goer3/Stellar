package v1

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 获取系统用户授权接口列表
func GetSystemUserApiListHandler(ctx *gin.Context) {
	// 获取用户的角色 ID
	systemRoleId, err := utils.ExtractUintResultFromContext(ctx, "systemRoleId")
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色信息失败")
		return
	}

	// 如果角色是超级管理员，则返回所有 Api 列表
	if systemRoleId == common.SystemSuperAdministratorRoleId {
		var systemApis []model.SystemApi
		common.DB.Find(&systemApis)
	} else {
		var systemRole model.SystemRole
		common.DB.Preload("SystemApi").Where("id = ?", systemRoleId).First(&systemRole)
	}

	// 从 Context 中获取当前用户信息
	response.Success()
}
