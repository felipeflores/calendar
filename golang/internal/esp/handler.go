package esp

import (
	"fmt"
	"iot/rest"
	"net/http"

	"iot/internal/esp/service"
)

type EspHandler struct {
	es service.EspService
}

func NewEspHandler(es service.EspService) *EspHandler {
	return &EspHandler{
		es: es,
	}
}

func (h *EspHandler) StartEsp(w http.ResponseWriter, r *http.Request) error {
	port, err := rest.GetStr(r, "port")
	if err != nil {
		return err
	}

	err = h.es.StartEsp(port)
	if err != nil {
		return err
	}

	return nil
}

func (h *EspHandler) GetNetworks(w http.ResponseWriter, r *http.Request) error {

	fmt.Println(r.URL.Path)
	wifi, err := h.es.GetNetworks()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Retorno: %v", wifi))
	// return ferrors.NewBadRequest(errors.New("Errou aqui"))
	w.WriteHeader(http.StatusOK)
	return rest.SendJSON(w, wifi)
}
