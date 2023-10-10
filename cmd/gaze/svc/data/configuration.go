package data

import (
	"errors"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/db"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"github.com/tidwall/buntdb"
)

const dbVersion = 1

type Configuration struct {
	DatabaseVersion int
}

func InitializeDataConfiguration() error {

	// Check For the Configuration Entry, if it doesn't exist new database initialization
	_, err := db.GetItemString("gaze_configuration")
	if err != nil && !errors.Is(err, buntdb.ErrNotFound) {
		logging.Debug("an unexpected error occured")
		return err
	}

	if err != nil && errors.Is(err, buntdb.ErrNotFound) {
		logging.Debug("Configuration Entry does not exist, new installation of gaze")
		return newGazeInstall()
	}

	logging.Debug("Database Is Pre-Configured")

	return nil
}

func newGazeInstall() error {
	// Create Indexes

	// User Index
	logging.Trace("creating `users` index")
	err := db.CreateIndex("users")
	if err != nil {
		return err
	}

	return nil
}
