package response_formatter

import log "notification/pkg/logger"

type ResponseFormatter struct {
	logger log.Logger
}

type Pagination struct {
	Total       int64 `json:"total" example:"20"`
	Count       int   `json:"count" example:"5"`
	PerPage     int   `json:"per_page" example:"5"`
	CurrentPage int   `json:"current_page" example:"2"`
	TotalPages  int64 `json:"total_pages" example:"4"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Response struct {
	Data *any  `json:"data,omitempty" swaggerignore:"true"`
	Meta *Meta `json:"meta,omitempty" swaggerignore:"true"`
}

type ResponseError struct {
	Code         int    `json:"code,omitempty" swaggerignore:"true"`
	Errors       any    `json:"errors,omitempty" swaggerignore:"true"`
	ErrorMessage string `json:"error_message"`
}

func NewResponseFormatter(logger log.Logger) ResponseFormatter {
	return ResponseFormatter{
		logger: logger,
	}
}
