package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware(c *gin.Context) {
	// Use recover to capture any panics
	defer func() {
		if r := recover(); r != nil {
			// Recovered from a panic; capture the error and report it to Sentry
			if err, ok := r.(error); ok {
				sentry.CaptureException(err)
			}
		}
	}()

	// Handle the request
	c.Next()

	// If an error occurred, capture it and report it to Sentry
	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			sentry.CaptureException(err.Err)
		}
	}
}
