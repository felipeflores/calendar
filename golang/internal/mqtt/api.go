package esp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"iot/internal/mqtt/service"
	"iot/pkg/middleware"
	"iot/pkg/mqtt"
)

type Config struct {
	Logger     *zap.Logger
	Router     *mux.Router
	Middleware *middleware.Middleware
	Mqqt       *mqtt.MqttService
}

func NewMqttApi(c Config) {
	handler := NewMqttHandler(*service.NewMqttService(c.Mqqt))
	SetRoutes(handler, c.Router, c.Middleware)
}

func SetRoutes(
	handler *MqttHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	fmt.Println("SetRouteMqtt")
	r := router.PathPrefix("/v1/mqtt").Subrouter()

	r.Handle(
		"",
		handlers.CompressHandler(
			mw.HandlerError(handler.StartConfigMqtt),
		),
	).Methods(http.MethodPost)

}
