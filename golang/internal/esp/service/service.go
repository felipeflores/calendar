package service

import (
	"fmt"
	"strings"
	"time"

	"iot/pkg/serial"

	"iot/internal/esp/model"
)

type Control string

const (
	START_WIFI Control = "START_WIFI"
	END_WIFI   Control = "END_WIFI"
)

type EspService struct {
	serialService serial.SerialService
}

func NewEspService(serialService serial.SerialService) *EspService {
	return &EspService{
		serialService: serialService,
	}
}

func (es *EspService) StartEsp(port string) error {
	err := es.serialService.Start(port)
	if err != nil {
		return err
	}
	return nil
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func (es *EspService) GetNetworks() (model.Wifi, error) {
	fmt.Println("GetNetworks")
	err := es.serialService.ReadCommand("networks")
	if err != nil {
		return model.Wifi{}, err
	}
	var wifi model.Wifi
	for {

		buf := strings.Clone(es.serialService.GetBuffer())
		if strings.Contains(buf, string(START_WIFI)) && strings.Contains(buf, string(END_WIFI)) {
			indexStart := strings.Index(buf, string(START_WIFI))
			indexEnd := strings.Index(buf, string(END_WIFI))

			wifi = model.Wifi{
				IndexStart: indexStart,
				IndexEnd:   indexEnd,
				Valor:      substr(buf, indexStart, indexEnd),
				Output:     buf,
			}

			startBuf := substr(buf, 0, indexStart)
			endBuf := substr(buf, indexEnd+8, len(buf))

			var sb = strings.Builder{}
			sb.WriteString(startBuf)
			sb.WriteString(endBuf)

			es.serialService.SetBuffer("")

			break
		}
		time.Sleep(time.Second * 5)
	}

	return wifi, nil
}
