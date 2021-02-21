package storage

import "github.com/clintjedwards/scheduler/model"

// Bucket represents the name of a section of key/value pairs
// usually a grouping of some sort
// ex. A key/value pair of userid-userdata would belong in the users bucket
type Bucket string

const (
	// EmployeesBucket represents the container in which employees are managed
	EmployeesBucket Bucket = "employees"
	// SchedulesBucket represents the container in which schedules are managed
	SchedulesBucket Bucket = "schedules"
	// PositionsBucket represents the container that holds employment positions
	PositionsBucket Bucket = "positions"
)

const (
	// SchedulesOrderKey is a schedules bucket key that stores the order
	SchedulesOrderKey = "_order"
)

// EngineType represents the different possible storage engines available
type EngineType string

const (
	// BoltEngine represents a bolt storage engine.
	// A file based key-value store.(https://github.com/boltdb/bolt)
	BoltEngine EngineType = "bolt"

	// MemoryEngine represents an in-memory storage engine
	// Used mostly for dev work
	MemoryEngine EngineType = "memory"
)

// Engine represents backend storage implementations where items can be persisted
type Engine interface {
	GetAllEmployees() (map[string]*model.Employee, error)
	GetEmployee(id string) (*model.Employee, error)
	AddEmployee(id string, employee *model.Employee) error
	UpdateEmployee(id string, employee *model.PatchEmployee) error
	DeleteEmployee(id string) error

	GetAllPositions() (map[string]*model.Position, error)
	GetPosition(id string) (*model.Position, error)
	AddPosition(id string, Position *model.Position) error
	UpdatePosition(id string, Position *model.Position) error
	DeletePosition(id string) error

	// GetAllSchedules returns an unpagined map of all schedules with the order they were added
	GetAllSchedules() (schedules *ScheduleMap, err error)
	GetSchedule(id string) (*model.Schedule, error)
	AddSchedule(id string, schedule *model.Schedule) error
	UpdateSchedule(id string, schedule *model.Schedule) error
	DeleteSchedule(id string) error
}
