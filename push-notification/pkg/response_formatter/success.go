package response_formatter

import (
	"github.com/gin-gonic/gin"
)

func (r ResponseFormatter) Success(c *gin.Context, data any, statusCode int) {
	c.JSON(statusCode, Response{
		Data: &data,
	})
}
