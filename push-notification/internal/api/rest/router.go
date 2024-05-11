package rest

import (
	"notification/internal/api/rest/handlers/health"
	"notification/internal/api/rest/handlers/template"
	"notification/internal/api/rest/middleware"
)

func (s *Server) SetupMonitoringRoutes(healthHandler *health.HealthHandler) {
	r := s.engine
	r.GET("/health", healthHandler.CheckHealth)
}

func (s *Server) SetupAPIRoutes(
	templateHandler template.Handler,
	internalMiddleware middleware.InternalMiddleware,
	throttleMiddleware middleware.ThrottleMiddleware,
) {
	r := s.engine

	v1 := r.Group("/v1")
	{

		internal := v1.Group("/internal", internalMiddleware.Handle)

		internal.POST("/sms/template", throttleMiddleware.Throttle, templateHandler.SendSmsTemplated)

		// TODO: implement simple send
		//internal.POST("/sms/send", throttleMiddleware.Throttle, templateHandler.SendSms)

		// TODO: implement sms history
		//internal.POST("/sms/report/", SMsHandler.Reports)
	}
	// TODO: implement client apis
}
