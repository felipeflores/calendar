package serial

import (
	"errors"
	"fmt"
	"iot/pkg/ferrors"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

type SerialService struct {
	serialPort serial.Port
	sb         strings.Builder
}

func NewSerialService() *SerialService {
	return &SerialService{}
}

func GetPorts() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	if len(ports) == 0 {
		return nil, ferrors.NewNotFound(errors.New("Sem portas"))
	}
	return ports, nil
}

func (ss *SerialService) Start(portName string) error {
	fmt.Println("Chamou aqui")
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(portName, mode)
	if err != nil {
		return err
	}

	ss.serialPort = port
	ss.sb = strings.Builder{}
	go ss.SerialMonitor()
	return nil
}

func (ss *SerialService) Close() error {
	fmt.Println("Fechou")
	err := ss.serialPort.Close()
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}

func (ss *SerialService) SerialMonitor() error {
	err := ss.serialPort.SetReadTimeout(time.Second * 10)
	if err != nil {
		return err
	}
	for {
		buff := make([]byte, 100)
		n, err := ss.serialPort.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n > 0 {
			result := string(buff[:n])

			fmt.Print(result)

			ss.sb.WriteString(result)
		}
	}
	return nil
}

func (ss *SerialService) ReadCommand(command string) error {
	fmt.Println("Comando")
	_, err := ss.serialPort.Write([]byte(fmt.Sprintf("%s\r", command)))
	if err != nil {
		return err
	}
	return nil
}

func (ss *SerialService) GetBuffer() string {
	return ss.sb.String()
}

func (ss *SerialService) SetBuffer(data string) {
	fmt.Println("Clear")
	ss.sb.Reset()
	// ss.sb.WriteString(data)
}
