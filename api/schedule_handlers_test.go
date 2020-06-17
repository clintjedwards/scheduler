package api

import (
	"fmt"
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
	mockAPI.storage.AddPosition("4", &proto.Position{
		Id:          "1",
		PrimaryName: "baker",
	})
	mockAPI.storage.AddPosition("4", &proto.Position{
		Id:          "2",
		PrimaryName: "porter",
	})
	mockAPI.storage.AddPosition("4", &proto.Position{
		Id:          "3",
		PrimaryName: "second base",
	})

	info.mockAPI = mockAPI

	return
}

func (info *testHarness) TestGenerateSchedule(t *testing.T) {
	t.Run("GenerateSchedule", func(t *testing.T) {
		schedule, _ := info.mockAPI.generateSchedule(time.Now(), proto.GenerateScheduleRequest{
			StartDate:      "01-01-2001",
			Length:         5,
			EmployeeFilter: []string{},
			DayToPositionShiftMapping: &proto.DayToPositionShiftMapping{
				Monday: &proto.PositionShiftMap{
					PositionShiftMap: map[string]*proto.Shifts{
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
				Tuesday: &proto.PositionShiftMap{
					PositionShiftMap: map[string]*proto.Shifts{
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
				Wednesday: &proto.PositionShiftMap{
					PositionShiftMap: map[string]*proto.Shifts{
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
				Thursday: &proto.PositionShiftMap{
					PositionShiftMap: map[string]*proto.Shifts{
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
				Friday: &proto.PositionShiftMap{
					PositionShiftMap: map[string]*proto.Shifts{
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
			},
		})

		for time, positionMap := range schedule.Timetable {
			fmt.Println(time)
			for positionID, shifts := range positionMap.PositionShiftMap {
				fmt.Printf("\tposition %s:\n", positionID)
				for _, shift := range shifts.Shifts {
					fmt.Printf("\t\tshift: %s-%s\n", shift.StartTime, shift.EndTime)
					fmt.Printf("\t\t\temployee: %s\n", shift.Employee)
				}
			}
		}
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
