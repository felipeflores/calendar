package service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"iot/pkg/serial"

	"iot/internal/esp/model"
)

type Control string

const (
	START_WIFI Control = "START_WIFI"
	END_WIFI   Control = "END_WIFI"
	START_INFO Control = "START_INFO"
	END_INFO   Control = "END_INFO"
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

func (es *EspService) Reset() error {
	err := es.serialService.ReadCommand("reset")
	if err != nil {
		return err
	}
	return nil
}

func (es *EspService) GetNetworks() (model.Wifi, error) {
	fmt.Println("GetNetworks")
	err := es.serialService.ReadCommand("networks")
	if err != nil {
		return model.Wifi{}, err
	}
	wifi := model.Wifi{}
	for {

		buf := strings.Clone(es.serialService.GetBuffer())
		if strings.Contains(buf, string(START_WIFI)) && strings.Contains(buf, string(END_WIFI)) {
			indexStart := strings.Index(buf, string(START_WIFI))
			indexEnd := strings.Index(buf, string(END_WIFI))

			buf = (substr(buf, indexStart-4, indexEnd-10))

			x := strings.Replace(buf, string(END_WIFI), "", 1)
			x = strings.Replace(x, string(START_WIFI), "", 1)

			re := regexp.MustCompile(`\r?\n`)
			x = re.ReplaceAllString(x, " ")

			fmt.Println(x)

			in := []byte(x)

			err := json.Unmarshal(in, &wifi)
			if err != nil {
				fmt.Println("errou")
				return model.Wifi{}, err
			}

			// wifi = model.Wifi{|
			// 	IndexStart: indexStart,
			// 	IndexEnd:   indexEnd,
			// 	Valor:      substr(buf, indexStart, indexEnd),
			// 	Output:     buf,
			// }

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

func (es *EspService) GetInfo() (model.Info, error) {
	err := es.serialService.ReadCommand("info")
	if err != nil {
		return model.Info{}, err
	}
	info := model.Info{}
	for {

		buf := strings.Clone(es.serialService.GetBuffer())
		if strings.Contains(buf, string(START_INFO)) && strings.Contains(buf, string(END_INFO)) {
			re := regexp.MustCompile(`\r?\n`)
			x := re.ReplaceAllString(buf, " ")

			re = regexp.MustCompile(`\\`)
			x = re.ReplaceAllString(x, " ")

			indexStart := strings.Index(x, string(START_INFO))
			indexEnd := strings.Index(x, string(END_INFO))

			x = (substr(x, indexStart, indexEnd-indexStart))

			x = strings.Replace(x, string(START_INFO), "", 1)
			x = strings.Replace(x, string(END_INFO), "", 1)

			in := []byte(x)

			err := json.Unmarshal(in, &info)
			if err != nil {
				return model.Info{}, err
			}

			es.serialService.SetBuffer("")

			break
		}
		time.Sleep(time.Second * 5)
	}

	return info, nil
}
