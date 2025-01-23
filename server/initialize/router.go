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
	gin.SetMode(common.Config.System.Mode)

	// 创建一个没有中间件的 Gin 路由引擎
	r := gin.New()
	r.Use(middleware.AccessLog)       // 访问日志中间件
	r.Use(middleware.Cors)            // 跨域中间件
	r.Use(middleware.Exception)       // 异常处理中间件
	auth, err := middleware.JWTAuth() // 初始化 JWT 中间件
	if err != nil {
		panic("JWT 认证中间件初始化失败: " + err.Error())
	}
	// 路由注册
	{
		// 开放路由，不需要认证
		publicRouterGroup := r.Group(common.SYSTEM_API_PREFIX)
		router.Public(publicRouterGroup, auth)
	}
	{
		// 开放路由，需要认证，但是不需要授权的路由
		publicAuthRouterGroup := r.Group(common.SYSTEM_API_PREFIX)
		publicAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.PublicAuth(publicAuthRouterGroup, auth)
	}
	{
		// 系统路由
		systemApiPrefix := common.SYSTEM_API_PREFIX + "/system"
		{
			// 系统职位-路由
			systemJobPositionApiPrefix := systemApiPrefix + "/job-position"
			{
				{
					// 系统职位，登录路由
					systemJobPositionAuthRouterGroup := r.Group(systemJobPositionApiPrefix)
					systemJobPositionAuthRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemJobPositionAuth(systemJobPositionAuthRouterGroup, auth)
				}
				{
					// 系统职位，登录和授权路由
					systemJobPositionAuthAndPermissionRouterGroup := r.Group(systemJobPositionApiPrefix)
					systemJobPositionAuthAndPermissionRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemJobPositionAuthAndPermission(systemJobPositionAuthAndPermissionRouterGroup, auth)
				}
			}
			// 系统部门-路由
			systemDepartmentApiPrefix := systemApiPrefix + "/department"
			{
				{
					// 系统部门，登录路由
					systemDepartmentAuthRouterGroup := r.Group(systemDepartmentApiPrefix)
					systemDepartmentAuthRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemDepartmentAuth(systemDepartmentAuthRouterGroup, auth)
				}
				{
					// 系统部门，登录和授权路由
					systemDepartmentAuthAndPermissionRouterGroup := r.Group(systemDepartmentApiPrefix)
					systemDepartmentAuthAndPermissionRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemDepartmentAuthAndPermission(systemDepartmentAuthAndPermissionRouterGroup, auth)
				}
			}
			// 系统用户-路由
			systemUserApiPrefix := systemApiPrefix + "/user"
			{
				{
					// 系统用户，登录路由
					systemUserAuthRouterGroup := r.Group(systemUserApiPrefix)
					systemUserAuthRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemUserAuth(systemUserAuthRouterGroup, auth)
				}
				{
					// 系统用户，登录和授权路由
					systemUserAuthAndPermissionRouterGroup := r.Group(systemUserApiPrefix)
					systemUserAuthAndPermissionRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemUserAuthAndPermission(systemUserAuthAndPermissionRouterGroup, auth)
				}
			}
			// 系统角色-路由
			systemRoleApiPrefix := systemApiPrefix + "/role"
			{
				{
					// 系统角色，登录路由
					systemRoleAuthRouterGroup := r.Group(systemRoleApiPrefix)
					systemRoleAuthRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemRoleAuth(systemRoleAuthRouterGroup, auth)
				}
				{
					// 系统角色，登录和授权路由
					systemRoleAuthAndPermissionRouterGroup := r.Group(systemRoleApiPrefix)
					systemRoleAuthAndPermissionRouterGroup.Use(auth.MiddlewareFunc())
					router.SystemRoleAuthAndPermission(systemRoleAuthAndPermissionRouterGroup, auth)
				}
			}
		}
	}
	return r
}
