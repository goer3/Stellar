package v1

import (
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 获取系统用户授权接口列表
func GetSystemUserAPIListHandler(ctx *gin.Context) {
	// 从 Context 中获取当前用户信息
	response.Success()
}
