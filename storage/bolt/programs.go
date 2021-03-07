package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
)

// GetAllPrograms returns an unpaginated list of current links
func (db *Bolt) GetAllPrograms() (map[string]*model.Program, error) {
	results := map[string]*model.Program{}

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.ProgramsBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var program model.Program

			err := json.Unmarshal(value, &program)
			if err != nil {
				log.Error().Err(err).Str("id", string(key)).Msg("could not unmarshal database object")
				// We don't return an error here so that we can at least return a partial list
				return nil
			}

			results[string(key)] = &program
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetProgram returns a single program by id
func (db *Bolt) GetProgram(id string) (*model.Program, error) {

	var storedProgram model.Program

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.ProgramsBucket))

		programRaw := bucket.Get([]byte(id))
		if programRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := json.Unmarshal(programRaw, &storedProgram)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedProgram, nil
}

// AddProgram stores a new program
func (db *Bolt) AddProgram(id string, program *model.Program) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.ProgramsBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		programRaw, err := json.Marshal(program)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), programRaw)
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

// DeleteProgram removes a program from the database
func (db *Bolt) DeleteProgram(id string) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.ProgramsBucket))

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
