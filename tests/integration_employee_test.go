package tests

import (
	"context"
	"testing"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/stretchr/testify/require"
)

func (info *testHarness) TestAddEmployee(t *testing.T) {
	t.Run("AddEmployee", func(t *testing.T) {

		request := proto.AddEmployeeRequest{
			Name:      "obama",
			Positions: map[string]bool{},
		}

		response, err := info.client.AddEmployee(context.Background(), &request)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotEmpty(t, response)
		require.NotEmpty(t, response.Employee.Positions)
	})
}

func (info *testHarness) TestGetEmployee(t *testing.T) {
	t.Run("GetEmployee", func(t *testing.T) {

		request := proto.AddEmployeeRequest{
			Name:      "michelle",
			Positions: map[string]bool{},
		}

		response, err := info.client.AddEmployee(context.Background(), &request)
		if err != nil {
			require.NoError(t, err)
		}

		expectedResponse := proto.GetEmployeeResponse{
			Employee: &proto.Employee{
				Id:        response.Employee.Id,
				Name:      "michelle",
				Positions: map[string]bool{},
				Created:   response.Employee.Created,
				Modified:  response.Employee.Modified,
			},
		}

		getResponse, err := info.client.GetEmployee(context.Background(), &proto.GetEmployeeRequest{
			Id: response.Employee.Id,
		})
		require.NoError(t, err)
		require.NotNil(t, getResponse)
		require.Equal(t, expectedResponse.Employee, getResponse.Employee)
	})
}
