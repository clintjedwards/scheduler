package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	go_proto "github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
)

// GetAllEmployees returns an unpaginated list of current links
func (db *Bolt) GetAllEmployees() (map[string]*proto.Employee, error) {
	results := map[string]*proto.Employee{}

	db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var employee proto.Employee

			err := go_proto.Unmarshal(value, &employee)
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
func (db *Bolt) GetEmployee(id string) (*proto.Employee, error) {

	var storedEmployee proto.Employee

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		employeeRaw := bucket.Get([]byte(id))
		if employeeRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := go_proto.Unmarshal(employeeRaw, &storedEmployee)
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
func (db *Bolt) AddEmployee(id string, employee *proto.Employee) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		employeeRaw, err := go_proto.Marshal(employee)
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

// UpdateEmployee alters employee infromation
func (db *Bolt) UpdateEmployee(id string, employee *proto.Employee) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.EmployeesBucket))

		// First check if key exists
		currentEmployee := bucket.Get([]byte(id))
		if currentEmployee == nil {
			return utils.ErrEntityNotFound
		}

		employeeRaw, err := go_proto.Marshal(employee)
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
