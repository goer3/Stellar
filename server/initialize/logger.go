package initialize

import (
	"fmt"
	"os"
	"stellar/common"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志级别映射
var zapLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

// 日志日期格式调整
func zapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(common.TIME_FORMAT_MILLISECOND))
}

// 创建日志写入器
func createLogWriter(c common.LoggerConfiguration) zapcore.WriteSyncer {
	// 如果未启用日志，则输出到控制台
	if !c.Enabled {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	// 如果开启日志滚动，则使用日期格式化日志文件名
	var filename string
	if c.Rolling.Enabled {
		now := time.Now()
		filename = fmt.Sprintf("%s-%04d-%02d-%02d.log", c.PathPrefix, now.Year(), now.Month(), now.Day())
	} else {
		filename = c.PathPrefix + ".log"
	}

	// 配置 lumberjack
	hook := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    c.Rolling.MaxSize,
		MaxAge:     c.Rolling.MaxAge,
		MaxBackups: c.Rolling.MaxBackups,
		Compress:   c.Rolling.Compress,
	}

	// 如果不启用滚动，重置相关参数
	if !c.Rolling.Enabled {
		hook.MaxSize = 0
		hook.MaxAge = 0
		hook.MaxBackups = 0
		hook.Compress = false
	}

	fileWriter := zapcore.AddSync(hook)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), fileWriter)
}

// NewLogger 创建新的日志记录器
func NewLogger(c common.LoggerConfiguration) *zap.SugaredLogger {
	// 创建日志配置
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapLocalTimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	// 获取写入器
	ws := createLogWriter(c)

	// 获取日志级别
	level, ok := zapLevelMap[c.Level]
	if !ok {
		level = zapcore.InfoLevel // 默认使用 INFO 级别
	}

	// 创建编码器
	var encoder zapcore.Encoder
	if c.Format == "json" {
		encoder = zapcore.NewJSONEncoder(config)
	} else {
		encoder = zapcore.NewConsoleEncoder(config)
	}

	// 创建日志核心
	core := zapcore.NewCore(encoder, ws, level)

	// 创建日志记录器
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

// 系统日志初始化
func SystemLogger() {
	logger := NewLogger(common.Config.Log.System)
	common.SystemLog = logger
}

// 访问日志初始化
func AccessLogger() {
	logger := NewLogger(common.Config.Log.Access)
	common.AccessLog = logger
}
