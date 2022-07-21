package ferrors

import "net/http"

type ErrBadRequest struct {
	error
	message string
}

func NewBadRequest(err error) *ErrBadRequest {
	return &ErrBadRequest{
		error:   err,
		message: http.StatusText(http.StatusBadRequest),
	}
}

func (*ErrBadRequest) BadRequest() bool {
	return true
}
