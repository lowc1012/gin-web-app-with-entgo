package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLogger = New(os.Stderr, zap.InfoLevel, "production")

var (
	Info    = DefaultLogger.Info
	Infow   = DefaultLogger.Infow
	Warn    = DefaultLogger.Warn
	Warnw   = DefaultLogger.Warnw
	Error   = DefaultLogger.Error
	Errorw  = DefaultLogger.Errorw
	DPanic  = DefaultLogger.DPanic
	DPanicw = DefaultLogger.DPanicw
	Panic   = DefaultLogger.Panic
	Panicw  = DefaultLogger.Panicw
	Fatal   = DefaultLogger.Fatal
	Fatalw  = DefaultLogger.Fatalw
	Debug   = DefaultLogger.Debug
	Debugw  = DefaultLogger.Debugw
)

func New(writer io.Writer, level zapcore.Level, env string, extraOpts ...zap.Option) *zap.SugaredLogger {
	var cfg zapcore.EncoderConfig
	var encoder zapcore.Encoder
	opts := make([]zap.Option, 0, len(extraOpts))

	switch env {
	case "production":
		cfg = zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewJSONEncoder(cfg)
	case "development", "mock":
		cfg = zap.NewDevelopmentEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewConsoleEncoder(cfg)
		opts = append(opts, zap.WithCaller(true))
	case "test":
		cfg = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			NameKey:        "logger",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}
		encoder = zapcore.NewJSONEncoder(cfg)
		opts = append(opts, zap.WithCaller(true))
	}

	opts = append(opts, extraOpts...)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(writer),
		level,
	)

	return zap.New(core, opts...).Sugar()
}

func Reset(l *zap.SugaredLogger) {
	DefaultLogger = l
	Info = DefaultLogger.Info
	Infow = DefaultLogger.Infow
	Warn = DefaultLogger.Warn
	Warnw = DefaultLogger.Warnw
	Error = DefaultLogger.Error
	Errorw = DefaultLogger.Errorw
	DPanic = DefaultLogger.DPanic
	DPanicw = DefaultLogger.DPanicw
	Panic = DefaultLogger.Panic
	Panicw = DefaultLogger.Panicw
	Fatal = DefaultLogger.Fatal
	Fatalw = DefaultLogger.Fatalw
	Debug = DefaultLogger.Debug
	Debugw = DefaultLogger.Debugw
}

func Init() *zap.SugaredLogger {
	var logLevel zapcore.Level

	// read global loglevel
	switch strings.ToLower(config.Global.LogLevel) {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn", "warning":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "dpanic":
		logLevel = zapcore.DPanicLevel
	case "panic":
		logLevel = zapcore.PanicLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	DefaultLogger = New(
		os.Stderr,
		logLevel,
		config.Global.Env,
	)

	Reset(DefaultLogger)

	return DefaultLogger
}

func StdInfo(format string, v ...any) {
	Info(fmt.Sprintf(format, v...))
}
