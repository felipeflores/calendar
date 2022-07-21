package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service Service  `yaml:"service"`
	Server  Server   `yaml:"server"`
	Cors    corsInfo `yaml:"cors"`
	Mqtt    Mqtt     `yaml:"mqtt"`
}

type Service struct {
	Name string `yaml:"name"`
}

type Server struct {
	Address         string        `yaml:"address"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type corsInfo struct {
	AllowedHeaders []string `yaml:"allowed_headers"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedOrigins []string `yaml:"allowed_origins"`
	ExposedHeaders []string `yaml:"exposed_headers"`
	MaxAge         int      `yaml:"max_age"`
}

type Mqtt struct {
	Broker   string    `yaml:"broker"`
	Port     int       `yaml:"port"`
	ClientID string    `yaml:"client_id"`
	Event    MqttEvent `yaml:"events"`
}
type MqttEvent struct {
	Calendar string `yaml:"calendar"`
}

func LoadConfig(configFile string) (*Config, error) {
	fmt.Println(configFile)
	var cfg Config
	if err := loadConfigFile(configFile, &cfg); err != nil {
		return nil, err
	}

	fmt.Println("aqui agora o q tem", cfg)
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadConfigFile(configFile string, cfg *Config) error {
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
	return yaml.Unmarshal(ymlFile, &cfg)
}
