package memory

import (
	"encoding/json"

	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/clintjedwards/toolkit/listutil"
	"github.com/rs/zerolog/log"
)

// GetAllSchedules returns an unpaginated list of current schedules
func (db *Memory) GetAllSchedules() (schedules *storage.ScheduleMap, err error) {

	schedules = &storage.ScheduleMap{
		Schedules: map[string]*models.Schedule{},
		Order:     []string{},
	}

	for id, rawSchedule := range db.store[storage.SchedulesBucket] {

		if id == storage.SchedulesOrderKey {
			var order []string
			if len(rawSchedule) == 0 {
				schedules.Order = order
				continue
			}
			err := json.Unmarshal(rawSchedule, &order)
			if err != nil {
				log.Error().Err(err).Msg("could not unmarshal order object")
				return nil, err
			}
			schedules.Order = order
			continue
		}

		var schedule models.Schedule

		err := json.Unmarshal(rawSchedule, &schedule)
		if err != nil {
			log.Error().Err(err).Str("id", string(id)).Msg("could not unmarshal database object")
			return nil, err
		}

		schedules.Schedules[id] = &schedule
	}

	return schedules, nil
}

// GetSchedule returns a single schedule by id
func (db *Memory) GetSchedule(id string) (*models.Schedule, error) {

	var storedSchedule models.Schedule

	rawSchedule, ok := db.store[storage.SchedulesBucket][id]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := json.Unmarshal(rawSchedule, &storedSchedule)
	if err != nil {
		return nil, err
	}

	return &storedSchedule, err
}

// AddSchedule stores a new schedule
func (db *Memory) AddSchedule(id string, schedule *models.Schedule) error {
	if _, ok := db.store[storage.SchedulesBucket][id]; ok {
		return utils.ErrEntityExists
	}

	scheduleRaw, err := json.Marshal(schedule)
	if err != nil {
		return err
	}

	var order []string
	rawOrder := db.store[storage.SchedulesBucket][storage.SchedulesOrderKey]
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

	db.store[storage.SchedulesBucket][id] = scheduleRaw
	db.store[storage.SchedulesBucket][storage.SchedulesOrderKey] = rawOrder
	return nil
}

// UpdateSchedule alters schedule infromation
func (db *Memory) UpdateSchedule(id string, schedule *models.Schedule) error {
	_, ok := db.store[storage.SchedulesBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	scheduleRaw, err := json.Marshal(schedule)
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

	var order []string
	rawOrder := db.store[storage.SchedulesBucket][storage.SchedulesOrderKey]
	if len(rawOrder) != 0 {
		err := json.Unmarshal(rawOrder, &order)
		if err != nil {
			log.Error().Err(err).Msg("could not unmarshal order object")
			return err
		}
	}

	order = listutil.RemoveStringFromList(order, id)
	rawOrder, err := json.Marshal(order)
	if err != nil {
		log.Error().Err(err).Msg("could not marshal order object")
		return err
	}

	db.store[storage.SchedulesBucket][id] = nil
	db.store[storage.SchedulesBucket][storage.SchedulesOrderKey] = rawOrder
	return nil
}
