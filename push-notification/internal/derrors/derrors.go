package derrors

import (
	"errors"
	"net/http"
)

var (
	ErrValidation     = errors.New("inputs are not valid")
	ErrUnhandled      = errors.New("unhandled error")
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequestBody = errors.New("bad request body")

	ErrNotFound = errors.New("not found")
)

var codes = map[error]int{
	ErrValidation:     http.StatusUnprocessableEntity,
	ErrUnhandled:      599,
	ErrInternalServer: http.StatusInternalServerError,
	ErrBadRequestBody: http.StatusBadRequest,
	ErrNotFound:       http.StatusNotFound,
}

var status = map[error]int{
	ErrValidation:     http.StatusUnprocessableEntity,
	ErrUnhandled:      http.StatusInternalServerError,
	ErrInternalServer: http.StatusInternalServerError,
	ErrBadRequestBody: http.StatusBadRequest,
	ErrNotFound:       http.StatusNotFound,
}

func ToCode(e error) int {
	return codes[e]
}
func ToStatus(e error) int {
	status := status[e]
	// status not defined
	if status == 0 {
		return 499
	}
	return status
}
