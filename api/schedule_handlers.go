package api

import (
	"context"
	"fmt"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type schedule struct {
	employees map[string]*proto.Employee
	// schedule mapping is represented as a multidimensional map
	// first level: key -> dates : value ->
	// second level: key -> positions : value -> array of shift to employee mappings
	mapping map[time.Time]map[string]proto.Shifts
}

// 	startDate, _ := time.Parse("01-02-2006", request.StartDate)
// 		schedule[currentDate.Format("01-02-2006")] = day
// 		currentDate = currentDate.AddDate(0, 0, 1)

// As we iterate over the dates we should loop through the positions
// As we loop through the positions we should loop through the shifts
// As we loop through the shifts for each position we should choose employees with the highest weight number
// We can do this by just looping through them all for starters, no need for something like a heap just yet
// who is also eligible for the position
// once we have who is eligible for the position we should place them and add them to a temporary list of people
// already placed on that day
// We should document all of this above the function
func (api *API) generateSchedule(startDate time.Time, length int, settings proto.GenerateScheduleSettings) error {

	employees, err := api.getEmployees(settings.EmployeeFilter)
	if err != nil {
		return err
	}

	newSchedule := schedule{
		employees: employees,
		mapping:   map[time.Time]map[string]proto.Shifts{},
	}

	// We start from the start date and then iterate through each day until we hit the desired length
	for i := 0; i < length; i++ {
		switch startDate.Weekday() {
		case time.Monday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Monday)
		case time.Tuesday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Tuesday)
		case time.Wednesday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Wednesday)
		case time.Thursday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Thursday)
		case time.Friday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Friday)
		case time.Saturday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Saturday)
		case time.Sunday:
			newSchedule.scheduleDay(startDate, &settings.PositionShiftMap.Sunday)
		default:
			log.Error().Msgf("could not generate day %s; not a valid weekday", startDate.Weekday())
		}
		startDate = startDate.AddDate(0, 0, 1)
	}

	for time, positionMap := range newSchedule.mapping {
		fmt.Println(time.Weekday())
		for positionID, shifts := range positionMap {
			fmt.Printf("\tposition %s:\n", positionID)
			for _, shift := range shifts.Shifts {
				fmt.Printf("\t\tshift: %s-%s\n", shift.StartTime, shift.EndTime)
				fmt.Printf("\t\t\temployee: %s\n", shift.Employee)
			}
		}
	}

	//fmt.Println(newSchedule.mapping)

	return nil
}

// scheduleDay will insert employees into roles they are best suited for
// this alters the given map.
//
func (sch schedule) scheduleDay(date time.Time, positionShiftMap *map[string]*proto.Shifts) {
	testmapping := map[string]proto.Shifts{}

	availableEmployees := map[string]proto.Employee{}
	// a list of already scheduled employees so that they aren't scheduled for multiple positions per day
	for id, employee := range sch.employees {
		availableEmployees[id] = *employee
	}

	for positionID, shifts := range *positionShiftMap {
		for _, shift := range shifts.Shifts {
			for id, employee := range availableEmployees {
				_, exists := employee.Positions[positionID]
				if !exists {
					continue
				}

				shift.Employee = id
				delete(availableEmployees, id)
				break
			}
		}
	}

	for position, shift := range *positionShiftMap {
		testmapping[position] = *shift
	}
	sch.mapping[date] = testmapping
	return
}

// getEmployees returns a map of eligible employees for future scheduling.
// TODO(clintjedwards): This datastructure should be more advanced, something like a heap would work nicely here
// since eventually employees will have a weight. We might have to calculate weights for each
// position ahead of time and then use that here later
func (api *API) getEmployees(employeeFilter []string) (map[string]*proto.Employee, error) {

	eligibleEmployees := map[string]*proto.Employee{}

	employees, err := api.storage.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	// if the employee filter is empty, include all employees
	if len(employeeFilter) == 0 {
		eligibleEmployees = employees
	} else {
		for _, employee := range employeeFilter {
			eligibleEmployees[employee] = employees[employee]
		}
	}

	return eligibleEmployees, nil
}

// GenerateSchedule returns a potential schedule marked as PENDING.
// It is possible to either discard this schedule or update/accept it.
func (api *API) GenerateSchedule(ctx context.Context, request *proto.GenerateScheduleRequest) (*proto.GenerateScheduleResponse, error) {

	// Validate request input
	_, err := time.Parse("01-02-2006", request.StartDate)
	if err != nil {
		return &proto.GenerateScheduleResponse{}, status.Error(codes.FailedPrecondition,
			"could not parse start date; should be in format mm-dd-yyyy")
	}

	//generateSchedule(startDate, *request.Settings)

	return &proto.GenerateScheduleResponse{}, nil
}
