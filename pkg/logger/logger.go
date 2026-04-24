package logger

import (
	"os"
	"simplebank/pkg/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSettings) *LoggerZap {
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	logLevel := config.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName, // Log file path
		MaxSize:    config.MaxSize,     // Max size in megabytes before log rotation
		MaxBackups: config.MaxBackups,  // Max number of backup files to keep
		MaxAge:     config.MaxAge,      // Max age in days before log rotation
		Compress: config.Compress,      // Compress old log files, disable by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {

	// Customize the log format
	encoderConfig := zap.NewProductionEncoderConfig()

	// Use ISO8601 format for timestamps
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Set the key for the timestamp field to "timestamp"
	encoderConfig.TimeKey = "timestamp"

	// Use capital letters for log levels (e.g., INFO, ERROR)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Use short caller format (file:line)
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Return a JSON encoder with the customized configuration
	return zapcore.NewJSONEncoder(encoderConfig)

}

