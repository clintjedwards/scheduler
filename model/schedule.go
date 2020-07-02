package model

import (
	"time"

	"github.com/rs/zerolog/log"
)

// Alloc represents an employee/position tuple that is inserted into timeslots
type Alloc struct {
	EmployeeID string `json:"employee_id"`
	PositionID string `json:"position_id"`
}

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
	timeslots := []string{"0000", "0030", "0100", "0130", "0200", "0230", "0300", "0330", "0400",
		"0430", "0500", "0530", "0600", "0630", "0700", "0730", "0800", "0830", "0900", "0930",
		"1000", "1030", "1100", "1130", "1200", "1230", "1300", "1330", "1400", "1430", "1500",
		"1530", "1600", "1630", "1700", "1730", "1800", "1830", "1900", "1930", "2000", "2030",
		"2100", "2130", "2200", "2230", "2300", "2330", "2400"}

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
	_ = s.NewDay(date)

	// exit early if we don't actually need to schedule anything for this date
	if program == nil || employees == nil {
		return nil
	}

	// Copy the employee set so it's easy to delete an employee once they have been scheduled
	availableEmployees := map[string]Employee{}
	for employeeID, employee := range employees {
		availableEmployees[employeeID] = employee
	}

	for positionID, shifts := range program {
		for _, shift := range shifts {
			for id, employee := range availableEmployees {
				if _, exists := employee.Positions[positionID]; !exists {
					continue
				}

				// TODO(clintjedwards): for this specific shift we need to iterate through all
				// possible 30 min increments and then populate the day structure with an alloc
				// for each of those increments
				delete(availableEmployees, id)
				break
			}
			if shift.Employee == "" {
				// TODO(clintjedwards): Return a collection of errors to the calling function,
				/// so it can be assigned to the day and reported back to the user
				log.Warn().Msgf("could not find eligible employee for postion: %s; shift %s-%s",
					positionID, shift.Start, shift.End)
			}
		}
	}

	return nil
}

// getTimeslots returns all 30 min timeslots within a given shift period
func getTimeslots(shift Shift) []string {

	return nil
}
