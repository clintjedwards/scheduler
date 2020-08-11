package model

import "fmt"

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
	StartDate string         `json:"start_date"` //format: mm-dd-yyyy
	Status    EmployeeStatus `json:"status"`
	// Unavailabilities represents time periods that an employee cannot work expressed as cron expressions
	Unavailabilities []string `json:"unavailabilities"`
	// Positions is a set of positions ids that the employee is allowed to work
	Positions map[string]struct{} `json:"positions"`
	// Preferences are used to weight employees in scheduling. The key of the dictionary
	// is the preferences type and the value can be the current setting.
	// example POSITION => "$somePositionID"
	Preferences map[string]string `json:"preferences"`
	Created     int64             `json:"created"`
	Modified    int64             `json:"modified"`
}

// IsValid ensures the bare minimum for an employee is present
func (e *Employee) IsValid() error {
	if e.Name == "" {
		return fmt.Errorf("employee must include name")
	}

	return nil
}

// Position represents an employment position
type Position struct {
	ID            string `json:"id"`
	PrimaryName   string `json:"primary_name"`
	SecondaryName string `json:"secondary_name"`
	Description   string `json:"description"`
	// Metadata is extra data kept about a position that might be useful to consumers
	// of the API. For position specifically, this is a good way to add
	// frontend features like setting custom colors for a position.
	Metadata map[string]string `json:"metadata"`
}

// Shift is a unit of time in which an employee can be scheduled for
type Shift struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	PositionID string `json:"position_id"` // position id
	EmployeeID string `json:"employee_id"` // employee id
}

// Program represents a mapping of position to n shifts for certain days.
// Days left empty will have no potential shifts and therefore not be scheduled.
type Program struct {
	Monday    []Shift `json:"monday"`
	Tuesday   []Shift `json:"tuesday"`
	Wednesday []Shift `json:"wednesday"`
	Thursday  []Shift `json:"thursday"`
	Friday    []Shift `json:"friday"`
	Saturday  []Shift `json:"saturday"`
	Sunday    []Shift `json:"sunday"`
}
