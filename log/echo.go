package log

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func NewEchoLogger(logger *zap.Logger) *echoLogger {
	return &echoLogger{logger: logger}
}

type echoLogger struct {
	logger *zap.Logger
}

func (e *echoLogger) SetHeader(h string) {
	panic("implement me")
}

func (e *echoLogger) Output() io.Writer {
	return os.Stderr
}

func (e *echoLogger) SetOutput(w io.Writer) {
	// panic("implement SetOutput")
}

func (e *echoLogger) Prefix() string {
	panic("implement Prefix")
}

func (e *echoLogger) SetPrefix(p string) {
	panic("implement SetPrefix")
}

func (e *echoLogger) Level() log.Lvl {
	panic("implement Level")
}

func (e *echoLogger) SetLevel(lvl log.Lvl) {
	panic("implement SetLevel")
}

func (e *echoLogger) Print(i ...interface{}) {
	panic("implement Print")
}

func (e *echoLogger) Printf(format string, args ...interface{}) {
	e.logger.Sugar().Infof(format, args)
}

func (e *echoLogger) Printj(json log.JSON) {
	panic("implement Printj")
}

func (e *echoLogger) Debug(i ...interface{}) {
	panic("implement Debug")
}

func (e *echoLogger) Debugf(format string, args ...interface{}) {
	panic("implement Debugf")
}

func (e *echoLogger) Debugj(json log.JSON) {
	panic("implement Debugj")
}

func (e *echoLogger) Info(i ...interface{}) {
	panic("implement Info")
}

func (e *echoLogger) Infof(format string, args ...interface{}) {
	panic("implement Infof")
}

func (e *echoLogger) Infoj(json log.JSON) {
	panic("implement Infoj")
}

func (e *echoLogger) Warn(i ...interface{}) {
	panic("implement Warn")
}

func (e *echoLogger) Warnf(format string, args ...interface{}) {
	panic("implement Warnf")
}

func (e *echoLogger) Warnj(json log.JSON) {
	panic("implement Warnj")
}

func (e *echoLogger) Error(i ...interface{}) {
	e.logger.Sugar().Errorf("", i)
}

func (e *echoLogger) Errorf(format string, args ...interface{}) {
	panic("implement Errorf")
}

func (e *echoLogger) Errorj(json log.JSON) {
	panic("implement Errorj")
}

func (e *echoLogger) Fatal(i ...interface{}) {
	panic("implement Fatal")
}

func (e *echoLogger) Fatalj(json log.JSON) {
	panic("implement Fatalj")
}

func (e *echoLogger) Fatalf(format string, args ...interface{}) {
	panic("implement Fatalf")
}

func (e *echoLogger) Panic(i ...interface{}) {
	panic("implement Panic")
}

func (e *echoLogger) Panicj(json log.JSON) {
	panic("implement Panicj")
}

func (e *echoLogger) Panicf(format string, args ...interface{}) {
	panic("implement Panicf")
}
