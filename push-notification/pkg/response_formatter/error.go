package response_formatter

import (
	"errors"

	log "notification/pkg/logger"

	"notification/internal/derrors"

	"github.com/gin-gonic/gin"
)

func (r ResponseFormatter) Error(c *gin.Context, e error) {

	if errors.Is(e, derrors.ErrInternalServer) {
		r.logger.Error(e, log.J{
			"errors": c.Errors.Errors(),
		})

		c.JSON(derrors.ToStatus(derrors.ErrInternalServer), ResponseError{
			Code:         derrors.ToCode(derrors.ErrInternalServer),
			ErrorMessage: derrors.ErrInternalServer.Error(),
		})
		return
	}
	code := derrors.ToCode(e)
	if code == 0 {
		c.JSON(derrors.ToStatus(derrors.ErrUnhandled), ResponseError{
			Code:         derrors.ToCode(derrors.ErrUnhandled),
			ErrorMessage: derrors.ErrUnhandled.Error(),
		})
		return
	}
	c.JSON(derrors.ToStatus(e), ResponseError{
		Code:         derrors.ToCode(e),
		ErrorMessage: e.Error(),
	})
}

func (r ResponseFormatter) Errors(c *gin.Context, m any, code int) {
	c.JSON(code, ResponseError{
		Errors: m,
	})
}

func (r ResponseFormatter) ErrorMessage(c *gin.Context, m string, code int) {
	c.JSON(code, ResponseError{
		ErrorMessage: m,
	})
}

func (r ResponseFormatter) ErrorsWithMessage(c *gin.Context, m any, message string, code int) {
	c.JSON(code, ResponseError{
		ErrorMessage: message,
		Errors:       m,
	})
}
