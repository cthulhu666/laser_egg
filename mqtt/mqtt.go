package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/cthulhu666/laser-egg/cmd/config"
	"github.com/cthulhu666/laser-egg/laseregg"
	"github.com/cthulhu666/laser-egg/target"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

type client struct {
	paho.Client
	topic string
}

func (c client) Send(measurement api.Measurement) error {
	payload, err := json.Marshal(measurement)
	if err != nil {
		return err
	}
	token := c.Client.Publish(c.topic, 0, false, payload)
	ok := token.WaitTimeout(5 * time.Second)
	if !ok {
		return fmt.Errorf("timeout")
	}
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func New(config config.Mqtt) (target.Target, error) {
	if config.Debug {
		paho.DEBUG = log.New(os.Stdout, "[DEBUG]", 0)
		paho.WARN = log.New(os.Stdout, "[WARN]", 0)
		paho.ERROR = log.New(os.Stdout, "[ERROR]", 0)
	}
	opts := paho.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", config.Hostname, config.Port)).
		SetUsername(config.Username).
		SetPassword(config.Password)

	c := paho.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client{c, config.Topic}, nil
}
