package log

import (
	"Meow-backend/pkg/utils"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// WriterConsole console输出
	WriterConsole = "console"
	// WriterFile 文件输出
	WriterFile = "file"

	logSuffix      = ".log"
	warnLogSuffix  = "_warn.log"
	errorLogSuffix = "_error.log"

	// defaultSkip zapLogger 包装了一层 zap.Logger，默认要跳过一层
	defaultSkip = 1
)

const (
	// RotateTimeDaily 按天切割
	RotateTimeDaily = "daily"
	// RotateTimeHourly 按小时切割
	RotateTimeHourly = "hourly"
)

var (
	hostname string
	logDir   string
)

// For mapping config zapLogger to app zapLogger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Prevent data race from occurring during zap.AddStacktrace
var zapStacktraceMutex sync.Mutex

func getLoggerLevel(cfg *LoggerConfig) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// zapLogger zapLogger struct
type sugarZapLogger struct {
	sugarLogger *zap.SugaredLogger
}

// newZapLogger new zap zapLogger
func newZapLogger(cfg *LoggerConfig, opts ...Option) (*zap.Logger, error) {
	for _, opt := range opts {
		opt(cfg)
	}
	return buildLogger(cfg, defaultSkip), nil
}

// newLoggerWithCallerSkip new zapLogger with caller skip
func newLoggerWithCallerSkip(cfg *LoggerConfig, skip int, opts ...Option) (Logger, error) {
	for _, opt := range opts {
		opt(cfg)
	}
	return &sugarZapLogger{sugarLogger: buildLogger(cfg, defaultSkip+skip).Sugar()}, nil
}

func buildLogger(cfg *LoggerConfig, skip int) *zap.Logger {
	logDir = cfg.LoggerDir
	if strings.HasSuffix(logDir, "/") {
		logDir = strings.TrimRight(logDir, "/")
	}

	var encoderCfg zapcore.EncoderConfig
	if cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == WriterConsole {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	var cores []zapcore.Core
	var options []zap.Option
	// init option
	hostname, _ = os.Hostname()
	option := zap.Fields(
		zap.String("ip", utils.GetLocalIP()),
		zap.String("app_id", cfg.ServiceName),
		zap.String("instance_id", hostname),
	)
	options = append(options, option)

	writers := strings.Split(cfg.Writers, ",")
	for _, w := range writers {
		switch w {
		case WriterConsole:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))
		case WriterFile:
			// info
			cores = append(cores, getInfoCore(encoder, cfg))

			// warning
			core, option := getWarnCore(encoder, cfg)
			cores = append(cores, core)
			if option != nil {
				options = append(options, option)
			}

			// error
			core, option = getErrorCore(encoder, cfg)
			cores = append(cores, core)
			if option != nil {
				options = append(options, option)
			}
		default:
			// console
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))
			// file
			cores = append(cores, getAllCore(encoder, cfg))
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	if !cfg.DisableCaller {
		caller := zap.AddCaller()
		options = append(options, caller)
	}

	// 跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(skip)
	options = append(options, addCallerSkip)

	// 构造日志
	return zap.New(combinedCore, options...)
}

func getAllCore(encoder zapcore.Encoder, cfg *LoggerConfig) zapcore.Core {
	allWriter := getLogWriterWithTime(cfg, GetLogFile(cfg.Filename, logSuffix))
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
}

func getInfoCore(encoder zapcore.Encoder, cfg *LoggerConfig) zapcore.Core {
	infoWrite := getLogWriterWithTime(cfg, GetLogFile(cfg.Filename, logSuffix))
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
}

func getWarnCore(encoder zapcore.Encoder, cfg *LoggerConfig) (zapcore.Core, zap.Option) {
	warnWrite := getLogWriterWithTime(cfg, GetLogFile(cfg.Filename, warnLogSuffix))
	var stacktrace zap.Option
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if !cfg.DisableCaller {
			zapStacktraceMutex.Lock()
			stacktrace = zap.AddStacktrace(zapcore.WarnLevel)
			zapStacktraceMutex.Unlock()
		}
		return lvl == zapcore.WarnLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(warnWrite), warnLevel), stacktrace
}

func getErrorCore(encoder zapcore.Encoder, cfg *LoggerConfig) (zapcore.Core, zap.Option) {
	errorFilename := GetLogFile(cfg.Filename, errorLogSuffix)
	errorWrite := getLogWriterWithTime(cfg, errorFilename)
	var stacktrace zap.Option
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if !cfg.DisableCaller {
			zapStacktraceMutex.Lock()
			stacktrace = zap.AddStacktrace(zapcore.ErrorLevel)
			zapStacktraceMutex.Unlock()
		}
		return lvl >= zapcore.ErrorLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), errorLevel), stacktrace
}

// getLogWriterWithTime 按时间(小时)进行切割
func getLogWriterWithTime(cfg *LoggerConfig, filename string) io.Writer {
	logFullPath := filename
	rotationPolicy := cfg.LogRollingPolicy
	backupCount := cfg.LogBackupCount
	// 默认
	var (
		rotateDuration time.Duration
		// 时间格式使用shell的date时间格式
		timeFormat string
	)
	if rotationPolicy == RotateTimeHourly {
		rotateDuration = time.Hour
		timeFormat = ".%Y%m%d%H"
	} else if rotationPolicy == RotateTimeDaily {
		rotateDuration = time.Hour * 24
		timeFormat = ".%Y%m%d"
	}
	hook, err := rotatelogs.New(
		logFullPath+timeFormat,
		rotatelogs.WithLinkName(logFullPath),        // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(backupCount),   // 文件最大保存份数
		rotatelogs.WithRotationTime(rotateDuration), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// Debug zapLogger
func (l *sugarZapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Info zapLogger
func (l *sugarZapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Warn zapLogger
func (l *sugarZapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Error zapLogger
func (l *sugarZapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *sugarZapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *sugarZapLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *sugarZapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *sugarZapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *sugarZapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *sugarZapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *sugarZapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *sugarZapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugarLogger.With(f...)
	return &sugarZapLogger{newLogger}
}
