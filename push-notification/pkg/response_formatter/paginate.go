package response_formatter

import (
	"math"
	"net/http"
	"reflect"

	"notification/pkg/response_formatter/paginator"

	"github.com/gin-gonic/gin"
)

func (r ResponseFormatter) Paginate(c *gin.Context, total int64, paginator paginator.Paginate, data interface{}) {

	count := reflect.ValueOf(data).Len()

	totalPages := int64(math.Ceil(float64(total) / float64(paginator.Size)))
	if totalPages <= 0 {
		totalPages = 0
	}

	c.JSON(http.StatusOK, Response{
		Data: &data,
		Meta: &Meta{
			Pagination{
				Total:       total,
				Count:       count,
				PerPage:     paginator.Size,
				CurrentPage: paginator.Page,
				TotalPages:  totalPages,
			},
		},
	})
}

func (r ResponseFormatter) PaginateWithCount(c *gin.Context, total int64, count int, paginator paginator.Paginate, data interface{}) {
	totalPages := int64(math.Ceil(float64(total) / float64(paginator.Size)))
	if totalPages <= 0 {
		totalPages = 0
	}

	c.JSON(http.StatusOK, Response{
		Data: &data,
		Meta: &Meta{
			Pagination{
				Total:       total,
				Count:       count,
				PerPage:     paginator.Size,
				CurrentPage: paginator.Page,
				TotalPages:  totalPages,
			},
		},
	})
}
