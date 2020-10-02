package storage

import "github.com/clintjedwards/scheduler/model"

// ScheduleMap is a simple implementation of an ordered map.
// The order field contains keys where the last key inserted is added to the back of the slice.
// TODO(clintjedwards): Make the order sorted by schedule date instead of insertion order
type ScheduleMap struct {
	Schedules map[string]*model.Schedule `json:"schedules"`
	Order     []string                   `json:"order"`
}

// Set inserts an element in the map in order
func (sch *ScheduleMap) Set(key string, value *model.Schedule) {
	sch.Schedules[key] = value
	sch.Order = append(sch.Order, key)
}
