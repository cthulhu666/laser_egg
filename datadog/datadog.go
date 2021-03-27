package datadog

import (
	"github.com/cthulhu666/laser-egg/cmd/config"
	"github.com/cthulhu666/laser-egg/laseregg"
	"github.com/cthulhu666/laser-egg/target"
	dd "github.com/zorkian/go-datadog-api"
)

type datadog struct {
	*dd.Client
}

func New(cfg config.DataDog) (target.Target, error) {

	return datadog{
		Client: dd.NewClient(cfg.ApiKey, cfg.AppKey),
	}, nil
}

func (d datadog) Send(measurement api.Measurement) error {
	ts := dd.Float64(float64(measurement.Ts.Unix()))

	series := []dd.Metric{
		{
			Metric:   dd.String("air.pm10"),
			Points:   []dd.DataPoint{dd.DataPoint{ts, dd.Float64(float64(measurement.Pm10))}},
			Type:     dd.String("gauge"),
			Host:     nil,
			Tags:     []string{"test:test"},
			Unit:     nil,
			Interval: nil,
		},
		{
			Metric:   dd.String("air.pm25"),
			Points:   []dd.DataPoint{dd.DataPoint{ts, dd.Float64(float64(measurement.Pm25))}},
			Type:     dd.String("gauge"),
			Host:     nil,
			Tags:     []string{"test:test"},
			Unit:     nil,
			Interval: nil,
		},
	}
	err := d.Client.PostMetrics(series)
	if err != nil {
		return err
	}
	return nil
}