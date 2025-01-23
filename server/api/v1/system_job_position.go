package v1

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统职位列表
func GetSystemJobPositionListHandler(c *gin.Context) {
	var systemJobPositionList []model.SystemJobPosition
	if err := common.DB.Find(&systemJobPositionList).Error; err != nil {
		response.FailedWithMessage("获取系统职位列表失败")
		return
	}
	response.SuccessWithData(gin.H{"list": systemJobPositionList})
}
