package api

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"

	"github.com/rs/zerolog/log"
)

// API represents the grpc backend service
type API struct {
	config  *config.Config
	storage storage.Engine
	// we add this so we aren't forced to immediately implement all methods
	// for a valid api server
	proto.UnimplementedSchedulerAPIServer
}

// NewAPI inits a grpc api service
func NewAPI(config *config.Config, storage storage.Engine) *API {

	return &API{
		config:  config,
		storage: storage,
	}
}

// CreateGRPCServer creates a grpc server with all the proper settings; TLS enabled
func CreateGRPCServer(api *API) *grpc.Server {

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	proto.RegisterSchedulerAPIServer(grpcServer, api)

	return grpcServer
}

// InitGRPCService starts a GPRC server
func InitGRPCService(config *config.Config, server *grpc.Server) {

	listen, err := net.Listen("tcp", config.GRPCURL)
	if err != nil {
		log.Panic().Err(err).Msg("could not init tcp listener")
	}

	log.Info().Str("url", config.GRPCURL).Msg("starting grpc service")

	log.Fatal().Err(server.Serve(listen)).Msg("server exited abnormally")
}
