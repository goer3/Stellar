package router

import (
	"github.com/gin-gonic/gin"
)

// 开放路由，不需要授权，但是需要认证
func PublicAuth(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}
