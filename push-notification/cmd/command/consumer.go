package command

import (
	"context"
	"fmt"
	"notification/internal/config"
	"notification/internal/constants"
	"notification/internal/consumers/sms_consumer"
	"notification/internal/factories/sms_resolver"
	"notification/internal/repositories"
	"notification/internal/tasks"
	"notification/pkg/asynq"

	log "notification/pkg/logger"
	"notification/pkg/logger/arvanlog"
	"notification/pkg/mysql"

	"github.com/spf13/cobra"
)

type Consumer struct {
	highPriority bool
	lowPriority  bool
	logger       log.Logger
}

func (cmd Consumer) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	cmd.logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:consumer:command")

	consumerCmd := &cobra.Command{
		Use:   "consumer",
		Short: "run auth consumer",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(ctx, cfg)
		},
	}

	consumerCmd.Flags().BoolVarP(&cmd.highPriority, "sms-high-priority", "", false, "Run high consumer")
	consumerCmd.Flags().BoolVarP(&cmd.lowPriority, "sms-low-priority", "", false, "Run low consumer")

	return consumerCmd
}
func (cmd Consumer) main(ctx context.Context, cfg *config.Config) {
	if cmd.highPriority {
		cmd.HighPriorityConsumer(ctx, cfg)
	}
	if cmd.lowPriority {
		cmd.LowPriorityConsumer(ctx, cfg)
	}
}

func (cmd Consumer) HighPriorityConsumer(ctx context.Context, cfg *config.Config) {
	db, err := mysql.NewClient(ctx, &cfg.Database.MySQL)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}
	gormDB, err := mysql.NewGormWithInstance(db, cfg.AppDebug)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}

	err = mysql.Migrate(db)
	if err != nil {
		cmd.logger.Fatal(fmt.Errorf("mysql migration failed: %w", err))
	}

	smsRepository := repositories.NewSmsRepository(gormDB)
	logger := arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:provider-resolver")
	resolver := sms_resolver.NewResolver(cfg.SMSProvider, cfg.AppEnv, logger)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:consumer:sms:high-priority")
	highConsumer := sms_consumer.New(cfg.SMSProvider, logger, smsRepository, resolver)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:asynq-high-priority-server")
	server := asynq.NewServer(logger, cfg.Database.Redis, constants.QUEUE_PRIORITY_HIGH, cfg.SmsConsumer.AsynqHighWorkerCount)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:high-priority-worker")
	worker := tasks.NewWorker(server, highConsumer, logger)
	if err := worker.StartWorker(constants.QUEUE_PRIORITY_HIGH); err != nil {
		cmd.logger.Error(err)
		return
	}
}

func (cmd Consumer) LowPriorityConsumer(ctx context.Context, cfg *config.Config) {
	db, err := mysql.NewClient(ctx, &cfg.Database.MySQL)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}
	gormDB, err := mysql.NewGormWithInstance(db, cfg.AppDebug)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}

	err = mysql.Migrate(db)
	if err != nil {
		cmd.logger.Fatal(fmt.Errorf("mysql migration failed: %w", err))
	}

	smsRepository := repositories.NewSmsRepository(gormDB)
	logger := arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:provider-resolver")
	resolver := sms_resolver.NewResolver(cfg.SMSProvider, cfg.AppEnv, logger)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:consumer:sms:low-priority")
	highConsumer := sms_consumer.New(cfg.SMSProvider, logger, smsRepository, resolver)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:asynq-low-priority-server")
	server := asynq.NewServer(logger,
		cfg.Database.Redis,
		constants.QUEUE_PRIORITY_LOW,
		cfg.SmsConsumer.AsynqLowWorkerCount,
	)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:sms:low-priority-worker")
	worker := tasks.NewWorker(server, highConsumer, logger)
	if err := worker.StartWorker(constants.QUEUE_PRIORITY_LOW); err != nil {
		cmd.logger.Error(err)
		return
	}
}
