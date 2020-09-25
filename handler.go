package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

// DataPacket is a struct for the packet sent to the browser
type DataPacket struct {
	Humidity    []float64 `json:"humidity_data"`
	Temperature []float64 `json:"temperature_data"`
	Light       []float64 `json:"light_data"`
	Moisture1   []float64 `json:"moisture1_data"`
	Moisture2   []float64 `json:"moisture2_data"`
	Hum         float64   `json:"humidity_current"`
	Temp        float64   `json:"temperature_current"`
	Lux         float64   `json:"light_current"`
	M1          float64   `json:"m1_current"`
	M2          float64   `json:"m2_current"`
}

func queueToData(queue *DataQueue) ([]byte, error) {
	dp := DataPacket{}
	for row := queue.Queue.Front(); row != nil; row = row.Next() {
		r := row.Value.(*DataRow)
		dp.Humidity = append(dp.Humidity, r.Humidity)
		dp.Temperature = append(dp.Temperature, r.Temperature)
		dp.Light = append(dp.Light, r.Light)
		dp.Moisture1 = append(dp.Moisture1, r.Moisture1)
		dp.Moisture2 = append(dp.Moisture2, r.Moisture2)
	}
	last := queue.Queue.Back()
	l := last.Value.(*DataRow)
	dp.Hum = l.Humidity
	dp.Temp = l.Temperature
	dp.Lux = l.Light
	dp.M1 = l.Moisture1
	dp.M2 = l.Moisture2
	packet, err := json.Marshal(dp)
	if err != nil {
		return nil, err
	}
	return packet, nil
}

func handler(queue DataQueue) http.HandlerFunc {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Failed to parse index.html")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := queueToData(&queue)
		dataStr := string(data)
		if err != nil {
			log.Printf("Failed to marshal data: %v\n", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		t.Execute(w, dataStr)
	}
}
