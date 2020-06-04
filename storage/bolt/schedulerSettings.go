package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	go_proto "github.com/golang/protobuf/proto"
)

//GetSchedulerSettings() (*proto.SchedulerSettings, error)
//UpdateSchedulerSettings(settings proto.SchedulerSettings) error

// GetSchedulerSettings returns a single schedulersettings by id
func (db *Bolt) GetSchedulerSettings(id string) (*proto.SchedulerSettings, error) {

	var storedSchedulerSettings proto.SchedulerSettings

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulerSettingsBucket))

		schedulersettingsRaw := bucket.Get([]byte(storage.SettingsKey))
		if schedulersettingsRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := go_proto.Unmarshal(schedulersettingsRaw, &storedSchedulerSettings)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedSchedulerSettings, nil
}

// UpdateSchedulerSettings stores a new settings
func (db *Bolt) UpdateSchedulerSettings(id string, schedulersettings *proto.SchedulerSettings) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.SchedulerSettingsBucket))

		// First check if key exists
		exists := bucket.Get([]byte(storage.SettingsKey))
		if exists != nil {
			return utils.ErrEntityExists
		}

		schedulersettingsRaw, err := go_proto.Marshal(schedulersettings)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), schedulersettingsRaw)
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