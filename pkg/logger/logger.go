package logger

import (
	// 引入必要的包

	"Campusforum/settings"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

// Init 初始化日志
func Init(cfg *settings.LogConfig, mode string) (err error) {
	// 初始化日志
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		// 如果是开发模式，将日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),                                     //传到日志文件
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), //传到控制台
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l) //传到日志文件
	}

	// 创建 Logger，并将调用信息添加到日志中
	lg = zap.New(core, zap.AddCaller())
	// 替换全局 Logger
	zap.ReplaceGlobals(lg)
	// 输出初始化成功的日志
	zap.L().Info("init logger success")
	return
}

// getEncoder 获取日志编码器
func getEncoder() zapcore.Encoder {
	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 获取日志写入器
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// 根据时间生成文件名
	filename = "logs/" + generateLogFilename(filename)
	// 配置 Lumberjack 日志写入器
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	// 将 Lumberjack 日志写入器添加到 Zap 的 Syncer 中
	return zapcore.AddSync(lumberJackLogger)
}

// generateLogFilename 根据时间生成日志文件名
func generateLogFilename(filename string) string {
	// 使用当前时间生成文件名
	return filename + time.Now().Format("2006-01-02") + ".log"
}

// GinLogger 记录 Gin 框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		// 计算请求处理耗时
		cost := time.Since(start)
		// 记录请求相关信息
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery 恢复可能出现的 panic，并记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 延迟处理 panic
		defer func() {
			if err := recover(); err != nil {
				// 检查是否为断开连接错误
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取 HTTP 请求内容
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					// 如果连接已断开，则记录错误并中止请求
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				// 处理 panic，并记录日志
				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		// 执行后续中间件和处理程序
		c.Next()
	}
}
