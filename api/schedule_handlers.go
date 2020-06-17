package api

import (
	"context"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type schedule struct {
// 	employees map[string]*proto.Employee
// 	// schedule mapping is represented as a multidimensional map
// 	// first level: key -> dates : value ->
// 	// second level: key -> positions : value -> array of shift to employee mappings
// 	mapping map[time.Time]map[string]*proto.Shifts
// }

// func (sch schedule) string() {
// 	for time, positionMap := range sch.mapping {
// 		fmt.Println(time.Weekday())
// 		for positionID, shifts := range positionMap {
// 			fmt.Printf("\tposition %s:\n", positionID)
// 			for _, shift := range shifts.Shifts {
// 				fmt.Printf("\t\tshift: %s-%s\n", shift.StartTime, shift.EndTime)
// 				fmt.Printf("\t\t\temployee: %s\n", shift.Employee)
// 			}
// 		}
// 	}
// }

func (api *API) generateSchedule(startDate time.Time, settings proto.GenerateScheduleRequest) (proto.Schedule, error) {

	employees, err := api.getEmployees(settings.EmployeeFilter)
	if err != nil {
		return proto.Schedule{}, err
	}

	newSchedule := proto.Schedule{
		StartDate:                 startDate.Format("01-02-2006"),
		DayToPositionShiftMapping: settings.DayToPositionShiftMapping,
		Status:                    proto.Schedule_PENDING,
		Preferences:               settings.Preferences,
		EmployeeFilter:            settings.EmployeeFilter,
		Timetable:                 map[string]*proto.PositionShiftMap{},
	}

	//TODO(clintjedwards): settings can be nil which can cause panics, we should validate settings and create
	// non nil values before passing it to other functions
	// We start from the start date and then iterate through each day until we hit the desired length
	for i := 0; i < int(settings.Length); i++ {
		simpleDate := startDate.Format("01-02-2006")

		switch startDate.Weekday() {
		case time.Monday:
			if settings.DayToPositionShiftMapping.Monday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Monday.PositionShiftMap)
		case time.Tuesday:
			if settings.DayToPositionShiftMapping.Tuesday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Tuesday.PositionShiftMap)
		case time.Wednesday:
			if settings.DayToPositionShiftMapping.Wednesday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Wednesday.PositionShiftMap)
		case time.Thursday:
			if settings.DayToPositionShiftMapping.Thursday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Thursday.PositionShiftMap)
		case time.Friday:
			if settings.DayToPositionShiftMapping.Friday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Friday.PositionShiftMap)
		case time.Saturday:
			if settings.DayToPositionShiftMapping.Saturday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Saturday.PositionShiftMap)
		case time.Sunday:
			if settings.DayToPositionShiftMapping.Sunday == nil {
				continue
			}
			newSchedule.Timetable[simpleDate] = scheduleDay(employees, settings.DayToPositionShiftMapping.Sunday.PositionShiftMap)
		default:
			log.Error().Msgf("could not generate day %s; not a valid weekday", startDate.Weekday())
		}

		startDate = startDate.AddDate(0, 0, 1)
	}

	newSchedule.EndDate = startDate.Format("01-02-2006")
	return newSchedule, nil
}

// scheduleDay will insert employees into roles they are best suited for
// this alters the given map.
func scheduleDay(employees map[string]*proto.Employee, positionShiftMap map[string]*proto.Shifts) *proto.PositionShiftMap {
	day := map[string]*proto.Shifts{}

	for positionID, shifts := range positionShiftMap {
		day[positionID] = shifts
	}

	availableEmployees := map[string]proto.Employee{}
	for id, employee := range employees {
		availableEmployees[id] = *employee
	}

	for positionID, shifts := range day {
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

	return &proto.PositionShiftMap{
		PositionShiftMap: day,
	}
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
	startDate, err := time.Parse("01-02-2006", request.StartDate)
	if err != nil {
		return &proto.GenerateScheduleResponse{}, status.Error(codes.FailedPrecondition,
			"could not parse start date; should be in format mm-dd-yyyy")
	}

	schedule, err := api.generateSchedule(startDate, *request)
	if err != nil {
		return &proto.GenerateScheduleResponse{}, status.Error(codes.Internal,
			"could not generate schedule")
	}

	return &proto.GenerateScheduleResponse{
		Schedule: &schedule,
	}, nil
}
