package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/storage"
)

//TODO(clintjedwards): make sure bolt is 1-1 with memory package

// Bolt is a representation of the bolt datastore
type Bolt struct {
	store *bolt.DB
}

// Init creates a new boltdb with given settings
func Init(configuration interface{}) (Bolt, error) {
	db := Bolt{}
	conf, ok := configuration.(*config.BoltConfig)
	if !ok {
		return Bolt{}, fmt.Errorf("incorrect config type expected 'config.BoltConfig'; got '%T'", configuration)
	}

	store, err := bolt.Open(conf.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return Bolt{}, err
	}

	// Create root bucket if not exists
	err = store.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.EmployeesBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.SchedulesBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.PositionsBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return Bolt{}, err
	}

	db.store = store

	return db, nil
}
