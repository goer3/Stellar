package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 开放路由
func Public(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", v1.HealthHandler)   // 健康检查
	rg.GET("/info", v1.SystemInfoHandler) // 系统信息
	rg.POST("/login", auth.LoginHandler)  // 用户登录
	return rg
}
