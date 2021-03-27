package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"time"
)

type Configuration struct {
	PollingInterval time.Duration `env:"POLLING_INTERVAL,default=5m"`

	DataDog
	LaserEgg
	Mqtt
}

type DataDog struct {
	ApiKey string `env:"DD_API_KEY"`
	AppKey string `env:"DD_APP_KEY"`
}

type LaserEgg struct {
	Id  string `env:"LASEREGG_ID"`
	Key string `env:"LASEREGG_KEY"`
}

type Mqtt struct {
	Hostname string `env:"MQTT_HOSTNAME"`
	Port	 int16 	`env:"MQTT_PORT"`
	Username string `env:"MQTT_USERNAME"`
	Password string `env:"MQTT_PASSWORD"`
	Topic	 string `env:"MQTT_TOPIC,default=laseregg"`
	Debug    bool	`env:"MQTT_DEBUG,default=false"`
}

func Load() Configuration {
	var config Configuration
	if err := envdecode.StrictDecode(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
