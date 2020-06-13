package app

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

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

	schedulerAPI := api.NewAPI(config, storage)
	grpcServer := api.CreateGRPCServer(schedulerAPI)

	initCombinedService(config, grpcServer)
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

// initCombinedService starts a long running combined grpc/http (grpc-web compatible) service
// with all proper settings
func initCombinedService(config *config.Config, server *grpc.Server) {
	wrappedGrpc := grpcweb.WrapServer(server)

	router := mux.NewRouter()

	if config.Frontend {
		frontend := frontend.NewFrontend()
		frontend.RegisterUIRoutes(router)
	}

	combinedHandler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.Header.Get("Content-Type"), "application/grpc") || wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}
		router.ServeHTTP(resp, req)
	})

	var modifiedHandler http.Handler
	if config.Debug {
		modifiedHandler = handlers.LoggingHandler(os.Stdout, combinedHandler)
	} else {
		modifiedHandler = combinedHandler
	}

	httpServer := http.Server{
		Addr:         config.URL,
		Handler:      modifiedHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Str("url", config.URL).Msg("starting grpc/http combined service")
	log.Fatal().Err(httpServer.ListenAndServe()).Msg("server exited abnormally")
}
