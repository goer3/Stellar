package v1

import (
	"net/http"
	"stellar/common"

	"github.com/gin-gonic/gin"
)

// 健康检查接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// 系统信息接口
func SystemInfoHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"systemName":             common.SYSTEM_NAME,
		"systemDescription":      common.SYSTEM_DESCRIPTION,
		"systemVersion":          common.SystemVersion,
		"systemGoVersion":        common.SYSTEM_GO_VERSION,
		"systemDeveloperName":    common.SYSTEM_DEVELOPER_NAME,
		"systemDeveloperEmail":   common.SYSTEM_DEVELOPER_EMAIL,
		"systemGithubRepository": common.SYSTEM_GITHUB_REPOSITORY,
	})
}
