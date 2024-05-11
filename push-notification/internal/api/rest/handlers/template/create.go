package template

import (
	"net/http"
	"notification/internal/api/rest/requests"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h Handler) CreateTemplate(c *gin.Context) {
	var req requests.CreateTemplate

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO: implement validator for request

	res, err := h.templateService.CreteTemplate(c, req)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	result := h.templateTransformer.Transform(c, res)
	h.responseFormatter.Success(c, result, http.StatusCreated)
}
