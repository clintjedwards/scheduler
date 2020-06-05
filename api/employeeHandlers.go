package api

import (
	"context"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

// AddEmployee adds a new employee to the scheduler service
func (api *API) AddEmployee(ctx context.Context, request *proto.AddEmployeeRequest) (*proto.AddEmployeeResponse, error) {

	newEmployee := proto.Employee{
		Id:             string(utils.GenerateRandString(api.config.IDLength)),
		Name:           request.Name,
		Notes:          request.Notes,
		StartDate:      request.StartDate,
		Unavailability: request.Unavailability,
		Positions:      request.Positions,
		Preferences:    request.Preferences,
		Created:        time.Now().Unix(),
		Modified:       time.Now().Unix(),
	}

	err := api.storage.AddEmployee(newEmployee.Id, &newEmployee)
	if err != nil {
		if err == utils.ErrEntityExists {
			return &proto.AddEmployeeResponse{}, status.Error(codes.AlreadyExists, "could not add employee; employee exists")
		}
		log.Error().Err(err).Msg("could not add employee")
		return &proto.AddEmployeeResponse{}, status.Error(codes.Internal, "could not add employee")
	}

	return &proto.AddEmployeeResponse{
		Employee: &newEmployee,
	}, nil
}
