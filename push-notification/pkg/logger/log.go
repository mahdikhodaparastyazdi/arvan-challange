package log

import (
	"strings"

	"notification/pkg/logger/zerologger"
)

type LogLevel int
type LogLevelStr string

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const (
	DebugLevelStr LogLevelStr = "DEBUG"
	InfoLevelStr  LogLevelStr = "INFO"
	WarnLevelStr  LogLevelStr = "WARN"
	ErrorLevelStr LogLevelStr = "ERROR"
	FatalLevelStr LogLevelStr = "FATAL"
)

type Logger interface {
	Fatal(message any, data ...J)
	Error(message any, data ...J)
	Warn(message any, data ...J)
	Info(message any, data ...J)
	Debug(message any, data ...J)
}

type ClonableLogger interface {
	Logger
	CloneWithPrefix(prefix string) ClonableLogger
}

type J = map[string]any

var logger Logger

func init() {
	logger = zerologger.NewStdOutLogger()
}

func SetDefaultLogger(l Logger) {
	logger = l
}

func Fatal(message any, data ...J) {
	logger.Fatal(message, data...)
}

func Error(message any, data ...J) {
	logger.Error(message, data...)
}

func Warn(message any, data ...J) {
	logger.Warn(message, data...)
}

func Info(message any, data ...J) {
	logger.Info(message, data...)
}

func Debug(message any, data ...J) {
	logger.Debug(message, data...)
}

func StrLogLevelToInt(level LogLevelStr) LogLevel {
	switch strings.ToUpper(string(level)) {
	case string(DebugLevelStr):
		return DebugLevel

	case string(InfoLevelStr):
		return InfoLevel

	case string(WarnLevelStr):
		return WarnLevel

	case string(ErrorLevelStr):
		return ErrorLevel

	case string(FatalLevelStr):
		return FatalLevel

	}
	return 0
}
