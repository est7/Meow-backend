package log

import (
	"Meow-backend/internal/interfaces"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log is A global variable so that log functions can be directly accessed
// 公共的全局 zapLogger 变量
var log interfaces.Logger

// 私有的 zap zapLogger 变量
var zapLogger *zap.Logger

//var zapLogger *zap.Logger

type Logger = interfaces.Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields = interfaces.Fields

func InitZapLogger(cfg *LoggerConfig, mode interfaces.Mode, opts ...Option) interfaces.Logger {
	var err error

	switch mode {
	case interfaces.DebugMode:
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout", "/var/log/debug.log"}
	case interfaces.TestMode:
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout", "/var/log/test.log"}
	case interfaces.ReleaseMode:
		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"stdout", "/var/log/release.log"}
	default:
		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"stdout", "/var/log/default.log"}
	}

	zapLogger, err = newZapLogger(cfg, opts...)
	if err != nil {
		panic(fmt.Sprintf("init newZapLogger err: %v", err))
	}

	log, err = newLoggerWithCallerSkip(cfg, 1, opts...)
	if err != nil {
		panic(fmt.Sprintf("init newLoggerWithCallerSkip err: %v", err))
	}

	return log
}

// GetLogger return a log
func GetLogger() Logger {
	return log
}

// GetZapLogger return raw zap zapLogger
func GetZapLogger() *zap.Logger {
	return zapLogger
}

// WithContext is a zapLogger that can log msg and log span for trace
// WithContext 是一个 zap Logger，可以记录消息和日志范围以进行跟踪
func WithContext(ctx context.Context) Logger {
	//return zap zapLogger

	if span := trace.SpanFromContext(ctx); span != nil {
		logger := spanLogger{span: span, logger: zapLogger}

		spanCtx := span.SpanContext()
		logger.spanFields = []zapcore.Field{
			zap.String("trace_id", spanCtx.TraceID().String()),
			zap.String("span_id", spanCtx.SpanID().String()),
		}

		return logger
	}
	return GetLogger()
}

// Debug zapLogger
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info zapLogger
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn zapLogger
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error zapLogger
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal zapLogger
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf zapLogger
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof zapLogger
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf zapLogger
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf zapLogger
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf zapLogger
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// WithFields zapLogger
// output more field, eg:
//
//	contextLogger := log.WithFields(log.Fields{"key1": "value1"})
//	contextLogger.Info("print multi field")
//
// or more sample to use:
//
//	log.WithFields(log.Fields{"key1": "value1"}).Info("this is a test log")
//	log.WithFields(log.Fields{"key1": "value1"}).Infof("this is a test log, user_id: %d", userID)
func WithFields(keyValues Fields) Logger {
	return GetLogger().WithFields(keyValues)
}
