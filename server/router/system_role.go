package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 系统角色登录路由
func SystemRoleAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/api-list", v1.GetSystemRoleApiListHandler) // 获取系统角色授权接口列表
	return rg
}

// 系统角色登录和授权路由
func SystemRoleAuthAndPermission(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	return rg
}
