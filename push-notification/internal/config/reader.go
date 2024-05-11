package config

import (
	"errors"
	"fmt"

	"notification/internal/constants"
	log "notification/pkg/logger"
	"notification/pkg/redis"
	"notification/pkg/sentry"

	"notification/pkg/mysql"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	appEnv := loadString("APP_ENV")
	debug := loadBool("APP_DEBUG")
	constants.SmsPrice = loadInt("SMS_PRICE")
	cfg := Config{
		AppEnv:   AppEnv(appEnv),
		Locale:   loadString("LOCALE"),
		AppDebug: debug,
		Tz:       loadString("TZ"),
		LogLevel: log.LogLevelStr(loadString("LOG_LEVEL")),
		HTTP: HTTP{
			Host: loadString("API_HTTP_HOST"),
			Port: loadInt("API_HTTP_PORT"),
		},
		Throttle: Throttle{
			Count:     loadInt("SMS_LIMIT_COUNT"),
			PerMinute: loadInt("SMS_LIMIT_PERIOD_MINS"),
		},
		Database: Database{
			MySQL: mysql.Config{
				Host:         loadString("MYSQL_HOST"),
				Port:         loadInt("MYSQL_PORT"),
				Username:     loadString("MYSQL_USER"),
				Password:     loadString("MYSQL_PASSWORD"),
				DatabaseName: loadString("MYSQL_DATABASE"),
				Timezone:     loadString("TZ"),
			},
			Redis: redis.Config{
				Host:     loadString("REDIS_HOST"),
				Port:     loadString("REDIS_PORT"),
				Password: loadString("REDIS_PASSWORD"),
				Database: loadInt("REDIS_DATABASE"),
			},
		},
		Sentry: &sentry.Config{
			Dsn:                loadString("SENTRY_DSN"),
			EnableTracing:      loadBool("SENTRY_ENABLE_TRACING"),
			TracesSampleRate:   loadFloat64("SENTRY_TRACES_SAMPLE_RATE"),
			Active:             loadBool("SENTRY_ACTIVE"),
			Debug:              debug,
			Environment:        appEnv,
			SampleRate:         loadFloat64("SENTRY_SAMPLE_RATE"),
			ProfilesSampleRate: loadFloat64("SENTRY_PROFILES_SAMPLE_RATE"),
		},
		TrackerService: TrackerService{
			BaseUrl: loadString("TRACKER_BASE_URL"),
			ApiKey:  loadString("TRACKER_API_KEY"),
		},
		SmsConsumer: SmsConsumer{
			AsynqHighWorkerCount: loadInt("ASYNQ_SMS_HIGH_WORKER_COUNT"),
			AsynqLowWorkerCount:  loadInt("ASYNQ_SMS_LOW_WORKER_COUNT"),
			AsynqLowMaxRetry:     loadInt("ASYNQ_SMS_JOB_LOW_MAX_RETRY"),
			AsynqHighMaxRetry:    loadInt("ASYNQ_SMS_JOB_HIGH_MAX_RETRY"),
			AsynqTimeoutSeconds:  loadDuration("ASYNQ_SMS_JOB_TIMEOUT_IN_SECONDS"),
		},
		MessagesTTL: MessagesTTL{
			OTP: loadDuration("DEFAULT_OTP_TTL"),
			SMS: loadDuration("DEFAULT_SMS_TTL"),
		},
		TrustedProxies: loadStringSlice("TRUSTED_PROXIES"),
	}

	return &cfg, nil
}
