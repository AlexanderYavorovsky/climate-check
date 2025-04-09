package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/database"
	"server/internal/jsonutils"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) PostMeasurement(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Humidity    float64 `json:"humidity"`
		Temperature float64 `json:"temperature"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		jsonutils.RespondWithError(w, 400, fmt.Sprintf("error passing JSON: %v", err))
		return
	}

	measurement, err := apiCfg.DB.CreateMeasurement(r.Context(), database.CreateMeasurementParams{
		ID:              uuid.New(),
		MeasurementTime: time.Now().UTC(),
		Humidity:        params.Humidity,
		Temperature:     params.Temperature,
	})
	if err != nil {
		jsonutils.RespondWithError(w, 400, fmt.Sprintf("Couldn't post measurement: %v", err))
		return
	}

	jsonutils.RespondWithJSON(w, 200, models.DatabaseMeasurementToMeasurement(measurement))
}

func (apiCfg *ApiConfig) GetMeasurements(w http.ResponseWriter, r *http.Request) {
	measurements, err := apiCfg.DB.GetMeasurements(r.Context())
	if err != nil {
		jsonutils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get measurements: %v", err))
		return
	}
	jsonutils.RespondWithJSON(w, 201, models.DatabaseMeasurementsToMeasurements(measurements))
}
