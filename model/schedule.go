package model

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

// Schedule represents a generated timetable mapping of positions => shift => employee
type Schedule struct {
	ID      string  `json:"id"`
	Start   string  `json:"start"`
	End     string  `json:"end"`
	Program Program `json:"program"`
	// Prefereces can be used to weight employees during scheduling.
	// They key is the preference type and the value is the current setting.
	// Example PREFER_MORE_EXPERIENCE => true
	Preferences map[string]string `json:"preferences"`
	// EmployeeFilter can be used to specify employees to use in scheduling.
	// An empty filter will assume all available employees can be scheduled.
	EmployeeFilter []string `json:"employee_filter"`
	// TimeTable is the resulting schedule that has been generated with the other settings.
	// Its used so that we can model a schedule for display purposes. Rather than track
	// things by shift, timetable tracks things by variable increments in the day.
	//
	// Timetable is a multidimentional map where:
	// the first key is a date in format mm-dd-yyyy.
	// the second key represents time slots ranging over the typical time.
	// the second key's value represents employees who are scheduled in that time period
	TimeTable map[string]map[string][]Alloc `json:"time_table"`
	Created   int64                         `json:"created"`
	Modified  int64                         `json:"modified"`
}

// NewSchedule creates a schedule with provided settings.
// We use this so we can populate settings from the user and prepopulate the timetable with time slots.
func NewSchedule(id string, settings Schedule) *Schedule {
	newSchedule := &Schedule{}
	newSchedule.ID = id
	newSchedule.Created = time.Now().Unix()
	newSchedule.Modified = time.Now().Unix()
	newSchedule.EmployeeFilter = settings.EmployeeFilter
	newSchedule.Start = settings.Start
	newSchedule.End = settings.End
	newSchedule.Preferences = settings.Preferences
	newSchedule.Program = settings.Program
	newSchedule.TimeTable = map[string]map[string][]Alloc{}

	return newSchedule
}

// NewDay instantiates a new day and populates it with timeslots
func (s *Schedule) NewDay(date string) map[string][]Alloc {
	timeslots, err := parseTimeSlots("0000", "2330")
	if err != nil {
		log.Fatal().Msg("could not parse time slots for new day")
	}
	s.TimeTable[date] = map[string][]Alloc{}

	for _, time := range timeslots {
		s.TimeTable[date][time] = []Alloc{}
	}

	return s.TimeTable[date]
}

// ScheduleDay will insert employees into roles they are best suited for per day; this alters the
// schedule datastructure.
func (s *Schedule) ScheduleDay(dateTime time.Time, employees map[string]Employee) error {

	date := dateTime.Format("01-02-2006")

	// Figure out which program we should have on any given day. We might be able to turn
	// this into an enum and drop all of this code someday.
	program := map[string][]Shift{}
	switch weekday := dateTime.Weekday(); weekday {
	case time.Monday:
		program = s.Program.Monday
	case time.Tuesday:
		program = s.Program.Tuesday
	case time.Wednesday:
		program = s.Program.Wednesday
	case time.Thursday:
		program = s.Program.Thursday
	case time.Friday:
		program = s.Program.Friday
	case time.Saturday:
		program = s.Program.Saturday
	case time.Sunday:
		program = s.Program.Sunday
	default:
		log.Error().Msgf("could not generate day %q; not a valid weekday", weekday)
	}

	// populate the intended date with all needed timestamps
	day := s.NewDay(date)

	// exit early if we don't actually need to schedule anything for this date
	if program == nil || employees == nil {
		return nil
	}

	employeeSet := newEmployeeSet(employees)

	for positionID, shifts := range program {
		for _, shift := range shifts {
			id, err := employeeSet.nextAvailableEmployee(positionID)
			if err != nil {
				return err
			}

			timeslots, err := parseTimeSlots(shift.Start, shift.End)
			if err != nil {
				return err
			}

			alloc := Alloc{
				EmployeeID: id,
				PositionID: positionID,
			}

			for _, time := range timeslots {
				day[time] = append(day[time], alloc)
			}
		}
	}

	return nil
}

// parseTimeSlots returns all 30 min timeslots within a given time period
// format is of 4 digit 24 hour: 0300
func parseTimeSlots(startTime, endTime string) ([]string, error) {
	// the parsing format according to golang time package
	// https://golang.org/pkg/time/#Parse
	const format = "1504"

	current, err := time.Parse(format, startTime)
	end, err := time.Parse(format, endTime)
	if err != nil {
		return nil, err
	}

	timeSlots := []string{}

	for {
		if current.After(end) {
			break
		}

		timeSlots = append(timeSlots, current.Format(format))
		current = current.Add(time.Minute * 30)
	}

	return timeSlots, nil
}

// employee set represents a set of available employees
type employeeSet map[string]Employee

func newEmployeeSet(employees map[string]Employee) employeeSet {

	availableEmployees := employeeSet{}
	for employeeID, employee := range employees {
		availableEmployees[employeeID] = employee
	}

	return availableEmployees
}

// nextAvailableEmployee returns an employee that is eligible to work the
// given position. Will return an error if no employees are found for a given position
// TODO(clintjedwards): return custom error types so that we can eventually return these
// reasons to the user
func (e employeeSet) nextAvailableEmployee(positionID string) (string, error) {

	for id, employee := range e {
		if _, exists := employee.Positions[positionID]; !exists {
			continue
		}

		delete(e, id)
		return id, nil
	}

	return "", fmt.Errorf("could not find an eligible employee for position %s", positionID)
}
