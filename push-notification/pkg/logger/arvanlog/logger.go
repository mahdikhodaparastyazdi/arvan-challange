package arvanlog

import (
	"os"

	log "notification/pkg/logger"

	"github.com/rs/zerolog"
)

const skipStackFrameCount = 2

type ArvanlogLogger struct {
	arvanlog zerolog.Logger
	prefix   string
}

func NewStdOutLogger(logLevel log.LogLevelStr, prefix string) ArvanlogLogger {
	return ArvanlogLogger{
		arvanlog: zerolog.New(os.Stdout).With().Timestamp().Logger().Level(convertToZeroLogLevel(log.StrLogLevelToInt(logLevel))),
		prefix:   prefix,
	}
}

func NewFromArvanlog(z zerolog.Logger) ArvanlogLogger {
	return ArvanlogLogger{
		arvanlog: z,
	}
}

func (z ArvanlogLogger) Fatal(message any, data ...map[string]any) {
	z.arvanlog.Fatal().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ArvanlogLogger) Error(message any, data ...map[string]any) {
	z.arvanlog.Error().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ArvanlogLogger) Warn(message any, data ...map[string]any) {
	z.arvanlog.Warn().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ArvanlogLogger) Info(message any, data ...map[string]any) {
	z.arvanlog.Info().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ArvanlogLogger) Debug(message any, data ...map[string]any) {
	z.arvanlog.Debug().Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ArvanlogLogger) CloneWithPrefix(prefix string) log.ClonableLogger {
	z.prefix += ":" + prefix
	return z
}
