package models

import (
	"server/internal/database"
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	ID          uuid.UUID `json:"id"`
	Time        time.Time `json:"time"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
}

func DatabaseMeasurementToMeasurement(m database.Measurement) Measurement {
	return Measurement{
		ID:          m.ID,
		Time:        m.MeasurementTime,
		Humidity:    m.Humidity,
		Temperature: m.Temperature,
	}
}

func DatabaseMeasurementsToMeasurements(ms []database.Measurement) []Measurement {
	res := []Measurement{}
	for _, m := range ms {
		res = append(res, DatabaseMeasurementToMeasurement(m))
	}
	return res
}
