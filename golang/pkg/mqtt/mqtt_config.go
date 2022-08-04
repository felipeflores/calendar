package mqtt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Mqtt struct {
	Broker   string    `json:"broker"`
	Port     int       `json:"port"`
	ClientID string    `json:"client_id"`
	Event    MqttEvent `json:"events"`
}
type MqttEvent struct {
	Calendar string `json:"calendar"`
}

func LoadConfig(configFile string) (*Mqtt, error) {
	fmt.Println(configFile)
	var cfg Mqtt
	if err := loadConfigFile(configFile, &cfg); err != nil {
		return nil, err
	}

	fmt.Println("aqui agora o q tem", cfg)
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadConfigFile(configFile string, cfg *Mqtt) error {
	fmt.Println("Vem", configFile)
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	ymlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("errou %v", err)
		return err
	}
	fmt.Println("Nao deu erro")
	return json.Unmarshal(ymlFile, &cfg)
}
