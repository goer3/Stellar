package v1

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统职位列表
func GetSystemJobPositionListHandler(c *gin.Context) {
	var jobPositionList []model.SystemJobPosition
	if err := common.DB.Find(&jobPositionList).Error; err != nil {
		response.FailedWithMessage("获取系统职位列表失败")
		return
	}
	response.SuccessWithData(gin.H{"list": jobPositionList})
}
