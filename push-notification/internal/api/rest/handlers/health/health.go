package health

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	token := c.Query("token")
	services := strings.Split(c.Query("services"), ",")

	startTime := time.Now()

	for _, s := range services {
		if err := h.healthService.CheckHealth(c.Request.Context(), token, s); err != nil {
			c.JSON(http.StatusInternalServerError, healthCheckResponse{
				Healthy:     false,
				Message:     fmt.Sprintf("%s is not ok", s),
				ElapsedTime: time.Since(startTime),
			})
			return
		}
	}

	c.JSON(http.StatusOK, healthCheckResponse{
		Healthy:     true,
		Message:     "ok",
		ElapsedTime: time.Since(startTime),
	})
}
