package main

import (
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

func getMeasurements(serverIp, serverPort string) {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/measurements", serverIp, serverPort))
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
		fmt.Println(m.Time.Local(), m.Humidity, m.Temperature)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")

	getMeasurements(serverIp, serverPort)
	//TODO: print histogram
	//TODO: limit by date/...
}
