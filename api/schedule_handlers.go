package api

import (
	"context"
	"fmt"
	"time"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/mitchellh/copystructure"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (api *API) generateSchedule(settings proto.GenerateScheduleRequest) (proto.Schedule, error) {

	startTime, err := time.Parse("01-02-2006", settings.Start)
	if err != nil {
		return proto.Schedule{}, fmt.Errorf("could not parse start date; should be in format mm-dd-yyyy: %w", err)
	}

	employees, err := api.getEmployees(settings.EmployeeFilter)
	if err != nil {
		return proto.Schedule{}, err
	}

	newSchedule := proto.Schedule{
		Start:          settings.Start,
		Program:        settings.Program,
		Status:         proto.Schedule_PENDING,
		Preferences:    settings.Preferences,
		EmployeeFilter: settings.EmployeeFilter,
		Timetable:      map[string]*proto.PositionShiftMap{},
	}

	//TODO(clintjedwards): settings can be nil which can cause panics, we should validate settings and create
	// non nil values before passing it to other functions
	// We start from the start date and then iterate through each day until we hit the desired length
	for i := 0; i < int(settings.Length); i++ {
		date := startTime.Format("01-02-2006")

		switch startTime.Weekday() {
		case time.Monday:
			if settings.Program.Monday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Monday.PositionShiftMap)
		case time.Tuesday:
			if settings.Program.Tuesday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Tuesday.PositionShiftMap)
		case time.Wednesday:
			if settings.Program.Wednesday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Wednesday.PositionShiftMap)
		case time.Thursday:
			if settings.Program.Thursday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Thursday.PositionShiftMap)
		case time.Friday:
			if settings.Program.Friday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Friday.PositionShiftMap)
		case time.Saturday:
			if settings.Program.Saturday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Saturday.PositionShiftMap)
		case time.Sunday:
			if settings.Program.Sunday == nil {
				break
			}
			newSchedule.Timetable[date] = scheduleDay(employees, settings.Program.Sunday.PositionShiftMap)
		default:
			log.Error().Msgf("could not generate day %s; not a valid weekday", startTime.Weekday())
		}
		startTime = startTime.AddDate(0, 0, 1)
	}

	newSchedule.End = startTime.Format("01-02-2006")
	return newSchedule, nil
}

// scheduleDay will insert employees into roles they are best suited for
// this alters the given map.
func scheduleDay(employees map[string]*proto.Employee, positionShiftMap map[string]*proto.Shifts) *proto.PositionShiftMap {
	availableEmployees := map[string]proto.Employee{}

	rawDay, err := copystructure.Copy(positionShiftMap)
	if err != nil {
		log.Fatal().Msg("could not copy days")
	}
	day := rawDay.(map[string]*proto.Shifts)

	for employeeID, employee := range employees {
		availableEmployees[employeeID] = *employee
	}

	for positionID, shifts := range day {
		for _, shift := range shifts.Shifts {
			for id, employee := range availableEmployees {
				if _, exists := employee.Positions[positionID]; !exists {
					continue
				}

				shift.Employee = id
				delete(availableEmployees, id)
				break
			}
			if shift.Employee == "" {
				// TODO(clintjedwards): Return a collection of errors to the calling function,
				/// so it can be assigned to the day and reported back to the user
				log.Warn().Msgf("could not find eligible employee for shift %s-%s", shift.StartTime, shift.EndTime)
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

	schedule, err := api.generateSchedule(*request)
	if err != nil {
		return &proto.GenerateScheduleResponse{}, status.Error(codes.Internal,
			"could not generate schedule")
	}

	schedule.Id = string(utils.GenerateRandString(api.config.IDLength))
	api.storage.AddSchedule(schedule.Id, &schedule)

	return &proto.GenerateScheduleResponse{
		Schedule: &schedule,
	}, nil
}

// ListSchedules returns an unpaginated list of schedules with order defined
func (api *API) ListSchedules(ctx context.Context, request *proto.ListSchedulesRequest) (*proto.ListSchedulesResponse, error) {

	schedules, err := api.storage.GetAllSchedules()
	if err != nil {
		return &proto.ListSchedulesResponse{}, status.Error(codes.Internal, "failed to retrieve schedules from database")
	}

	return &proto.ListSchedulesResponse{
		Schedules: schedules.Schedules,
		Order:     schedules.Order,
	}, nil
}
