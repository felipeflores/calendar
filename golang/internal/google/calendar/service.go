package calendar

import (
	"context"
	"fmt"
	"time"

	"iot/pkg/automation"
	"iot/pkg/google"
	errorsGoogle "iot/pkg/google/errors"
	"iot/pkg/mqtt"
)

type CalendarService struct {
	calendarGoogle    *google.CalendarGoogle
	automationService *automation.AutomationService
	mqttClient        *mqtt.MqttService
	event             mqtt.MqttEvent
}

func NewCalendarService(cg *google.CalendarGoogle, mqttClient *mqtt.MqttService, event mqtt.MqttEvent) *CalendarService {
	return &CalendarService{
		calendarGoogle:    cg,
		automationService: automation.NewAutomationService(),
		mqttClient:        mqttClient,
		event:             event,
	}
}

func (cs *CalendarService) GetEvents() {
	ctx := context.Background()
	currentEvent := ""
	for {
		time.Sleep(10 * time.Second)

		currentEvent = cs.GetEvent(ctx, currentEvent)
	}
}

func (cs *CalendarService) GetEvent(ctx context.Context, currentEvent string) string {
	fmt.Println(fmt.Sprintf("Entrou nos events current=%s", currentEvent))
	event, err := cs.calendarGoogle.Get(ctx)
	if err != nil {
		_, ok := err.(*errorsGoogle.ErrNotConfigured)
		if !ok {
			fmt.Printf("errou %s", err)
			cs.mqttClient.Publish(cs.event.Calendar, "off")
		} else {
			return currentEvent
		}
	}
	fmt.Println("start", event.End.DateTime)
	fmt.Println("start", event.End.TimeZone)
	fmt.Println("end", event.End.DateTime)
	fmt.Println("end", event.End.TimeZone)
	fmt.Println(fmt.Sprintf("Evento %s", event.Summary))

	loc, err := time.LoadLocation(event.Start.TimeZone)
	if err != nil {
		fmt.Println(err)
		cs.mqttClient.Publish(cs.event.Calendar, "off")
	}

	endTime, err := time.ParseInLocation(time.RFC3339, event.End.DateTime, loc)
	if err != nil {
		fmt.Println(err)
		cs.mqttClient.Publish(cs.event.Calendar, "off")
	}
	startTime, err := time.ParseInLocation(time.RFC3339, event.Start.DateTime, loc)
	if err != nil {
		fmt.Println(err)
		cs.mqttClient.Publish(cs.event.Calendar, "off")
	}
	now := time.Now().In(loc)

	if now.After(startTime) && now.Before(endTime) {
		fmt.Println(fmt.Sprintf("Encontrou no meio do esquema current=%s evento=%s", currentEvent, event.Summary))
		if currentEvent != event.Summary {
			currentEvent = event.Summary
			fmt.Printf("Existe %s", cs.event)
			cs.mqttClient.Publish(cs.event.Calendar, "on")
		}

		// cs.powerOn(ctx)
	} else {
		// cs.powerOff(ctx)
		if currentEvent != "" {
			currentEvent = ""
			cs.mqttClient.Publish(cs.event.Calendar, "off")
		}
	}
	return currentEvent
}
