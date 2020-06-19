package storage

import "github.com/clintjedwards/scheduler/proto"

// ScheduleMap is a simple implementation of an ordered map.
// The order field contains keys where the last key inserted is added to the back of the slice.
// TODO(clintjedwards): Make the order sorted by schedule date instead of insertion order
type ScheduleMap struct {
	Schedules map[string]*proto.Schedule
	Order     []string
}

// Set inserts an element in the map in order
func (sch *ScheduleMap) Set(key string, value *proto.Schedule) {
	sch.Schedules[key] = value
	sch.Order = append(sch.Order, key)
}
