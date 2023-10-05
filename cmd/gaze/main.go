package main

import (
	"github.com/ben-dow/Gaze/cmd/gaze/api/rest"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/config"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/db"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"log"
	"net/http"
)

func main() {
	config.InitializeConfiguration()
	logging.Info("Configuration Loaded")

	db.InitializeDatabase()
	logging.Info("Database Loaded")

	api := rest.NewRestApi()
	logging.Info("API Initialized")

	logging.Info("Starting Server")
	log.Fatal(http.ListenAndServe(":3000", api))
}
