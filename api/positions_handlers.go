package api

import (
	"context"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//TODO(clintjedwards): validate params of all handlers

// ListPositions returns all positions unpaginated
func (api *API) ListPositions(ctx context.Context, request *proto.ListPositionsRequest) (*proto.ListPositionsResponse, error) {

	positions, err := api.storage.GetAllPositions()
	if err != nil {
		return &proto.ListPositionsResponse{}, status.Error(codes.Internal, "failed to retrieve positions from database")
	}

	return &proto.ListPositionsResponse{
		Positions: positions,
	}, nil
}

// AddPosition adds a new position to the scheduler service
func (api *API) AddPosition(ctx context.Context, request *proto.AddPositionRequest) (*proto.AddPositionResponse, error) {

	newPosition := proto.Position{
		Id:            string(utils.GenerateRandString(api.config.IDLength)),
		PrimaryName:   request.PrimaryName,
		SecondaryName: request.SecondaryName,
		Description:   request.Description,
	}

	err := api.storage.AddPosition(newPosition.Id, &newPosition)
	if err != nil {
		if err == utils.ErrEntityExists {
			return &proto.AddPositionResponse{}, status.Error(codes.AlreadyExists, "could not add position; position exists")
		}
		log.Error().Err(err).Msg("could not add position")
		return &proto.AddPositionResponse{}, status.Error(codes.Internal, "could not add position")
	}

	return &proto.AddPositionResponse{
		Position: &newPosition,
	}, nil
}

// GetPosition returns a single position by id
func (api *API) GetPosition(ctx context.Context, request *proto.GetPositionRequest) (*proto.GetPositionResponse, error) {

	// Validate user input
	if request.Id == "" {
		return &proto.GetPositionResponse{},
			status.Error(codes.FailedPrecondition, "id required")
	}

	position, err := api.storage.GetPosition(request.Id)
	if err != nil {
		if err == utils.ErrEntityNotFound {
			return &proto.GetPositionResponse{}, status.Error(codes.NotFound, "could not find position")
		}
		log.Error().Err(err).Msg("could not get position")
		return &proto.GetPositionResponse{}, status.Error(codes.Internal, "could not get position")
	}

	return &proto.GetPositionResponse{
		Position: position,
	}, nil
}
