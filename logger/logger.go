package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

var (
	Int    = zap.Int
	String = zap.String
	Error  = zap.Error
	Bool   = zap.Bool
	Any    = zap.Any
)

type LoggerI interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type loggerImpl struct {
	zap *zap.Logger
}

func NewLogger(namespace string) LoggerI {
	logger := loggerImpl{
		zap: newZapLogger(namespace),
	}

	return &logger
}

func (l *loggerImpl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *loggerImpl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *loggerImpl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *loggerImpl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *loggerImpl) DPanic(msg string, fields ...Field) {
	l.zap.DPanic(msg, fields...)
}

func (l *loggerImpl) Panic(msg string, fields ...Field) {
	l.zap.Panic(msg, fields...)
}

func (l *loggerImpl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

func GetNamed(l LoggerI, name string) LoggerI {
	switch v := l.(type) {
	case *loggerImpl:
		v.zap = v.zap.Named(name)
		return v
	default:
		l.Info("logger.GetNamed: invalid logger type")
		return l
	}
}

func WithFields(l LoggerI, fields ...Field) LoggerI {
	switch v := l.(type) {
	case *loggerImpl:
		return &loggerImpl{
			zap: v.zap.With(fields...),
		}
	default:
		l.Info("logger.WithFields: invalid logger type")
		return l
	}
}

func Cleanup(l LoggerI) error {
	switch v := l.(type) {
	case *loggerImpl:
		return v.zap.Sync()
	default:
		l.Info("logger.Cleanup: invalid logger type")
		return nil
	}
}
