package memory

import "github.com/clintjedwards/scheduler/storage"

// Memory stores the datastructures that enable the in-memory datastructure
type Memory struct {
	store map[storage.Bucket]map[string][]byte
}

// Init creates a new in-memory database with given settings
func Init() (Memory, error) {
	db := Memory{
		store: map[storage.Bucket]map[string][]byte{
			storage.EmployeesBucket: {},
			storage.SchedulesBucket: {
				storage.SchedulesOrderKey: []byte{},
			},
			storage.PositionsBucket: {},
		},
	}

	return db, nil
}
