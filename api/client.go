package api

import (
	"encoding/json"
	"fmt"
	"github.com/cthulhu666/laser-egg/cmd/config"
	"io/ioutil"
	"net/http"
	"time"
)

type Measurement struct {
	Id   string `json:"id"`
	Info 		`json:"info.aqi"`
}

type Info struct {
	Ts   time.Time 	`json:"ts"`
	Data 			`json:"data"`
}

type Data struct {
	Pm10 uint `json:"pm10"`
	Pm25 uint `json:"pm25"`
}

func GetMeasurement(cfg config.LaserEgg) (m Measurement, err error) {
	url := fmt.Sprintf("https://api.kaiterra.com/v1/lasereggs/%s?key=%s", cfg.Id, cfg.Key)
	resp, err := http.Get(url)
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return m, err
	}
	if err := json.Unmarshal(body, &m); err != nil {
		return m, err
	}
	return m, err
}
