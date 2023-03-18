package util

import (
	"github.com/bavelee/mfc/cfg"
	"os"
)
import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 10:44
 * @Desc:
 */

func SetupGlobalLogger() *zap.Logger {
	fileEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    10,
		MaxBackups: 1,
		MaxAge:     10,
	})

	consoleConfig := zap.NewProductionEncoderConfig()
	consoleConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

	var core zapcore.Core
	if os.Getenv("RUNTIME_MODE") == "DEBUG" {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, fileWriter, zap.DebugLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel))
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, fileWriter, zap.InfoLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.InfoLevel))
	}
	logger := zap.New(core, zap.AddCaller())
	_ = logger.Sync()
	zap.ReplaceGlobals(logger)
	return logger
}
