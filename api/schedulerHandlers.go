package api

import (
	"context"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (api *API) GetSchedulerSettings(ctx context.Context, request *proto.GetSchedulerSettingsRequest) (*proto.GetSchedulerSettingsResponse, error) {

	return &proto.GetSchedulerSettingsResponse{}, nil
}

func (api *API) SetSchedulerSettings(ctx context.Context, request *proto.SetSchedulerSettingsRequest) (*proto.SetSchedulerSettingsResponse, error) {

	settings := proto.SchedulerSettings{
		Positions: request.Positions,
	}

	err := api.storage.UpdateSchedulerSettings(settings)
	if err != nil {
		if err != nil {
			if err == utils.ErrEntityNotFound {
				defer zap.S().Fatalw("could not update settings due to missing bucket|key", "error", err)
				return &proto.SetSchedulerSettingsResponse{}, status.Error(codes.NotFound, "could not update settings; settings key not found.")
			}
			zap.S().Fatalw("could not update settings", "error", err)
			return &proto.SetSchedulerSettingsResponse{}, status.Error(codes.Internal, "could not update settings")
		}
	}

	return &proto.SetSchedulerSettingsResponse{}, nil
}
