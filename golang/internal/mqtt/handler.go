package esp

import (
	"net/http"

	"iot/internal/mqtt/model"
	"iot/internal/mqtt/service"
	"iot/rest"
)

type MqttHandler struct {
	ms service.MqttService
}

func NewMqttHandler(ms service.MqttService) *MqttHandler {
	return &MqttHandler{
		ms: ms,
	}
}

func (h *MqttHandler) StartConfigMqtt(w http.ResponseWriter, r *http.Request) error {

	mqtt := model.MqttRequest{}
	err := rest.DeserializeJSON(r, &mqtt)
	if err != nil {
		return err
	}

	err = h.ms.StartConfigMqtt(mqtt)
	if err != nil {
		return err
	}

	return nil
}
