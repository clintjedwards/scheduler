package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
)

// GetAllPositions returns an unpaginated list of current links
func (db *Bolt) GetAllPositions() (map[string]*models.Position, error) {
	results := map[string]*models.Position{}

	db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.PositionsBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var position models.Position

			err := json.Unmarshal(value, &position)
			if err != nil {
				log.Error().Err(err).Str("id", string(key)).Msg("could not unmarshal database object")
				// We don't return an error here so that we can at least return a partial list
				return nil
			}

			results[string(key)] = &position
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	return results, nil
}

// GetPosition returns a single position by id
func (db *Bolt) GetPosition(id string) (*models.Position, error) {

	var storedPosition models.Position

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.PositionsBucket))

		positionRaw := bucket.Get([]byte(id))
		if positionRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := json.Unmarshal(positionRaw, &storedPosition)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedPosition, nil
}

// AddPosition stores a new position
func (db *Bolt) AddPosition(id string, position *models.Position) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.PositionsBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		positionRaw, err := json.Marshal(position)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), positionRaw)
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

// UpdatePosition alters position infromation
func (db *Bolt) UpdatePosition(id string, position *models.Position) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.PositionsBucket))

		// First check if key exists
		currentPosition := bucket.Get([]byte(id))
		if currentPosition == nil {
			return utils.ErrEntityNotFound
		}

		positionRaw, err := json.Marshal(position)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), positionRaw)
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

// DeletePosition removes a position from the database
func (db *Bolt) DeletePosition(id string) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.PositionsBucket))

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
