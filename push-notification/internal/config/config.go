package config

import (
	"time"

	log "notification/pkg/logger"
	"notification/pkg/mysql"

	"notification/pkg/redis"

	"notification/pkg/sentry"
)

type LogLevel string

type AppEnv string

const (
	ProductionEnv AppEnv = "production"
	StageEnv      AppEnv = "stage"
	DevelopEnv    AppEnv = "develop"
	LocalEnv      AppEnv = "locale"
)

type (
	Config struct {
		AppEnv         AppEnv
		Locale         string
		AppDebug       bool
		LogLevel       log.LogLevelStr
		HealthToken    string
		HTTP           HTTP
		Database       Database
		SMSProvider    SMSProvider
		Sentry         *sentry.Config
		Tz             string
		Throttle       Throttle
		TrackerService TrackerService
		SmsConsumer    SmsConsumer
		MessagesTTL    MessagesTTL
		TrustedProxies []string
	}

	HTTP struct {
		Host string
		Port int
	}

	Database struct {
		MySQL mysql.Config
		Redis redis.Config
	}

	Sentry struct {
		Active           bool
		Dsn              string
		EnableTracing    bool
		TracesSampleRate float64
	}

	Throttle struct {
		Count     int
		PerMinute int
	}

	SMSProvider struct {
	}

	SmsConsumer struct {
		AsynqHighWorkerCount int
		AsynqLowWorkerCount  int
		AsynqLowMaxRetry     int
		AsynqHighMaxRetry    int
		AsynqTimeoutSeconds  time.Duration
	}

	MessagesTTL struct {
		OTP time.Duration
		SMS time.Duration
	}

	TrackerService struct {
		BaseUrl string
		ApiKey  string
	}
)
