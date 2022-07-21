package admin

import (
	"iot/pkg/google"
	"iot/pkg/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Config struct {
	Logger         *zap.Logger
	Router         *mux.Router
	Middleware     *middleware.Middleware
	CalendarGoogle *google.CalendarGoogle
}

func NewAdminApi(c Config) {
	handler := NewAdminHandler(c.CalendarGoogle)
	SetRoutes(handler, c.Router, c.Middleware)
}

func SetRoutes(
	handler *AdminHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	r := router.PathPrefix("/v1/admin").Subrouter()

	r.Handle(
		"/test",
		handlers.CompressHandler(
			mw.HandlerError(handler.Get),
		),
	).Methods(http.MethodGet)
}
