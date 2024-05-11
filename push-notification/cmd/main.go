package main

import (
	"context"
	"fmt"
	"notification/cmd/command"
	"notification/internal/app"
	"notification/internal/config"
	"notification/internal/version"
	"os/signal"
	"syscall"

	log "notification/pkg/logger"
	sentryPkg "notification/pkg/sentry"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Printf("Version: %v\nRelease Date: %v\nCommit Hash: %v\n\n\n", version.Version, version.ReleaseDate, version.CommitHash)
	const description = "arvan notification"
	root := &cobra.Command{Short: description}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app.InitLogger()

	sentry := sentryPkg.NewSentry(&sentryPkg.Config{
		Dsn:              cfg.Sentry.Dsn,
		EnableTracing:    cfg.Sentry.EnableTracing,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
		Active:           cfg.Sentry.Active,
	})
	err = sentry.InitSentry()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	root.AddCommand(
		command.Version{}.Command(),
		command.Server{}.Command(ctx, cfg),
		command.Consumer{}.Command(ctx, cfg),
	)

	if err := root.Execute(); err != nil {
		log.Fatal(fmt.Sprintf("failed go execute root command: \n %s", err.Error()))
	}
}
