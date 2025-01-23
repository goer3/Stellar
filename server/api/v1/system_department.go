package v1

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统部门列表
func GetSystemDepartmentListHandler(c *gin.Context) {
	var departmentList []model.SystemDepartment
	if err := common.DB.Find(&departmentList).Error; err != nil {
		response.FailedWithMessage("获取系统部门列表失败")
		return
	}
	response.SuccessWithData(gin.H{"list": departmentList})
}
