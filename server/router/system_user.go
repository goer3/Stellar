package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 系统用户登录路由
func SystemUserAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	return rg
}

// 系统用户登录和授权路由
func SystemUserAuthAndPermission(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	return rg
}
