package v1

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 获取系统角色授权接口列表
func GetSystemRoleApiListHandler(ctx *gin.Context) {
	// 获取用户的角色 ID
	systemRoleId, err := utils.ExtractUintResultFromContext(ctx, "systemRoleId")
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色信息失败")
		return
	}
	// 如果角色是超级管理员，则返回所有 Api 列表
	var systemApiList []string
	if systemRoleId == common.SystemSuperAdministratorRoleId {
		var systemApis []model.SystemApi
		if err := common.DB.Find(&systemApis).Error; err != nil {
			response.FailedWithMessage("获取系统Api列表失败")
			return
		}
		for _, systemApi := range systemApis {
			systemApiList = append(systemApiList, systemApi.Path)
		}
	} else {
		var systemRole model.SystemRole
		if err := common.DB.Preload("SystemApis").Where("id = ?", systemRoleId).First(&systemRole).Error; err != nil {
			response.FailedWithMessage("获取系统角色信息失败")
			return
		}
		for _, systemApi := range systemRole.SystemApis {
			systemApiList = append(systemApiList, systemApi.Path)
		}
	}
	response.SuccessWithData(gin.H{
		"list": systemApiList,
	})
}
