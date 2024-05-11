package health

import (
	"context"
	"time"
)

type healthServiceInterface interface {
	CheckHealth(ctx context.Context, token string, service string) error
}

type HealthHandler struct {
	healthService healthServiceInterface
}

func NewHealthHandler(service healthServiceInterface) *HealthHandler {
	return &HealthHandler{
		healthService: service,
	}
}

type healthCheckResponse struct {
	Healthy     bool          `json:"healthy"`
	Message     string        `json:"message"`
	ElapsedTime time.Duration `json:"elapsed_time"`
}
