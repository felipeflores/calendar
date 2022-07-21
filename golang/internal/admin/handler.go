package admin

import (
	"fmt"
	"iot/pkg/google"
	"net/http"
)

type AdminHandler struct {
	calendarGoogle *google.CalendarGoogle
}

func NewAdminHandler(cg *google.CalendarGoogle) *AdminHandler {
	return &AdminHandler{
		calendarGoogle: cg,
	}
}

func (h *AdminHandler) Get(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	h.calendarGoogle.Get(ctx)
	// err =
	fmt.Println("Passou aqui222")

	// return ferrors.NewBadRequest(errors.New("Errou aqui"))
	return nil
}
