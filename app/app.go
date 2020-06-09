package app

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/clintjedwards/scheduler/api"
	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/storage/bolt"
	"github.com/clintjedwards/scheduler/storage/memory"
)

// StartServices initializes all required services,the raw GRPC service, and the metrics endpoint
func StartServices() {
	config, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get config in order to start services")
	}

	storage, err := InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		log.Panic().Err(err).Msg("could not init storage")
	}

	log.Info().Str("engine", config.Database.Engine).Msg("storage engine initialized")

	schedulerAPI := api.NewAPI(config, storage)
	grpcServer := api.CreateGRPCServer(schedulerAPI)

	api.InitGRPCService(config, grpcServer)
}

// InitStorage creates a storage object with the appropriate engine
func InitStorage(engineType storage.EngineType) (storage.Engine, error) {

	config, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	switch engineType {
	case storage.BoltEngine:

		boltStorageEngine, err := bolt.Init(config.Database.Bolt)
		if err != nil {
			return nil, err
		}

		return &boltStorageEngine, nil
	case storage.MemoryEngine:

		memoryStorageEngine, err := memory.Init()
		if err != nil {
			return nil, err
		}

		return &memoryStorageEngine, nil
	default:
		return nil, fmt.Errorf("storage backend \"%s\" not implemented", engineType)
	}
}
