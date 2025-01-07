package initialize

import (
	"stellar/common"
	"stellar/middleware"
	"stellar/router"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func Router() *gin.Engine {
	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)

	// 创建一个没有中间件的 Gin 路由引擎
	r := gin.New()
	r.Use(middleware.AccessLog) // 访问日志中间件
	r.Use(middleware.Cors)      // 跨域中间件
	r.Use(middleware.Exception) // 异常处理中间件
	{
		// 开放路由，不需要认证
		publicRouterGroup := r.Group(common.SYSTEM_API_PREFIX)
		router.Public(publicRouterGroup)
	}
	{
		// 开放路由，需要认证，但是不需要授权的路由
		publicAuthRouterGroup := r.Group(common.SYSTEM_API_PREFIX)
		router.PublicAuth(publicAuthRouterGroup)
	}
	return r
}
