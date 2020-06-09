package memory

import (
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	go_proto "github.com/golang/protobuf/proto"
)

// GetSchedulerSettings returns a single schedulersettings by id
func (db *Memory) GetSchedulerSettings() (*proto.SchedulerSettings, error) {
	var storedSchedulerSettings proto.SchedulerSettings

	rawSettings, ok := db.store[storage.SchedulerSettingsBucket][storage.SettingsKey]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := go_proto.Unmarshal(rawSettings, &storedSchedulerSettings)
	if err != nil {
		return nil, err
	}

	return &storedSchedulerSettings, err
}

// UpdateSchedulerSettings stores a new settings
func (db *Memory) UpdateSchedulerSettings(schedulersettings *proto.SchedulerSettings) error {

	rawSettings, err := go_proto.Marshal(schedulersettings)
	if err != nil {
		return err
	}

	db.store[storage.SchedulerSettingsBucket][storage.SettingsKey] = rawSettings
	return nil
}
