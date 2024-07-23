package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

type FxLogger struct {
	*Logger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger
func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger()
		globalLogger = &logger
	}
	return *globalLogger
}

// GetFxLogger get the fx logger
func (l Logger) GetFxLogger() FxLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// newLogger sets up logger
func newLogger() Logger {

	config := zap.NewDevelopmentConfig()
	env := os.Getenv("ENV")

	if env == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

// Printf prits go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	l.Infof(str, args)
}
