package datadog

import (
	"github.com/cthulhu666/laser-egg/api"
	"log"
)

func HandleMeasurement(dog Datadog, ch <-chan api.Measurement) error {
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				log.Panic("ch closed?")
			}
			log.Printf("[Datadog] sending measurement: %v", m)
			dog.Send(m)
		}
	}
}
