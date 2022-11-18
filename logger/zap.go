package logger

import (
	"os"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/term"
)

func newZapLogger(namespace string) *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})

	logStdErrorWriter := zapcore.Lock(os.Stderr)
	logStdInfoWriter := zapcore.Lock(os.Stdout)

	b := term.IsTerminal(int(os.Stderr.Fd()))

	core := zapcore.NewTee(
		zapcore.NewCore(logging.NewEncoder(4, b), logStdErrorWriter, highPriority),
		zapcore.NewCore(logging.NewEncoder(4, b), logStdInfoWriter, lowPriority),
	)

	logger := zap.New(
		core,
		zap.AddCaller(), zap.AddCallerSkip(1),
	)

	logger = logger.Named(namespace)

	zap.RedirectStdLog(logger)

	return logger
}
