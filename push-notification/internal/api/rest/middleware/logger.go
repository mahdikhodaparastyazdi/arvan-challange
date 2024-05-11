package middleware

import (
	"time"

	log "notification/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		data := log.J{
			"client_id":   param.ClientIP,
			"method":      param.Method,
			"status_code": param.StatusCode,
			"path":        param.Path,
			"latency":     param.Latency.String(),
		}

		const serverErrorStatus = 500
		if c.Writer.Status() >= serverErrorStatus {
			log.Error(param.ErrorMessage, data)
		} else {
			log.Info(param.ErrorMessage, data)
		}
	}
}

type GinRecoveryWriter struct{}

func (w *GinRecoveryWriter) Write(p []byte) (int, error) {
	log.Error(string(p))
	return len(p), nil
}
