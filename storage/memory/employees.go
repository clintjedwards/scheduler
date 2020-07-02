package memory

import (
	"encoding/json"

	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
)

// GetAllEmployees returns an unpaginated list of current employees
func (db *Memory) GetAllEmployees() (map[string]*model.Employee, error) {
	results := map[string]*model.Employee{}

	for id, rawEmployee := range db.store[storage.EmployeesBucket] {
		var employee model.Employee

		err := json.Unmarshal(rawEmployee, &employee)
		if err != nil {
			log.Error().Err(err).Str("id", string(id)).Msg("could not unmarshal database object")
			return nil, err
		}

		results[id] = &employee
	}

	return results, nil
}

// GetEmployee returns a single employee by id
func (db *Memory) GetEmployee(id string) (*model.Employee, error) {
	var storedEmployee model.Employee

	rawEmployee, ok := db.store[storage.EmployeesBucket][id]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := json.Unmarshal(rawEmployee, &storedEmployee)
	if err != nil {
		return nil, err
	}

	return &storedEmployee, err
}

// AddEmployee stores a new employee
func (db *Memory) AddEmployee(id string, employee *model.Employee) error {
	_, ok := db.store[storage.EmployeesBucket][id]
	if ok {
		return utils.ErrEntityExists
	}

	employeeRaw, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	db.store[storage.EmployeesBucket][id] = employeeRaw
	return nil
}

// UpdateEmployee alters employee infromation
func (db *Memory) UpdateEmployee(id string, employee *model.Employee) error {
	_, ok := db.store[storage.EmployeesBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	employeeRaw, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	db.store[storage.EmployeesBucket][id] = employeeRaw
	return nil
}

// DeleteEmployee removes a employee from the database
func (db *Memory) DeleteEmployee(id string) error {

	_, ok := db.store[storage.EmployeesBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	db.store[storage.EmployeesBucket][id] = nil
	return nil
}
