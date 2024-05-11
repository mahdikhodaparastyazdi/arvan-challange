package middleware

import (
	"context"
	"net/http"
	"notification/internal/config"
	"notification/internal/constants"
	"notification/internal/derrors"
	"time"

	response_formatter "notification/pkg/response_formatter"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ThrottleMiddleware struct {
	smsService        smsServiceInterface
	responseFormatter response_formatter.ResponseFormatter
	count             int
	perMin            int
}

type smsServiceInterface interface {
	CountSince(ctx context.Context, mobile string, duration time.Time) (int, error)
}

type Data struct {
	Mobile string `json:"mobile"`
}

func NewThrottleMiddleware(
	us smsServiceInterface,
	cfg config.Throttle,
	responseFormatter response_formatter.ResponseFormatter,
) ThrottleMiddleware {
	return ThrottleMiddleware{
		smsService:        us,
		count:             cfg.Count,
		perMin:            cfg.PerMinute,
		responseFormatter: responseFormatter,
	}
}

func (t *ThrottleMiddleware) Throttle(c *gin.Context) {
	var data Data

	SMSRate := c.GetString("sms_rate")
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		_ = c.Error(err)
		t.responseFormatter.Error(c, derrors.ErrValidation)
		return
	}

	actualCount, err := t.smsService.CountSince(c, data.Mobile, time.Now().Add(-1*time.Minute))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if SMSRate == constants.SMS_RATE_LOW && actualCount >= t.count {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		return
	}
	c.Next()
}
