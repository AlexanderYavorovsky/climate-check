package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"server/internal/database"
	"server/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// init config
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	port := os.Getenv("SERVER_PORT")
	dbURL := os.Getenv("DB_URL")

	// init storage
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := database.New(conn)
	apiCfg := handlers.ApiConfig{
		DB: db,
	}

	// init router
	router := chi.NewRouter()
	router.Get("/healthz", handlers.HandlerHealthz)
	router.Get("/measurements", apiCfg.GetMeasurements)
	router.Post("/measurements", apiCfg.PostMeasurement)

	// run server
	srv := &http.Server{
		Handler:     router,
		Addr:        ":" + port,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Server starting on port %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
