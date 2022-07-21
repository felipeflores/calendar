package admin

import (
	"iot/pkg/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Config struct {
	Logger     *zap.Logger
	Router     *mux.Router
	Middleware *middleware.Middleware
}

func NewPortApi(c Config) {
	handler := NewPortHandler()
	SetRoutes(handler, c.Router, c.Middleware)
}

func SetRoutes(
	handler *PortHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	r := router.PathPrefix("/v1/ports").Subrouter()

	r.Handle(
		"",
		handlers.CompressHandler(
			mw.HandlerError(handler.GetPorts),
		),
	).Methods(http.MethodGet)
}
