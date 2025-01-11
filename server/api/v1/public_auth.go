package v1

import (
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 校验 Token 是否过期
func TokenVerificationHandler(ctx *gin.Context) {
	response.Success()
}
