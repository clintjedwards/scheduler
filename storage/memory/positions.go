package memory

import (
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/clintjedwards/scheduler/utils"
	go_proto "github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
)

// GetAllPositions returns an unpaginated list of current positions
func (db *Memory) GetAllPositions() (map[string]*proto.Position, error) {
	results := map[string]*proto.Position{}

	for id, rawPosition := range db.store[storage.PositionsBucket] {
		var position proto.Position

		err := go_proto.Unmarshal(rawPosition, &position)
		if err != nil {
			log.Error().Err(err).Str("id", string(id)).Msg("could not unmarshal database object")
			return nil, err
		}

		results[id] = &position
	}

	return results, nil
}

// GetPosition returns a single position by id
func (db *Memory) GetPosition(id string) (*proto.Position, error) {
	var storedPosition proto.Position

	rawPosition, ok := db.store[storage.PositionsBucket][id]
	if !ok {
		return nil, utils.ErrEntityNotFound
	}

	err := go_proto.Unmarshal(rawPosition, &storedPosition)
	if err != nil {
		return nil, err
	}

	return &storedPosition, err
}

// AddPosition stores a new position
func (db *Memory) AddPosition(id string, position *proto.Position) error {
	_, ok := db.store[storage.PositionsBucket][id]
	if ok {
		return utils.ErrEntityExists
	}

	positionRaw, err := go_proto.Marshal(position)
	if err != nil {
		return err
	}

	db.store[storage.PositionsBucket][id] = positionRaw
	return nil
}

// UpdatePosition alters position information
func (db *Memory) UpdatePosition(id string, position *proto.Position) error {
	_, ok := db.store[storage.PositionsBucket][id]
	if !ok {
		return utils.ErrEntityNotFound
	}

	positionRaw, err := go_proto.Marshal(position)
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
