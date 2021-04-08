package meilogrus

import (
	"github.com/sirupsen/logrus"
)

// MeiLogger meilogrus默认采用的日志实例
var MeiLogger = NewMeiLogger()

// NewMeiLogger NewMeiLogger
func NewMeiLogger() *logrus.Logger {
	mlHelper := new(MeiFileLog)
	ml, _ := mlHelper.GetLogger()
	return ml
}

// Debug Debug
func Debug(args ...interface{}) {
	MeiLogger.Debug(args...)
}

// Info Info
func Info(args ...interface{}) {
	MeiLogger.Info(args...)
}

// Warn Warn
func Warn(args ...interface{}) {
	MeiLogger.Warn(args...)
}

// Error Error
func Error(args ...interface{}) {
	MeiLogger.Error(args...)
}

// Fatal Fatal
func Fatal(args ...interface{}) {
	MeiLogger.Fatal(args...)
}
