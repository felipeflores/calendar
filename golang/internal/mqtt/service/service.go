package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"iot/internal/mqtt/model"
	"iot/pkg/mqtt"
)

type MqttService struct {
	Mqqt *mqtt.MqttService
}

func NewMqttService(mqqt *mqtt.MqttService) *MqttService {
	return &MqttService{
		Mqqt: mqqt,
	}
}

func (ms *MqttService) StartConfigMqtt(m model.MqttRequest) error {
	port, err := strconv.Atoi(m.Port)
	if err != nil {
		return err
	}
	mqttConfig := mqtt.Mqtt{
		Broker:   m.Broker,
		Port:     port,
		ClientID: fmt.Sprintf("%s-%s", m.Name, m.Info.ChipID),
		Event: mqtt.MqttEvent{
			Calendar: fmt.Sprintf("%s/%s", m.Info.ChipID, m.EventsCalendar),
		},
	}

	mqttFile := "mqtt.json"

	file, _ := json.MarshalIndent(mqttConfig, "", " ")
	_ = ioutil.WriteFile(mqttFile, file, 0644)

	return ms.Mqqt.Setup(mqttConfig.ClientID, mqttConfig.Broker, mqttConfig.Port)
}
