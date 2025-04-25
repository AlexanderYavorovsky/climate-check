package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/tarm/serial"
)

const BytesReadCnt = 11

func getHumidityTemperature(data []byte) (float64, float64, error) {
	humidity, err := strconv.ParseFloat(string(data[:5]), 64)
	if err != nil {
		return 0, 0, err
	}
	temperature, err := strconv.ParseFloat(string(data[6:]), 64)
	if err != nil {
		return 0, 0, err
	}
	return humidity, temperature, nil
}

func sendToServer(serverIp, serverPort string, data []byte) {
	humidity, temperature, err := getHumidityTemperature(data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Hum:%v Temp:%v\n", humidity, temperature)

	dataMap := map[string]interface{}{
		"humidity":    humidity,
		"temperature": temperature,
	}
	dataJSON, err := json.Marshal(dataMap)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post("http://"+serverIp+":"+serverPort+"/measurements", "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")

	c := &serial.Config{
		Name:        "/dev/ttyACM0",
		Baud:        1000000,
		ReadTimeout: time.Minute * 2,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 2048)
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n == BytesReadCnt {
			log.Println("Sending to server")
			sendToServer(serverIp, serverPort, buf[:n])
		}
	}
}
