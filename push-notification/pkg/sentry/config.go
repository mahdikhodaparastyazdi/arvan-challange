package sentry

type Config struct {
	Dsn                string
	EnableTracing      bool
	TracesSampleRate   float64
	Active             bool
	Debug              bool
	Environment        string
	SampleRate         float64
	ProfilesSampleRate float64
}
