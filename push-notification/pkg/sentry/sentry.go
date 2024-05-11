package sentry

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

type Sentry struct {
	Dsn                string
	EnableTracing      bool
	TracesSampleRate   float64
	Active             bool
	Debug              bool
	Environment        string
	SampleRate         float64
	ProfilesSampleRate float64
}

func NewSentry(cfg *Config) *Sentry {
	return &Sentry{
		Dsn:                cfg.Dsn,
		EnableTracing:      cfg.EnableTracing,
		TracesSampleRate:   cfg.TracesSampleRate,
		Active:             cfg.Active,
		Environment:        cfg.Environment,
		SampleRate:         cfg.SampleRate,
		ProfilesSampleRate: cfg.ProfilesSampleRate,
		Debug:              cfg.Debug,
	}
}

func (s *Sentry) InitSentry() error {
	if !s.Active {
		return nil
	}

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           s.Dsn,
		EnableTracing: s.EnableTracing,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate:   s.TracesSampleRate,
		Debug:              s.Debug,
		AttachStacktrace:   true,
		SampleRate:         s.SampleRate,
		ProfilesSampleRate: s.ProfilesSampleRate,
		Environment:        s.Environment,
	}); err != nil {
		return fmt.Errorf("Sentry initialization failed: %v", err)
	}

	return nil
}
