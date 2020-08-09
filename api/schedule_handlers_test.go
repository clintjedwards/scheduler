package api

import (
	"fmt"
	"testing"

	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/storage/memory"
	"github.com/rs/zerolog/log"
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
	mockAPI.storage.AddEmployee("1", &model.Employee{
		ID:     "1",
		Name:   "Clint",
		Status: model.EmployeeActive,
		Positions: map[string]struct{}{
			"1": {},
			"2": {},
			"3": {},
		},
	})
	mockAPI.storage.AddEmployee("2", &model.Employee{
		ID:     "2",
		Name:   "Caroline",
		Status: model.EmployeeActive,
		Positions: map[string]struct{}{
			"1": {},
			"2": {},
			"3": {},
		},
	})
	mockAPI.storage.AddEmployee("3", &model.Employee{
		ID:     "3",
		Name:   "Shane",
		Status: model.EmployeeActive,
		Positions: map[string]struct{}{
			"1": {},
			"2": {},
			"3": {},
		},
	})
	mockAPI.storage.AddEmployee("4", &model.Employee{
		ID:   "4",
		Name: "Shanaya",
		Positions: map[string]struct{}{
			"1": {},
			"2": {},
			"3": {},
		},
	})
	mockAPI.storage.AddPosition("4", &model.Position{
		ID:          "1",
		PrimaryName: "baker",
	})
	mockAPI.storage.AddPosition("4", &model.Position{
		ID:          "2",
		PrimaryName: "porter",
	})
	mockAPI.storage.AddPosition("4", &model.Position{
		ID:          "3",
		PrimaryName: "second base",
	})

	info.mockAPI = mockAPI

	return
}

func (info *testHarness) TestGenerateSchedule(t *testing.T) {
	t.Run("GenerateSchedule", func(t *testing.T) {
		sch := model.Schedule{
			Start:          "06-19-1990",
			End:            "06-22-1990",
			EmployeeFilter: []string{},
			Program: model.Program{
				Monday: []model.Shift{
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "1",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "2",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "3",
					},
				},
				Tuesday: []model.Shift{
					{
						Start:      "0900",
						End:        "1400",
						PositionID: "1",
					},
					{
						Start:      "0900",
						End:        "1400",
						PositionID: "2",
					},
					{
						Start:      "0900",
						End:        "1400",
						PositionID: "3",
					},
				},
				Wednesday: []model.Shift{
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "1",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "2",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "3",
					},
				},
				Thursday: []model.Shift{
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "1",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "2",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "3",
					},
				},
				Friday: []model.Shift{
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "1",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "2",
					},
					{
						Start:      "0800",
						End:        "1300",
						PositionID: "3",
					},
				},
			},
		}

		newSchedule := model.NewSchedule("test", sch)
		err := info.mockAPI.generateSchedule(newSchedule)
		if err != nil {
			t.Error(err)
		}
		for date, shifts := range newSchedule.TimeTable {
			fmt.Println(date)
			for _, shift := range shifts {
				fmt.Printf("\temployee: %s\n", shift.EmployeeID)
				fmt.Printf("\tposition: %s\n", shift.PositionID)
				fmt.Printf("\tstart: %s\n", shift.Start)
				fmt.Printf("\tend: %s\n\n", shift.End)
			}
		}
	})
}

func TestScheduleHandlers(t *testing.T) {
	info := testHarness{}
	info.setup()

	info.TestGenerateSchedule(t)
}
