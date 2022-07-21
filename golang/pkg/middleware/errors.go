package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

type (
	badrequest interface {
		BadRequest() bool
	}
	notfound interface {
		NotFound() bool
	}
)

func httpStatusCode(err error) int {
	fmt.Println(err)
	switch err.(type) {
	case badrequest:
		return http.StatusBadRequest
	case notfound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
