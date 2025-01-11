package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 开放路由
func Public(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", v1.HealthHandler)           // 健康检查
	rg.GET("/information", v1.InformationHandler) // 系统信息
	rg.POST("/login", auth.LoginHandler)          // 用户登录
	return rg
}

// 开放路由，不需要授权，但是需要认证
func PublicAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/logout", auth.LogoutHandler)                      // 注销登录
	rg.GET("/token-verification", v1.TokenVerificationHandler) // 校验 Token 是否过期
	return rg
}
