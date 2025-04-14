package main

import (
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
}
