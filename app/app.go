package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/clintjedwards/scheduler/api"
	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/frontend"
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

	startHTTPService(config, storage)
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
		return nil, fmt.Errorf("storage backend %q not implemented", engineType)
	}
}

func startHTTPService(config *config.Config, storage storage.Engine) {
	router := mux.NewRouter()
	api := api.NewAPI(config, storage)

	api.RegisterEmployeeRoutes(router)
	api.RegisterPositionRoutes(router)
	api.RegisterScheduleRoutes(router)
	api.RegisterSystemRoutes(router)

	// we put frontend routes last since it serves as a catch-all and
	// mux checks for route matches in the order they are registered
	if config.Frontend {
		frontend := frontend.NewFrontend()
		frontend.RegisterUIRoutes(router)
	}

	var modifiedHandler http.Handler
	if config.Debug {
		modifiedHandler = handlers.LoggingHandler(os.Stdout, router)
	} else {
		modifiedHandler = router
	}

	server := http.Server{
		Addr:         config.URL,
		Handler:      modifiedHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Str("url", config.URL).Msg("starting http service")

	// Run our server in a goroutine and listen for signals
	// that indicate graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("server exited abnormally")

		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// shutdown gracefully with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	os.Exit(0)
}
