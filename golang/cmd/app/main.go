package main

import (
	"flag"
	"log"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	internalAdmin "iot/internal/admin"
	"iot/internal/config"
	internalEsp "iot/internal/esp"
	internalGoogle "iot/internal/google"
	"iot/internal/google/calendar"
	internalPorts "iot/internal/ports"
	"iot/pkg/google"
	"iot/pkg/httpserver"
	"iot/pkg/middleware"
	"iot/pkg/mqtt"
	"iot/pkg/serial"
	"iot/rest"
)

func main() {
	// set environment variable GEEKS
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "386")

	var configFile string

	flag.StringVar(&configFile, "c", "env.yaml", "env file path")
	flag.Parse()

	c, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error on load config file %v", err)
	}

	mw := middleware.New()

	clientGoogle := google.NewClientGoogle()
	calendarGoogle := google.NewCalendarGoogle(clientGoogle)

	mqttClient, err := mqtt.NewMqttService(c.Mqtt.ClientID, c.Mqtt.Broker, c.Mqtt.Port)
	if err != nil {
		panic(err)
	}

	// Init the mux router
	router := mux.NewRouter()

	rest.SetupRoutes(router)

	internalAdmin.NewAdminApi(internalAdmin.Config{
		Router:         router,
		Middleware:     mw,
		CalendarGoogle: calendarGoogle,
	})

	internalPorts.NewPortApi(internalPorts.Config{
		Router:     router,
		Middleware: mw,
	})

	serialService := serial.NewSerialService()

	internalEsp.NewEspApi(internalEsp.Config{
		Router:     router,
		Middleware: mw,
		SerialPort: serialService,
	})

	service := calendar.NewCalendarService(calendarGoogle, mqttClient, c.Mqtt.Event)

	// go service.GetEvents()

	internalGoogle.NewGoogle(router, mw, calendarGoogle, service)

	go httpserver.Start()

	httpserver.Run(
		httpserver.Config{
			Address:         c.Server.Address,
			IdleTimeout:     c.Server.IdleTimeout,
			ReadTimeout:     c.Server.ReadTimeout,
			WriteTimeout:    c.Server.WriteTimeout,
			ShutdownTimeout: c.Server.ShutdownTimeout,
		},
		handlers.RecoveryHandler()(
			handlers.CORS(
				handlers.AllowedHeaders(c.Cors.AllowedHeaders),
				handlers.AllowedMethods(c.Cors.AllowedMethods),
				handlers.ExposedHeaders(c.Cors.ExposedHeaders),
				handlers.AllowedOrigins(c.Cors.AllowedOrigins),
				handlers.MaxAge(c.Cors.MaxAge),
			)(router),
		),
	)

	// err = http.ListenAndServe(c.Server.Address, router)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Subiu")

	// router.PathPrefix("/health").
	// 	Methods(http.MethodGet).
	// 	Handler(handlers.CompressHandler)
}
