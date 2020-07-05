package model

import (
	"fmt"
	"time"
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
	// Its used so that we can model a schedule for display purposes.
	//
	// Timetable is a multidimentional map where:
	// The first key is a date in format mm-dd-yyyy.
	// The second key is a listing of all shifts that occur on that day.
	TimeTable map[string][]Shift `json:"time_table"`
	Created   int64              `json:"created"`
	Modified  int64              `json:"modified"`
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
	newSchedule.TimeTable = map[string][]Shift{}

	return newSchedule
}

// ScheduleDay will insert employees into roles they are best suited for per day; this alters the
// schedule datastructure.
func (s *Schedule) ScheduleDay(dateTime time.Time, employees map[string]Employee) error {
	date := dateTime.Format("01-02-2006")

	// Figure out which program we should have on any given day. We might be able to turn
	// this into an enum and drop all of this code someday.
	program := []Shift{}
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
		return fmt.Errorf("could not generate day %q; not a valid weekday", weekday)
	}

	// exit early if we don't actually need to schedule anything for this date
	if program == nil || employees == nil {
		return nil
	}

	employeeSet := newEmployeeSet(employees)

	for _, shift := range program {
		id, err := employeeSet.pop(shift.PositionID)
		if err != nil {
			return err
		}

		s.TimeTable[date] = append(s.TimeTable[date], Shift{
			Start:      shift.Start,
			End:        shift.End,
			PositionID: shift.PositionID,
			EmployeeID: id,
		})
	}

	return nil
}

// employee set represents a set of available employees
type employeeSet map[string]Employee

// newEmployeeSet creates a new set datastructure for managing available employees
func newEmployeeSet(employees map[string]Employee) employeeSet {

	availableEmployees := employeeSet{}
	for employeeID, employee := range employees {
		availableEmployees[employeeID] = employee
	}

	return availableEmployees
}

// pop returns an employee that is eligible to work the given position.
// Will return an error if no employees are found.
// TODO(clintjedwards): return custom error types so that we can eventually return these
// reasons to the user
func (e employeeSet) pop(positionID string) (string, error) {

	for id, employee := range e {
		if _, exists := employee.Positions[positionID]; !exists {
			continue
		}

		delete(e, id)
		return id, nil
	}

	return "", fmt.Errorf("could not find an eligible employee for position %s", positionID)
}
