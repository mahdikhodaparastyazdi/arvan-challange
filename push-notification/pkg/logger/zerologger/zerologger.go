package zerologger

import (
	"os"

	"github.com/rs/zerolog"
)

const skipStackFrameCount = 2

type ZerologLogger struct {
	zerolog zerolog.Logger
}

func NewStdOutLogger() ZerologLogger {
	return ZerologLogger{
		zerolog: zerolog.New(os.Stdout),
	}
}

func NewFromZerolog(z zerolog.Logger) ZerologLogger {
	return ZerologLogger{
		zerolog: z,
	}
}

func (z ZerologLogger) Fatal(message any, data ...map[string]any) {
	if len(data) > 0 {
		z.zerolog.Fatal().Caller(skipStackFrameCount).Interface("data", data[0]).Msgf("%v", message)
		return
	}
	z.zerolog.Fatal().Caller(skipStackFrameCount).Msgf("%v", message)
}

func (z ZerologLogger) Error(message any, data ...map[string]any) {
	z.zerolog.Error().Caller(skipStackFrameCount).Interface("data", data).Msgf("%v", message)
}

func (z ZerologLogger) Warn(message any, data ...map[string]any) {
	z.zerolog.Warn().Caller(skipStackFrameCount).Interface("data", data).Msgf("%v", message)
}

func (z ZerologLogger) Info(message any, data ...map[string]any) {
	z.zerolog.Info().Caller(skipStackFrameCount).Interface("data", data).Msgf("%v", message)
}

func (z ZerologLogger) Debug(message any, data ...map[string]any) {
	z.zerolog.Debug().Interface("data", data).Msgf("%v", message)
}
