package main

import (
	histter "client/internal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	ID          uuid.UUID `json:"id"`
	Time        time.Time `json:"time"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
}

func getMeasurements() {
	resp, err := http.Get("http://localhost:8080/measurements")
	if err != nil {
		log.Println(err.Error())
		return
	}

	var measurements []Measurement
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&measurements)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, m := range measurements {
		fmt.Println(m.Humidity, m.Temperature)
	}
}

func main() {
	fmt.Println("TODO")

	arr := []float64{25.4, 19.8, 23.4, 21.0, 18.7, 10.1, 12, 34, 26, 22, 28}
	hist := histter.MakeHistogram(arr, '.', 10)
	for _, s := range hist {
		fmt.Println(s)
	}
}
