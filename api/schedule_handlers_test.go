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
		Name: "Clint",
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("2", &proto.Employee{
		Id:   "2",
		Name: "Caroline",
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("3", &proto.Employee{
		Id:   "3",
		Name: "Shane",
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("4", &proto.Employee{
		Id:   "4",
		Name: "Shanaya",
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.UpdateSchedulerSettings(&proto.SchedulerSettings{
		Positions: []*proto.Position{
			{
				Id:          "1",
				PrimaryName: "Baker",
			},
			{
				Id:          "2",
				PrimaryName: "Porter",
			},
			{
				Id:          "3",
				PrimaryName: "Retail",
			},
		},
	})

	info.mockAPI = mockAPI

	return
}

func (info *testHarness) TestGenerateSchedule(t *testing.T) {
	t.Run("GenerateSchedule", func(t *testing.T) {
		info.mockAPI.generateSchedule(time.Now(), 5, proto.GenerateScheduleSettings{
			EmployeeFilter: []string{},
			PositionShiftMap: &proto.PositionShiftMap{
				Monday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Tuesday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Wednesday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Thursday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Friday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Saturday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
				Sunday: map[string]*proto.Shifts{
					"1": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"2": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
					"3": {
						Shifts: []*proto.Shift{
							{
								StartTime: "0800",
								EndTime:   "1300",
							},
						},
					},
				},
			},
		})
	})
}

func (info *testHarness) TestGetEligibleEmployees(t *testing.T) {
	t.Run("GetEligibleEmployees", func(t *testing.T) {
		employees, err := info.mockAPI.getEmployees([]string{"1", "2"})

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
