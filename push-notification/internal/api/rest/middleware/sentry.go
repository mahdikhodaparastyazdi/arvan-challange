package middleware

import (
	"net/http"

	log "notification/pkg/logger"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use recover to capture any panics
		defer func() {
			if r := recover(); r != nil {
				// Recovered from a panic; capture the error and report it to Sentry
				if err, ok := r.(error); ok {
					sentry.CaptureException(err)
				}
				// Respond with a 500 Internal Server Error
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		// Handle the request
		c.Next()

		// If an error occurred, capture it and report it to Sentry
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.IsType(gin.ErrorTypePublic) {
					log.Info("Public error", log.J{
						"error": err.Err.Error(),
					})
					continue
				}
				sentry.CaptureException(err.Err)
			}
		}
	}
}
