package template

import (
	"errors"
	"net/http"
	"notification/internal/api/rest/requests"
	"notification/internal/constants"
	"notification/internal/derrors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h Handler) SendSmsTemplated(c *gin.Context) {
	var req requests.SendWithTemplateRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		h.responseFormatter.Error(c, derrors.ErrValidation)
		return
	}
	// TODO: implement validator for request
	// TODO: implement validator for mobile

	clientId, err := strconv.Atoi(c.GetString("client_id"))
	if err != nil {
		return
	}
	err = h.templateService.SendTemplate(c, req, uint(clientId))
	if err == nil {
		h.responseFormatter.Success(c, nil, http.StatusOK)
		return
	}

	_ = c.Error(err)
	if errors.Is(err, constants.ErrTemplateNotFound) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if errors.Is(err, constants.ErrProviderIsNotActive) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusBadRequest)
		return
	}
	if errors.Is(err, constants.ErrTemplateProviderNotDefined) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusBadRequest)
		return
	}
	if errors.Is(err, constants.ErrExpiryDateTime) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	h.responseFormatter.Error(c, derrors.ErrInternalServer)
}
