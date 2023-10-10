package db

import (
	"encoding/json"
	"fmt"
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

type GazeDataItem interface {
	GetKey() string
}

type GazeIndexedDataItem interface {
	GazeDataItem
	GetIndex() string
}

func CreateIndex(name string) error {
	patternStr := fmt.Sprintf("%s:*", name)
	return gazedb.CreateIndex(name, patternStr)
}

func StoreItem(item GazeDataItem) error {
	marshaledItem, err := json.Marshal(item)
	if err != nil {
		return err
	}

	itemStr := string(marshaledItem)

	err = gazedb.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(item.GetKey(), itemStr, nil)
		return err
	})

	return err
}

func StoreIndexedItem(item GazeIndexedDataItem) error {
	marshaledItem, err := json.Marshal(item)
	if err != nil {
		return err
	}

	itemStr := string(marshaledItem)

	keyStr := fmt.Sprintf("%s:%s", item.GetIndex(), item.GetKey())
	err = gazedb.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(keyStr, itemStr, nil)
		return err
	})

	return err
}

func GetItem(key string, out interface{}) error {
	val, err := GetItemString(key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), out)

	return err
}

func GetItemString(key string) (string, error) {
	var dbValue string

	err := gazedb.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key, true)
		if err == nil {
			dbValue = val
		}
		return err
	})

	if err != nil {
		return "", err
	}

	return dbValue, nil
}
