package model

type MqttRequest struct {
	Info           Chip   `json:"info"`
	Name           string `json:"name"`
	Broker         string `json:"broker"`
	Port           string `json:"port"`
	EventsCalendar string `json:"eventsCalendar"`
}

type Chip struct {
	ChipID string `json:"chip_id"`
}
