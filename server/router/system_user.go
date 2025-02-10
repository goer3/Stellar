package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 系统用户登录路由
func SystemUserAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	return rg
}

// 系统用户登录和授权路由
func SystemUserAuthAndPermission(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", v1.GetSystemUserListHandler)
	rg.GET("/detail", v1.GetSystemUserDetailHandler)
	rg.POST("/create", v1.CreateSystemUserHandler)
	rg.POST("/batch-import", v1.BatchImportSystemUserHandler)
	rg.PUT("/update", v1.UpdateSystemUserDetailHandler)
	rg.PUT("/batch-modify", v1.BatchModifySystemUserHandler)
	return rg
}
