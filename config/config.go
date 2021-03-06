package config

import "github.com/kelseyhightower/envconfig"

// BoltConfig represents a on-disk key/value store
// https://github.com/boltdb/bolt
type BoltConfig struct {
	// file path for database file
	Path string `envconfig:"bolt_database_path" default:"/tmp/scheduler.db"`
}

// DatabaseConfig defines config settings for database
type DatabaseConfig struct {
	// The database engine used by the backend
	// possible values are: bolt, memory
	Engine string `envconfig:"database_engine" default:"bolt"`
	Bolt   *BoltConfig
}

// Config represents overall configuration objects of the application
type Config struct {
	Debug bool `envconfig:"debug" default:"false"`
	// Possible values "debug", "info", "warn", "error", "fatal", "panic"
	LogLevel string `envconfig:"loglevel" default:"info"`
	// the length of all randomly generated ids
	IDLength int    `envconfig:"id_length" default:"5"`
	URL      string `envconfig:"url" default:"localhost:8080"`
	Frontend bool   `envconfig:"frontend" default:"true"`
	Database *DatabaseConfig
}

// FromEnv parses environment variables into the config object based on envconfig name
func FromEnv() (*Config, error) {
	var config Config
	// We don't use the prefix field here because it doesn't work for nested structs
	err := envconfig.Process("scheduler", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
