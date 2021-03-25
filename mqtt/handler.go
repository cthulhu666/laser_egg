package mqtt

import (
	"github.com/cthulhu666/laser-egg/api"
	"log"
)

func HandleMeasurement(client interface{}, ch <-chan api.Measurement) error {
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				log.Panic("ch closed?")
			}
			log.Printf("[MQTT] sending measurement: %v", m)
			// TODO: actually send
		}
	}
}
