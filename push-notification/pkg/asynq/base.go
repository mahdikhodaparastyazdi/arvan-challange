package asynq

import (
	"fmt"

	log "notification/pkg/logger"
	"notification/pkg/redis"

	"github.com/hibiken/asynq"
)

func NewClient(cfg redis.Config) *asynq.Client {
	if cfg.Port == "" {
		cfg.Port = "6379"
	}
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
}

func NewServer(logger log.Logger, cfg redis.Config, queueName string, workerCount int) *asynq.Server {
	if cfg.Port == "" {
		cfg.Port = "6379"
	}
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	}, asynq.Config{
		Queues: map[string]int{
			queueName: 10,
		},
		Concurrency: workerCount,
		Logger: asynqLogger{
			logger: logger,
		},
	})
}

type asynqLogger struct {
	logger log.Logger
}

func (a asynqLogger) Debug(args ...interface{}) {
	a.logger.Debug(args[1], log.J{
		"sender": "asynq",
		"args":   args[1:],
	})
}

func (a asynqLogger) Info(args ...interface{}) {
	a.logger.Info(args[0], log.J{
		"sender": "asynq",
		"args":   args[1:],
	})
}

func (a asynqLogger) Warn(args ...interface{}) {
	a.logger.Warn(args[0], log.J{
		"sender": "asynq",
		"args":   args[1:],
	})
}

func (a asynqLogger) Error(args ...interface{}) {
	a.logger.Error(args[0], log.J{
		"sender": "asynq",
		"args":   args[1:],
	})
}

func (a asynqLogger) Fatal(args ...interface{}) {
	a.logger.Fatal(args[0], log.J{
		"sender": "asynq",
		"args":   args[1:],
	})
}
