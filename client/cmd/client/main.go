package main

import (
	histter "client/internal"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	ID          uuid.UUID `json:"id"`
	Time        time.Time `json:"time"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
}

func getMeasurements(serverIp, serverPort string) []Measurement {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/measurements", serverIp, serverPort))
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var measurements []Measurement
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&measurements)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return measurements
}

func getHumidity(measurements []Measurement) []float64 {
	var humidity []float64
	for _, m := range measurements {
		humidity = append(humidity, m.Humidity)
	}
	return humidity
}

func getTemperature(measurements []Measurement) []float64 {
	var temperature []float64
	for _, m := range measurements {
		temperature = append(temperature, m.Temperature)
	}
	return temperature
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")

	lastNMeasurements := 10
	measurements := getMeasurements(serverIp, serverPort)
	humidity := getHumidity(measurements[len(measurements)-lastNMeasurements:])
	fmt.Println("Humidity, %: ", humidity)
	temperature := getTemperature(measurements[len(measurements)-lastNMeasurements:])
	histter.PrintHistogram(histter.MakeHistogram(humidity, '.', 10, 100))
	fmt.Println("Temperature, C: ", temperature)
	histter.PrintHistogram(histter.MakeHistogram(temperature, '.', 10, 30))

	//TODO: limit by date/...
}
