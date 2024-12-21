package logger

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// DPanic, Panic and Fatal level can not be set by user.
const (
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

// Config contains logger configuration.
type Config struct {
	Lvl        string `yaml:"lvl" default:"info"`
	File       string `yaml:"file"`
	Dev        bool   `yaml:"dev"`
	MaxSize    int    `yaml:"max_size" default:"64" validate:"number,gte=0"`
	MaxBackups int    `yaml:"max_backups" default:"30" validate:"number,gte=0"`
	MaxAge     int    `yaml:"max_age" default:"90" validate:"number,gte=0"`
	Compress   bool   `yaml:"compress"`

	ZapLvl zapcore.Level
}

var ErrUnknownLogLvl = errors.New("unknown log error")

func (cfg *Config) AfterLoad() error {
	switch cfg.Lvl {
	case DebugLevelStr:
		cfg.ZapLvl = zap.DebugLevel
	case InfoLevelStr:
		cfg.ZapLvl = zap.InfoLevel
	case WarningLevelStr:
		cfg.ZapLvl = zap.WarnLevel
	case ErrorLevelStr:
		cfg.ZapLvl = zap.ErrorLevel
	default:
		return fmt.Errorf("%w: %s", ErrUnknownLogLvl, cfg.Lvl)
	}

	return nil
}

func New(cfg *Config) (*zap.Logger, func() error, error) {
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize, // MB
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, // days
		Compress:   cfg.Compress,
	})
	core := zapcore.NewCore(
		// use NewConsoleEncoder for human-readable output
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), ws),
		zap.NewAtomicLevelAt(cfg.ZapLvl),
	)

	var zapLogger *zap.Logger
	if cfg.Dev {
		zapLogger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		zapLogger = zap.New(core)
	}

	zap.ReplaceGlobals(zapLogger)

	return zapLogger, func() error { return zapLogger.Sync() }, nil
}
