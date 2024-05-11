package command

import (
	"context"
	"fmt"
	"notification/internal/api/rest"
	template_handler "notification/internal/api/rest/handlers/template"
	"notification/internal/api/rest/middleware"
	"notification/internal/api/rest/transformer"
	"notification/internal/config"
	"notification/internal/repositories"
	clientservice "notification/internal/services/client"
	sms_service "notification/internal/services/sms"
	template_service "notification/internal/services/template"
	"notification/internal/tasks"
	"notification/pkg/asynq"

	log "notification/pkg/logger"
	"notification/pkg/logger/arvanlog"
	"notification/pkg/mysql"
	response_formatter "notification/pkg/response_formatter"

	"github.com/spf13/cobra"
)

type Server struct {
	logger log.Logger
}

func (cmd Server) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	cmd.logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:server:command")

	return &cobra.Command{
		Use:   "server",
		Short: "run setting server",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(ctx, cfg)
		},
	}
}

func (cmd Server) main(ctx context.Context, cfg *config.Config) {
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

	smsRepo := repositories.NewSmsRepository(gormDB)
	templateRepo := repositories.NewTemplateRepository(gormDB)
	providerRepo := repositories.NewProviderRepository(gormDB)
	providerTemplateRepo := repositories.NewProviderTemplate(gormDB)
	clientRepo := repositories.NewClientRepository(gormDB)

	clientService := clientservice.New(clientRepo)
	smsService := sms_service.New(smsRepo)

	asynqClient := asynq.NewClient(cfg.Database.Redis)
	queue := tasks.NewQueue(asynqClient, cfg.SmsConsumer.AsynqHighMaxRetry, cfg.SmsConsumer.AsynqTimeoutSeconds)
	logger := arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:api:service:template-service")
	templateService := template_service.New(clientRepo, smsRepo, templateRepo, queue, providerRepo, providerTemplateRepo, cfg.MessagesTTL, logger)

	templateTransformer := transformer.NewTemplateTransformer()

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification::api:response-formatter-pkg")
	responseFormatter := response_formatter.NewResponseFormatter(logger)
	templateHandler := template_handler.New(templateService, responseFormatter, templateTransformer)

	internalMiddleware := middleware.NewInternalMiddleware(clientService)
	throttleMiddleware := middleware.NewThrottleMiddleware(smsService, cfg.Throttle, responseFormatter)

	logger = arvanlog.NewStdOutLogger(cfg.LogLevel, "notification:api:server")
	server := rest.New(logger, cfg.TrustedProxies)
	server.SetupAPIRoutes(
		templateHandler,
		internalMiddleware,
		throttleMiddleware,
	)
	if err := server.Serve(ctx, fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
		cmd.logger.Fatal(err)
	}
}
