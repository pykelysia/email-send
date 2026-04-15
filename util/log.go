package util

import (
	"email-send/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// InitLogger 初始化日志
func InitLogger(c *config.Config) error {
	// 创建编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	level := zapcore.InfoLevel

	// 创建编码器
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建日志写入器
	var cores []zapcore.Core

	// 添加控制台输出
	consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
	cores = append(cores, consoleCore)

	// 如果配置了日志文件路径，添加文件输出
	if c.LogConfig.LogPath != "" {
		// 确保日志目录存在
		logDir := c.LogConfig.LogPath
		if dir := os.Getenv("LOG_DIR"); dir != "" {
			logDir = dir + "/" + c.LogConfig.LogPath
		}

		// 创建日志文件
		file, err := os.OpenFile(logDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// 如果无法打开文件，仍继续使用控制台日志
			logger.Warnf("无法打开日志文件 %s: %v", logDir, err)
		} else {
			// 文件编码器（无颜色，用于文件输出）
			fileEncoderConfig := encoderConfig
			fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			fileEncoder := zapcore.NewConsoleEncoder(fileEncoderConfig)
			fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level)
			cores = append(cores, fileCore)
		}
	}

	// 创建多输出核心
	core := zapcore.NewTee(cores...)

	// 创建 logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = zapLogger.Sugar()

	return nil
}

// Info 输出 Info 级别日志
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof 输出 Info 级别格式化日志
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Debug 输出 Debug 级别日志
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf 输出 Debug 级别格式化日志
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Warn 输出 Warn 级别日志
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf 输出 Warn 级别格式化日志
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Error 输出 Error 级别日志
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf 输出 Error 级别格式化日志
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// Fatal 输出 Fatal 级别日志
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf 输出 Fatal 级别格式化日志
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Print 打印日志
func Print(args ...interface{}) {
	logger.Info(args...)
}

// Printf 打印格式化日志
func Printf(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Sync 刷新日志缓冲区
func Sync() {
	if logger != nil {
		logger.Sync()
	}
}
