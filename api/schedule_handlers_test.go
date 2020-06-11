package api

import (
	"testing"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage/memory"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

type testHarness struct {
	mockAPI *API
}

func (info *testHarness) setup() {
	memoryStorageEngine, err := memory.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get config in order to start services")
	}

	mockAPI := NewAPI(nil, &memoryStorageEngine)
	mockAPI.storage.AddEmployee("1", &proto.Employee{
		Id:   "1",
		Name: "testEmployee1",
	})
	mockAPI.storage.AddEmployee("2", &proto.Employee{
		Id:   "2",
		Name: "testEmployee2",
	})
	mockAPI.storage.AddEmployee("3", &proto.Employee{
		Id:   "3",
		Name: "testEmployee3",
	})
	mockAPI.storage.AddEmployee("4", &proto.Employee{
		Id:   "4",
		Name: "testEmployee4",
	})

	info.mockAPI = mockAPI

	return
}

func (info *testHarness) TestGenerateSchedule(t *testing.T) {
	t.Run("GenerateSchedule", func(t *testing.T) {
		info.mockAPI.generateSchedule(time.Now(), 5, proto.GenerateScheduleSettings{
			EmployeeFilter: []string{},
		})
	})
}

func (info *testHarness) TestGetEligibleEmployees(t *testing.T) {
	t.Run("GetEligibleEmployees", func(t *testing.T) {
		employees, err := info.mockAPI.getEligibleEmployees([]string{"1", "2"})

		require.NoError(t, err)
		require.NotNil(t, employees)
		require.Contains(t, employees, "1", "")
		require.Contains(t, employees, "2", "")
		require.NotContains(t, employees, "3", "")
		require.NotContains(t, employees, "4", "")
	})
}

func TestScheduleHandlers(t *testing.T) {
	info := testHarness{}
	info.setup()

	info.TestGenerateSchedule(t)
	info.TestGetEligibleEmployees(t)
}
