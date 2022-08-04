package google

import (
	"fmt"
	"net/http"

	"iot/pkg/google"
	googleModel "iot/pkg/google/model"

	"iot/internal/google/calendar"
	"iot/internal/google/signup"
	"iot/rest"
)

type GoogleHandler struct {
	CalendarService *calendar.CalendarService
	SignupService   *signup.SignupService
}

func NewGoogleHandler(
	cg *google.CalendarGoogle,
	calendarService *calendar.CalendarService,
	signupService *signup.SignupService,
) *GoogleHandler {
	return &GoogleHandler{
		CalendarService: calendarService,
		SignupService:   signupService,
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

func (h *GoogleHandler) SetCredentials(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	credentials := googleModel.Credentials{}
	err := rest.DeserializeJSON(r, &credentials)
	if err != nil {
		return err
	}

	url, err := h.SignupService.SetCredentials(ctx, credentials)
	if err != nil {
		return err
	}

	return rest.SendJSON(w, signup.Auth{URL: url})
}

func (h *GoogleHandler) SetCode(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	code := signup.Code{}
	err := rest.DeserializeJSON(r, &code)
	if err != nil {
		return err
	}

	return h.SignupService.SetCode(ctx, code)
}
