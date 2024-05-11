package template

import (
	"net/http"
	"notification/internal/derrors"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllTemplates(c *gin.Context) {
	templates, err := h.templateService.GetAllTemplates(c)
	if err != nil {
		_ = c.Error(err)
		h.responseFormatter.Error(c, derrors.ErrInternalServer)
		return
	}

	result := h.templateTransformer.TransformMany(c, templates)

	h.responseFormatter.Success(c, result, http.StatusOK)
}
