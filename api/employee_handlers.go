package api

import (
	"context"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddEmployee adds a new employee to the scheduler service
func (api *API) AddEmployee(ctx context.Context, request *proto.AddEmployeeRequest) (*proto.AddEmployeeResponse, error) {

	//TODO(clintjedwards): validate params

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

// GetEmployee returns a single employee by id
func (api *API) GetEmployee(ctx context.Context, request *proto.GetEmployeeRequest) (*proto.GetEmployeeResponse, error) {

	// Validate user input
	if request.Id == "" {
		return &proto.GetEmployeeResponse{},
			status.Error(codes.FailedPrecondition, "id required")
	}

	employee, err := api.storage.GetEmployee(request.Id)
	if err != nil {
		if err == utils.ErrEntityNotFound {
			return &proto.GetEmployeeResponse{}, status.Error(codes.NotFound, "could not find employee")
		}
		log.Error().Err(err).Msg("could not get employee")
		return &proto.GetEmployeeResponse{}, status.Error(codes.Internal, "could not get employee")
	}

	return &proto.GetEmployeeResponse{
		Employee: employee,
	}, nil
}
