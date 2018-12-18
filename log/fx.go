package log

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type fxLogger struct {
	logger *zap.Logger
}

func (l fxLogger) Printf(s string, args ...interface{}) {
	l.logger.Sugar().Debugf(s, args...)
}

func NewFXLogger(logger *zap.Logger) fx.Printer {
	return fxLogger{logger: logger}
}
