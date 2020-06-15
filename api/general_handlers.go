package api

import (
	"context"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/toolkit/version"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var appVersion = "v0.0.dev_<build_time>_<commit>"

// GetSystemInfo returns system information and health
func (api *API) GetSystemInfo(context context.Context, request *proto.GetSystemInfoRequest) (*proto.GetSystemInfoResponse, error) {

	info, err := version.Parse(appVersion)
	if err != nil {
		log.Error().Err(err).Msg("could not parse version")
		return &proto.GetSystemInfoResponse{}, status.Error(codes.Internal, "could not get system information")
	}

	return &proto.GetSystemInfoResponse{
		BuildTime:       info.Epoch,
		Commit:          info.Hash,
		DebugEnabled:    api.config.Debug,
		FrontendEnabled: api.config.Frontend,
		Semver:          info.Semver,
	}, nil
}
