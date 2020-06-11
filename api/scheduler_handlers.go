package api

import (
	"context"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetSchedulerSettings returns global application settings that effect all schedules
func (api *API) GetSchedulerSettings(ctx context.Context, request *proto.GetSchedulerSettingsRequest) (*proto.GetSchedulerSettingsResponse, error) {

	settings, err := api.storage.GetSchedulerSettings()
	if err != nil {
		if err == utils.ErrEntityNotFound {
			defer log.Fatal().Err(err).Msg("could not update settings due to missing bucket|key")
			return &proto.GetSchedulerSettingsResponse{}, status.Error(codes.NotFound, "could not update settings; settings key not found.")
		}
		log.Error().Err(err).Msg("could not get settings ")
		return &proto.GetSchedulerSettingsResponse{}, status.Error(codes.Internal, "could not update settings")
	}

	return &proto.GetSchedulerSettingsResponse{
		Settings: settings,
	}, nil
}

// SetSchedulerSettings sets global application settings that effect all schedules
func (api *API) SetSchedulerSettings(ctx context.Context, request *proto.SetSchedulerSettingsRequest) (*proto.SetSchedulerSettingsResponse, error) {

	for _, position := range request.Positions {
		position.Id = string(utils.GenerateRandString(api.config.IDLength))
	}

	settings := proto.SchedulerSettings{
		Positions: request.Positions,
	}

	err := api.storage.UpdateSchedulerSettings(&settings)
	if err != nil {
		log.Error().Err(err).Msg("could not update settings ")
		return &proto.SetSchedulerSettingsResponse{}, status.Error(codes.Internal, "could not update settings")
	}

	return &proto.SetSchedulerSettingsResponse{
		Settings: &settings,
	}, nil
}
