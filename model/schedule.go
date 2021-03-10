package model

import (
	"fmt"
	"log"
	"time"

	"github.com/clintjedwards/avail"
)

// Schedule represents a generated timetable mapping of positions => shift => employee
type Schedule struct {
	ID      string  `json:"id"`
	Start   string  `json:"start"` // yyyy-mm-dd
	End     string  `json:"end"`   // yyyy-mm-dd
	Program Program `json:"program"`
	// Preferences can be used to weight employees during scheduling.
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
	// The first key is a date in format yyyy-mm-dd.
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
func (sch *Schedule) ScheduleDay(dateTime time.Time, employees map[string]*Employee) error {
	date := dateTime.Format("2006-01-02")

	// Figure out which program we should have on any given day. We might be able to turn
	// this into an enum and drop all of this code someday.
	var program []Shift
	switch weekday := dateTime.Weekday(); weekday {
	case time.Monday:
		program = sch.Program.Monday
	case time.Tuesday:
		program = sch.Program.Tuesday
	case time.Wednesday:
		program = sch.Program.Wednesday
	case time.Thursday:
		program = sch.Program.Thursday
	case time.Friday:
		program = sch.Program.Friday
	case time.Saturday:
		program = sch.Program.Saturday
	case time.Sunday:
		program = sch.Program.Sunday
	default:
		return fmt.Errorf("could not generate day %q; not a valid weekday", weekday)
	}

	// exit early if we don't actually need to schedule anything for this date
	if program == nil || employees == nil {
		return nil
	}

	// filter out employees that aren't eligible for this date
	employees = filterUnavailableEmployees(dateTime, employees)

	employeeSet := newEmployeeSet(employees)
	for _, shift := range program {
		id, err := employeeSet.pop(shift.PositionID)
		if err != nil {
			return err
		}

		sch.TimeTable[date] = append(sch.TimeTable[date], Shift{
			Start:      shift.Start,
			End:        shift.End,
			PositionID: shift.PositionID,
			EmployeeID: id,
		})
	}

	return nil
}

func filterUnavailableEmployees(date time.Time, employees map[string]*Employee) map[string]*Employee {
	filtered := map[string]*Employee{}

	for id, employee := range employees {
		eligible := true

		for _, expression := range employee.Unavailabilities {
			avail, err := avail.New(expression)
			if err != nil {
				log.Printf("(UPDATEME) avail error: %v", err)
				continue
			}

			// if date is within user unavail range user cannot work on that day
			if avail.Able(date) {
				eligible = false
			}
		}

		if eligible {
			filtered[id] = employee
		}
	}

	return filtered
}

// employee set represents a set of available employees
type employeeSet map[string]*Employee

// newEmployeeSet creates a new set datastructure for managing available employees
func newEmployeeSet(employees map[string]*Employee) employeeSet {

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
