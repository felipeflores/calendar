package google

import (
	"fmt"
	"net/http"

	"iot/internal/google/calendar"
	"iot/pkg/google"
)

type GoogleHandler struct {
	CalendarService *calendar.CalendarService
}

func NewGoogleHandler(cg *google.CalendarGoogle, calendarService *calendar.CalendarService) *GoogleHandler {
	return &GoogleHandler{
		CalendarService: calendarService,
	}
}

func (h *GoogleHandler) GetCalendarEvents(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	h.CalendarService.GetEvent(ctx, "")
	// err =
	fmt.Println("Passou aqui222")

	// return ferrors.NewBadRequest(errors.New("Errou aqui"))
	return nil
}
