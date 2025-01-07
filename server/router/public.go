package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 开放路由
func Public(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", v1.HealthHandler) // 健康检查
	return rg
}
