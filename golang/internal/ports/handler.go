package admin

import (
	"iot/pkg/serial"
	"iot/rest"
	"net/http"
)

type PortHandler struct {
}

func NewPortHandler() *PortHandler {

	return &PortHandler{}
}

func (h *PortHandler) GetPorts(w http.ResponseWriter, r *http.Request) error {
	ports, err := serial.GetPorts()
	if err != nil {
		return err
	}
	return rest.SendJSON(w, ports)
}
