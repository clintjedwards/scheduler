package memory

import (
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	go_proto "github.com/golang/protobuf/proto"
)

// GetSchedule returns a single schedule by id
func (db *Memory) GetSchedule(id string) (*proto.Schedule, error) {

	var storedSchedule proto.Schedule

	rawSchedule, ok := db.store[storage.SchedulesBucket][id]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := go_proto.Unmarshal(rawSchedule, &storedSchedule)
	if err != nil {
		return nil, err
	}

	return &storedSchedule, err
}

// AddSchedule stores a new schedule
func (db *Memory) AddSchedule(id string, schedule *proto.Schedule) error {
	_, ok := db.store[storage.SchedulesBucket][id]
	if ok {
		return utils.ErrEntityExists
	}

	scheduleRaw, err := go_proto.Marshal(schedule)
	if err != nil {
		return err
	}

	db.store[storage.SchedulesBucket][id] = scheduleRaw
	return nil
}

// UpdateSchedule alters schedule infromation
func (db *Memory) UpdateSchedule(id string, schedule *proto.Schedule) error {
	_, ok := db.store[storage.SchedulesBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	scheduleRaw, err := go_proto.Marshal(schedule)
	if err != nil {
		return err
	}

	db.store[storage.SchedulesBucket][id] = scheduleRaw
	return nil
}

// DeleteSchedule removes a schedule from the database
func (db *Memory) DeleteSchedule(id string) error {
	_, ok := db.store[storage.SchedulesBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	db.store[storage.SchedulesBucket][id] = nil
	return nil
}
