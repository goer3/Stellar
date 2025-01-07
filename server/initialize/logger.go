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
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(common.TIME_FORMAT_MILLISECOND))
}

// 日志初始化
func NewLogger(c common.LoggerConfiguration) *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()       // 新建配置
	config.EncodeTime = ZapLocalTimeEncoder          // 调整时间
	config.EncodeLevel = zapcore.CapitalLevelEncoder // 关闭颜色
	var ws zapcore.WriteSyncer
	if c.Enabled {
		var fileWriter zapcore.WriteSyncer
		if c.Rolling.Enabled {
			// 开启日志滚动，需要配置滚动参数
			now := time.Now()
			filename := fmt.Sprintf("%s-%04d-%02d-%02d.log", c.PathPrefix, now.Year(), now.Month(), now.Day())
			hook := &lumberjack.Logger{
				Filename:   filename,
				MaxSize:    c.Rolling.MaxSize,
				MaxAge:     c.Rolling.MaxAge,
				MaxBackups: c.Rolling.MaxBackups,
				Compress:   c.Rolling.Compress,
			}
			defer func(hook *lumberjack.Logger) {
				_ = hook.Close()
			}(hook)
			fileWriter = zapcore.AddSync(hook)
		} else {
			// 不开启日志滚动，则不配置滚动参数
			filename := c.PathPrefix + ".log"
			hook := &lumberjack.Logger{
				Filename:   filename,
				MaxSize:    0,
				MaxAge:     0,
				MaxBackups: 0,
				Compress:   false,
			}
			defer func(hook *lumberjack.Logger) {
				_ = hook.Close()
			}(hook)
			fileWriter = zapcore.AddSync(hook)
		}
		// 输出到控制台和文件
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), fileWriter)
	} else {
		// 只输出到控制台
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	// 整合日志输出信息
	var core zapcore.Core
	if c.Format == "json" {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(config), ws, zapcore.Level(zapLevelMap[c.Level]))
	} else {
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(config), ws, zapcore.Level(zapLevelMap[c.Level]))
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
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
