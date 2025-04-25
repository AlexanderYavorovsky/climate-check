package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"client/internal/histter"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Measurement struct {
	ID          uuid.UUID `json:"id"`
	Time        time.Time `json:"time"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
}

func getMeasurements(serverIP, serverPort string) []Measurement {
	//TODO: use context with timeout
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/measurements", serverIP, serverPort))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

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
	humidity := make([]float64, len(measurements))
	for i, m := range measurements {
		humidity[i] = m.Humidity
	}
	return humidity
}

func getTemperature(measurements []Measurement) []float64 {
	temperature := make([]float64, len(measurements))
	for i, m := range measurements {
		temperature[i] = m.Temperature
	}
	return temperature
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	serverIP := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")

	lastNMeasurements := 10
	measurements := getMeasurements(serverIP, serverPort)
	humidity := getHumidity(measurements[len(measurements)-lastNMeasurements:])
	fmt.Println("Humidity, %: ", humidity)
	temperature := getTemperature(measurements[len(measurements)-lastNMeasurements:])
	histter.PrintHistogram(histter.MakeHistogram(humidity, '.', 10, 100))
	fmt.Println("Temperature, C: ", temperature)
	histter.PrintHistogram(histter.MakeHistogram(temperature, '.', 10, 30))

	// TODO: limit by date/...
}
