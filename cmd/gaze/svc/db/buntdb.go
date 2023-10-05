package db

import (
	"github.com/ben-dow/Gaze/cmd/gaze/svc/config"
	"github.com/tidwall/buntdb"
)

var gazedb *buntdb.DB

func InitializeDatabase() error {
	if gazedb == nil {
		db, err := buntdb.Open(config.GetConfiguration().DatabaseLocation)
		if err != nil {
			return err
		}
		gazedb = db
	}

	return nil
}
