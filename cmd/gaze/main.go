package main

import (
	"context"
	"github.com/ben-dow/Gaze/cmd/gaze/api/rest"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/config"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/data"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/db"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	config.InitializeConfiguration()
	logging.Debug("Configuration Loaded")

	err := db.InitializeDatabase()
	if err != nil {
		logging.Error("could not initialize datastore connection %v", err)
		return
	}
	logging.Debug("Database Loaded")

	err = data.InitializeDataConfiguration()
	if err != nil {
		logging.Error("could not initialize data configuration", err)
	}
	logging.Debug("Data Configuration Loaded")

	api := rest.NewRestApi(config.GetConfiguration().ServerAddress)
	logging.Debug("API Initialized")

	logging.Debug("Starting API Server")
	apiServerWg := &sync.WaitGroup{}
	api.Start(apiServerWg)

	logging.Info("Started")
	for {
		select {
		case <-shutdown:
			logging.Info("Shutting Down")

			// Context for Shutdown
			// Shutdowns must complete within 15 seconds
			ctx, cncl := context.WithTimeout(context.Background(), time.Second*15)

			logging.Trace("Stopping API Server")
			err := api.Stop(ctx)
			if err != nil {
				logging.Error("could not shutdown api server. %v", err)
			}

			// Wait for wait groups to exit
			apiServerWg.Wait()
			logging.Trace("Shutdown Complete")

			cncl()
			return
		}
	}
}
