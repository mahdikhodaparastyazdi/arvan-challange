package middleware

import (
	"encoding/json"
	"net/http"

	log "notification/pkg/logger"

	"github.com/gin-gonic/gin"
)

type PublicAuthMiddleware struct {
	logger log.Logger
}

func NewPublicAuthMiddleware(logger log.Logger) PublicAuthMiddleware {
	return PublicAuthMiddleware{
		logger: logger,
	}
}

type response struct {
	UserID uint `json:"user_id"`
}

func (p PublicAuthMiddleware) respondUnauthorized(c *gin.Context) {
	p.logger.Error("failed to retrieve X-Auth-Data information")
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "forbidden"})
}

func (p PublicAuthMiddleware) Handle(c *gin.Context) {
	auth := c.GetHeader("X-Auth-Data")

	if auth == "" {
		p.respondUnauthorized(c)
		return
	}

	var resp response
	err := json.Unmarshal([]byte(auth), &resp)
	if err != nil || resp.UserID == 0 {
		p.respondUnauthorized(c)
		return
	}

	c.Set("user_id", resp.UserID)

	c.Next()
}
