/**
 * @date: 2022/3/12
 * @desc: ...
 */

package cronLogger

import (
	"DataCompare/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

// CronLogger 初始化logger
func CronLogger(logFileName string) *zap.Logger {
	var (
		level        = "info"                  // debug, info, warn, error, dpanic, panic, fatal
		logPath      = "./logs/" + logFileName // 日志绝对路径
		logInConsole = utils.IsDebugEnv()      // 打印到控制台
		maxSize      = 1                       // 在进行切割之前，日志文件的最大大小（以MB为单位)
		maxBackups   = 10                      // 保留旧文件的最大个数
		maxAge       = 30                      // 保留旧文件的最大天数
		compress     = false                   // 是否压缩/归档旧文件
	)
	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	logLevel := getLogLevel(level)
	//atomicLevel.SetLevel(logLevel)
	// 状态配置
	writeSyncer := getLogWriter(logPath, maxSize, maxBackups, maxAge, compress, logInConsole)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	return zap.New(core, zap.AddCaller())
}

// getLogLevel 获取日志级别
func getLogLevel(logLevel string) zapcore.Level {
	logLevel = strings.ToLower(logLevel)
	var LogLevelInfo = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	level, ok := LogLevelInfo[logLevel]
	if !ok {
		level = zapcore.InfoLevel
	}
	return level
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "cronLogger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // FullCallerEncoder 全路径编码器  ShortCallerEncoder 简单路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter
func getLogWriter(logPath string, maxSize, maxBackups, maxAge int, compress, logInConsole bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,    // 日志的位置
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩/归档旧文件
	}
	defer lumberJackLogger.Close()
	syncer := zapcore.AddSync(lumberJackLogger)
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return syncer
}
