package memory

import (
	"encoding/json"

	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/rs/zerolog/log"
)

// GetAllPositions returns an unpaginated list of current positions
func (db *Memory) GetAllPositions() (map[string]*models.Position, error) {
	results := map[string]*models.Position{}

	for id, rawPosition := range db.store[storage.PositionsBucket] {
		var position models.Position

		err := json.Unmarshal(rawPosition, &position)
		if err != nil {
			log.Error().Err(err).Str("id", string(id)).Msg("could not unmarshal database object")
			return nil, err
		}

		results[id] = &position
	}

	return results, nil
}

// GetPosition returns a single position by id
func (db *Memory) GetPosition(id string) (*models.Position, error) {
	var storedPosition models.Position

	rawPosition, ok := db.store[storage.PositionsBucket][id]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := json.Unmarshal(rawPosition, &storedPosition)
	if err != nil {
		return nil, err
	}

	return &storedPosition, err
}

// AddPosition stores a new position
func (db *Memory) AddPosition(id string, position *models.Position) error {
	_, ok := db.store[storage.PositionsBucket][id]
	if ok {
		return utils.ErrEntityExists
	}

	positionRaw, err := json.Marshal(position)
	if err != nil {
		return err
	}

	db.store[storage.PositionsBucket][id] = positionRaw
	return nil
}

// UpdatePosition alters position information
func (db *Memory) UpdatePosition(id string, position *models.Position) error {
	_, ok := db.store[storage.PositionsBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	positionRaw, err := json.Marshal(position)
	if err != nil {
		return err
	}

	db.store[storage.PositionsBucket][id] = positionRaw
	return nil
}

// DeletePosition removes a position from the database
func (db *Memory) DeletePosition(id string) error {

	_, ok := db.store[storage.PositionsBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	db.store[storage.PositionsBucket][id] = nil
	return nil
}
