package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 健康检查接口函数
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
