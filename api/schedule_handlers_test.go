package api

import (
	"fmt"
	"testing"

	"github.com/clintjedwards/scheduler/model"
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
	mockAPI.storage.AddEmployee("1", &model.Employee{
		ID:     "1",
		Name:   "Clint",
		Status: model.EmployeeActive,
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("2", &model.Employee{
		ID:     "2",
		Name:   "Caroline",
		Status: model.EmployeeActive,
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("3", &model.Employee{
		ID:     "3",
		Name:   "Shane",
		Status: model.EmployeeActive,
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},
	})
	mockAPI.storage.AddEmployee("4", &model.Employee{
		ID:   "4",
		Name: "Shanaya",
		Positions: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
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
			End:            "06-19-1990",
			EmployeeFilter: []string{},
			Program: model.Program{
				Monday: map[string][]model.Shift{
					"1": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"2": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"3": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
				},
				Tuesday: map[string][]model.Shift{
					"1": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"2": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"3": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
				},
				Wednesday: map[string][]model.Shift{
					"1": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"2": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"3": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
				},
				Thursday: map[string][]model.Shift{
					"1": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"2": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"3": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
				},
				Friday: map[string][]model.Shift{
					"1": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"2": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
					"3": {
						{
							Start: "0800",
							End:   "1300",
						},
					},
				},
			},
		}

		newSchedule := model.NewSchedule("test", sch)
		err := info.mockAPI.generateSchedule(newSchedule)
		if err != nil {
			t.Error(err)
		}
		for date, times := range newSchedule.TimeTable {
			fmt.Println(date)
			for time, allocs := range times {
				fmt.Printf("\ttime %s:\n", time)
				for _, alloc := range allocs {
					fmt.Printf("\t\t\temployee: %s\n", alloc.EmployeeID)
				}
			}
		}
	})
}

func (info *testHarness) TestGetEligibleEmployees(t *testing.T) {
	t.Run("GetEligibleEmployeesFilter", func(t *testing.T) {
		employees, err := info.mockAPI.getEmployees([]string{"1", "2"})

		require.NoError(t, err)
		require.NotNil(t, employees)
		require.Contains(t, employees, "1", "")
		require.Contains(t, employees, "2", "")
		require.NotContains(t, employees, "3", "")
		require.NotContains(t, employees, "4", "")
	})

	t.Run("GetEligibleEmployeesActive", func(t *testing.T) {
		employees, err := info.mockAPI.getEmployees([]string{})

		require.NoError(t, err)
		require.NotNil(t, employees)
		require.Contains(t, employees, "1", "")
		require.Contains(t, employees, "2", "")
		require.Contains(t, employees, "3", "")
		require.NotContains(t, employees, "4", "")
	})
}

func TestScheduleHandlers(t *testing.T) {
	info := testHarness{}
	info.setup()

	info.TestGenerateSchedule(t)
	info.TestGetEligibleEmployees(t)
}
