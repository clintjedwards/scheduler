package model

// EmployeeStatus represents the working state an employee is currently in
type EmployeeStatus string

const (
	// EmployeeActive represents an employee that can be scheduled
	EmployeeActive EmployeeStatus = "active"
	// EmployeeDisabled represents an employee that is not currently working
	EmployeeDisabled EmployeeStatus = "disabled"
)

// Employee represents a schedulable employee
type Employee struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Notes     string         `json:"notes"`
	StartDate string         `json:"start_date"`
	Status    EmployeeStatus `json:"status"`
	// Unavailable represents a mapping of date(format: mm-dd-yyyy) with
	// time range(format: 00:00-24:00) when the employee will not be able to be
	// scheduled
	Unavailable map[string]string `json:"unavailable"`
	// Positions is a set of positions ids that the employee is allowed to work
	Positions map[string]bool `json:"positions"`
	// Preferences are used to weight employees in scheduling. The key of the dictionary
	// is the preferences type and the value can be the current setting.
	// example POSITION => "$somePositionID"
	Preferences map[string]string `json:"preferences"`
	Created     int64             `json:"created"`
	Modified    int64             `json:"modified"`
}

// Position represents an employment position
type Position struct {
	ID            string `json:"id"`
	PrimaryName   string `json:"primary_name"`
	SecondaryName string `json:"secondary_name"`
	Description   string `json:"description"`
}

// Shift is a unit of time in which an employee can be scheduled for
type Shift struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Employee string `json:"employee"`
}

// Program represents a mapping of position to n shifts for certain days.
// Days left empty will have no potential shifts and therefore not be scheduled.
type Program struct {
	Monday    map[string][]Shift `json:"monday"`
	Tuesday   map[string][]Shift `json:"tuesday"`
	Wednesday map[string][]Shift `json:"wednesday"`
	Thursday  map[string][]Shift `json:"thursday"`
	Friday    map[string][]Shift `json:"friday"`
	Saturday  map[string][]Shift `json:"saturday"`
	Sunday    map[string][]Shift `json:"sunday"`
}

// Alloc represents an employee/position tuple that is inserted into timeslots
type Alloc struct {
	EmployeeID string `json:"employee_id"`
	PositionID string `json:"position_id"`
}
