package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	internalAdmin "iot/internal/admin"
	"iot/internal/config"
	internalEsp "iot/internal/esp"
	internalGoogle "iot/internal/google"
	"iot/internal/google/calendar"
	internalMqtt "iot/internal/mqtt"
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
	err = clientGoogle.GetClient()
	if err != nil {
		fmt.Println(fmt.Sprintf("Sem token %v", err))
	}
	calendarGoogle := google.NewCalendarGoogle(clientGoogle)

	var mqttConfigFile string

	flag.StringVar(&mqttConfigFile, "mqqtConfig", "mqtt.json", "json mqtt file path")
	flag.Parse()

	mqqtConfig, err := mqtt.LoadConfig(mqttConfigFile)
	if err != nil {
		log.Fatalf("Error on load config file %v", err)
	}
	mqttClient := mqtt.NewMqttService()
	if err != nil {
		panic(err)
	}
	err = mqttClient.Setup(mqqtConfig.ClientID, mqqtConfig.Broker, mqqtConfig.Port)
	if err != nil {
		fmt.Println(fmt.Sprintf("Erro setup mqtt %v", err))
	}
	// Init the mux router
	router := mux.NewRouter()

	rest.SetupRoutes(router)

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

	internalMqtt.NewMqttApi(internalMqtt.Config{
		Router:     router,
		Middleware: mw,
		Mqqt:       mqttClient,
	})

	internalAdmin.NewAdminApi(internalAdmin.Config{
		Router:         router,
		Middleware:     mw,
		CalendarGoogle: calendarGoogle,
	})

	service := calendar.NewCalendarService(calendarGoogle, mqttClient, mqqtConfig.Event)

	go service.GetEvents()

	internalGoogle.NewGoogle(router, mw, calendarGoogle, service, clientGoogle)

	// go httpserver.Start()

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
