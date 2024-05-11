package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "notification/pkg/logger"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"notification/internal/api/rest/middleware"

	"github.com/gin-gonic/gin"
)

const headerTimeout = 10 * time.Second

type Server struct {
	engine *gin.Engine
	logger log.Logger
}

func New(logger log.Logger, trustedProxies []string) *Server {
	r := gin.New()
	err := r.SetTrustedProxies(trustedProxies)
	if err != nil {
		log.Fatal("trusted proxies has error", log.J{
			"error": err.Error(),
		})
	}

	r.RedirectTrailingSlash = false

	r.Use(middleware.GinLogger())

	r.Use(gin.RecoveryWithWriter(&middleware.GinRecoveryWriter{}))
	r.Use(middleware.SentryMiddleware())

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	return &Server{
		engine: r,
		logger: logger,
	}
}

func (s *Server) Serve(ctx context.Context, address string) error {
	srv := &http.Server{
		Addr:              address,
		Handler:           s.engine,
		ReadHeaderTimeout: headerTimeout,
	}

	s.logger.Info(fmt.Sprintf("rest server starting at: %s", address))
	srvError := make(chan error)
	go func() {
		srvError <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		s.logger.Info("rest server is shutting down")
		return srv.Shutdown(ctx)
	case err := <-srvError:
		return err
	}
}
