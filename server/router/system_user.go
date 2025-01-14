package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 系统用户登录路由
func SystemUserAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/api-list", v1.GetSystemUserAPIListHandler) // 获取系统用户授权接口列表
	return rg
}

// 系统用户登录和授权路由
func SystemUserAuthAndPermission(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	return rg
}
