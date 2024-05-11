package app

import (
	"os"

	log "notification/pkg/logger"
	zerologLogger "notification/pkg/logger/zerologger"

	"github.com/rs/zerolog"
)

func InitLogger() {
	z := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.SetDefaultLogger(zerologLogger.NewFromZerolog(z))
}
