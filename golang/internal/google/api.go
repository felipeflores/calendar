package google

import (
	"iot/internal/google/calendar"
	"iot/internal/google/signup"
	"iot/pkg/google"
	"iot/pkg/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewGoogle(router *mux.Router,
	middleware *middleware.Middleware,
	calendarGoogle *google.CalendarGoogle,
	calendarService *calendar.CalendarService,
	clientGoogle *google.ClientGoogle,
) {
	signupService := signup.NewSignupService(clientGoogle)

	handler := NewGoogleHandler(calendarGoogle, calendarService, signupService)
	SetRoutes(handler, router, middleware)
}

func SetRoutes(
	handler *GoogleHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	r := router.PathPrefix("/v1/google").Subrouter()

	r.Handle(
		"/signup",
		handlers.CompressHandler(
			mw.HandlerError(handler.SetCredentials),
		),
	).Methods(http.MethodPost)

	r.Handle(
		"/code",
		handlers.CompressHandler(
			mw.HandlerError(handler.SetCode),
		),
	).Methods(http.MethodPost)

	rCalendar := r.PathPrefix("/calendar").Subrouter()
	rCalendar.Handle(
		"/events",
		handlers.CompressHandler(
			mw.HandlerError(handler.GetCalendarEvents),
		),
	).Methods(http.MethodGet)
}
