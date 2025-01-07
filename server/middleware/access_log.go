package middleware

import (
	"fmt"
	"stellar/common"
	"time"

	"github.com/gin-gonic/gin"
)

// 访问日志中间件
func AccessLog(ctx *gin.Context) {
	// 请求时间，毫秒时间戳
	startMillisecondTimestamp := time.Now().UnixMilli()
	// 处理请求
	ctx.Next()
	// 结束时间，毫秒时间戳
	endMillisecondTimestamp := time.Now().UnixMilli()
	// 执行耗时，毫秒
	execMillisecond := endMillisecondTimestamp - startMillisecondTimestamp
	// 请求方式
	method := ctx.Request.Method
	// 请求地址
	requestURI := ctx.Request.RequestURI
	// 状态码
	status := ctx.Writer.Status()
	// 来源 IP
	clientIP := ctx.ClientIP()
	// 完整的日志
	var logString string
	if common.Config.Log.Access.Format == "json" {
		logJson := map[string]interface{}{
			"method":          method,
			"requestURI":      requestURI,
			"status":          status,
			"execMillisecond": execMillisecond,
			"clientIP":        clientIP,
		}
		if method == "OPTIONS" {
			common.AccessLog.With("request", logJson).Debug()
		} else {
			common.AccessLog.With("request", logJson).Info()
		}
	} else {
		logString = fmt.Sprintf("%s\t%s\t%d\t%dms\t%s",
			method,
			requestURI,
			status,
			execMillisecond,
			clientIP,
		)
		// 打印日志，OPTIONS 请求使用 DEBUG
		if method == "OPTIONS" {
			common.AccessLog.Debug(logString)
		} else {
			common.AccessLog.Info(logString)
		}
	}
}
