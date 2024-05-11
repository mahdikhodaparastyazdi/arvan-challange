package arvanlog

import (
	log "notification/pkg/logger"

	"github.com/rs/zerolog"
)

func convertToZeroLogLevel(level log.LogLevel) zerolog.Level {
	switch level {
	case log.DebugLevel:
		return zerolog.DebugLevel
	case log.InfoLevel:
		return zerolog.InfoLevel
	case log.WarnLevel:
		return zerolog.WarnLevel
	case log.ErrorLevel:
		return zerolog.ErrorLevel
	case log.FatalLevel:
		return zerolog.FatalLevel

	}

	return 0
}
