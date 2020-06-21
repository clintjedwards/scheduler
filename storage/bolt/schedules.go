package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
)

// GetSchedule returns a single schedule by id
func (db *Bolt) GetSchedule(id string) (*models.Schedule, error) {

	var storedSchedule models.Schedule

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulesBucket))

		scheduleRaw := bucket.Get([]byte(id))
		if scheduleRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := json.Unmarshal(scheduleRaw, &storedSchedule)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedSchedule, nil
}

// AddSchedule stores a new schedule
func (db *Bolt) AddSchedule(id string, schedule *models.Schedule) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulesBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		scheduleRaw, err := json.Marshal(schedule)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), scheduleRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateSchedule alters schedule infromation
func (db *Bolt) UpdateSchedule(id string, schedule *models.Schedule) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulesBucket))

		// First check if key exists
		currentSchedule := bucket.Get([]byte(id))
		if currentSchedule == nil {
			return utils.ErrEntityNotFound
		}

		scheduleRaw, err := json.Marshal(schedule)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), scheduleRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteSchedule removes a schedule from the database
func (db *Bolt) DeleteSchedule(id string) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulesBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists == nil {
			return utils.ErrEntityNotFound
		}

		err := bucket.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
