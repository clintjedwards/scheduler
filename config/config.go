package config

import "github.com/kelseyhightower/envconfig"

// BoltConfig represents a on-disk key/value store
// https://github.com/boltdb/bolt
type BoltConfig struct {
	// file path for database file
	Path string `envconfig:"database_path" default:"/tmp/scheduler.db"`
}

// DatabaseConfig defines config settings for database
type DatabaseConfig struct {
	// The database engine used by the backend
	// possible values are: bolt, memory
	Engine string `envconfig:"database_engine" default:"memory"`
	Bolt   *BoltConfig
}

// Config represents overall configuration objects of the application
type Config struct {
	Debug bool `envconfig:"debug" default:"false"`
	// Possible values "debug", "info", "warn", "error", "fatal", "panic"
	LogLevel    string `envconfig:"loglevel" default:"info"`
	TLSCertPath string `envconfig:"tls_cert_path" default:"./localhost.crt"`
	TLSKeyPath  string `envconfig:"tls_key_path" default:"./localhost.key"`
	// the length of all randomly generated ids
	IDLength int    `envconfig:"id_length" default:"5"`
	HTTPURL  string `envconfig:"http_url" default:"localhost:8080"`
	GRPCURL  string `envconfig:"grpc_url" default:"localhost:8081"`
	Database *DatabaseConfig
}

// FromEnv parses environment variables into the config object based on envconfig name
func FromEnv() (*Config, error) {
	var config Config
	err := envconfig.Process("scheduler", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
