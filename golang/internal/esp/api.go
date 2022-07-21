package esp

import (
	"fmt"
	"iot/internal/esp/service"
	"iot/pkg/middleware"
	"iot/pkg/serial"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Config struct {
	Logger     *zap.Logger
	Router     *mux.Router
	Middleware *middleware.Middleware
	SerialPort *serial.SerialService
}

func NewEspApi(c Config) {
	handler := NewEspHandler(*service.NewEspService(*c.SerialPort))
	SetRoutes(handler, c.Router, c.Middleware)
}

func SetRoutes(
	handler *EspHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	fmt.Println("SetRouteESP")
	r := router.PathPrefix("/v1/esp").Subrouter()

	r.Handle(
		"/networks",
		handlers.CompressHandler(
			mw.HandlerError(handler.GetNetworks),
		),
	).Methods(http.MethodGet)

	r.Handle(
		"/{port}/start",
		handlers.CompressHandler(
			mw.HandlerError(handler.StartEsp),
		),
	).Methods(http.MethodPost)

}
