package api

import (
	"context"
	"fmt"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isMonday() {

}

type shiftEmployeeMap struct {
	shift    proto.Shift
	employee string // employee id
}

// 	startDate, _ := time.Parse("01-02-2006", request.StartDate)
// 		schedule[currentDate.Format("01-02-2006")] = day
// 		currentDate = currentDate.AddDate(0, 0, 1)

// We should then setup the employee pool with default weights attached
// We need to then iterate over the dates we can do this by just setting a counter to 7 and iterating until we hit that
// remembering to increment the date at the end
// As we iterate over the dates we should loop through the positions
// As we loop through the positions we should loop through the shifts
// As we loop through the shifts for each position we should choose employees with the highest weight number
// We can do this by just looping through them all for starters, no need for something like a heap just yet
// who is also eligible for the position
// once we have who is eligible for the position we should place them and add them to a temporary list of people
// already placed on that day
// We should document all of this above the function
func (api *API) generateSchedule(startDate time.Time, length int, settings proto.GenerateScheduleSettings) error {

	employees, err := api.getEligibleEmployees(settings.EmployeeFilter)
	if err != nil {
		return err
	}

	// schedule is represented as a multidimensional map
	// first level: key -> dates : value ->
	// second level: key -> positions : value -> array of shift to employee mappings
	_ = map[time.Time]map[string][]shiftEmployeeMap{}

	// We start from the start date and then iterate through each day until we hit the desired length
	for i := 0; i < length; i++ {
		fmt.Println(startDate)
		startDate = startDate.AddDate(0, 0, 1)
	}

	fmt.Println(employees)
	return nil
}

// getEligibleEmployees returns a map of eligible employees for future scheduling.
// TODO(clintjedwards): This datastructure should be more advanced, something like a heap would work nicely here
// since eventually employees will have a weight
func (api *API) getEligibleEmployees(employeeFilter []string) (map[string]*proto.Employee, error) {

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

func generateDay() {

	// a list of employees that have already been scheduled for this day
	//alreadyScheduled := map[string]bool{}

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
