package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
)

// GetAllSchedules returns an unpaginated list of current links
func (db *Bolt) GetAllSchedules() (schedules *storage.ScheduleMap, err error) {
	schedules = &storage.ScheduleMap{
		Schedules: map[string]*models.Schedule{},
		Order:     []string{},
	}

	db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulesBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			if string(key) == storage.SchedulesOrderKey {
				var order []string
				if len(value) == 0 {
					schedules.Order = order
					return nil
				}
				err := json.Unmarshal(value, &order)
				if err != nil {
					log.Error().Err(err).Msg("could not unmarshal order object")
					return err
				}
				schedules.Order = order
				return nil
			}

			var schedule models.Schedule

			err := json.Unmarshal(value, &schedule)
			if err != nil {
				log.Error().Err(err).Str("id", string(key)).Msg("could not unmarshal database object")
				// We don't return an error here so that we can at least return a partial list
				return err
			}

			schedules.Schedules[string(key)] = &schedule
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	return schedules, nil
}

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

		var order []string
		rawOrder := bucket.Get([]byte(storage.SchedulesOrderKey))
		if len(rawOrder) != 0 {
			err = json.Unmarshal(rawOrder, &order)
			if err != nil {
				log.Error().Err(err).Msg("could not unmarshal order object")
				return err
			}
		}

		order = append(order, id)
		rawOrder, err = json.Marshal(order)
		if err != nil {
			log.Error().Err(err).Msg("could not marshal order object")
			return err
		}

		err = bucket.Put([]byte(storage.SchedulesOrderKey), rawOrder)
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
