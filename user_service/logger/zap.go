package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewProductionConfig().Build(zap.AddCallerSkip(1))
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	sugar := logger.Sugar()

	return &zapLogger{logger: sugar}
}

func (l *zapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}
