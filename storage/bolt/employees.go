package bolt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/fatih/structs"
	"github.com/rs/zerolog/log"
)

// GetAllEmployees returns an unpaginated list of current links
func (db *Bolt) GetAllEmployees() (map[string]*model.Employee, error) {
	results := map[string]*model.Employee{}

	db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var employee model.Employee

			err := json.Unmarshal(value, &employee)
			if err != nil {
				log.Error().Err(err).Str("id", string(key)).Msg("could not unmarshal database object")
				// We don't return an error here so that we can at least return a partial list
				return nil
			}

			results[string(key)] = &employee
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	return results, nil
}

// GetEmployee returns a single employee by id
func (db *Bolt) GetEmployee(id string) (*model.Employee, error) {

	var storedEmployee model.Employee

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		employeeRaw := bucket.Get([]byte(id))
		if employeeRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := json.Unmarshal(employeeRaw, &storedEmployee)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedEmployee, nil
}

// AddEmployee stores a new employee
func (db *Bolt) AddEmployee(id string, employee *model.Employee) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		employeeRaw, err := json.Marshal(employee)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), employeeRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateEmployee patches employee information
func (db *Bolt) UpdateEmployee(id string, patchEmployee *model.PatchEmployee) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		// First check if key exists
		currentEmployeeRaw := bucket.Get([]byte(id))
		if currentEmployeeRaw == nil {
			return utils.ErrEntityNotFound
		}

		var currentEmployee model.Employee
		err := json.Unmarshal(currentEmployeeRaw, &currentEmployee)
		if err != nil {
			return err
		}

		err = patchStruct(&currentEmployee, patchEmployee)
		if err != nil {
			return err
		}

		currentEmployee.Modified = time.Now().Unix()

		employeeRaw, err := json.Marshal(&currentEmployee)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), employeeRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func patchStruct(target, patch interface{}) error {

	var dst = structs.New(target)
	var fields = structs.New(patch).Fields() // work stack

	for N := len(fields); N > 0; N = len(fields) {
		var srcField = fields[N-1] // pop the top
		fields = fields[:N-1]

		if !srcField.IsExported() {
			continue
		}
		if srcField.IsEmbedded() {
			fields = append(fields, srcField.Fields()...)
			continue
		}
		if srcField.IsZero() {
			continue
		}

		var name = srcField.Name()

		var dstField, ok = dst.FieldOk(name)
		if !ok {
			continue
		}
		var srcValue = reflect.ValueOf(srcField.Value())
		srcValue = reflect.Indirect(srcValue)
		if skind, dkind := srcValue.Kind(), dstField.Kind(); skind != dkind {
			err := fmt.Errorf("field `%v` types mismatch while patching: %v vs %v", name, dkind, skind)
			return err
		}

		err := dstField.Set(srcValue.Interface())
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteEmployee removes a employee from the database
func (db *Bolt) DeleteEmployee(id string) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists == nil {
			return utils.ErrEntityNotFound
		}

		err := bucket.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
