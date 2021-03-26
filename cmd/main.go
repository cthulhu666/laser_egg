package main

import (
	"github.com/cthulhu666/laser-egg/api"
	"github.com/cthulhu666/laser-egg/cmd/config"
	"github.com/cthulhu666/laser-egg/datadog"
	"github.com/cthulhu666/laser-egg/mqtt"
	"log"
	"sync"
	"time"
)

type application struct {
	cfg     config.Configuration
	datadog datadog.Datadog

	datadogCh chan api.Measurement
	mqttCh    chan api.Measurement

	lastMeasurement api.Measurement
}

func main() {
	cfg := config.Load()

	dd, err := datadog.New(cfg.DataDog)
	if err != nil {
		log.Panic(err)
	}

	app := application{
		cfg:       cfg,
		datadog:   dd,
		datadogCh: make(chan api.Measurement, 1),
		mqttCh:    make(chan api.Measurement, 1),
	}

	ticker := time.NewTicker(cfg.PollingInterval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := app.update(); err != nil {
					log.Printf("[Main] failed to update: %v", err)
					continue
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	var wg sync.WaitGroup

	goWithWg(func() error {
		return datadog.HandleMeasurement(app.datadog, app.datadogCh)
	}, "Datadog", &wg)

	goWithWg(func() error {
		return mqtt.HandleMeasurement(nil, app.mqttCh)
	}, "MQTT", &wg)

	wg.Wait()
}

func (app *application) update() error {
	log.Println("[Main] Polling Kaiterra API")
	m, err := api.GetMeasurement(app.cfg.LaserEgg)
	if err != nil {
		return err
	}

	if m.Ts.After(app.lastMeasurement.Ts) {
		app.datadogCh <- m
		app.mqttCh <- m
		app.lastMeasurement = m
	} else {
		log.Println("[Main] Measurement unchanged, skipping.")
	}

	return nil
}

func goWithWg(f func() error, description string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := f(); err != nil {
			log.Printf("CRITICAL: error in background function: %v", err)
		}
		log.Printf("%s exited", description)
	}()
}
