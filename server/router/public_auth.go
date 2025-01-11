package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 开放路由，不需要授权，但是需要认证
func PublicAuth(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.POST("/logout", auth.LogoutHandler) // 注销登录
	return rg
}
