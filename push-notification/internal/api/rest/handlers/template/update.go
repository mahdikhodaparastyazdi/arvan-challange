package template

import (
	"net/http"
	"notification/internal/api/rest/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) UpdateTemplate(c *gin.Context) {
	tmplID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateTemplate
	req.ID = uint(tmplID)

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	// TODO: implement validator for request

	if err := h.templateService.UpdateTemplate(c, req); err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	h.responseFormatter.Success(c, nil, http.StatusNoContent)
}
