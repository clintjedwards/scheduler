package storage

import (
	"github.com/clintjedwards/scheduler/proto"
)

// Bucket represents the name of a section of key/value pairs
// usually a grouping of some sort
// ex. A key/value pair of userid-userdata would belong in the users bucket
type Bucket string

const (
	// EmployeesBucket represents the container in which employees are managed
	EmployeesBucket Bucket = "employees"
	// SchedulesBucket represents the container in which schedules are managed
	SchedulesBucket Bucket = "schedules"
	// SchedulerSettingsBucket represents the container in which scheduler settings are managed
	SchedulerSettingsBucket Bucket = "schedulersettings"
)

const (
	// SettingsKey is a root database key that stores app settings in the SchedulerSettings bucket
	SettingsKey = "settings"
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
	GetAllEmployees() (map[string]*proto.Employee, error)
	GetEmployee(id string) (*proto.Employee, error)
	AddEmployee(id string, employee *proto.Employee) error
	UpdateEmployee(id string, employee *proto.Employee) error
	DeleteEmployee(id string) error
	GetSchedulerSettings() (*proto.SchedulerSettings, error)
	UpdateSchedulerSettings(settings *proto.SchedulerSettings) error
	GetSchedule(id string) (*proto.Schedule, error)
	AddSchedule(id string, schedule *proto.Schedule) error
	UpdateSchedule(id string, schedule *proto.Schedule) error
	DeleteSchedule(id string) error
}
